import { motion, useAnimationControls } from "framer-motion"
import { CalendarClock } from "lucide-react"

import { InfoCard } from "@/components"

export const ActionStaking = () => {
	return (
		<InfoCard
			icon={<CalendarClock size={24} className="opacity-40" />}
			text="Stake."
			description="Earn rewards by staking your tokens in pools."
			className="relative z-[99999] row-span-2 overflow-hidden"
		>
			<div className="absolute inset-0 bottom-1/2 overflow-hidden">
				<motion.div
					className="absolute left-1/2 top-[-24rem] flex h-24 w-24 -translate-x-1/2 items-center justify-center rounded-full border-[2px] border-dashed border-plug-green/40 font-bold text-plug-green"
					animate={{
						top: ["-35%", "100%"]
					}}
					transition={{
						duration: 4,
						repeat: Infinity,
						repeatDelay: 3,
						ease: "linear"
					}}
				>
					<p className="relative">$DAI</p>
				</motion.div>
			</div>
			<div className="absolute inset-0 top-1/2 z-[999] h-[4px] overflow-hidden bg-plug-yellow" />
			<div className="absolute inset-0 top-1/2 overflow-hidden">
				<motion.div
					className="absolute left-1/2 top-[-24rem] flex h-24 w-24 -translate-x-1/2 items-center justify-center rounded-full border-[2px] border-dashed border-plug-green bg-plug-yellow font-bold text-plug-green"
					animate={{
						top: ["-35%", "100%"]
					}}
					transition={{
						duration: 4,
						repeat: Infinity,
						repeatDelay: 3,
						ease: "linear",
						delay: 3
					}}
				>
					<p className="relative">$sDAI</p>
				</motion.div>
			</div>

			<div className="absolute bottom-[20%] left-0 right-0 top-0 bg-gradient-to-b from-plug-white/0 to-plug-white" />
			<div className="absolute bottom-0 left-0 right-0 top-[80%] bg-plug-white" />
		</InfoCard>
	)
}

export default ActionStaking
