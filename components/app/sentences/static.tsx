import { FC } from "react"

import { usePlugs } from "@/contexts"

export const StaticFragment: FC<{ index: number; fragmentIndex: number }> = ({
	index,
	fragmentIndex
}) => {
	const { fragments } = usePlugs()

	return <span>{fragments[index][fragmentIndex]} </span>
}
