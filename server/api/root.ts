import { EventEmitter } from 'stream'

import { observable } from '@trpc/server/observable'

import canvasRouter from '@/server/api/routers/canvas'
import {
	createTRPCRouter,
	protectedProcedure,
	publicProcedure
} from '@/server/api/trpc'

import { emitter } from '../emitter'

let counter = 0

export const appRouter = createTRPCRouter({
	canvas: canvasRouter,

	public: publicProcedure.query(() => {
		return 'you can see this message!'
	}),
	protected: protectedProcedure.query(() => {
		return 'This is protected content. You can access this content because you are signed in.'
	}),
	randomNumber: publicProcedure.subscription(() => {
		return observable<number>(emit => {
			const int = setInterval(() => {
				emit.next(Math.random())
			}, 2000)

			return () => {
				clearInterval(int)
			}
		})
	}),

	increment: protectedProcedure.mutation(() => {
		counter++
		emitter.emit('increment', counter)
		return counter
	}),

	onIncremenet: protectedProcedure.subscription(() => {
		return observable<number>(emit => {
			const listener = (count: number) => {
				emit.next(count)
			}

			emitter.on('increment', listener)

			return () => {
				emitter.off('increment', listener)
			}
		})
	})
})

export type AppRouter = typeof appRouter
