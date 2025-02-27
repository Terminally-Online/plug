import { TRPCError } from "@trpc/server"

import { z } from "zod"

import { Plug } from "@prisma/client"

import { Action, createIntent, deleteIntent, getIntent, Intent, toggleIntent } from "@/lib"
import { createTRPCRouter, protectedProcedure } from "@/server/api/trpc"
import { subscription, subscriptions } from "@/server/subscription"

type IntentWithPlug = Intent & { plug: Plug }

export const activity = createTRPCRouter({
	get: protectedProcedure.query(async ({ ctx }) => {
		const intents = await getIntent({ address: ctx.session.address })

		const intentIds = intents.map(intent => intent.id)
		const plugs = await ctx.db.plug.findMany({
			where: { intentIds: { hasSome: intentIds } }
		})

		return intents
			.map(intent => {
				const plug = plugs.find(plug => plug.intentIds.includes(intent.id))
				if (!plug) return
				return { ...intent, plug }
			})
			.filter(intent => intent != undefined)
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
			const plug = await ctx.db.plug.findUniqueOrThrow({
				where: {
					id: input.plugId
				}
			})
			const intent = await createIntent({
				chainId: input.chainId,
				from: ctx.session.address,
				status: "active",
				inputs: JSON.parse(plug.actions),
				frequency: input.frequency,
				startAt: input.startAt.toISOString(),
				endAt: input.endAt?.toISOString()
			})

			const updated = {
				...intent,
				plug: await ctx.db.plug.update({
					where: { id: input.plugId, socketId: ctx.session.address },
					data: {
						intentIds: {
							push: intent.id
						}
					}
				})
			}

			ctx.emitter.emit(subscriptions.execution.update, updated)

			return updated
		}),

	toggle: protectedProcedure.input(z.object({ id: z.string() })).mutation(async ({ input, ctx }) => {
		const plug = await ctx.db.plug.findFirstOrThrow({
			where: {
				socketId: ctx.session.address,
				intentIds: { has: input.id }
			}
		})

		if (!plug) throw new TRPCError({ code: "NOT_FOUND" })

		const intent = { ...(await toggleIntent(input)), plug }

		// ctx.emitter.emit(subscriptions.execution.update, intent)

		// return intent
	}),

	delete: protectedProcedure.input(z.object({ id: z.string() })).mutation(async ({ input, ctx }) => {
		const intent = await deleteIntent(input)

		ctx.emitter.emit(subscriptions.execution.delete, intent)

		return intent
	}),

	onActivity: subscription<IntentWithPlug>("protected", subscriptions.execution.update),
	onDelete: subscription<IntentWithPlug>("protected", subscriptions.execution.delete)
})
