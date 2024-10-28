import { Session } from "next-auth"
import { createContext, FC, PropsWithChildren } from "react"

import { useSetAtom } from "jotai"

import { api } from "@/server/client"
import { socketModelAtom } from "@/state"

export const DataContext = createContext({
	refetch: () => {}
})

/**
 * This is the data layer for the application. It is implemented as a context for simplicity
 * of use. In reality, the state is atomic with the use of jotai so that we do not trigger
 * rerenders where is not needed.
 */
export const DataProvider: FC<PropsWithChildren<{ session: Session | null }>> = ({ session, children }) => {
	const setSocket = useSetAtom(socketModelAtom)

	const { refetch } = api.socket.get.useQuery(undefined, {
		enabled: session !== null,
		onSuccess: data => setSocket(data)
	})

	return (
		<DataContext.Provider
			value={{
				refetch
			}}
		>
			{children}
		</DataContext.Provider>
	)
}
