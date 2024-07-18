import { z } from "zod"

import { Prisma } from "@prisma/client"
import { TRPCError } from "@trpc/server"
import { observable } from "@trpc/server/observable"

import { getBalances } from "@/lib"
import { createTRPCRouter, protectedProcedure } from "@/server/api/trpc"

export const userSocket = Prisma.validator<Prisma.UserSocketDefaultArgs>()({})
export type UserSocket = Prisma.UserSocketGetPayload<typeof userSocket>

const events = {
	add: "add-socket",
	rename: "rename-socket"
}

const TEMPORARY_ADDRESS = "0x62180042606624f02d8a130da8a3171e9b33894d"

const CURRENT_IMPLEMENTATION_VERSION = 0
const FACTORY_ADDRESS = ""
const FACTORY_ABI = []

export const socket = createTRPCRouter({
	// TODO: Implement the real functionality for this.
	// address: protectedProcedure.query(async ({ ctx }) => {
	// 	const { address } = ctx.session

	// 	const { nextSocketSalt, nextSocketAddress } = await ctx.db.user.upsert({
	// 		where: {
	// 			address
	// 		},
	// 		update: {},
	// 		create: {
	// 			address
	// 		}
	// 	})
	// 	if (!nextSocketSalt || !nextSocketAddress) {
	// 		// * Need to cast the type to be viem compatible.
	// 		const salt: string = bytesToHex(
	// 			toBytes(
	// 				encodePacked(
	// 					["address", "uint80", "uint16"],
	// 					[
	// 						ctx.session.address as `0x${string}`,
	// 						BigInt(Date.now() / 1000),
	// 						CURRENT_IMPLEMENTATION_VERSION
	// 					]
	// 				),
	// 				{ size: 32 }
	// 			)
	// 		)
	// 		// TODO: Need to call into the contract to get the next
	// 		//		 address that will be deployed.
	// 		const address = ""
	// 		// TODO: Save the address and salt to the database.
	// 		await ctx.db.user.update({
	// 			where: {
	// 				address
	// 			},
	// 			data: {
	// 				address,
	// 				nextSocketSalt: salt,
	// 				nextSocketAddress: address
	// 			}
	// 		})
	// 		return { salt, address }
	// 	}
	// 	return {
	// 		salt: nextSocketSalt,
	// 		address: nextSocketAddress
	// 	}
	// }),
	all: protectedProcedure.query(async ({ ctx }) => {
		try {
			return await ctx.db.userSocket.findMany({
				where: {
					userAddress: ctx.session.address
				}
			})
		} catch (error) {
			throw new TRPCError({
				code: "BAD_REQUEST"
			})
		}
	}),
	get: protectedProcedure.input(z.string()).query(async ({ ctx, input }) => {
		try {
			const socket = await ctx.db.userSocket.findUnique({
				where: {
					userAddress_socketAddress: {
						userAddress: ctx.session.address,
						socketAddress: input
					}
				}
			})

			if (!socket) throw new TRPCError({ code: "NOT_FOUND" })

			return socket
		} catch (error) {
			throw new TRPCError({
				code: "BAD_REQUEST"
			})
		}
	}),
	add: protectedProcedure
		.input(z.string().optional())
		.mutation(async ({ ctx, input }) => {
			try {
				const sockets = await ctx.db.userSocket.findMany({
					where: { userAddress: ctx.session.address }
				})
				const count = sockets.length
				const name = `Socket #${count + 1}`

				const socketAddress = input || TEMPORARY_ADDRESS

				const socket = await ctx.db.userSocket.create({
					data: {
						userAddress: ctx.session.address,
						socketAddress,
						name
					}
				})

				ctx.emitter.emit(events.add, socket)

				return socket
			} catch (error) {
				throw new TRPCError({
					code: "BAD_REQUEST"
				})
			}
		}),
	onAdd: protectedProcedure.subscription(async ({ ctx }) => {
		return observable<UserSocket>(emit => {
			const handleAdd = (data: UserSocket) => {
				if (data.userAddress === ctx.session.address) emit.next(data)
			}
			ctx.emitter.on(events.add, handleAdd)
			return () => ctx.emitter.off(events.add, handleAdd)
		})
	}),
	rename: protectedProcedure
		.input(
			z.object({
				address: z.string(),
				name: z.string()
			})
		)
		.mutation(async ({ ctx, input }) => {
			try {
				const socket = await ctx.db.userSocket.update({
					where: {
						userAddress_socketAddress: {
							userAddress: ctx.session.address,
							socketAddress: input.address
						}
					},
					data: {
						name: input.name
					}
				})

				ctx.emitter.emit(events.rename, socket)

				return socket
			} catch (error) {
				throw new TRPCError({
					code: "BAD_REQUEST"
				})
			}
		}),
	onRename: protectedProcedure.subscription(async ({ ctx }) => {
		return observable<UserSocket>(emit => {
			const handleRename = (data: UserSocket) => {
				if (data.userAddress === ctx.session.address) emit.next(data)
			}
			ctx.emitter.on(events.rename, handleRename)
			return () => ctx.emitter.off(events.rename, handleRename)
		})
	}),
	balances: protectedProcedure.input(z.string()).query(async ({ input }) => {
		try {
			if (input === "") return []

			return await getBalances(input)
		} catch (e) {
			throw new TRPCError({ code: "BAD_REQUEST" })
		}
	})
})
