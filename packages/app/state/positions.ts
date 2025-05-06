import { useMemo } from "react"

import { atom, useAtom, useAtomValue, useSetAtom } from "jotai"

import { useResponse } from "@/lib/hooks/useResponse"
import { api, RouterOutputs } from "@/server/client"

import { useSocket } from "./authentication"
import { atomFamily, atomWithStorage } from "jotai/utils"

type Positions = RouterOutputs["service"]["zerion"]["wallet"]["positions"]["data"]
type Collections = RouterOutputs["service"]["zerion"]["wallet"]["nftCollections"]["data"]

const positionsStorageFamily = atomFamily((address: string) =>
	atomWithStorage<Positions>(`plug.positions.${address}`, [])
)
const positionsLoaderFamily = atomFamily((address: string) => atom(
	async (get) => {
		const storedPositions = get(positionsStorageFamily(address))
		return storedPositions
	}
))
const collectionsStorageFamily = atomFamily((address: string) =>
	atomWithStorage<Collections>(`plug.collections.${address}`, [])
)

const countAtom = atom(1)
const asyncAtom = atom(async (get) => get(countAtom) * 2)

export const collectiblesFamily = atomFamily((address: string) => atom(get => get(collectionsStorageFamily(address))))
export const positionsFamily = atomFamily((address: string) => atom(get => get(positionsStorageFamily(address))))

export const updateCollectiblesAtom = atomFamily((address: string) =>
	atom(null, (_, set, collections: Collections) => {
		set(collectionsStorageFamily(address), collections)
	})
)
export const updatePositionsAtom = atomFamily((address: string) =>
	atom(null, (_, set, positions: Positions) => {
		set(positionsStorageFamily(address), positions)
	})
)

export const useFetchHoldingsForAddress = (address: string, enabled: boolean = true) => {
	const updateCollectibles = useSetAtom(updateCollectiblesAtom(address))
	const updatePositions = useSetAtom(updatePositionsAtom(address))

	// const positions = useResponse(
	// 	() =>
	// 		api.socket.balances.positions.useQuery(address, {
	// 			enabled: enabled && !!address && address.startsWith("0x"),
	// 		}),
	// 	{ onSuccess: updatePositions, onError: error => console.error('error', error) }
	// )
	//
	// const collections = useResponse(
	// 	() =>
	// 		api.socket.balances.collectibles.useQuery(address, {
	// 			enabled: enabled && !!address && address.startsWith("0x"),
	// 		}),
	// 	{ onSuccess: updateCollectibles }
	// )
	//
	// return {
	// 	positions, collections
	// }
	return {}
}

export const useInitializeHoldingsFetching = (
	options: { address?: string; enabled: boolean } = { address: "", enabled: true }
) => {
	return useFetchHoldingsForAddress(options.address || "", options.enabled && options.address !== undefined)
}

export const useHoldings = (providedAddress?: string) => {
	const { socket } = useSocket()
	const address = providedAddress || socket?.socketAddress || ""
	// const [activeFetches] = useAtom(activeFetchesAtom)
	//
	// const collectibles = useAtomValue(collectiblesFamily(address))
	// const tokens = useAtomValue(tokensFamily(address))
	// const protocols = useAtomValue(protocolsFamily(address))
	// const lastUpdate = useAtomValue(lastUpdateFamily(address))

	// const isCurrentlyFetching = address.startsWith("0x") && activeFetches.has(address)

	return useMemo(
		() => ({
			// address,
			collectibles: [],
			tokens: [],
			protocols: [],
			// isLoading: !isCurrentlyFetching && lastUpdate.positions === 0 && lastUpdate.collectibles === 0,
			// isSuccess: lastUpdate.positions > 0 && lastUpdate.collectibles > 0,
			// lastUpdate
		}),
		[]
	)
}
