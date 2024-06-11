import type { FC, PropsWithChildren } from "react"

import { WagmiProvider } from "wagmi"

import { createWeb3Modal } from "@web3modal/wagmi/react"
import { defaultWagmiConfig } from "@web3modal/wagmi/react/config"

import { base, mainnet, optimism } from "viem/chains"

const projectId =
	process.env.NEXT_PUBLIC_WALLETCONNECT_ID ||
	"b17c8bdfe7719b0f3551627ff43a0af1"

const metadata = {
	name: "Plug",
	description: '"IF This, Then That" for Ethereum blockchains and protocols.',
	url: "https://onplug.io",
	icons: ["https://onplug.io/favicon.ico"]
}

const chains = [mainnet, base, optimism] as const

const config = defaultWagmiConfig({
	chains,
	projectId,
	metadata,
	ssr: true,
	enableEIP6963: true,
	enableEmail: true
})

createWeb3Modal({ wagmiConfig: config, projectId })

export const WalletProvider: FC<PropsWithChildren> = ({ children }) => {
	return <WagmiProvider config={config}>{children}</WagmiProvider>
}

export default WalletProvider
