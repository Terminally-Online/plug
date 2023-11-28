import React, { FC, PropsWithChildren, useState } from 'react'

import {
	DndContext,
	KeyboardSensor,
	Modifiers,
	MouseSensor,
	PointerActivationConstraint,
	TouchSensor,
	useDraggable,
	useSensor,
	useSensors
} from '@dnd-kit/core'
import type { Coordinates } from '@dnd-kit/utilities'

import { Axis, Draggable } from '@/components/drag/draggable/draggable'
import { Wrapper } from '@/components/drag/wrapper/wrapper'
import CanvasStore from '@/lib/store'

export type DraggableProps = {
	activationConstraint?: PointerActivationConstraint
	modifiers?: Modifiers
	buttonStyle?: React.CSSProperties
	style?: React.CSSProperties
}

export const DraggableComponents: FC<PropsWithChildren<DraggableProps>> = ({
	activationConstraint,
	modifiers,
	style,
	buttonStyle
}) => {
	const [coordinates, setCoordinates] = useState<Coordinates>({
		a: {
			x: 5000 - 120,
			y: 5000 - 60
		},
		b: {
			x: 5000 - 120,
			y: 5000 + 50
		}
	})
	const mouseSensor = useSensor(MouseSensor, {
		activationConstraint
	})
	const touchSensor = useSensor(TouchSensor, {
		activationConstraint
	})
	const keyboardSensor = useSensor(KeyboardSensor, {})
	const sensors = useSensors(mouseSensor, touchSensor, keyboardSensor)

	return (
		<DndContext
			sensors={sensors}
			onDragEnd={props => {
				const { active, delta } = props

				setCoordinates(previousCoordinates => {
					const { x, y } = previousCoordinates[active.id]

					return {
						...previousCoordinates,
						[active.id]: {
							x: x + delta.x,
							y: y + delta.y
						}
					}
				})
			}}
			modifiers={modifiers}
		>
			{Object.keys(coordinates).map(key => {
				const { x, y } = coordinates[key]

				return (
					<DraggableItem
						id={key}
						top={y - CanvasStore.screen.y}
						left={x - CanvasStore.screen.x}
						style={style}
						buttonStyle={buttonStyle}
					/>
				)
			})}
		</DndContext>
	)
}

export type DraggableItemProps = {
	id: string
	style?: React.CSSProperties
	buttonStyle?: React.CSSProperties
	top?: number
	left?: number
}

export const DraggableItem: FC<DraggableItemProps> = ({
	id,
	style,
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
			style={{ ...style, top, left }}
			buttonStyle={buttonStyle}
			transform={transform}
			{...attributes}
		/>
	)
}
