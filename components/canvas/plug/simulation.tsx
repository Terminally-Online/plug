import type { FC } from "react"
import { memo, useState } from "react"

import { LightningBoltIcon, UpdateIcon } from "@radix-ui/react-icons"

import { Pin } from "@/lib/types"

export type PlugSimulationState = {
	success: boolean
}

export type PlugSimulationProps = {
	pins: Array<Pin>
	onSimulation: (state: PlugSimulationState) => void
}

export const PlugSimulation: FC<PlugSimulationProps> = ({
	pins,
	onSimulation
}) => {
	const [isLoading, setIsLoading] = useState(false)

	const endsWithThen = pins[pins.length - 1].type === "then"

	if (!endsWithThen) return null

	const handleClick = async () => {
		const request: Promise<{ success: boolean }> = new Promise(resolve => {
			// * Update the loading state.
			setIsLoading(true)

			// TODO: Remove this once the endpoint is live.
			// ? Temporary fake request to simulate a real request.
			setTimeout(() => {
				resolve({ success: true })
			}, 3000)
		})

		const response = await request

		setIsLoading(false)

		onSimulation(response)

		// TODO: Handle the rest of the stuff.
	}

	return (
		<button
			type="button"
			className="flex w-full flex-row items-center justify-center gap-2 rounded-sm border-[1px] border-stone-950 bg-stone-800 p-2 text-sm font-bold transition-all duration-200 ease-in-out hover:bg-white hover:text-stone-950 disabled:bg-white/60 disabled:text-stone-950"
			onClick={handleClick}
			disabled={isLoading}
		>
			{isLoading ? (
				<>
					<UpdateIcon className="h-3 w-3 animate-spin opacity-60" />
					Simulating
				</>
			) : (
				<>
					<LightningBoltIcon className="h-3 w-3 opacity-60" />
					Simulate
				</>
			)}
		</button>
	)
}

export default memo(PlugSimulation)
