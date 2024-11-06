import { DateRange } from "react-day-picker"

import { actions, categories } from "../constants"
import { frequencies } from "../functions"

export type Option = {
	icon: JSX.Element | undefined
	label: string
	value: string | number
	imagePath?: string
}

export type Value = string | Option | undefined | null

export type Action = {
	protocol: keyof typeof categories
	action: keyof (typeof actions)[keyof typeof categories]
	values: Array<Value>
}

export type Actions = Array<Action>

export type Schedule = {
	date: DateRange | undefined
	repeats: (typeof frequencies)[0]
}
