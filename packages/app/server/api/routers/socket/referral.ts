import { TRPCError } from "@trpc/server"

import { z } from "zod"

import { createTRPCRouter, protectedProcedure } from "@/server/api/trpc"

export const referral = createTRPCRouter({
	request: protectedProcedure.mutation(async ({ ctx }) => {
		return ctx.db.socketIdentity.update({
			where: { socketId: ctx.session.address },
			data: {
				requestedAt: new Date()
			}
		})
	}),

	submit: protectedProcedure.input(z.string()).mutation(async ({ ctx, input }) => {
		const currentUser = await ctx.db.socketIdentity.findFirst({
			where: { socketId: ctx.session.address }
		})

		if (currentUser?.referrerId) 
			throw new TRPCError({
				code: "BAD_REQUEST",
				message: "You are already approved and cannot use another referral code."
			})

		const referrer = await ctx.db.socketIdentity.findFirst({
			where: {
				referralCode: input
			}
		})

		if (!referrer || !referrer.approvedAt || referrer.socketId === ctx.session.address) throw new TRPCError({
			code: "BAD_REQUEST",
			message: "Referral code provided is not valid."
		})

		return ctx.db.socketIdentity.update({
			where: { socketId: ctx.session.address },
			data: {
				approvedAt: new Date(),
				referrerId: referrer.socketId
			}
		})
	})
})
