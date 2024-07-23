import { FC, PropsWithChildren } from "react"

import { AccordionContent } from "@/components"
import { cn } from "@/lib"

type Props = {
	loading?: boolean
	expanded?: boolean
	onExpand?: () => void
	noPaddingChildren?: React.ReactNode
	noPadding?: boolean
	accordion?: React.ReactNode
} & PropsWithChildren &
	React.HTMLAttributes<HTMLButtonElement>

export const Accordion: FC<Props> = ({
	loading = false,
	expanded = false,
	onExpand = () => {},
	noPaddingChildren,
	noPadding = false,
	children,
	className,
	accordion
}) => (
	<button
		className={cn(
			"group flex h-min w-full flex-col items-center overflow-hidden rounded-[16px] border-[1px] border-grayscale-0",
			expanded && "bg-grayscale-0 hover:bg-white",
			loading
				? "animate-loading bg-gradient-animated bg-[length:200%_200%]"
				: "transition-all duration-200 ease-in-out",
			loading === false &&
				expanded === false &&
				"bg-white hover:border-white hover:bg-grayscale-0",
			loading === false ? "cursor-pointer" : "cursor-default",
			className
		)}
		onClick={onExpand}
	>
		{noPaddingChildren}

		<div
			className={cn(
				"flex h-min w-full flex-col",
				noPadding === false && "p-4"
			)}
		>
			{children}

			{accordion && (
				<AccordionContent expanded={expanded}>
					{accordion}
				</AccordionContent>
			)}
		</div>
	</button>
)
