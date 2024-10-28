import { TRPCError } from "@trpc/server"

import { z } from "zod"

import { categories } from "@/lib"
import { anonymousProtectedProcedure, createTRPCRouter } from "@/server/api/trpc"

export const events = {
	edit: "edit-plug",
	queue: "queue-plug"
} as const

const getTags = (actions: string) => {
	const parsed: Array<{ categoryName: string; actionName: string }> = JSON.parse(actions)

	return Array.from(
		new Set(
			parsed
				.map(action => action.categoryName)
				.map(categoryName => categories[categoryName].tags)
				.flat()
		)
	)
}

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

			const tags = getTags(input.actions)

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
