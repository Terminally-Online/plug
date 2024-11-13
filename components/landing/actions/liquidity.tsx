import React from "react"

import { motion } from "framer-motion"
import { CalendarClock } from "lucide-react"

import { InfoCard } from "../cards"

const generatePoints = (width: number, height: number, segments = 10) => {
	const points = []
	for (let i = 0; i <= segments; i++) {
		const x = (i / segments) * width
		points.push(x)
	}
	return points
}

export const ActionLiquidity = () => {
	const width = 500
	const height = 200
	const points = generatePoints(width, height)

	const createPath = (values: number[]) => {
		const points = values.map((v, i) => `${(i * width) / 10} ${v}`).join(" L ")
		return `M 0 50 L ${points} L ${width} 200 L 0 200 Z`
	}

	return (
		<InfoCard
			icon={<CalendarClock size={24} className="opacity-40" />}
			text="Provide Liquidity."
			description="Deposit your crypto into liquidity pools to earn swap fees and yield."
			className="relative z-[99999] col-span-2 h-[320px] sm:h-[320px] 2xl:h-[300px]"
		>
			<svg className="absolute inset-0 h-full w-full" viewBox="0 0 500 200" preserveAspectRatio="none">
				<motion.path
					fill="#D2F38A"
					initial={{
						d: createPath(points.map(() => 50))
					}}
					animate={{
						d: [
							createPath(points.map((_, i) => 50 + Math.sin(i * 0.5) * 30)),
							createPath(points.map((_, i) => 50 + Math.cos(i * 0.5) * 40)),
							createPath(points.map((_, i) => 50 + Math.sin(i * 0.8) * 35))
						]
					}}
					transition={{
						duration: 6,
						repeat: Infinity,
						repeatType: "reverse",
						ease: "easeInOut",
						times: [0, 0.5, 1]
					}}
				/>
			</svg>
			<div className="absolute bottom-[50%] left-0 right-0 top-0 bg-gradient-to-b from-plug-white/0 to-plug-white" />
			<div className="absolute bottom-0 left-0 right-0 top-[50%] bg-plug-white" />
		</InfoCard>
	)
}

export default ActionLiquidity
