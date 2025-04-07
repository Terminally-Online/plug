import { useCallback, useMemo } from "react"

import { Connector } from "wagmi"

import { CONNECTION, isMobileWeb, useRecentConnectorId } from "@/lib"
import { useConnect } from "./useConnect"

type ConnectorID = (typeof CONNECTION)[keyof typeof CONNECTION]

const SHOULD_THROW = { shouldThrow: true } as const

export function getConnectorWithId(
	connectors: readonly Connector[],
	id: ConnectorID,
	options: { shouldThrow: true }
): Connector
export function getConnectorWithId(connectors: readonly Connector[], id: ConnectorID): Connector | undefined
export function getConnectorWithId(
	connectors: readonly Connector[],
	id: ConnectorID,
	options?: { shouldThrow: true }
): Connector | undefined {
	const connector = connectors.find(c => c.id === id)
	if (!connector && options?.shouldThrow) {
		throw new Error(`Expected connector ${id} missing from wagmi context.`)
	}
	return connector
}

/** Returns a wagmi `Connector` with the given id. If `shouldThrow` is passed, an error will be thrown if the connector is not found. */
export function useConnectorWithId(id: ConnectorID, options: { shouldThrow: true }): Connector
export function useConnectorWithId(id: ConnectorID): Connector | undefined
export function useConnectorWithId(id: ConnectorID, options?: { shouldThrow: true }): Connector | undefined {
	const connection = useConnect()
	return useMemo(
		() =>
			options?.shouldThrow
				? getConnectorWithId(connection.connectors, id, options)
				: getConnectorWithId(connection.connectors, id),
		[connection.connectors, id, options]
	)
}

function getInjectedConnectors(connectors: readonly Connector[]) {
	let isCoinbaseWalletBrowser = false
	const injectedConnectors = connectors.filter(c => {
		// Special-case: Ignore coinbase eip6963-injected connector; coinbase connection is handled via the SDK connector.
		if (c.id === CONNECTION.COINBASE_RDNS) {
			if (isMobileWeb) {
				isCoinbaseWalletBrowser = true
			}
			return false
		}

		return c.type === CONNECTION.INJECTED_CONNECTOR_TYPE && c.id !== CONNECTION.INJECTED_CONNECTOR_ID
	})

	// Special-case: Return deprecated window.ethereum connector when no eip6963 injectors are present.
	const fallbackInjector = getConnectorWithId(connectors, CONNECTION.INJECTED_CONNECTOR_ID, { shouldThrow: true })
	if (!injectedConnectors.length && Boolean(window.ethereum)) {
		return { injectedConnectors: [fallbackInjector], isCoinbaseWalletBrowser }
	}

	return { injectedConnectors, isCoinbaseWalletBrowser }
}

type InjectableConnector = Connector & { isInjected?: boolean }
export function useOrderedConnections(excludeWalletConnectConnections = false): InjectableConnector[] {
	const { connectors }  = useConnect()
	const recentConnectorId = useRecentConnectorId()

	const sortByRecent = useCallback(
		(a: Connector, b: Connector) => {
			if (a.id === recentConnectorId) {
				return -1
			} else if (b.id === recentConnectorId) {
				return 1
			} else {
				return 0
			}
		},
		[recentConnectorId]
	)

	return useMemo(() => {
		const { injectedConnectors: injectedConnectorsBase, isCoinbaseWalletBrowser } = getInjectedConnectors(
			connectors
		)
		const injectedConnectors = injectedConnectorsBase.map(c => ({ ...c, isInjected: true }))

		const coinbaseSdkConnector = getConnectorWithId(
			connectors,
			CONNECTION.COINBASE_SDK_CONNECTOR_ID,
			SHOULD_THROW
		)
		const walletConnectConnector = getConnectorWithId(
			connectors,
			CONNECTION.WALLET_CONNECT_CONNECTOR_ID,
			SHOULD_THROW
		)

		if (!coinbaseSdkConnector || !walletConnectConnector) {
			throw new Error("Expected connector(s) missing from wagmi context.")
		}

		// Special-case: Only display the injected connector for in-wallet browsers.
		if (isMobileWeb && injectedConnectors.length === 1) {
			return injectedConnectors
		}

		// Special-case: Only display the Coinbase connector in the Coinbase Wallet.
		if (isCoinbaseWalletBrowser) {
			return [coinbaseSdkConnector]
		}

		const orderedConnectors: InjectableConnector[] = [...injectedConnectors, coinbaseSdkConnector]

		if (!excludeWalletConnectConnections) {
			orderedConnectors.push(walletConnectConnector)
		}

		orderedConnectors.sort(sortByRecent)

		return orderedConnectors
	}, [connectors, excludeWalletConnectConnections, sortByRecent])
}
