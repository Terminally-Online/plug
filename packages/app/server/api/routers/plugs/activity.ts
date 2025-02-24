import { TRPCError } from "@trpc/server"

import { z } from "zod"

import { Prisma } from "@prisma/client"

import { createIntent } from "@/lib"
import { createTRPCRouter, protectedProcedure } from "@/server/api/trpc"
import { subscription, subscriptions } from "@/server/subscription"

const execution = Prisma.validator<Prisma.IntentDefaultArgs>()({
	include: {
		plug: true,
		runs: { orderBy: { createdAt: "desc" } }
	}
})
export type Execution = Prisma.IntentGetPayload<typeof execution>

export const activity = createTRPCRouter({
	get: protectedProcedure.query(async ({ ctx }) => {
		return await ctx.db.intent.findMany({
			where: { plug: { socketId: ctx.session.address } },
			orderBy: { createdAt: "desc" },
			include: { plug: true, runs: { orderBy: { createdAt: "desc" } } }
		})
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
		const execution = await ctx.db.intent.findUnique({
			where: { id: input.id },
			include: { plug: true }
		})

		if (!execution) throw new TRPCError({ code: "NOT_FOUND" })

		const toggled = await ctx.db.intent.update({
			where: { id: input.id },
			data: { status: execution.status.trim() !== "active" ? "active" : "paused" },
			include: {
				plug: true,
				runs: { orderBy: { createdAt: "desc" } }
			}
		})

		ctx.emitter.emit(subscriptions.execution.update, toggled)

		return toggled
	}),

	delete: protectedProcedure.input(z.object({ id: z.string() })).mutation(async ({ input, ctx }) => {
		const execution = await ctx.db.intent.delete({
			where: { id: input.id }
		})

		ctx.emitter.emit(subscriptions.execution.delete, execution)

		return execution
	}),

	onActivity: subscription<Execution>("protected", subscriptions.execution.update),
	onDelete: subscription<Execution>("protected", subscriptions.execution.delete)
})
