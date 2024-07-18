import { FC, useMemo } from "react"

import Image from "next/image"

import { useSession } from "next-auth/react"

import { ChevronRight, CircleHelp } from "lucide-react"

import { Button, Frame, Search } from "@/components"
import { Option, useFrame, usePlugs, Value } from "@/contexts"
import {
	categories,
	cn,
	formatInputName,
	formatTitle,
	getIndexes,
	actions as staticActions
} from "@/lib"

type Props = {
	index: number
	fragmentIndex: number
}

export const DynamicFragment: FC<Props> = ({ index, fragmentIndex }) => {
	const { data: session } = useSession()
	const { frameVisible, handleFrameVisible } = useFrame()
	const { id, plug, actions, fragments, dynamic, handle } = usePlugs()

	const action = actions[index]
	const fragment = fragments[index][fragmentIndex]

	const category = categories[action.categoryName]
	const staticAction = staticActions[action.categoryName][action.actionName]

	const Icon = staticAction.icon || CircleHelp

	const own = plug && session && session.address === plug.userAddress

	const [childIndex, parentIndex] = useMemo(
		() => getIndexes(fragment),
		[fragment]
	)

	const inputName = formatInputName(staticAction.inputs[parentIndex].name)

	const label = useMemo(() => {
		const value = action.values[parentIndex]

		if (!value || value === "") return inputName

		return value instanceof Object
			? formatTitle(value.label).toLowerCase()
			: value
	}, [action, parentIndex, inputName])

	const options: Array<Option> | undefined = useMemo(() => {
		if (!action.values || !staticAction.options) return undefined

		if (childIndex === null)
			return (staticAction.options as Array<Array<Option>>)[parentIndex]

		const childValue = action.values[childIndex]

		if (childValue === undefined || childValue instanceof Object === false)
			return undefined

		return (
			staticAction.options as Array<{
				[key: string]: Array<Option>
			}>
		)[parentIndex][childValue.value]
	}, [staticAction, childIndex, parentIndex, action])

	// This loops through the fragments and updates the respective value based
	// on the stringified index value of the fragment. This admittedly is a bit
	// confusing to read on first glance, but it's a way to update child values
	// when a parent value changes.
	//
	// Example:
	// {0}    will update the value at index 0.
	// {0=>1} will update the value at index 1 and set the value to undefined
	//        when the value at index 0 changes which is signalled by the
	//        upper index value.
	const handleValue = (value: Value) => {
		const hasChanged =
			value instanceof Object &&
			actions[index].values[parentIndex] instanceof Object
				? // @ts-ignore -- Don't know why this is necessary or how to resolve it, but it is.
					value.value !== actions[index].values[parentIndex].value
				: value !== actions[index].values[parentIndex]

		// If the value is the same as the current value, then we don't need
		// to update the database with a non-changing value.
		if (hasChanged === true)
			handle.action.edit({
				id,
				actions: JSON.stringify(
					actions.map((action, actionIndex) => ({
						...action,
						values: action.values.map((actionValue, valueIndex) => {
							if (actionIndex === index) {
								// If it is the action and fragment we are editing, return
								// the new value and save it to the database.
								if (valueIndex === parentIndex) return value

								// If it is the action, but another fragment depends on the one that is changing,
								// then we need to set the value to undefined if the value is actually changing.
								// If the value is not changing, then we can just return the original value.
								const [childIndex] = getIndexes(
									dynamic[index][valueIndex]
								)

								if (parentIndex === childIndex) return undefined
							}

							// If the action is not the one we are editing, return the
							// original value that is already in the array.
							return actionValue
						})
					}))
				)
			})

		// If the value selected was from a list of options, close the frame.
		if (value instanceof Object) handleFrameVisible(undefined)
	}

	return (
		<>
			<button
				className={cn(
					"rounded-lg bg-gradient-to-tr px-2 py-1 font-bold text-plug-green transition-all duration-200 ease-in-out",
					own === true ? "cursor-pointer" : "cursor-default"
				)}
				style={{
					background: `linear-gradient(to right, rgba(0,239,54,0.1), rgba(147,223,0,0.1))`
				}}
				onClick={() =>
					own
						? handleFrameVisible(`${index}-${fragmentIndex}`)
						: undefined
				}
			>
				{label}
			</button>

			<Frame
				className="scrollbar-hide z-[2] max-h-[calc(100vh-80px)] overflow-y-auto"
				icon={
					<Image
						src={category.image}
						alt={action.categoryName}
						width={24}
						height={24}
						className="rounded-md"
					/>
				}
				label={`${formatTitle(action.actionName)}${action.values.length > 1 ? `: ${formatTitle(inputName)}` : ""}`}
				visible={frameVisible === `${index}-${fragmentIndex}`}
			>
				<div className="flex flex-col gap-4">
					{options === undefined &&
						action.values[parentIndex] instanceof Object ===
							false && (
							<Search
								icon={<Icon size={14} />}
								placeholder={formatTitle(inputName)}
								// @ts-ignore
								search={action.values[parentIndex]}
								handleSearch={handleValue}
							/>
						)}

					{options !== undefined && (
						<div className="flex w-full flex-col gap-2">
							{options.map((option, optionIndex) => (
								<button
									key={`${index}-${optionIndex}`}
									className="group flex w-full items-center text-left font-bold"
									onClick={() => handleValue(option)}
								>
									<div className="flex flex-row items-center gap-4">
										{option.imagePath && (
											<Image
												src={option.imagePath}
												alt=""
												width={64}
												height={64}
												className="w-6"
											/>
										)}
										{formatTitle(option.label)}
									</div>

									<Button
										variant="secondary"
										className="ml-auto p-1 group-hover:bg-grayscale-100"
										onClick={() => handleValue(option)}
									>
										<ChevronRight
											size={14}
											className="float-right"
										/>
									</Button>
								</button>
							))}
						</div>
					)}
				</div>
			</Frame>
		</>
	)
}
