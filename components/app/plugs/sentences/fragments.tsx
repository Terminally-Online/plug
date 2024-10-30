import { FC, useMemo } from "react"

import { DynamicFragment, StaticFragment } from "@/components"
import { ACTION_REGEX } from "@/contexts"
import { Action } from "@/lib"
import { actions as staticActions } from "@/lib/constants"

export const Fragments: FC<{
	index: number
	item: string
	actionIndex: number
	action: Action
	preview: boolean
}> = ({ index, item, actionIndex, action, preview }) => {
	// Split all of the sentence fragments into an appropriate array based on the
	// regex shape that enables the f-string like syntax.
	const fragments = useMemo(() => {
		const staticAction = staticActions[action.categoryName][action.actionName]

		return staticAction ? (staticAction["sentence"].split(ACTION_REGEX) as string[]) : []
	}, [action])

	const dynamic = useMemo(() => {
		return fragments.filter(fragment => fragment.match(ACTION_REGEX))
	}, [fragments])

	let dynamicIndex = -1

	return (
		<>
			{fragments.map((fragment, fragmentIndex) => {
				if (fragment.match(ACTION_REGEX)) {
					dynamicIndex++
					return (
						<DynamicFragment
							key={`${actionIndex}-${fragmentIndex}`}
							item={item}
							index={index}
							actionIndex={actionIndex}
							fragmentIndex={fragmentIndex}
							dynamicIndex={dynamicIndex}
							fragment={fragment}
							action={action}
							dynamic={dynamic}
							preview={preview}
						/>
					)
				}
				return <StaticFragment key={`${actionIndex}-${fragmentIndex}`} fragment={fragment} />
			})}
		</>
	)
}
