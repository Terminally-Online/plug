import { FC, HTMLAttributes, useMemo } from "react"

import { motion, MotionProps } from "framer-motion"

import { useBalances } from "@/contexts"
import { cn } from "@/lib"

import { SocketPositionItem } from "./position-item"

export const SocketPositionList: FC<
	HTMLAttributes<HTMLDivElement> &
		MotionProps & { id: string; expanded?: boolean }
> = ({ id, expanded, className, ...props }) => {
	const { positions } = useBalances()

	const visibilePositions = useMemo(() => {
		if (positions === undefined) return Array(3).fill(undefined)

		if (expanded) return Object.keys(positions.defi)

		return Object.keys(positions.defi).slice(0, 3)
	}, [expanded, positions])

	if (positions === undefined) return null

	return (
		<motion.div
			className={cn("flex flex-col gap-2", className)}
			initial="hidden"
			animate="visible"
			variants={{
				hidden: { opacity: 0 },
				visible: {
					opacity: 1,
					transition: {
						staggerChildren: 0.05
					}
				}
			}}
			{...(props as MotionProps)}
		>
			{visibilePositions.map((protocol: string) => (
				<SocketPositionItem
					key={protocol}
					id={id}
					position={positions.defi[protocol]}
				/>
			))}
		</motion.div>
	)
}
