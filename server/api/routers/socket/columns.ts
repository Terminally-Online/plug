import { z } from "zod"

import { TRPCError } from "@trpc/server"

import { createTRPCRouter, protectedProcedure } from "../../trpc"

export const COLUMN_KEYS = {
	ADD: "ADD",
	PLUGS: "PLUGS",
	DISCOVER: "DISCOVER",
	MY_PLUGS: "MY_PLUGS",
	ACTIVITY: "ACTIVITY",
	ASSETS: "ASSETS",
	TOKENS: "TOKENS",
	COLLECTIBLES: "COLLECTIBLES",
	POSITIONS: "POSITIONS",
	EARNINGS: "EARNINGS",
	SETTINGS: "SETTINGS"
}

export const DEFAULT_COLUMNS = [
	{ key: COLUMN_KEYS.PLUGS, index: 0 },
	{ key: COLUMN_KEYS.ACTIVITY, index: 1 }
]

export const columns = createTRPCRouter({
	add: protectedProcedure
		.input(z.object({ key: z.string(), id: z.string().optional() }))
		.mutation(async ({ input, ctx }) => {
			return await ctx.db.$transaction(async tx => {
				if (Object.values(COLUMN_KEYS).includes(input.key) === false)
					throw new TRPCError({
						code: "BAD_REQUEST",
						message: "Invalid column key."
					})

				const columns = await tx.consoleColumn.findMany({
					where: { socketId: ctx.session.address }
				})

				if (columns === null) throw new TRPCError({ code: "NOT_FOUND" })

				if (input.id === undefined)
					return await tx.userSocket.update({
						where: { id: ctx.session.address },
						data: {
							columns: {
								createMany: {
									data: [
										{
											key: input.key,
											index: columns.length
										}
									]
								}
							}
						},
						include: { columns: { orderBy: { index: "asc" } } }
					})

				return await tx.userSocket.update({
					where: { id: ctx.session.address },
					data: {
						columns: {
							update: {
								where: { id: input.id },
								data: {
									key: input.key
								}
							}
						}
					},
					include: { columns: { orderBy: { index: "asc" } } }
				})
			})
		}),
	remove: protectedProcedure
		.input(z.string())
		.mutation(async ({ input, ctx }) => {
			return await ctx.db.$transaction(async tx => {
				const column = await tx.consoleColumn.delete({
					where: { id: input }
				})

				return await tx.userSocket.update({
					where: { id: ctx.session.address },
					data: {
						columns: {
							updateMany: {
								where: { index: { gt: column.index } },
								data: { index: { decrement: 1 } }
							}
						}
					},
					include: { columns: { orderBy: { index: "asc" } } }
				})
			})
		}),
	resize: protectedProcedure
		.input(
			z.object({
				id: z.string(),
				width: z.number()
			})
		)
		.mutation(async ({ input, ctx }) => {
			return await ctx.db.userSocket.update({
				where: { id: ctx.session.address },
				data: {
					columns: {
						update: {
							where: { id: input.id },
							data: { width: input.width }
						}
					}
				},
				include: { columns: { orderBy: { index: "asc" } } }
			})
		}),
	move: protectedProcedure
		.input(
			z.object({
				from: z.number(),
				to: z.number()
			})
		)
		.mutation(async ({ input, ctx }) => {
			return await ctx.db.$transaction(async tx => {
				const columns = await tx.consoleColumn.findMany({
					where: { socketId: ctx.session.address },
					orderBy: { index: "asc" }
				})

				const [movedColumn] = columns.splice(input.from, 1)
				columns.splice(input.to, 0, movedColumn)

				return await tx.userSocket.update({
					where: { id: ctx.session.address },
					data: {
						columns: {
							deleteMany: { index: { gte: 0 } },
							createMany: {
								data: columns.map((column, index) => ({
									...column,
									consoleId: undefined,
									index,
									socketId: undefined
								}))
							}
						}
					},
					include: { columns: { orderBy: { index: "asc" } } }
				})
			})
		})
})
