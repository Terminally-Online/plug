import { FC, PropsWithChildren, useEffect, useMemo, useState } from "react"

import Image from "next/image"

import { motion } from "framer-motion"
import { CalendarClock } from "lucide-react"

import { InfoCard } from "@/components"
import { formatTitle } from "@/lib"

const actions = [
	[
		{
			protocol: "plug",
			elements: (
				<>
					Run <span className="rounded-sm bg-[#00E100]/10 p-2 py-1 text-[#00E100]">1</span> time a{" "}
					<span className="rounded-sm bg-[#00E100]/10 p-2 py-1 text-[#00E100]">year</span> .
				</>
			)
		},
		{
			protocol: "uniswap",
			elements: (
				<>
					Can swap <span className="rounded-sm bg-[#00E100]/10 p-2 py-1 text-[#00E100]">100</span>{" "}
					<span className="rounded-sm bg-[#00E100]/10 p-2 py-1 text-[#00E100]">$USDC</span> to{" "}
					<span className="rounded-sm bg-[#00E100]/10 p-2 py-1 text-[#00E100]">$ETH</span> .
				</>
			)
		},
		{
			protocol: "plug",
			elements: (
				<>
					Has <span className="rounded-sm bg-[#00E100]/10 p-2 py-1 text-[#00E100]">0.2</span>{" "}
					<span className="rounded-sm bg-[#00E100]/10 p-2 py-1 text-[#00E100]">$ETH</span> or greater.
				</>
			)
		},
		{
			protocol: "ens",
			elements: (
				<>
					Can renew <span className="rounded-sm bg-[#00E100]/10 p-2 py-1 text-[#00E100]">nftchance.eth</span>{" "}
					.
				</>
			)
		}
	],
	[
		{
			protocol: "plug",
			elements: (
				<>
					Run <span className="rounded-sm bg-[#00E100]/10 p-2 py-1 text-[#00E100]">1</span> time a{" "}
					<span className="rounded-sm bg-[#00E100]/10 p-2 py-1 text-[#00E100]">day</span> .
				</>
			)
		},
		{
			protocol: "plug",
			elements: (
				<>
					Run after <span className="rounded-sm bg-[#00E100]/10 p-2 py-1 text-[#00E100]">10</span>{" "}
					<span className="rounded-sm bg-[#00E100]/10 p-2 py-1 text-[#00E100]">PM</span>{" "}
					<span className="rounded-sm bg-[#00E100]/10 p-2 py-1 text-[#00E100]">UTC</span> .
				</>
			)
		},
		{
			protocol: "plug",
			elements: (
				<>
					Run before <span className="rounded-sm bg-[#00E100]/10 p-2 py-1 text-[#00E100]">12/31/2024</span> .
				</>
			)
		},
		{
			protocol: "plug",
			elements: (
				<>
					Swap <span className="rounded-sm bg-[#00E100]/10 p-2 py-1 text-[#00E100]">36,000</span>{" "}
					<span className="rounded-sm bg-[#00E100]/10 p-2 py-1 text-[#00E100]">$USDC</span> to{" "}
					<span className="rounded-sm bg-[#00E100]/10 p-2 py-1 text-[#00E100]">$ETH</span> .
				</>
			)
		},
		{
			protocol: "nouns",
			elements: (
				<>
					Can bid with <span className="rounded-sm bg-[#00E100]/10 p-2 py-1 text-[#00E100]">9</span>{" "}
					<span className="rounded-sm bg-[#00E100]/10 p-2 py-1 text-[#00E100]">$ETH</span> .
				</>
			)
		},
		{
			protocol: "nouns",
			elements: (
				<>
					Bid on Noun with{" "}
					<span className="rounded-sm bg-[#00E100]/10 p-2 py-1 text-[#00E100]">Pineapple Hat</span> .
				</>
			)
		}
	],
	[
		{
			protocol: "plug",
			elements: (
				<>
					Run <span className="rounded-sm bg-[#00E100]/10 p-2 py-1 text-[#00E100]">1</span> time a{" "}
					<span className="rounded-sm bg-[#00E100]/10 p-2 py-1 text-[#00E100]">hour</span> .
				</>
			)
		},
		{
			protocol: "yearn",
			elements: (
				<>
					<span className="rounded-sm bg-[#00E100]/10 p-2 py-1 text-[#00E100]">$USDC</span> pool is above{" "}
					<span className="rounded-sm bg-[#00E100]/10 p-2 py-1 text-[#00E100]">72%</span> APY.
				</>
			)
		},
		{
			protocol: "yearn",
			elements: (
				<>
					Can deposit <span className="rounded-sm bg-[#00E100]/10 p-2 py-1 text-[#00E100]">10,000</span>{" "}
					<span className="rounded-sm bg-[#00E100]/10 p-2 py-1 text-[#00E100]">$USDC</span> .
				</>
			)
		}
	]
]

const ExecutionAction: FC<PropsWithChildren<{ index: number; indexes: number; protocol: string }>> = ({
	index,
	indexes,
	protocol,
	children
}) => {
	return (
		<>
			<motion.div
				className="relative flex w-full flex-row items-center gap-4 rounded-lg"
				initial={{ opacity: 0, transform: "translateY(20px)" }}
				animate={{ opacity: 1, transform: "translateY(0px)" }}
				transition={{
					duration: 0.2,
					delay: index * 0.05,
					repeat: Infinity,
					repeatType: "reverse",
					repeatDelay: 3 * indexes
				}}
			>
				<motion.div
					className="bg-gradient-radial absolute inset-0 rounded-lg from-plug-green via-plug-yellow to-white opacity-0 blur-lg filter"
					initial={{ opacity: 0 }}
					animate={{
						opacity: [0, 0.1, 0.2, 0],
						background: [
							"radial-gradient(circle at 0% 0%, #00E100 0%, #A3F700 50%, #FFFFFF 100%)",
							"radial-gradient(circle at 100% 0%, #00E100 0%, #A3F700 50%, #FFFFFF 100%)",
							"radial-gradient(circle at 100% 100%, #00E100 0%, #A3F700 50%, #FFFFFF 100%)",
							"radial-gradient(circle at 0% 100%, #00E100 0%, #A3F700 50%, #FFFFFF 100%)",
							"radial-gradient(circle at 0% 0%, #00E100 0%, #A3F700 50%, #FFFFFF 100%)"
						]
					}}
					transition={{
						duration: 2,
						ease: "linear",
						delay: 0.1 + index * 2,
						repeat: Infinity
					}}
				/>

				<motion.div
					className="bg-gradient-radial absolute inset-0 rounded-lg from-plug-green via-plug-yellow to-white opacity-0"
					initial={{ opacity: 0 }}
					animate={{
						opacity: [0, 1, 1, 0],
						background: [
							"radial-gradient(circle at 0% 0%, #00E100 0%, #A3F700 50%, #FFFFFF 100%)",
							"radial-gradient(circle at 100% 0%, #00E100 0%, #A3F700 50%, #FFFFFF 100%)",
							"radial-gradient(circle at 100% 100%, #00E100 0%, #A3F700 50%, #FFFFFF 100%)",
							"radial-gradient(circle at 0% 100%, #00E100 0%, #A3F700 50%, #FFFFFF 100%)",
							"radial-gradient(circle at 0% 0%, #00E100 0%, #A3F700 50%, #FFFFFF 100%)"
						]
					}}
					transition={{
						duration: 2,
						ease: "linear",
						delay: 0.1 + index * 2,
						repeat: Infinity
					}}
				/>

				<div className="relative z-10 m-[2px] flex w-full flex-row items-center gap-4 rounded-lg bg-white p-3">
					<Image
						src={`/protocols/${protocol}.png`}
						alt={formatTitle(protocol)}
						width={48}
						height={48}
						className="h-6 w-6 rounded-sm"
					/>
					<p className="font-bold text-black/40">{children}</p>
				</div>
			</motion.div>

			{index !== indexes && (
				<motion.div
					className="h-4 w-[2px]"
					initial={{ opacity: 0, transform: "translateY(20px)" }}
					animate={{ opacity: 1, transform: "translateY(0px)" }}
					transition={{
						duration: 0.2,
						delay: index * 0.05,
						repeat: Infinity,
						repeatType: "reverse",
						repeatDelay: 3 * indexes
					}}
				>
					<motion.div
						className="h-full w-full bg-grayscale-100"
						initial={{ background: "linear-gradient(30deg, #EBECEC, #EBECEC)" }}
						animate={{ background: "linear-gradient(30deg, #00E100, #A3F700)" }}
						transition={{
							duration: 0.2,
							ease: "linear",
							delay: 2 + index * 2,
							repeat: Infinity,
							repeatType: "reverse",
							repeatDelay: 3 * indexes
						}}
					/>
				</motion.div>
			)}
		</>
	)
}

const ExecutionActions = () => {
	const [currentActionSet, setCurrentActionSet] = useState(0)

	useEffect(() => {
		const interval = setInterval(
			() => {
				setCurrentActionSet(prev => (prev + 1) % actions.length)
			},
			(actions[currentActionSet].length * 3 + 0.5 * 2) * 1000
		)

		return () => clearInterval(interval)
	}, [currentActionSet])

	return (
		<div className="flex flex-col items-center">
			{actions[currentActionSet].map((action, index) => (
				<ExecutionAction
					key={`${currentActionSet}-${index}`}
					index={index + 1}
					indexes={actions[currentActionSet].length}
					protocol={action.protocol}
				>
					{action.elements}
				</ExecutionAction>
			))}
		</div>
	)
}

export const Execution = () => {
	return (
		<InfoCard
			icon={<CalendarClock size={24} className="opacity-40" />}
			text='"If this, then that” driven execution.'
			description="Stop missing opportunities and have your transactions execute no matter where you are or what you're doing. If you want it done, it will be delivered on a silver platter."
			className="col-span-2 h-[540px] xl:col-span-4 xl:row-span-2 xl:h-full"
		>
			<div className="flex h-[80%] w-full flex-col items-center justify-center gap-2">
				<ExecutionActions />
			</div>
		</InfoCard>
	)
}
