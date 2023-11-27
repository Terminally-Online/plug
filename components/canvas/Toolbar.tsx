import { useEffect } from 'react'

import CanvasStore from '@/lib/store'

export const Toolbar = () => {
	const zoom = Math.round(CanvasStore.scale.x * 100)

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
			<div className="fixed top-8 left-0 right-0 bg-red-300 border-b-[1px] border-stone-950">
				<div className="relative bg-stone-900 left-0 text-white flex flex-row items-stretch">
					<p className="p-4 ml-auto text-center text-xs border-l-[1px] border-stone-950 text-white/60 tabular-nums">
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
		</>
	)
}

export default Toolbar
