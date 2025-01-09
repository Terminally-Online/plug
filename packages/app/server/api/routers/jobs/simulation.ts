import { TRPCError } from "@trpc/server"

import { z } from "zod"

import { Action } from "@/lib"

import { apiKeyProcedure, createTRPCRouter } from "../../trpc"

const DEFAULT_SIMULATION_FREQUENCY = 1 * 60 * 1000

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

			// NOTE: We have to do a little shape transformation as our solver api expects a specific shape.
			const parsedExecutions = executions.map(queuedWorkflow => {
				const inputs = JSON.parse(queuedWorkflow.actions as string).map((action: Action) => ({
					protocol: action.protocol,
					action: action.action,
					...Object.entries(action.values ?? []).reduce(
						(acc, [_, value]) => {
							if (!value || !value.name) return acc

							acc[value.name] = value.value

							return acc
						},
						{} as Record<string, string>
					)
				}))

				return {
					id: queuedWorkflow.id,
					chainId: 1,
					from: queuedWorkflow.workflow.socket.socketAddress,
					inputs
				}
			})

			// NOTE: We update these executions in the database when they are called upon by
			// the solver so that we do not process the same simulation twice. We do not create a simulation
			// record in the database until we get back the response from the solver.
			await tx.execution.updateMany({
				where: {
					id: {
						in: executions.map(execution => execution.id)
					}
				},
				data: { nextSimulationAt: new Date(now.getTime() + DEFAULT_SIMULATION_FREQUENCY) }
			})

			return parsedExecutions
		})
	}),

	simulated: apiKeyProcedure
		.input(
			z.array(
				z.object({
					id: z.string(),
					status: z.string(),
					error: z.string().optional(),
					gasEstimate: z.number().optional()
				})
			)
		)
		.mutation(async ({ input, ctx }) => {
			return await ctx.db.$transaction(async tx => {
				return await Promise.all(
					input.map(async simulation => {
						const { id } = await tx.simulation.create({
							data: {
								status: simulation.status,
								executionId: simulation.id,
								error: simulation.error,
								gasEstimate: simulation.gasEstimate
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
