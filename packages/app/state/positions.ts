import { useCallback, useEffect, useMemo } from "react"

import { atom, useAtom, useAtomValue, useSetAtom } from "jotai"

import { useResponse } from "@/lib/hooks/useResponse"
import { api, RouterOutputs } from "@/server/client"

import { useSocket } from "./authentication"
import { atomFamily, atomWithStorage, selectAtom } from "jotai/utils"

type Balances = RouterOutputs["socket"]["balances"]

const CACHE_DURATION = 30 * 1000

interface LastUpdated {
	positions: number
	collectibles: number
}

const activeFetchesAtom = atom<Set<string>>(new Set<string>())
const collectiblesStorageFamily = atomFamily((address: string) =>
	atomWithStorage<Balances["collectibles"]>(`plug.collectibles.${address}`, [])
)
const positionsStorageFamily = atomFamily((address: string) =>
	atomWithStorage<Balances["positions"]>(`plug.positions.${address}`, { tokens: [], protocols: [] })
)
const lastUpdateStorageFamily = atomFamily((address: string) =>
	atomWithStorage<LastUpdated>(`plug.lastUpdated.${address}`, {
		positions: 0,
		collectibles: 0
	})
)

export const collectiblesFamily = atomFamily((address: string) => atom(get => get(collectiblesStorageFamily(address))))
export const positionsFamily = atomFamily((address: string) => atom(get => get(positionsStorageFamily(address))))
export const tokensFamily = atomFamily((address: string) =>
	selectAtom(
		positionsFamily(address),
		positions => positions.tokens,
		(a, b) => {
			if (a.length !== b.length) return false
			return a.every((token, i) => token.symbol === b[i].symbol && token.balance === b[i].balance)
		}
	)
)
export const protocolsFamily = atomFamily((address: string) =>
	selectAtom(
		positionsFamily(address),
		positions => positions.protocols,
		(a, b) => {
			if (a.length !== b.length) return false
			return a.every((protocol, i) => protocol.name === b[i].name && protocol.color === b[i].color)
		}
	)
)

export const lastUpdateFamily = atomFamily((address: string) => atom(get => get(lastUpdateStorageFamily(address))))

export const updateCollectiblesAtom = atomFamily((address: string) =>
	atom(null, (_, set, newCollectibles: Balances["collectibles"]) => {
		set(collectiblesStorageFamily(address), newCollectibles)
		set(lastUpdateStorageFamily(address), prev => ({ ...prev, collectibles: Date.now() }))
	})
)

export const updatePositionsAtom = atomFamily((address: string) =>
	atom(null, (_, set, newPositions: Balances["positions"]) => {
		set(positionsStorageFamily(address), newPositions)
		set(lastUpdateStorageFamily(address), prev => ({ ...prev, positions: Date.now() }))
	})
)

export const useFetchHoldingsForAddress = (address: string, enabled: boolean = true) => {
	const setActiveFetches = useSetAtom(activeFetchesAtom)
	const updateCollectibles = useSetAtom(updateCollectiblesAtom(address))
	const updatePositions = useSetAtom(updatePositionsAtom(address))

	const handleUpdateCollectibles = useCallback(
		(newCollectibles: Balances["collectibles"]) => {
			updateCollectibles(newCollectibles)
		},
		[updateCollectibles]
	)

	const handleUpdatePositions = useCallback(
		(newPositions: Balances["positions"]) => {
			updatePositions(newPositions)
		},
		[updatePositions]
	)

	const {
		data: positionsData,
		isLoading: isLoadingPositions,
		isSuccess: isSuccessPositions,
		isFetching: isFetchingPositions,
		refetch: refetchPositions
	} = useResponse(
		() =>
			api.socket.balances.positions.useQuery(address, {
				enabled: enabled && !!address && address.startsWith("0x"),
				refetchInterval: CACHE_DURATION,
				staleTime: CACHE_DURATION
			}),
		{ onSuccess: handleUpdatePositions, onError: error => console.error('error', error) }
	)

	const {
		data: collectiblesData,
		isLoading: isLoadingCollectibles,
		isSuccess: isSuccessCollectibles,
		isFetching: isFetchingCollectibles,
		refetch: refetchCollectibles
	} = useResponse(
		() =>
			api.socket.balances.collectibles.useQuery(address, {
				enabled: enabled && !!address && address.startsWith("0x"),
				refetchInterval: CACHE_DURATION,
				staleTime: CACHE_DURATION
			}),
		{ onSuccess: handleUpdateCollectibles }
	)

	useEffect(() => {
		if (positionsData) handleUpdatePositions(positionsData)
		if (collectiblesData) handleUpdateCollectibles(collectiblesData)
	}, [positionsData, collectiblesData, handleUpdatePositions, handleUpdateCollectibles])

	useEffect(() => {
		if (!address || !enabled) return

		setActiveFetches((prev: Set<string>) => {
			const newFetches = new Set(prev)
			newFetches.add(address)
			return newFetches
		})

		return () => {
			setActiveFetches((prev: Set<string>) => {
				const newFetches = new Set(prev)
				newFetches.delete(address)
				return newFetches
			})
		}
	}, [address, enabled, setActiveFetches])

	return {
		isLoading: isLoadingPositions || isLoadingCollectibles,
		isSuccess: isSuccessPositions && isSuccessCollectibles,
		isFetching: isFetchingPositions || isFetchingCollectibles,
		refetch: useCallback(async () => {
			const [posResult, collResult] = await Promise.all([refetchPositions(), refetchCollectibles()])
			return {
				success: posResult.isSuccess && collResult.isSuccess
			}
		}, [refetchPositions, refetchCollectibles])
	}
}

export const useInitializeHoldingsFetching = (
	options: { address?: string; enabled: boolean } = { address: "", enabled: true }
) => {
	return useFetchHoldingsForAddress(options.address || "", options.enabled && options.address !== undefined)
}

export const useHoldings = (providedAddress?: string) => {
	const { socket } = useSocket()
	const address = providedAddress || socket?.socketAddress || ""
	const [activeFetches] = useAtom(activeFetchesAtom)

	const collectibles = useAtomValue(collectiblesFamily(address))
	const tokens = useAtomValue(tokensFamily(address))
	const protocols = useAtomValue(protocolsFamily(address))
	const lastUpdate = useAtomValue(lastUpdateFamily(address))

	const isCurrentlyFetching = address.startsWith("0x") && activeFetches.has(address)

	return useMemo(
		() => ({
			address,
			collectibles,
			tokens,
			protocols,
			isLoading: !isCurrentlyFetching && lastUpdate.positions === 0 && lastUpdate.collectibles === 0,
			isSuccess: lastUpdate.positions > 0 && lastUpdate.collectibles > 0,
			lastUpdate
		}),
		[address, collectibles, tokens, protocols, isCurrentlyFetching, lastUpdate]
	)
}
