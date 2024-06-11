import type { FC } from "react"

import { ChevronLeft, Plug } from "lucide-react"

import { cn } from "@/lib/utils"

import { Button } from "../buttons"

type Props = {
	variant?: "raw" | "frame"
	size?: "md" | "lg"
	back?: string
	icon?: JSX.Element
	label: string
	nextPadded?: boolean
	nextHref?: string
	nextOnClick?: () => void
	nextLabel?: string | JSX.Element
}

const variants: Record<NonNullable<Props["variant"]>, string> = {
	raw: "pb-[20px] pt-[30px]",
	frame: "pb-[20px]"
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
	nextLabel
}) => {
	const base = "font-bold"

	return (
		<div
			className={cn(
				"flex w-full flex-row items-center gap-2",
				variants[variant]
			)}
		>
			{back && (
				<Button variant="secondary" href={back} className="mr-2 p-1">
					<ChevronLeft size={14} className="opacity-60" />
				</Button>
			)}

			{icon && icon}
			<h1 className={cn(base, sizes[size])}>{label}</h1>

			{nextLabel && (nextHref || nextOnClick) && (
				<Button
					variant="secondary"
					sizing={size}
					href={nextHref}
					onClick={nextOnClick}
					className={cn(
						"ml-auto outline-none",
						size === "md" && nextPadded === true
							? "px-2 py-1 text-xs"
							: "p-1"
					)}
				>
					{nextLabel}
				</Button>
			)}
		</div>
	)
}
