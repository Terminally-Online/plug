import { FC, PropsWithChildren, useEffect, useState } from "react"

import { motion } from "framer-motion"
import { CalendarClock } from "lucide-react"

import { Image } from "@/components/app/utils/image"
import { InfoCard } from "@/components/landing/cards/info"
import { formatTitle } from "@/lib"

const actions = [
	[
		{
			protocol: "plug",
			elements: (
				<>
					Run <span className="rounded-sm bg-[#D2F38A]/10 p-2 py-1 text-[#385842]">1</span> time a{" "}
					<span className="rounded-sm bg-[#D2F38A]/10 p-2 py-1 text-[#385842]">year</span> .
				</>
			)
		},
		{
			protocol: "uniswap",
			elements: (
				<>
					Can swap <span className="rounded-sm bg-[#D2F38A]/10 p-2 py-1 text-[#385842]">100</span>{" "}
					<span className="rounded-sm bg-[#D2F38A]/10 p-2 py-1 text-[#385842]">$USDC</span> to{" "}
					<span className="rounded-sm bg-[#D2F38A]/10 p-2 py-1 text-[#385842]">$ETH</span> .
				</>
			)
		},
		{
			protocol: "plug",
			elements: (
				<>
					Has <span className="rounded-sm bg-[#D2F38A]/10 p-2 py-1 text-[#385842]">0.2</span>{" "}
					<span className="rounded-sm bg-[#D2F38A]/10 p-2 py-1 text-[#385842]">$ETH</span> or greater.
				</>
			)
		},
		{
			protocol: "ens",
			elements: (
				<>
					Can renew <span className="rounded-sm bg-[#D2F38A]/10 p-2 py-1 text-[#385842]">nftchance.eth</span>{" "}
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
					Run <span className="rounded-sm bg-[#D2F38A]/10 p-2 py-1 text-[#385842]">1</span> time a{" "}
					<span className="rounded-sm bg-[#D2F38A]/10 p-2 py-1 text-[#385842]">day</span> .
				</>
			)
		},
		{
			protocol: "plug",
			elements: (
				<>
					Run after <span className="rounded-sm bg-[#D2F38A]/10 p-2 py-1 text-[#385842]">10</span>{" "}
					<span className="rounded-sm bg-[#D2F38A]/10 p-2 py-1 text-[#385842]">PM</span>{" "}
					<span className="rounded-sm bg-[#D2F38A]/10 p-2 py-1 text-[#385842]">UTC</span> .
				</>
			)
		},
		{
			protocol: "plug",
			elements: (
				<>
					Run before <span className="rounded-sm bg-[#D2F38A]/10 p-2 py-1 text-[#385842]">12/31/2024</span> .
				</>
			)
		},
		{
			protocol: "uniswap",
			elements: (
				<>
					Swap <span className="rounded-sm bg-[#D2F38A]/10 p-2 py-1 text-[#385842]">36,000</span>{" "}
					<span className="rounded-sm bg-[#D2F38A]/10 p-2 py-1 text-[#385842]">$USDC</span> to{" "}
					<span className="rounded-sm bg-[#D2F38A]/10 p-2 py-1 text-[#385842]">$ETH</span> .
				</>
			)
		},
		{
			protocol: "nouns",
			elements: (
				<>
					Can bid with <span className="rounded-sm bg-[#D2F38A]/10 p-2 py-1 text-[#385842]">9</span>{" "}
					<span className="rounded-sm bg-[#D2F38A]/10 p-2 py-1 text-[#385842]">$ETH</span> .
				</>
			)
		},
		{
			protocol: "nouns",
			elements: (
				<>
					Bid on Noun with{" "}
					<span className="rounded-sm bg-[#D2F38A]/10 p-2 py-1 text-[#385842]">Pineapple Hat</span> .
				</>
			)
		}
	],
	[
		{
			protocol: "plug",
			elements: (
				<>
					Run <span className="rounded-sm bg-[#D2F38A]/10 p-2 py-1 text-[#385842]">1</span> time a{" "}
					<span className="rounded-sm bg-[#D2F38A]/10 p-2 py-1 text-[#385842]">hour</span> .
				</>
			)
		},
		{
			protocol: "yearn",
			elements: (
				<>
					<span className="rounded-sm bg-[#D2F38A]/10 p-2 py-1 text-[#385842]">$USDC</span> pool is above{" "}
					<span className="rounded-sm bg-[#D2F38A]/10 p-2 py-1 text-[#385842]">72%</span> APY.
				</>
			)
		},
		{
			protocol: "yearn",
			elements: (
				<>
					Can deposit <span className="rounded-sm bg-[#D2F38A]/10 p-2 py-1 text-[#385842]">10,000</span>{" "}
					<span className="rounded-sm bg-[#D2F38A]/10 p-2 py-1 text-[#385842]">$USDC</span> .
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
					className="bg-gradient-radial absolute inset-0 rounded-lg from-plug-green via-plug-yellow to-plug-white opacity-0"
					initial={{ opacity: 0 }}
					animate={{
						opacity: [0, 1, 1, 0],
						background: [
							"radial-gradient(circle at 0%  #385842 0%, #D2F38A 50%, #FFFFFF 100%)",
							"radial-gradient(circle at 100% 0%, #385842 0%, #D2F38A 50%, #FFFFFF 100%)",
							"radial-gradient(circle at 100% 100%, #385842 0%, #D2F38A 50%, #FFFFFF 100%)",
							"radial-gradient(circle at 0% 100%, #385842 0%, #d2f38a 50%, #FFFFFF 100%)",
							"radial-gradient(circle at 0% 0%, #D2F38A 0%, #385842 50%, #FFFFFF 100%)"
						]
					}}
					transition={{
						duration: 2,
						ease: "linear",
						delay: 0.1 + index * 2,
						repeat: Infinity
					}}
				/>

				<div className="relative z-10 m-[2px] flex w-full flex-row items-center gap-2 rounded-lg bg-plug-white p-3 lg:gap-4">
					<Image
						src={`/protocols/${protocol}.png`}
						alt={formatTitle(protocol)}
						width={48}
						height={48}
						className="h-6 w-6 rounded-sm"
					/>
					<p className="text-[14px] font-bold text-plug-green/40 lg:text-[16px]">{children}</p>
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
						className="h-full w-full bg-plug-green/10"
						initial={{ background: "linear-gradient(30deg, #EBECEC, #EBECEC)" }}
						animate={{ background: "linear-gradient(30deg, #D2F38A, #385842)" }}
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
			text='"If this, then that” execution.'
			description="Control every granular detail and squeeze the maximum value out of every transaction."
			className="relative z-[999] col-span-2 row-span-2 h-full min-h-[580px] xl:col-span-4"
		>
			<div className="flex h-[80%] w-full select-none flex-col items-center justify-center gap-2">
				<ExecutionActions />
			</div>

			<div className="absolute bottom-[30%] left-0 right-0 top-[60%] bg-gradient-to-b from-plug-white/0 to-plug-white" />
			<div className="absolute bottom-0 left-0 right-0 top-[70%] bg-plug-white" />
		</InfoCard>
	)
}
