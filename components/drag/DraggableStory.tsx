import React, { useState } from 'react'

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

interface Props {
	initialCoordinates?: Coordinates
	activationConstraint?: PointerActivationConstraint
	axis?: Axis
	handle?: boolean
	modifiers?: Modifiers
	buttonStyle?: React.CSSProperties
	style?: React.CSSProperties
}

export function DraggableStory({
	initialCoordinates,
	activationConstraint,
	axis,
	handle,
	modifiers,
	style,
	buttonStyle
}: Props) {
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
				console.log(active)
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
						axis={axis}
						handle={handle}
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

interface DraggableItemProps {
	id: string
	handle?: boolean
	style?: React.CSSProperties
	buttonStyle?: React.CSSProperties
	axis?: Axis
	top?: number
	left?: number
}

function DraggableItem({
	axis,
	id,
	style,
	top,
	left,
	handle,
	buttonStyle
}: DraggableItemProps) {
	const { attributes, isDragging, listeners, setNodeRef, transform } =
		useDraggable({
			id
		})

	return (
		<Draggable
			id={id}
			ref={setNodeRef}
			dragging={isDragging}
			handle={handle}
			listeners={listeners}
			style={{ ...style, top, left }}
			buttonStyle={buttonStyle}
			transform={transform}
			axis={axis}
			{...attributes}
		/>
	)
}
