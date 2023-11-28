import type { FC } from 'react'
import { useEffect, useMemo, useState } from 'react'

import { createSnapModifier } from '@dnd-kit/modifiers'

import Toolbar from '@/components/viewport/toolbar'
import { useTabs } from '@/contexts/TabsProvider'
import { api } from '@/lib/api'
import { ItemTypes } from '@/lib/constants'
import useRenderLoop from '@/lib/hooks/useRenderLoop'
import CanvasStore from '@/lib/store'

import { DraggableComponents } from '../drag/draggable/draggable'
import { Grid } from '../drag/grid/grid'
import { Scaler } from './scaler'

export type CanvasProps = {
	frame: string
	id: string
}

// * This component manages the base canvas and handles the addition of a new component.
export const Canvas: FC<CanvasProps> = ({ frame, id }) => {
	const { handleAdd } = useTabs()

	const [initialCanvas] = api.canvas.get.useSuspenseQuery(id)
	const [canvas, setCanvas] = useState(initialCanvas)
	const [gridSize, setGridSize] = useState(30)

	const addComponent = api.canvas.component.add.useMutation({
		onSuccess(component) {
			setCanvas(previousCanvas => {
				return {
					...previousCanvas,
					components: [...previousCanvas.components, component]
				}
			})
		}
	})

	const components = useMemo(() => {
		return canvas.components.reduce(
			(acc, component) => {
				acc[component.id] = component
				return acc
			},
			{} as Record<string, (typeof canvas.components)[0]>
		)
	}, [canvas])

	const handleClick = (e: React.MouseEvent<HTMLDivElement>) => {
		const isShiftClick = e.shiftKey
		const isCtrlClick = e.ctrlKey
		const isCommandClick = e.metaKey
		const isAltClick = e.altKey

		const ready =
			isShiftClick || isCommandClick || isCtrlClick || isAltClick

		if (!ready) return

		// TODO: This is not right :)
		const top =
			-70 + Math.round(CanvasStore.pointer.y / gridSize) * gridSize
		const left =
			-10 + Math.round(CanvasStore.pointer.x / gridSize) * gridSize

		const type = ItemTypes.Plug

		addComponent.mutate({
			id,
			component: {
				left,
				top,
				type,
				width: gridSize * 12 - 1,
				height: gridSize * 4 - 1,
				content: ''
			}
		})
	}

	useEffect(() => {
		if (!canvas) return

		handleAdd({
			label: canvas.name,
			color: canvas.color,
			href: `/canvas/${canvas.id}`,
			active: true
		})
	}, [canvas, handleAdd])

	if (!canvas) return null

	return (
		<>
			<Scaler onClick={handleClick}>
				<Grid size={gridSize} onSizeChange={setGridSize}>
					<DraggableComponents
						id={id}
						initialComponents={components}
						gridSize={gridSize}
					/>
				</Grid>
			</Scaler>

			<Toolbar />
		</>
	)
}

export default Canvas
