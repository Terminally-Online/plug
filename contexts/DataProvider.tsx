import { createContext, FC, PropsWithChildren} from "react"

import { api } from "@/server/client"

import { useSetAtom } from "jotai"

import { socketModelAtom } from "@/state"

export const DataContext = createContext({})

/**
 * This is the data layer for the application. It is implemented as a context for simplicity of use.
 * In reality, the state is atomic with the use of jotai so that we do not trigger rerenders where
 * is not needed.
 */
export const DataProvider: FC<PropsWithChildren> = ({ children }) => {
	const setSocket = useSetAtom(socketModelAtom)

	api.socket.get.useQuery(undefined, {
		onSettled: data => setSocket(data)
	})

	return <DataContext.Provider value={{}}>{children}</DataContext.Provider>
}
