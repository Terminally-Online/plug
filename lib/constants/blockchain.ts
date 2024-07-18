import { formatChainName } from "@/lib/functions/format"

import {
	base,
	baseSepolia,
	mainnet,
	optimism,
	optimismSepolia,
	sepolia
} from "wagmi/chains"

export const mainnets = [mainnet, base, optimism].map(chain => ({
	...chain,
	name: formatChainName(chain.name)
}))
export const testnets = [sepolia, baseSepolia, optimismSepolia].map(chain => ({
	...chain,
	name: formatChainName(chain.name)
}))
export const chains = [...mainnets, ...testnets].map(chain => ({
	...chain,
	name: formatChainName(chain.name)
}))
