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

			// TODO: We need to factor the tags based on the protocols/actions from the solver.
			const tags: string[] = []

			// TODO: Attempt a JSON parse on the actions to confirm it is valid before saving it to the database.

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
		})
})
