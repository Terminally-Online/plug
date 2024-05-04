import type { FC, PropsWithChildren } from "react"
import React, { useEffect, useMemo, useState } from "react"

import type { DragEndEvent } from "@dnd-kit/core"
import {
	DndContext,
	KeyboardSensor,
	MouseSensor,
	PointerActivationConstraint,
	TouchSensor,
	useSensor,
	useSensors
} from "@dnd-kit/core"
import { createSnapModifier } from "@dnd-kit/modifiers"

// import { Component } from "@prisma/client"
import { api } from "@/lib/api"

import { DraggableComponent } from "./draggable-component"

// export type DraggableComponentsProps = {
// 	id: string
// 	initialComponents: Record<string, Component>
// 	gridSize: number
// 	activationConstraint?: PointerActivationConstraint
// }
//
// export type DraggableMoveProps = {
// 	id: string
// 	top?: number
// 	left?: number
// }

export const DraggableComponents = () => <></>

// export const DraggableComponents: FC<
// 	PropsWithChildren<DraggableComponentsProps>
// > = ({ id, initialComponents, activationConstraint, gridSize }) => {
// 	const [components, setComponents] = useState(initialComponents)
//
// 	const moveComponent = api.canvas.component.move.useMutation({
// 		onMutate(componentPosition) {
// 			const { component } = componentPosition
// 			const { id, top, left } = component
// 			handleMove({ id, top, left })
// 		}
// 	})
//
// 	api.canvas.component.onMove.useSubscription(id, {
// 		onData(component) {
// 			console.log(component)
// 		}
// 	})
//
// 	const snapToGrid = useMemo(() => createSnapModifier(gridSize), [gridSize])
//
// 	const mouseSensor = useSensor(MouseSensor, {
// 		activationConstraint
// 	})
// 	const touchSensor = useSensor(TouchSensor, {
// 		activationConstraint
// 	})
// 	const keyboardSensor = useSensor(KeyboardSensor, {})
// 	const sensors = useSensors(mouseSensor, touchSensor, keyboardSensor)
//
// 	const handleMove = ({ id, top, left }: DraggableMoveProps) => {
// 		setComponents(previousComponents => ({
// 			...previousComponents,
// 			[id]: {
// 				...previousComponents[id],
// 				top: top ?? previousComponents[id].top,
// 				left: left ?? previousComponents[id].left
// 			}
// 		}))
// 	}
//
// 	const handleDrag = ({ active, delta }: DragEndEvent) => {
// 		const { top, left } = components[active.id]
//
// 		moveComponent.mutate({
// 			id,
// 			component: {
// 				id: active.id as string,
// 				top: top + delta.y,
// 				left: left + delta.x
// 			}
// 		})
// 	}
//
// 	// * Re-render when a parent component updates the base components.
// 	// ? Is assumed to be the head of the canvas changes.
// 	useEffect(() => {
// 		setComponents(initialComponents)
// 	}, [initialComponents])
//
// 	return (
// 		<DndContext
// 			sensors={sensors}
// 			onDragEnd={handleDrag}
// 			modifiers={[snapToGrid]}
// 		>
// 			{Object.keys(components).map(key => {
// 				const { left, top } = components[key]
//
// 				return (
// 					<DraggableComponent
// 						key={key}
// 						id={key}
// 						left={left}
// 						top={top}
// 						gridSize={gridSize}
// 					/>
// 				)
// 			})}
// 		</DndContext>
// 	)
// }
