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

const whereWithSearch = (
	where: Prisma.CanvasWhereInput,
	fieldName: string,
	search: string | string[] | undefined
): Record<"where", Prisma.CanvasWhereInput> => {
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

const events = { add: "add-canvas", update: "update-canvas" }

export default createTRPCRouter({
	all: protectedProcedure
		.input(z.union([z.string(), z.array(z.string())]).optional())
		.query(async ({ ctx, input: search }) => {
			const userId = ctx.session.user.name

			try {
				return await ctx.db.canvas.findMany({
					...whereWithSearch({ userId }, "name", search),
					orderBy: {
						updatedAt: "desc"
					}
				})
			} catch (e) {
				throw new TRPCError({ code: "BAD_REQUEST" })
			}
		}),

	infinite: protectedProcedure
		.input(
			z.object({
				cursor: z.string().nullish(),
				limit: z.number().optional().default(10),
				sort: z.union([z.string(), z.array(z.string())]).optional(),
				search: z.union([z.string(), z.array(z.string())]).optional()
			})
		)
		.query(async ({ ctx, input }) => {
			try {
				const userId = ctx.session.user.name

				const { cursor, search, sort } = input

				const searchArray: (string | undefined)[] = Array.isArray(
					search
				)
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

				if (Array.isArray(sort))
					throw new TRPCError({ code: "NOT_IMPLEMENTED" })

				let orderBy = {}
				if (sort === "newest") orderBy = { createdAt: "desc" }
				else if (sort === "oldest") orderBy = { createdAt: "asc" }
				else if (sort === "active") orderBy = { updatedAt: "desc" }
				else orderBy = { updatedAt: "desc" }

				const count = await ctx.db.canvas.count({ where })

				const limit = input.limit + 1

				const canvases = await ctx.db.canvas.findMany({
					...whereWithSearch({ userId }, "name", search),
					orderBy,
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
			} catch (e) {
				throw new TRPCError({ code: "BAD_REQUEST" })
			}
		}),

	get: protectedProcedure.input(z.string()).query(async ({ ctx, input }) => {
		try {
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
		} catch (e) {
			throw new TRPCError({ code: "BAD_REQUEST" })
		}
	}),

	add: protectedProcedure
		.input(
			z.object({
				name: z.string().optional().default("Untitled Canvas"),
				public: z.boolean().optional().default(true),
				color: z
					.string()
					.optional()
					.default(
						`#${Math.floor(Math.random() * 16777215).toString(16)}`
					)
			})
		)
		.mutation(async ({ ctx, input }) => {
			const userId = ctx.session.user.name

			try {
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

				emitter.emit(events.add, canvas)

				return canvas
			} catch (e) {
				throw new TRPCError({ code: "BAD_REQUEST" })
			}
		}),

	update: protectedProcedure
		.input(
			z.object({
				id: z.string(),
				name: z.string().optional(),
				color: z.string().optional(),
				public: z.boolean().optional(),
				components: z.array(ComponentSchema).optional()
			})
		)
		.mutation(async ({ ctx, input }) => {
			const userId = ctx.session.user.name

			try {
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
				emitter.emit(events.update, updatedCanvas)

				return updatedCanvas
			} catch (e) {
				throw new TRPCError({ code: "BAD_REQUEST" })
			}
		}),

	// ? Allow a user to listen to the creation of their own canvases.
	onAdd: protectedProcedure.subscription(({ ctx }) => {
		return observable<CanvasWithComponents>(emit => {
			const handleCreate = (canvas: CanvasWithComponents) => {
				const canView = canvas.userId === ctx.session.user.name

				if (!canView) return

				emit.next(canvas)
			}

			emitter.on(events.add, handleCreate)

			return () => {
				emitter.off(events.add, handleCreate)
			}
		})
	}),

	// ? Allow a user to listen to the updates of their own canvases.
	// * This would include name changes, color changes, etc.
	//	  -- The components are not streamed here.
	onUpdate: protectedProcedure.subscription(({ ctx }) => {
		return observable<CanvasWithComponents>(emit => {
			const handleUpdate = (canvas: CanvasWithComponents) => {
				const canView =
					canvas.userId === ctx.session.user.name || canvas.public

				if (!canView) return

				emit.next(canvas)
			}

			emitter.on(events.update, handleUpdate)

			return () => {
				emitter.off(events.update, handleUpdate)
			}
		})
	}),

	// ? Extend the component router as base canvas route.
	component: componentRouter
})
