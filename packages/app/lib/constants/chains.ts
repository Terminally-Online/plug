import { createPublicClient, extractChain, http } from "viem"
import { anvil, base, mainnet, optimism } from "viem/chains"

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

export const ANVIL_RPC = "http://127.0.0.1:8545"
export const chains = {
	[mainnet.id]: {
		...mainnet,
		alchemyPrefix: "eth-mainnet",
		color: "#393939",
		logo: "/blockchain/ethereum.png",
		rpcUrls: {
			...mainnet.rpcUrls,
			...getAppRPCs("eth-mainnet")
		}
	} as const satisfies Chain,
	[anvil.id]: {
		...anvil,
		alchemyPrefix: "eth-mainnet-forked",
		color: "#FAFF00",
		logo: "/protocols/plug.png",
		rpcUrls: {
			...anvil.rpcUrls,
			[RPCType.AppOnly]: {
				http: [ANVIL_RPC],
				webSocket: [ANVIL_RPC]
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
			...getAppRPCs("opt-mainnet")
		}
	} as const satisfies Chain,
	[base.id]: {
		...base,
		alchemyPrefix: "base-mainnet",
		color: "#0052FF",
		logo: "/blockchain/base.png",
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
