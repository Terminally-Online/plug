import { FC, PropsWithChildren } from "react"

import { cn } from "@/lib"

type Props = PropsWithChildren & React.HTMLAttributes<HTMLDivElement>

export const Footer: FC<Props> = ({ className, children, ...props }) => {
	return (
		<div
			className={cn(
				"fixed bottom-0 left-0 right-0 z-[1] min-w-24 bg-gradient-to-t from-white to-white/0 p-4 py-4 pt-8",
				className
			)}
			{...props}
		>
			{children}
		</div>
	)
}
