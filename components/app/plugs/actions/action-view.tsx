import { FC, useMemo } from "react"

import Image from "next/image"

import { Plus } from "lucide-react"

import { Button, Sentence } from "@/components"
import { useFrame, usePlugs, useSockets } from "@/contexts"
import { actions, categories, formatTitle, getValues } from "@/lib"

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

export const ActionView: FC<{ id: string }> = ({ id }) => {
	const { handleFrame } = useFrame({ id, key: "actions" })
	const { plug, own, actions, handle } = usePlugs(id)

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

	if (plug === undefined) return null

	return (
		<div className="mb-72 flex flex-col">
			{actions.map((_, index) => (
				<Sentence id={id} key={index} index={index} />
			))}

			{own && (
				<div className="mt-12">
					<h4 className="mb-2 font-bold opacity-40">Next Action Suggestions</h4>
					<div className="flex flex-col gap-2">
						{suggestions.map((suggestion, idx) => (
							<button
								key={idx}
								className="flex items-center gap-4 rounded-lg bg-grayscale-0 p-4 font-bold"
								onClick={() =>
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
								<Image
									src={categories[suggestion.categoryName].image}
									alt={suggestion.categoryName}
									width={24}
									height={24}
									className="rounded-md"
								/>
								<p className="flex w-full flex-wrap items-center gap-[8px] truncate">
									{formatTitle(suggestion.actionName)}
								</p>
								<Button
									variant="secondary"
									className="group p-1"
									onClick={() =>
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
									<Plus size={14} />
								</Button>
							</button>
						))}
					</div>
				</div>
			)}
		</div>
	)
}
