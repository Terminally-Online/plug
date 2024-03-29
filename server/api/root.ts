import accountRouter from "@/server/api/routers/account"
import canvasRouter from "@/server/api/routers/canvas"
import { createTRPCRouter } from "@/server/api/trpc"

export const appRouter = createTRPCRouter({
	account: accountRouter,
	canvas: canvasRouter
})

export type AppRouter = typeof appRouter
