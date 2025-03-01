import { FC, useMemo } from "react"

import { Sentence } from "@/components/app/plugs/sentences/sentence"
import { Callout } from "@/components/app/utils/callout"
import { Image } from "@/components/app/utils/image"
import { Accordion } from "@/components/shared/utils/accordion"
import { Action, formatTitle, getValues } from "@/lib"
import { useActions } from "@/state/actions"
import { columnByIndexAtom } from "@/state/columns"
import { usePlugStore } from "@/state/plugs"
import { useAtom } from "jotai"

const getProtocolFrequency = (actions: Pick<Action, "protocol" | "action">[]): Record<string, number> => {
	const protocolFrequency: Record<string, number> = {}

	actions.forEach(action => (protocolFrequency[action.protocol] = (protocolFrequency[action.protocol] || 0) + 1))

	return protocolFrequency
}

export const ActionView: FC<{ index: number }> = ({ index }) => {
	const [column] = useAtom(columnByIndexAtom(index))
	const [solverActions] = useActions()

	const { item } = column ?? {}
	const { plug, own, actions, handle } = usePlugStore(item)

	const baseSuggestions = useMemo(
		() =>
			Object.entries(solverActions).flatMap(([protocol, actions]) => {
				// const chains = actions.metadata.chains
				// if (!chains.some(chain => connectedChains.map(c => c.id as number).includes(chain.chainIds))) return []				
				
				return Object.keys(actions.schema).map(action => ({
					protocol,
					action
				}))
			}),
		[solverActions]
	)

	const suggestions = useMemo(() => {
		if (actions.length === 0) return baseSuggestions.slice(0, 5)

		const protocolFrequency = getProtocolFrequency(actions)
		const selectedActions = new Set(actions.map(action => `${action.protocol}-${action.action}`))
		const mostRecentAction = actions[actions.length - 1]

		return baseSuggestions
			.filter(suggestion => !selectedActions.has(`${suggestion.protocol}-${suggestion.action}`))
			.sort((a, b) => {
				if (a.protocol === mostRecentAction.protocol) return -1
				if (b.protocol === mostRecentAction.protocol) return 1
				return (protocolFrequency[b.protocol] || 0) - (protocolFrequency[a.protocol] || 0)
			})
			.slice(0, 3)
	}, [baseSuggestions, actions])

	if (!item || !plug) return null

	return (
		<div className="mb-72 flex flex-col">
			<Callout.EmptyPlug index={index} isEmpty={actions.length === 0} />

			{actions.map((action, actionIndex) => (
				<Sentence
					key={`${index}-${actionIndex}-${action.id}-sentence`}
					index={index}
					item={item}
					actionIndex={actionIndex}
					action={action}
				/>
			))}

			{own && actions.length > 0 && suggestions.length > 0 && (
				<div className="mt-12">
					<h4 className="mb-2 font-bold opacity-40">Suggestions</h4>
					<div className="flex flex-col gap-2">
						{suggestions.map((suggestion, suggestionIndex) => (
							<Accordion
								key={`${suggestionIndex}-suggestion`}
								className="flex items-center gap-4 font-bold"
								onExpand={() =>
									handle.action.edit({
										id: plug.id,
										actions: JSON.stringify([
											...actions,
											{
												...suggestion,
												id: Math.floor(Math.random() * 100_000_000_000),
												...getValues(
													solverActions[suggestion.protocol].schema[suggestion.action]
														.sentence
												)
											}
										])
									})
								}
							>
								<div className="flex items-center gap-4">
									<div className="relative h-6 w-10">
										<Image
											src={solverActions[suggestion.protocol].metadata.icon}
											alt={suggestion.protocol}
											width={64}
											height={64}
											className="absolute h-6 w-6 rounded-sm blur-xl filter"
										/>
										<Image
											src={solverActions[suggestion.protocol].metadata.icon}
											alt={suggestion.protocol}
											width={64}
											height={64}
											className="absolute h-6 w-6 rounded-sm"
										/>
									</div>
									<p className="flex w-full flex-wrap items-center gap-[8px] truncate">
										{formatTitle(suggestion.action)}
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
