import { TRPCError } from "@trpc/server"

import { z } from "zod"

import { Prisma, Simulation } from "@prisma/client"

import { Action, getNextSimulationAt } from "@/lib"

import { apiKeyProcedure, createTRPCRouter } from "../../trpc"

export const simulation = createTRPCRouter({
	simulateNext: apiKeyProcedure.input(z.object({ count: z.number() }).nullish()).mutation(async ({ input, ctx }) => {
		return await ctx.db.$transaction(async tx => {
			const executions = await tx.execution.findMany({
				where: {
					nextSimulationAt: { lte: new Date() },
					status: "active"
				},
				take: input?.count ?? 100,
				select: {
					id: true,
					actions: true,
					workflow: {
						select: {
							socket: { select: { id: true, socketAddress: true } }
						}
					}
				}
			})

			await tx.execution.updateMany({
				where: { id: { in: executions.map(e => e.id) } },
				data: { status: "processing" }
			})

			return executions.map(queuedWorkflow => {
				const inputs = JSON.parse(queuedWorkflow.actions as string).map((action: Action) => ({
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
				}))

				return {
					id: queuedWorkflow.id,
					chainId: 1,
					from: queuedWorkflow.workflow.socket.socketAddress,
					inputs
				}
			})
		})
	}),

	simulated: apiKeyProcedure
		.input(
			z.array(
				z.object({
					id: z.string(),
					status: z.string(),
					error: z.string().optional(),
					errors: z.array(z.string()).optional(),
					gasEstimate: z.number().optional()
				})
			)
		)
		.mutation(async ({ input, ctx }) => {
			return await ctx.db.$transaction(async tx => {
				return await Promise.all(
					input.map(async simulation => {
						const execution = await tx.execution.findUnique({
							where: { id: simulation.id },
							include: { workflow: true, simulations: true }
						})

						if (!execution) throw new TRPCError({ code: "NOT_FOUND" })
						if (execution.status !== "processing") throw new TRPCError({ code: "BAD_REQUEST" })

						const nextSimulation = getNextSimulationAt(execution, simulation)

						await tx.execution.update({
							where: { id: simulation.id },
							data: {
								status: nextSimulation?.nextSimulationAt ? "active" : "completed",
								nextSimulationAt: nextSimulation?.nextSimulationAt ?? null,
								periodEndAt: nextSimulation?.periodEndAt ?? execution.periodEndAt
							}
						})

						const { id } = await tx.simulation.create({
							data: {
								status: simulation.status,
								executionId: simulation.id,
								error: simulation.error,
								errors: simulation.errors,
								gasEstimate: simulation.gasEstimate
							},
							select: { id: true }
						})

						return id
					})
				)
			})
		})
})
