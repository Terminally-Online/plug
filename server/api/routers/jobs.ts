import { z } from "zod"

import { getMetadataForToken } from "@/lib/opensea/metadata"

import { apiKeyProcedure, createTRPCRouter } from "../trpc"

const CLEANUP_OLDER_THAN_DAYS = 7

export const jobs = createTRPCRouter({
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
	fetchMetadata: apiKeyProcedure
		.input(z.object({ count: z.number().min(1).max(100) }))
		.mutation(async ({ input, ctx }) => {
			const collectiblesWithoutMetadata = await ctx.db.collectible.findMany({
				where: {
					collectibleMetadata: null
				},
				take: input.count,
				include: {
					collection: true
				}
			})

			const results = await Promise.all(
				collectiblesWithoutMetadata.map(async collectible => {
					try {
						const metadata = await getMetadataForToken({
							address: collectible.collectionAddress,
							chain: collectible.collectionChain,
							tokenId: collectible.tokenId
						})
						return { success: true, tokenId: collectible.tokenId, metadata }
					} catch (error) {
						return { success: false, tokenId: collectible.tokenId, error: (error as Error).message }
					}
				})
			)

			return results
		})
})
