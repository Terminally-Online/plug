import { createTRPCRouter } from "../../trpc"
import { actions } from "./actions"
import { killer } from "./kill"
import { tokens } from "./tokens"

export const solver = createTRPCRouter({
	actions,
	killer,
	tokens
})
