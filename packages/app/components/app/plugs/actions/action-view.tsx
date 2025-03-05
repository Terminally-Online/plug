import { FC, useCallback, useMemo } from "react"

import { Sentence } from "@/components/app/plugs/sentences/sentence"
import { Callout } from "@/components/app/utils/callout"
import { Image } from "@/components/app/utils/image"
import { Accordion } from "@/components/shared/utils/accordion"
import { cn, Action, formatTitle, getValues, useConnect, ActionSchemaCoils } from "@/lib"
import { useActions } from "@/state/actions"
import { columnByIndexAtom } from "@/state/columns"
import { useAtom, useSetAtom } from "jotai"
import { editPlugAtom, plugByIdAtom, plugsAtom } from "@/state/plugs"
import { api } from "@/server/client"
import { DragDropContext, Draggable, Droppable, DropResult } from "@hello-pangea/dnd"
import { Sentences } from "../sentences/sentences"

const getProtocolFrequency = (actions: Pick<Action, "protocol" | "action">[]): Record<string, number> => {
	const protocolFrequency: Record<string, number> = {}

	actions.forEach(action => (protocolFrequency[action.protocol] = (protocolFrequency[action.protocol] || 0) + 1))

	return protocolFrequency
}

export const ActionView: FC<{ index: number }> = ({ index }) => {
	const { account: { session } } = useConnect()

	const [column] = useAtom(columnByIndexAtom(index))
	const [solverActions] = useActions()

	const setPlugs = useSetAtom(plugsAtom)
	const [plug] = useAtom(plugByIdAtom(column?.item ?? ""))
	const own = plug && session && session.address === plug.socketId || false
	const editPlug = useSetAtom(editPlugAtom)
	const actionMutation = api.plugs.action.edit.useMutation({
		onSuccess: result => editPlug(result)
	})
	const edit = useCallback(
		(...params: Parameters<typeof actionMutation.mutate>) => actionMutation.mutate(...params),
		[actionMutation]
	)

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
		if (!plug || plug?.actions.length === 0) return baseSuggestions.slice(0, 5)

		const protocolFrequency = getProtocolFrequency(plug.actions)
		const selectedActions = new Set(plug.actions.map(action => `${action.protocol}-${action.action}`))
		const mostRecentAction = plug.actions[plug.actions.length - 1]

		return baseSuggestions
			.filter(suggestion => !selectedActions.has(`${suggestion.protocol}-${suggestion.action}`))
			.sort((a, b) => {
				if (a.protocol === mostRecentAction.protocol) return -1
				if (b.protocol === mostRecentAction.protocol) return 1
				return (protocolFrequency[b.protocol] || 0) - (protocolFrequency[a.protocol] || 0)
			})
			.slice(0, 3)
	}, [baseSuggestions, plug])


	const handleDragEnd = (result: DropResult) => {
		if (!result.destination || !plug) return

		const newActions = [...plug.actions]
		const [removed] = newActions.splice(result.source.index, 1)
		newActions.splice(result.destination.index, 0, removed)

		const actions = JSON.stringify(newActions)

		setPlugs(prev => prev.map(p => plug && p.id === column?.item ? { ...p, actions, updatedAt: new Date() } : p))
		edit({
			id: plug.id,
			actions
		})
	}

	if (!plug) return null

	return (
		<div className="mb-72 flex flex-col">
			<Callout.EmptyPlug index={index} isEmpty={plug.actions.length === 0} />

			<Sentences index={index} />

			{own && plug.actions.length > 0 && suggestions.length > 0 && (
				<div className="mt-12">
					<h4 className="mb-2 font-bold opacity-40">Suggestions</h4>
					<div className="flex flex-col gap-2">
						{suggestions.map((suggestion, suggestionIndex) => (
							<Accordion
								key={`${suggestionIndex}-suggestion`}
								className="flex items-center gap-4 font-bold"
								onExpand={() =>
									edit({
										id: plug.id,
										actions: JSON.stringify([
											...plug.actions,
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
