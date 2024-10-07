import { FC, PropsWithChildren } from "react"

import { createClient } from "viem"
import { createConfig, http, WagmiProvider } from "wagmi"
import { coinbaseWallet, safe, walletConnect } from "wagmi/connectors"

import { ChainIds, chains } from "@/lib/constants/chains"
import { injectedWithFallback } from "@/lib/functions/wallet/connector"
import { RPCType } from "@/lib/types/chain"

declare module "wagmi" {
	interface Register {
		config: typeof wagmiConfig
	}
}

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

export const wagmiConfig = createConfig({
	chains: [chains[ChainIds.Mainnet]],
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
