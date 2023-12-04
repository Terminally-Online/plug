import { z } from "zod"

import { Prisma } from "@prisma/client"
import { TRPCError } from "@trpc/server"
import { observable } from "@trpc/server/observable"

import componentRouter, {
	ComponentSchema
} from "@/server/api/routers/component"
import { createTRPCRouter, protectedProcedure } from "@/server/api/trpc"
import { emitter } from "@/server/emitter"

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

const whereWithSearch = (
	where: Prisma.CanvasWhereInput,
	fieldName: string,
	search: string | string[] | undefined
) => {
	const searchArray: (string | undefined)[] = Array.isArray(search)
		? search
		: search
		  ? search.split(" ")
		  : []

	const searchSyntax = searchArray.join(" | ")

	const searchClause =
		searchSyntax && searchSyntax.length > 0
			? {
					[`${fieldName}`]: {
						search: searchSyntax
					}
			  }
			: {}

	return {
		where: {
			...where,
			...searchClause
		}
	}
}

export default createTRPCRouter({
	all: protectedProcedure
		.input(z.union([z.string(), z.array(z.string())]).optional())
		.query(async ({ ctx, input: search }) => {
			const userId = ctx.session.user.name

			return await ctx.db.canvas.findMany({
				...whereWithSearch({ userId }, "name", search),
				orderBy: {
					updatedAt: "desc"
				}
			})
		}),

	infinite: protectedProcedure
		.input(
			z.object({
				cursor: z.string().nullish(),
				limit: z.number().optional().default(10),
				search: z.union([z.string(), z.array(z.string())]).optional()
			})
		)
		.query(async ({ ctx, input }) => {
			const userId = ctx.session.user.name

			const { cursor, search } = input

			const searchArray: (string | undefined)[] = Array.isArray(search)
				? search
				: search
				  ? search.split(" ")
				  : []

			const syntaxSearch = searchArray.join(" | ")

			let where
			if (search !== undefined && search !== "")
				where = {
					name: { search: syntaxSearch },
					userId
				}
			else where = { userId }

			const count = await ctx.db.canvas.count({ where })

			const limit = input.limit + 1

			const canvases = await ctx.db.canvas.findMany({
				...whereWithSearch({ userId }, "name", search),
				orderBy: {
					updatedAt: "desc"
				},
				cursor: cursor
					? {
							id: cursor
					  }
					: undefined,
				take: limit
			})

			let nextCursor: typeof cursor | undefined = undefined
			if (canvases.length > 10) {
				const nextItem = canvases.pop()
				nextCursor = nextItem!.id
			}

			return {
				items: canvases,
				nextCursor,
				count
			}
		}),

	get: protectedProcedure.input(z.string()).query(async ({ ctx, input }) => {
		const canvas = await ctx.db.canvas.findUnique({
			where: {
				id: input
			},
			include: { components: true }
		})

		if (!canvas) throw new TRPCError({ code: "NOT_FOUND" })

		if (canvas.public) return canvas

		const userId = ctx.session.user.name

		if (!userId) throw new TRPCError({ code: "UNAUTHORIZED" })

		if (canvas.userId !== userId)
			throw new TRPCError({ code: "UNAUTHORIZED" })

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

			emitter.emit("create-canvas", canvas)

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

			if (!canvas) throw new TRPCError({ code: "NOT_FOUND" })

			if (canvas.userId !== userId)
				throw new TRPCError({ code: "FORBIDDEN" })

			// * Update the fields that were passed in.
			const updatedCanvas: CanvasWithComponents =
				await ctx.db.canvas.update({
					where: {
						id: input.id
					},
					data: {
						...canvas,
						...input,
						components: undefined
					},
					include: { components: true }
				})

			// * Emit an update event.
			emitter.emit("update", updatedCanvas)

			return updatedCanvas
		}),
	onCreate: protectedProcedure.subscription(() => {
		return observable<CanvasWithComponents>(emit => {
			emitter.on("create-canvas", emit.next)

			return () => {
				emitter.off("create-canvas", emit.next)
			}
		})
	}),
	onUpdate: protectedProcedure.subscription(() => {
		return observable<CanvasWithComponents>(emit => {
			emitter.on("update", emit.next)

			return () => {
				emitter.off("update", emit.next)
			}
		})
	}),
	component: componentRouter
})
