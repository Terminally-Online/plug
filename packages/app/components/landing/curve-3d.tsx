import { useRef } from "react"

import { motion, useScroll, useTransform } from "framer-motion"

import { ActionBorrow } from "@/components/landing/actions/borrow"
import { ActionBridge } from "@/components/landing/actions/bridge"
import { ActionDiscover } from "@/components/landing/actions/discover"
import { ActionEarn } from "@/components/landing/actions/earn"
import { ActionTrade } from "@/components/landing/actions/harvest"
import { ActionLiquidity } from "@/components/landing/actions/liquidity"
import { ActionStaking } from "@/components/landing/actions/staking"
import { ActionSwap } from "@/components/landing/actions/swap"
import { useMediaQuery } from "@/lib"

export const Curve3D = () => {
	const containerRef = useRef<HTMLDivElement>(null)
	const { xl } = useMediaQuery()
	const { scrollYProgress } = useScroll({
		target: containerRef,
		offset: ["start end", "end start"]
	})

	const diagonal = Math.sqrt(Math.pow(xl ? 2800 : 3200, 2) + Math.pow(xl ? 1400 : 1600, 2))
	const maxRadius = diagonal / 2

	const pathLength = useTransform(scrollYProgress, [0, 0.35], [0, 1])
	const circleRadius = useTransform(scrollYProgress, [0.34, 0.35, 0.6], [0, 30, maxRadius])
	const textOpacity = useTransform(scrollYProgress, [0.5, 0.55], [0, 1])

	return (
		<div
			className="relative mb-12 h-full w-full bg-plug-yellow py-24 xl:h-screen xl:bg-plug-white xl:py-0"
			ref={containerRef}
		>
			<div className="absolute inset-0 z-[99999] mt-24 hidden overflow-visible xl:flex">
				<svg viewBox="0 0 1827 976" fill="none" className="absolute inset-0 overflow-visible">
					<path
						d="M-18 251C307.5 398 263.104 121.272 556 159.5C816.5 193.5 715.132 473.288 913 486"
						stroke="url(#paint0_linear_4614_166)"
						strokeWidth="60"
						strokeLinecap="round"
					/>
					<motion.path
						d="M-18 251C307.5 398 263.104 121.272 556 159.5C816.5 193.5 715.132 473.288 913 486L952.5 489"
						stroke="#FEFFF7"
						strokeWidth="60"
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
						strokeWidth="60"
						strokeLinecap="round"
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
				className="flex flex-col items-center justify-center gap-12 px-8 text-plug-green lg:px-24 xl:absolute xl:inset-0 xl:z-[100000] xl:pt-[50%] 2xl:pt-[60%]"
				style={{
					opacity: xl ? textOpacity : 1
				}}
			>
				<div className="mr-auto max-w-[720px] xl:mx-auto xl:text-center">
					<h2 className="mb-4 text-[52px] font-black leading-tight text-[#385842]">
						Every common crypto usecase on autopilot.
					</h2>
					<p className="mr-auto max-w-[480px] text-xl font-bold text-plug-green/40 xl:mx-auto">
						When using Plug, you have everything at your fingertips only a few clicks away.
					</p>
				</div>

				<div className="grid w-full grid-cols-2 gap-2 xl:max-w-[1200px] xl:grid-cols-4">
					<ActionEarn />
					<ActionLiquidity />
					<ActionStaking />
					<ActionDiscover />
					<ActionTrade />
					<ActionSwap />
					<ActionBridge />
					<ActionBorrow />
				</div>
			</motion.div>
		</div>
	)
}
