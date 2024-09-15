import { getCsrfToken, signIn } from "next-auth/react"
import { createContext, PropsWithChildren, useCallback, useContext } from "react"

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

import { useDisconnect } from "@/lib/hooks/wallet/useDisconnect"

const ConnectionContext = createContext<
	| {
			connection: UseConnectReturnType<ResolvedRegister["config"]>
			account: UseAccountReturnType<ResolvedRegister["config"]>
			sign: UseSignMessageReturnType
			prove: (address?: string) => Promise<void>
	  }
	| undefined
>(undefined)

export function ConnectionProvider({ children }: PropsWithChildren) {
	const connection = useConnectWagmi({
		mutation: {
			onError(error) {
				if (error instanceof UserRejectedRequestError) connection.reset()
			}
		}
	})

	const chainId = useChainId()
	const account = useAccount()
	const { disconnect } = useDisconnect()
	const sign = useSignMessage()

	const prove = useCallback(
		async (address?: string) => {
			try {
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
						onSuccess: signature =>
							signIn("credentials", {
								message,
								signature,
								chainId,
								redirect: true,
								callbackUrl: "/app/"
							}),
						onError: (_, account) => {
							if (account.connector) account.connector.disconnect()

							connection.reset()
							sign.reset()
							disconnect()
						}
					}
				)
			} catch (e) {
				console.error("hit the error here")
				connection.reset()
				sign.reset()
				disconnect()
			}
		},
		[connection, chainId, account, sign, disconnect]
	)

	// useEffect(() => {
	// 	if (!accountDrawer.isOpen && connection.isPending) {
	// 		connection.reset()
	// 		disconnect()
	// 	}
	// }, [connection, accountDrawer.isOpen, disconnect])

	return (
		<ConnectionContext.Provider value={{ connection, account, sign, prove }}>{children}</ConnectionContext.Provider>
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
