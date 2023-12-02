import { useEffect } from 'react'

import { SewingPinIcon } from '@radix-ui/react-icons'

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
			<div className="fixed left-0 right-0 top-12 border-b-[1px] border-stone-950">
				<div className="relative left-0 flex flex-row items-stretch bg-stone-900 text-white">
					<button
						className="group ml-auto w-min border-l-[1px] border-stone-950 p-4 text-center text-xs tabular-nums text-white/60 transition-all duration-200 ease-in-out hover:bg-white hover:text-stone-950"
						onClick={handleCenter}
					>
						<SewingPinIcon
							className="group:hover:opacity-100 opacity-60"
							width={16}
							height={16}
						/>
					</button>

					<p className="border-l-[1px] border-stone-950 p-4 text-center text-xs tabular-nums text-white/60">
						{zoom}%
					</p>
				</div>
			</div>

			<div className="fixed bottom-0 left-0 right-0 flex flex-row items-center gap-2 border-t-[1px] border-stone-950 bg-stone-900">
				<div className="mx-auto p-4 text-white">
					<p className="text-xs opacity-60">
						Tip: Double click anywhere to start a new plug
					</p>
				</div>
			</div>

			{
				<div className="fixed right-0 top-24 z-10 m-2 rounded-sm bg-red-400 p-2 font-bold text-red-700">
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
