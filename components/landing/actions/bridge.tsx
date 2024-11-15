import { motion } from "framer-motion"
import { BringToFront } from "lucide-react"

import { InfoCard } from "@/components"

export const ActionBridge = () => {
	const leftPoints = Array(18).fill(null)
	const centerPoints = Array(1).fill(null)
	const rightPoints = Array(18).fill(null)

	return (
		<InfoCard
			icon={<BringToFront size={24} className="opacity-40" />}
			text="Bridge."
			description="Move your crypto quickly between chains."
			className="relative z-[99999] col-span-2 h-[320px] overflow-hidden sm:h-[320px] xl:col-span-1 2xl:h-[300px]"
		>
			<div className="absolute inset-0 top-2 h-1/2 w-full">
				{/* Changed the container to stretch beyond boundaries */}
				<div className="absolute left-[-10%] right-[-10%] h-full w-[120%]">
					<svg
						viewBox="0 0 1200 300"
						className="mt-[-5%] h-full w-full"
						preserveAspectRatio="xMidYMid slice"
						style={{ overflow: "visible" }}
					>
						<defs>
							<linearGradient id="flow-gradient" x1="0" y1="0" x2="1" y2="0">
								<stop offset="0%" stopColor="#D2F38A" />
								<stop offset="100%" stopColor="#D2F38A" />
							</linearGradient>
							<linearGradient id="flow-gradient-reversed" x1="0" y1="0" x2="1" y2="0">
								<stop offset="0%" stopColor="#D2F38A" />
								<stop offset="100%" stopColor="#D2F38A" />
							</linearGradient>
						</defs>
						{leftPoints.map((_, i) =>
							centerPoints.map((_, j) => (
								<motion.path
									key={`flow-${i}-${j}`}
									d={`M 0 ${i * 40} C 300 ${i * 40}, 420 ${j * 45}, 600 ${j * 45}`}
									stroke="url(#flow-gradient)"
									strokeWidth="16"
									fill="none"
									initial={{ pathLength: 0 }}
									animate={{ pathLength: [0, 1, 0] }}
									transition={{
										duration: 2,
										delay: i * 0.1 + j * 0.1,
										repeat: Infinity,
										repeatDelay: i * 0.1 + j * 0.1,
										ease: "easeIn"
									}}
								/>
							))
						)}
						{rightPoints.map((_, i) =>
							centerPoints.map((_, j) => (
								<motion.path
									key={`flow-right-${i}-${j}`}
									d={`M 600 ${j * 45} C 780 ${j * 45}, 900 ${i * 40}, 1200 ${i * 40}`}
									stroke="url(#flow-gradient-reversed)"
									strokeWidth="16"
									fill="none"
									initial={{ pathLength: 0 }}
									animate={{ pathLength: [0, 1, 0] }}
									transition={{
										duration: 2,
										delay: 2 + i * 0.1 + j * 0.1,
										repeat: Infinity,
										repeatDelay: 2 + i * 0.1 + j * 0.1,
										ease: "easeOut"
									}}
								/>
							))
						)}
					</svg>
				</div>
			</div>
			<div className="absolute bottom-[50%] left-0 right-0 top-0 bg-gradient-to-b from-plug-white/0 to-plug-white" />
			<div className="absolute bottom-0 left-0 right-0 top-[50%] bg-plug-white" />
		</InfoCard>
	)
}

export default ActionBridge
