import { FC, PropsWithChildren } from "react"

import { Chain, createClient } from "viem"
import { base, berachain, mainnet, optimism } from "viem/chains"
import { createConfig, WagmiProvider, webSocket } from "wagmi"
import { coinbaseWallet, safe, walletConnect } from "wagmi/connectors"

import { env } from "@/env"
import { chains } from "@/lib/constants/chains"
import { injectedWithFallback } from "@/lib/functions/wallet/connector"
import { RPCType } from "@/lib/types"

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

export const connectedChains = [
	chains[mainnet.id], chains[base.id], chains[optimism.id], chains[berachain.id]
]

export const wagmiConfig = createConfig({
	chains: Object.values(connectedChains) as Chain[] as [Chain, ...Chain[]],
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
			pollingInterval: 200,
			transport: webSocket(chain.rpcUrls[RPCType.AppOnly].webSocket?.[0] ?? "")
		})
	}
})

export const WalletProvider: FC<PropsWithChildren> = ({ children }) => {
	return <WagmiProvider config={wagmiConfig}>{children}</WagmiProvider>
}

export default WalletProvider
