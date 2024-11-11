import { TRPCError } from "@trpc/server"

import { z } from "zod"

import { anonymousProtectedProcedure, createTRPCRouter } from "@/server/api/trpc"

export const events = {
	edit: "edit-plug",
	queue: "queue-plug"
} as const

export const action = createTRPCRouter({
	edit: anonymousProtectedProcedure
		.input(
			z.object({
				id: z.string().optional(),
				actions: z.string()
			})
		)
		.mutation(async ({ input, ctx }) => {
			if (input.id === undefined) throw new TRPCError({ code: "BAD_REQUEST" })

			// NOTE: We wrap this endpoint in a try/catch because we do not want JSON.parse to crash the server.
			try {
				// TODO: For this to work we need to be sending the full action object to the database.
				// const actions = JSON.parse(input.actions)
				const tags: string[] = []

				const plug = await ctx.db.workflow.update({
					where: { id: input.id, socketId: ctx.session.address },
					data: {
						actions: input.actions,
						tags,
						updatedAt: new Date()
					}
				})

				ctx.emitter.emit(events.edit, plug)

				return plug
			} catch {
				throw new TRPCError({ code: "BAD_REQUEST" })
			}
		})
})
