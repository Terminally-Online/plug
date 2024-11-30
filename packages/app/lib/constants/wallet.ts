import { useAtomValue } from "jotai"

import { atomWithStorage } from "jotai/utils"

// TODO: This should be moved to state/connectors.ts

export const CONNECTION = {
	WALLET_CONNECT_CONNECTOR_ID: "walletConnect",
	UNISWAP_WALLET_CONNECT_CONNECTOR_ID: "uniswapWalletConnect",
	INJECTED_CONNECTOR_ID: "injected",
	INJECTED_CONNECTOR_TYPE: "injected",
	COINBASE_SDK_CONNECTOR_ID: "coinbaseWalletSDK",
	COINBASE_RDNS: "com.coinbase.wallet",
	METAMASK_RDNS: "io.metamask",
	UNISWAP_EXTENSION_RDNS: "org.uniswap.app",
	SAFE_CONNECTOR_ID: "safe"
} as const

export const CONNECTOR_ICON_OVERRIDE_MAP: { [id in string]?: string } = {
	[CONNECTION.METAMASK_RDNS]: "/wallets/metamask-icon.svg",
	[CONNECTION.COINBASE_SDK_CONNECTOR_ID]: "/wallets/coinbase-icon.svg",
	[CONNECTION.WALLET_CONNECT_CONNECTOR_ID]: "/wallets/walletconnect-icon.svg"
}

// Used to track which connector was used most recently for UI states.
export const recentConnectorIdAtom = atomWithStorage<string | undefined>("plug.recentConnectorId", undefined)
export function useRecentConnectorId() {
	return useAtomValue(recentConnectorIdAtom)
}
