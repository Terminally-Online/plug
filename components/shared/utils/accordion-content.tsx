import { FC, useRef } from "react"

import { cn } from "@/lib/utils"

type Props = { expanded: boolean } & React.HTMLAttributes<HTMLSpanElement>

export const AccordionContent: FC<Props> = ({
	expanded,
	className,
	style,
	children,
	...props
}) => {
	const ref = useRef<HTMLDivElement>(null)

	return (
		<span
			className={cn("w-full overflow-y-hidden transition-all", className)}
			style={{
				opacity: expanded ? 1 : 0,
				height: expanded ? ref.current?.offsetHeight || 0 : 0,
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
