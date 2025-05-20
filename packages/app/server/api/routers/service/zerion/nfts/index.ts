import { createTRPCRouter } from "@/server/api/trpc"

import { detail } from "./detail"
import { list } from "./list"

export const nfts = createTRPCRouter({
	list,
	detail
})
