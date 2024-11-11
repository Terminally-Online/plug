import { z } from "zod"

import { createTRPCRouter, protectedProcedure } from "@/server/api/trpc"

export const stats = createTRPCRouter({
	get: protectedProcedure.query(async ({ ctx }) => {
		const now = new Date()
		const periods = Array.from({ length: 4 })
			.map((_, i) => {
				const date = new Date()
				date.setDate(now.getDate() - i * 7)
				return date
			})
			.reverse()

		const getWeekStart = (date: Date) => {
			const result = new Date(date)
			result.setDate(date.getDate() - ((date.getDay() + 6) % 7)) // Get Monday
			result.setUTCHours(0, 0, 0, 0)
			return result
		}

		const [referralCounts, viewCounts] = await Promise.all([
			// Existing referral counts query
			Promise.all(
				periods.map(async date => {
					const weekStart = getWeekStart(date)
					const weekEnd = new Date(weekStart)
					weekEnd.setDate(weekStart.getDate() + 7)

					return ctx.db.socketIdentity.count({
						where: {
							approvedAt: {
								gte: weekStart,
								lt: weekEnd
							},
							referrerId: ctx.session.address
						}
					})
				})
			),
			// Simplified view counts query - now we just need to look up the exact weeks
			Promise.all(
				periods.map(async date => {
					const weekStart = getWeekStart(date)

					const views = await ctx.db.view.aggregate({
						where: {
							workflow: {
								socketId: ctx.session.address
							},
							date: weekStart
						},
						_sum: {
							views: true
						}
					})

					return views._sum.views || 0
				})
			)
		])

		return {
			counts: {
				referrals: referralCounts,
				views: viewCounts
			},
			periods: periods.map(date => {
				const weekStart = getWeekStart(date)
				return {
					weekStart: weekStart.toISOString(),
					weekEnd: new Date(weekStart.getTime() + 6 * 24 * 60 * 60 * 1000).toISOString()
				}
			})
		}
	})
})
