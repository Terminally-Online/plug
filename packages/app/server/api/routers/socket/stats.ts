import { createTRPCRouter, protectedProcedure } from "@/server/api/trpc"
import { Prisma } from "@prisma/client"

function getWeekStart(date: Date) {
	const result = new Date(date)
	result.setDate(date.getDate() - ((date.getDay() + 6) % 7))
	result.setUTCHours(0, 0, 0, 0)
	return result
}

const QUERY_TIMEOUT = 10000;

export const stats = createTRPCRouter({
	get: protectedProcedure
		.query(async ({ ctx }) => {
			const address = ctx.session.address;
			const now = new Date()

			const periodRanges = Array.from({ length: 4 })
				.map((_, i) => {
					const date = new Date(now)
					date.setDate(now.getDate() - i * 7)
					const weekStart = getWeekStart(date)
					const weekEnd = new Date(weekStart)
					weekEnd.setDate(weekStart.getDate() + 7)
					return {
						weekStart,
						weekEnd,
						weekStartIso: weekStart.toISOString(),
						weekEndIso: new Date(weekStart.getTime() + 6 * 24 * 60 * 60 * 1000).toISOString()
					}
				})
				.reverse()

			const periodStarts = periodRanges.map(p => p.weekStart.toISOString())

			const createTimedQuery = <T>(query: Promise<T>): Promise<T> => {
				const timeout = new Promise<never>((_, reject) =>
					setTimeout(() => reject(new Error('Query timeout')), QUERY_TIMEOUT)
				);
				return Promise.race([query, timeout]) as Promise<T>;
			};

			try {
				const [
					plugCreationData,
					viewsData,
					referralsData,
					forksData,
				] = await Promise.all([
					createTimedQuery(ctx.db.$queryRaw`
							SELECT 
								${Prisma.raw(periodStarts.map((start, i) => {
						const end = i < periodStarts.length - 1 ? periodStarts[i + 1] : now.toISOString()
						return `COUNT(CASE WHEN "createdAt" >= '${start}' AND "createdAt" < '${end}' THEN 1 END) as period_${i}`
					}).join(', '))}
							FROM "Plug"
							WHERE "socketId" = ${address}
						`),

					createTimedQuery(ctx.db.$queryRaw`
							SELECT 
								${Prisma.raw(periodStarts.map((start, i) => {
						const end = i < periodStarts.length - 1 ? periodStarts[i + 1] : now.toISOString()
						return `COALESCE(SUM(CASE WHEN v."date" >= '${start}' AND v."date" < '${end}' THEN v."views" ELSE 0 END), 0) as period_${i}`
					}).join(', '))}
							FROM "View" v
							JOIN "Plug" p ON v."plugId" = p."id"
							WHERE p."socketId" = ${address}
						`),

					createTimedQuery(ctx.db.$queryRaw`
							SELECT 
								${Prisma.raw(periodStarts.map((start, i) => {
						const end = i < periodStarts.length - 1 ? periodStarts[i + 1] : now.toISOString()
						return `COUNT(CASE WHEN "approvedAt" >= '${start}' AND "approvedAt" < '${end}' THEN 1 END) as period_${i}`
					}).join(', '))}
							FROM "SocketIdentity"
							WHERE "referrerId" = ${address}
						`),

					createTimedQuery(ctx.db.$queryRaw`
							SELECT 
								${Prisma.raw(periodStarts.map((start, i) => {
						const end = i < periodStarts.length - 1 ? periodStarts[i + 1] : now.toISOString()
						return `COUNT(CASE WHEN f."createdAt" >= '${start}' AND f."createdAt" < '${end}' THEN 1 END) as period_${i}`
					}).join(', '))}
							FROM "Plug" f
							JOIN "Plug" p ON f."plugForkedId" = p."id"
							WHERE p."socketId" = ${address}
							AND f."socketId" != ${address}
						`)
				]);

				const extractPeriodData = (data: any) => {
					if (!data || !data[0]) return periodRanges.map(() => 0)
					return periodStarts.map((_, i) => Number(data[0][`period_${i}`] || 0))
				}

				return {
					counts: {
						plugs: extractPeriodData(plugCreationData),
						views: extractPeriodData(viewsData),
						referrals: extractPeriodData(referralsData),
						forks: extractPeriodData(forksData)
					},
					periods: periodRanges.map(period => ({
						weekStart: period.weekStartIso,
						weekEnd: period.weekEndIso
					}))
				};

			} catch (error) {
				console.error("Stats query error:", error);

				return {
					counts: {
						plugs: periodRanges.map(() => 0),
						views: periodRanges.map(() => 0),
						referrals: periodRanges.map(() => 0),
						forks: periodRanges.map(() => 0)
					},
					periods: periodRanges.map(period => ({
						weekStart: period.weekStartIso,
						weekEnd: period.weekEndIso
					}))
				};
			}
		}
		)
})
