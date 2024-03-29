import { type FC, useState } from "react"

import { signOut } from "next-auth/react"
import { getCsrfToken, signIn, useSession } from "next-auth/react"

import { SiweMessage } from "siwe"
import { useAccount, useChainId, useSignMessage } from "wagmi"

import { useWeb3Modal } from "@web3modal/wagmi/react"

export type SiweProps = Partial<{
	callbackUrl: string
	redirect: boolean
}>

const Siwe: FC<SiweProps> = ({
	callbackUrl = "/protected",
	redirect = false
} = {}) => {
	const chainId = useChainId()
	const { data: session } = useSession()
	const { open } = useWeb3Modal()
	const { address, isConnected } = useAccount()
	const { signMessageAsync } = useSignMessage()

	const [error, setError] = useState<unknown | null>(null)

	const isAuthenticated = session !== undefined && session !== null

	const handleLogin = async (e: React.MouseEvent<HTMLButtonElement>) => {
		e.preventDefault()

		console.log("handling", isAuthenticated)

		// ? This user is already authenticated.
		if (isAuthenticated) return

		setError(null)

		if (isConnected) {
			try {
				const message = new SiweMessage({
					domain: window.location.host,
					address,
					statement:
						"Sign in to Plug by signing this message to prove your identity.",
					uri: window.location.origin,
					version: "1",
					chainId: chainId,
					nonce: await getCsrfToken()
				})
				const signature = await signMessageAsync({
					message: message.prepareMessage()
				})

				signIn("credentials", {
					message: JSON.stringify(message),
					redirect,
					signature,
					callbackUrl
				})
			} catch (error: unknown) {
				setError(error)
			}
		}

		// * The user must be connected before they can be authenticated.
		open()
	}

	return (
		<>
			{!session ? (
				<button onClick={handleLogin}>Sign-in</button>
			) : (
				<button onClick={() => signOut()}>Sign out</button>
			)}

			{error}
		</>
	)
}

export default Siwe
