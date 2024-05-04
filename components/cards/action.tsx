import type { FC } from "react"

import { motion, MotionProps } from "framer-motion"
import { twMerge } from "tailwind-merge"

import { colors } from "@/lib/constants"

type Props = {
	size?: "md" | "lg"
	color?: keyof typeof colors
	glow?: boolean
	title: string
	className?: string
} & MotionProps

const sizes: Record<NonNullable<Props["size"]>, string> = {
	md: "text-lg font-bold",
	lg: "text-md lg:text-xl font-bold min-h-[140px] lg:min-h-[200px]"
}

export const ActionCard: FC<Props> = ({
	size = "md",
	color = "blue",
	glow = false,
	title,
	className,
	...props
}) => {
	const base = `rounded-lg p-4 lg:p-8 text-white text-left flex flex-col justify-end`

	return (
		<motion.button
			className={twMerge(base, sizes[size], className)}
			style={{
				backgroundColor: colors[color],
				boxShadow: glow ? `0 0 20px ${colors[color]}` : "none"
			}}
			{...props}
		>
			<span className="w-[90%]">{title}</span>
		</motion.button>
	)
}

export default ActionCard
