import { FC, useMemo } from "react"

import { DynamicFragment, StaticFragment } from "@/components"
import { Action } from "@/lib"
import { api } from "@/server/client"
import { ACTION_REGEX } from "@/state"

export const Fragments: FC<{
	index: number
	item: string
	action: Action
	actionIndex: number
	preview: boolean
}> = ({ index, item, action, actionIndex, preview }) => {
	const { data: actions } = api.solver.actions.get.useQuery({
		protocol: action.protocol,
		action: action.action
	})

	const protocol = actions?.[action.protocol]

	const fragments = useMemo(() => {
		if (!protocol) return []

		return protocol.schema[action.action].sentence.split(ACTION_REGEX) as string[]
	}, [protocol, action])

	const dynamic = useMemo(() => {
		return fragments.filter(fragment => fragment.match(ACTION_REGEX))
	}, [fragments])

	let dynamicIndex = -1

	if (!protocol) return null

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
							protocol={protocol}
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
