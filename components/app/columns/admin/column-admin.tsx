import { FC, HTMLAttributes } from "react"

export const ConsoleAdmin: FC<HTMLAttributes<HTMLDivElement> & { index: number }> = ({ index, ...props }) => {
	return <div {...props}></div>
}
