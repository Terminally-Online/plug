import { useEffect, useState } from "react"

import Image from "next/image"
import Link from "next/link"

import { motion } from "framer-motion"
import { Book, Twitter } from "lucide-react"

import { Button, LandingContainer } from "@/components"
import { routes } from "@/lib"

import { HeroShapes } from "./shapes"

const EARLY_ACCESS = process.env.NEXT_PUBLIC_EARLY_ACCESS === "false" ? false : true

const totalBars = 100
const variants = {
	okayTrader: (index: number) => {
		const baseValue = 15
		const tradingPeriods = [20, 40, 60, 80]
		const tradingResults = [-10, 15, -5, 20]

		let result = baseValue
		for (let i = 0; i < tradingPeriods.length; i++) {
			if (index >= tradingPeriods[i]) {
				result += tradingResults[i]
			}
		}

		const noise = Math.random() * 5 - 1
		result += noise

		return Math.max(0, Math.min(100, result))
	},
	goodTrader: (index: number) => {
		const baseGrowth = (index / (totalBars - 1)) * 45
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
		const previousGrowth = index > 0 ? variants.supernaturalTrader(index - 1) : baseGrowth
		const smoothedGrowth = previousGrowth * (1 - smoothingFactor) + baseGrowth * smoothingFactor

		const noise = Math.random() * 0.5 - 0.25
		return Math.min(100, Math.max(0, smoothedGrowth + noise))
	}
}

const HeroBarChart = () => {
	const [currentVariant, setCurrentVariant] = useState<keyof typeof variants>("okayTrader")

	useEffect(() => {
		const interval = setInterval(() => {
			setCurrentVariant(prev => {
				const variantKeys = Object.keys(variants) as Array<keyof typeof variants>
				const currentIndex = variantKeys.indexOf(prev)
				return variantKeys[(currentIndex + 1) % variantKeys.length]
			})
		}, 5000)

		return () => clearInterval(interval)
	}, [])

	return (
		<div className="z-2 absolute bottom-0 left-0 right-0 top-0 -mx-8 flex flex-row justify-between">
			{Array.from({ length: totalBars }).map((_, index) => {
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
			})}
		</div>
	)
}

export const Hero = () => (
	<div className="relative z-[2] flex h-full min-h-screen w-screen">
		<HeroShapes />
		<HeroBarChart />

		<div className="z-2 relative w-full">
			<LandingContainer className="flex h-full flex-col py-8 text-white">
				<div className="flex flex-row items-center gap-4">
					<Link className="mr-4" href={routes.index}>
						<Image src="/white-icon.svg" alt="Logo" width={24} height={24} />
					</Link>
					<a href={routes.documentation} target="_blank" rel="noreferrer">
						<Book size={18} className="opacity-80 transition-opacity duration-200 hover:opacity-100" />
					</a>
					<a href={routes.twitter} target="_blank" rel="noreferrer">
						<Twitter size={18} className="opacity-80 transition-opacity duration-200 hover:opacity-100" />
					</a>

					<Button
						variant="none"
						className="ml-auto w-max rounded-md border-[1px] border-white/20 bg-white/20 px-4 py-2 text-center text-sm font-black text-white"
						href={EARLY_ACCESS ? routes.earlyAccess : routes.app}
					>
						Enter App
					</Button>
				</div>

				<div className="my-auto flex min-h-[calc(100vh-180px)] items-center pb-6">
					<div className="my-auto flex flex-col gap-16">
						<motion.h1
							className="max-w-[75%] text-[3.5rem] font-black leading-tight text-white md:text-[72px] lg:text-[96px]"
							initial={{ y: 20, opacity: 0 }}
							animate={{
								y: [0, 20],
								opacity: [0, 1]
							}}
							transition={{ duration: 0.3 }}
						>
							Supernatural returns with unparalleled control.
						</motion.h1>

						<motion.p
							className="max-w-[52%] text-[1.25rem] font-bold text-white/80 md:text-[24px]"
							initial={{ y: -20, opacity: 0 }}
							animate={{
								y: [0, -20],
								opacity: [0, 1]
							}}
							transition={{
								duration: 0.3,
								delay: 0.15
							}}
						>
							Use the blockchain like never before with an &ldquo;if this, then that&rdquo; platform that
							enables you to set rules so that everything executes instantly whether you&apos;re at the
							computer or not.
						</motion.p>

						<Button
							variant="none"
							className="w-max rounded-md border-[1px] border-white/30 bg-white/20 px-8 py-3 text-center font-black text-white"
							href={EARLY_ACCESS ? routes.earlyAccess : routes.app}
						>
							Enter App
						</Button>
					</div>
				</div>
			</LandingContainer>
		</div>
	</div>
)
