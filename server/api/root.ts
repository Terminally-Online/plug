import { inferRouterInputs, inferRouterOutputs } from "@trpc/server"

import { jobs } from "@/server/api/routers/jobs"
import { misc } from "@/server/api/routers/misc"
import { plugs } from "@/server/api/routers/plugs"
import { socket } from "@/server/api/routers/socket"
import { createTRPCRouter } from "@/server/api/trpc"

export const appRouter = createTRPCRouter({
	plugs,
	socket,
	jobs,
	misc
})

export type AppRouter = typeof appRouter

export type RouterInputs = inferRouterInputs<AppRouter>
export type RouterOutputs = inferRouterOutputs<AppRouter>
