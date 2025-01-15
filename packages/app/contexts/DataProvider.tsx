import { Session } from "next-auth"
import { createContext, FC, PropsWithChildren, useContext } from "react"

import { useSetAtom } from "jotai"

import { api } from "@/server/client"
import { actionsAtom } from "@/state/actions"
import { socketModelAtom } from "@/state/authentication"

/**
 * This is the data layer for the application. It is implemented as a context for simplicity
 * of use. In reality, the state is atomic with the use of jotai so that we do not trigger
 * rerenders where is not needed.
 */
interface DataContextType {
	socketQuery: ReturnType<typeof api.socket.get.useQuery>
}

export const DataContext = createContext<DataContextType>({} as DataContextType)

export const DataProvider: FC<PropsWithChildren<{ session: Session | null }>> = ({ session, children }) => {
	const setSocket = useSetAtom(socketModelAtom)
	const setActions = useSetAtom(actionsAtom)

	// Initialize socket query with exposed reference
	const socketQuery = api.socket.get.useQuery(undefined, {
		enabled: session !== null,
		onSuccess: data => setSocket(data)
	})

	api.solver.actions.getSchemas.useQuery(undefined, {
		onSuccess: data => setActions(data)
	})

	return (
		<DataContext.Provider
			value={{
				socketQuery
			}}
		>
			{children}
		</DataContext.Provider>
	)
}

/**
 * Hook to access the data context
 */
export const useData = () => useContext(DataContext)
