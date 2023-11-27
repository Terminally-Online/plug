import { EventEmitter } from 'events'
import { z } from 'zod'

import { Prisma } from '@prisma/client'
import { observable } from '@trpc/server/observable'

import { createTRPCRouter, protectedProcedure } from '@/server/api/trpc'

const ee = new EventEmitter()

export const ComponentSchema = z.object({
	id: z.string(),
	top: z.number(),
	left: z.number(),
	type: z.union([z.literal('PLUG'), z.literal('BOX'), z.literal('MARKDOWN')]),
	width: z.number(),
	height: z.number(),
	content: z.string(),
	createdAt: z.string().optional(),
	updatedAt: z.string().optional()
})

export const component = Prisma.validator<Prisma.ComponentDefaultArgs>()({})

export type Component = Prisma.CanvasGetPayload<typeof component>

export default createTRPCRouter({
	add: protectedProcedure
		.input(
			z.object({
				id: z.string(),
				component: ComponentSchema.omit({ id: true })
			})
		)
		.mutation(async ({ ctx, input }) => {
			const component = await ctx.db.component.create({
				data: {
					...input.component,
					canvasId: input.id
				}
			})

			ee.emit(input.id, component)

			return component
		}),

	move: protectedProcedure
		.input(
			z.object({
				id: z.string(),
				component: z.object({
					id: z.string(),
					top: z.number(),
					left: z.number()
				})
			})
		)
		.mutation(async ({ ctx, input }) => {
			const component = await ctx.db.component.update({
				where: {
					id: input.component.id,
					canvasId: input.id
				},
				data: {
					top: input.component.top,
					left: input.component.left
				}
			})

			ee.emit('move', input.component.id)

			return component
		}),

	randomNumber: protectedProcedure.subscription(() => {
		return observable<number>(emit => {
			const interval = setInterval(() => {
				emit.next(Math.random())
			}, 1000)

			return () => {
				clearInterval(interval)
			}
		})
	}),

	onMove: protectedProcedure.subscription(() => {
		return observable<string>(emit => {
			console.log('ready to move')

			const onMove = (data: string) => {
				console.log('move', data)
				emit.next(data)
			}

			ee.on('move', onMove)

			return () => {
				ee.off('move', onMove)
			}
		})
	})
})
