import { mainnet } from "viem/chains"
import { normalize } from "viem/ens"

import { TRPCError } from "@trpc/server"

import { z } from "zod"

import { createClient, SOCKET_BASE_QUERY } from "@/lib"
import { anonymousProtectedProcedure, createTRPCRouter } from "@/server/api/trpc"

import { balances } from "./balances"
import { companion } from "./companion"
import { referral } from "./referral"

export const TEMPORARY_ADDRESS = "0x62180042606624f02d8a130da8a3171e9b33894d"
const ENS_CACHE_TIME = 24 * 60 * 60 * 1000

const client = createClient(mainnet.id)

export const socket = createTRPCRouter({
	get: anonymousProtectedProcedure.query(async ({ ctx }) => {
		let ens = await ctx.db.eNS.findUnique({
			where: { socketId: ctx.session.address }
		})

		let name = ens?.name ?? ""
		let avatar = ens?.avatar ?? ""

		if (
			(ens === null || Date.now() - ens.updatedAt.getTime() > ENS_CACHE_TIME) &&
			ctx.session.address.startsWith("0x")
		) {
			try {
				name = (await client.getEnsName({ address: ctx.session.address as `0x${string}` })) || ""
				if (name) {
					avatar = (await client.getEnsAvatar({ name: normalize(name) })) || ""
				}
			} catch (error) {
				// NOTE: We silently swallow errors here because we don't want to interrupt the
				// user's session if the ENS name is not available. There may be a better
				// way to handle this in the future. For now, it is fine since it works and
				// logs about failed ENS lookups don't really matter.
			}
		}

		await ctx.db.userSocket.upsert({
			where: { id: ctx.session.address },
			create: {
				id: ctx.session.address,
				socketAddress: TEMPORARY_ADDRESS,
				identity: {
					create: {
						ens: {
							create: {
								name,
								avatar
							}
						},
						companion: {
							create: {
								name: "New Companion",
								feedCount: 0,
								treatsFed: 0
							}
						}
					}
				}
			},
			update: {
				updatedAt: new Date(),
				identity: {
					upsert: {
						create: {
							ens: {
								connectOrCreate: {
									where: { socketId: ctx.session.address },
									create: {
										name,
										avatar
									}
								}
							},
							companion: {
								connectOrCreate: {
									where: { socketId: ctx.session.address },
									create: {
										name: "New Companion",
										feedCount: 0,
										treatsFed: 0
									}
								}
							}
						},
						update: {
							ens: {
								update: {
									where: { socketId: ctx.session.address },
									data: { name, avatar, updatedAt: new Date() }
								}
							},
							companion: {
								upsert: {
									where: { socketId: ctx.session.address },
									create: {
										name: "New Companion",
										feedCount: 0,
										treatsFed: 0
									},
									update: {}
								}
							}
						}
					}
				}
			}
		})

		const socket = await ctx.db.userSocket.findFirst({
			where: { id: ctx.session.address },
			...SOCKET_BASE_QUERY
		})

		if (!socket) {
			throw new TRPCError({
				code: "NOT_FOUND",
				message: "Socket not found"
			})
		}

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
										ens: {
											name: {
												contains: input.search,
												mode: "insensitive"
											}
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
	companion,
	referral
})
