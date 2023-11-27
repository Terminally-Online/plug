import { memo, useState } from 'react'

import { MinusIcon, PlusIcon } from '@radix-ui/react-icons'

const max = 200
const min = 1

export const Toolbar = () => {
	const [zoom, setZoom] = useState(100)

	const canZoom = (zoom: number) => {
		if (zoom < min) return false
		if (zoom > max) return false
		return true
	}

	const handleZoom = (zoom: number) => {
		if (!canZoom(zoom)) return

		setZoom(zoom)
	}

	return (
		<div className="fixed bottom-0 left-0 right-0 flex flex-row gap-2 items-center border-t-[1px] border-stone-950">
			<div className="relative bg-stone-900 left-0 text-white flex flex-row items-center">
				<button
					type="button"
					className="p-2 text-white/60 hover:bg-white hover:text-stone-950 active:bg-white active:text-stone-950 transition-all duration-200 ease-in-out border-[1px] border-t-[0px] border-stone-950 cursor-pointer disabled:cursor-not-allowed"
					onClick={() => handleZoom(zoom - 10)}
					disabled={!canZoom(zoom - 10)}
				>
					<MinusIcon width={16} height={16} />
				</button>

				<input
					className="text-center text-xs w-[50px] h-full bg-stone-800 border-b-[1px] border-stone-950 text-white/60 outline-none"
					min="1"
					max="100"
					value={zoom}
					onChange={e => handleZoom(parseInt(e.target.value))}
				/>

				<button
					type="button"
					className="p-2 text-white/60 hover:bg-white hover:text-stone-950 active:bg-white active:text-stone-950 transition-all duration-200 ease-in-out border-[1px] border-t-[0px] border-stone-950"
				>
					<PlusIcon width={16} height={16} />
				</button>
			</div>

			<div className="mx-auto text-white p-2">
				<p className="text-xs opacity-60">
					Tip: Double click anywhere to start a new plug
				</p>
			</div>
		</div>
	)
}

export default memo(Toolbar)
