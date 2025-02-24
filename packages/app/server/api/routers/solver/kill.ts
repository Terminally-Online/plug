import { TRPCError } from "@trpc/server"

import { kill, killed } from "@/lib"

import { createTRPCRouter, protectedProcedure } from "../../trpc"

export const killer = createTRPCRouter({
	killed: protectedProcedure.query(async () => await killed()),

	kill: protectedProcedure.mutation(async ({ ctx }) => {
		const socket = await ctx.db.socket.findFirst({
			where: { id: ctx.session?.address }
		})

		if (!socket?.admin) throw new TRPCError({ code: "UNAUTHORIZED" })

		return await kill()
	})
})
