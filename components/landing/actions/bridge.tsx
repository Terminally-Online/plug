import { motion } from "framer-motion"
import { CalendarClock } from "lucide-react"

import { InfoCard } from "../cards"

const LINES = 40

export const ActionBridge = () => {
	return (
		<InfoCard
			icon={<CalendarClock size={24} className="opacity-40" />}
			text="Bridge."
			description="Move your crypto quickly between chains."
			className="relative z-[99999] h-[320px] sm:h-[320px] 2xl:h-[300px]"
		>
			<div className="absolute bottom-1/2 left-0 right-0 top-0 flex">
				{Array.from({ length: LINES }).map((_, index) => (
					<motion.div
						key={index}
						className="absolute mr-auto h-[2px] w-24 bg-gradient-to-r from-plug-yellow/40 to-plug-yellow"
						initial={{
							top: `${Math.random() * 100}%`
						}}
						animate={{
							left: 0,
							x: ["-100%", "200%"]
						}}
						transition={{
							duration: 0.5,
							repeat: Infinity,
							ease: "easeInOut",
							delay: Math.random()
						}}
					/>
				))}
			</div>

			<div className="absolute bottom-[50%] left-0 right-0 top-0 bg-gradient-to-b from-plug-white/0 to-plug-white" />
			<div className="absolute bottom-0 left-0 right-0 top-[50%] bg-plug-white" />
		</InfoCard>
	)
}
