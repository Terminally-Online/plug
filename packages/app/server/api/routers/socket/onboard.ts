import { z } from "zod"

import { anonymousProtectedProcedure, apiKeyProcedure, createTRPCRouter } from "../../trpc"

export const onboard = createTRPCRouter({
	onboard: anonymousProtectedProcedure
		.input(z.object({
			onboardingColor: z.string().optional()
		}))
		.mutation(async ({ ctx, input }) => {
			return await ctx.db.socketIdentity.update({
				where: { socketId: ctx.session.address },
				data: { 
					onboardingAt: new Date(),
					onboardingColor: input.onboardingColor
				}
			})
		}),

	onboarding: apiKeyProcedure.query(async ({ ctx }) => {
		return await ctx.db.socketIdentity.findMany({
			where: {
				onboardedAt: null,
				onboardingAt: { not: null }
			}
		})
	}),

	onboarded: apiKeyProcedure.input(z.array(z.string())).mutation(async ({ ctx, input }) => {
		await ctx.db.socketIdentity.updateMany({
			where: {
				socketId: {
					in: input
				}
			},
			data: {
				onboardedAt: new Date()
			}
		})
	})
})
