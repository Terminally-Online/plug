"use client"

import { type FC, type PropsWithChildren, useEffect, useState } from "react"

import Link from "next/link"
import { useRouter } from "next/router"

import { AnimatePresence, motion } from "framer-motion"
import { XIcon } from "lucide-react"
import { useAccount, useEnsName } from "wagmi"

import Button from "./button"

export const WalletConnector: FC<PropsWithChildren> = ({ children }) => {
	const router = useRouter()

	const { address } = useAccount()

	const [client, setClient] = useState(false)

	const connect = router.query.connect
	const isConnecting = connect === "true"

	const handleClose = () => {
		router.push(router.pathname, undefined, { shallow: true })
	}

	useEffect(() => {
		setClient(true)
	}, [])

	return (
		<>
			{isConnecting && client ? (
				<motion.div
					initial={{ opacity: 0 }}
					animate={{ opacity: 1 }}
					exit={{ opacity: 0 }}
					transition={{ duration: 0.2 }}
				>
					<motion.div
						className="absolute bottom-0 left-0 right-0 top-0 z-[999998] cursor-pointer bg-stone-900 bg-opacity-5 backdrop-blur-xl"
						onClick={handleClose}
					/>

					<motion.div
						className="absolute bottom-0 right-0 top-0 z-[999999] flex flex flex h-full w-[420px] flex-col justify-center gap-2 border-l-[1px] border-stone-950 bg-stone-900 bg-opacity-90 backdrop-blur-xl"
						initial={{ opacity: 0, right: "-10%" }}
						animate={{ opacity: 1, right: 0 }}
						exit={{ opacity: 0 }}
						transition={{ duration: 0.2, delay: 0.2 }}
					>
						<button
							className="group absolute right-4 top-4 cursor-pointer text-white"
							onClick={handleClose}
						>
							<XIcon
								width={16}
								height={16}
								className="opacity-60 transition-all duration-200 ease-in-out group-hover:opacity-80"
							/>
						</button>

						<div className="mt-auto flex flex-col gap-4 p-12">
							<h1 className="text-3xl">
								{address ? "Finish logging in..." : "Log in."}
							</h1>
							<p className="opacity-60">
								{address
									? "You're almost done! Just sign a message when used to prove your ownership of this wallet and be used to create a Plug session."
									: "Connect your wallet to continue. If you don't have an account yet, one will automatically be made for you."}
							</p>
						</div>

						<Button />

						<p className="mt-auto p-12 opacity-60">
							By continuing, you agree to our{" "}
							<Link href="/terms">Terms of Service</Link> and{" "}
							<Link href="/privacy">Privacy Policy</Link>.
						</p>
					</motion.div>
				</motion.div>
			) : null}

			{children}
		</>
	)
}

export default WalletConnector
