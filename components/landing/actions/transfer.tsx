import React, { useEffect, useRef, useState } from "react"

import { motion, useScroll, useTransform } from "framer-motion"
import { Handshake } from "lucide-react"

import { InfoCard } from "@/components"

interface Circle {
	x: number
	y: number
	radius: number
}

interface Token extends Circle {
	id: number
	rotation: number
	size: number
}

const distance = (a: Circle, b: Circle): number => {
	return Math.sqrt(Math.pow(a.x - b.x, 2) + Math.pow(a.y - b.y, 2))
}

const circleCollides = (circle: Circle, circles: Circle[], minDistance: number): boolean => {
	for (let other of circles) {
		const dist = distance(circle, other)
		if (dist < minDistance) return true
	}
	return false
}

interface PackingArea {
	x: number
	y: number
	width: number
	height: number
	density: number
}

const generatePackedCircles = (
	count: number,
	width: number,
	height: number,
	baseRadius: number
): (Circle & { size: number })[] => {
	const circles: (Circle & { size: number })[] = []
	const maxAttempts = 150
	const padding = baseRadius

	const packingAreas: PackingArea[] = [
		{
			x: padding,
			y: 0,
			width: width - padding * 2,
			height: height,
			density: 1
		}
	] as const

	// Helper function to get hexagonal grid position
	const getHexPosition = (row: number, col: number, area: PackingArea): { x: number; y: number } => {
		const hexWidth = baseRadius * 1.732 // √3
		const hexHeight = baseRadius * 1.5

		return {
			x: area.x + col * hexWidth + (row % 2) * (hexWidth / 2),
			y: area.y + row * hexHeight
		}
	}

	// Place tokens in each area using hexagonal pattern as starting positions
	for (const area of packingAreas) {
		const areaCount = Math.floor(count * area.density)
		let placed = 0

		// Calculate how many rows and columns we need for hexagonal packing
		const rows = Math.ceil(Math.sqrt(areaCount))
		const cols = Math.ceil(areaCount / rows)

		// Try placing tokens in a hexagonal pattern with some randomization
		for (let row = 0; row < rows && placed < areaCount; row++) {
			for (let col = 0; col < cols && placed < areaCount; col++) {
				let attempts = 0
				let circle: Circle & { size: number }

				do {
					// Generate random size for this token
					const sizeMultiplier = 0.7 + Math.random() * 0.6 // Random between 0.7 and 1.3
					const tokenSize = baseRadius * sizeMultiplier

					// Get base hexagonal position
					const hexPos = getHexPosition(row, col, area)

					// Add slight randomization to the hexagonal position
					const jitter = tokenSize * 0.2
					circle = {
						x: hexPos.x + (Math.random() - 0.5) * jitter,
						y: hexPos.y + (Math.random() - 0.5) * jitter,
						radius: tokenSize,
						size: tokenSize * 2 // Store the diameter
					}

					// Add some horizontal spread based on vertical position
					const spreadFactor = 1 - circle.y / height
					const xSpread = area.width * (0.3 + spreadFactor * 0.2)
					const xCenter = area.x + area.width / 2
					circle.x = xCenter + (circle.x - xCenter) * (1 + spreadFactor)

					attempts++
				} while (circleCollides(circle, circles, circle.radius * 0.5) && attempts < maxAttempts)

				if (attempts < maxAttempts) {
					circles.push(circle)
					placed++
				}
			}
		}

		// Fill any gaps with random positions
		while (placed < areaCount) {
			let attempts = 0
			let circle: Circle & { size: number }

			do {
				const spreadFactor = 1 - (area.y + area.height / 2) / height
				const xSpread = area.width * (0.4 + spreadFactor * 0.3)
				const xCenter = area.x + area.width / 2

				const sizeMultiplier = 1.5 + Math.random()
				const tokenSize = baseRadius * sizeMultiplier

				circle = {
					x: xCenter + (Math.random() - 0.5) * xSpread,
					y: area.y + Math.random() * area.height,
					radius: tokenSize,
					size: tokenSize * 2
				}
				attempts++
			} while (circleCollides(circle, circles, circle.radius * 0.5) && attempts < maxAttempts)

			if (attempts < maxAttempts) {
				circles.push(circle)
				placed++
			} else {
				break // Prevent infinite loop if we can't place more tokens
			}
		}
	}

	return circles
}

const BASE_TOKEN_RADIUS = 32
const TOKENS = 40
const SYMBOLS = ["ETH", "BTC", "SOL", "AVAX", "MATIC", "UNI", "AAVE", "SUSHI", "YFI", "COMP", "MKR", "CRV"]

export const ActionTransfer: React.FC = () => {
	const [packedTokens, setPackedTokens] = useState<Token[]>([])
	const containerRef = useRef<HTMLDivElement>(null)

	const { scrollYProgress } = useScroll({
		target: containerRef,
		offset: ["start end", "end start"]
	})
	const pathLength = useTransform(scrollYProgress, [-500, 0], [1, 0])

	useEffect(() => {
		const positions = generatePackedCircles(TOKENS, 320, 320, BASE_TOKEN_RADIUS)
		const tokens = positions.map((pos, index) => ({
			...pos,
			id: index,
			rotation: -15 + Math.random() * 30
		}))
		setPackedTokens(tokens)
	}, [])

	return (
		<div ref={containerRef}>
			<InfoCard
				icon={<Handshake size={24} className="opacity-40" />}
				text="Earn."
				description="Earn yield and creator rewards constantly."
				className="relative z-[99999] h-[320px] sm:h-[320px] 2xl:h-[300px]"
			>
				<div className="absolute inset-0 z-[-1] overflow-hidden">
					{packedTokens.map(token => (
						<motion.div
							key={token.id}
							className="absolute flex items-center justify-center"
							style={{
								width: token.size,
								height: token.size,
								left: token.x - token.radius,
								top: token.y - token.radius,
								transform: `rotate(${token.rotation}deg)`
							}}
							animate={{
								transform: [
									"translateY(-500%)",
									"translateY(0%)",
									"translateY(0%)",
									"translateY(0%)",
									"translateY(0%)",
									"translateY(0%)",
									"translateY(500%)"
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
							<div className="flex h-full w-full items-center justify-center rounded-full border-[2px] border-dashed border-plug-green/40 bg-plug-yellow font-bold text-plug-green">
								<p className="relative text-xs">
									${SYMBOLS[Math.floor(Math.random() * SYMBOLS.length)]}
								</p>
							</div>
						</motion.div>
					))}
				</div>

				{/* Gradient overlays */}
				<div className="absolute bottom-[50%] left-0 right-0 top-0 bg-gradient-to-b from-plug-white/0 to-plug-white" />
				<div className="absolute bottom-0 left-0 right-0 top-[50%] bg-plug-white" />
			</InfoCard>
		</div>
	)
}
