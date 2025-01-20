import { FC, memo, PropsWithChildren, useCallback, useEffect, useMemo, useRef, useState } from "react"

import { AnimatePresence, motion } from "framer-motion"

import { cn } from "@/lib"

const Sparkle = ({ style, delay, color }: { style: any; delay: number; color: string }) => {
	const waveX = useMemo(() => Array.from({ length: 8 }, () => (Math.random() - 0.5) * 30), [])
	const waveY = useMemo(() => Array.from({ length: 8 }, () => (Math.random() - 0.5) * 30), [])

	return (
		<motion.div
			initial={{ scale: 0, rotate: 0, opacity: 0, x: 0, y: 0 }}
			animate={{
				scale: [0, 1, 0.8, 1, 0],
				rotate: [0, 180, 270, 360],
				opacity: [0, 1, 0.8, 1, 0],
				x: waveX,
				y: waveY
			}}
			transition={{
				duration: 3,
				ease: "easeInOut",
				repeat: Infinity,
				repeatType: "loop",
				delay: delay,
				times: [0, 0.3, 0.5, 0.8, 1]
			}}
			style={{
				position: "absolute",
				width: "10px",
				height: "10px",
				borderRadius: "50%",
				background: color,
				boxShadow: `0 0 12px 4px ${color}`,
				...style
			}}
		/>
	)
}

export const SparklingText: FC<
	PropsWithChildren<
		React.HTMLAttributes<HTMLDivElement> & {
			item: string
			color?: string
			sparkles?: boolean
			sparkleKey?: number
		}
	>
> = memo(({ children, item, color, sparkles = true, sparkleKey, ...props }) => {
	const containerRef = useRef<HTMLDivElement>(null)
	const previousSparkleKey = useRef<number | undefined>()
	const previousItem = useRef<string | undefined>()
	const hasAnimated = useRef(false)

	const [sparkleCount, setSparkleCount] = useState(0)
	const [bounds, setBounds] = useState({ top: 0, left: 0, width: 0, height: 0 })

	useEffect(() => {
		if (!containerRef.current) return

		const updateBounds = () => {
			const rect = containerRef.current?.getBoundingClientRect()
			if (rect) {
				setBounds({
					top: 0,
					left: 0,
					width: rect.width,
					height: rect.height
				})
				setSparkleCount(Math.max(4, Math.floor(rect.width * 0.7)))
			}
		}

		updateBounds()
		const observer = new ResizeObserver(updateBounds)
		observer.observe(containerRef.current)

		return () => {
			observer.disconnect()
		}
	}, [children])

	const generateSparklePositions = useCallback(() => {
		const positions = []
		for (let i = 0; i < sparkleCount; i++) {
			const xProgress = Math.random()
			const yOffset = (Math.random() - 0.5) * (bounds.height * 0.8)

			positions.push({
				left: xProgress * bounds.width,
				top: bounds.height * 0.5 + yOffset,
				delay: Math.random() * 3
			})
		}
		return positions
	}, [sparkleCount, bounds])

	useEffect(() => {
		if (item !== previousItem.current) {
			hasAnimated.current = false
			previousItem.current = item
			previousSparkleKey.current = sparkleKey
			return
		}

		if (sparkleKey !== previousSparkleKey.current) {
			if (previousSparkleKey.current !== undefined) {
				hasAnimated.current = true
			}
			previousSparkleKey.current = sparkleKey
		}
	}, [sparkleKey, item])

	return (
		<p
			{...props}
			ref={containerRef}
			className={cn(props.className, "relative w-full truncate overflow-ellipsis whitespace-nowrap")}
		>
			{children}

			{sparkles && previousSparkleKey.current !== undefined && hasAnimated.current && (
				<AnimatePresence>
					<motion.div
						key={sparkleKey}
						style={{
							position: "absolute",
							top: 0,
							left: 0,
							right: 0,
							bottom: 0,
							pointerEvents: "none",
							zIndex: 1000
						}}
						initial={{ opacity: 0 }}
						animate={{ opacity: 1 }}
						exit={{ opacity: 0 }}
						transition={{ duration: 0.1 }}
						onAnimationStart={() => {
							console.log("Animation starting:", {
								sparkleKey,
								previousKey: previousSparkleKey.current
							})
						}}
					>
						<motion.div
							initial={{ opacity: 1 }}
							animate={{ opacity: 0 }}
							transition={{ duration: 0.2, delay: 2.8 }}
						>
							{generateSparklePositions().map((pos, i) => (
								<Sparkle
									key={i}
									delay={pos.delay}
									color={color ?? "#FFF"}
									style={{
										position: "absolute",
										left: pos.left,
										top: pos.top
									}}
								/>
							))}
						</motion.div>
					</motion.div>
				</AnimatePresence>
			)}
		</p>
	)
})

SparklingText.displayName = "SparklingText"
