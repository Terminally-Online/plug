import { FC } from "react"

import { usePlugs } from "@/contexts"

export const StaticFragment: FC<{
	id: string
	index: number
	fragmentIndex: number
}> = ({ id, index, fragmentIndex }) => {
	const { fragments } = usePlugs(id)

	return <span>{fragments[index][fragmentIndex]} </span>
}
