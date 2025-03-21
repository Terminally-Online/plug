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
			result.setDate(date.getDate() - ((date.getDay() + 6) % 7))
			result.setUTCHours(0, 0, 0, 0)
			return result
		}

		const [referralCounts, viewCounts, plugCreationCounts, forkCounts] = await Promise.all([
			// Referrals - how many users were referred by this user
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

			// Views - how many times plugs were viewed
			Promise.all(
				periods.map(async date => {
					const weekStart = getWeekStart(date)
					const weekEnd = new Date(weekStart)
					weekEnd.setDate(weekStart.getDate() + 7)

					const views = await ctx.db.view.aggregate({
						where: {
							plug: {
								socketId: ctx.session.address
							},
							date: {
								gte: weekStart,
								lt: weekEnd
							}
						},
						_sum: {
							views: true
						}
					})

					return views._sum.views || 0
				})
			),

			// Plugs Created - how many new plugs the creator made
			Promise.all(
				periods.map(async date => {
					const weekStart = getWeekStart(date)
					const weekEnd = new Date(weekStart)
					weekEnd.setDate(weekStart.getDate() + 7)

					return ctx.db.plug.count({
						where: {
							socketId: ctx.session.address,
							createdAt: {
								gte: weekStart,
								lt: weekEnd
							}
						}
					})
				})
			),

			// Forks - how many times others forked the creator's plugs
			Promise.all(
				periods.map(async date => {
					const weekStart = getWeekStart(date)
					const weekEnd = new Date(weekStart)
					weekEnd.setDate(weekStart.getDate() + 7)

					// Get all plugs created by this user
					const userPlugs = await ctx.db.plug.findMany({
						where: { socketId: ctx.session.address },
						select: { id: true }
					})
					
					const userPlugIds = userPlugs.map(plug => plug.id)

					// Count plugs created by others that forked from user's plugs
					return ctx.db.plug.count({
						where: {
							plugForkedId: {
								in: userPlugIds
							},
							socketId: {
								not: ctx.session.address // exclude self-forks
							},
							createdAt: {
								gte: weekStart,
								lt: weekEnd
							}
						}
					})
				})
			)
		])

		return {
			counts: {
				referrals: referralCounts,
				views: viewCounts,
				plugs: plugCreationCounts,
				forks: forkCounts
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
