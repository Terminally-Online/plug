import { FC, memo, PropsWithChildren } from "react"

import { useClient } from "wagmi"

import { AccordionContent } from "@/components/shared/utils/accordion-content"
import { cn } from "@/lib"

type Props = PropsWithChildren<{
	loading?: boolean
	expanded?: boolean
	onExpand?: () => void
	noPaddingChildren?: React.ReactNode
	noPadding?: boolean
	accordion?: React.ReactNode
}> &
	React.HTMLAttributes<HTMLButtonElement>

export const Accordion: FC<Props> = memo(
	({
		loading = false,
		expanded = false,
		onExpand = () => {},
		noPaddingChildren,
		noPadding = false,
		children,
		className,
		accordion,
		...props
	}) => {
		const isClient = useClient()

		if (!isClient) return null

		return (
			<button
				className={cn(
					"group flex w-full flex-col items-center overflow-hidden rounded-[16px] border-[1px] border-plug-green/10 outline-none",
					expanded && "bg-plug-green/5 hover:bg-white",
					loading
						? "animate-loading bg-gradient-animated bg-[length:200%_200%]"
						: "transition-all duration-200 ease-in-out",
					loading === false && expanded === false && "bg-white hover:border-white hover:bg-plug-green/5",
					loading === false ? "cursor-pointer" : "cursor-default",
					className
				)}
				onClick={onExpand}
				{...props}
			>
				{noPaddingChildren}

				<div className={cn("flex h-full w-full flex-col", noPadding === false && "p-4")}>
					{children}

					{accordion && <AccordionContent expanded={expanded}>{accordion}</AccordionContent>}
				</div>
			</button>
		)
	}
)

Accordion.displayName = "Accordion"
