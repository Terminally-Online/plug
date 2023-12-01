import { FC, memo } from 'react'

import { PlusIcon } from '@radix-ui/react-icons'

import type { Pin } from '@/lib/types'

export type PinAppendageProps = {
	pin: Pin
	onClick: () => void
	gridSize: number
	isAvailable?: boolean
}

export const PinAppendage: FC<PinAppendageProps> = ({
	pin,
	onClick,
	gridSize,
	isAvailable = true
}) => {
	return (
		<div className="relative flex flex-col items-center justify-center">
			<div
				className="w-[1px] bg-stone-950"
				style={{
					height: `${gridSize * 2}px`
				}}
			/>

			{isAvailable && pin.type === 'if' && (
				<button
					type="button"
					className="text-xs absolute border-[1px] border-stone-950 bg-stone-900 text-white/40 hover:bg-white hover:text-stone-950 rounded-full w-[20px] h-[20px] flex items-center justify-center transition-all duration-2 ease-in-out"
					onClick={onClick}
				>
					<PlusIcon width={10} height={10} />
				</button>
			)}
		</div>
	)
}

export default memo(PinAppendage)
