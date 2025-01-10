import { TRPCError } from "@trpc/server"

import { z } from "zod"

import { Prisma, Simulation } from "@prisma/client"

import { Action } from "@/lib"

import { apiKeyProcedure, createTRPCRouter } from "../../trpc"

const MINUTE = 60 * 1000
const HOUR = 60 * MINUTE
const DAY = 24 * HOUR

const execution = Prisma.validator<Prisma.ExecutionDefaultArgs>()({
	include: {
		workflow: true,
		simulations: { orderBy: { createdAt: "desc" } }
	}
})
export type Execution = Prisma.ExecutionGetPayload<typeof execution>

const getNextSimulationAt = (
	execution: Execution,
	simulation: { id: string; status: string; error?: string; errors?: string[]; gasEstimate?: number }
) => {
	const now = new Date()
	const workflowFrequency = execution.workflow.frequency * MINUTE
	const executionFrequency = execution.frequency * DAY

	// Early exits
	if (execution.endAt && now >= execution.endAt) return null
	if (execution.frequency <= 0 && simulation.status === "success") return null

	// Compute next period end if available
	const nextPeriodEnd = execution.periodEndAt ? new Date(execution.periodEndAt.getTime() + executionFrequency) : null

	console.log("execution", execution)
	console.log("simulation", simulation)

	// Handle zero/negative frequency failure case
	if (execution.frequency <= 0) {
		const potentialNext = new Date(now.getTime() + workflowFrequency)

		if (!execution.endAt || potentialNext > execution.endAt)
			return {
				periodEndAt: execution.periodEndAt,
				nextSimulationAt: potentialNext
			}

		return null
	}

	// Handle success case with positive frequency
	if (simulation.status === "success") {
		if (!nextPeriodEnd) return null
		if (execution.endAt && nextPeriodEnd > execution.endAt)
			return {
				periodEndAt: execution.endAt,
				nextSimulationAt: null
			}

		return {
			periodEndAt: nextPeriodEnd,
			nextSimulationAt: nextPeriodEnd
		}
	}

	// Handle failure case with positive frequency
	const potentialNext = new Date(now.getTime() + workflowFrequency)

	// If we can retry within current period
	if (execution.periodEndAt && potentialNext < execution.periodEndAt) {
		return {
			periodEndAt: execution.periodEndAt,
			nextSimulationAt: potentialNext
		}
	}

	// If we need to move to next period
	if (nextPeriodEnd) {
		if (execution.endAt && nextPeriodEnd > execution.endAt) {
			return {
				periodEndAt: execution.endAt,
				nextSimulationAt: null
			}
		}

		return {
			periodEndAt: nextPeriodEnd,
			nextSimulationAt: new Date(nextPeriodEnd.getTime() + workflowFrequency)
		}
	}

	return null
}

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
