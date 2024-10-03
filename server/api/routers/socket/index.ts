import { createPublicClient, http } from "viem"
import { mainnet } from "viem/chains"
import { normalize } from "viem/ens"

import { anonymousProtectedProcedure, createTRPCRouter } from "@/server/api/trpc"
import { TRPCError } from "@trpc/server"

import { z } from "zod"

import { SOCKET_BASE_QUERY } from "@/lib"

import { balances } from "./balances"

const TEMPORARY_ADDRESS = "0x62180042606624f02d8a130da8a3171e9b33894d"

const client = createPublicClient({
	chain: mainnet,
	transport: http(process.env.ALCHEMY_API_URL)
})

export const socket = createTRPCRouter({
	get: anonymousProtectedProcedure.query(async ({ input, ctx }) => {
		// const columnsToCreate = ctx.session.user.demo
		// 	? DEFAULT_DEMO_VIEWS
		// 	: ctx.session.user.anonymous
		// 		? DEFAULT_ANONYMOUS_VIEWS
		// 		: DEFAULT_VIEWS
		// const { data: name } = useEnsName({
		// 	chainId: mainnet.id,
		// 	address: ctx.session?.address as `0x${string}`
		// })
		// const { data: avatar } = useEnsAvatar({
		// 	chainId: mainnet.id,
		// 	name: normalize(name ?? "") || undefined
		// })

		const name = ""
		const avatar = ""
		// TODO(#423): Removed ENS avatar and name and never re-implemented it.

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
							ens: name
								? {
										connectOrCreate: {
											where: { name },
											create: {
												name,
												avatar
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

		if (name)
			await ctx.db.eNS.upsert({
				where: { name },
				create: {
					name,
					avatar: avatar ? avatar : undefined,
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
					avatar: avatar ? avatar : undefined,
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
