import { getCsrfToken, signIn, useSession } from "next-auth/react"
import { createContext, PropsWithChildren, useCallback, useContext, useEffect } from "react"

import { UserRejectedRequestError } from "viem"
import { createSiweMessage } from "viem/siwe"
import {
	ResolvedRegister,
	useAccount,
	UseAccountReturnType,
	useChainId,
	UseConnectReturnType,
	useConnect as useConnectWagmi,
	useSignMessage,
	UseSignMessageReturnType
} from "wagmi"

import { useAtom, useSetAtom } from "jotai"

import { CONNECTION, getConnectorWithId, WalletConnectProvider } from "@/lib"
import { useDisconnect } from "@/lib/hooks/wallet/useDisconnect"
import {
	authenticationLoadingAtom,
	authenticationResponseAtom,
	useColumns,
	walletConnectProviderAtom,
	walletConnectURIAtom
} from "@/state"

const ConnectionContext = createContext<
	| {
			connection: UseConnectReturnType<ResolvedRegister["config"]>
			account: UseAccountReturnType<ResolvedRegister["config"]> & {
				isAuthenticated: boolean
			}
			sign: UseSignMessageReturnType
			prove: (index: number, from?: string, address?: string) => Promise<void>
	  }
	| undefined
>(undefined)

export function ConnectionProvider({ children }: PropsWithChildren) {
	const [walletConnectProvider, setWalletConnectProvider] = useAtom(walletConnectProviderAtom)
	const [walletConnectURI, setWalletConnectURI] = useAtom(walletConnectURIAtom)
	const setAuthenticationLoading = useSetAtom(authenticationLoadingAtom)
	const setAuthenticationResponse = useSetAtom(authenticationResponseAtom)

	const connection = useConnectWagmi({
		mutation: {
			onError(error) {
				if (error instanceof UserRejectedRequestError) connection.reset()
			}
		}
	})

	const chainId = useChainId()
	const account = useAccount()
	const sign = useSignMessage()
	const { disconnect } = useDisconnect()

	const { navigate } = useColumns()
	const { data: session } = useSession()

	const isAuthenticated = (account.address === session?.user.id && session?.user.id?.startsWith("0x")) || false

	/**
	 * Trigger and handle the signing of a message to prove ownership of the address.
	 * @param address The address to prove ownership of.
	 * @returns A promise that resolves when the signing is complete.
	 */
	const prove = useCallback(
		async (index: number, from?: string, address?: string) => {
			try {
				setAuthenticationLoading(false)

				const nonce = await getCsrfToken()

				if (!nonce) throw new Error("Could not get nonce.")

				const expirationTime = new Date(new Date().getTime() + 5 * 60 * 1000)
				const domain = window.location.host
				const uri = window.location.origin
				const statement = `Access the Plug platform by proving your ownership of the address: ${account.address}.`
				const message = createSiweMessage({
					address: (address ?? account.address) as `0x${string}`,
					chainId,
					domain,
					nonce,
					uri,
					version: "1",
					statement,
					expirationTime
				})

				sign.signMessage(
					{
						message
					},
					{
						onSuccess: async signature => {
							setAuthenticationLoading(true)
							setAuthenticationResponse(undefined)

							console.log("Authenticating...")
							const authenticationResponse = await signIn("credentials", {
								message,
								signature,
								chainId,
								redirect: true,
								callbackUrl: `${window.location.origin}/app`
							})
							console.log("Authenticated", authenticationResponse)

							setAuthenticationResponse(authenticationResponse)
							navigate({ index, key: from })
						},
						onError: (e, account) => {
							console.error("Authentication error", e)
							if (account.connector) account.connector.disconnect()

							connection.reset()
							sign.reset()
							disconnect()
						}
					}
				)
			} catch (e) {
				setAuthenticationLoading(false)

				connection.reset()
				sign.reset()
				disconnect()
			}
		},
		[navigate, connection, chainId, account, sign, disconnect, setAuthenticationLoading, setAuthenticationResponse]
	)

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
	const init = useCallback(async () => {
		const connector = getConnectorWithId(connection.connectors, CONNECTION.WALLET_CONNECT_CONNECTOR_ID, {
			shouldThrow: true
		})

		const provider = (await connector?.getProvider?.()) as WalletConnectProvider

		setWalletConnectProvider(provider)

		await provider.connect()
	}, [connection.connectors, setWalletConnectProvider])

	useEffect(() => {
		if (account.address) return

		if (!walletConnectProvider) {
			init()
		} else if (walletConnectURI === undefined) {
			walletConnectProvider.once("display_uri", setWalletConnectURI)
		}
	}, [account.address, walletConnectProvider, walletConnectURI, init, setWalletConnectURI])

	return (
		<ConnectionContext.Provider value={{ connection, account: { ...account, isAuthenticated }, sign, prove }}>
			{children}
		</ConnectionContext.Provider>
	)
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
