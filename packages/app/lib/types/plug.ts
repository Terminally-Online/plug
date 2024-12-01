import { DateRange } from "react-day-picker"

import { frequencies } from "@/lib"

export type ActionSchema = {
	metadata: {
		icon: string
	}
	schema: {
		[action: string]: {
			sentence: string
			options?: Record<
				string,
				{
					value: string
					name: string
					label: string
					info?: string
					icon: string
				}[]
			>
		}
	}
}

export type ActionSchemas = {
	[protocol: string]: ActionSchema
}

export type Action = {
	protocol: string
	action: string
	values: Record<string, string | undefined>
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