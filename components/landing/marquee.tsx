import { useEffect, useRef, useState } from "react"

import { motion, PanInfo, useSpring, useTransform } from "framer-motion"
import normalizeWheel from "normalize-wheel"
import { useRafLoop } from "react-use"

const text =
	"PLUG PUBLIC ALPHA COMING SOON // EXPERIENCE THE BLOCKCHAIN LIKE NEVER BEFORE WITH PLUG // "
const count = 10

export function MarqueeItem({
	speed,
	children
}: {
	speed: any
	children: React.ReactNode
}) {
	const itemRef = useRef<HTMLDivElement>(null)
	const rectRef = useRef<DOMRect | null>(null)
	const x = useRef(0)

	const width = typeof window !== "undefined" ? window.innerWidth : 0
	const height = typeof window !== "undefined" ? window.innerHeight : 0

	const setX = () => {
		if (!itemRef.current || !rectRef.current) {
			return
		}

		const xPercentage = (x.current / rectRef.current.width) * 100

		if (xPercentage < -100) {
			x.current = 0
		}

		if (xPercentage > 0) {
			x.current = -rectRef.current.width
		}

		itemRef.current.style.transform = `translate3d(${xPercentage}%, 0, 0)`
	}

	const loop = () => {
		x.current -= speed.get()

		setX()
	}

	useEffect(() => {
		if (!itemRef.current) return

		rectRef.current = itemRef.current.getBoundingClientRect()
	}, [itemRef, rectRef, width, height])

	return (
		<motion.div
			className="user-select-none whitespace-nowrap pr-[0.25rem] text-lg leading-[-0.05rem] text-black/60 dark:text-white/60"
			ref={itemRef}
		>
			{children}
		</motion.div>
	)
}

export function MarqueeClient({
	speed = 1,
	threshold = 0.014,
	wheelFactor = 1.8,
	dragFactor = 1.2
}) {
	const ref = useRef<HTMLDivElement>(null)
	const slowDownRef = useRef(false)
	const isScrolling = useRef<NodeJS.Timeout | null>(null)

	const x = useRef(0)

	const windowWidth = typeof window !== "undefined" ? window.innerWidth : 0

	const speedSpring = useSpring(speed, {
		damping: 40,
		stiffness: 90,
		mass: 5
	})
	const skewX = useTransform(
		speedSpring,
		[-windowWidth * 0.05, 0, windowWidth * 0.05],
		[1, 0, 1]
	)

	const handleOnWheel = (e: React.WheelEvent<HTMLDivElement> | undefined) => {
		const normalized = normalizeWheel(e)

		x.current = normalized.pixelY * wheelFactor

		if (isScrolling.current) {
			window.clearTimeout(isScrolling.current)
		}

		isScrolling.current = setTimeout(() => {
			speedSpring.set(speed)
		}, 30)
	}

	const handleDragStart = () => {
		slowDownRef.current = true
		ref.current?.classList.add("drag")
		speedSpring.set(0)
	}

	const handleOnDrag = (_: unknown, info: PanInfo) => {
		speedSpring.set(dragFactor * -info.delta.x)
	}

	const handleDragEnd = () => {
		slowDownRef.current = false
		ref.current?.classList.remove("drag")
		speedSpring.set(speed)
	}

	const loop = () => {
		if (slowDownRef.current || Math.abs(x.current) < threshold) return

		x.current *= 0.66

		if (x.current < 0) x.current = Math.min(x.current, 0)
		else x.current = Math.max(x.current, 0)

		speedSpring.set(speed + x.current)
	}

	useEffect(() => {
		let index = 0

		const interval = setInterval(() => {
			const title = `${text.slice(index)} ${text.slice(0, index)}`

			document.title = title

			index = index + 1 > text.length ? 0 : index + 1
		}, 100)

		return () => clearInterval(interval)
	}, [])

	// useRafLoop(loop)

	return (
		<div className="mb-auto border-b-[1px] border-stone-950 py-4 backdrop-blur-lg">
			<motion.div
				className="flex cursor-grab items-center"
				ref={ref}
				style={{ skewX }}
				onWheel={handleOnWheel}
				drag="x"
				dragPropagation={true}
				dragConstraints={{ left: 0, right: 0 }}
				onDragStart={handleDragStart}
				onDrag={handleOnDrag}
				onDragEnd={handleDragEnd}
				dragElastic={0.000001}
			>
				{Array.from(Array(count).keys()).map((_, i) => (
					<MarqueeItem key={i} speed={speedSpring}>
						{text}
					</MarqueeItem>
				))}
			</motion.div>
		</div>
	)
}

export const Marquee = () => {
	const [visible, setVisible] = useState(false)

	useEffect(() => {
		setVisible(true)
	}, [])

	return <div>{visible && <MarqueeClient />}</div>
}

export default Marquee
