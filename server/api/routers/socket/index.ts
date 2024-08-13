import { TRPCError } from "@trpc/server"

import { DEFAULT_VIEWS, VIEW_KEYS } from "@/lib"
import { createTRPCRouter, protectedProcedure } from "@/server/api/trpc"

import { balances } from "./balances"
import { columns } from "./columns"

const TEMPORARY_ADDRESS = "0x62180042606624f02d8a130da8a3171e9b33894d"

export const socket = createTRPCRouter({
	get: protectedProcedure.query(async ({ ctx }) => {
		const { columns } = await ctx.db.userSocket.upsert({
			where: {
				id: ctx.session.address
			},
			create: {
				id: ctx.session.address,
				socketAddress: TEMPORARY_ADDRESS,
				columns: {
					createMany: {
						data: DEFAULT_VIEWS
					}
				}
			},
			update: {},
			select: { columns: true }
		})

		// NOTE: Make sure the socket always has the global home column that is
		//       denoted by the -1 index as the column view should never show it.
		if (columns.find(column => column.index === -1) === undefined) {
			await ctx.db.consoleColumn.create({
				data: {
					key: VIEW_KEYS.HOME,
					index: -1,
					socketId: ctx.session.address
				}
			})
		}

		const socket = await ctx.db.userSocket.findFirst({
			where: { id: ctx.session.address },
			include: { columns: true }
		})

		if (socket === null) throw new TRPCError({ code: "NOT_FOUND" })

		return socket
	}),
	balances,
	columns
})
