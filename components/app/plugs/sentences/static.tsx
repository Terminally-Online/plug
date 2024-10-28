import { FC } from "react"

export const StaticFragment: FC<{
	fragment: string
}> = ({ fragment }) => {
	return <span>{fragment} </span>
}
