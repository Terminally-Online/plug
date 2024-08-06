import { chains, nativeTokenAddress, tokens } from "../constants"

export const getBlockExplorerUrl = (chainId: number) => {
	return chains.find(chain => chain.id === chainId)?.blockExplorers.default
		.url
}

export const getBlockExplorerAddress = (
	chainId: number,
	address: string | undefined
) => {
	if (!address) return ""
	return `${getBlockExplorerUrl(chainId)}/address/${address}`
}

export const getBlockExplorerTransaction = (
	chainId: number,
	tx: string | undefined
) => {
	if (!tx) return ""
	return `${getBlockExplorerUrl(chainId)}/tx/${tx}`
}

export const getBlockExplorerBlock = (
	chainId: number,
	block: string | undefined
) => {
	if (!block) return ""
	return `${getBlockExplorerUrl(chainId)}/block/${block}`
}

export const getChainImage = (chainId: number | string) => {
	if (typeof chainId === "string") return `/blockchain/${chainId}.png`

	switch (chainId) {
		case 1:
			return "/blockchain/ethereum.png"
		case 11155111:
			return "/blockchain/ethereum.png"
		case 10:
			return "/blockchain/optimism.png"
		case 11155420:
			return "/blockchain/optimism.png"
		case 8453:
			return "/blockchain/base.png"
		case 84532:
			return "/blockchain/base.png"
		default:
			return "/blockchain/ethereum.png"
	}
}

export const getNativeAssetImage = (chainId: number) => {
	const nativeAddress = nativeTokenAddress

	const nativeImage = tokens.find(
		token =>
			token.address.toLowerCase() === nativeAddress.toLowerCase() &&
			token.chainId === chainId
	)?.logoURI

	if (nativeImage) return nativeImage

	return `https://assets.smold.app/api/token/1/${nativeAddress}/logo-128.png`
}

export const getAssetColor = (symbol: string) => {
	const colors = [
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

	return colors[symbol.charCodeAt(0) % colors.length]
}
