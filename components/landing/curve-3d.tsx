import { useRef } from "react"

import { motion, useScroll, useTransform } from "framer-motion"
import { CalendarClock } from "lucide-react"

import { ActionBorrow, ActionBridge, ActionDiscover, ActionStaking, ActionSwap } from "./actions"
import { ActionLiquidity } from "./actions/liquidity"
import { InfoCard } from "./cards"

export const Curve3D = () => {
	const containerRef = useRef<HTMLDivElement>(null)
	const { scrollYProgress } = useScroll({
		target: containerRef,
		offset: ["start end", "end start"]
	})

	const pathLength = useTransform(scrollYProgress, [0, 0.35], [0, 1])

	const diagonal = Math.sqrt(Math.pow(2800, 2) + Math.pow(1400, 2))
	const maxRadius = diagonal / 2

	const circleRadius = useTransform(scrollYProgress, [0.34, 0.35, 0.6], [0, 30, maxRadius])

	const textOpacity = useTransform(scrollYProgress, [0.5, 0.55], [0, 1])

	return (
		<div className="relative h-screen w-full" ref={containerRef}>
			<div className="absolute inset-0 z-[99999] overflow-visible">
				<svg
					viewBox="0 0 1827 976"
					fill="none"
					xmlns="http://www.w3.org/2000/svg"
					className="h-full w-full overflow-visible"
				>
					<motion.path
						style={{ pathLength }}
						d="M0 251C234 385 219.145 251 508.814 251C732.341 251 718.957 473.761 913 486.5"
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
				className="absolute inset-0 z-[100000] flex flex-col items-center justify-center gap-12 pt-[60%]"
				style={{
					opacity: textOpacity
				}}
			>
				<div className="max-w-[720px] text-center">
					<h2 className="mb-4 text-[52px] font-black text-[#385842]">
						Every common crypto usecase at your fingertips.
					</h2>
					<p className="text-xl font-bold text-plug-green/40">Scroll down to explore more</p>
				</div>

				<div className="grid max-w-[1200px] grid-cols-4 gap-4">
					<InfoCard
						icon={<CalendarClock size={24} className="opacity-40" />}
						text="Liquidity."
						description="Stake and manage your crypto assets."
						className="relative z-[99999] h-[320px] sm:h-[320px] 2xl:h-[300px]"
					>
						<div className="absolute bottom-[50%] left-0 right-0 top-0 bg-gradient-to-b from-plug-white/0 to-plug-white" />
						<div className="absolute bottom-0 left-0 right-0 top-[50%] bg-plug-white" />
					</InfoCard>

					<ActionBorrow />
					<ActionStaking />
					<ActionDiscover />
					<ActionBridge />
					<ActionSwap />
					<ActionLiquidity />
					<InfoCard
						icon={<CalendarClock size={24} className="opacity-40" />}
						text="Harvest."
						description="Move your crypto quickly between chains."
						className="relative z-[99999] h-[320px] sm:h-[320px] 2xl:h-[300px]"
					>
						<div className="absolute bottom-[50%] left-0 right-0 top-0 bg-gradient-to-b from-plug-white/0 to-plug-white" />
						<div className="absolute bottom-0 left-0 right-0 top-[50%] bg-plug-white" />
					</InfoCard>
				</div>
			</motion.div>
		</div>
	)
}

export default Curve3D
