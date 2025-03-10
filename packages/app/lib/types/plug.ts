import { DateRange } from "react-day-picker"

import { frequencies } from "@/lib"

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
