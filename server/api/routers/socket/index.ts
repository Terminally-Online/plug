import { createPublicClient, http } from "viem"
import { mainnet } from "viem/chains"
import { normalize } from "viem/ens"

import { anonymousProtectedProcedure, createTRPCRouter } from "@/server/api/trpc"
import { TRPCError } from "@trpc/server"

import { z } from "zod"

import { SOCKET_BASE_QUERY } from "@/lib"

import { balances } from "./balances"

const TEMPORARY_ADDRESS = "0x62180042606624f02d8a130da8a3171e9b33894d"
const ENS_CACHE_TIME = 24 * 60 * 60 * 1000 // 24 hours

const client = createPublicClient({
	chain: mainnet,
	transport: http(process.env.ALCHEMY_API_URL)
})

export const socket = createTRPCRouter({
	get: anonymousProtectedProcedure.query(async ({ ctx }) => {
		let ens = await ctx.db.eNS.findFirst({
			where: { identity: { socketId: ctx.session.address } },
			orderBy: { updatedAt: "desc" }
		})

		if (!ens) {
			ens = await ctx.db.eNS.create({
				data: {
					name: "",
					avatar: "",
					identity: { connect: { socketId: ctx.session.address } }
				}
			})
		}

		let name = ens.name
		let avatar = ens.avatar

		if (Date.now() - ens.updatedAt.getTime() > ENS_CACHE_TIME) {
			try {
				const newName = await client.getEnsName({ address: ctx.session.address as `0x${string}` })
				if (newName) {
					name = newName
					avatar = (await client.getEnsAvatar({ name: normalize(name) })) || ""
				}
			} catch (error) {
				console.error(`Error fetching ENS data for ${ctx.session.address}:`, error)
			} finally {
				// NOTE: We silently swallow errors here because we don't want to interrupt the
				// user's session if the ENS name is not available. There may be a better
				// way to handle this in the future. For now, it is fine since it works and
				// logs about failed ENS lookups don't really matter.
				await ctx.db.eNS.update({
					where: { name: ens.name },
					data: { name, avatar, updatedAt: new Date() }
				})
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
							connect: { name: ens.name }
						}
					}
				}
			},
			update: {
				updatedAt: new Date(), // Add this line
				identity: {
					upsert: {
						create: {
							ens: {
								connect: { name: ens.name }
							}
						},
						update: {
							ens: {
								connect: { name: ens.name }
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
