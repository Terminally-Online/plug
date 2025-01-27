import { FC, useRef } from "react"

import { useClient } from "wagmi"

import { cn } from "@/lib/utils"

type Props = { expanded: boolean } & React.HTMLAttributes<HTMLSpanElement>

export const AccordionContent: FC<Props> = ({ expanded, className, style, children, ...props }) => {
	const ref = useRef<HTMLDivElement>(null)

	const isClient = useClient()

	if (!isClient || !expanded) return null

	return (
		<span
			className={cn("w-full overflow-y-hidden transition-all", className)}
			style={{
				opacity: 1,
				height: ref.current?.offsetHeight,
				...style
			}}
			{...props}
		>
			<div ref={ref} className="w-full pt-4">
				{children}
			</div>
		</span>
	)
}
