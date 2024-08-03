import { createTRPCRouter, protectedProcedure } from "@/server/api/trpc"

import { balances } from "./balances"
import { columns, DEFAULT_COLUMNS } from "./columns"

const TEMPORARY_ADDRESS = "0x62180042606624f02d8a130da8a3171e9b33894d"

export const socket = createTRPCRouter({
	get: protectedProcedure.query(async ({ ctx }) => {
		return await ctx.db.userSocket.upsert({
			where: {
				id: ctx.session.address
			},
			create: {
				id: ctx.session.address,
				socketAddress: TEMPORARY_ADDRESS,
				columns: {
					createMany: {
						data: DEFAULT_COLUMNS
					}
				}
			},
			update: {},
			include: { columns: true, collectibles: true }
		})
	}),
	balances,
	columns
})
