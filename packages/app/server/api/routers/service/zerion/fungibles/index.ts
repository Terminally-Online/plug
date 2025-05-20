import { createTRPCRouter } from "@/server/api/trpc"

import { chart } from "./chart"
import { detail } from "./detail"
import { list } from "./list"

export const fungibles = createTRPCRouter({
	list,
	detail,
	chart
})
