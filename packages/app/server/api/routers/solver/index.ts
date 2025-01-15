import { createTRPCRouter } from "../../trpc"
import { actions } from "./actions"
import { tokens } from "./tokens"

export const solver = createTRPCRouter({
	actions,
	tokens
})
