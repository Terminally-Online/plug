import { createTRPCRouter } from "../../trpc"
import { actions } from "./actions"
import { killer } from "./kill"

export const solver = createTRPCRouter({
	actions,
	killer,
})
