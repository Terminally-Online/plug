import React, { useCallback, useEffect, useRef } from "react"

import ReactCanvasConfetti from "react-canvas-confetti"

const duration = 15 * 1000
const animationEnd = Date.now() + duration

export default function Fireworks({ enabled = false }: { enabled: boolean }) {
	const ref = useRef<any>(null)

	const getInstance = useCallback((instance: any) => {
		ref.current = instance
	}, [])

	useEffect(() => {
		if (enabled !== true || !ref.current) return

		const randomInRange = (min: number, max: number): number => {
			return Math.random() * (max - min) + min
		}

		const getAnimationSettings = (
			originXA: number,
			originXB: number,
			particleCount: number
		) => {
			return {
				startVelocity: 30,
				spread: 360,
				ticks: 60,
				zIndex: 0,
				particleCount: particleCount,
				origin: {
					x: randomInRange(originXA, originXB),
					y: Math.random() - 0.2
				}
			}
		}

		const interval: NodeJS.Timeout = setInterval(() => {
			const timeLeft = animationEnd - Date.now()

			if (timeLeft <= 0) {
				return clearInterval(interval)
			}

			const particleCount = 50 * (timeLeft / duration)

			ref.current(getAnimationSettings(0.1, 0.3, particleCount))
			ref.current(getAnimationSettings(0.7, 0.9, particleCount))
		}, 400)

		return () => {
			clearInterval(interval)
		}
	}, [ref, enabled])

	return (
		<ReactCanvasConfetti
			refConfetti={getInstance}
			className="pointer-events-none fixed left-0 top-0 z-[50] h-screen w-screen"
		/>
	)
}
