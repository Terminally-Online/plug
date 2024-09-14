import { FC, HTMLAttributes, PropsWithChildren, useCallback, useEffect } from "react"

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

export const AuthButton: FC<HTMLAttributes<HTMLButtonElement> & PropsWithChildren<ButtonProps>> = ({
	callbackUrl = "/app/",
	redirect = true,
	className,
	...props
}) => {
	// const { open } = useWeb3Modal()

	// const { address, isConnected } = useAccount()
	// const chainId = useChainId()

	// const { signMessageAsync, isLoading, isError } = useSignMessage()

	// const { data: session } = useSession()
	// const { disconnect } = useDisconnect({
	// 	mutation: {
	// 		onSuccess: () => signOut()
	// 		// { callbackUrl: "/" }
	// 	}
	// })

	// useEffect(() => {
	// 	if (isConnected === false || isLoading || isError) return

	// 	handleLogin()
	// }, [isConnected, isLoading, isError, handleLogin])

	return <></>
}

export default AuthButton
