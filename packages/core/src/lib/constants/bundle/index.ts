import {
	NetworkBase,
	NetworkConfig,
	NetworkReferences,
	Retries
} from '@/src/lib'

export const mainnet = 1
export const optimism = 10
export const polygon = 137
export const base = 8453
export const arbitrum = 42161
export const zora = 7777777
export const degen = 666666666

// * Default retry configuration for the Engine.
export const DEFAULT_NETWORK_RETRIES: Retries = {
	retries: 3,
	delay: 1000
}

// * Resource to acquire public RPC node URLs to use as default:
//   - https://chainlist.org/
export const DEFAULT_NETWORKS: Record<number, NetworkBase> = {
	[mainnet]: {
		key: 'mainnet',
		rpc: 'wss://ethereum.publicnode.com',
		explorer: 'https://api.etherscan.io/api',
		explorerHasApiKey: true
	},
	[optimism]: {
		key: 'optimism',
		rpc: 'wss://optimism.publicnode.com',
		explorer: 'https://api-optimistic.etherscan.io/api',
		explorerHasApiKey: true
	},
	[polygon]: {
		key: 'polygon',
		rpc: 'wss://polygon-bor.publicnode.com',
		explorer: 'https://api.polygonscan.com/api',
		explorerHasApiKey: true
	},
	[base]: {
		key: 'base',
		rpc: 'wss://base.publicnode.com',
		explorer: 'https://api.basescan.org/api',
		explorerHasApiKey: true
	},
	[arbitrum]: {
		key: 'arbitrum',
		rpc: 'wss://arbitrum-one.publicnode.com',
		explorer: 'https://api.arbiscan.io/api',
		explorerHasApiKey: true
	},
	[zora]: {
		key: 'zora',
		rpc: 'https://rpc.zora.energy',
		explorer: 'https://explorer.zora.energy/api',
		explorerHasApiKey: false
	},
	[degen]: {
		key: 'degen',
		rpc: 'https://rpc.degen.tips',
		explorer: 'https://explorer.degen.tips/api',
		explorerHasApiKey: false
	}
}

export const DEFAULT_NETWORK_REFERENCES: NetworkReferences = {
	artifacts: './artifacts',
	// @ts-ignore
	references: {}
}

export const DEFAULT_NETWORK_CONFIG: NetworkConfig = {
	collectors: [],
	executors: [],
	processes: {}
}
