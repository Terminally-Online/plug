import { FC, useMemo, useState } from "react"

import Image from "next/image"

import { ChevronRight, CircleHelp } from "lucide-react"

import { Frame } from "@/components/app/frames/base"
import { Action, Option, Value } from "@/components/app/sentences"
import { getIndexes } from "@/components/app/sentences/fragments"
import { Button } from "@/components/buttons"
import { Search } from "@/components/inputs"
import { actionCategories, actions } from "@/lib/constants"
import { formatInputName, formatTitle } from "@/lib/functions"

export const DynamicFragment: FC<{
	action: Action
	fragment: string
	values: Array<Value>
	handleValue: (value: Value) => void
}> = ({ action, fragment, values, handleValue }) => {
	const category = actionCategories[action.categoryName]
	const staticAction = actions[action.categoryName][action.actionName]

	const Icon = staticAction.icon || CircleHelp

	const [valuesVisible, setValuesVisible] = useState(false)

	const [childIndex, index] = useMemo(() => getIndexes(fragment), [fragment])

	const inputName = formatInputName(staticAction.inputs[index].name)

	const label = useMemo(() => {
		const value = values[index]

		if (value === undefined || value === "") return inputName

		return value instanceof Object
			? formatTitle(value.label).toLowerCase()
			: value
	}, [values, index, inputName])

	const options: Array<Option> | undefined = useMemo(() => {
		if (!values || !staticAction.options) return undefined

		if (childIndex === null)
			return (staticAction.options as Array<Array<Option>>)[index]

		const childValue = values[childIndex]

		if (childValue === undefined || childValue instanceof Object === false)
			return undefined

		return (
			staticAction.options as Array<{
				[key: string]: Array<Option>
			}>
		)[index][childValue.value]
	}, [staticAction, childIndex, index, values])

	return (
		<>
			<button
				className="cursor-pointer rounded-lg bg-gradient-to-tr px-2 py-1 font-bold transition-all duration-200 ease-in-out"
				style={{
					background: `linear-gradient(to right, rgba(0,239,54,0.1), rgba(147,223,0,0.1))`,
					color: `#00EF35`
				}}
				onClick={() => setValuesVisible(!valuesVisible)}
			>
				{label}
			</button>

			<Frame
				className="z-[2]"
				icon={
					<Image
						src={category.image}
						alt={action.categoryName}
						width={24}
						height={24}
						className="rounded-md"
					/>
				}
				label={`${formatTitle(action.actionName)}${values.length >= 1 ? `: ${formatTitle(inputName)}` : ""}`}
				visible={valuesVisible}
				handleVisibleToggle={() => setValuesVisible(!valuesVisible)}
			>
				<div className="flex flex-col gap-4">
					{/* TODO: Fix this. */}
					{/* {options === undefined && values[index] instanceof String &&
						values[index] instanceof Object === false && (
							<Search
								icon={<Icon size={14} className="opacity-60" />}
								placeholder={formatTitle(inputName)}
								search={values[index]}
								handleSearch={(value: string) =>
									handleValue(value)
								}
							/>
						)} */}

					{options !== undefined && (
						<div className="flex w-full flex-col gap-2">
							{options.map((option, optionIndex) => (
								<button
									key={optionIndex}
									className="group flex w-full text-left font-bold"
									onClick={() => {
										handleValue(option)
										setValuesVisible(false)
									}}
								>
									{formatTitle(option.label)}
									<Button
										variant="secondary"
										className="ml-auto p-1 group-hover:bg-grayscale-100"
										onClick={() => {
											handleValue(option)
											setValuesVisible(false)
										}}
									>
										<ChevronRight
											size={14}
											className="float-right opacity-60"
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
