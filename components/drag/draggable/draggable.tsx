import type { CSSProperties, FC, PropsWithChildren } from 'react'
import React, { forwardRef, useMemo, useState } from 'react'

import classNames from 'classnames'

import type { DraggableSyntheticListeners } from '@dnd-kit/core'
import {
	DndContext,
	KeyboardSensor,
	MouseSensor,
	PointerActivationConstraint,
	TouchSensor,
	useDraggable,
	useSensor,
	useSensors
} from '@dnd-kit/core'
import { createSnapModifier } from '@dnd-kit/modifiers'
import type { Transform } from '@dnd-kit/utilities'
import { Component } from '@prisma/client'

import { api } from '@/lib/api'
import CanvasStore from '@/lib/store'
import { cn } from '@/lib/utils'

import { draggable } from './draggable-svg'
import styles from './draggable.module.css'

export type DraggableComponentsProps = {
	id: string
	initialComponents: Record<string, Component>
	gridSize: number
	activationConstraint?: PointerActivationConstraint
}

export type DraggableMoveProps = {
	id: string
	top: number
	left: number
}

export const DraggableComponents: FC<
	PropsWithChildren<DraggableComponentsProps>
> = ({ id, initialComponents, activationConstraint, gridSize }) => {
	const [components, setComponents] = useState(initialComponents)

	const moveComponent = api.canvas.component.move.useMutation({
		onMutate(componentPosition) {
			const { component } = componentPosition
			const { id, top, left } = component
			handleMove({ id, top, left })
		}
	})

	const snapToGrid = useMemo(() => createSnapModifier(gridSize), [gridSize])

	const buttonStyle = {
		marginLeft: gridSize - 20 + 1,
		marginTop: gridSize - 20 + 1,
		width: gridSize * 12 - 1,
		height: gridSize * 4 - 1
	}

	const mouseSensor = useSensor(MouseSensor, {
		activationConstraint
	})
	const touchSensor = useSensor(TouchSensor, {
		activationConstraint
	})
	const keyboardSensor = useSensor(KeyboardSensor, {})
	const sensors = useSensors(mouseSensor, touchSensor, keyboardSensor)

	const handleMove = ({ id, top, left }: DraggableMoveProps) => {
		setComponents(previousComponents => ({
			...previousComponents,
			[id]: {
				...previousComponents[id],
				top,
				left
			}
		}))
	}

	return (
		<DndContext
			sensors={sensors}
			onDragEnd={props => {
				const { active, delta } = props
				const { top, left } = components[active.id]
				moveComponent.mutate({
					id,
					component: {
						id: active.id as string,
						top: top + delta.y,
						left: left + delta.x
					}
				})
			}}
			modifiers={[snapToGrid]}
		>
			{Object.keys(components).map(key => {
				const { left, top } = components[key]

				return (
					<DraggableComponent
						id={key}
						top={top - CanvasStore.screen.y}
						left={left - CanvasStore.screen.x}
						width={gridSize * 12 - 1}
						height={gridSize * 4 - 1}
						buttonStyle={buttonStyle}
					/>
				)
			})}
		</DndContext>
	)
}

export type DraggableComponentProps = {
	id: string
	buttonStyle?: CSSProperties
	top: number
	left: number
	width: number
	height: number
}

export const DraggableComponent: FC<DraggableComponentProps> = ({
	id,
	buttonStyle,
	top,
	left
}) => {
	const { attributes, isDragging, listeners, setNodeRef, transform } =
		useDraggable({
			id
		})

	return (
		<Draggable
			id={id}
			ref={setNodeRef}
			dragging={isDragging}
			listeners={listeners}
			style={{ alignItems: 'flex-start', top, left }}
			buttonStyle={buttonStyle}
			transform={transform}
			{...attributes}
		/>
	)
}

export type DraggableProps = {
	id: string
	dragOverlay?: boolean
	dragging?: boolean
	listeners?: DraggableSyntheticListeners
	style?: CSSProperties
	buttonStyle?: CSSProperties
	transform?: Transform | null
}

export const Draggable = forwardRef<HTMLButtonElement, DraggableProps>(
	function Draggable(
		{
			id,
			dragOverlay,
			dragging,
			listeners,
			transform,
			style,
			buttonStyle,
			...props
		},
		ref
	) {
		return (
			<div
				id={id}
				className={cn(
					'items-start',
					classNames(
						styles.Draggable,
						dragOverlay && styles.dragOverlay,
						dragging && styles.dragging
					)
				)}
				style={
					{
						...style,
						'--translate-x': `${transform?.x ?? 0}px`,
						'--translate-y': `${transform?.y ?? 0}px`
					} as CSSProperties
				}
			>
				<button
					{...props}
					{...listeners}
					ref={ref}
					style={buttonStyle}
					aria-label="Draggable"
					data-cypress="draggable-item"
				>
					{draggable}
				</button>
			</div>
		)
	}
)
