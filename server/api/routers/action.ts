import { z } from "zod"

import { TRPCError } from "@trpc/server"

import { createTRPCRouter, protectedProcedure } from "@/server/api/trpc"

import { events } from "./plug"

export const action = createTRPCRouter({
	edit: protectedProcedure
		.input(
			z.object({
				id: z.string().optional(),
				actions: z.string()
			})
		)
		.mutation(async ({ input, ctx }) => {
			try {
				if (input.id === undefined)
					throw new TRPCError({ code: "BAD_REQUEST" })

				const plug = await ctx.db.workflow.update({
					where: { id: input.id },
					data: { actions: input.actions, updatedAt: new Date() }
				})

				ctx.emitter.emit(events.edit, plug)

				return plug
			} catch (error) {
				throw new TRPCError({ code: "BAD_REQUEST" })
			}
		})
})
