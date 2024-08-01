import { z } from "zod"

import { TRPCError } from "@trpc/server"

import { createTRPCRouter, protectedProcedure } from "../trpc"

const keys = {
	ADD: "ADD",
	PLUGS: "PLUGS",
	ACTIVITY: "ACTIVITY",
	ASSETS: "ASSETS",
	TOKENS: "TOKENS",
	COLLECTIBLES: "COLLECTIBLES",
	POSITIONS: "POSITIONS",
	EARNINGS: "EARNINGS",
	SETTINGS: "SETTINGS"
}

export const console = createTRPCRouter({
	get: protectedProcedure.query(async ({ ctx }) => {
		return await ctx.db.console.upsert({
			where: { id: ctx.session.address },
			create: {
				id: ctx.session.address,
				columns: {
					createMany: {
						data: [
							{ key: keys.PLUGS, index: 0 },
							{ key: keys.ACTIVITY, index: 1 }
						]
					}
				}
			},
			update: {},
			include: { columns: true }
		})
	}),
	add: protectedProcedure
		.input(z.object({ key: z.string(), id: z.string().optional() }))
		.mutation(async ({ input, ctx }) => {
			return await ctx.db.$transaction(async tx => {
				if (Object.values(keys).includes(input.key) === false)
					throw new TRPCError({
						code: "BAD_REQUEST",
						message: "Invalid column key."
					})

				const columns = await tx.consoleColumn.findMany({
					where: { consoleId: ctx.session.address }
				})

				if (columns === null) throw new TRPCError({ code: "NOT_FOUND" })

				if (input.id === undefined)
					return await tx.console.update({
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

				return await tx.console.update({
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

				const console = await tx.console.update({
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

				if (console === null) throw new TRPCError({ code: "NOT_FOUND" })

				return console
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
			return await ctx.db.console.update({
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
					where: { consoleId: ctx.session.address },
					orderBy: { index: "asc" }
				})

				const [movedColumn] = columns.splice(input.from, 1)
				columns.splice(input.to, 0, movedColumn)

				return await tx.console.update({
					where: { id: ctx.session.address },
					data: {
						columns: {
							deleteMany: { index: { gte: 0 } },
							createMany: {
								data: columns.map((column, index) => ({
									...column,
									consoleId: undefined,
									index
								}))
							}
						}
					},
					include: { columns: { orderBy: { index: "asc" } } }
				})
			})
		})
})
