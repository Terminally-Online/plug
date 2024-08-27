import { z } from "zod"

import { getPositions } from "@/lib"
import { anonymousProtectedProcedure, createTRPCRouter } from "@/server/api/trpc"
import { getDominantColor } from "@/server/color"

export const misc = createTRPCRouter({
	featureRequest: anonymousProtectedProcedure
		.input(z.object({ context: z.string(), message: z.string().optional() }))
		.mutation(async ({ input, ctx }) => {
			return await ctx.db.featureRequest.create({
				data: {
					userAddress: ctx.session.address,
					context: input.context,
					message: input.message
				}
			})
		}),
	search: anonymousProtectedProcedure.input(z.string().optional()).query(async ({ input, ctx }) => {
		// TODO: Handle private plugs and exclude them while combining the self-results.
		const plugs = await ctx.db.workflow.findMany({
			where: {
				name: {
					contains: input,
					mode: "insensitive",
					notIn: ["Untitled Plug", ""]
				}
			},
			orderBy: { updatedAt: "desc" }
		})

		const { tokens } = ctx.session.user.anonymous ? { tokens: [] } : await getPositions(ctx.session.address, input)

		const collectibles = ctx.session.user.anonymous
			? { collectibles: [] }
			: await ctx.db.openseaCollection.findMany({
					where: {
						AND: [
							{
								OR: [
									{
										collectibles: {
											some: {
												cacheSocketId: ctx.session.address,
												OR: [
													{
														name: {
															contains: input,
															mode: "insensitive"
														}
													},
													{
														description: {
															contains: input,
															mode: "insensitive"
														}
													},
													{
														cacheChain: {
															contains: input,
															mode: "insensitive"
														}
													}
												]
											}
										}
									},
									{
										name: {
											contains: input,
											mode: "insensitive"
										}
									}
								]
							},
							{
								collectibles: {
									some: {
										cacheSocketId: ctx.session.address
									}
								}
							}
						]
					},
					include: {
						collectibles: {
							where: {
								cacheSocketId: ctx.session.address,
								OR: [
									{
										name: {
											contains: input,
											mode: "insensitive"
										}
									},
									{
										description: {
											contains: input,
											mode: "insensitive"
										}
									},
									{
										cacheChain: {
											contains: input,
											mode: "insensitive"
										}
									}
								]
							}
						}
					},
					orderBy: { createdAt: "desc" }
				})

		return {
			plugs,
			tokens,
			collectibles
		}
	}),
	extractDominantColor: anonymousProtectedProcedure.input(z.string()).query(async ({ input }) => await getDominantColor(input))
})
