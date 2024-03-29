"use client"

import { useEffect, useState } from "react"

export const useMouse = () => {
	const [mouse, setMouse] = useState({ x: 0, y: 0 })

	const isMoved = mouse.x !== 0 || mouse.y !== 0

	useEffect(() => {
		const handleMouseMove = (e: MouseEvent) => {
			setMouse({ x: e.clientX, y: e.clientY })
		}

		window.addEventListener("mousemove", handleMouseMove)

		return () => {
			window.removeEventListener("mousemove", handleMouseMove)
		}
	}, [])

	return { mouse, isMoved }
}

export default useMouse
