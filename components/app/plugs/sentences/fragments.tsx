import { FC } from "react"

import { DynamicFragment, StaticFragment } from "@/components"
import { ACTION_REGEX, usePlugs } from "@/contexts"

export const Fragments: FC<{
	index: number
	item: string
	actionIndex: number
}> = ({ index, item, actionIndex }) => {
	const { fragments } = usePlugs(item)

	return (
		<>
			{fragments[actionIndex].map((fragment, fragmentIndex) =>
				fragment.match(ACTION_REGEX) ? (
					<DynamicFragment
						key={`${actionIndex}-${fragmentIndex}`}
						item={item}
						index={index}
						actionIndex={actionIndex}
						fragmentIndex={fragmentIndex}
					/>
				) : (
					<StaticFragment
						key={`${actionIndex}-${fragmentIndex}`}
						item={item}
						actionIndex={actionIndex}
						fragmentIndex={fragmentIndex}
					/>
				)
			)}
		</>
	)
}
