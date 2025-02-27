import { createPublicClient, extractChain, http } from "viem"
import { base, mainnet, optimism } from "viem/chains"

import { env } from "@/env"
import { Chain, RPCType } from "@/lib/types"

const getAppRPCs = (prefix: string) => {
	return {
		[RPCType.AppOnly]: {
			http: [`https://${prefix}.g.alchemy.com/v2/${env.NEXT_PUBLIC_ALCHEMY_KEY}`],
			webSocket: [`wss://${prefix}.g.alchemy.com/v2/${env.NEXT_PUBLIC_ALCHEMY_KEY}`]
		}
	}
}

export const ANVIL_RPC = "127.0.0.1:8545"
export const chains = {
	[mainnet.id]: {
		...mainnet,
		alchemyPrefix: "eth-mainnet",
		color: "#393939",
		logo: "https://cdn.onplug.io/blockchain/ethereum.png",
		rpcUrls: {
			...mainnet.rpcUrls,
			...getAppRPCs("eth-mainnet")
		}
	} as const satisfies Chain,
	[31337]: {
		id: 31337,
		name: "Plug",
		nativeCurrency: { name: "Ether", symbol: "ETH", decimals: 18 },
		blockExplorers: {
			default: {
				name: "Etherscan",
				url: "https://etherscan.io",
				apiUrl: "https://api.etherscan.io/api"
			}
		},
		contracts: {
			ensRegistry: {
				address: "0x00000000000C2E074eC69A0dFb2997BA6C7d2e1e"
			},
			ensUniversalResolver: {
				address: "0xce01f8eee7E479C928F8919abD53E553a36CeF67",
				blockCreated: 19_258_213
			},
			multicall3: {
				address: "0xca11bde05977b3631167028862be2a173976ca11",
				blockCreated: 14_353_601
			}
		},
		alchemyPrefix: "eth-mainnet-forked",
		color: "#FAFF00",
		logo: "https://cdn.onplug.io/protocols/plug.png",
		rpcUrls: {
			default: {
				http: [`http://${ANVIL_RPC}`],
				webSocket: [`ws://${ANVIL_RPC}`]
			},
			[RPCType.AppOnly]: {
				http: [`http://${ANVIL_RPC}`],
				webSocket: [`ws://${ANVIL_RPC}`]
			}
		}
	} as const satisfies Chain,
	[optimism.id]: {
		...optimism,
		alchemyPrefix: "opt-mainnet",
		color: "#FF0420",
		logo: "https://cdn.onplug.io/blockchain/optimism.png",
		rpcUrls: {
			...optimism.rpcUrls,
			...getAppRPCs("opt-mainnet")
		}
	} as const satisfies Chain,
	[base.id]: {
		...base,
		alchemyPrefix: "base-mainnet",
		color: "#0052FF",
		logo: "https://cdn.onplug.io/blockchain/base.png",
		rpcUrls: {
			...base.rpcUrls,
			...getAppRPCs("base-mainnet")
		}
	} as const satisfies Chain
}

export const chainsArray = Object.values(chains)

export type ChainId = keyof typeof chains

export const createClient = (chainId: ChainId) => {
	const chain = extractChain({
		chains: chainsArray,
		id: chainId
	})

	return createPublicClient({
		chain,
		transport: http(chain.rpcUrls[RPCType.AppOnly].http[0])
	})
}
