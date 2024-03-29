import { Network } from "alchemy-sdk"
import { z } from "zod"

export const NetworkSchema = z.union([
	z.literal(Network.ETH_MAINNET),
	z.literal(Network.ETH_GOERLI),
	z.literal(Network.ETH_SEPOLIA),
	z.literal(Network.OPT_MAINNET),
	z.literal(Network.OPT_GOERLI),
	z.literal(Network.OPT_SEPOLIA),
	z.literal(Network.ARB_MAINNET),
	z.literal(Network.ARB_GOERLI),
	z.literal(Network.ARB_SEPOLIA),
	z.literal(Network.MATIC_MAINNET),
	z.literal(Network.MATIC_MUMBAI),
	z.literal(Network.MATIC_AMOY),
	z.literal(Network.ASTAR_MAINNET),
	z.literal(Network.POLYGONZKEVM_MAINNET),
	z.literal(Network.POLYGONZKEVM_TESTNET),
	z.literal(Network.BASE_MAINNET),
	z.literal(Network.BASE_GOERLI),
	z.literal(Network.BASE_SEPOLIA)
])
