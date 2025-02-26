import { z } from "zod"

import { Plug } from "@prisma/client"

import { createIntent, deleteIntent, getIntent, toggleIntent, Intent, Action } from "@/lib"
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

		return intents.map(intent => ({
			...intent,
			plug: plugs.find(plug => plug.intentIds.includes(intent.id))!!
		}))
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
					id: input.plugId,
				}
			})

			console.log('got plug', plug)

			const actions = JSON.parse(plug.actions as string).map((action: Action) => ({
				protocol: action.protocol,
				action: action.action,
				...Object.entries(action.values ?? []).reduce(
					(acc, [_, value]) => {
						if (!value?.name) return acc
						acc[value.name] = value.value
						return acc
					},
					{} as Record<string, string>
				)
			})) as Array<Record<string, string>>

			const intent = await createIntent({
				chainId: input.chainId,
				from: ctx.session.address,
				status: "active",
				actions,
				frequency: input.frequency,
				startAt: input.startAt,
				endAt: input.endAt,
			})

			console.log('created intent')

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

			console.log('updated', updated)

			ctx.emitter.emit(subscriptions.execution.update, updated)

			return updated
		}),

	toggle: protectedProcedure.input(z.object({ id: z.string() })).mutation(async ({ input, ctx }) => {
		const intent = await toggleIntent(input)

		ctx.emitter.emit(subscriptions.execution.update, intent)

		return intent
	}),

	delete: protectedProcedure.input(z.object({ id: z.string() })).mutation(async ({ input, ctx }) => {
		const intent = await deleteIntent(input)

		ctx.emitter.emit(subscriptions.execution.delete, intent)

		return intent
	}),

	onActivity: subscription<IntentWithPlug>("protected", subscriptions.execution.update),
	onDelete: subscription<IntentWithPlug>("protected", subscriptions.execution.delete)
})
