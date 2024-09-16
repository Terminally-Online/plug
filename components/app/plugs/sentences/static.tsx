import { FC } from "react"

import { usePlugs } from "@/contexts"

export const StaticFragment: FC<{
	item: string
	actionIndex: number
	fragmentIndex: number
}> = ({ item, actionIndex, fragmentIndex }) => {
	const { fragments } = usePlugs(item)

	return <span>{fragments[actionIndex][fragmentIndex]} </span>
}
