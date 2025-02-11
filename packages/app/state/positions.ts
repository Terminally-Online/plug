import { useCallback, useEffect } from "react"

import { useAtomValue, useSetAtom } from "jotai"

import { api, RouterOutputs } from "@/server/client"

import { useSocket } from "./authentication"
import { atomFamily, atomWithStorage } from "jotai/utils"

type Balances = RouterOutputs["socket"]["balances"]

const CACHE_DURATION = 30 * 1000

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
	const setLastUpdateCache = useSetAtom(lastUpdateCacheFamily(address))

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

	const {
		data: positionsData,
		isLoading: isLoadingPositions,
		isSuccess: isSuccessPositions,
		isFetching: isFetchingPositions,
		refetch: refetchPositions
	} = api.socket.balances.positions.useQuery(address, {
		enabled: !!address && address.startsWith("0x"),
		onSuccess: updatePositions,
		refetchInterval: CACHE_DURATION,
		staleTime: CACHE_DURATION
	})

	const {
		data: collectiblesData,
		isLoading: isLoadingCollectibles,
		isSuccess: isSuccessCollectibles,
		isFetching: isFetchingCollectibles,
		refetch: refetchCollectibles
	} = api.socket.balances.collectibles.useQuery(address, {
		enabled: !!address && address.startsWith("0x"),
		onSuccess: updateCollectibles,
		refetchInterval: CACHE_DURATION,
		staleTime: CACHE_DURATION
	})

	useEffect(() => {
		if (positionsData) updatePositions(positionsData)
		if (collectiblesData) updateCollectibles(collectiblesData)
	}, [positionsData, collectiblesData, updatePositions, updateCollectibles])

	const isLoading = isLoadingPositions || isLoadingCollectibles || isFetchingPositions || isFetchingCollectibles
	const isSuccess = isSuccessPositions && isSuccessCollectibles

	const refetch = useCallback(() => {
		refetchPositions()
		refetchCollectibles()
	}, [refetchPositions, refetchCollectibles])

	return { isLoading, isSuccess, refetch }
}

export const useHoldings = (providedAddress?: string) => {
	const { socket } = useSocket()
	const address = providedAddress || socket?.socketAddress || ""

	console.log("useHoldings called:", {
		providedAddress,
		socketAddress: socket?.socketAddress,
		resolvedAddress: address
	})

	const { isLoading, isSuccess, refetch: refetchHoldings } = useFetchHoldings(address ?? socket?.socketAddress)

	const collectibles = useAtomValue(collectiblesFamily(address))
	const positions = useAtomValue(positionsFamily(address))
	const lastUpdate = useAtomValue(lastUpdateCacheFamily(address))
	return {
		address,
		collectibles,
		tokens: positions.tokens,
		protocols: positions.protocols,
		isLoading,
		isSuccess,
		refetchHoldings,
		lastUpdate
	}
}
