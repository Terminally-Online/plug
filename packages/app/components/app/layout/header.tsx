import { FC } from "react"

import { ChevronLeft } from "lucide-react"

import { Button } from "@/components/shared/buttons/button"
import { cn } from "@/lib/utils"

type Props = {
	variant?: "raw" | "frame"
	size?: "sm" | "md" | "lg"
	onBack?: () => void
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
	frame: "pb-0"
}

const sizes: Record<NonNullable<Props["size"]>, string> = {
	sm: "",
	md: "text-lg",
	lg: "text-xl"
}

export const Header: FC<Props> = ({
	variant = "raw",
	size = "md",
	onBack,
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
		<div className={cn("flex w-full select-none flex-row items-center gap-4", variants[variant], className)}>
			{
				<Button variant="secondary" onClick={onBack} className="mr-2 rounded-[10px] p-1">
					<ChevronLeft size={14} />
				</Button>
			}

			{icon && icon}

			<div className="z-[1] flex w-full flex-row items-center gap-4 truncate overflow-ellipsis whitespace-nowrap">
				{label instanceof Object ? label : <p className={cn(base, sizes[size])}>{label}</p>}

				{children}
				{nextEmpty === false && nextLabel && (nextHref || nextOnClick) && (
					<Button
						variant="secondary"
						sizing={size}
						href={nextHref}
						onClick={nextOnClick}
						className={cn(
							"outline-none h-8 w-8 p-0 flex items-center justify-center",
							size === "md" && nextPadded === true ? "rounded-sm" : "rounded-sm",
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
		</div>
	)
}
