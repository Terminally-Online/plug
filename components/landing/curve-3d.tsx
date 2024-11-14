import { useRef } from "react"

import { motion, useScroll, useTransform } from "framer-motion"

import { useMediaQuery } from "@/lib"

import {
	ActionBorrow,
	ActionBridge,
	ActionDiscover,
	ActionEarn,
	ActionStaking,
	ActionSwap,
	ActionTrade
} from "./actions"
import { ActionLiquidity } from "./actions/liquidity"

export const Curve3D = () => {
	const containerRef = useRef<HTMLDivElement>(null)
	const { xl } = useMediaQuery()
	const { scrollYProgress } = useScroll({
		target: containerRef,
		offset: ["start end", "end start"]
	})

	const diagonal = Math.sqrt(Math.pow(2800, 2) + Math.pow(1400, 2))
	const maxRadius = diagonal / 2

	const pathLength = useTransform(scrollYProgress, [0, 0.35], [0, 1])
	const circleRadius = useTransform(scrollYProgress, [0.34, 0.35, 0.6], [0, 30, maxRadius])
	const textOpacity = useTransform(scrollYProgress, [0.5, 0.55], [0, 1])

	return (
		<div className="relative w-full xl:h-screen" ref={containerRef}>
			<div className="absolute inset-0 z-[99999] mt-24 hidden overflow-visible xl:flex">
				<svg viewBox="0 0 1827 976" fill="none" className="absolute inset-0 overflow-visible">
					<path
						d="M-18 251C307.5 398 263.104 121.272 556 159.5C816.5 193.5 715.132 473.288 913 486"
						stroke="url(#paint0_linear_4614_166)"
						stroke-width="60"
						stroke-linecap="round"
					/>
					<motion.path
						d="M-18 251C307.5 398 263.104 121.272 556 159.5C816.5 193.5 715.132 473.288 913 486L952.5 489"
						stroke="#FEFFF7"
						stroke-width="60"
						stroke-dasharray="4 4"
						animate={{ strokeDashoffset: [60, 0] }}
						transition={{
							duration: 0.5,
							repeat: Infinity,
							ease: "linear"
						}}
					/>
					<motion.path
						style={{ pathLength }}
						d="M-18 251C307.5 398 263.104 121.272 556 159.5C816.5 193.5 715.132 473.288 913 486"
						stroke="url(#paint0_linear_4614_166)"
						stroke-width="60"
						stroke-linecap="round"
					/>
					<motion.circle cx="913" cy="486.5" fill="#D2F38A" r={circleRadius} className="overflow-visible" />
					<defs>
						<linearGradient
							id="paint0_linear_4614_166"
							x1="774.577"
							y1="456.812"
							x2="61.8676"
							y2="240.295"
							gradientUnits="userSpaceOnUse"
						>
							<stop stopColor="#D2F38A" />
							<stop offset="1" stopColor="#385842" />
						</linearGradient>
					</defs>
				</svg>
			</div>

			<motion.div
				className="absolute inset-0 z-[100000] flex flex-col items-center justify-center gap-12 xl:pt-[60%]"
				style={{
					opacity: xl ? textOpacity : 1
				}}
			>
				<div className="text-center xl:max-w-[720px]">
					<h2 className="mb-4 text-[52px] font-black text-[#385842]">
						Every common crypto usecase on autopilot.
					</h2>
					<p className="mx-auto max-w-[480px] text-xl font-bold text-plug-green/40">
						When using Plug, you have everything at your fingertips only a few clicks away.
					</p>
				</div>

				<div className="grid max-w-[1200px] grid-cols-2 gap-4 xl:grid-cols-4">
					<ActionEarn />
					<ActionBorrow />
					<ActionStaking />
					<ActionDiscover />
					<ActionTrade />
					<ActionSwap />
					<ActionBridge />
					<ActionLiquidity />
				</div>
			</motion.div>
		</div>
	)
}

export default Curve3D
