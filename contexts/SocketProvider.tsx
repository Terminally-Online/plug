import { createContext, FC, PropsWithChildren, useContext } from "react"

import { api } from "@/server/client"

import { useSetAtom } from "jotai"

import { socketModelAtom } from "@/state"

export const SocketContext = createContext({})

export const SocketProvider: FC<PropsWithChildren> = ({ children }) => {
	const setSocket = useSetAtom(socketModelAtom)

	api.socket.get.useQuery(undefined, {
		onSettled: data => setSocket(data)
	})

	return <SocketContext.Provider value={{}}>{children}</SocketContext.Provider>
}
