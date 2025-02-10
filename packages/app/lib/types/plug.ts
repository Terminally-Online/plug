import { DateRange } from "react-day-picker"

import { frequencies } from "@/lib"
import { connectedChains } from "@/contexts"
import { Option } from "@/state/plugs"

export type Options = {
	value: string
	name: string
	label: string
	icon: { [key: string]: string }
	info?: { label: string, value: string }
}[]

export type ActionSchema = {
	metadata: {
		icon: string
		tags: Array<string>
		chains: Array<{ 
			name: string, 
			chainIds: [typeof connectedChains[number]['id']], 
			explorer: string, 
			icon: { [key: string]: string } 
		}>
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
} & Partial<Options[number]>

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
