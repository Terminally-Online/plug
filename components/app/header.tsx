import { type FC, useRef } from "react"

import { AnimatePresence, motion, useTransform } from "framer-motion"
import { useScroll } from "framer-motion"
import { ChevronLeft } from "lucide-react"

import { cn } from "@/lib/utils"

import { Button } from "../buttons"

type Props = {
	variant?: "raw" | "frame"
	size?: "md" | "lg"
	back?: string
	icon?: JSX.Element
	label: string | JSX.Element
	nextPadded?: boolean
	nextHref?: string
	nextOnClick?: () => void
	nextLabel?: string | JSX.Element
	nextEmpty?: boolean
} & React.HTMLAttributes<HTMLDivElement>

const variants: Record<NonNullable<Props["variant"]>, string> = {
	raw: "pb-4 pt-8",
	frame: "pb-4"
}

const sizes: Record<NonNullable<Props["size"]>, string> = {
	md: "text-lg",
	lg: "text-xl"
}

export const Header: FC<Props> = ({
	variant = "raw",
	size = "md",
	back,
	icon,
	label,
	nextPadded = true,
	nextHref,
	nextOnClick,
	nextLabel,
	nextEmpty = false,
	className,
	children
}) => {
	const base = "font-bold truncate"

	return (
		<div
			className={cn(
				"sticky top-0 z-[2] flex w-full flex-row items-center gap-4 bg-white",
				variants[variant],
				className
			)}
		>
			{back && (
				<Button variant="secondary" href={back} className="mr-2 p-1">
					<ChevronLeft size={14} className="opacity-60" />
				</Button>
			)}

			{icon && icon}
			{label instanceof Object ? (
				label
			) : (
				<p className={cn(base, sizes[size])}>{label}</p>
			)}

			{children}
			{nextEmpty === false && nextLabel && (nextHref || nextOnClick) && (
				<Button
					variant="secondary"
					sizing={size}
					href={nextHref}
					onClick={nextOnClick}
					className={cn(
						"outline-none",
						size === "md" && nextPadded === true
							? "px-2 py-1 text-xs"
							: "p-1",
						children === undefined && "ml-auto"
					)}
				>
					{nextLabel}
				</Button>
			)}

			{nextEmpty === true && nextLabel && (nextHref || nextOnClick) && (
				<div className="ml-auto">{nextLabel}</div>
			)}
		</div>
	)
}
