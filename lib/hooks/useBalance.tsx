import { useMemo } from "react"

import { isAddress } from "viem"
import { erc20Abi } from "viem"
import { useReadContracts } from "wagmi"

import { TOKENS } from "@/lib/tokens"

export const useBalance = ({
	chainId,
	tokenAddress,
	address
}: {
	chainId: number
	tokenAddress: string | undefined
	address: `0x${string}`
}) => {
	const typedAddress = tokenAddress as `0x${string}`

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
			(isAddress(typedAddress) && data) || []

		console.log(data)

		if (tokenAddress === undefined) return undefined

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
			address: tokenAddress,
			name: name.result,
			symbol: symbol.result,
			decimals: decimals.result,
			balance: balance.result,
			chainId,
			logoURI
		}
	}, [chainId, typedAddress, tokenAddress, data])

	const tokens = useMemo(() => {
		const staticTokens = TOKENS.filter(token => token.chainId === chainId)

		if (metadata === undefined) return staticTokens

		return [metadata, ...staticTokens]
	}, [chainId, metadata])

	return { tokens, metadata }
}

export default useBalance
