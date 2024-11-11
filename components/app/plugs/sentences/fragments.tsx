import { FC, useMemo } from "react"

import { DynamicFragment, StaticFragment } from "@/components"
import { Action } from "@/lib"
import { ACTION_REGEX, useActions } from "@/state"

export const Fragments: FC<{
	index: number
	item: string
	action: Action
	actionIndex: number
	preview: boolean
}> = ({ index, item, action, actionIndex, preview }) => {
	const [actions] = useActions()

	const fragments = useMemo(() => {
		return actions[action.protocol].schema[action.action].sentence.split(ACTION_REGEX) as string[]
	}, [actions, action])

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
