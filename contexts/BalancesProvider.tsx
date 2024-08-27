import { createContext, FC, PropsWithChildren, useContext } from "react"

import { api, RouterOutputs } from "@/server/client"

import { useSockets } from "./SocketProvider"

const REFETCH_INTERVAL = 5 * 60 * 1000

export const BalancesContext = createContext<{
	collectibles: RouterOutputs["socket"]["balances"]["collectibles"]
	positions: RouterOutputs["socket"]["balances"]["positions"]
}>({
	collectibles: [],
	positions: {
		tokens: [],
		protocols: []
	}
})

export const BalancesProvider: FC<PropsWithChildren> = ({ children }) => {
	const { address, socket } = useSockets()

	const enabled = socket && socket.socketAddress !== undefined && socket.id.startsWith("anonymous") === false

	const { data: collectibles } = api.socket.balances.collectibles.useQuery(socket?.socketAddress, {
		enabled
	})

	const { data: positions } = api.socket.balances.positions.useQuery(address, {
		enabled,
		refetchInterval: REFETCH_INTERVAL
	})

	return (
		<BalancesContext.Provider
			value={{
				collectibles: collectibles ?? [],
				positions: positions ?? {
					tokens: [],
					protocols: []
				}
			}}
		>
			{children}
		</BalancesContext.Provider>
	)
}

export const useBalances = () => useContext(BalancesContext)
