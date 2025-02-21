import { TRPCError } from "@trpc/server"

import { z } from "zod"

import { Prisma } from "@prisma/client"

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
				startAt: z.date(),
				endAt: z.date().optional()
			})
		)
		.mutation(async ({ input, ctx }) => {
			const workflow = await ctx.db.plug.findUnique({
				where: {
					id: input.plugId,
					socketId: ctx.session.address
				}
			})

			if (!workflow) throw new TRPCError({ code: "NOT_FOUND" })

			// NOTE: We have a try/catch here so that if someone posts in invalid JSON, we don't
			// crash the server and instead return the proper 400 (Bad Request) error. We do not
			// need to utilize the parsed JSON here because we are only using it to validate the
			// JSON string will be fine when sent to the Solver backend.
			try {
				JSON.parse(workflow.actions)
				const execution = await ctx.db.intent.create({
					data: {
						plugId: input.plugId,
						chainId: input.chainId,
						actions: workflow.actions,

						frequency: input.frequency,
						startAt: input.startAt,
						endAt: input.endAt,
						periodEndAt: input.frequency
							? new Date(input.startAt.getTime() + input.frequency * 60 * 1000 * 60 * 24)
							: null,
						nextSimulationAt: input.startAt
					},
					include: {
						plug: true,
						runs: { orderBy: { createdAt: "desc" } }
					}
				})

				ctx.emitter.emit(subscriptions.execution.update, execution)

				return execution
			} catch (error) {
				throw new TRPCError({ code: "BAD_REQUEST" })
			}
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
