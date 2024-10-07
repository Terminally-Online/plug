import { base, mainnet, optimism } from "viem/chains"

import { Chain, RPCType } from "@/lib/types"

export enum ChainIds {
	Mainnet = mainnet.id
}

export const chains: Record<ChainIds, Chain> = {
	[mainnet.id]: {
		...mainnet,
		alchemyPrefix: "eth-mainnet",
		color: "#393939",
		logo: "/blockchain/ethereum.png",
		rpcUrls: {
			...mainnet.rpcUrls,
			[RPCType.AppOnly]: {
				http: [`https://eth-mainnet.g.alchemy.com/v2/${process.env.ALCHEMY_API_KEY}`],
				webSocket: [`wss://eth-mainnet.g.alchemy.com/v2/${process.env.ALCHEMY_API_KEY}`]
			}
		}
	} as const satisfies Chain,
	[optimism.id]: {
		...optimism,
		alchemyPrefix: "opt-mainnet",
		color: "#FF0420",
		logo: "/blockchain/optimism.png",
		rpcUrls: {
			...optimism.rpcUrls,
			[RPCType.AppOnly]: {
				http: [`https://opt-mainnet.g.alchemy.com/v2/${process.env.ALCHEMY_API_KEY}`],
				webSocket: [`wss://opt-mainnet.g.alchemy.com/v2/${process.env.ALCHEMY_API_KEY}`]
			}
		}
	} as const satisfies Chain,
	[base.id]: {
		...base,
		alchemyPrefix: "base-mainnet",
		color: "#0052FF",
		logo: "/blockchain/base.png",
		rpcUrls: {
			...base.rpcUrls,
			[RPCType.AppOnly]: {
				http: [`https://base-mainnet.g.alchemy.com/v2/${process.env.ALCHEMY_API_KEY}`],
				webSocket: [`wss://base-mainnet.g.alchemy.com/v2/${process.env.ALCHEMY_API_KEY}`]
			}
		}
	} as const satisfies Chain
}

export const chainsArray = Object.values(chains)

export type ChainId = keyof typeof chains
