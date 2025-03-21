import { TRPCError } from "@trpc/server"

import { z } from "zod"

import { Plug } from "@prisma/client"

import { createIntent, deleteIntent, getIntent, IntentResponseIntent, toggleIntent, toggleIntentStatus } from "@/lib"
import { createTRPCRouter, protectedProcedure } from "@/server/api/trpc"
import { subscription, subscriptions } from "@/server/subscription"

type IntentWithPlug = IntentResponseIntent & { plug: Plug }

export const activity = createTRPCRouter({
	get: protectedProcedure.query(async ({ ctx }) => {
		const socket = await ctx.db.socket.findFirstOrThrow({ where: { id: ctx.session.address } })
		const intents = await getIntent({ addresses: [ctx.session.address, socket.socketAddress] })
		console.log('intents', intents)
		const intentIds = intents.map(intent => intent.id)
		console.log('intent ids', intentIds)
		const plugs = await ctx.db.plug.findMany({
			where: { intentIds: { hasSome: intentIds } }
		})
		console.log('plugs', plugs)

		return intents
			.map(intent => {
				const plug = plugs.find(plug => plug.intentIds.includes(intent.id))
				return { ...intent, plug }
			})
	}),

	queue: protectedProcedure
		.input(
			z.object({
				plugId: z.string(),
				chainId: z.number(),
				frequency: z.number(),
				startAt: z.coerce.date(),
				endAt: z.coerce.date().optional(),
				socket: z.boolean().optional()
			})
		)
		.mutation(async ({ input, ctx }) => {
			const { socket, ...plug } = await ctx.db.plug.findUniqueOrThrow({
				where: {
					id: input.plugId
				},
				include: { socket: true }
			})
			const intent = await createIntent({
				chainId: input.chainId,
				from: input.socket ? socket.socketAddress : ctx.session.address,
				status: "active",
				inputs: JSON.parse(plug.actions),
				frequency: input.frequency,
				startAt: input.startAt.toISOString(),
				endAt: input.endAt?.toISOString()
			})
			const updated = await ctx.db.plug.update({
				where: { id: input.plugId, socketId: ctx.session.address },
				data: {
					intentIds: {
						push: intent.id
					}
				}
			})
			const spread = { ...intent, plug: updated }

			ctx.emitter.emit(subscriptions.execution.update, spread)

			return spread
		}),

	toggleSaved: protectedProcedure.input(z.object({ id: z.string() })).mutation(async ({ input, ctx }) => {
		const plug = await ctx.db.plug.findFirstOrThrow({
			where: {
				socketId: ctx.session.address,
				intentIds: { has: input.id }
			}
		})

		if (!plug) throw new TRPCError({ code: "NOT_FOUND" })

		const toggled = await toggleIntent(input)
		const intent = { ...toggled, plug }

		ctx.emitter.emit(subscriptions.execution.update, intent)

		return intent
	}),

	toggleStatus: protectedProcedure.input(z.object({ id: z.string() })).mutation(async ({ input, ctx }) => {
		const plug = await ctx.db.plug.findFirstOrThrow({
			where: {
				socketId: ctx.session.address,
				intentIds: { has: input.id }
			}
		})

		if (!plug) throw new TRPCError({ code: "NOT_FOUND" })

		const intent = { ...(await toggleIntentStatus(input)), plug }

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
