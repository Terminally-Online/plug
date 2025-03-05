import { FC, useCallback, useMemo } from "react"

import { Sentence } from "@/components/app/plugs/sentences/sentence"
import { useConnect, ActionSchemaCoils, InputValue } from "@/lib"
import { useActions } from "@/state/actions"
import { columnByIndexAtom } from "@/state/columns"
import { useAtom, useSetAtom } from "jotai"
import { editPlugAtom, plugByIdAtom, plugsAtom } from "@/state/plugs"
import { api } from "@/server/client"
import { DragDropContext, Draggable, Droppable, DropResult } from "@hello-pangea/dnd"
import { Callout } from "@/components/app/utils/callout"

type SentenceProps = { index: number }

export const Sentences: FC<SentenceProps> = ({ index }) => {
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

	const availableCoils = useMemo(() => {
		if (!plug || !solverActions) return {}
		
		const coils: Record<string, {type: string, actionIndex: number, coil: any}> = {}
		
		plug.actions.forEach((action, actionIndex) => {
			const actionSchema = solverActions[action.protocol]?.schema[action.action]
			if (!actionSchema || !actionSchema.coils) return
			
			actionSchema.coils.forEach(coil => {
				if (coil.slice?.name) {
					coils[coil.slice.name] = {
						type: coil.slice.type,
						actionIndex,
						coil
					}
				}
			})
		})
		
		return coils
	}, [plug, solverActions])
	
	const handleValueChange = (actionIndex: number, inputIndex: string, value: string, additionalData: any = {}) => {
		if (!plug) return
		
		const isLinked = value.startsWith("<-{") && value.endsWith("}")
		
		if (isLinked) {
			const coilName = value.substring(3, value.length - 1)
			const coilInfo = availableCoils[coilName]
			
			if (!coilInfo || coilInfo.actionIndex >= actionIndex) {
				console.error("Invalid coil reference or circular dependency", {
					coilName,
					actionIndex,
					availableCoils
				})
				return
			}
			
			const actionSchema = solverActions[plug.actions[actionIndex].protocol]?.schema[plug.actions[actionIndex].action]
			if (actionSchema) {
				const inputType = actionSchema.sentence
					.split('{')
					.filter(part => part.includes('}'))
					.map(part => {
						const match = part.match(/(\d+)<([^:]+):([^>]+)>/)
						if (match && match[1] === inputIndex) {
							return match[3] 
						}
						return null
					})
					.filter(Boolean)[0]
				
				if (inputType && coilInfo.type && !isTypeCompatible(inputType, coilInfo.type)) {
					console.error(`Type mismatch: Input expects ${inputType} but coil provides ${coilInfo.type}`)
				}
			}
		}
		
		const newActions = [...plug.actions]
		const action = newActions[actionIndex]
		
		if (!action.values) {
			action.values = {}
		}
		
		const isNumber = additionalData.isNumber || false
		
		action.values[inputIndex] = {
			index: inputIndex,
			key: inputIndex,
			name: inputIndex,
			value: isNumber && !isLinked ? parseFloat(value) : value,
			...additionalData
		}
		
		setPlugs(prev => prev.map(p => plug && p.id === column?.item ? { ...p, actions: JSON.stringify(newActions), updatedAt: new Date() } : p))
		edit({
			id: plug.id,
			actions: JSON.stringify(newActions)
		})
	}
	
	// TODO: This is a simplified version - real implementation would be more thorough
	const isTypeCompatible = (inputType: string, coilType: string): boolean => {
		if (inputType === coilType) return true
		
		if (inputType.includes('uint') || inputType.includes('int')) {
			return coilType.includes('uint') || coilType.includes('int')
		}
		
		if (inputType === 'address') {
			return coilType === 'address'
		}
		
		if (inputType.includes('bytes') || inputType === 'string') {
			return coilType.includes('bytes') || coilType === 'string'
		}
		
		return false
	}
	
	const validateType = (coilName: string, expectedType: string): boolean => {
		const coilInfo = availableCoils[coilName]
		if (!coilInfo) return false
		
		return coilInfo.type === expectedType
	}

	const handleDragEnd = (result: DropResult) => {
		if (!result.destination || !plug || !own) return

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

	return <div className="mb-72 flex flex-col">
		<Callout.EmptyPlug index={index} isEmpty={plug.actions.length === 0} />
		
		<DragDropContext onDragEnd={handleDragEnd}>
			<Droppable droppableId={`${index}-${plug.id}-items`} direction="vertical">
				{provided => (
					<div ref={provided.innerRef} className="flex flex-col" {...provided.droppableProps}>
						{plug.actions.map((action, actionIndex) => {
							const values = Object.values(plug.actions[actionIndex + 1]?.values ?? {})
							const linked = values.filter(val => val?.value?.startsWith("<-{"))

							let prevCoils: ActionSchemaCoils | undefined = []
							if (actionIndex > 0) {
								const prevAction = plug.actions[actionIndex - 1]
								prevCoils = solverActions[prevAction.protocol]?.schema[prevAction.action]?.coils || []
							}

							return <Draggable
								key={String(action.id)}
								draggableId={String(action.id)}
								index={actionIndex}
							>
								{(provided, snapshot) => (
									<div
										ref={provided.innerRef}
										className="flex h-full w-full flex-row rounded-lg"
										{...provided.draggableProps}
										style={{
											...provided.draggableProps.style,
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
												handleValueChange(actionIndex, inputIndex, value, additionalData)
											}
											validateType={validateType}
											availableCoils={availableCoils}
											{...provided.dragHandleProps}
										/>
									</div>
								)}
							</Draggable>
						})}

						{provided.placeholder}
					</div>
				)}
			</Droppable>
		</DragDropContext>
	</div>
}
