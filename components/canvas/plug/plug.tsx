import type { FC, PropsWithChildren } from "react"
import { useCallback, useState } from "react"

import { pins } from "@/lib/constants"
import type { Pin as PinType } from "@/lib/types"
import { cn } from "@/lib/utils"

import PinAppendage from "../pin/appendage"
import Pin from "../pin/pin"
import PlugSimulation from "./simulation"

export type PlugProps = {
	id: string
	selecting: string | null
	gridSize: number
	preview?: boolean
} & JSX.IntrinsicElements["div"]

export const Plug: FC<PropsWithChildren<PlugProps>> = ({
	id,
	children,
	preview,
	gridSize,
	selecting,
	...rest
}) => {
	// * Deconstruct the values that were sent from the database.
	// TODO: Acknowledge that this is dangerous and not ideal.
	const values = JSON.parse(children as string) as Array<PinType>
	values

	const [selectedPins, setSelectedPins] = useState([pins[0].pins[0]])
	const [signature, setSignature] = useState<string | undefined>()
	signature
	setSignature

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
			if (pin.type === "then")
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

	return (
		<div
			className={cn(
				"flex w-full cursor-move flex-col items-stretch text-white"
			)}
			role={preview ? "PlugPreview" : "Plug"}
			{...rest}
		>
			{selectedPins.map((pin, index) => (
				<div key={index}>
					<Pin
						selectedPin={pin}
						pins={availablePins}
						gridSize={gridSize}
						onPinChange={newPin => handleChange(index, newPin)}
					/>

					<PinAppendage
						pin={pin}
						onClick={() => handleAddition(index)}
						gridSize={gridSize}
						isAvailable={availablePins.length > 0}
					/>
				</div>
			))}

			<PlugSimulation pins={selectedPins} />
		</div>
	)
}

export default Plug
