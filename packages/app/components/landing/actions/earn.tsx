import { FC, useEffect, useState } from "react"

import { motion } from "framer-motion"
import { Handshake } from "lucide-react"

import { InfoCard } from "@/components"

interface Circle {
	xPercent: number // Changed from x to xPercent
	yPercent: number // Changed from y to yPercent
	radius: number
}

interface Token extends Circle {
	id: number
	rotation: number
	size: number
}

const distance = (a: Circle, b: Circle): number => {
	// Convert percentage distances to a 0-100 scale for calculation
	return Math.sqrt(Math.pow(a.xPercent - b.xPercent, 2) + Math.pow(a.yPercent - b.yPercent, 2))
}

const circleCollides = (circle: Circle, circles: Circle[], minDistance: number): boolean => {
	for (let other of circles) {
		const dist = distance(circle, other)
		if (dist < minDistance) return true
	}
	return false
}

const generatePackedCircles = (count: number, baseRadius: number): (Circle & { size: number })[] => {
	const circles: (Circle & { size: number })[] = []
	const maxAttempts = 150

	// Helper function to get hexagonal grid position in percentages
	const getHexPosition = (row: number, col: number): { xPercent: number; yPercent: number } => {
		const hexWidth = 15 // Percentage of container width
		const hexHeight = 15 // Percentage of container height

		return {
			xPercent: col * hexWidth + (row % 2) * (hexWidth / 2),
			yPercent: row * hexHeight
		}
	}

	// Calculate rows and columns based on count
	const rows = Math.ceil(Math.sqrt(count))
	const cols = Math.ceil(count / rows)

	let placed = 0

	// Try placing tokens in a hexagonal pattern
	for (let row = 0; row < rows && placed < count; row++) {
		for (let col = 0; col < cols && placed < count; col++) {
			let attempts = 0
			let circle: Circle & { size: number }

			do {
				const sizeMultiplier = 0.7 + Math.random() * 0.6
				const tokenSize = baseRadius * sizeMultiplier

				// Get base hexagonal position
				const hexPos = getHexPosition(row, col)

				// Add randomization to the position
				const jitter = 5 // 5% jitter
				circle = {
					xPercent: hexPos.xPercent + (Math.random() - 0.5) * jitter,
					yPercent: hexPos.yPercent + (Math.random() - 0.5) * jitter,
					radius: tokenSize,
					size: tokenSize * 2
				}

				// Spread tokens out more towards the top
				const spreadFactor = 1 - circle.yPercent / 100
				circle.xPercent = 50 + (circle.xPercent - 50) * (1 + spreadFactor)

				// Ensure tokens stay within bounds (with some overflow allowed)
				circle.xPercent = Math.max(-10, Math.min(110, circle.xPercent))
				circle.yPercent = Math.max(-10, Math.min(110, circle.yPercent))

				attempts++
			} while (circleCollides(circle, circles, 10) && attempts < maxAttempts)

			if (attempts < maxAttempts) {
				circles.push(circle)
				placed++
			}
		}
	}

	// Fill any remaining gaps with random positions
	while (placed < count) {
		let attempts = 0
		let circle: Circle & { size: number }

		do {
			const sizeMultiplier = 1.5 + Math.random()
			const tokenSize = baseRadius * sizeMultiplier

			circle = {
				xPercent: Math.random() * 120 - 10, // -10% to 110%
				yPercent: Math.random() * 120 - 10, // -10% to 110%
				radius: tokenSize,
				size: tokenSize * 2
			}
			attempts++
		} while (circleCollides(circle, circles, 10) && attempts < maxAttempts)

		if (attempts < maxAttempts) {
			circles.push(circle)
			placed++
		} else {
			break
		}
	}

	return circles
}

const BASE_TOKEN_RADIUS = 32
const TOKENS = 40
const SYMBOLS = ["ETH", "BTC", "SOL", "AVAX", "MATIC", "UNI", "AAVE", "SUSHI", "YFI", "COMP", "MKR", "CRV"]

export const ActionEarn: FC = () => {
	const [packedTokens, setPackedTokens] = useState<Token[]>([])

	useEffect(() => {
		const positions = generatePackedCircles(TOKENS, BASE_TOKEN_RADIUS)
		const tokens = positions.map((pos, index) => ({
			...pos,
			id: index,
			rotation: -15 + Math.random() * 30
		}))
		setPackedTokens(tokens)
	}, [])

	return (
		<InfoCard
			icon={<Handshake size={24} className="opacity-40" />}
			text="Earn."
			description="Earn yield and creator rewards constantly."
			className="relative z-[99999] col-span-2 h-[320px] overflow-hidden sm:h-[320px] xl:col-span-1 2xl:h-[300px]"
		>
			<div className="absolute inset-0 z-[-1]">
				{packedTokens.map(token => (
					<motion.div
						key={token.id}
						className="absolute flex items-center justify-center"
						style={{
							width: token.size,
							height: token.size,
							left: `${token.xPercent}%`,
							top: `${token.yPercent}%`,
							transform: `translate(-50%, -50%) rotate(${token.rotation}deg)`
						}}
						animate={{
							transform: [
								"translate(-50%, -500%) rotate(0deg)",
								"translate(-50%, -50%) rotate(0deg)",
								"translate(-50%, -50%) rotate(0deg)",
								"translate(-50%, -50%) rotate(0deg)",
								"translate(-50%, -50%) rotate(0deg)",
								"translate(-50%, -50%) rotate(0deg)",
								"translate(-50%, 500%) rotate(0deg)"
							]
						}}
						transition={{
							duration: 6,
							ease: "easeInOut",
							repeat: Infinity,
							repeatDelay: packedTokens.length * 0.1,
							delay: (packedTokens.length - token.id) * 0.1
						}}
					>
						<div className="flex h-full w-full items-center justify-center rounded-full border-[1px] border-dashed border-plug-green/40 bg-plug-yellow font-bold text-plug-green">
							<p className="relative text-xs">${SYMBOLS[Math.floor(Math.random() * SYMBOLS.length)]}</p>
						</div>
					</motion.div>
				))}
			</div>

			{/* Gradient overlays */}
			<div className="absolute bottom-[50%] left-0 right-0 top-0 bg-gradient-to-b from-plug-white/0 to-plug-white" />
			<div className="absolute bottom-0 left-0 right-0 top-[50%] bg-plug-white" />
		</InfoCard>
	)
}
