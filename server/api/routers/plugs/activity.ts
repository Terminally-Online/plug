import { TRPCError } from "@trpc/server"

import { z } from "zod"

import { Prisma } from "@prisma/client"

import { createTRPCRouter, protectedProcedure } from "@/server/api/trpc"
import { subscription, subscriptions } from "@/server/subscription"

const execution = Prisma.validator<Prisma.ExecutionDefaultArgs>()({
	include: {
		workflow: true
	}
})
export type Execution = Prisma.ExecutionGetPayload<typeof execution>

export const activity = createTRPCRouter({
	get: protectedProcedure.query(async ({ ctx }) => {
		return await ctx.db.execution.findMany({
			where: { workflow: { socketId: ctx.session.address } },
			orderBy: { createdAt: "desc" },
			include: { workflow: true }
		})
	}),

	queue: protectedProcedure
		.input(
			z.object({
				workflowId: z.string(),
				startAt: z.date(),
				endAt: z.date().optional(),
				frequency: z.number().optional()
			})
		)
		.mutation(async ({ input, ctx }) => {
			const workflow = await ctx.db.workflow.findUnique({
				where: {
					id: input.workflowId,
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
				const execution = await ctx.db.execution.create({
					data: {
						workflowId: input.workflowId,
						actions: workflow.actions,
						startAt: input.startAt,
						endAt: input.endAt,
						frequency: input.frequency,
						nextSimulationAt: input.startAt
					},
					include: {
						workflow: true
					}
				})

				ctx.emitter.emit(subscriptions.plugs.queue, execution)

				return execution
			} catch (error) {
				throw new TRPCError({ code: "BAD_REQUEST" })
			}
		}),

	onActivity: subscription<Execution>("protected", subscriptions.plugs.queue)
})
