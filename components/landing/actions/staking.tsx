import { useEffect, useState } from "react"

import { motion, useAnimationControls } from "framer-motion"
import { CalendarClock } from "lucide-react"

import { InfoCard } from "@/components"

export const ActionStaking = () => {
	const controls = useAnimationControls()

	const [duration, setDuration] = useState(6)
	const [loopCount, setLoopCount] = useState(0)

	useEffect(() => {
		const animate = async () => {
			while (true) {
				await controls.start({
					top: "100%",
					transition: {
						duration: duration,
						ease: "easeIn"
					}
				})

				controls.set({ top: "-24rem" })

				setLoopCount(prev => prev + 1)

				if (loopCount >= 20) {
					setDuration(2)
					setLoopCount(0)
				} else {
					setDuration(prev => Math.max(prev * 0.9, 0.4))
				}
			}
		}

		animate()
	}, [controls, loopCount, duration])

	return (
		<InfoCard
			icon={<CalendarClock size={24} className="opacity-40" />}
			text="Staking."
			description="Earn rewards by staking your tokens in pools."
			className="relative z-[99999] row-span-2 overflow-hidden"
		>
			<div className="absolute inset-0">
				<motion.div
					animate={controls}
					className="absolute left-1/2 top-[-24rem] flex h-24 w-24 -translate-x-1/2 items-center justify-center rounded-full border-[2px] border-dashed border-plug-green bg-plug-yellow font-bold text-plug-green"
				>
					<p>$USDC</p>
				</motion.div>
			</div>
			<div className="absolute bottom-[20%] left-0 right-0 top-0 bg-gradient-to-b from-plug-white/0 to-plug-white" />
			<div className="absolute bottom-0 left-0 right-0 top-[80%] bg-plug-white" />
		</InfoCard>
	)
}

export default ActionStaking
