import { useEffect, useMemo, useState } from "react"

import { motion } from "framer-motion"

import { useMediaQuery } from "@/lib"

const createVariants = (totalBars: number) => ({
	okayTrader: (index: number) => {
		const baseValue = 10
		const tradingPeriods = [20, 40, 60, 80]
		const tradingResults = [-10, 15, -5, 20]
		let result = baseValue
		let left = 0
		let right = tradingPeriods.length - 1
		while (left <= right) {
			const mid = Math.floor((left + right) / 2)
			if (index >= tradingPeriods[mid]) {
				result += tradingResults[mid]
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
		const noise = (Math.random() - 0.5) * 5
		return Math.max(0, Math.min(100, result + noise))
	},
	goodTrader: (index: number) => {
		const baseGrowth = (index / (totalBars - 1)) * 40
		const volatility = Math.sin(index * 0.2) * 10
		const noise = Math.random() * 4 - 1
		return Math.max(0, Math.min(100, baseGrowth + volatility + noise))
	},
	exceptionalTrader: (index: number) => {
		const baseGrowth =
			((Math.exp(index / (totalBars / 2.5)) - 1) / (Math.exp((totalBars - 1) / (totalBars / 2.5)) - 1)) * 100
		const noise = Math.random() * 3 - 1.5
		return Math.max(0, Math.min(100, baseGrowth + noise))
	},
	supernaturalTrader: (index: number): number => {
		const growthRate = 0.05
		const baseGrowth = ((Math.exp(growthRate * index) - 1) / (Math.exp(growthRate * totalBars) - 1)) * 150
		const smoothingFactor = 0.2
		const previousGrowth = index > 0 ? createVariants(totalBars).supernaturalTrader(index - 1) : baseGrowth
		const smoothedGrowth = previousGrowth * (1 - smoothingFactor) + baseGrowth * smoothingFactor
		const noise = Math.random() * 0.5 - 0.25
		return Math.min(100, Math.max(0, smoothedGrowth + noise))
	}
})

export const HeroBarChart = () => {
	const { md } = useMediaQuery()

	const [currentVariant, setCurrentVariant] = useState<keyof ReturnType<typeof createVariants>>("okayTrader")

	const totalBars = md ? 100 : 60

	const variants = useMemo(() => createVariants(totalBars), [totalBars])

	useEffect(() => {
		const interval = setInterval(() => {
			setCurrentVariant(prev => {
				const variantKeys = Object.keys(variants) as Array<keyof typeof variants>
				const currentIndex = variantKeys.indexOf(prev)
				return variantKeys[(currentIndex + 1) % variantKeys.length]
			})
		}, 5000)

		return () => clearInterval(interval)
	}, [variants])

	const bars = useMemo(
		() =>
			Array.from({ length: totalBars }).map((_, index) => {
				const height = variants[currentVariant](index)
				const opacity = (index / (totalBars - 1)) * 0.4

				return (
					<motion.div
						key={index}
						className="mt-auto w-1 rounded-full bg-gradient-to-t from-white/0 to-white"
						style={{ opacity }}
						initial={{ height: 0 }}
						animate={{ height: `${Math.max(1, height)}%` }}
						transition={{
							duration: 1,
							delay: index * 0.03,
							ease: "easeOut"
						}}
					/>
				)
			}),
		[currentVariant, totalBars, variants]
	)

	return <div className="z-2 absolute bottom-0 left-0 right-0 top-0 -mx-8 flex flex-row justify-between">{bars}</div>
}
