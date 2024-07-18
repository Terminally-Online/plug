import { Alchemy, Network } from "alchemy-sdk"
import { hexToBigInt } from "viem"

import { formatBalance, nativeTokenAddress, tokens } from "@/lib"

export const networks = [
	Network.ETH_MAINNET,
	Network.OPT_MAINNET,
	Network.BASE_MAINNET
]

const getChainId = (chain: Network) => {
	switch (chain) {
		case Network.ETH_MAINNET:
			return 1
		case Network.OPT_MAINNET:
			return 10
		case Network.BASE_MAINNET:
			return 8453
		default:
			return 1
	}
}

const getChainName = (chain: Network) => {
	switch (chain) {
		case Network.ETH_MAINNET:
			return "Ethereum"
		case Network.OPT_MAINNET:
			return "Optimism"
		case Network.BASE_MAINNET:
			return "Base"
		default:
			return "Ethereum"
	}
}

export const getBalancesForChain = async (address: string, chain: Network) => {
	const chainId = getChainId(chain)
	const chainName = getChainName(chain)

	const alchemy = new Alchemy({
		apiKey: process.env.ALCHEMY_API_KEY,
		network: chain
	})

	const balance = await alchemy.core.getBalance(address)
	const { tokenBalances } = await alchemy.core.getTokenBalances(address)

	const combinedTokens = [
		{
			contractAddress: nativeTokenAddress,
			chain: chainId,
			chainName,
			tokenBalance: balance
		},
		...tokenBalances
	]

	const combinedBalances = combinedTokens.map(token => {
		// Find the token metadata.
		const staticToken = tokens.find(
			t =>
				t.address.toLowerCase() ===
					token.contractAddress.toLowerCase() && t.chainId === chainId
		)

		// If we cannot find the token metadata do not include it.
		if (!staticToken || token.tokenBalance === null) return undefined

		// Determine the balance of the token held on the specific chain.
		const balance = hexToBigInt(token.tokenBalance as `0x${string}`)
		const balanceFormatted = formatBalance(balance, staticToken.decimals)

		return {
			...staticToken,
			address: token.contractAddress,
			chainId,
			chainName,
			balance,
			balanceFormatted
		}
	})

	// Filter out the tokens with a balance of 0.
	return combinedBalances.filter(
		token => (token?.balance ?? BigInt(0)) !== BigInt(0)
	)
}

type ChainBalance = {
	address: string
	chainId: number
	chainName: string
	decimals: number
	balance: bigint
	balanceFormatted: number
}

type CombinedBalances = Array<{
	address: string
	name: string
	symbol: string
	logoURI: string
	balance: bigint
	balanceFormatted: number
	chains: Array<ChainBalance>
}>

export const getBalances = async (address: string, chains = networks) => {
	return (
		(
			await Promise.all(
				chains.map(chain => getBalancesForChain(address, chain))
			)
		)
			.flat()
			// Flatten all the tokens with the same symbol into one.
			.reduce((acc, token) => {
				if (token === undefined) return acc

				const existingToken = acc.find(t => t.symbol === token.symbol)

				const chain = {
					address: token.address,
					chainId: token.chainId,
					chainName: token.chainName,
					decimals: token.decimals,
					balance: token.balance,
					balanceFormatted: token.balanceFormatted
				}

				if (existingToken) {
					existingToken.balance += token.balance
					existingToken.balanceFormatted += token.balanceFormatted

					existingToken.chains.push(chain)
				} else {
					acc.push({
						...token,
						chains: [chain]
					})
				}

				return acc
			}, [] as CombinedBalances)
			// Calculate the percent of holdings across each chain.
			.map(token => {
				const totalBalance = token.chains.reduce(
					(acc, chain) => acc + Number(chain.balance),
					0
				)

				return {
					...token,
					chains: token.chains.map(chain => ({
						...chain,
						percentage: Number.parseFloat(
							Number(
								(Number(chain.balance) / totalBalance) * 100
							).toFixed(2)
						)
					}))
				}
			})
			// Sort by the largest to smallest holdings.
			.sort(
				(a, b) =>
					(Number(b?.balanceFormatted) ?? 0) -
					(Number(a?.balanceFormatted) ?? 0)
			)
	)
}
