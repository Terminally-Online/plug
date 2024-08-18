import {
	FC,
	HTMLAttributes,
	PropsWithChildren,
	useCallback,
	useEffect
} from "react"

import { getCsrfToken, signIn, signOut, useSession } from "next-auth/react"

import { SiweMessage } from "siwe"
import { useAccount, useChainId, useDisconnect, useSignMessage } from "wagmi"

import { useWeb3Modal } from "@web3modal/wagmi/react"

import { Button } from "@/components"
import { cn } from "@/lib"

export type ButtonProps = {
	callbackUrl?: string
	redirect?: boolean
}

export const AuthButton: FC<
	HTMLAttributes<HTMLButtonElement> & PropsWithChildren<ButtonProps>
> = ({ callbackUrl = "/app/", redirect = true, className, ...props }) => {
	const { open } = useWeb3Modal()

	const { address, isConnected } = useAccount()
	const chainId = useChainId()

	const { signMessageAsync, isLoading, isError } = useSignMessage()

	const { data: session } = useSession()
	const { disconnect } = useDisconnect({
		mutation: {
			onSuccess: () => signOut()
			// { callbackUrl: "/" }
		}
	})

	const handleLogin = useCallback(async () => {
		if (!isConnected) {
			open()
			return
		}

		try {
			const message = new SiweMessage({
				domain: window.location.host,
				address,
				statement: `Access the Plug platform by proving your ownership of the address: ${address}.`,
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
		} catch (e) {
			console.error(e)
		}
	}, [
		redirect,
		address,
		chainId,
		isConnected,
		open,
		signMessageAsync,
		callbackUrl
	])

	useEffect(() => {
		if (isConnected === false || isLoading || isError) return

		handleLogin()
	}, [isConnected, isLoading, isError, handleLogin])

	return (
		<>
			{session?.address ? (
				<Button
					variant="destructive"
					className={cn(className ? className : "w-full")}
					onClick={() => disconnect()}
					{...props}
				>
					Logout
				</Button>
			) : (
				<Button
					className={cn(className ? className : "w-full")}
					onClick={handleLogin}
					{...props}
				>
					{isConnected
						? isLoading
							? "Signing Message..."
							: "Sign Message"
						: "Connect Wallet"}
				</Button>
			)}
		</>
	)
}

export default AuthButton
