import { mainnet } from "viem/chains"
import { normalize } from "viem/ens"

import { TRPCError } from "@trpc/server"

import { z } from "zod"

import { createClient, SOCKET_BASE_QUERY } from "@/lib"
import { anonymousProtectedProcedure, createTRPCRouter } from "@/server/api/trpc"

import { onboard } from "./onboard"
import { referral } from "./referral"
import { stats } from "./stats"
import { env } from "@/env"
import { getSocketAddress, getSocketFactory, getSocketImplementation, getSocketSalt } from "@/lib/functions/socket"

const ENS_CACHE_TIME = 24 * 60 * 60 * 1000
export const MAGIC_NONCE = BigInt(1738)

const client = createClient(mainnet.id)

const getDeployment = async (admin: `0x${string}`) => {
	const { deployment: { address: factory } } = getSocketFactory()
	const { deployment: { address: implementation } } = getSocketImplementation()
	// TODO MASON AND CHANCE: This doesn't seem to add up, it looks like the salt also includes the delegate address and the socket implementation address in the factory deploy method. Does the reference here and in seed.ts need to be updated or am I missing something
	const { hex: salt } = getSocketSalt(
		MAGIC_NONCE,
		admin as `0x${string}`,
	)
	const { address: socketAddress } = await getSocketAddress(salt as `0x${string}`)

	return {
		socketAddress,
		deploymentFactory: factory,
		deploymentNonce: parseInt(MAGIC_NONCE.toString()),
		deploymentDelegate: env.SOLVER_DELEGATE_ADDRESS,
		deploymentImplementation: implementation,
		deploymentSalt: salt
	} as const
}

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


		let socketAddress = ctx.session.address
		let factory = undefined
		let nonce = undefined
		let delegate = undefined
		let implementation = undefined
		let salt = undefined
		if (ctx.session.address.startsWith("0x")) {
			console.log("ctx.session.address", ctx.session.address)
			const { 
				socketAddress: deploymentSocketAddress,
				deploymentFactory,
				deploymentNonce,
				deploymentDelegate, 
				deploymentImplementation, 
				deploymentSalt  
			} = await getDeployment(ctx.session.address as `0x${string}`)

			socketAddress = deploymentSocketAddress
			factory = deploymentFactory
			nonce = deploymentNonce
			delegate = deploymentDelegate
			implementation = deploymentImplementation
			salt = deploymentSalt
		}

		await ctx.db.socket.upsert({
			where: { id: ctx.session.address },
			create: {
				id: ctx.session.address,
				socketAddress,
				deploymentFactory: factory,
				deploymentNonce: parseInt(MAGIC_NONCE.toString()),
				deploymentDelegate: delegate,
				deploymentImplementation: implementation,
				deploymentSalt: salt,
				identity: {
					create: {
						ens: {
							create: {
								name,
								avatar
							}
						},
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
						},
						update: {
							ens: {
								update: {
									where: { socketId: ctx.session.address },
									data: { name, avatar, updatedAt: new Date() }
								}
							},
						}
					}
				}
			}
		})

		const socket = await ctx.db.socket.findFirst({
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
			return await ctx.db.socket.findMany({
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

	referral,
	stats,
	onboard
})
