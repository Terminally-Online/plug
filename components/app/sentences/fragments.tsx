import { FC, useEffect, useMemo, useState } from "react"

import { Action, Value } from "."

import { useActions } from "@/contexts/ActionProvider"

import { DynamicFragment } from "./dynamic"
import { StaticFragment } from "./static"

type Props = {
	action: Action
}

const regex = /({\d+(?:=>\d+)?})/g

export const getIndexes = (fragment: string) => {
	const sanitized = fragment.replace("{", "").replace("}", "").split("=>")

	if (sanitized.length > 1)
		return [sanitized[0], sanitized[1]].map(Number) as [number, number]

	return [null, Number(sanitized[0])] as [null, number]
}

// NOTE: This entire component is the collection of fragments that make up a sentence.
export const Fragments: FC<Props> = ({ action }) => {
	const { handleEdit } = useActions()

	const parsed = action.data ? JSON.parse(action.data) : undefined

	// Split all of the sentence fragments into an appropriate array based on the
	// regex shape that enables the f-string like syntax.
	const fragments = useMemo(
		() =>
			!parsed || !parsed["sentence"]
				? []
				: (parsed["sentence"].split(regex) as string[]),
		[parsed]
	)

	// Filter down to only the dynamic fragments that match the regex pattern
	// so that we can use the carried index value to update the correct indexes
	// when one is updated by the user.
	const dynamicFragments = useMemo(
		() => fragments.filter(fragment => fragment.match(regex)),
		[fragments]
	)

	// TODO: Values are assumed to be undefined right now rather than stored in the database.
	const [values, setValues] = useState<Array<Value>>(
		Array(dynamicFragments.length).fill(undefined)
	)

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
	const handleValue = (fragment: string, value: Value) => {
		const [, index] = getIndexes(fragment)

		setValues(prev =>
			prev.map((v, i) =>
				index === getIndexes(dynamicFragments[i])[0]
					? undefined
					: i === index
						? value
						: v
			)
		)
	}

	// useEffect(() => {
	// 	// handleEdit(values)
	// }, [values, handleEdit])

	return (
		<>
			{fragments.map((fragment, index) =>
				fragment.match(regex) ? (
					<DynamicFragment
						key={index}
						action={action}
						fragment={fragment}
						values={values}
						handleValue={value => handleValue(fragment, value)}
					/>
				) : (
					<StaticFragment key={index} fragment={fragment} />
				)
			)}
		</>
	)
}
