import { z } from "zod"

import { getPositions } from "@/lib"
import { anonymousProtectedProcedure, createTRPCRouter } from "@/server/api/trpc"
import { getDominantColor } from "@/server/color"

export const misc = createTRPCRouter({
	search: anonymousProtectedProcedure.input(z.string().optional()).query(async ({ input, ctx }) => {
		const socket = await ctx.db.socket.findFirst({
			where: { id: ctx.session.address }
		})

		if (socket === null) return { plugs: [], tokens: [], collectibles: [] }

		const addForkCounts = async <T extends { id: string }>(plugs: T[]) => {
			const forkCounts = await Promise.all(
				plugs.map(async plug => ({
					id: plug.id,
					forks: await ctx.db.plug.count({
						where: { plugForkedId: plug.id }
					})
				}))
			)

			return plugs.map(plug => ({
				...plug,
				forkCount: forkCounts.find(f => f.id === plug.id)?.forks ?? 0
			}))
		}

		const plugs = await addForkCounts(
			await ctx.db.plug.findMany({
				where: {
					name: {
						contains: input,
						mode: "insensitive",
						notIn: ["Untitled Plug", ""]
					},
					OR: [
						// Show user's own plugs (including private)
						{
							socketId: ctx.session.address
						},
						// Show other users' public plugs
						{
							socketId: {
								not: ctx.session.address
							},
							isPrivate: false,
							actions: { not: "[]" }
						}
					]
				},
				orderBy: { updatedAt: "desc" }
			})
		)

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
