import { TRPCError } from "@trpc/server"

import { z } from "zod"

import { subscriptions } from "@/server/subscription"

import { apiKeyProcedure, createTRPCRouter } from "../../trpc"

export const simulation = createTRPCRouter({
	simulate: apiKeyProcedure.input(z.object({ id: z.string() })).mutation(async ({ input, ctx }) => {
		return await ctx.db.$transaction(async tx => {
			const execution = await tx.execution.findUnique({
				where: {
					id: input.id
				},
				include: {
					workflow: true
				}
			})

			if (execution === null) throw new TRPCError({ code: "NOT_FOUND" })

			const now = new Date()
			const nextSimulationAt = new Date(now.getTime() + execution.frequency * 60 * 1000)

			return await tx.execution.update({
				where: {
					id: execution.id
				},
				data: {
					nextSimulationAt
				}
			})
		})
	}),

	simulateNext: apiKeyProcedure.input(z.object({ count: z.number() }).nullish()).mutation(async ({ input, ctx }) => {
		const now = new Date()

		return await ctx.db.$transaction(async tx => {
			const executions = await tx.execution.findMany({
				where: {
					nextSimulationAt: {
						lte: now
					},
					status: "active"
				},
				take: input?.count ?? 100,
				select: {
					id: true,
					actions: true,
					workflow: { select: { socket: { select: { id: true, socketAddress: true } } } }
				}
			})

			const parsedExecutions = executions.map(queuedWorkflow => ({
				...queuedWorkflow,
				actions: JSON.parse(queuedWorkflow.actions as string)
			}))

			// NOTE: We update these executions in the database when they are called upon by
			// the solver so that we do not process the same simulation twice. We do not create a simulation
			// record in the database until we get back the response from the solver.
			// await Promise.all(
			// 	executions.map(simulatedExecution =>
			// 		tx.execution.update({
			// 			where: { id: simulatedExecution.id },
			// 			data: {
			// 				nextSimulationAt: new Date(now.getTime() + simulatedExecution.frequency * 60 * 1000)
			// 			}
			// 		})
			// 	)
			// )

			return parsedExecutions
		})
	}),

	simulated: apiKeyProcedure
		.input(z.array(z.object({ id: z.string(), status: z.string() })))
		.mutation(async ({ input, ctx }) => {
			return await ctx.db.$transaction(async tx => {
				return await Promise.all(
					input.map(async simulation => {
						const { id } = await tx.simulation.create({
							data: {
								status: simulation.status,
								executionId: simulation.id
							},
							select: {
								id: true
							}
						})

						return id
					})
				)
			})
		})
})
