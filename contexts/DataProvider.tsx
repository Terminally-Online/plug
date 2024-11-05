import { Session } from "next-auth"
import { createContext, FC, PropsWithChildren } from "react"

import { useSetAtom } from "jotai"

import { api } from "@/server/client"
import { actionsAtom, socketModelAtom } from "@/state"

export const DataContext = createContext({})

/**
 * This is the data layer for the application. It is implemented as a context for simplicity
 * of use. In reality, the state is atomic with the use of jotai so that we do not trigger
 * rerenders where is not needed.
 */
export const DataProvider: FC<PropsWithChildren<{ session: Session | null }>> = ({ session, children }) => {
	const setSocket = useSetAtom(socketModelAtom)
	const setActions = useSetAtom(actionsAtom)

	api.socket.get.useQuery(undefined, {
		enabled: session !== null,
		onSuccess: data => setSocket(data)
	})

	api.solver.actions.get.useQuery(undefined, {
		onSuccess: data => setActions(data)
	})

	return <DataContext.Provider value={{}}>{children}</DataContext.Provider>
}
