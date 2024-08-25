import { FC } from "react"

import { DynamicFragment, StaticFragment } from "@/components"
import { ACTION_REGEX, usePlugs } from "@/contexts"

export const Fragments: FC<{
	id: string
	index: number
}> = ({ id, index }) => {
	const { fragments } = usePlugs(id)

	return (
		<>
			{fragments[index].map((fragment, fragmentIndex) =>
				fragment.match(ACTION_REGEX) ? (
					<DynamicFragment key={`${index}-${fragmentIndex}`} id={id} index={index} fragmentIndex={fragmentIndex} />
				) : (
					<StaticFragment key={`${index}-${fragmentIndex}`} id={id} index={index} fragmentIndex={fragmentIndex} />
				)
			)}
		</>
	)
}
