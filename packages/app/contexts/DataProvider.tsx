import { Session } from "next-auth"
import { createContext, FC, PropsWithChildren, useContext } from "react"

import { useAtom, useSetAtom } from "jotai"

import { useResponse } from "@/lib/hooks/useResponse"
import { api } from "@/server/client"
import { actionsAtom } from "@/state/actions"
import { socketModelAtom } from "@/state/authentication"
import { plugsAtom, usePlugSubscriptions } from "@/state/plugs"
import { useInitializeHoldingsFetching } from "@/state/positions"
import { useAccount } from "@/lib/hooks/account/useAccount"

type DataContextType = {}

export const DataContext = createContext<DataContextType>({} as DataContextType)

export const DataProvider: FC<PropsWithChildren<{ session: Session | null }>> = ({ children }) => {
	const { address, isAuthenticated } = useAccount()

	const [socket, setSocket] = useAtom(socketModelAtom)
	const setActions = useSetAtom(actionsAtom)
	const setPlugs = useSetAtom(plugsAtom)

	const enabled = isAuthenticated

	useResponse(
		() =>
			api.socket.get.useQuery(undefined, { enabled }),
		{ onSuccess: socket => setSocket(socket) }
	)

	useResponse(
		() =>
			api.solver.actions.schemas.useQuery(
				{ chainId: 8453 },
				{ enabled }
			),
		{ onSuccess: actions => setActions(actions) }
	)

	useInitializeHoldingsFetching({
		address,
		enabled
	})
	useInitializeHoldingsFetching({
		address: socket.socketAddress,
		enabled
	})

	useResponse(() => api.plugs.all.useQuery({ target: "mine" }, { enabled }), {
		onSuccess: data => setPlugs(prev => [...prev, ...data.filter(d => !prev.some(p => p.id === d.id))])
	})

	usePlugSubscriptions({ enabled })

	return <DataContext.Provider value={{}}>{children}</DataContext.Provider>
}

export const useData = () => useContext(DataContext)
