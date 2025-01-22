import { env } from "@/env"
import { ChainId, chains } from "@/lib"

export const getBlockExplorerUrl = (chainId: keyof typeof chains) => {
	return chains[chainId].blockExplorers.default.url
}

export const getBlockExplorerAddress = (chainId: ChainId, address: string | undefined) => {
	if (!address) return ""
	return `${getBlockExplorerUrl(chainId)}/address/${address}`
}

export const getBlockExplorerTransaction = (chainId: ChainId, tx: string | undefined) => {
	if (!tx) return ""
	return `${getBlockExplorerUrl(chainId)}/tx/${tx}`
}

export const getBlockExplorerBlock = (chainId: ChainId, block: string | undefined) => {
	if (!block) return ""
	return `${getBlockExplorerUrl(chainId)}/block/${block}`
}

// Build map of normalized chain names to chain IDs
const CHAIN_NAME_MAP = Object.entries(chains).reduce((acc, [id, chain]) => {
	const chainId = Number(id) as ChainId
	acc[chain.name.toLowerCase()] = chainId
	acc[chain.alchemyPrefix] = chainId
	return acc
}, {} as Record<string, ChainId>)

// Add additional aliases
const CHAIN_ALIASES: Record<string, ChainId> = {
	"eth": 1,
	"ethereum": 1,
	"mainnet": 1,
	"anvil": 31337,
	"local": 31337,
	"fork": 31337,
	"plug": 31337,
	"op": 10
} as const

const NORMALIZED_CHAIN_MAP = { ...CHAIN_NAME_MAP, ...CHAIN_ALIASES } as const

export const getChainId = (chainName: string | undefined): ChainId => {
	if (!chainName) return env.NEXT_PUBLIC_DEVELOPMENT ? 31337 : 1
	return NORMALIZED_CHAIN_MAP[chainName.toLowerCase()] ?? (env.NEXT_PUBLIC_DEVELOPMENT ? 31337 : 1)
}

export const getChainName = (chainId: ChainId) => {
	const chain = chains[chainId]

	if (chainId === 31337) return "Plug"
	return chain.name
}

export const ASSET_COLORS = [
	"#f87171",
	"#fb923c",
	"#fbbf24",
	"#facc15",
	"#a3e635",
	"#4ade80",
	"#34d399",
	"#2dd4bf",
	"#22d3ee",
	"#38bdf8",
	"#60a5fa",
	"#818cf8",
	"#a78bfa",
	"#c084fc",
	"#e879f9",
	"#f472b6",
	"#fb7185"
]

export const getAssetColor = (symbol: string) => {
	return ASSET_COLORS[symbol.charCodeAt(0) % ASSET_COLORS.length]
}
