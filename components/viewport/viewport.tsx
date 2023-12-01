import {
	FC,
	memo,
	PointerEvent,
	Suspense,
	useEffect,
	useRef,
	WheelEvent
} from 'react'

import useSize from '@react-hook/size'

import useRenderLoop from '@/lib/hooks/useRenderLoop'
import CanvasStore from '@/lib/store'

import Canvas from '../canvas/canvas'

export type ViewportProps = {
	id: string
}

// * This component handles the render loop and the camera interactions.
export const Viewport: FC<ViewportProps> = ({ id }) => {
	const canvasRef = useRef<HTMLDivElement>(null)
	// ? While this appears to do nothing, it is actually used to trigger a
	//   re-render on the component that hold it as a prop. If it is not
	//   assigned to a prop, then it does do nothing.
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
			className="w-screen h-full text-black dark:text-white overflow-hidden overscroll-none"
			ref={canvasRef}
			onWheel={handleWheel}
			onPointerMove={handlerPointerMove}
		>
			<Suspense fallback={<div>Loading...</div>}>
				<Canvas frame={frame} id={id} />
			</Suspense>
		</div>
	)
}

export default memo(Viewport)
