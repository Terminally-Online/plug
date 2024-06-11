import { plug } from "@/server/api/routers/plug"
import { socket } from "@/server/api/routers/socket"
import { createTRPCRouter } from "@/server/api/trpc"

export const appRouter = createTRPCRouter({
	socket,
	plug
})

export type AppRouter = typeof appRouter
