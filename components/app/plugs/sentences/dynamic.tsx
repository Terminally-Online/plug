import { useSession } from "next-auth/react"
import { FC, useMemo } from "react"

import { ChevronRight, CircleHelp } from "lucide-react"

import { Button, Checkbox, Frame, Image, Search } from "@/components"
import { Option, usePlugs, Value } from "@/contexts"
import { Action, categories, cn, formatInputName, formatTitle, getIndexes, actions as staticActions } from "@/lib"
import { useColumns } from "@/state"

export const DynamicFragment: FC<{
	item: string
	index: number
	actionIndex: number
	fragmentIndex: number
	action: Action
	fragment: string
	dynamic: string[]
	preview: boolean
}> = ({ index, item, actionIndex, fragmentIndex, action, fragment, dynamic, preview }) => {
	const { data: session } = useSession()
	const { isFrame, frame } = useColumns(index, `${actionIndex}-${fragmentIndex}`)
	const { plug, actions, handle } = usePlugs(item)

	const category = categories[action.categoryName]
	const staticAction = staticActions[action.categoryName][action.actionName]

	const Icon = staticAction.icon || CircleHelp

	const own = plug && session && session.address === plug.socketId

	const [childIndex, parentIndex] = useMemo(() => getIndexes(fragment), [fragment])

	const inputName = formatInputName(staticAction.inputs[parentIndex]?.name)

	const label = useMemo(() => {
		const value = action.values[parentIndex]

		if (!value || value === "") return inputName

		return value instanceof Object ? formatTitle(value.label).toLowerCase() : value
	}, [action, parentIndex, inputName])

	const isReady = action && Boolean(action.values[parentIndex])

	const options: Array<Option> | undefined = useMemo(() => {
		if (!action.values || !staticAction.options) return undefined

		if (childIndex === null) return (staticAction.options as Array<Array<Option>>)[parentIndex]

		const childValue = action.values[childIndex]

		if (childValue === undefined || childValue instanceof Object === false) return undefined

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
			value instanceof Object && action.values[parentIndex] instanceof Object
				? value.value !== action.values[parentIndex].value
				: value !== action.values[parentIndex]

		if (hasChanged) {
			handle.action.edit({
				id: plug?.id,
				actions: JSON.stringify(
					actions.map((action, nestedActionIndex) => ({
						...action,
						values: action.values.map((actionValue, valueIndex) => {
							if (actionIndex === nestedActionIndex) {
								// If this is the value being changed
								if (valueIndex === parentIndex) return value

								// Check for dependencies using dynamic fragments
								const fragment = dynamic[valueIndex]
								if (!fragment) return actionValue

								// Get dependency info for this value
								const [thisChildIndex] = getIndexes(fragment)

								// Reset this value if:
								// 1. It directly depends on the changed value (parentIndex matches childIndex)
								// 2. It depends on any value that we're resetting (recursive dependency)
								const shouldReset =
									// Direct dependency on the changed value
									parentIndex === thisChildIndex ||
									// Indirect dependency through another value that's being reset
									action.values.some((_, idx) => {
										if (idx === valueIndex) return false
										const [depChildIndex] = getIndexes(dynamic[idx])
										return (
											depChildIndex === thisChildIndex &&
											(idx === parentIndex || action.values[idx] === undefined)
										)
									})

								return shouldReset ? undefined : actionValue
							}

							return actionValue
						})
					}))
				)
			})
		}
	}

	return (
		<>
			<button
				className={cn(
					"rounded-sm bg-gradient-to-tr px-2 py-1 font-bold transition-all duration-200 ease-in-out",
					preview && isReady === false ? "text-plug-red" : "text-plug-green",
					own === true ? "cursor-pointer" : "cursor-default"
				)}
				style={{
					background:
						preview && isReady === false
							? "linear-gradient(to right, rgba(255,0,0,0.1), rgba(255,0,0,0.1))"
							: `linear-gradient(to right, rgba(0,239,54,0.1), rgba(147,223,0,0.1))`
				}}
				onClick={() => (own ? frame() : undefined)}
			>
				{label}
			</button>

			<Frame
				index={index}
				icon={
					<Image
						src={category.image}
						alt={action.categoryName}
						width={24}
						height={24}
						className="rounded-sm"
					/>
				}
				label={`${formatTitle(action.actionName)}${action.values.length > 1 ? `: ${formatTitle(inputName)}` : ""}`}
				visible={isFrame}
			>
				<div className="flex flex-col gap-4">
					{options === undefined && action.values[parentIndex] instanceof Object === false && (
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
								<div
									key={`${index}-${actionIndex}-${optionIndex}`}
									className="flex flex-row items-center gap-4"
								>
									<Checkbox
										checked={
											// @ts-ignore - Action value can be string | Option but really only Option
											option.value === action.values[parentIndex]?.value
										}
										handleChange={() => handleValue(option)}
									/>

									<button
										key={`${index}-${actionIndex}-${optionIndex}`}
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
									</button>
								</div>
							))}
						</div>
					)}

					<Button
						className="py-4"
						onClick={() => {
							if (options === undefined) frame()
						}}
					>
						{options === undefined ? "Done" : "Continue"}
					</Button>
				</div>
			</Frame>
		</>
	)
}
