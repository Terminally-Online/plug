import { createContext, PropsWithChildren, useContext, useEffect } from "react"

import { UserRejectedRequestError } from "viem"
import { ResolvedRegister, UseConnectReturnType, useConnect as useConnectWagmi } from "wagmi"
import { useDisconnect } from "wagmi"

import { useAtom } from "jotai"

import { CONNECTION, WalletConnectProvider } from "@/lib"
import { walletConnectURIAtom } from "@/state/authentication"
import { useSidebar } from "@/state/sidebar"

import { getConnectorWithId } from "./useConnections"

const ConnectionContext = createContext<UseConnectReturnType<ResolvedRegister["config"]> | undefined>(undefined)

export function ConnectionProvider({ children }: PropsWithChildren) {
	const [walletConnectURI, setWalletConnectURI] = useAtom(walletConnectURIAtom)

	const { is } = useSidebar()

	const connection = useConnectWagmi({
		mutation: {
			onError(error) {
				if (error instanceof UserRejectedRequestError) connection.reset()
			}
		}
	})
	const { disconnect } = useDisconnect()

	/**
	 * Prepare the wallet connect provider for displaying the QR code outside of the general context.
	 *
	 * Without calling `.connect()` on the provider, the QR code will not be displayed as we would need it
	 * to be triggered by a user clicking a button.
	 *
	 * Additionally, if we do not load it preemptively, the QR code
	 * will have a loading cycle. This way, it is constantly prepared and ready to be displayed.
	 * @see {@link https://www.npmjs.com/package/@walletconnect/ethereum-provider}
	 */
	useEffect(() => {
		if (walletConnectURI) return

		const init = async () => {
			const connector = getConnectorWithId(connection.connectors, CONNECTION.WALLET_CONNECT_CONNECTOR_ID, {
				shouldThrow: true
			})
			const provider = (await connector?.getProvider?.()) as WalletConnectProvider
			provider.once("display_uri", setWalletConnectURI)

			await provider.connect()
		}

		init()
	}, [connection.connectors, walletConnectURI, setWalletConnectURI])

	/**
	 * When the user had an active connection, but closes the authentication context we go ahead and
	 * reset the connection to clear out the process that the user has signalled to be abandoned.
	 */
	useEffect(() => {
		if (!is.authenticating && connection.isPending) {
			connection.reset()
			disconnect()
		}
	}, [connection, is.authenticating, disconnect])

	// useEffect(() => {
	// 	const isAnonymous = !socket.id.startsWith("0x")
	// 	const isPotentiallyExpired = !address || socket.id === address
	// 	if (isAnonymous || isPotentiallyExpired) return
	//
	// 	disconnect()
	// }, [socket, address, disconnect])

	return <ConnectionContext.Provider value={connection}>{children}</ConnectionContext.Provider>
}

/**
 * Wraps wagmi.useConnect in a singleton provider to provide the same connect state to all callers.
 * @see {@link https://wagmi.sh/react/api/hooks/useConnect}
 * @see {@link https://wagmi.sh/react/api/hooks/useAccount}
 * @see {@link https://wagmi.sh/react/api/hooks/useSignMessage}
 */
export function useConnect() {
	const value = useContext(ConnectionContext)
	if (!value) {
		throw new Error("useConnect must be used within a ConnectionProvider")
	}
	return value
}
