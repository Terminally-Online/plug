import { useEffect, useState } from "react"

import { motion, useAnimationFrame } from "framer-motion"
import { HandCoins } from "lucide-react"

import { Counter, InfoCard } from "@/components"

export const ActionBorrow = () => {
	const [percentage, setPercentage] = useState(25)
	const [activeIndex, setActiveIndex] = useState(0)
	const widthValues = [80, 90, 50, 70, 80, 90, 40, 20]
	const animationDuration = 12000

	useAnimationFrame(time => {
		const progress = (time % animationDuration) / animationDuration
		const index = Math.floor(progress * widthValues.length)
		const currentWidth = widthValues[index]
		setPercentage(currentWidth)
	})

	// Calculate target index based on percentage
	const getTargetIndex = (percent: number) => {
		return Math.floor((percent / 100) * 24)
	}

	useEffect(() => {
		const targetIndex = getTargetIndex(percentage)

		// Animate pills sequentially
		const animatePills = async () => {
			if (targetIndex > activeIndex) {
				// Animate upwards
				for (let i = activeIndex; i <= targetIndex; i++) {
					setActiveIndex(i)
					await new Promise(resolve => setTimeout(resolve, 50)) // Adjust speed here
				}
			} else {
				// Animate downwards
				for (let i = activeIndex; i >= targetIndex; i--) {
					setActiveIndex(i)
					await new Promise(resolve => setTimeout(resolve, 50)) // Adjust speed here
				}
			}
		}

		animatePills()
	}, [percentage, activeIndex])

	const getIsActive = (index: number) => {
		return index <= activeIndex
	}

	return (
		<InfoCard
			icon={<HandCoins size={24} className="opacity-40" />}
			text="Borrow & Lend."
			description="Realize the full value of your onchain assets by supplying and borrowing with decentralized lending markets."
			className="relative z-[99999] col-span-2 h-[320px] sm:h-[320px] 2xl:h-[300px]"
		>
			<motion.p className="absolute left-0 right-0 top-2 z-[9999] flex w-full justify-between whitespace-nowrap px-8 pb-4 font-bold">
				<span className="opacity-40">Health Factor:</span>
				<span className="flex flex-row items-center pl-2">
					<Counter count={Math.round(percentage)} />%
				</span>
			</motion.p>

			<div className="absolute inset-0 bottom-1/2 mt-10 flex flex-row gap-2">
				{Array.from({ length: 24 }).map((_, index) => (
					<motion.div
						key={index}
						className="relative h-16 w-full rounded-full md:h-32"
						animate={{
							backgroundColor: getIsActive(index) ? "#D2F38A" : "#FF0000",
							opacity: getIsActive(index) ? 1 : 0.6
						}}
						transition={{
							duration: 0.2,
							ease: "easeInOut"
						}}
					/>
				))}
			</div>

			<div className="absolute bottom-[50%] left-0 right-0 top-0 bg-gradient-to-b from-plug-white/0 to-plug-white" />
			<div className="absolute bottom-0 left-0 right-0 top-[50%] bg-plug-white" />
		</InfoCard>
	)
}
