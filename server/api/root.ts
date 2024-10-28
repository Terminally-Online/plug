import { inferRouterInputs, inferRouterOutputs } from "@trpc/server"

import { misc } from "@/server/api/routers/misc"
import { plugs } from "@/server/api/routers/plugs"
import { socket } from "@/server/api/routers/socket"
import { createTRPCRouter } from "@/server/api/trpc"

import { jobs } from "./routers/jobs"

export const appRouter = createTRPCRouter({
	plugs,
	socket,
	jobs,
	misc
})

export type AppRouter = typeof appRouter

export type RouterInputs = inferRouterInputs<AppRouter>
export type RouterOutputs = inferRouterOutputs<AppRouter>
