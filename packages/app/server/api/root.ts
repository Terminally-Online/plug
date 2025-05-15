import { inferRouterInputs, inferRouterOutputs } from "@trpc/server"

import { misc } from "@/server/api/routers/misc"
import { plugs } from "@/server/api/routers/plugs"
import { service } from "@/server/api/routers/service"
import { socket } from "@/server/api/routers/socket"
import { solver } from "@/server/api/routers/solver"
import { createTRPCRouter } from "@/server/api/trpc"

export const appRouter = createTRPCRouter({
	plugs,
	service,
	socket,
	misc,
	solver,
})

export type AppRouter = typeof appRouter

export type RouterInputs = inferRouterInputs<AppRouter>
export type RouterOutputs = inferRouterOutputs<AppRouter>
