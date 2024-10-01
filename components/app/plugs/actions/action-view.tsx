import Image from "next/image"
import { FC, useMemo } from "react"

import { Plus } from "lucide-react"

import { Accordion, Button, Callout, Sentence } from "@/components"
import { usePlugs } from "@/contexts"
import { actions, categories, formatTitle, getValues } from "@/lib"
import { useColumns } from "@/state"

const baseSuggestions = Object.entries(actions).flatMap(([categoryName, actions]) =>
	Object.keys(actions).map(actionName => ({
		categoryName,
		actionName
	}))
)

const getProtocolFrequency = (
	actions: Array<{
		categoryName: string
		actionName: string
	}>
): Record<string, number> => {
	const protocolFrequency: Record<string, number> = {}

	actions.forEach(action => {
		protocolFrequency[action.categoryName] = (protocolFrequency[action.categoryName] || 0) + 1
	})

	return protocolFrequency
}

export const ActionView: FC<{ index: number }> = ({ index }) => {
	const { column } = useColumns(index)
	const { item } = column ?? {}
	const { plug, own, actions, handle } = usePlugs(item)

	const suggestions = useMemo(() => {
		const protocolFrequency = getProtocolFrequency(actions)
		const selectedActions = new Set(actions.map(action => `${action.categoryName}-${action.actionName}`))

		if (actions.length === 0) {
			return baseSuggestions.slice(0, 5)
		}

		const mostRecentAction = actions[actions.length - 1]

		return baseSuggestions
			.filter(suggestion => !selectedActions.has(`${suggestion.categoryName}-${suggestion.actionName}`))
			.sort((a, b) => {
				if (a.categoryName === mostRecentAction.categoryName) return -1
				if (b.categoryName === mostRecentAction.categoryName) return 1
				return (protocolFrequency[b.categoryName] || 0) - (protocolFrequency[a.categoryName] || 0)
			})
			.slice(0, 3)
	}, [actions])

	if (!item || !plug) return null

	return (
		<div className="mb-72 flex flex-col">
			<Callout.EmptyPlug index={index} isEmpty={actions.length === 0} />

			{actions.map((_, actionIndex) => (
				<Sentence key={index} index={index} item={item} actionIndex={actionIndex} />
			))}

			{own && (
				<div className="mt-12">
					<h4 className="mb-2 font-bold opacity-40">Next Action Suggestions</h4>
					<div className="flex flex-col gap-2">
						{suggestions.map((suggestion, idx) => (
							<Accordion
								key={idx}
								className="flex items-center gap-4 font-bold"
								onExpand={() =>
									handle.action.edit({
										id: plug.id,
										actions: JSON.stringify([
											...actions,
											{
												...suggestion,
												values: getValues(suggestion.categoryName, suggestion.actionName)
											}
										])
									})
								}
							>
								<div className="flex items-center gap-4">
									<div className="relative h-6 w-10">
										<Image
											src={categories[suggestion.categoryName].image}
											alt={suggestion.categoryName}
											width={64}
											height={64}
											className="absolute h-6 w-6 rounded-sm blur-xl filter"
										/>
										<Image
											src={categories[suggestion.categoryName].image}
											alt={suggestion.categoryName}
											width={64}
											height={64}
											className="absolute h-6 w-6 rounded-sm"
										/>
									</div>
									<p className="flex w-full flex-wrap items-center gap-[8px] truncate">
										{formatTitle(suggestion.actionName)}
									</p>
								</div>
							</Accordion>
						))}
					</div>
				</div>
			)}
		</div>
	)
}
