import { FC, HTMLAttributes, useState } from "react"

import { SocketTokenList } from "../tokens"

export const SocketTokens: FC<HTMLAttributes<HTMLDivElement>> = ({
	...props
}) => {
	const [expanded, setExpanded] = useState(false)

	return <SocketTokenList expanded={expanded} />
}
