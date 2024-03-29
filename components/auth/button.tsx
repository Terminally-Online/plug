"use client"

import type { FC, PropsWithChildren } from "react"

import Image from "next/image"

import { getCsrfToken, signIn } from "next-auth/react"

import { AnimatePresence } from "framer-motion"
import { motion } from "framer-motion"
import {
	ClipboardSignatureIcon,
	FileSignatureIcon,
	WalletCardsIcon
} from "lucide-react"
import { SiweMessage } from "siwe"
import {
	useAccount,
	useChainId,
	useEnsAvatar,
	useEnsName,
	useSignMessage
} from "wagmi"

import { useWeb3Modal } from "@web3modal/wagmi/react"

export type ButtonProps = {
	callbackUrl?: string
	redirect?: boolean
}

export const Button: FC<PropsWithChildren<ButtonProps>> = ({
	callbackUrl = "/canvas/",
	redirect = true
}) => {
	const { open } = useWeb3Modal()
	const { address, isConnected } = useAccount()
	const chainId = useChainId()
	const { data: name } = useEnsName({ address })
	// TODO: Fix the name retrieval.
	const { data: avatar } = useEnsAvatar({ name: 'nftchance.eth' })

	const { error, signMessageAsync, isLoading, isError } = useSignMessage()

	const handleLogin = async () => {
		if (!isConnected) {
			open()
			return
		}

		try {
			const message = new SiweMessage({
				domain: window.location.host,
				address,
				statement: `Access the Plug platform by proving your ownership of the address: ${
					name ?? address
				}.`,
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
	}

	return (
		<>
			<button
				type="button"
				onClick={handleLogin}
				className="group flex w-full flex-row items-center justify-center gap-4 border-y-[1px] border-stone-950 bg-white p-4 uppercase text-stone-950 transition-all transition-all duration-200 duration-200 ease-in-out ease-in-out hover:bg-stone-950 hover:text-white"
			>
				{isConnected ? (
					<>
						{isLoading ? (
							<>
								<ClipboardSignatureIcon
									width={16}
									height={16}
									className="opacity-40 group-hover:opacity-60"
								/>{" "}
								Proving Ownership...
							</>
						) : (
							<>
								<FileSignatureIcon
									width={16}
									height={16}
									className="opacity-40 group-hover:opacity-60"
								/>{" "}
								Prove Account Ownership
							</>
						)}
					</>
				) : (
					<>
						<WalletCardsIcon
							width={16}
							height={16}
							className="opacity-40 group-hover:opacity-60"
						/>{" "}
						Connect Wallet
					</>
				)}
			</button>

			{address ? (
				<div className="mt-2 flex w-full justify-center px-12 text-center">
					<p className="flex flex-row items-center gap-2">
						<span className="opacity-60">of</span>
						<a
							className="group flex flex-row items-center gap-2"
							href={`https://etherscan.io/address/${address}`}
							target="_blank"
							rel="noopener noreferrer"
						>
							{name && avatar ? (
								<Image
									src={avatar}
									alt={name}
									className="h-5 w-5 rounded-full"
								/>
							) : address ? (
								<></>
							) : (
								<></>
							)}
							<span className="tabular-nums opacity-60 transition-all duration-200 ease-in-out group-hover:opacity-100">
								{name
									? name
									: `${address.slice(0, 8)}...${address.slice(
											-4
									  )}`}
							</span>
						</a>
					</p>
				</div>
			) : null}

			<AnimatePresence>
				{error && isError && (
					<motion.p
						className="mt-2 flex w-full justify-center px-12 text-center text-red-500"
						initial={{ opacity: 0 }}
						animate={{ opacity: 1 }}
						exit={{ opacity: 0 }}
					>
						{error.message}
					</motion.p>
				)}
			</AnimatePresence>
		</>
	)
}

export default Button
