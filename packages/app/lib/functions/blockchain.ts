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

export const getChainId = (chainName: string) => {
	switch (chainName) {
		case "anvil":
			return 31337
		case "ethereum":
			return 1
		case "optimism":
			return 10
		case "base":
			return 8453
		default:
			return 1
	}
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
