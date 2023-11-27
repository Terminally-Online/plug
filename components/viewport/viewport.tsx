import {
	FC,
	memo,
	PointerEvent,
	Suspense,
	useEffect,
	useMemo,
	useRef,
	useState,
	WheelEvent
} from 'react'

import { createSnapModifier } from '@dnd-kit/modifiers'
import useSize from '@react-hook/size'

import useRenderLoop from '@/lib/hooks/useRenderLoop'
import CanvasStore from '@/lib/store'

import Canvas from '../canvas/Canvas'
import { Position } from '../canvas/Position'
import { DraggableStory, SnapToGrid } from '../drag/DraggableStory'
import { Grid } from '../drag/grid/grid'

export type ViewportProps = {
	id: string
}

export const Viewport: FC<ViewportProps> = ({ id }) => {
	const canvasRef = useRef<HTMLDivElement>(null)
	// ? While this appears to do nothing, it is actually used to trigger a
	//   re-render on the component that hold it as a prop. If it is not
	//   assigned to a prop, then it does do nothing.
	const frame = useRenderLoop(60)

	const [width, height] = useSize(canvasRef)

	const [gridSize, setGridSize] = useState(30)

	const style = {
		alignItems: 'flex-start'
	}

	const buttonStyle = {
		marginLeft: gridSize - 20 + 1,
		marginTop: gridSize - 20 + 1,
		width: gridSize * 8 - 1,
		height: gridSize * 2 - 1
	}

	const snapToGrid = useMemo(() => createSnapModifier(gridSize), [gridSize])

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
			className="w-full h-full text-black dark:text-white overflow-hidden overscroll-none"
			ref={canvasRef}
			onWheel={handleWheel}
			onPointerMove={handlerPointerMove}
		>
			<Suspense fallback={<div>Loading...</div>}>
				<Canvas frame={frame} id={id}>
					<DraggableStory
						modifiers={[snapToGrid]}
						style={style}
						buttonStyle={buttonStyle}
						key={gridSize}
					/>
				</Canvas>

				<Position id="grid" left={0} top={0} width={800} height={800}>
					<Grid size={gridSize} onSizeChange={setGridSize} />
				</Position>
			</Suspense>
			.{' '}
			{
				<div className="fixed top-0 right-0 text-red-700 bg-red-400 text-red-700 font-bold p-2 m-2 z-10 rounded-sm">
					<p>User id: tester</p>

					<p>
						Camera: {Math.round(CanvasStore.camera.x)},{' '}
						{Math.round(CanvasStore.camera.y)},{' '}
						{Math.round(CanvasStore.camera.z)}
					</p>
					<p>
						Container: {CanvasStore.container.width} x{' '}
						{CanvasStore.container.height}
					</p>
					<p>Locked: {CanvasStore.camera.locked.toString()}</p>
					<p>
						Scale: {Math.round(CanvasStore.scale.x)},{' '}
						{Math.round(CanvasStore.scale.y)}
					</p>
					<p>
						Screen: {Math.round(CanvasStore.screen.x)},{' '}
						{Math.round(CanvasStore.screen.y)}
					</p>
					<p>
						Pointer: {Math.round(CanvasStore.pointer.x)},{' '}
						{Math.round(CanvasStore.pointer.y)}
					</p>
				</div>
			}
		</div>
	)
}

export default memo(Viewport)
