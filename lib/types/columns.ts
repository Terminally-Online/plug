import { MinimalUserSocketModel } from "@/prisma/types"

export type Column = {
	key: string
	index: number
	width?: number
	item?: string
	from?: string
	viewAs?: MinimalUserSocketModel
	frame?: string
}
