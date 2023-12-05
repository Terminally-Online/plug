import { useEffect, useState } from "react"

import { AnimatePresence, motion } from "framer-motion"
import { Puzzle, Receipt, Sparkles } from "lucide-react"

import Grid from "@/components/landing/grid"
import type { NextPageWithLayout } from "@/lib/types"

import PageLayout from "./index.layout"

const chains = [
	["ethereum", 7500],
	["optimism", 5000],
	["base", 5000],
	["polygon", 5000],
	["arbitrum", 3000],
	["avalanche", 3000],
	["bsc", 3000]
]

const Page: NextPageWithLayout = () => {
	const [chain, setChain] = useState(chains[0])

	useEffect(() => {
		const delay = chain[1] as number

		const timeout = setTimeout(() => {
			const index = chains.indexOf(chain)
			const next = index === chains.length - 1 ? 0 : index + 1

			setChain(chains[next])
		}, delay)

		return () => clearTimeout(timeout)
	}, [chain])

	return (
		<>
			<Grid />

			<div className="mt-auto flex h-full flex-col">
				<div className="mx-4 mb-12 h-full lg:mx-8">
					<motion.h1 className="text-shadow-blur mb-12 mt-12 text-4xl shadow-white dark:shadow-black lg:mt-0 lg:text-7xl">
						<span className="text-black/60 dark:text-white/60">
							<motion.span
								className="text-black dark:text-white"
								animate={{ opacity: [0.6, 1] }}
								transition={{
									duration: 0.4,
									delay: 0.4,
									ease: "easeInOut"
								}}
							>
								PLUG
							</motion.span>
							-AND-
							<motion.span
								className="text-black dark:text-white"
								animate={{ opacity: [0.6, 1] }}
								transition={{
									duration: 0.4,
									delay: 0.4,
									ease: "easeInOut"
								}}
							>
								PLAY
							</motion.span>
							<br /> {'"'}
						</span>
						<motion.span
							className="text-black dark:text-white"
							animate={{ opacity: [0.6, 1] }}
							transition={{
								duration: 0.4,
								delay: 0.8,
								ease: "easeInOut"
							}}
						>
							IF THIS
						</motion.span>
						<span className="text-black/60 dark:text-white/60">
							,
						</span>
						<motion.span
							className="text-black dark:text-white"
							animate={{ opacity: [0.6, 1] }}
							transition={{
								duration: 0.4,
								delay: 1.2,
								ease: "easeInOut"
							}}
						>
							THEN THAT
						</motion.span>
						<span className="text-black/60 dark:text-white/60">
							{'"'}
						</span>
						<br />
						<span className="text-black/60 dark:text-white/60">
							FOR
						</span>
						<motion.span
							className="relative w-full"
							animate={{ opacity: [0.6, 1] }}
							transition={{
								duration: 0.4,
								delay: 1.6,
								ease: "easeInOut"
							}}
						>
							{" "}
							<AnimatePresence>
								{chains.map(
									(c, i) =>
										chain === c && (
											<motion.span
												key={i}
												initial={{
													display: "none",
													opacity: 0
												}}
												animate={{
													display: "inline-block",
													opacity: 1
												}}
												exit={{
													display: "none",
													opacity: 0
												}}
												transition={{
													duration: 0.4,
													delayChildren: 1.6,
													ease: "easeInOut"
												}}
											>
												{`${chain[0]}`.toUpperCase()}
											</motion.span>
										)
								)}
							</AnimatePresence>{" "}
						</motion.span>
						<span className="text-black/60 dark:text-white/60">
							TRANSACTIONS.
						</span>
					</motion.h1>

					<div className="col-span-12 grid grid-cols-10 gap-4 lg:grid-cols-12 lg:gap-12 xl:grid-cols-10">
						<div className="text-shadow-blur col-span-9 shadow-white dark:shadow-black lg:col-span-4 xl:col-span-3">
							<h3 className="mb-4 flex flex-row items-center gap-4 bg-gradient-to-b from-black to-black/60 bg-clip-text text-2xl text-transparent dark:from-white dark:to-white/60">
								<Sparkles
									className="text-black/40 dark:text-white/40"
									size={24}
								/>
								CONDITIONAL EXECUTION
							</h3>
							<p className="ml-10 bg-gradient-to-t from-black to-black/60 bg-clip-text text-justify text-transparent dark:from-white dark:to-white/60">
								Unlocks the power of conditional execution for
								transactions on EVM blockchains. From simple to
								complex, Plug can scale till your heart{'\''}s
								content.
							</p>
						</div>

						<div className="text-shadow-blur col-span-9 shadow-white dark:shadow-black lg:col-span-4 xl:col-span-3">
							<h3 className="mb-4 flex flex-row items-center gap-4 bg-gradient-to-b from-black to-black/60 bg-clip-text text-2xl text-transparent dark:from-white dark:to-white/60">
								<Puzzle
									className="text-black/40 dark:text-white/40"
									size={24}
								/>
								SEAMLESS INTEGRATION
							</h3>
							<p className="ml-10 bg-gradient-to-t from-black to-black/60 bg-clip-text text-justify text-transparent dark:from-white dark:to-white/60">
								Designed to work with any smart contract
								deployed past, present or future, Plug has
								several levels of integration to chose from. No
								need to deploy a new contract.
							</p>
						</div>

						<div className="text-shadow-blur col-span-9 shadow-white dark:shadow-black lg:col-span-4 xl:col-span-3">
							<h3 className="mb-4 flex flex-row items-center gap-4 bg-gradient-to-b from-black to-black/60 bg-clip-text text-2xl text-transparent dark:from-white dark:to-white/60">
								<Receipt
									className="text-black/40 dark:text-white/40"
									size={24}
								/>
								ALL EVM BLOCKCHAINS
							</h3>
							<p className="ml-10 bg-gradient-to-t from-black to-black/60 bg-clip-text text-justify text-transparent dark:from-white dark:to-white/60">
								Plug was designed to operate on all EVM
								compatible blockchains such as Optimism, Base,
								etc.. With this, cross-chain functionality is
								included right out of the box.
							</p>
						</div>
					</div>
				</div>
			</div>
		</>
	)
}

Page.getLayout = page => <PageLayout>{page}</PageLayout>

export default Page
