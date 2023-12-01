import type { CSSProperties, FC, PropsWithChildren } from 'react'
import React, { forwardRef, useEffect, useMemo, useState } from 'react'

import classNames from 'classnames'

import type { DragEndEvent, DraggableSyntheticListeners } from '@dnd-kit/core'
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

import Plug from '@/components/canvas/blocks/Plug'
import { api } from '@/lib/api'
import { inBounds } from '@/lib/functions/math-utils'
import CanvasStore from '@/lib/store'
import { cn } from '@/lib/utils'

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
		marginTop: gridSize - 0 + 1,
		width: gridSize * 8 - 1,
		height: gridSize * 3 - 1
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

	const handleDrag = ({ active, delta }: DragEndEvent) => {
		const { top, left } = components[active.id]

		moveComponent.mutate({
			id,
			component: {
				id: active.id as string,
				top: top + delta.y,
				left: left + delta.x
			}
		})
	}

	// * Re-render when a parent component updates the base components.
	// ? Is assumed to be the head of the canvas changes.
	useEffect(() => {
		setComponents(initialComponents)
	}, [initialComponents])

	return (
		<DndContext
			sensors={sensors}
			onDragEnd={handleDrag}
			modifiers={[snapToGrid]}
		>
			{Object.keys(components).map(key => {
				const { left, top } = components[key]

				if (
					!inBounds(
						{
							left,
							top,
							width: buttonStyle.width,
							height: buttonStyle.height
						},
						{
							left: CanvasStore.screen.x,
							top: CanvasStore.screen.y,
							width: screen.width,
							height: screen.height
						}
					)
				)
					return null

				return (
					<DraggableComponent
						key={key}
						id={key}
						top={top - CanvasStore.screen.y}
						left={left - CanvasStore.screen.x}
						width={gridSize * 8 - 1}
						height={gridSize * 3 - 1}
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
	left,
	width,
	height
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
			style={{
				position: 'absolute',
				alignItems: 'flex-start',
				top,
				left
			}}
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

const parse = (value: string | number | undefined) => {
	if (typeof value === 'number') return value

	return parseInt(value ?? '0')
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
		const left = parse(style?.left)
		const top = parse(style?.top)

		return (
			<div
				id={id}
				className={cn(
					'items-start w-max',
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
					className="absolute inline-block flex items-center content-center appearance-none border-none outline-none"
					style={buttonStyle}
					aria-label="Draggable"
					data-cypress="draggable-item"
				>
					<Plug id={'test'} selecting={null}>
						{'[]'}
					</Plug>

					<p className="absolute top-[110%] left-0 bg-red-400 text-red-700 font-bold rounded-sm p-2">
						{buttonStyle?.width} x {buttonStyle?.height} @
						{Math.round(left)} x {Math.round(top)}{' '}
					</p>
				</button>
			</div>
		)
	}
)
