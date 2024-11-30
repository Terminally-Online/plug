import { useCallback, useEffect } from "react"

import { useAtom, useAtomValue, useSetAtom } from "jotai"

import { api, RouterOutputs } from "@/server/client"

import { useSocket } from "./authentication"
import { atomFamily, atomWithStorage } from "jotai/utils"

type Balances = RouterOutputs["socket"]["balances"]

const CACHE_DURATION = 5 * 60 * 1000

const collectiblesFamily = atomFamily((address: string) =>
	atomWithStorage<Balances["collectibles"]>(`plug.collectibles.${address}`, [])
)
const positionsFamily = atomFamily((address: string) =>
	atomWithStorage<Balances["positions"]>(`plug.positions.${address}`, { tokens: [], protocols: [] })
)
const lastUpdateCacheFamily = atomFamily((address: string) =>
	atomWithStorage<{ positions: number; collectibles: number }>(`plug.lastUpdated.${address}`, {
		positions: 0,
		collectibles: 0
	})
)

const useFetchHoldings = (address: string) => {
	const setCollectibles = useSetAtom(collectiblesFamily(address))
	const setPositions = useSetAtom(positionsFamily(address))
	const [lastUpdateCache, setLastUpdateCache] = useAtom(lastUpdateCacheFamily(address))

	const updateCollectibles = useCallback(
		(newCollectibles: Balances["collectibles"]) => {
			setCollectibles(newCollectibles)
			setLastUpdateCache(prev => ({ ...prev, collectibles: Date.now() }))
		},
		[setCollectibles, setLastUpdateCache]
	)

	const updatePositions = useCallback(
		(newPositions: Balances["positions"]) => {
			setPositions(newPositions)
			setLastUpdateCache(prev => ({ ...prev, positions: Date.now() }))
		},
		[setPositions, setLastUpdateCache]
	)

	const shouldFetchPositions = useCallback(() => {
		return Date.now() - lastUpdateCache.positions > CACHE_DURATION
	}, [lastUpdateCache.positions])

	const shouldFetchCollectibles = useCallback(() => {
		return Date.now() - lastUpdateCache.collectibles > CACHE_DURATION
	}, [lastUpdateCache.collectibles])

	const {
		data: positionsData,
		isLoading: isLoadingPositions,
		refetch: refetchPositions
	} = api.socket.balances.positions.useQuery(address, {
		enabled: !!address && address.startsWith("0x") && shouldFetchPositions(),
		onSuccess: updatePositions,
		refetchInterval: CACHE_DURATION,
		staleTime: CACHE_DURATION
	})

	const {
		data: collectiblesData,
		isLoading: isLoadingCollectibles,
		refetch: refetchCollectibles
	} = api.socket.balances.collectibles.useQuery(address, {
		enabled: !!address && address.startsWith("0x") && shouldFetchCollectibles(),
		onSuccess: updateCollectibles,
		refetchInterval: CACHE_DURATION,
		staleTime: CACHE_DURATION
	})

	useEffect(() => {
		if (positionsData) updatePositions(positionsData)
		if (collectiblesData) updateCollectibles(collectiblesData)
	}, [positionsData, collectiblesData, updatePositions, updateCollectibles])

	const isLoading = isLoadingPositions || isLoadingCollectibles

	const refetch = useCallback(() => {
		if (shouldFetchPositions()) refetchPositions()
		if (shouldFetchCollectibles()) refetchCollectibles()
	}, [shouldFetchPositions, shouldFetchCollectibles, refetchPositions, refetchCollectibles])

	return { isLoading, refetch }
}

export const useHoldings = (providedAddress?: string) => {
	const { socket } = useSocket()
	const address = providedAddress || socket?.socketAddress || ""

	const { isLoading, refetch: refetchHoldings } = useFetchHoldings(address ?? socket?.socketAddress)

	const collectibles = useAtomValue(collectiblesFamily(address))
	const positions = useAtomValue(positionsFamily(address))

	return {
		address,
		collectibles,
		tokens: positions.tokens,
		protocols: positions.protocols,
		isLoading,
		refetchHoldings
	}
}
