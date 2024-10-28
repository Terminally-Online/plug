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
		const socket = await ctx.db.userSocket.findFirst({
			where: { id: ctx.session.address }
		})

		if (socket === null) return { plugs: [], tokens: [], collectibles: [] }

		// TODO(#403): Handle private plugs and exclude them while combining the self-results.
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

		const { tokens } = ctx.session.user.anonymous
			? { tokens: [] }
			: await getPositions(ctx.session.address, socket.socketAddress, input)

		const cacheId = `${ctx.session.address}-${socket.socketAddress}`

		const collectibles = ctx.session.user.anonymous
			? []
			: await ctx.db.collection.findMany({
					where: {
						AND: [
							{
								OR: [
									{
										collectibles: {
											some: {
												cacheId,
												OR: [
													{
														name: {
															contains: input,
															mode: "insensitive"
														}
													}
													// {
													// 	description: {
													// 		contains: input,
													// 		mode: "insensitive"
													// 	}
													// },
													// {
													// 	cacheChain: {
													// 		contains: input,
													// 		mode: "insensitive"
													// 	}
													// }
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
										cacheId
									}
								}
							}
						]
					},
					include: {
						collectibles: {
							where: {
								cacheId,
								OR: [
									{
										name: {
											contains: input,
											mode: "insensitive"
										}
									}
									// {
									// 	description: {
									// 		contains: input,
									// 		mode: "insensitive"
									// 	}
									// },
									// {
									// 	cacheChain: {
									// 		contains: input,
									// 		mode: "insensitive"
									// 	}
									// }
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
	extractDominantColor: anonymousProtectedProcedure
		.input(z.string())
		.query(async ({ input }) => await getDominantColor(input))
})
