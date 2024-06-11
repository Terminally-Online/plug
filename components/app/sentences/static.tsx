import { FC } from "react"

export const StaticFragment: FC<{ fragment: string }> = ({ fragment }) => (
	<span>{fragment} </span>
)
