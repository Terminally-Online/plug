import { z } from "zod"

import { createTRPCRouter, protectedProcedure } from "@/server/api/trpc"
import { getDominantColor } from "@/server/color"

export const misc = createTRPCRouter({
	featureRequest: protectedProcedure
		.input(
			z.object({ context: z.string(), message: z.string().optional() })
		)
		.mutation(async ({ input, ctx }) => {
			const featureRequest = await ctx.db.featureRequest.create({
				data: {
					userAddress: ctx.session.address,
					context: input.context,
					message: input.message
				}
			})

			return featureRequest
		}),
	search: protectedProcedure
		.input(z.string().optional())
		.query(async ({ input, ctx }) => {
			return {
				plugs: await ctx.db.workflow.findMany({
					where: {
						name: {
							contains: input,
							mode: "insensitive",
							notIn: ["Untitled Plug", ""]
						}
					},
					orderBy: { updatedAt: "desc" }
				}),
				tokens: await ctx.db.tokenBalance.findMany({
					where: {
						name: {
							contains: input,
							mode: "insensitive",
							not: undefined
						}
					}
				}),
				collectibles: await ctx.db.openseaCollection.findMany({
					where: {
						name: {
							contains: input,
							mode: "insensitive",
							not: undefined
						}
					}
				})
			}
		}),
	extractDominantColor: protectedProcedure
		.input(z.string())
		.query(async ({ input }) => {
			return await getDominantColor(input)
		})
})
