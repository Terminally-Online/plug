import type { FC } from "react"

import { twMerge } from "tailwind-merge"

type Props = {
	variant?: "primary" | "secondary" | "white" | "disabled" | "destructive"
	text: string
} & React.HTMLProps<HTMLButtonElement>

const variants: Record<NonNullable<Props["variant"]>, string> = {
	primary: "bg-gradient-to-tr from-[#00EF35] to-[#93DF00] text-white",
	secondary: "bg-[#D9D9D9]/20 text-black hover:bg-[#D9D9D9]/40",
	white: "bg-white text-black hover:bg-opacity-80",
	disabled: "bg-gradient-to-tr from-[#D9D9D940] to-[#D9D9D9]",
	destructive: "bg-gradient-to-tr from-[#EF0E00] to-[#DF5000] text-white"
}

export const MainButton: FC<Props> = ({
	variant = "primary",
	text,
	className
}) => {
	const base =
		"rounded-full py-[12px] px-[40px] font-bold transition-bg transition-text duration-200 hover:text-opacity-80"

	return (
		<button className={twMerge(base, className, variants[variant])}>
			{text}
		</button>
	)
}

export default MainButton
