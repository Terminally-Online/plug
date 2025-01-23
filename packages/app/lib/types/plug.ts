import { DateRange } from "react-day-picker"

import { frequencies } from "@/lib"
import { connectedChains } from "@/contexts"

export type Options = {
	value: string
	name: string
	label: string
	info?: string
	icon: string
}[]

export type ActionSchema = {
	metadata: {
		icon: string
		tags: Array<string>
		chains: Array<typeof connectedChains[number]['id']>
	}
	schema: Record<
		string,
		{
			type: string
			sentence: string
			options?: Record<string, Options | Record<string, Options>>
		}
	>
}

export type ActionSchemas = {
	[protocol: string]: ActionSchema
}

export type Action = {
	protocol: string
	action: string
	values: Record<string, { value: string; name: string } | undefined>
}

export type Actions = Array<Action>

export type Schedule = {
	date: DateRange | undefined
	repeats: (typeof frequencies)[0]
}

export type Transfer = {
	recipient?: string
	percentage?: number
	precise?: string
	tokenId?: string
}
