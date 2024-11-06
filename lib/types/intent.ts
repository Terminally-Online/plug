import { DateRange } from "react-day-picker"

import { frequencies } from "@/lib"

export type Option = {
	icon: JSX.Element | undefined
	label: string
	value: string | number
	imagePath?: string
}

export type Value = string | Option | undefined | null

export type Action = {
	protocol: string
	action: string
	values: Array<Value>
}

export type Actions = Array<Action>

export type Schedule = {
	date: DateRange | undefined
	repeats: (typeof frequencies)[0]
}
