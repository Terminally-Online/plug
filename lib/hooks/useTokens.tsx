import { useMemo } from "react"

import { erc20Abi, isAddress } from "viem"
import { useReadContracts } from "wagmi"

import { TOKENS } from "@/lib/tokens"

import { Search } from "../types/balances"
import { levenshteinDistance } from "../utils"

// TODO: If you change chains in the middle of the process we could serve you
//		 the wrong asset address.

export const useTokens = ({
	chainId,
	address,
	query,
	asset
}: {
	chainId: number
	address: `0x${string}`
	query: string
	asset?: Search["asset"]
}) => {
	const typedAddress = (asset?.address ?? "0x") as `0x${string}`

	const { data } = useReadContracts({
		allowFailure: true,
		contracts: [
			{
				address: typedAddress,
				abi: erc20Abi,
				functionName: "name"
			},
			{
				address: typedAddress,
				abi: erc20Abi,
				functionName: "symbol"
			},
			{
				address: typedAddress,
				abi: erc20Abi,
				functionName: "decimals"
			},
			{
				address: typedAddress,
				abi: erc20Abi,
				functionName: "balanceOf",
				args: [address]
			}
		]
	})

	const metadata = useMemo(() => {
		const [name, symbol, decimals, balance] =
			(isAddress(query) && data) || []

		if (query === undefined) return undefined

		// NOTE: Use the logoURI of a similar token in name -- This could end up being a bad decision, but
		//		 it improves the UX for the user. If they have a token that is not in the list, it will
		//		 still show up with a logo that exists on another chain.
		let logoURI = ""
		if (symbol?.result !== undefined) {
			const found = TOKENS.find(token => token.symbol === symbol.result)

			if (found) logoURI = found.logoURI
		}

		if (
			name?.result === undefined ||
			symbol?.result === undefined ||
			decimals?.result === undefined ||
			balance?.result === undefined
		)
			return undefined

		return {
			address: query,
			name: name.result,
			symbol: symbol.result,
			decimals: decimals.result,
			balance: balance.result,
			chainId,
			logoURI
		}
	}, [chainId, query, typedAddress, data])

	const all = useMemo(
		() => TOKENS.filter(token => token.chainId === chainId),
		[chainId]
	)

	const tokens = useMemo(() => {
		const staticTokens = all
			.filter(token => {
				if (query === "") return true

				if (isAddress(query)) {
					const found = token.address === query
					if (found === true && metadata === undefined) return true
					return false
				}

				return (
					token.name.toLowerCase().includes(query.toLowerCase()) ||
					token.symbol.toLowerCase().includes(query.toLowerCase())
				)
			})
			.sort((a, b) => {
				const distanceA = levenshteinDistance(a.symbol, query)
				const distanceB = levenshteinDistance(b.symbol, query)

				return distanceA - distanceB
			})

		if (metadata === undefined) return staticTokens

		return [metadata, ...staticTokens]
	}, [chainId, query, all, metadata])

	return { all, tokens, metadata }
}

export default useTokens
