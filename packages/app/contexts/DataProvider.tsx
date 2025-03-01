import { Session } from "next-auth"
import { createContext, FC, PropsWithChildren, useContext } from "react"

import { useSetAtom } from "jotai"

import { api } from "@/server/client"
import { actionsAtom } from "@/state/actions"
import { socketModelAtom } from "@/state/authentication"
import { useResponse } from "@/lib/hooks/useResponse"

type DataContextType = {}
export const DataContext = createContext<DataContextType>({} as DataContextType)

export const DataProvider: FC<PropsWithChildren<{ session: Session | null }>> = ({ session, children }) => {
	const setSocket = useSetAtom(socketModelAtom)
	const setActions = useSetAtom(actionsAtom)

	useResponse(
		() => api.socket.get.useQuery(undefined, {
			enabled: session !== null,
		}),
		{ onSuccess: socket => setSocket(socket) }
	)

	useResponse(
		() => api.solver.actions.schemas.useQuery(
			{ chainId: 8453 },
			{
				enabled: session !== null,
				refetchInterval: 5 * 60 * 1000
			}
		),
		{ onSuccess: actions => setActions(actions) }
	)

	return (
		<DataContext.Provider value={{}}>
			{children}
		</DataContext.Provider>
	)
}

export const useData = () => useContext(DataContext)
