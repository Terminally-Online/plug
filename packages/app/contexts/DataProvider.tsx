import { Session } from "next-auth"
import { createContext, FC, PropsWithChildren, useContext } from "react"

import { useAtom, useSetAtom } from "jotai"

import { useConnect } from "@/lib"
import { useResponse } from "@/lib/hooks/useResponse"
import { api } from "@/server/client"
import { actionsAtom } from "@/state/actions"
import { socketModelAtom } from "@/state/authentication"
import { plugsAtom, usePlugSubscriptions } from "@/state/plugs"
import { useInitializeHoldingsFetching } from "@/state/positions"

type DataContextType = {}

export const DataContext = createContext<DataContextType>({} as DataContextType)

export const DataProvider: FC<PropsWithChildren<{ session: Session | null }>> = ({ children }) => {
	const {
		account: { session }
	} = useConnect()

	const [socket, setSocket] = useAtom(socketModelAtom)
	const setActions = useSetAtom(actionsAtom)
	const setPlugs = useSetAtom(plugsAtom)

	useResponse(
		() =>
			api.socket.get.useQuery(undefined, {
				enabled: session !== null
			}),
		{ onSuccess: socket => setSocket(socket) }
	)

	useResponse(
		() =>
			api.solver.actions.schemas.useQuery(
				{ chainId: 8453 },
				{
					enabled: session !== null
				}
			),
		{ onSuccess: actions => setActions(actions) }
	)

	useResponse(() => api.plugs.all.useQuery({ target: "mine" }, { enabled: session !== null }), {
		onSuccess: data => setPlugs(prev => [...prev, ...data.filter(d => !prev.some(p => p.id === d.id))])
	})

	usePlugSubscriptions({ enabled: session !== null })

	useInitializeHoldingsFetching({
		address: session?.address,
		enabled: session !== null && session?.address.startsWith("0x")
	})
	useInitializeHoldingsFetching({
		address: socket?.socketAddress,
		enabled: session !== null && session?.address.startsWith("0x")
	})

	return <DataContext.Provider value={{}}>{children}</DataContext.Provider>
}

export const useData = () => useContext(DataContext)
