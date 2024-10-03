import { createPublicClient, http } from "viem"
import { mainnet } from "viem/chains"
import { normalize } from "viem/ens"

import { anonymousProtectedProcedure, createTRPCRouter } from "@/server/api/trpc"
import { TRPCError } from "@trpc/server"

import { z } from "zod"

import { SOCKET_BASE_QUERY } from "@/lib"

import { balances } from "./balances"

const TEMPORARY_ADDRESS = "0x62180042606624f02d8a130da8a3171e9b33894d"
const ENS_CACHE_TIME = 30 * 24 * 60 * 60 * 1000

const client = createPublicClient({
	chain: mainnet,
	transport: http(process.env.ALCHEMY_API_URL)
})

export const socket = createTRPCRouter({
	get: anonymousProtectedProcedure.query(async ({ ctx }) => {
		const ens = await ctx.db.eNS.findFirst({
			where: { identity: { socketId: ctx.session.address } },
			orderBy: { updatedAt: "desc" }
		})

		let name = ens?.name || ""
		let avatar = ens?.avatar || ""

		if (ens === null || (ens && ens.updatedAt > new Date(Date.now() - ENS_CACHE_TIME))) {
			try {
				name = (await client.getEnsName({ address: ctx.session.address as `0x${string}` })) || ""
				if (name) {
					avatar = (await client.getEnsAvatar({ name: normalize(name) })) || ""
				}
			} finally {
				// NOTE: We silently swallow errors here because we don't want to interrupt the
				//       user's session if the ENS name is not available. There may be a better
				//       way to handle this in the future. For now, it is fine since it works and
				//       logs about failed ENS lookups don't really matter.
			}
		}

		await ctx.db.userSocket.upsert({
			where: { id: ctx.session.address },
			create: {
				id: ctx.session.address,
				socketAddress: TEMPORARY_ADDRESS,
				identity: {
					connectOrCreate: {
						where: { socketId: ctx.session.address },
						create: {
							ens: {
								connectOrCreate: {
									where: { name },
									create: {
										name,
										avatar
									}
								}
							}
						}
					}
				}
			},
			update: {
				identity: {
					upsert: {
						create: {
							ens: {
								create: {
									name,
									avatar
								}
							}
						},
						update: {
							ens: {
								upsert: {
									create: {
										name,
										avatar
									},
									update: {
										name,
										avatar
									}
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
