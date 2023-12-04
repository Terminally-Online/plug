import { FC, memo } from "react"

import { PlusIcon } from "@radix-ui/react-icons"

import type { Pin } from "@/lib/types"

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

			{isAvailable && pin.type === "if" && (
				<button
					type="button"
					className="duration-2 absolute flex h-[20px] w-[20px] items-center justify-center rounded-full border-[1px] border-stone-950 bg-stone-900 text-xs text-white/40 transition-all ease-in-out hover:bg-white hover:text-stone-950"
					onClick={onClick}
				>
					<PlusIcon width={10} height={10} />
				</button>
			)}
		</div>
	)
}

export default memo(PinAppendage)
