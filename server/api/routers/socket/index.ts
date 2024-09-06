import { z } from "zod"

import { TRPCError } from "@trpc/server"

import { DEFAULT_ANONYMOUS_VIEWS, DEFAULT_VIEWS, SOCKET_BASE_QUERY, VIEW_KEYS } from "@/lib"
import { anonymousProtectedProcedure, createTRPCRouter } from "@/server/api/trpc"

import { balances } from "./balances"
import { columns } from "./columns"

const TEMPORARY_ADDRESS = "0x62180042606624f02d8a130da8a3171e9b33894d"

export const socket = createTRPCRouter({
	get: anonymousProtectedProcedure
		.input(
			z.object({
				name: z.string().nullish(),
				avatar: z.string().nullish()
			})
		)
		.query(async ({ input, ctx }) => {
			const { columns } = await ctx.db.userSocket.upsert({
				where: {
					id: ctx.session.address
				},
				create: {
					id: ctx.session.address,
					socketAddress: TEMPORARY_ADDRESS,
					columns: {
						createMany: {
							data: ctx.session.user.anonymous ? DEFAULT_ANONYMOUS_VIEWS : DEFAULT_VIEWS
						}
					},
					identity: {
						connectOrCreate: {
							where: { socketId: ctx.session.address },
							create: {
								ens: input.name
									? {
											connectOrCreate: {
												where: { name: input.name },
												create: {
													name: input.name,
													avatar: input.avatar
												}
											}
										}
									: undefined
							}
						}
					}
				},
				update: {
					identity: {
						upsert: {
							where: { socketId: ctx.session.address },
							create: {},
							update: {}
						}
					}
				},
				select: { columns: true }
			})

			if (input.name)
				await ctx.db.eNS.upsert({
					where: { name: input.name },
					create: {
						name: input.name,
						avatar: input.avatar ? input.avatar : undefined,
						identity: {
							connectOrCreate: {
								where: { socketId: ctx.session.address },
								create: {
									socketId: ctx.session.address
								}
							}
						}
					},
					update: {
						avatar: input.avatar ? input.avatar : undefined,
						identity: {
							connectOrCreate: {
								where: { socketId: ctx.session.address },
								create: {
									socketId: ctx.session.address
								}
							}
						}
					}
				})

			// NOTE: Make sure the socket always has the global home column that is
			//       denoted by the -1 index as the column view should never show it.
			if (columns.find(column => column.index === -1) === undefined) {
				await ctx.db.consoleColumn.create({
					data: {
						key: VIEW_KEYS.HOME,
						index: -1,
						socketId: ctx.session.address
					}
				})
			}

			const socket = await ctx.db.userSocket.findFirst({
				where: { id: ctx.session.address },
				...SOCKET_BASE_QUERY
			})

			if (socket === null) throw new TRPCError({ code: "NOT_FOUND" })

			return socket
		}),
	search: anonymousProtectedProcedure
		.input(z.object({ search: z.string(), limit: z.number().optional().default(3) }))
		.query(async ({ input, ctx }) => {
			return await ctx.db.userSocket.findMany({
				where: {
					AND: [
						{
							OR: [
								{
									id: {
										contains: input.search,
										mode: "insensitive"
									}
								},
								{
									socketAddress: {
										contains: input.search,
										mode: "insensitive"
									}
								},
								{
									identity: {
										ensName: {
											contains: input.search,
											mode: "insensitive"
										}
									}
								}
							]
						},
						{
							NOT: {
								id: {
									contains: "anonymous",
									mode: "insensitive"
								}
							}
						},
						{
							NOT: {
								id: {
									contains: "demo",
									mode: "insensitive"
								}
							}
						},
						{
							NOT: {
								id: ctx.session.address
							}
						}
					]
				},
				orderBy: { updatedAt: "desc" },
				take: input.limit,
				...SOCKET_BASE_QUERY
			})
		}),
	balances,
	columns
})
