import { Schedule, Transfer } from "@/lib"

export type Column = {
	id: number
	key: string
	index: number
	width?: number
	item?: string
	from?: string
	frame?: string
	schedule?: Schedule
	transfer?: Transfer
}
