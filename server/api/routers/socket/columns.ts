import { z } from "zod"

import { TRPCError } from "@trpc/server"

import { VIEW_KEYS } from "@/lib"

import { createTRPCRouter, protectedProcedure } from "../../trpc"

export const columns = createTRPCRouter({
	add: protectedProcedure
		.input(z.object({ key: z.string(), id: z.string().optional() }))
		.mutation(async ({ input, ctx }) => {
			return await ctx.db.$transaction(async tx => {
				if (Object.values(VIEW_KEYS).includes(input.key) === false)
					throw new TRPCError({
						code: "BAD_REQUEST",
						message: "Invalid column key."
					})

				const columns = await tx.consoleColumn.findMany({
					where: { socketId: ctx.session.address }
				})

				if (input.id === undefined)
					return await tx.userSocket.update({
						where: { id: ctx.session.address },
						data: {
							columns: {
								createMany: {
									data: [
										{
											key: input.key,
											index: columns.length - 1
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
	navigate: protectedProcedure
		.input(
			z.object({
				id: z.string().optional(),
				key: z.string(),
				item: z.string().optional(),
				from: z.string().optional()
			})
		)
		.mutation(async ({ input, ctx }) => {
			if (input.id === undefined)
				throw new TRPCError({ code: "BAD_REQUEST" })

			return await ctx.db.userSocket.update({
				where: { id: ctx.session.address },
				data: {
					columns: {
						update: {
							where: { id: input.id },
							data: {
								key: input.key,
								item: input.item ?? null,
								from: input.from ?? null
							}
						}
					}
				}
			})
		}),
	remove: protectedProcedure
		.input(z.string())
		.mutation(async ({ input, ctx }) => {
			return await ctx.db.$transaction(async tx => {
				const column = await tx.consoleColumn.delete({
					where: { id: input, socketId: ctx.session.address }
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
				const consoleColumns = await tx.consoleColumn.findMany({
					where: { socketId: ctx.session.address },
					orderBy: { index: "asc" }
				})

				const columns = consoleColumns.slice(1, consoleColumns.length)
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
