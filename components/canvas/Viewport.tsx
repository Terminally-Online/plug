'use client'

import Canvas from './Canvas'
import useRenderLoop from '@/lib/hooks/useRenderLoop'

import {
	FC,
	memo,
	PointerEvent,
	Suspense,
	useEffect,
	useRef,
	WheelEvent
} from 'react'
import { DndProvider } from 'react-dnd'
import { HTML5Backend } from 'react-dnd-html5-backend'

import CanvasStore from '@/lib/store'

import useSize from '@react-hook/size'

export type ViewportProps = {
	id: string
}

const Viewport: FC<ViewportProps> = ({ id }) => {
	const canvasRef = useRef<HTMLDivElement>(null)

	const frame = useRenderLoop(60)
	const [width, height] = useSize(canvasRef)

	const wheelListener = (e: WheelEvent) => {
		e.stopPropagation()

		const friction = 1
		const event = e as WheelEvent
		const deltaX = event.deltaX * friction
		const deltaY = event.deltaY * friction

		if (!event.ctrlKey) {
			CanvasStore.moveCamera(deltaX, deltaY)
		} else {
			CanvasStore.zoomCamera(deltaX, deltaY)
		}
	}

	const pointerListener = (event: PointerEvent) => {
		CanvasStore.movePointer(event.clientX, event.clientY)
	}

	useEffect(() => {
		if (width === 0 || height === 0) return

		CanvasStore.initialize(width, height)
	}, [width, height])

	return (
		<div className="bg-stone-900 w-full h-full text-black dark:text-white">
			<div
				className="w-full h-full relative overflow-hidden overscroll-none"
				ref={canvasRef}
				onWheel={wheelListener}
				onPointerMove={pointerListener}
			>
				<DndProvider backend={HTML5Backend}>
					<Suspense fallback={<div>Loading...</div>}>
						<Canvas frame={frame} id={id} />
					</Suspense>
				</DndProvider>
			</div>
		</div>
	)
}

export default memo(Viewport)
