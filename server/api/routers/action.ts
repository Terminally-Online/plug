import { z } from "zod"

import { anonymousProtectedProcedure, createTRPCRouter } from "@/server/api/trpc"
import { TRPCError } from "@trpc/server"

import { categories } from "@/lib"

import { events } from "./plug"

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
			try {
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
			} catch (error) {
				throw new TRPCError({ code: "BAD_REQUEST" })
			}
		})
})
