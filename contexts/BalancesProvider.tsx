import { createContext, FC, PropsWithChildren, useContext } from "react"

import { api, RouterOutputs } from "@/server/client"

import { useSockets } from "./SocketProvider"

export const BalancesContext = createContext<{
	tokens: RouterOutputs["socket"]["tokens"] | undefined
	collectibles: RouterOutputs["socket"]["collectibles"] | undefined
}>({
	tokens: [],
	collectibles: {}
})

export const BalancesProvider: FC<PropsWithChildren> = ({ children }) => {
	const { socket } = useSockets()

	const { data: tokens } = api.socket.tokens.useQuery(socket?.socketAddress)
	const { data: collectibles } = api.socket.collectibles.useQuery(
		socket?.socketAddress
	)

	return (
		<BalancesContext.Provider
			value={{
				tokens,
				collectibles
			}}
		>
			{children}
		</BalancesContext.Provider>
	)
}

export const useBalances = () => useContext(BalancesContext)
