import { useEffect, useState } from "react"

import { motion } from "framer-motion"

import useMouse from "@/lib/hooks/useMouse"

export default function Grid() {
	const { mouse, isMoved } = useMouse()

	const [dimensions, setDimensions] = useState({ width: 0, height: 0 })
	const [position, setPosition] = useState({ x: 0, y: 0 })

	const cells = 10
	const cell = dimensions.width / cells
	const cols = Math.floor(dimensions.width / cell)
	const ratio = dimensions.width / dimensions.height
	const rows = Math.ceil(cols / ratio)
	const radius = Math.ceil((cell * 2) / 3)

	const power =
		Math.pow(mouse.x - position.x, 2) + Math.pow(mouse.y - position.y, 2)
	const base = 1.5
	const opacity = base - Math.sqrt(power) / 1000

	useEffect(() => {
		if (mouse.x === position.x && mouse.y === position.y) return

		const speed = 0.05

		requestAnimationFrame(() => {
			setPosition(prevPos => ({
				x: prevPos.x + (mouse.x - prevPos.x) * speed,
				y: prevPos.y + (mouse.y - prevPos.y) * speed
			}))
		})
	}, [mouse, position])

	useEffect(() => {
		setDimensions({ width: window.innerWidth, height: window.innerHeight })
	}, [])

	return (
		<motion.div
			className={`fixed left-0 top-0 z-[-1] grid h-screen w-screen gap-[1px] bg-stone-950`}
			style={{
				gridTemplateColumns: `repeat(${cols}, 1fr)`,
				gridTemplateRows: `repeat(${rows}, 1fr)`
			}}
			transition={{ duration: 0.4, delay: 0.4, ease: "easeInOut" }}
		>
			{isMoved && (
				<motion.div
					className={`bg-gradient-radial absolute z-[-2] rounded-full bg-white from-white to-transparent blur-xl`}
					style={{
						width: radius * 2,
						height: radius * 2,
						left: position.x,
						top: position.y,
						opacity,
						transform: `translate(-50%, -50%)`
					}}
					animate={{ opacity }}
				/>
			)}

			<div className="absolute left-0 top-0 z-[-1] h-screen w-screen bg-gradient-to-tr from-stone-900 to-transparent" />

			{Array.from({ length: 1 + cols * rows }).map((_, i) => (
				<div key={i} className="bg-stone-900" />
			))}
		</motion.div>
	)
}
