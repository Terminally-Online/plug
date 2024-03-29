import type { FC, PropsWithChildren } from "react"
import { createContext, useContext, useMemo } from "react"

import { VaultProvider } from "."
import { useBalance as useNativeBalance } from "wagmi"

import { api } from "@/lib/api"
import { truncateBalance } from "@/lib/blockchain"
import { useBalance } from "@/lib/hooks/useBalance"
import { useDebounce } from "@/lib/hooks/useDebounce"
import { NATIVE_TOKEN_ADDRESS } from "@/lib/tokens"
import { Search } from "@/lib/types/balances"
import { formatNumber } from "@/lib/utils"

import { DomainProvider } from "./DomainProvider"

const INITIAL_SEARCH: Search = {
	query: "",
	isSearching: false,
	asset: undefined
}

export const BalancesContext = createContext<{
	search: Search
	debouncedSearch: Search
	handleSearch: (search: Search) => void
}>({
	search: INITIAL_SEARCH,
	debouncedSearch: INITIAL_SEARCH,
	handleSearch: () => {}
})

export const BalancesProvider: FC<PropsWithChildren> = ({ children }) => {
	const {
		debounce: handleSearch,
		value: search,
		debounced: debouncedSearch
	} = useDebounce({ initial: INITIAL_SEARCH })

	return (
		<VaultProvider>
			<DomainProvider>
				<BalancesContext.Provider
					value={{
						search,
						debouncedSearch,
						handleSearch
					}}
				>
					{children}
				</BalancesContext.Provider>
			</DomainProvider>
		</VaultProvider>
	)
}

export const useBalances = ({
	chainId,
	address,
	direction,
	amount
}: {
	chainId: number
	address: `0x${string}`
	direction?: 1 | -1
	amount?: number
}) => {
	const { search, debouncedSearch, handleSearch } =
		useContext(BalancesContext)

	const { data: balances } = api.account.balances.useQuery(address)

	const nativeAsset = useMemo(() => {
		if (balances === undefined) return undefined

		return balances.find(
			balance =>
				balance?.chain === chainId &&
				balance?.address === NATIVE_TOKEN_ADDRESS
		)
	}, [balances])

	const decimals = useMemo(() => {
		if (search.asset) return search.asset.decimals
		if (nativeAsset) return nativeAsset.decimals
		return 0
	}, [search.asset, nativeAsset])

	const symbol = useMemo(() => {
		if (search.asset) return search.asset.symbol
		if (nativeAsset) return nativeAsset.symbol
		return ""
	}, [search.asset, nativeAsset])

	const balance = useMemo(() => {
		if (balances && search.asset) {
			const asset = search.asset
			const tokenBalance = balances.find(
				balance =>
					balance?.chain === chainId &&
					balance?.address === asset.address
			)

			return tokenBalance?.balance || BigInt(0)
		}
		if (nativeAsset) return nativeAsset.balance

		return BigInt(0)
	}, [search.asset, nativeAsset])

	const amountBigInt = useMemo(() => {
		if (!direction || !decimals) return BigInt(0)

		return (
			BigInt(direction) *
			(amount ? BigInt(amount * 10 ** decimals) : BigInt(0))
		)
	}, [direction, amount, decimals])

	const preBalance = useMemo(
		() => formatNumber(truncateBalance(balance, decimals)),
		[balance, decimals]
	)

	const changedBalance = useMemo(
		() => (balance ?? BigInt(0)) + amountBigInt,
		[balance, amountBigInt]
	)

	const postBalance = useMemo(() => {
		if (changedBalance < BigInt(0)) return "0"

		return formatNumber(truncateBalance(changedBalance, decimals))
	}, [amountBigInt, balance, decimals])

	return {
		chainId,
		address,
		symbol,
		decimals,
		search,
		debouncedSearch,
		preBalance,
		postBalance,
		balances,
		handleSearch
	}
}
