import { Schedule, Transfer } from "@/lib"

export type Column = {
	id: number
	key: string
	index: number
	chain?: number
	width?: number
	item?: string
	from?: string
	frame?: string
	schedule?: Schedule
	transfer?: Transfer
}
