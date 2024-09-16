import { anonymousProtectedProcedure, createTRPCRouter } from "@/server/api/trpc"
import { TRPCError } from "@trpc/server"

import { z } from "zod"

import { SOCKET_BASE_QUERY } from "@/lib"

import { balances } from "./balances"

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
			// const columnsToCreate = ctx.session.user.demo
			// 	? DEFAULT_DEMO_VIEWS
			// 	: ctx.session.user.anonymous
			// 		? DEFAULT_ANONYMOUS_VIEWS
			// 		: DEFAULT_VIEWS

			await ctx.db.userSocket.upsert({
				where: {
					id: ctx.session.address
				},
				create: {
					id: ctx.session.address,
					socketAddress: TEMPORARY_ADDRESS,
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
				}
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
	balances
})
