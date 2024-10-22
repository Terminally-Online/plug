import { TRPCError } from "@trpc/server"

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
	simulate: apiKeyProcedure.input(z.object({ id: z.string() })).mutation(async ({ input, ctx }) => {
		return await ctx.db.$transaction(async tx => {
			const workflow = await tx.workflow.findUnique({
				where: {
					id: input.id
				}
			})

			if (workflow === null) throw new TRPCError({ code: "NOT_FOUND" })

			return await tx.workflow.update({
				where: {
					id: workflow.id
				},
				data: {
					nextSimulationAt: new Date(Date.now() + workflow.frequency * 60 * 1000)
				}
			})
		})
	}),
	simulateNext: apiKeyProcedure
		.input(z.object({ count: z.number() }).nullish())
		.mutation(async ({ input, ctx }) => {
			const now = new Date()

			return await ctx.db.$transaction(async tx => {
				const workflows = await tx.workflow.findMany({
					where: {
						nextSimulationAt: {
							lte: now
						}
					},
					take: input?.count ?? 100
				})

				await Promise.all(
					workflows.map(workflow =>
						tx.workflow.update({
							where: { id: workflow.id },
							data: {
								nextSimulationAt: new Date(now.getTime() + workflow.frequency * 60 * 1000)
							}
						})
					)
				)

				return workflows
			})
		})
})
