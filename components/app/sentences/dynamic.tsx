import { FC, useEffect, useMemo, useState } from "react"

import Image from "next/image"

import { Action, Option, StaticAction, Value } from "."
import { ChevronRight, InfinityIcon } from "lucide-react"

import { Button } from "@/components/buttons"
import { Search } from "@/components/inputs"
import { actionCategories, actions } from "@/lib/constants"
import { formatInputName, formatTitle } from "@/lib/functions"

import { Frame } from "../frames/base"
import { getIndexes } from "./fragments"

export const DynamicFragment: FC<{
	action: Action
	fragment: string
	values: Array<Value>
	handleValue: (value: Value) => void
}> = ({ action, fragment, values, handleValue }) => {
	const category =
		actionCategories[action.categoryName as keyof typeof actionCategories]
	const staticCategory =
		actions[action.categoryName as keyof typeof actionCategories]
	const staticAction: StaticAction =
		staticCategory[action.actionName as keyof typeof staticCategory]

	const Icon = staticAction.icon || InfinityIcon

	const [childIndex, index] = useMemo(() => getIndexes(fragment), [fragment])

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

	const inputName = formatInputName(staticAction.inputs[index].name)
	const label =
		values[index] === undefined || values[index] === ""
			? inputName
			: values[index] instanceof Object
				? formatTitle(values[index]?.label ?? "").toLowerCase()
				: values[index]

	const [valuesVisible, setValuesVisible] = useState(false)

	return (
		<>
			<button
				className="cursor-pointer rounded-lg bg-gradient-to-tr px-2 py-1 font-bold transition-all duration-200 ease-in-out"
				style={{
					background: `linear-gradient(to right, rgba(0,239,54,0.1), rgba(147,223,0,0.1))`,
					color: `#00E100`
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
					{options === undefined &&
						values[index] instanceof Object === false && (
							<Search
								icon={<Icon size={14} className="opacity-60" />}
								placeholder={formatTitle(inputName)}
								search={values[index]}
								handleSearch={(value: string) =>
									handleValue(value)
								}
							/>
						)}

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
