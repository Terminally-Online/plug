import { z } from "zod"

import { Prisma } from "@prisma/client"
import { TRPCError } from "@trpc/server"
import { observable } from "@trpc/server/observable"

import { getBalances, getTokens } from "@/lib"
import { getCollectibles } from "@/lib/functions/opensea"
import { createTRPCRouter, protectedProcedure } from "@/server/api/trpc"

export const userSocket = Prisma.validator<Prisma.UserSocketDefaultArgs>()({})
export type UserSocket = Prisma.UserSocketGetPayload<typeof userSocket>

const events = {
	add: "add-socket",
	rename: "rename-socket"
}

const TEMPORARY_ADDRESS = "0x62180042606624f02d8a130da8a3171e9b33894d"

export const socket = createTRPCRouter({
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
	}),

	tokens: protectedProcedure
		.input(z.string().optional())
		.query(async ({ input }) => {
			try {
				if (input === undefined) return []

				return await getTokens(input)
			} catch (e) {
				console.error(e)
				throw new TRPCError({ code: "BAD_REQUEST" })
			}
		}),

	collectibles: protectedProcedure
		.input(z.string().optional())
		.query(async ({ input }) => {
			try {
				if (input === undefined) return {}

				return await getCollectibles(input)
			} catch (e) {
				console.error(e)
				throw new TRPCError({ code: "BAD_REQUEST" })
			}
		})
})
