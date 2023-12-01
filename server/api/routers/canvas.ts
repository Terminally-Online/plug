import { z } from 'zod'

import { Prisma } from '@prisma/client'
import { TRPCError } from '@trpc/server'
import { observable } from '@trpc/server/observable'

import componentRouter, {
	ComponentSchema
} from '@/server/api/routers/component'
import { createTRPCRouter, protectedProcedure } from '@/server/api/trpc'
import { emitter } from '@/server/emitter'

export const canvasWithComponents =
	Prisma.validator<Prisma.CanvasDefaultArgs>()({
		include: { components: true }
	})

export type CanvasWithComponents = Prisma.CanvasGetPayload<
	typeof canvasWithComponents
>

export const CanvasSchema = z.object({
	id: z.string(),
	name: z.string(),
	public: z.boolean(),
	color: z.string(),
	components: z.array(ComponentSchema),
	createdAt: z.string().optional(),
	updatedAt: z.string().optional()
})

export default createTRPCRouter({
	all: protectedProcedure
		.input(z.union([z.string(), z.array(z.string())]).optional())
		.query(async ({ ctx, input: search }) => {
			const userId = ctx.session.user.name

			// ? I am not sure what to do about arrays. Prisma syntax is enough for now.
			if (Array.isArray(search))
				throw new TRPCError({ code: 'BAD_REQUEST' })

			try {
				if (search !== undefined && search !== '')
					return await ctx.db.canvas.findMany({
						where: {
							name: { search },
							userId
						},
						orderBy: {
							updatedAt: 'desc'
						}
					})

				return await ctx.db.canvas.findMany({
					where: {
						userId
					},
					orderBy: {
						updatedAt: 'desc'
					}
				})
			} catch (e) {
				return []
			}
		}),
	get: protectedProcedure.input(z.string()).query(async ({ ctx, input }) => {
		const canvas = await ctx.db.canvas.findUnique({
			where: {
				id: input
			},
			include: { components: true }
		})

		if (!canvas) throw new TRPCError({ code: 'NOT_FOUND' })

		if (canvas.public) return canvas

		const userId = ctx.session.user.name

		if (!userId) throw new TRPCError({ code: 'UNAUTHORIZED' })

		if (canvas.userId !== userId)
			throw new TRPCError({ code: 'UNAUTHORIZED' })

		return canvas as CanvasWithComponents
	}),
	create: protectedProcedure
		.input(
			z.object({
				name: z.string(),
				public: z.boolean(),
				color: z.string()
			})
		)
		.mutation(async ({ ctx, input }) => {
			const userId = ctx.session.user.name

			console.log(userId, ctx.session)

			// * Create the canvas in the database.
			const canvas = await ctx.db.canvas.create({
				data: {
					name: input.name,
					public: input.public,
					color: input.color,
					user: {
						connectOrCreate: {
							where: { id: userId },
							create: { id: userId }
						}
					}
				},
				include: { components: true }
			})

			emitter.emit('create-canvas', canvas)

			return canvas
		}),
	update: protectedProcedure
		.input(
			z.object({
				id: z.string(),
				name: z.string().optional(),
				color: z.string().optional(),
				public: z.boolean().optional(),
				components: z.array(ComponentSchema)
			})
		)
		.mutation(async ({ ctx, input }) => {
			const userId = ctx.session.user.name

			const canvas = await ctx.db.canvas.findUnique({
				where: {
					id: input.id
				},
				include: { components: true }
			})

			if (!canvas) throw new TRPCError({ code: 'NOT_FOUND' })

			if (canvas.userId !== userId)
				throw new TRPCError({ code: 'FORBIDDEN' })

			// * Update the fields that were passed in.
			const updatedCanvas: CanvasWithComponents =
				await ctx.db.canvas.update({
					where: {
						id: input.id
					},
					data: {
						...canvas,
						...input,
						// * TODO: For now we are not updating components
						components: undefined
					},
					include: { components: true }
				})

			// * Emit an update event.
			emitter.emit('update', updatedCanvas)

			return updatedCanvas
		}),
	onCreate: protectedProcedure.subscription(() => {
		return observable<CanvasWithComponents>(emit => {
			emitter.on('create-canvas', emit.next)

			return () => {
				emitter.off('create-canvas', emit.next)
			}
		})
	}),
	onUpdate: protectedProcedure.subscription(() => {
		return observable<CanvasWithComponents>(emit => {
			emitter.on('update', emit.next)

			return () => {
				emitter.off('update', emit.next)
			}
		})
	}),
	component: componentRouter
})
