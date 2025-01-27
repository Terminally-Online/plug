import { z } from "zod"

import { getMetadataForToken } from "@/lib/opensea/metadata"

import { apiKeyProcedure, createTRPCRouter } from "../../trpc"

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
})
