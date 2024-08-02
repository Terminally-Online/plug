import { inferRouterInputs, inferRouterOutputs } from "@trpc/server"

import { misc } from "@/server/api/routers/misc"
import { plug } from "@/server/api/routers/plug"
import { socket } from "@/server/api/routers/socket"
import { createTRPCRouter } from "@/server/api/trpc"

export const appRouter = createTRPCRouter({
	misc,
	plug,
	socket
})

export type AppRouter = typeof appRouter

export type RouterInputs = inferRouterInputs<AppRouter>
export type RouterOutputs = inferRouterOutputs<AppRouter>
