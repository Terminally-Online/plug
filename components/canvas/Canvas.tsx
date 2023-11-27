'use client'

import { type FC, useEffect, useMemo, useState } from 'react'

import { useDrop } from 'react-dnd'

import Plug from '@/components/canvas/blocks/Plug'
import { Drag } from '@/components/canvas/Drag'
import { Position } from '@/components/canvas/Position'
import Toolbar from '@/components/canvas/Toolbar'
import { useTabs } from '@/contexts/TabsProvider'
import { api } from '@/lib/api'
import { ItemTypes } from '@/lib/constants'
import { snapToGrid } from '@/lib/functions/snap-to-grid'
import CanvasStore from '@/lib/store'
import type { DragItem } from '@/lib/types'

export type CanvasProps = {
	frame: string
	id: string
}

export const Canvas: FC<CanvasProps> = ({ id }) => {
	const { handleAdd } = useTabs()

	const [initialCanvas] = api.canvas.get.useSuspenseQuery(id)
	const [canvas, setCanvas] = useState(initialCanvas)

	const addComponent = api.canvas.component.add.useMutation()
	const moveComponent = api.canvas.component.move.useMutation({
		onMutate(componentPosition) {
			const { component } = componentPosition

			const { id, top, left } = component

			handleMove({ id, top, left })
		}
	})

	const handleMove = ({ id, top, left }: {
		id: string
		top: number
		left: number
	}) => {
		setCanvas(prevCanvas => {
			if (!prevCanvas) return prevCanvas

			const index = prevCanvas.components.findIndex(c => c.id === id)

			if (index === -1) return prevCanvas

			const newCanvas = { ...prevCanvas }
			newCanvas.components[index] = {
				...newCanvas.components[index],
				top,
				left
			}

			return newCanvas
		})
	}

	const components = useMemo(() => {
		return canvas.components.reduce(
			(acc, component) => {
				acc[component.id] = component
				return acc
			},
			{} as Record<string, (typeof canvas.components)[0]>
		)
	}, [canvas])

	useEffect(() => {
		if (!canvas) return

		handleAdd({
			label: canvas.name,
			color: canvas.color,
			href: `/canvas/${canvas.id}`,
			active: true
		})
	}, [canvas, handleAdd])

	const [, drop] = useDrop(
		() => ({
			accept: [ItemTypes.Box, ItemTypes.Markdown],
			drop(item: DragItem, monitor) {
				const delta = monitor.getDifferenceFromInitialOffset()

				if (!delta) return

				let { id, top, left } = item

				left = Math.round(left + delta.x)
				top = Math.round(top + delta.y)

				if (snapToGrid) [left, top] = snapToGrid(left, top)

				moveComponent.mutate({
					id: canvas.id,
					component: { id, left, top }
				})
			}
		}),
		[moveComponent]
	)

	const handleClick = (e: React.MouseEvent<HTMLDivElement>) => {
		const isShiftClick = e.shiftKey
		const isCtrlClick = e.ctrlKey
		const isCommandClick = e.metaKey
		const isAltClick = e.altKey

		const ready =
			isShiftClick || isCommandClick || isCtrlClick || isAltClick

		if (!ready) return

		const width = 400
		const left = CanvasStore.pointer.x - width / 2
		const top = CanvasStore.pointer.y - width / 4
		const type = ItemTypes.Plug

		addComponent.mutate({
			id,
			component: {
				left,
				top,
				type,
				width,
				height: width,
				content: ''
			}
		})
	}

	if (!canvas) return null

	return (
		<>
			<div
				ref={drop}
				className="relative w-screen h-screen overscroll-none"
				style={{
					transform: `scale(${
						(CanvasStore.scale.x, CanvasStore.scale.y)
					})`,
					transformOrigin: 'top left'
				}}
				onClick={handleClick}
			>
				{components &&
					Object.keys(components).map(key => {
						const componentTypes = {
							[ItemTypes.Plug]: Plug
						}

						const component = components[key]
						const Component = componentTypes[component.type]

						return (
							<Position key={component.id} {...component}>
								<Plug
									id={component.id}
									// onClick={() => {
									// 	isSelecting.mutate({
									// 		id: canvas.id,
									// 		component: {
									// 			id: component.id,
									// 			selecting: username
									// 		}
									// 	})
									// }}
									selecting={component.selectingId}
								>
									{JSON.stringify(component.content)}
								</Plug>
							</Position>
						)
					})}

				<Drag />
			</div>

			<Toolbar />
		</>
	)
}

export default Canvas
