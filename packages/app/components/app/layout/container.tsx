import { FC, PropsWithChildren } from "react"

import { cn } from "@/lib/utils"

export const Container: FC<PropsWithChildren & React.HTMLProps<HTMLDivElement> & { column?: boolean }> = ({
	column = false,
	children,
	className,
	...props
}) => (
	<div className={cn("flex flex-col", column === false && "mx-4", className)} {...props}>
		{children}
	</div>
)

export default Container
