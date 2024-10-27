import { anonymousProtectedProcedure, createTRPCRouter, protectedProcedure } from "@/server/api/trpc"
import { TRPCError } from "@trpc/server"
import { observable } from "@trpc/server/observable"

import { z } from "zod"

import { Prisma } from "@prisma/client"

import { categories } from "@/lib"

const queuedWorkflow = Prisma.validator<Prisma.QueuedWorkflowDefaultArgs>()({
	include: {
		workflow: true
	}
})
export type QueuedWorkflow = Prisma.QueuedWorkflowGetPayload<typeof queuedWorkflow>

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

const subscription = (event: string) =>
	anonymousProtectedProcedure.subscription(async ({ ctx }) => {
		return observable<QueuedWorkflow>(emit => {
			const handleSubscription = (data: QueuedWorkflow) => {
				emit.next(data)
			}

			ctx.emitter.on(event, handleSubscription)

			return () => ctx.emitter.off(event, handleSubscription)
		})
	})

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
		}),

	activity: protectedProcedure.query(async ({ ctx }) => {
		return await ctx.db.queuedWorkflow.findMany({
			where: { socketId: ctx.session.address },
			orderBy: { createdAt: "desc" },
			include: { workflow: true }
		})
	}),

	queue: anonymousProtectedProcedure
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

			const queuedWorkflow = await ctx.db.queuedWorkflow.create({
				data: {
					workflowId: input.workflowId,
					socketId: ctx.session.address,
					startAt: input.startAt,
					endAt: input.endAt,
					frequency: input.frequency,
					nextSimulationAt: input.startAt
				},
				include: {
					workflow: true
				}
			})

			ctx.emitter.emit(events.queue, queuedWorkflow)

			return queuedWorkflow
		}),

	onActivity: subscription(events.queue)
})
