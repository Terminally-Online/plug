import { TRPCError } from "@trpc/server"

import { z } from "zod"

import { Prisma } from "@prisma/client"

import { createIntent, deleteIntent, getIntent, toggleIntent, Intent } from "@/lib"
import { createTRPCRouter, protectedProcedure } from "@/server/api/trpc"
import { subscription, subscriptions } from "@/server/subscription"

export const activity = createTRPCRouter({
	get: protectedProcedure.query(async ({ ctx }) => {
			const intents = await getIntent({
				address: ctx.session.address,
			})
		return [] as Array<Intent>

		// return await ctx.db.intent.findMany({
		// 	where: { plug: { socketId: ctx.session.address } },
		// 	orderBy: { createdAt: "desc" },
		// 	include: { plug: true, runs: { orderBy: { createdAt: "desc" } } }
		// })
	}),

	queue: protectedProcedure
		.input(
			z.object({
				plugId: z.string(),
				chainId: z.number(),
				frequency: z.number(),
				startAt: z.coerce.date(),
				endAt: z.coerce.date().optional()
			})
		)
		.mutation(async ({ input, ctx }) => {
			const plug = await ctx.db.plug.findUnique({
				where: {
					id: input.plugId,
					socketId: ctx.session.address
				}
			})

			if (!plug) throw new TRPCError({ code: "NOT_FOUND" })

			const intent = await createIntent({
				chainId: input.chainId,
				actions: plug.actions,
				frequency: input.frequency,
				startAt: input.startAt,
				endAt: input.endAt,
			})

			ctx.emitter.emit(subscriptions.execution.update, intent)

			// TODO: Add the intent id to the plugs intent id array

			return intent
		}),

	toggle: protectedProcedure.input(z.object({ id: z.string() })).mutation(async ({ input, ctx }) => {
		const intent = await toggleIntent(input)

		ctx.emitter.emit(subscriptions.execution.update, intent)

		return intent
	}),

	delete: protectedProcedure.input(z.object({ id: z.string() })).mutation(async ({ input, ctx }) => {
		const intent = await deleteIntent(input)

		ctx.emitter.emit(subscriptions.execution.delete, intent)

		// TODO: Remove the intent from the plugs intent id array

		return intent
	}),

	onActivity: subscription<Intent>("protected", subscriptions.execution.update),
	onDelete: subscription<Intent>("protected", subscriptions.execution.delete)
})
