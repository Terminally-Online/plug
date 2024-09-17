import { FC, PropsWithChildren } from "react"

import { createClient } from "viem"
import {
	arbitrum,
	avalanche,
	base,
	blast,
	bsc,
	celo,
	gnosis,
	linea,
	mainnet,
	mantle,
	mode,
	optimism,
	polygon,
	scroll,
	zkSync,
	zora
} from "viem/chains"
import { createConfig, http, WagmiProvider } from "wagmi"
import { coinbaseWallet, injected, safe, walletConnect } from "wagmi/connectors"

import { injectedWithFallback } from "@/lib/functions/wallet/connector"

export const WALLETCONNECT_PARAMS = {
	projectId: process.env.NEXT_PUBLIC_WALLETCONNECT_ID || "b17c8bdfe7719b0f3551627ff43a0af1",
	metadata: {
		name: "Plug",
		description: '"IF This, Then That" for Ethereum blockchains and protocols.',
		url: "https://onplug.io",
		icons: ["https://onplug.io/favicon.ico"]
	},
	showQrModal: false
}

export const wagmiChains = [
	mainnet,
	base,
	optimism,
	arbitrum,
	avalanche,
	blast,
	bsc,
	celo,
	gnosis,
	linea,
	polygon,
	scroll,
	mantle,
	mode,
	zora,
	zkSync
] as const

declare module "wagmi" {
	interface Register {
		config: typeof wagmiConfig
	}
}

export const wagmiConfig = createConfig({
	chains: wagmiChains,
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
			// TODO(#402): Update this to be an appOnly provider.
			transport: http(chain.rpcUrls.default.http[0])
		})
	}
})

export const WalletProvider: FC<PropsWithChildren> = ({ children }) => {
	return <WagmiProvider config={wagmiConfig}>{children}</WagmiProvider>
}

export default WalletProvider
