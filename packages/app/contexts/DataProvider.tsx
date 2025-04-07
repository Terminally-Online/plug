import { Session } from "next-auth"
import { createContext, FC, PropsWithChildren, useContext } from "react"

import { useSetAtom } from "jotai"

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
	const { address, isConnected, isAuthenticated } = useAccount()

	const setSocket = useSetAtom(socketModelAtom)
	const setActions = useSetAtom(actionsAtom)
	const setPlugs = useSetAtom(plugsAtom)

	const enabled = isConnected || isAuthenticated

	useResponse(
		() =>
			api.socket.get.useQuery(undefined, { enabled }),
		{ onSuccess: socket => setSocket(socket) }
	)

	useResponse(
		() =>
			api.solver.actions.schemas.useQuery(
				{ chainId: 8453 },
				{ enabled: isAuthenticated }
			),
		{ onSuccess: actions => setActions(actions) }
	)

	useResponse(() => api.plugs.all.useQuery({ target: "mine" }, { enabled }), {
		onSuccess: data => setPlugs(prev => [...prev, ...data.filter(d => !prev.some(p => p.id === d.id))])
	})

	usePlugSubscriptions({ enabled: isAuthenticated })

	useInitializeHoldingsFetching({
		address,
		enabled: isAuthenticated
	})
	useInitializeHoldingsFetching({
		address,
		enabled: isAuthenticated
	})

	return <DataContext.Provider value={{}}>{children}</DataContext.Provider>
}

export const useData = () => useContext(DataContext)
