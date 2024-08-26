import { FC, HTMLAttributes } from "react"

export const ConsoleAdmin: FC<HTMLAttributes<HTMLDivElement> & { id: string }> = ({ id, ...props }) => {
	return <div {...props}></div>
}
