'use client'

import Pin from './Pin'
import PinAppendage from './PinAppendage'
import PlugSimulation, { PlugSimulationState } from './PlugSimulation'

import type { FC, PropsWithChildren } from 'react'
import { memo, useCallback, useState } from 'react'

import { t } from '@/app/api/trpc/client'
import { pins } from '@/lib/constants'
import type { Pin as PinType } from '@/lib/types'
import { cn } from '@/lib/utils'

export type PlugProps = {
	id: string
	selecting: string | null
	preview?: boolean
} & JSX.IntrinsicElements['div']

export const Plug: FC<PropsWithChildren<PlugProps>> = ({
	id,
	children,
	preview,
	selecting,
	...rest
}) => {
	// * Deconstruct the values that were sent from the database.
	// TODO: Acknowledge that this is dangerous and not ideal.
	const values = JSON.parse(children as string) as Array<PinType>

	const [selectedPins, setSelectedPins] = useState([pins[0].pins[0]])
	const [simulation, setSimulation] = useState<PlugSimulationState | null>(
		null
	)

	const handlePost = () => {
		id
		// TODO: Post into the database when selectedPins is updated.
		//    NOTES: This will only orchestrate the state of the plug and not the actual canvas state of the plug (position, size, etc.)
		//    NOTES: Also do note that we cannot simply retrieve the state of the Plug from here because to have gotten here we needed to already know its position in order to properly place it on the Canvas.
	}

	// * Remove the selectedPins from the pins so that you can only choose each pin once.
	const availablePins = pins
		.map(pin => ({
			...pin,
			pins: pin.pins.filter(pin => !selectedPins.includes(pin))
		}))
		.filter(pin => pin.pins.length > 0)

	const handleChange = (index: number, pin: PinType) => {
		setSelectedPins(previousSelectedPins => {
			const newSelectedPins = [...previousSelectedPins]
			newSelectedPins[index] = pin

			// * If we have terminated the chain remove trailing pins.
			if (pin.type === 'then')
				newSelectedPins.splice(
					index + 1,
					newSelectedPins.length - index
				)

			return newSelectedPins
		})
	}

	const handleAddition = useCallback(
		(index: number) => {
			setSelectedPins(previousSelectedPins => {
				const newSelectedPins = [...previousSelectedPins]

				newSelectedPins.splice(index + 1, 0, availablePins[0].pins[0])

				return newSelectedPins
			})
		},
		[availablePins]
	)

	const handleSimulation = (state: PlugSimulationState) => {
		// TODO: Not sure what is done here yet, but something.
		setSimulation(state)
	}

	return (
		<div
			className={cn(
				'bg-stone-900 text-white cursor-move flex flex-col items-center justify-center',
				selecting ? 'border-[1px] border-red-500' : ''
			)}
			role={preview ? 'PlugPreview' : 'Plug'}
			{...rest}
		>
			{selectedPins.map((pin, index) => (
				<div key={index} className="w-full h-full">
					<Pin
						selectedPin={pin}
						pins={availablePins}
						onPinChange={newPin => handleChange(index, newPin)}
					/>

					<PinAppendage
						pin={pin}
						onClick={() => handleAddition(index)}
						isAvailable={availablePins.length > 0}
					/>
				</div>
			))}

			<PlugSimulation
				pins={selectedPins}
				onSimulation={handleSimulation}
			/>
		</div>
	)
}

export default memo(Plug)
