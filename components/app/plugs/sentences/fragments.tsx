import { FC } from "react"

import { ACTION_REGEX, usePlugs } from "@/contexts/PlugProvider"

import { DynamicFragment } from "./dynamic"
import { StaticFragment } from "./static"

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
