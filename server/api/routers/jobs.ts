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
			const queuedWorkflow = await tx.queuedWorkflow.findUnique({
				where: {
					id: input.id
				},
				include: {
					workflow: true
				}
			})

			if (queuedWorkflow === null) throw new TRPCError({ code: "NOT_FOUND" })

			const now = new Date()
			const nextSimulationAt = new Date(now.getTime() + queuedWorkflow.frequency * 60 * 1000)

			return await tx.queuedWorkflow.update({
				where: {
					id: queuedWorkflow.id
				},
				data: {
					nextSimulationAt
				}
			})
		})
	}),
	simulateNext: apiKeyProcedure.input(z.object({ count: z.number() }).nullish()).mutation(async ({ input, ctx }) => {
		const now = new Date()

		return await ctx.db.$transaction(async tx => {
			const queuedWorkflows = await tx.queuedWorkflow.findMany({
				where: {
					nextSimulationAt: {
						lte: now
					}
				},
				take: input?.count ?? 100,
				select: {
					id: true,
					workflow: { select: { actions: true, socket: { select: { id: true, socketAddress: true } } } }
				}
			})

			const parsedQueuedWorkflows = queuedWorkflows.map(queuedWorkflow => ({
				...queuedWorkflow,
				workflow: {
					...queuedWorkflow.workflow,
					actions: JSON.parse(queuedWorkflow.workflow.actions as string)
				}
			}))

			// TODO: Uncomment this when we are ready to save simulation results.
			// await Promise.all(
			// 	queuedWorkflows.map(queuedWorkflow =>
			// 		tx.queuedWorkflow.update({
			// 			where: { id: queuedWorkflow.id },
			// 			data: {
			// 				nextSimulationAt: new Date(now.getTime() + queuedWorkflow.frequency * 60 * 1000)
			// 			}
			// 		})
			// 	)
			// )

			return parsedQueuedWorkflows
		})
	})
})
