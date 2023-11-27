import {
	FC,
	memo,
	PointerEvent,
	Suspense,
	useEffect,
	useRef,
	useState,
	WheelEvent
} from 'react'

import {
	DndContext,
	KeyboardSensor,
	MouseSensor,
	TouchSensor,
	useSensor,
	useSensors
} from '@dnd-kit/core'
import { Coordinates } from '@dnd-kit/core/dist/types'
import useSize from '@react-hook/size'

import useRenderLoop from '@/lib/hooks/useRenderLoop'
import CanvasStore from '@/lib/store'

import {
	DraggableStory,
	SnapCenterToCursor,
	SnapToGrid
} from '../drag/DraggableStory'
import Canvas from './Canvas'
import { Draggable } from './Draggable'
import Droppable from './Droppable'

export type ViewportProps = {
	id: string
}

export const Viewport: FC<ViewportProps> = ({ id }) => {
	const canvasRef = useRef<HTMLDivElement>(null)
	const frame = useRenderLoop(60)

	const [width, height] = useSize(canvasRef)

	const handleWheel = (e: WheelEvent) => {
		e.stopPropagation()

		const deltaX = e.deltaX
		const deltaY = e.deltaY

		if (!e.ctrlKey) {
			CanvasStore.moveCamera(deltaX, deltaY)
		} else {
			CanvasStore.zoomCamera(deltaX, deltaY)
		}
	}

	const handlerPointerMove = (e: PointerEvent) => {
		CanvasStore.movePointer(e.clientX, e.clientY)
	}

	useEffect(() => {
		if (width === 0 || height === 0) return

		CanvasStore.initialize(width, height)
	}, [width, height])

	return (
		<div
			className="bg-stone-900 w-full h-full text-black dark:text-white overscroll-none w-full h-full relative overflow-hidden overscroll-none"
			ref={canvasRef}
			onWheel={handleWheel}
			onPointerMove={handlerPointerMove}
		>
			<SnapToGrid />
		</div>
	)
}

export default memo(Viewport)

// <Suspense fallback={<div>Loading...</div>}>
// 	<Canvas frame={frame} id={id} />
// </Suspense>
