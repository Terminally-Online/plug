import { FC, useCallback, useMemo } from "react"

import { useAtom, useSetAtom } from "jotai"

import { DragDropContext, Draggable, Droppable, DropResult } from "@hello-pangea/dnd"

import { Sentence } from "@/components/app/plugs/sentences/sentence"
import { SchemasResponseCoils } from "@/lib"
import { api } from "@/server/client"
import { useActions } from "@/state/actions"
import { columnByIndexAtom } from "@/state/columns"
import { editPlugAtom, plugByIdAtom, plugsAtom } from "@/state/plugs"
import { useAccount } from "@/lib/hooks/account/useAccount"

type SentenceProps = { index: number }

export const Sentences: FC<SentenceProps> = ({ index }) => {
	const { address } = useAccount()

	const [column] = useAtom(columnByIndexAtom(index))
	const [solverActions] = useActions()

	const setPlugs = useSetAtom(plugsAtom)
	const [plug] = useAtom(plugByIdAtom(column?.item ?? ""))
	const own = (plug && address === plug.socketId) || false
	const editPlug = useSetAtom(editPlugAtom)
	const actionMutation = api.plugs.action.edit.useMutation({
		onSuccess: result => editPlug(result)
	})
	const edit = useCallback(
		(...params: Parameters<typeof actionMutation.mutate>) => actionMutation.mutate(...params),
		[actionMutation]
	)

	const availableCoils = useMemo(() => {
		if (!plug || !solverActions) return {}

		const coils: Record<string, { type: string; actionIndex: number }> = {}

		plug.actions.forEach((action, actionIndex) => {
			const actionSchema = solverActions[action.protocol]?.schema[action.action]

			if (!actionSchema || !actionSchema.coils) return

			Object.keys(actionSchema.coils).forEach(name => {
				if (!actionSchema?.coils?.[name]) return

				coils[name] = {
					type: actionSchema.coils[name],
					actionIndex
				}
			})

			// actionSchema.coils.forEach(coil => {
			// 	if (coil.slice?.name) {
			// 		coils[coil.slice.name] = {
			// 			type: coil.slice.type,
			// 			actionIndex,
			// 			coil
			// 		}
			// 	}
			// })
		})

		return coils
	}, [plug, solverActions])

	const handleValueChange = (actionIndex: number, inputIndex: string, value: string, additionalData: any = {}) => {
		if (!plug) return

		const isLinked = value.startsWith("<-{") && value.endsWith("}")
		const newActions = [...plug.actions]
		const action = newActions[actionIndex]
		const isNumber = additionalData.isNumber || false

		// @ts-ignore
		if (!action.values) action.values = {}
		// @ts-ignore
		action.values[parseInt(inputIndex)] = {
			index: inputIndex,
			key: inputIndex,
			name: inputIndex,
			value: isNumber && !isLinked ? parseFloat(value) : value,
			...additionalData
		}

		setPlugs(prev =>
			prev.map(p =>
				plug && p.id === column?.item ? { ...p, actions: JSON.stringify(newActions), updatedAt: new Date() } : p
			)
		)
		edit({
			id: plug.id,
			actions: JSON.stringify(newActions)
		})
	}

	const validateType = (coilName: string, expectedType: string): boolean => {
		const coilInfo = availableCoils[coilName]
		if (!coilInfo) return false

		return coilInfo.type === expectedType
	}

	const handleDragEnd = (result: DropResult) => {
		if (!result.destination || !plug || !own) return

		if (result.source.index === result.destination.index) return

		const newActions = [...plug.actions]
		const [removed] = newActions.splice(result.source.index, 1)
		newActions.splice(result.destination.index, 0, removed)
		const actions = JSON.stringify(newActions)

		setPlugs(prev => prev.map(p => (plug && p.id === column?.item ? { ...p, actions, updatedAt: new Date() } : p)))
		edit({
			id: plug.id,
			actions
		})
	}

	const handleRemoveAction = (actionIndex: number) => {
		if (!plug) return

		const newActions = plug.actions.filter((_, i) => i !== actionIndex)
		const actions = JSON.stringify(newActions)

		setPlugs(prev => prev.map(p => (plug && p.id === column?.item ? { ...p, actions, updatedAt: new Date() } : p)))
		edit({
			id: plug.id,
			actions
		})
	}

	if (!plug) return null

	return (
		<div className="flex flex-col">
			<DragDropContext onDragEnd={handleDragEnd}>
				<Droppable droppableId={`${index}-${plug.id}-items`} direction="vertical">
					{provided => (
						<div ref={provided.innerRef} className="flex flex-col" {...provided.droppableProps}>
							{plug.actions.map((action, actionIndex) => {
								const values = Object.values(plug.actions[actionIndex + 1]?.values ?? {})
								const linked = values.filter(
									val => typeof val?.value === "string" && val?.value?.startsWith("<-{")
								)

								let prevCoils: SchemasResponseCoils | undefined = {}
								if (actionIndex > 0) {
									const prevAction = plug.actions[actionIndex - 1]
									prevCoils =
										solverActions[prevAction.protocol]?.schema[prevAction.action]?.coils || {}
								}

								return (
									<Draggable
										key={String(action.id)}
										draggableId={String(action.id)}
										index={actionIndex}
										isDragDisabled={column?.frame !== undefined}
									>
										{(provided, snapshot) => (
											<div
												ref={provided.innerRef}
												className="flex h-full w-full flex-row rounded-lg"
												{...provided.draggableProps}
												style={{
													...provided.draggableProps.style
												}}
											>
												<Sentence
													index={index}
													item={column?.item ?? ""}
													actionIndex={actionIndex}
													action={action}
													linked={linked}
													prevCoils={prevCoils}
													dragging={snapshot.isDragging}
													handleValueChange={(inputIndex, value, additionalData) =>
														handleValueChange(
															actionIndex,
															inputIndex,
															value,
															additionalData
														)
													}
													handleRemoveAction={() => handleRemoveAction(actionIndex)}
													validateType={validateType}
													availableCoils={availableCoils}
													{...provided.dragHandleProps}
												/>
											</div>
										)}
									</Draggable>
								)
							})}

							{provided.placeholder}
						</div>
					)}
				</Droppable>
			</DragDropContext>
		</div>
	)
}
