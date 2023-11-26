import { createTRPCRouter, publicProcedure } from '../trpc'
import canvasRouter from './canvas'

export const appRouter = createTRPCRouter({
	healthcheck: publicProcedure.query(() => 'healthy'),
	canvas: canvasRouter
})

export type AppRouter = typeof appRouter
