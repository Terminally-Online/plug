import { FC, PropsWithChildren } from "react"

import Link from "next/link"

import { cn } from "@/lib"

type Props = {
	variant?: "primary" | "secondary" | "white" | "disabled" | "destructive"
	sizing?: "sm" | "md" | "lg"
	href?: string
	external?: boolean
} & React.HTMLProps<HTMLButtonElement> &
	PropsWithChildren

const variants: Record<NonNullable<Props["variant"]>, string> = {
	primary:
		"relative bg-gradient-to-tr from-plug-green to-plug-yellow text-white before:absolute before:inset-0 before:bg-gradient-to-tr before:from-plug-green before:to-plug-yellow before:rounded-lg before:w-full before:h-full before:blur-sm before:z-[-1] hover:before:blur-md text-opacity-90 transition-all duration-200 ease-in-out before:transition-all before:duration-200 before:ease-in-out whitespace-nowrap",
	secondary:
		"border-[1px] border-grayscale-0 text-black hover:border-white hover:bg-grayscale-0 items-center flex justify-center text-opacity-60 whitespace-nowrap bg-white [&.active]:bg-grayscale-0 [&.active]:text-opacity-100 [&.active]:hover:bg-grayscale-100 [&.active]:hover:border-grayscale-0",
	white: "bg-white text-black hover:bg-opacity-80",
	disabled: "bg-gradient-to-tr from-[#D9D9D940] to-[#D9D9D9]",
	destructive: "bg-gradient-to-tr from-[#EF0E00] to-[#DF5000] text-white"
}

const sizings: Record<NonNullable<Props["sizing"]>, string> = {
	sm: "py-[8px] px-[24px] text-xs",
	md: "py-[10px] px-[32px] text-sm",
	lg: "py-[12px] px-[40px]"
}

export const Button: FC<Props> = ({
	variant = "primary",
	sizing = "lg",
	href,
	onClick,
	className,
	children,
	external = false,
	disabled = false
}) => {
	const base =
		"cursor-pointer outline-none rounded-lg font-bold transition-all duration-200 hover:text-opacity-100 select-none"

	if (onClick)
		return (
			<button
				onClick={disabled ? undefined : onClick}
				className={cn(
					variants[variant],
					sizings[sizing],
					base,
					className
				)}
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
				className={cn(
					variants[variant],
					sizings[sizing],
					base,
					className
				)}
			>
				{children}
			</a>
		)

	return (
		<Link
			href={href}
			className={cn(variants[variant], sizings[sizing], base, className)}
		>
			{children}
		</Link>
	)
}

export default Button
