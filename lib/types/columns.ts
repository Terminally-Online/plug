import { MinimalUserSocketModel } from "@/prisma/types"

import { Schedule } from "@/lib"

export type Column = {
	key: string
	index: number
	width?: number
	item?: string
	from?: string
	viewAs?: MinimalUserSocketModel
	frame?: string
	schedule?: Schedule
}
