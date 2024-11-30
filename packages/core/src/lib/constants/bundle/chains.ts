import {
	base,
	baseSepolia,
	mainnet,
	optimism,
	optimismSepolia,
	sepolia
} from "viem/chains"

export const chains = [
	mainnet,
	optimism,
	base,
	sepolia,
	optimismSepolia,
	baseSepolia
] as const
