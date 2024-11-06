import { useSession } from "next-auth/react"
import { FC, useMemo } from "react"

import { CircleHelp, Hash } from "lucide-react"

import { Button, Checkbox, Frame, Image, Search, TokenImage } from "@/components"
import { usePlugs, Value } from "@/contexts"
import { Action, cn, formatInputName, formatTitle, getIndexes } from "@/lib"
import { useActions, useColumns } from "@/state"

export const DynamicFragment: FC<{
	item: string
	index: number
	actionIndex: number
	fragmentIndex: number
	dynamicIndex: number
	action: Action
	fragment: string
	dynamic: string[]
	preview: boolean
}> = ({ index, item, action, actionIndex, dynamicIndex, fragment, dynamic, preview }) => {
	const { data: session } = useSession()
	const { isFrame, frame } = useColumns(index, `${actionIndex}-${dynamicIndex}`)
	const { plug, actions, handle } = usePlugs(item)

	const [solverActions] = useActions()

	const protocol = solverActions[action.protocol]
	const staticAction = protocol.schema[action.action]

	const own = plug && session && session.address === plug.socketId

	const [childIndex, parentIndex] = useMemo(() => getIndexes(fragment), [fragment])

	const inputName = formatInputName(staticAction.fields[parentIndex]?.name)

	const label = useMemo(() => {
		const value = action.values[parentIndex]

		if (!value || value === "") return inputName

		return value instanceof Object ? value.label.toLowerCase() : value
	}, [action, parentIndex, inputName])

	const isReady = action && Boolean(action.values[parentIndex])

	const options = useMemo(() => {
		if (!action.values || !staticAction.fields[parentIndex].options) return

		if (childIndex === null) return staticAction.fields[parentIndex].options

		// TODO: (#586) This is not correct -- Need to implement and validate pointer function indexes.
		// Right now we do not have any integrations that use pointers so we will come back to this later.
		// const childValue = action.values[childIndex]
		//
		// if (childValue === undefined || childValue instanceof Object === false) return
		//
		// return undefined
		//
		// return staticAction.fields[parentIndex].options[childValue.value]
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
							if (actionIndex !== nestedActionIndex) return actionValue

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
									if (!depChildIndex || !thisChildIndex) return false
									return (
										depChildIndex === thisChildIndex &&
										(idx === parentIndex || action.values[idx] === undefined)
									)
								})

							return shouldReset ? undefined : actionValue
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
					<div className="relative h-10 min-w-10">
						<Image
							src={protocol.metadata.icon}
							alt={action.protocol}
							width={64}
							height={64}
							className="absolute left-1/2 top-1/2 h-16 w-16 -translate-x-1/2 rounded-sm blur-2xl filter"
						/>
						<Image
							src={protocol.metadata.icon}
							alt={action.protocol}
							width={64}
							height={64}
							className="relative left-1/2 top-1/2 h-8 w-8 -translate-x-1/2 -translate-y-1/2 rounded-sm"
						/>
					</div>
				}
				label={
					<span className="relative">
						<span className="text-lg">
							<span className={cn(action.values.length > 1 && "opacity-40")}>
								{formatTitle(action.action)}
								{action.values.length > 1 && <span>:</span>}
							</span>
							{action.values.length > 1 && <span> {formatTitle(inputName)}</span>}
						</span>
					</span>
				}
				visible={isFrame}
				handleBack={dynamicIndex > 0 ? () => frame(`${actionIndex}-${dynamicIndex - 1}`) : undefined}
				hasOverlay
			>
				<div className="flex flex-col gap-4">
					{options === undefined && action.values[parentIndex] instanceof Object === false && (
						<Search
							icon={<Hash size={14} />}
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
										handleChange={() =>
											handleValue(
												// @ts-ignore - Action value can be string | Option but really only Option
												option.value === action.values[parentIndex]?.value ? null : option
											)
										}
									/>

									<button
										key={`${index}-${actionIndex}-${optionIndex}`}
										className="group flex w-full items-center text-left font-bold"
										onClick={() => handleValue(option)}
									>
										<div className="flex flex-row items-center gap-4">
											{option.icon && (
												<>
													{option.icon.startsWith(
														"https://token-icons.llamao.fi/icons/tokens/"
													) ? (
														<TokenImage
															logo={option.icon}
															symbol={option.label}
															size="xs"
															blur={false}
														/>
													) : (
														<Image
															src={option.icon}
															alt=""
															width={60}
															height={60}
															className="h-6 w-6 rounded-full"
															unoptimized
														/>
													)}
												</>
											)}
											{option.name}
										</div>
									</button>
								</div>
							))}
						</div>
					)}

					<Button
						variant={
							action.values[parentIndex] === "" ||
							(options !== undefined && action.values[parentIndex] === null)
								? "disabled"
								: "primary"
						}
						className="py-4"
						onClick={() => {
							frame(
								dynamicIndex + 1 < action.values.length
									? `${actionIndex}-${dynamicIndex + 1}`
									: undefined
							)
						}}
						disabled={
							action.values[parentIndex] === "" ||
							(options !== undefined && action.values[parentIndex] === null)
						}
					>
						{action.values[parentIndex] === "" ||
						(options !== undefined && action.values[parentIndex] === null)
							? "Missing required inputs"
							: action.values.length - 1 > dynamicIndex
								? "Next"
								: "Done"}
					</Button>
				</div>
			</Frame>
		</>
	)
}
