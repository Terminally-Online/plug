import { createTRPCRouter } from "../../trpc"
import { actions } from "./actions"

export const solver = createTRPCRouter({
	actions
})
