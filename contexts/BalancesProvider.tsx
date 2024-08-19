import { createContext, FC, PropsWithChildren, useContext } from "react"

import { api, RouterOutputs } from "@/server/client"

import { useSockets } from "./SocketProvider"

const DURATION = 5 * 60 * 1000

export const BalancesContext = createContext<{
	tokens: RouterOutputs["socket"]["balances"]["tokens"] | undefined
	collectibles:
		| RouterOutputs["socket"]["balances"]["collectibles"]
		| undefined
	positions: RouterOutputs["socket"]["balances"]["positions"] | undefined
}>({
	tokens: [],
	collectibles: [],
	positions: {
		tokens: [],
		defi: {}
	}
})

export const BalancesProvider: FC<PropsWithChildren> = ({ children }) => {
	const { address, socket } = useSockets()

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
			enabled: socket?.socketAddress !== undefined
		}
	)

	const { data: positions } = api.socket.balances.positions.useQuery(
		address,
		{
			enabled: address !== undefined,
			onSuccess: data => console.log("positions", data)
		}
	)

	return (
		<BalancesContext.Provider
			value={{
				tokens,
				collectibles,
				positions
			}}
		>
			{children}
		</BalancesContext.Provider>
	)
}

export const useBalances = () => useContext(BalancesContext)
