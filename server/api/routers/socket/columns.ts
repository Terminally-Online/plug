import { z } from "zod"

import { TRPCError } from "@trpc/server"

import { SOCKET_BASE_QUERY, VIEW_KEYS } from "@/lib"

import { anonymousProtectedProcedure, createTRPCRouter } from "../../trpc"

export const columns = createTRPCRouter({
	add: anonymousProtectedProcedure
		.input(
			z.object({
				key: z.string(),
				id: z.string().optional(),
				index: z.number().optional(),
				item: z.string().optional()
			})
		)
		.mutation(async ({ input, ctx }) => {
			console.log("adding column")

			return await ctx.db.$transaction(async tx => {
				if (Object.values(VIEW_KEYS).includes(input.key) === false)
					throw new TRPCError({
						code: "BAD_REQUEST",
						message: "Invalid column key."
					})

				if (input.id && input.index)
					throw new TRPCError({
						code: "BAD_REQUEST",
						message: "Cannot set both id and index."
					})

				const columns = await tx.consoleColumn.findMany({
					where: { socketId: ctx.session.address }
				})

				if (input.index !== undefined)
					return await tx.userSocket.update({
						where: { id: ctx.session.address },
						data: {
							columns: {
								updateMany: {
									where: {
										index: {
											gte: input.index
										}
									},
									data: {
										index: {
											increment: 1
										}
									}
								},
								create: {
									key: input.key,
									index: input.index,
									item: input.item
								}
							}
						},
						...SOCKET_BASE_QUERY
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
						...SOCKET_BASE_QUERY
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
					...SOCKET_BASE_QUERY
				})
			})
		}),
	navigate: anonymousProtectedProcedure
		.input(
			z.object({
				id: z.string().optional(),
				key: z.string(),
				item: z.string().optional(),
				from: z.string().optional()
			})
		)
		.mutation(async ({ input, ctx }) => {
			if (input.id === undefined) throw new TRPCError({ code: "BAD_REQUEST" })

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
				},
				...SOCKET_BASE_QUERY
			})
		}),
	remove: anonymousProtectedProcedure.input(z.string()).mutation(async ({ input, ctx }) => {
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
				...SOCKET_BASE_QUERY
			})
		})
	}),
	resize: anonymousProtectedProcedure
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
				...SOCKET_BASE_QUERY
			})
		}),
	move: anonymousProtectedProcedure
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
					...SOCKET_BASE_QUERY
				})
			})
		}),
	as: anonymousProtectedProcedure
		.input(
			z.object({
				id: z.string(),
				as: z.string()
			})
		)
		.mutation(async ({ input, ctx }) => {
			return await ctx.db.userSocket.update({
				where: { id: ctx.session.address },
				data: {
					columns: {
						update: {
							where: { id: input.id },
							data: { viewAsId: input.as === ctx.session.address ? null : input.as }
						}
					}
				},
				...SOCKET_BASE_QUERY
			})
		})
})
