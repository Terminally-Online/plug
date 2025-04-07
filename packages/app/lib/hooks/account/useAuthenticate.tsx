import { getCsrfToken, signIn, SignInResponse } from "next-auth/react"
import { useCallback } from "react"

import { Account, SignableMessage } from "viem"
import { createSiweMessage } from "viem/siwe"
import { Connector, useAccount, useChainId, useDisconnect, useSignMessage } from "wagmi"

import { useSetAtom } from "jotai"

import { authenticationLoadingAtom, authenticationResponseAtom } from "@/state/authentication"
import { useColumnActions } from "@/state/columns"

export const useAuthenticate = () => {
	const chainId = useChainId()
	const account = useAccount()

	const { signMessage, reset, error, failureReason, isPending: isLoading, isError } = useSignMessage()
	const { disconnect } = useDisconnect()

	const setAuthenticationLoading = useSetAtom(authenticationLoadingAtom)
	const setAuthenticationResponse = useSetAtom(authenticationResponseAtom)

	const { navigate } = useColumnActions()

	const createMessage = async (address?: string) => {
		if (!address) throw new Error("No wallet connected.")

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
		return message
	}

	const authenticate = useCallback(
		async (
			context?: {
				address?: string
			},
			options?: {
				onSuccess?: (response: SignInResponse | undefined) => void
				onError?: (error: unknown) => void
			}
		) => {
			try {
				const user = context?.address ?? account.address

				if (!user) throw new Error("No user to authenticate")

				setAuthenticationLoading(false)

				const handleAuthenticationSuccess = async (signature: string) => {
					setAuthenticationLoading(true)
					setAuthenticationResponse(undefined)

					const authenticationResponse = await signIn("credentials", {
						message,
						signature,
						chainId,
						redirect: true,
						callbackUrl: `${window.location.origin}/app${window.location.search}`
					})

					setAuthenticationResponse(authenticationResponse)
					options?.onSuccess?.(authenticationResponse)
				}
				const handleAuthenticationError = (
					error: unknown,
					account: {
						account?: `0x${string}` | Account | undefined
						message: SignableMessage
						connector?: Connector | undefined
					}
				) => {
					if (account.connector) account.connector.disconnect()

					reset()
					disconnect()

					options?.onError?.(error)
				}

				const message = await createMessage(user)
				signMessage(
					{ message },
					{
						onSuccess: handleAuthenticationSuccess,
						onError: handleAuthenticationError
					}
				)
			} catch (e) {
				setAuthenticationLoading(false)

				reset()
				disconnect()
			}
		},
		[
			navigate,
			chainId,
			account,
			signMessage,
			reset,
			disconnect,
			setAuthenticationLoading,
			setAuthenticationResponse
		]
	)

	return { authenticate, error, failureReason, isLoading, isError }
}
