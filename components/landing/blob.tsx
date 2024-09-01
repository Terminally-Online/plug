import { FC, HTMLAttributes, useEffect, useState } from "react"

import { motion, MotionProps, useMotionValue, useSpring } from "framer-motion"

import { cn } from "@/lib"

export const Blob: FC<
	HTMLAttributes<HTMLDivElement> & MotionProps & { top: string; left: string; width: string; height: string }
> = ({ top, left, width, height, className, ...props }) => {
	const [mousePosition, setMousePosition] = useState({ x: 0, y: 0 })
	const blobX = useMotionValue(0)
	const blobY = useMotionValue(0)

	const springConfig = { damping: 20, stiffness: 300 }
	const springX = useSpring(blobX, springConfig)
	const springY = useSpring(blobY, springConfig)

	useEffect(() => {
		const handleMouseMove = (event: MouseEvent) => {
			setMousePosition({ x: event.clientX, y: event.clientY })
		}

		window.addEventListener("mousemove", handleMouseMove)

		return () => {
			window.removeEventListener("mousemove", handleMouseMove)
		}
	}, [])

	useEffect(() => {
		const blobRect = document.getElementById("blob")?.getBoundingClientRect()
		if (blobRect) {
			const blobCenterX = blobRect.left + blobRect.width / 2
			const blobCenterY = blobRect.top + blobRect.height / 2
			const distanceX = mousePosition.x - blobCenterX
			const distanceY = mousePosition.y - blobCenterY
			const distance = Math.sqrt(distanceX ** 2 + distanceY ** 2)

			const maxDistance = 300
			const attractionStrength = 0.3

			const attractionX = (distanceX / distance) * Math.min(distance, maxDistance) * attractionStrength
			const attractionY = (distanceY / distance) * Math.min(distance, maxDistance) * attractionStrength

			blobX.set(attractionX)
			blobY.set(attractionY)
		}
	}, [mousePosition, blobX, blobY])

	return (
		<motion.div
			id="blob"
			className={cn(
				"absolute z-[-2] rounded-full bg-gradient-to-tr from-plug-green to-plug-yellow blur-[120px] filter",
				className
			)}
			style={{
				width: width.includes("%") ? `${width}` : `${width}px`,
				height: height.includes("%") ? `${height}` : `${height}px`,
				top: top.includes("%") ? `${top}` : `${top}px`,
				left: left.includes("%") ? `${left}` : `${left}px`,
				x: springX,
				y: springY
			}}
			{...props}
		/>
	)
}
