import { TRPCError } from "@trpc/server"

import { z } from "zod"

import { anonymousProtectedProcedure, createTRPCRouter, protectedProcedure } from "@/server/api/trpc"

export const referral = createTRPCRouter({
	request: protectedProcedure.mutation(async ({ ctx }) => {
		return ctx.db.socketIdentity.update({
			where: { socketId: ctx.session.address },
			data: {
				requestedAt: new Date()
			}
		})
	}),

	submit: protectedProcedure.input(z.object({ referrerAddress: z.string() })).mutation(async ({ ctx, input }) => {
		const referrer = await ctx.db.socketIdentity.findFirst({
			where: {
				referralCode: {
					equals: input.referrerAddress,
					mode: "insensitive"
				},
				approvedAt: { not: null }
			},
			select: { socketId: true }
		})

		if (!referrer) {
			throw new TRPCError({
				code: "BAD_REQUEST",
				message: "Referral code provided is not valid."
			})
		}

		return ctx.db.socketIdentity.update({
			where: { socketId: ctx.session.address },
			data: {
				approvedAt: new Date(),
				referrerId: referrer.socketId
			}
		})
	}),

	stats: anonymousProtectedProcedure.query(async ({ ctx }) => {
		const now = new Date()

		const periods = Array.from({ length: 4 })
			.map((_, i) => {
				const date = new Date()
				date.setDate(now.getDate() - i * 7)
				return date
			})
			.reverse()

		const referralCounts = await Promise.all(
			periods.map(async date => {
				const startOfWeek = new Date(date)
				startOfWeek.setDate(date.getDate() - date.getDay()) // Sunday
				startOfWeek.setHours(0, 0, 0, 0)

				const endOfWeek = new Date(startOfWeek)
				endOfWeek.setDate(startOfWeek.getDate() + 6) // Saturday
				endOfWeek.setHours(23, 59, 59, 999)

				const count = await ctx.db.socketIdentity.count({
					where: {
						approvedAt: {
							gte: startOfWeek,
							lte: endOfWeek
						},
						referrerId: ctx.session.address
					}
				})

				return count
			})
		)

		return {
			counts: referralCounts,
			periods: periods.map(date => {
				const startOfWeek = new Date(date)
				startOfWeek.setDate(date.getDate() - date.getDay())
				return {
					weekStart: startOfWeek.toISOString(),
					weekEnd: new Date(startOfWeek.getTime() + 6 * 24 * 60 * 60 * 1000).toISOString()
				}
			})
		}
	})
})
