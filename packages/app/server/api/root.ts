import { inferRouterInputs, inferRouterOutputs } from "@trpc/server"

import { biblo } from "@/server/api/routers/biblo"
import { jobs } from "@/server/api/routers/jobs"
import { misc } from "@/server/api/routers/misc"
import { plugs } from "@/server/api/routers/plugs"
import { socket } from "@/server/api/routers/socket"
import { solver } from "@/server/api/routers/solver"
import { createTRPCRouter } from "@/server/api/trpc"

export const appRouter = createTRPCRouter({
	biblo,
	plugs,
	socket,
	jobs,
	misc,
	solver
})

export type AppRouter = typeof appRouter

export type RouterInputs = inferRouterInputs<AppRouter>
export type RouterOutputs = inferRouterOutputs<AppRouter>
