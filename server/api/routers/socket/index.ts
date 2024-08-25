import { z } from "zod"

import { Prisma } from "@prisma/client"
import { TRPCError } from "@trpc/server"

import { DEFAULT_VIEWS, getFarcasterFollowing, SOCKET_BASE_INCLUDE, SOCKET_BASE_QUERY, VIEW_KEYS } from "@/lib"
import { createTRPCRouter, protectedProcedure } from "@/server/api/trpc"

import { balances } from "./balances"
import { columns } from "./columns"
import { graph } from "./graph"

const TEMPORARY_ADDRESS = "0x62180042606624f02d8a130da8a3171e9b33894d"

export const socket = createTRPCRouter({
	get: protectedProcedure
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
							data: DEFAULT_VIEWS
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

			// NOTE: This is not awaited because we will just do this as a background task.
			await getFarcasterFollowing(ctx.session.address)

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
	balances,
	columns,
	graph
})
