import { CalendarClock } from "lucide-react"
import { InfoCard } from "@/components"
import { motion, useAnimationControls } from "framer-motion"
import { useEffect, useState } from "react"

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
						ease: "easeIn",
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
					className="absolute top-[-24rem] w-24 h-24 rounded-full left-1/2 bg-plug-yellow -translate-x-1/2 flex items-center justify-center text-plug-green font-bold"
				>
					<p>$USDC</p>
				</motion.div>
			</div>
			<div className="absolute bottom-[20%] left-0 right-0 top-0 bg-gradient-to-b from-plug-white/0 to-plug-white" />
			<div className="absolute bottom-0 left-0 right-0 top-[80%] bg-plug-white" />
		</InfoCard>
	)
}

export default ActionStaking;
