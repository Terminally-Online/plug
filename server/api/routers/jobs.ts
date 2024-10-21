import { apiKeyProcedure, createTRPCRouter } from "../trpc"

const CLEANUP_OLDER_THAN_DAYS = 7

export const jobs = createTRPCRouter({
	cleanupAnonymousUsers: apiKeyProcedure.mutation(async ({ ctx }) => {
		const cutoffDate = new Date()
		cutoffDate.setDate(cutoffDate.getDate() - CLEANUP_OLDER_THAN_DAYS)

		return await ctx.db.userSocket.deleteMany({
			where: {
				id: {
					startsWith: "anonymous-"
				},
				createdAt: {
					lt: cutoffDate
				}
			}
		})
	})
})
