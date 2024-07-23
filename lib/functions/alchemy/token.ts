import { Alchemy, Network } from "alchemy-sdk"
import { hexToBigInt } from "viem"

import { formatBalance, getPrices, nativeTokenAddress, tokens } from "@/lib"

type Tokens = Array<{
	address: string
	name: string
	symbol: string
	logoURI: string
	balance: bigint
	balanceFormatted: number
	chains: Array<{
		address: string
		chainId: number
		chainName: string
		decimals: number
		balance: bigint
		balanceFormatted: number
	}>
}>

const getNetworkMetadata = (network: Network) => {
	switch (network) {
		case Network.ETH_MAINNET:
			return {
				chainId: 1,
				chainName: "Ethereum"
			}
		case Network.OPT_MAINNET:
			return {
				chainId: 10,
				chainName: "Optimism"
			}
		case Network.BASE_MAINNET:
			return {
				chainId: 8453,
				chainName: "Base"
			}
		default:
			return {
				chainId: 1,
				chainName: "Ethereum"
			}
	}
}

export const getTokensForChain = async (address: string, chain: Network) => {
	const networkMetadata = getNetworkMetadata(chain)

	const alchemy = new Alchemy({
		apiKey: process.env.ALCHEMY_API_KEY,
		network: chain
	})

	const balance = await alchemy.core.getBalance(address)
	const { tokenBalances } = await alchemy.core.getTokenBalances(address)

	const nativeAndERC20Balances = [
		{
			...networkMetadata,
			contractAddress: nativeTokenAddress,
			tokenBalance: balance
		},
		...tokenBalances
	]

	return nativeAndERC20Balances
		.map(token => {
			const staticToken = tokens.find(
				t =>
					t.address.toLowerCase() ===
						token.contractAddress.toLowerCase() &&
					t.chainId === networkMetadata.chainId
			)

			if (!staticToken || token.tokenBalance === null) return undefined

			const balance = hexToBigInt(token.tokenBalance as `0x${string}`)
			const balanceFormatted = formatBalance(
				balance,
				staticToken.decimals
			)

			return {
				...networkMetadata,
				...staticToken,
				address: token.contractAddress,
				balance,
				balanceFormatted
			}
		})
		.filter(token => (token?.balance ?? BigInt(0)) !== BigInt(0))
}

export const getTokens = async (
	address: string,
	networks = [Network.ETH_MAINNET, Network.OPT_MAINNET, Network.BASE_MAINNET]
) => {
	const responses = await Promise.allSettled(
		networks.map(chain => getTokensForChain(address, chain))
	)

	// Get token holders per network.
	const individualTokens = responses
		.flatMap(response => {
			if (response.status === "fulfilled") {
				return response.value
			}
			return []
		})
		.filter(token => token !== undefined)

	// Aggregate canonical tokens into one another.
	const multichainTokens = individualTokens.reduce((acc, token) => {
		const existingToken = acc.find(t => t.symbol === token.symbol)

		if (existingToken) {
			existingToken.balance += token.balance
			existingToken.balanceFormatted += token.balanceFormatted
			existingToken.chains.push(token)
		} else {
			acc.push({
				...token,
				chains: [token]
			})
		}

		return acc
	}, [] as Tokens)

	// Determine the percentage held of each token on each chain.
	const aggregatedTokens = multichainTokens.map(token => {
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

	// Build the URL parameter used to retrieve the prices in one request.
	const tokenKeys = multichainTokens
		.flatMap(token =>
			token.chains.flatMap(
				chain => `${chain.chainName.toLowerCase()}:${chain.address}`
			)
		)
		.join(",")

	// Get the prices for each token on each network.
	const priceData = await getPrices(tokenKeys)
	// Add the price, price change and token value to each token of each network.
	const pricedTokens = aggregatedTokens.map(token => {
		return {
			...token,
			chains: token.chains.map(chain => {
				const chainPriceData =
					priceData[
						`${chain.chainName.toLowerCase()}:${chain.address}`
					]

				return {
					...chain,
					price: chainPriceData?.price ?? 0,
					change: chainPriceData?.change ?? 0,
					value: chain.balanceFormatted * (chainPriceData?.price ?? 0)
				}
			})
		}
	})
	// Add the total value of the token to each token accounting for the value on each network.
	const valuedTokens = pricedTokens.map(token => {
		return {
			...token,
			totalValue: token.chains.reduce(
				(acc, chain) => acc + chain.value,
				0
			)
		}
	})

	// Sort the tokens by the total value.
	const sortedTokens = valuedTokens.sort(
		(a, b) => (Number(b?.totalValue) ?? 0) - (Number(a?.totalValue) ?? 0)
	)

	return sortedTokens
}
