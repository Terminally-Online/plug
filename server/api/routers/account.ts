import { Alchemy, Network } from "alchemy-sdk"
import { hexToBigInt } from "viem"
import { z } from "zod"

import { TRPCError } from "@trpc/server"

import { truncateBalance } from "@/lib/blockchain"
import { NATIVE_TOKEN_ADDRESS, TOKENS } from "@/lib/tokens"
import { formatNumber } from "@/lib/utils"
import { createTRPCRouter, protectedProcedure } from "@/server/api/trpc"

import vaultRouter from "./vault"

// TODO: Right now we only show tokens that we have in our official token list,
//       this should be updated so that users can import tokens that they may
//       need, but that we do not have metadata for.

const networks = [
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

const getBalancesForChain = async (address: string, chain: Network) => {
	const chainId = getChainId(chain)
	const chainName = getChainName(chain)

	const alchemy = new Alchemy({
		apiKey: process.env.ALCHEMY_API_KEY,
		network: chain
	})

	const balance = await alchemy.core.getBalance(address)
	const { tokenBalances } = await alchemy.core.getTokenBalances(address)

	const balances = [
		{
			contractAddress: NATIVE_TOKEN_ADDRESS,
			chain: chainId,
			chainName,
			tokenBalance: balance
		},
		...tokenBalances
	]

	return balances
		.map(token => {
			const { contractAddress, tokenBalance } = token

			const staticToken = TOKENS.find(
				t =>
					t.address.toLowerCase() === contractAddress.toLowerCase() &&
					t.chainId === chainId
			)

			if (!staticToken || tokenBalance === null) return undefined

			const balance = hexToBigInt(tokenBalance as `0x${string}`)
			const balanceFormatted = formatNumber(
				truncateBalance(balance, staticToken.decimals)
			)

			return {
				address: contractAddress,
				chain: chainId,
				chainName,
				name: staticToken.name,
				symbol: staticToken.symbol,
				decimals: staticToken.decimals,
				logoURI: staticToken.logoURI,
				balance,
				balanceFormatted
			}
		})
		.filter(token => (token?.balance ?? BigInt(0)) !== BigInt(0))
		.sort((a, b) => (Number(b?.balance) ?? 0) - (Number(a?.balance) ?? 0))
}

const getBalances = async (address: string, chains = networks) => {
	const balances = await Promise.all(
		chains.map(chain => getBalancesForChain(address, chain))
	)

	return balances.flat()
}

export default createTRPCRouter({
	// TODO: Only get balances for the networks that a vault is deployed on.
	balances: protectedProcedure.input(z.string()).query(async ({ input }) => {
		try {
			return await getBalances(input)
		} catch (e) {
			throw new TRPCError({ code: "BAD_REQUEST" })
		}
	}),

	vaults: vaultRouter
})
