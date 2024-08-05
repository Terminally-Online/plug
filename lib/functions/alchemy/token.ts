import { Alchemy, Network } from "alchemy-sdk"
import { hexToBigInt } from "viem"

import {
	formatBalance,
	getPriceKey,
	getPrices,
	nativeTokenAddress,
	tokens as staticTokens
} from "@/lib"
import { TokenBalanceModel, TokenPriceModel } from "@/prisma/types"
import { db } from "@/server/db"

const prohibitedNameInclusions = [
	".com",
	".io",
	".org",
	".net",
	".app",
	".gg",
	".xyz",
	".claims",
	".finance",
	".tech",
	".exchange",
	".wallet",
	".capital",
	".fund",
	".capital",
	".su",
	".cloud",
	".events"
]

const prohibitedSymbolInclusions = [
	...prohibitedNameInclusions,
	"claim",
	"airdrop",
	"visit"
]

const MINUTE = 60 * 1000
const TOKEN_BALANCES_CACHE_TIME = 3 * MINUTE

const getNetworkMetadata = (network: Network) => {
	switch (network) {
		case Network.ETH_MAINNET:
			return {
				id: 1,
				name: "ethereum"
			}
		case Network.OPT_MAINNET:
			return {
				id: 10,
				name: "optimism"
			}
		case Network.BASE_MAINNET:
			return {
				id: 8453,
				name: "base"
			}
		default:
			return {
				id: 1,
				name: "ethereum"
			}
	}
}

const getIsExcluded = (token: AlchemyTokenBalance) => {
	const isInvalidBalance = token.balance === 0

	const IsMissingName = token.name === undefined

	const isProhibitedName = prohibitedNameInclusions.some(inclusion =>
		token.name?.toLowerCase().includes(inclusion)
	)

	const isProhibitedSymbol = prohibitedSymbolInclusions.some(inclusion =>
		token.symbol?.toLowerCase().includes(inclusion)
	)

	return !(
		isInvalidBalance ||
		IsMissingName ||
		isProhibitedName ||
		isProhibitedSymbol
	)
}

const getAlchemyTokensForChain = async (
	alchemy: Alchemy,
	address: string,
	chain: {
		id: number
		name: string
	},
	pageKey?: string,
	balances: Array<AlchemyTokenBalance> = []
): Promise<Array<AlchemyTokenBalance>> => {
	if (pageKey === undefined) {
		const nativeBalance = await alchemy.core.getBalance(address)

		if (nativeBalance.toString() !== "0") {
			const balance = formatBalance(nativeBalance.toString(), 18)

			balances = [
				{
					contract: nativeTokenAddress,
					balance,
					name: "Ethereum",
					symbol: "ETH",
					decimals: 18,
					logo: "https://assets.smold.app/api/token/1/0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE/logo-128.png"
				}
			]
		}
	}

	const { tokens, pageKey: nextPageKey } =
		await alchemy.core.getTokensForOwner(
			address,
			pageKey
				? {
						pageKey: pageKey
					}
				: undefined
		)

	balances = [
		...balances,
		...tokens
			.map(token => {
				const staticToken = staticTokens.find(t => {
					const matchingAddress =
						t.address.toLowerCase() ===
						token.contractAddress.toLowerCase()

					const matchingChainId = t.chainId === chain.id

					return matchingAddress && matchingChainId
				})

				const decimals = staticToken?.decimals || token.decimals
				const name = staticToken?.name || token.name
				const logo = token.logo || staticToken?.logoURI
				const symbol = staticToken?.symbol || token.symbol

				return {
					contract: token.contractAddress,
					balance: Number(token.balance),
					decimals: decimals,
					name: name,
					symbol: symbol,
					logo: logo
				}
			})
			.filter(getIsExcluded)
	]

	return nextPageKey
		? getAlchemyTokensForChain(
				alchemy,
				address,
				chain,
				nextPageKey,
				balances
			)
		: balances
}

const getTokensForChain = async (address: string, network: Network) => {
	const chain = getNetworkMetadata(network)

	const cachedTokens = await db.tokenBalanceCache.findUnique({
		where: { socketId_chain: { socketId: address, chain: chain.name } },
		include: { tokens: true }
	})

	const cache =
		cachedTokens &&
		cachedTokens.updatedAt >
			new Date(Date.now() - TOKEN_BALANCES_CACHE_TIME)

	if (cache) return cachedTokens.tokens

	const balances = await getAlchemyTokensForChain(
		new Alchemy({
			apiKey: process.env.ALCHEMY_API_KEY,
			network
		}),
		address,
		chain
	)

	const socket = await db.userSocket.findFirst({
		where: { socketAddress: address }
	})

	if (socket === null) return []

	await db.tokenBalance.deleteMany({
		where: { cacheSocketId: socket.id, cacheChain: chain.name }
	})

	const tokenBalancesCache = await db.tokenBalanceCache.upsert({
		where: { socketId_chain: { socketId: socket.id, chain: chain.name } },
		create: {
			chain: chain.name,
			socketId: socket.id,
			tokens: {
				createMany: {
					data: balances
				}
			}
		},
		update: {
			tokens: {
				createMany: {
					data: balances
				}
			}
		},
		include: { tokens: true }
	})

	return tokenBalancesCache.tokens
}

const aggregateTokensByChain = async (balances: Array<TokenBalanceModel>) => {
	const groupedTokens = balances.reduce(
		(acc, token) => {
			if (
				token === undefined ||
				!token.name ||
				!token.symbol ||
				!token.balance ||
				!token.decimals
			)
				return acc

			const existingToken = acc.find(
				t => t.name === token.name && t.symbol === token.symbol
			)

			const chainToken = {
				chain: token.cacheChain,
				contract: token.contract,
				balance: token.balance,
				decimals: token.decimals
			}

			if (existingToken) {
				existingToken.balance += token.balance
				existingToken.chains.push(chainToken)
			} else {
				acc.push({
					name: token.name,
					symbol: token.symbol,
					balance: token.balance,
					logo: token.logo ?? undefined,
					chains: [chainToken]
				})
			}

			return acc
		},
		[] as Array<{
			name: string
			symbol: string
			balance: number
			logo: string | undefined
			chains: Array<{
				chain: string
				contract: string
				balance: number
				decimals: number
			}>
		}>
	)

	const aggregate = groupedTokens.map(token => {
		return {
			...token,
			chains: token.chains
				.map(chain => {
					const ratio = (chain.balance * 10000) / token.balance
					const decimal = parseInt(ratio.toString())
					const percentage = Number.parseFloat(
						(decimal / 100).toFixed(2)
					)

					return {
						...chain,
						percentage
					}
				})
				.sort((a, b) => b.percentage - a.percentage)
		}
	})

	const prices = (
		await getPrices(
			aggregate.flatMap(token =>
				token.chains.flatMap(chain =>
					getPriceKey(chain.chain, chain.contract)
				)
			)
		)
	).reduce(
		(acc, price) => ({
			...acc,
			[price.id]: price
		}),
		{} as Record<string, TokenPriceModel>
	)

	const tokens = aggregate.map(token => {
		const chains = token.chains.map(chain => {
			const price = prices[getPriceKey(chain.chain, chain.contract)]

			return {
				...chain,
				price: price?.price ?? 0,
				change: price?.change ?? 0,
				value: chain.balance * (price?.price ?? 0)
			}
		})

		const value = chains.reduce((acc, chain) => acc + chain.value, 0)

		return {
			...token,
			chains,
			value
		}
	})

	return tokens.sort((a, b) => Number(b.value) - Number(a.value))
}

export const getTokens = async (
	address: string,
	networks = [Network.ETH_MAINNET, Network.OPT_MAINNET, Network.BASE_MAINNET]
) => {
	const socket = await db.userSocket.findFirst({
		where: { socketAddress: address }
	})

	if (socket === null) return []

	await Promise.all(
		networks.map(async chain => await getTokensForChain(address, chain))
	)

	const balances = await db.tokenBalance.findMany({
		where: {
			cacheSocketId: socket.id,
			balance: { gt: 0 }
		}
	})

	return await aggregateTokensByChain(balances)
}
