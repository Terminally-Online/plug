import { z } from "zod"

import { getMetadataForToken } from "@/lib/opensea/metadata"

import { apiKeyProcedure, createTRPCRouter, protectedProcedure } from "../../trpc"
import { TRPCError } from "@trpc/server"

const CLEANUP_OLDER_THAN_DAYS = 7

export const maintenance = createTRPCRouter({
	anonymous: apiKeyProcedure.mutation(async ({ ctx }) => {
		const cutoffDate = new Date()
		cutoffDate.setDate(cutoffDate.getDate() - CLEANUP_OLDER_THAN_DAYS)

		return await ctx.db.userSocket.deleteMany({
			where: {
				id: {
					startsWith: "anonymous-"
				},
				createdAt: {
					lt: cutoffDate
				}
			}
		})
	}),

	collectibleMetadata: apiKeyProcedure
		.input(z.object({ count: z.number().min(1).max(100) }).nullish())
		.mutation(async ({ input, ctx }) => {
			const collectiblesWithoutMetadata = await ctx.db.collectible.findMany({
				where: {
					collectibleMetadata: null
				},
				take: input ? input.count : 50,
				include: {
					collection: true
				}
			})

			const results = await Promise.all(
				collectiblesWithoutMetadata.map(async collectible => {
					const token = {
						address: collectible.collectionAddress,
						chain: collectible.collectionChain,
						tokenId: collectible.tokenId
					}

					try {
						await getMetadataForToken(token)
						return { success: true, token }
					} catch (error) {
						return { success: false, token }
					}
				})
			)

			return results
		}),

	getIsSolverKilled: protectedProcedure.query(async ({ ctx }) => {
		// First check if the user is an admin
		const socket = await ctx.db.userSocket.findFirst({
			where: { id: ctx.session?.address }
		})
		if (!socket?.admin) {
			throw new TRPCError({
				code: 'UNAUTHORIZED',
				message: 'Admin access required'
			})
		}

		try {
			const response = await fetch(`${process.env.SOLVER_URL}/solver/kill`, {
				headers: {
					'x-api-key': process.env.ADMIN_API_KEY ?? "",
				}
			})

			if (!response.ok) {
				throw new Error(`HTTP error! status: ${response.status}`)
			}

			const data = await response.json()
			return { killed: data.killed }
		} catch (error) {
			throw new TRPCError({
				code: 'INTERNAL_SERVER_ERROR',
				message: `Failed to get solver kill status: ${error}`
			})
		}
	}),

	toggleSolverKill: protectedProcedure
		.mutation(async ({ ctx }) => {
			// Check if user is admin
			const socket = await ctx.db.userSocket.findFirst({
				where: { id: ctx.session?.address }
			})
			if (!socket?.admin) {
				throw new TRPCError({
					code: 'UNAUTHORIZED',
					message: 'Admin access required'
				})
			}

			try {
				const response = await fetch(`${process.env.SOLVER_URL}/solver/kill`, {
					method: 'POST',
					headers: {
						'x-api-key': process.env.ADMIN_API_KEY ?? "",
					},
				})

				if (!response.ok) {
					throw new Error(`HTTP error! status: ${response.status}`)
				}

				const data = await response.json()
				return { killed: data.killed }
			} catch (error) {
				throw new TRPCError({
					code: 'INTERNAL_SERVER_ERROR',
					message: `Failed to toggle solver kill: ${error}`
				})
			}
		})
})
