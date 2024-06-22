// import { z } from "zod"
//
// import { Prisma } from "@prisma/client"
// import { TRPCError } from "@trpc/server"
// import { observable } from "@trpc/server/observable"
//
// import componentRouter, {
// 	ComponentSchema
// } from "@/server/api/routers/component"
// import { createTRPCRouter, protectedProcedure } from "@/server/api/trpc"
// import { emitter } from "@/server/emitter"

// export const canvasWithComponents =
// 	Prisma.validator<Prisma.CanvasDefaultArgs>()({
// 		include: { components: true }
// 	})
//
// export type CanvasWithComponents = Prisma.CanvasGetPayload<
// 	typeof canvasWithComponents
// >
//
// const whereWithSearch = (
// 	where: Prisma.CanvasWhereInput,
// 	fieldName: string,
// 	search: string | string[] | undefined
// ): Record<"where", Prisma.CanvasWhereInput> => {
// 	const searchArray: (string | undefined)[] = Array.isArray(search)
// 		? search
// 		: search
// 			? search.split(" ")
// 			: []
//
// 	const searchSyntax = searchArray.join(" | ")
//
// 	const searchClause =
// 		searchSyntax && searchSyntax.length > 0
// 			? {
// 					[`${fieldName}`]: {
// 						search: searchSyntax
// 					}
// 				}
// 			: {}
//
// 	return {
// 		where: {
// 			...where,
// 			...searchClause
// 		}
// 	}
// }
//
// const events = { add: "add-canvas", update: "update-canvas" }

// export default createTRPCRouter({
// 	infinite: protectedProcedure
// 		.input(
// 			z.object({
// 				cursor: z.string().nullish(),
// 				limit: z.number().optional().default(10),
// 				sort: z.union([z.string(), z.array(z.string())]).optional(),
// 				search: z.union([z.string(), z.array(z.string())]).optional()
// 			})
// 		)
// 		.query(async ({ ctx, input }) => {
// 			try {
// 				const userId = ctx.session.user.name
//
// 				const { cursor, search, sort } = input
//
// 				const searchArray: (string | undefined)[] = Array.isArray(
// 					search
// 				)
// 					? search
// 					: search
// 						? search.split(" ")
// 						: []
//
// 				const syntaxSearch = searchArray.join(" | ")
//
// 				let where
// 				if (search !== undefined && search !== "")
// 					where = {
// 						name: { search: syntaxSearch },
// 						userId
// 					}
// 				else where = { userId }
//
// 				if (Array.isArray(sort))
// 					throw new TRPCError({ code: "NOT_IMPLEMENTED" })
//
// 				let orderBy = {}
// 				if (sort === "newest") orderBy = { createdAt: "desc" }
// 				else if (sort === "oldest") orderBy = { createdAt: "asc" }
// 				else if (sort === "active") orderBy = { updatedAt: "desc" }
// 				else orderBy = { updatedAt: "desc" }
//
// 				const count = await ctx.db.canvas.count({ where })
//
// 				const limit = input.limit + 1
//
// 				const canvases = await ctx.db.canvas.findMany({
// 					...whereWithSearch({ userId }, "name", search),
// 					orderBy,
// 					cursor: cursor
// 						? {
// 								id: cursor
// 							}
// 						: undefined,
// 					take: limit
// 				})
//
// 				let nextCursor: typeof cursor | undefined = undefined
// 				if (canvases.length > 10) {
// 					const nextItem = canvases.pop()
// 					nextCursor = nextItem!.id
// 				}
//
// 				return {
// 					items: canvases,
// 					nextCursor,
// 					count
// 				}
// 			} catch (e) {
// 				throw new TRPCError({ code: "BAD_REQUEST" })
// 			}
// 		})
// })
