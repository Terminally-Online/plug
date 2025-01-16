import { FC, PropsWithChildren } from "react"

import { createClient } from "viem"
import { createConfig, http, WagmiProvider } from "wagmi"
import { coinbaseWallet, safe, walletConnect } from "wagmi/connectors"

import { env } from "@/env"
import { ChainIds, chains } from "@/lib/constants/chains"
import { injectedWithFallback } from "@/lib/functions/wallet/connector"
import { RPCType } from "@/lib/types/chain"

declare module "wagmi" {
	interface Register {
		config: typeof wagmiConfig
	}
}

export const WALLETCONNECT_PARAMS = {
	projectId: env.NEXT_PUBLIC_WALLETCONNECT_ID,
	metadata: {
		name: "Plug",
		description: '"IF This, Then That" for Ethereum blockchains and protocols.',
		url: "https://onplug.io",
		icons: ["https://onplug.io/favicon.ico"]
	},
	showQrModal: false
}

// Use fork network in development, mainnet in production
const defaultChain = env.NEXT_PUBLIC_DEVELOPMENT ? chains[ChainIds.MainnetFork] : chains[ChainIds.Mainnet]

export const wagmiConfig = createConfig({
	chains: [defaultChain],
	connectors: [
		injectedWithFallback(),
		walletConnect(WALLETCONNECT_PARAMS),
		coinbaseWallet({
			appName: WALLETCONNECT_PARAMS.metadata.name,
			appLogoUrl: WALLETCONNECT_PARAMS.metadata.icons[0],
			reloadOnDisconnect: false,
			enableMobileWalletLink: true
		}),
		safe()
	],
	client({ chain }) {
		return createClient({
			chain,
			batch: { multicall: true },
			pollingInterval: 12_000,
			transport: http(chain.rpcUrls[RPCType.AppOnly].http[0])
		})
	}
})

export const WalletProvider: FC<PropsWithChildren> = ({ children }) => {
	return <WagmiProvider config={wagmiConfig}>{children}</WagmiProvider>
}

export default WalletProvider
