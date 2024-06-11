import { FC, useEffect, useMemo, useState } from "react"

import { Action, Value } from "."

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

// NOTE: This entire component is a lil' bit confusing to read at first glance, and
//       requires some Typescript trickery to get the values to update correctly, but
//       it works well for the use case. If you are confused, just go through each piece
//       again. This component is one of the few actually commented/documented because
//       it is one of the few that cannot be easily understood without context.
export const Fragments: FC<Props> = ({ action }) => {
	// Split all of the sentence fragments into an appropriate array based on the
	// regex shape that enables the f-string like syntax.
	const fragments = useMemo(() => {
		if (action.data === undefined) return []

		const parsed = JSON.parse(action.data)

		if (!parsed["sentence"]) return []

		return parsed["sentence"].split(regex) as string[]
	}, [action])

	// Filter down to only the dynamic fragments that match the regex pattern
	// so that we can use the carried index value to update the correct indexes
	// when one is updated by the user.
	const dynamicFragments = useMemo(
		() => fragments.filter(fragment => fragment.match(regex)),
		[fragments]
	)

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
