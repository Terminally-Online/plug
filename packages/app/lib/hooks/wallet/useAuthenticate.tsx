import { authenticationLoadingAtom, authenticationResponseAtom } from "@/state/authentication"
import { useColumnActions } from "@/state/columns"
import { useAtom, useSetAtom } from "jotai"
import { getCsrfToken, signIn } from "next-auth/react"
import { useCallback, useState } from "react"
import { Account, SignableMessage } from "viem"
import { createSiweMessage } from "viem/siwe"
import { Connector, useAccount, useChainId, useDisconnect, useSignMessage } from "wagmi"

export const useAuthenticate = () => {
	const chainId = useChainId()
	const account = useAccount()
	const { signMessage, reset } = useSignMessage()
	const { disconnect } = useDisconnect()

	const [authenticationLoading, setAuthenticationLoading] = useAtom(authenticationLoadingAtom)
	const [authenticationError, setAuthenticationError] = useState<unknown | undefined>()
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
		async (index: number, from?: string, address?: string) => {
			try {
				setAuthenticationLoading(false)

				const message = await createMessage(address ?? account.address)
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
					// NOTE: This navigate should have been done in an onSuccess callback.
					navigate({ index, key: from })
				}
				const handleAuthenticationError = (error: unknown, account: {
					account?: `0x${string}` | Account | undefined,
					message: SignableMessage,
					connector?: Connector | undefined
				}) => {
					setAuthenticationError(error)

					if (account.connector) account.connector.disconnect()

					reset()
					disconnect()
				}

				signMessage({ message }, {
					onSuccess: handleAuthenticationSuccess,
					onError: handleAuthenticationError
				})
			} catch (e) {
				setAuthenticationLoading(false)

				reset()
				disconnect()
			}
		},
		[navigate, chainId, account, signMessage, reset, disconnect, setAuthenticationLoading, setAuthenticationResponse]
	)

	return { authenticate, error: authenticationError, isLoading: authenticationLoading, isError: authenticationError !== undefined }
}
