import { FC } from "react"

import { DynamicFragment, StaticFragment } from "@/components"
import { ACTION_REGEX, usePlugs } from "@/contexts"

export const Fragments: FC<{
	index: number
}> = ({ index }) => {
	const { fragments } = usePlugs()

	return (
		<>
			{fragments[index].map((fragment, fragmentIndex) =>
				fragment.match(ACTION_REGEX) ? (
					<DynamicFragment
						key={`${index}-${fragmentIndex}`}
						index={index}
						fragmentIndex={fragmentIndex}
					/>
				) : (
					<StaticFragment
						key={`${index}-${fragmentIndex}`}
						index={index}
						fragmentIndex={fragmentIndex}
					/>
				)
			)}
		</>
	)
}
