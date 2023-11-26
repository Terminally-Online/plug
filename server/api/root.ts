import canvasRouter from '@/server/api/routers/canvas'
import { createTRPCRouter } from '@/server/api/trpc'

export const appRouter = createTRPCRouter({
	canvas: canvasRouter
})

export type AppRouter = typeof appRouter
