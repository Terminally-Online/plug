import { FC, PropsWithChildren } from "react"

import { useClient } from "wagmi"

import { AccordionContent } from "@/components"
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

export const Accordion: FC<Props> = ({
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
				"group relative flex h-min w-full flex-col items-center overflow-hidden rounded-[16px] border-[1px] border-grayscale-0 outline-none",
				expanded && "bg-grayscale-0 hover:bg-white",
				loading
					? "animate-loading bg-gradient-animated bg-[length:200%_200%]"
					: "transition-all duration-200 ease-in-out",
				loading === false && expanded === false && "bg-white hover:border-white hover:bg-grayscale-0/80",
				loading === false ? "cursor-pointer" : "cursor-default",
				className
			)}
			onClick={onExpand}
			{...props}
		>
			{noPaddingChildren}

			<div className={cn("flex h-min w-full flex-col", noPadding === false && "p-4")}>
				{children}

				{accordion && <AccordionContent expanded={expanded}>{accordion}</AccordionContent>}
			</div>

			<div className="absolute inset-0 rounded-[16px] shadow-[inset_4px_0_4px_0_rgba(255,255,255,.5)]" />
			<div className="absolute inset-0 rounded-[16px] shadow-[inset_0_4px_4px_0_rgba(255,255,255,0.5)]" />
			<div className="absolute inset-0 rounded-[16px] shadow-[inset_0_-4px_4px_0_rgba(255,255,255,0.5)]" />
		</button>
	)
}
