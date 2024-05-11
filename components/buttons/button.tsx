import type { FC, PropsWithChildren } from "react"

import Link from "next/link"

import { twMerge } from "tailwind-merge"

type Props = {
	variant?: "primary" | "secondary" | "white" | "disabled" | "destructive"
	href?: string
	external?: boolean
} & React.HTMLProps<HTMLButtonElement> &
	PropsWithChildren

const variants: Record<NonNullable<Props["variant"]>, string> = {
	primary: "bg-gradient-to-tr from-[#00EF35] to-[#93DF00] text-white",
	secondary: "bg-[#D9D9D9]/20 text-black hover:bg-[#D9D9D9]/40",
	white: "bg-white text-black hover:bg-opacity-80",
	disabled: "bg-gradient-to-tr from-[#D9D9D940] to-[#D9D9D9]",
	destructive: "bg-gradient-to-tr from-[#EF0E00] to-[#DF5000] text-white"
}

const base =
	"rounded-full py-[12px] px-[40px] font-bold transition-bg transition-text duration-200 hover:text-opacity-80"

export const Button: FC<Props> = ({
	variant = "primary",
	href,
	onClick,
	className,
	children,
	external = href && href.startsWith("http")
}) => {
	if (external)
		return (
			<a
				href={href}
				target="_blank"
				rel="noreferrer"
				className={twMerge(base, className, variants[variant])}
			>
				{children}
			</a>
		)

	if (href)
		return (
			<Link
				href={href}
				className={twMerge(base, className, variants[variant])}
			>
				{children}
			</Link>
		)

	return (
		<button
			onClick={onClick}
			className={twMerge(base, className, variants[variant])}
		>
			{children}
		</button>
	)
}

export default Button
