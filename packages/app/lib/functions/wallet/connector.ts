import { createConnector } from "wagmi"
import { injected } from "wagmi/connectors"

export function injectedWithFallback() {
	if (typeof window === "undefined") return injected()

	return createConnector(config => {
		const injectedConnector = injected()(config)

		return {
			...injectedConnector,
			connect(...params) {
				if (!window.ethereum) {
					window.open("https://metamask.io/", "inst_metamask")
				}
				return injectedConnector.connect(...params)
			},
			get icon() {
				return !window.ethereum || window.ethereum?.isMetaMask
					? "/wallets/metamask-icon.svg"
					: "/browser-wallet-light.svg"
			},
			get name() {
				return !window.ethereum
					? "Install MetaMask"
					: window.ethereum?.isMetaMask
						? "MetaMask"
						: "Browser Wallet"
			}
		}
	})
}
