import { FC, PropsWithChildren } from "react"

import { AccordionContent } from "@/components"
import { cn } from "@/lib"

type Props = {
	loading?: boolean
	expanded?: boolean
	onExpand?: () => void
	accordion?: React.ReactNode
} & PropsWithChildren &
	React.HTMLAttributes<HTMLButtonElement>

export const Accordion: FC<Props> = ({
	loading = false,
	expanded = false,
	onExpand = () => {},
	children,
	className,
	accordion
}) => (
	<button
		className={cn(
			"group group flex h-min w-full cursor-pointer flex-col items-center overflow-hidden rounded-[16px] border-[1px] border-grayscale-0 p-4 ",
			expanded && "bg-grayscale-0 hover:bg-white",
			loading
				? "animate-loading bg-gradient-animated bg-[length:200%_200%]"
				: "transition-all duration-200 ease-in-out",
			loading === false &&
				expanded === false &&
				"bg-white hover:border-white hover:bg-grayscale-0",
			className
		)}
		onClick={onExpand}
	>
		{children}

		{accordion && (
			<AccordionContent expanded={expanded}>{accordion}</AccordionContent>
		)}
	</button>
)
