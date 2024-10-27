import { anonymousProtectedProcedure, createTRPCRouter, protectedProcedure } from "@/server/api/trpc"
import { TRPCError } from "@trpc/server"

import { z } from "zod"

import { Prisma } from "@prisma/client"

import { categories } from "@/lib"

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

	getQueued: protectedProcedure.query(async ({ ctx }) => {
		const queuedWorkflows = await ctx.db.queuedWorkflow.findMany({
			where: { socketId: ctx.session.address },
			orderBy: { startAt: "desc" },
			include: { workflow: true }
		})

		return queuedWorkflows.map(qw => ({
			id: qw.id,
			text: qw.workflow.name || "Unnamed Workflow",
			status: "pending",
			time: qw.startAt.toISOString(),
			color: "blue"
		}))
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

			if (!workflow) {
				throw new TRPCError({
					code: "NOT_FOUND",
					message: "Workflow not found"
				})
			}

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

			const activityEvent = {
				id: queuedWorkflow.id,
				text: workflow.name || "Unnamed Workflow",
				status: "pending",
				time: queuedWorkflow.startAt.toISOString(),
				color: "blue"
			}

			ctx.emitter.emit(events.queue, activityEvent)

			return queuedWorkflow
		})
})
