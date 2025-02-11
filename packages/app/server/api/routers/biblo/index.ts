
import { createTRPCRouter } from "../../trpc"
import { chat } from "./chat"

export const biblo = createTRPCRouter({
	chat,
})
