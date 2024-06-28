import { bytesToHex, encodePacked, toBytes } from "viem"
import { z } from "zod"

import { Prisma } from "@prisma/client"
import { TRPCError } from "@trpc/server"
import { observable } from "@trpc/server/observable"

import { createTRPCRouter, protectedProcedure } from "@/server/api/trpc"

export const misc = createTRPCRouter({
	featureRequest: protectedProcedure
		.input(
			z.object({ context: z.string(), message: z.string().optional() })
		)
		.mutation(async ({ input, ctx }) => {
			try {
				const featureRequest = await ctx.db.featureRequest.create({
					data: {
						userAddress: ctx.session.address,
						context: input.context,
						message: input.message
					}
				})

				return featureRequest
			} catch (error) {
				throw new TRPCError({ code: "BAD_REQUEST" })
			}
		})
})
