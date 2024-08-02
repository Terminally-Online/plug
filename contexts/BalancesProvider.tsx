import { createContext, FC, PropsWithChildren, useContext } from "react"

import { api, RouterOutputs } from "@/server/client"

import { useSockets } from "./SocketProvider"

const DURATION = 10 * 60 * 1000

export const BalancesContext = createContext<{
	tokens: RouterOutputs["socket"]["balances"]["tokens"] | undefined
	collectibles:
		| RouterOutputs["socket"]["balances"]["collectibles"]
		| undefined
}>({
	tokens: [],
	collectibles: {}
})

export const BalancesProvider: FC<PropsWithChildren> = ({ children }) => {
	const { socket } = useSockets()

	const { data: tokens } = api.socket.balances.tokens.useQuery(
		socket?.socketAddress,
		{
			refetchInterval: DURATION,
			enabled: socket?.socketAddress !== undefined
		}
	)
	const { data: collectibles } = api.socket.balances.collectibles.useQuery(
		socket?.socketAddress,
		{
			refetchInterval: DURATION,
			enabled: socket?.socketAddress !== undefined
		}
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
