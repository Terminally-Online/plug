import { useEffect } from 'react'

import { HomeIcon, SewingPinIcon } from '@radix-ui/react-icons'

import CanvasStore from '@/lib/store'

export const Toolbar = () => {
	const zoom = Math.round(CanvasStore.scale.x * 100)

	const handleCenter = () => {
		CanvasStore.centerCamera()
	}

	useEffect(() => {
		const handleZoom = (e: KeyboardEvent) => {
			const keys = [
				['+', -10],
				['-', 10],
				['=', -10]
			]

			for (const [key, value] of keys) {
				if (e.key === key) {
					e.preventDefault()

					const zoom =
						typeof value === 'number' ? value : parseInt(value)

					CanvasStore.zoomCamera(zoom, zoom)
				}
			}
		}

		window.addEventListener('keydown', handleZoom)

		return () => window.removeEventListener('keydown', handleZoom)
	}, [])

	return (
		<>
			<div className="fixed top-8 left-0 right-0 border-b-[1px] border-stone-950">
				<div className="relative bg-stone-900 left-0 text-white flex flex-row items-stretch">
					<button
						className="group p-4 w-min ml-auto text-center text-xs border-l-[1px] border-stone-950 text-white/60 tabular-nums hover:bg-white hover:text-stone-950 transition-all duration-200 ease-in-out"
						onClick={handleCenter}
					>
						<SewingPinIcon
							className="opacity-60 group:hover:opacity-100"
							width={16}
							height={16}
						/>
					</button>

					<p className="p-4 text-center text-xs border-l-[1px] border-stone-950 text-white/60 tabular-nums">
						{zoom}%
					</p>
				</div>
			</div>

			<div className="bg-stone-900 fixed bottom-0 left-0 right-0 flex flex-row gap-2 items-center border-t-[1px] border-stone-950">
				<div className="mx-auto text-white p-4">
					<p className="text-xs opacity-60">
						Tip: Double click anywhere to start a new plug
					</p>
				</div>
			</div>

			{
				<div className="fixed top-24 right-0 bg-red-400 text-red-700 font-bold p-2 m-2 z-10 rounded-sm">
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
		</>
	)
}

export default Toolbar
