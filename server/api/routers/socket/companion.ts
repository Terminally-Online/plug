import { TRPCError } from "@trpc/server"

import { z } from "zod"

import { anonymousProtectedProcedure, createTRPCRouter } from "../../trpc"

export const companion = createTRPCRouter({
	feed: anonymousProtectedProcedure.mutation(async ({ ctx }) => {
		// TODO: This will error out when one does not exist.
		const existingCompanion = await ctx.db.companion.findUnique({
			where: { socketId: ctx.session.address }
		})

		if (existingCompanion && existingCompanion.lastFeedAt) {
			const timeSinceLastFeed = new Date().getTime() - existingCompanion.lastFeedAt.getTime()
			if (timeSinceLastFeed < 24 * 60 * 60 * 1000) {
				throw new TRPCError({
					code: "BAD_REQUEST",
					message: "Companion can only be fed once per day"
				})
			}
		}

		const treatsToFeed = Math.floor(Math.random() * 5 + 2 + Math.random() > 0.95 ? 20 * Math.random() * 0.3 : 0)
		// TODO: Already doing date calculations -- Might as well keep track of their streak
		// TODO: I don't know if this is actually the right way to see if the times are 24 hours within another.
		// const streaking = Number(new Date().toISOString()) - Number(existingCompanion?.lastFeedAt.toISOString()) > 24 * 60 * 60 * 1000
		// const streak = existingCompanion.streak++

		return await ctx.db.companion.upsert({
			where: { socketId: ctx.session.address },
			update: {
				feedCount: { increment: 1 },
				treatsFed: { increment: treatsToFeed },
				lastFeedAt: new Date()
			},
			create: {
				socketId: ctx.session.address,
				name: "New Companion",
				feedCount: 1,
				treatsFed: treatsToFeed
			}
		})
	})
})
