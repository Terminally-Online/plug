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
		const referrer = await ctx.db.socketIdentity.findFirst({
			where: {
				referralCode: input
			}
		})

		if (!referrer)
			throw new TRPCError({
				code: "BAD_REQUEST",
				message: "Referral code provided is not valid."
			})

		if (referrer.socketId === ctx.session.address)
			throw new TRPCError({
				code: "BAD_REQUEST",
				message:
					"You cannot refer yourself! I respect the effort though. Message @onplug_io on Twitter with a screenshot of this page for an immediate access."
			})

		if (!referrer.approvedAt)
			throw new TRPCError({
				code: "BAD_REQUEST",
				message: "Referral code provided is from a user not yet approved."
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
