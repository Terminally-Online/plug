// import accountRouter from "@/server/api/routers/account"
import { createTRPCRouter } from "@/server/api/trpc"

export const appRouter = createTRPCRouter({
	// account: accountRouter
})

export type AppRouter = typeof appRouter
