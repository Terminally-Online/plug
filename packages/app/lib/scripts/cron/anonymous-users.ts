import { db } from "@/server";

const CLEANUP_OLDER_THAN_DAYS = 3

const work = async () => {
	const cutoffDate = new Date()
	cutoffDate.setDate(cutoffDate.getDate() - CLEANUP_OLDER_THAN_DAYS)
	const removed = await db.socket.deleteMany({
		where: {
			id: {
				startsWith: "anonymous-"
			},
			createdAt: {
				lt: cutoffDate
			}
		}
	})
	console.log(`Cleaned up ${removed.count} anonymous sockets older than ${CLEANUP_OLDER_THAN_DAYS} days`)
}

work()
