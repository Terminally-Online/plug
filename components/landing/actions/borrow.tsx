import { useState } from "react"

import { motion, useAnimationFrame } from "framer-motion"
import { CalendarClock } from "lucide-react"

import { Counter } from "@/components/shared"

import { InfoCard } from "../cards"

export const ActionBorrow = () => {
	const [percentage, setPercentage] = useState(25)

	const widthValues = [80, 90, 50, 70, 80, 90, 40, 20]
	const animationDuration = 6000

	useAnimationFrame(time => {
		const progress = (time % animationDuration) / animationDuration
		const index = Math.floor(progress * widthValues.length)
		const currentWidth = widthValues[index]
		setPercentage(currentWidth)
	})
	return (
		<InfoCard
			icon={<CalendarClock size={24} className="opacity-40" />}
			text="Borrow & Lend."
			description="Realize the full value of your onchain assets by supplying and borrowing with decentralized lending markets."
			className="relative z-[99999] col-span-2 h-[320px] sm:h-[320px] 2xl:h-[300px]"
		>
			<div className="absolute inset-0 flex flex-row">
				<motion.p
					className="absolute top-2 ml-[28px] h-24 -translate-x-1/2 border-l-[2px] border-plug-green pb-4 font-bold"
					animate={{
						left: ["80%", "90%", "50%", "70%", "80%", "90%", "40%", "20%"]
					}}
					transition={{
						duration: 6,
						repeat: Infinity,
						repeatType: "reverse",
						ease: "easeInOut"
					}}
				>
					<span className="flex flex-row items-center pl-4">
						<Counter count={Math.round(percentage)} />%
					</span>
				</motion.p>
				<motion.div
					className="relative mt-[10%] flex h-24 items-center justify-center bg-plug-yellow"
					animate={{
						width: ["80%", "90%", "50%", "70%", "80%", "90%", "40%", "20%"]
					}}
					transition={{
						duration: 6,
						repeat: Infinity,
						repeatType: "reverse",
						ease: "easeInOut"
					}}
				>
					<p className="text-2xl font-black text-plug-green opacity-40">PROFIT</p>
				</motion.div>
				<motion.div
					className="relative mt-[10%] flex h-24 w-1/2 items-center justify-center bg-plug-red"
					animate={{
						width: ["20%", "10%", "50%", "30%", "20%", "10%", "60%", "80%"]
					}}
					transition={{
						duration: 6,
						repeat: Infinity,
						repeatType: "reverse",
						ease: "easeInOut"
					}}
				>
					<p className="text-2xl font-black text-white opacity-40">RISK</p>
				</motion.div>
			</div>

			<div className="absolute bottom-[50%] left-0 right-0 top-0 bg-gradient-to-b from-plug-white/0 to-plug-white" />
			<div className="absolute bottom-0 left-0 right-0 top-[50%] bg-plug-white" />
		</InfoCard>
	)
}
