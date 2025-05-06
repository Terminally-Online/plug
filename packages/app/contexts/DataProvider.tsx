import { Session } from "next-auth"
import { createContext, FC, PropsWithChildren, useContext } from "react"

import { useSetAtom } from "jotai"

import { useAccount } from "@/lib/hooks/account/useAccount"
import { useResponse } from "@/lib/hooks/useResponse"
import { api } from "@/server/client"
import { actionsAtom } from "@/state/actions"
import { socketModelAtom } from "@/state/authentication"
import { plugsAtom, usePlugSubscriptions } from "@/state/plugs"

type DataContextType = {}

export const DataContext = createContext<DataContextType>({} as DataContextType)

export const DataProvider: FC<PropsWithChildren<{ session: Session | null }>> = ({ children }) => {
	const { isAuthenticated } = useAccount()

	const setSocket = useSetAtom(socketModelAtom)
	const setActions = useSetAtom(actionsAtom)
	const setPlugs = useSetAtom(plugsAtom)

	useResponse(() => api.socket.get.useQuery(undefined, { enabled: isAuthenticated }), {
		onSuccess: socket => setSocket(socket)
	})

	useResponse(() => api.solver.actions.schemas.useQuery({ chainId: 8453 }, { enabled: isAuthenticated }), {
		onSuccess: actions => setActions(actions)
	})

	useResponse(
		() => api.plugs.all.useQuery({ target: "mine" }, { enabled: isAuthenticated, refetchOnWindowFocus: true }),
		{
			onSuccess: setPlugs
		}
	)

	usePlugSubscriptions({ enabled: isAuthenticated })

	return <DataContext.Provider value={{}}>{children}</DataContext.Provider>
}

export const useData = () => useContext(DataContext)
