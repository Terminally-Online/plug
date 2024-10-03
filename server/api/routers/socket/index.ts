import { createPublicClient, http } from "viem"
import { mainnet } from "viem/chains"

import { anonymousProtectedProcedure, createTRPCRouter } from "@/server/api/trpc"
import { TRPCError } from "@trpc/server"

import { z } from "zod"

import { SOCKET_BASE_QUERY } from "@/lib"

import { balances } from "./balances"

const TEMPORARY_ADDRESS = "0x62180042606624f02d8a130da8a3171e9b33894d"
const ENS_UPDATE_INTERVAL_DAYS = 30;

const client = createPublicClient({
	chain: mainnet,
	transport: http(process.env.ALCHEMY_API_URL)
})

function shouldUpdateENS(lastUpdated: Date | null): boolean {
	if (!lastUpdated) return true;
	const updateInterval = ENS_UPDATE_INTERVAL_DAYS * 24 * 60 * 60 * 1000;
	return Date.now() - lastUpdated.getTime() > updateInterval;
}

export const socket = createTRPCRouter({
	get: anonymousProtectedProcedure.query(async ({ input, ctx }) => {
		const existingSocket = await ctx.db.userSocket.findUnique({
			where: { id: ctx.session.address },
			include: { identity: { include: { ens: true } } },
		});

		const shouldUpdate = shouldUpdateENS(existingSocket?.identity?.ens?.updatedAt || null);

		let ensName = existingSocket?.identity?.ens?.name || null;
		let ensAvatar = existingSocket?.identity?.ens?.avatar || null;

		if (shouldUpdate) {
			try {
				ensName = await client.getEnsName({ address: ctx.session.address as `0x${string}` });
				if (ensName) {
					ensAvatar = await client.getEnsAvatar({ name: ensName });
				}
			} catch (error) {
				console.error(`Error fetching ENS data for ${ctx.session.address}:`, error);
			}
		}

		await ctx.db.userSocket.upsert({
			where: { id: ctx.session.address },
			create: {
				id: ctx.session.address,
				socketAddress: ctx.session.address,
				identity: {
					create: {
						ens: ensName ? {
							create: {
								name: ensName,
								avatar: ensAvatar || undefined
							}
						} : undefined
					}
				}
			},
			update: {
				identity: {
					upsert: {
						create: {
							ens: ensName ? {
								create: {
									name: ensName,
									avatar: ensAvatar || undefined
								}
							} : undefined
						},
						update: {
							ens: ensName ? {
								upsert: {
									create: {
										name: ensName,
										avatar: ensAvatar || undefined
									},
									update: {
										name: ensName,
										avatar: ensAvatar || undefined
									}
								}
							} : undefined
						}
					}
				}
			}
		});

		const socket = await ctx.db.userSocket.findFirst({
			where: { id: ctx.session.address },
			...SOCKET_BASE_QUERY
		});

		if (socket === null) throw new TRPCError({ code: "NOT_FOUND" });

		return socket;
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
