import { FC } from "react"

import { ChevronLeft } from "lucide-react"

import { Button } from "@/components"
import { cn } from "@/lib/utils"

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
	raw: "pb-4 pt-4",
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
				"flex w-full select-none flex-row items-center gap-4 bg-white",
				variants[variant],
				className
			)}
		>
			{back && (
				<Button
					variant="secondary"
					href={back}
					className="mr-2 rounded-[10px] p-1"
				>
					<ChevronLeft size={14} />
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
							? "rounded-sm px-2 py-1 text-xs"
							: "rounded-sm p-1",
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
