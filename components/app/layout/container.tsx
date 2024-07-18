import { FC, PropsWithChildren } from "react"

import { cn } from "@/lib/utils"

type Props = PropsWithChildren & React.HTMLProps<HTMLDivElement>

export const Container: FC<Props> = ({ children, className, ...props }) => (
	<div className={cn("mx-4 flex flex-col", className)} {...props}>
		{children}
	</div>
)

export default Container
