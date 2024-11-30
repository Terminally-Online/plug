import { TRPCError } from "@trpc/server"

import { anonymousProtectedProcedure, createTRPCRouter } from "../../trpc"

export const companion = createTRPCRouter({
	feed: anonymousProtectedProcedure.mutation(async ({ ctx }) => {
		const existingCompanion = await ctx.db.companion.findUnique({
			where: { socketId: ctx.session.address }
		})

		const now = new Date()

		if (existingCompanion && existingCompanion.lastFeedAt) {
			const timeSinceLastFeed = now.getTime() - existingCompanion.lastFeedAt.getTime()
			if (timeSinceLastFeed < 24 * 60 * 60 * 1000) {
				throw new TRPCError({
					code: "BAD_REQUEST",
					message: "Companion can only be fed once per day"
				})
			}
		}

		const treatsToFeed = Math.floor(Math.random() * 5 + 2 + Math.random() > 0.95 ? 20 * Math.random() * 0.3 : 0) + 2

		let newStreak = 1
		if (existingCompanion && existingCompanion.lastFeedAt) {
			const daysSinceLastFeed = Math.floor(
				(now.getTime() - existingCompanion.lastFeedAt.getTime()) / (24 * 60 * 60 * 1000)
			)
			if (daysSinceLastFeed === 1) {
				newStreak = existingCompanion.streak + 1
			} else if (daysSinceLastFeed > 1) {
				newStreak = 1
			} else {
				newStreak = existingCompanion.streak
			}
		}

		return await ctx.db.companion.upsert({
			where: { socketId: ctx.session.address },
			update: {
				feedCount: { increment: 1 },
				treatsFed: { increment: treatsToFeed },
				lastFeedAt: now,
				streak: newStreak
			},
			create: {
				socketId: ctx.session.address,
				name: "New Companion",
				feedCount: 1,
				treatsFed: treatsToFeed,
				lastFeedAt: now,
				streak: 1
			}
		})
	})
})
