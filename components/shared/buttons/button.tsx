import Link from "next/link"
import { FC, HTMLAttributes, PropsWithChildren } from "react"

import { cn } from "@/lib"

type Props = {
	variant?: "primary" | "primaryDisabled" | "secondary" | "white" | "disabled" | "destructive" | "none"
	sizing?: "sm" | "md" | "lg"
	href?: string
	external?: boolean
} & React.HTMLProps<HTMLButtonElement> &
	PropsWithChildren

const variants: Record<NonNullable<Props["variant"]>, string> = {
	primary:
		"relative bg-gradient-to-tr from-plug-green to-plug-yellow border-[1px] border-white text-white transition-all duration-200 ease-in-out before:transition-all before:duration-200 before:ease-in-out whitespace-nowrap",
	primaryDisabled:
		"relative bg-white text-plug-green border-[1px] border-plug-green transition-all duration-200 ease-in-out before:transition-all before:duration-200 before:ease-in-out whitespace-nowrap",
	secondary:
		"border-[1px] border-grayscale-100 text-black hover:border-white hover:bg-grayscale-100 items-center flex justify-center text-opacity-60 whitespace-nowrap [&.active]:bg-grayscale-0 [&.active]:text-opacity-100 [&.active]:hover:bg-grayscale-100 [&.active]:hover:border-grayscale-0",
	white: "bg-white text-black hover:bg-opacity-80",
	disabled: "bg-gradient-to-tr from-[#D9D9D940] to-[#D9D9D9]",
	destructive: "bg-gradient-to-tr from-[#EF0E00] to-[#DF5000] text-white",
	none: ""
}

const sizings: HTMLAttributes<HTMLButtonElement> & Record<NonNullable<Props["sizing"]>, string> = {
	sm: "py-[8px] px-[24px] text-xs rounded-sm before:rounded-sm",
	md: "py-[10px] px-[32px] text-sm rounded-md before:rounded-md",
	lg: "py-[12px] px-[40px] rounded-lg before:rounded-lg"
}

export const Button: FC<Props> = ({
	variant = "primary",
	sizing = "lg",
	href,
	onClick,
	className,
	children,
	external = false,
	disabled = false,
	...props
}) => {
	const base = "relative outline-none font-black transition-all duration-200 hover:text-opacity-100 select-none"

	if (onClick || disabled)
		return (
			<button
				{...props}
				type="button"
				onClick={disabled ? undefined : onClick}
				className={cn(variants[variant], sizings[sizing], base, className)}
				disabled={disabled}
			>
				{children}
			</button>
		)

	if (href === undefined || typeof href == String(undefined)) return null

	external = href ? href.startsWith("http") : false

	if (external)
		return (
			<a
				href={href}
				target="_blank"
				rel="noreferrer"
				className={cn(variants[variant], sizings[sizing], base, className)}
			>
				{children}
			</a>
		)

	return (
		<>
			<Link href={href} className={cn(variants[variant], sizings[sizing], base, className)}>
				{children}
			</Link>
		</>
	)
}

export default Button
