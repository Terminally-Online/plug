import { TRPCError } from "@trpc/server"

import { z } from "zod"

import { Action, getNextSimulationAt } from "@/lib"
import { subscriptions } from "@/server/subscription"

import { apiKeyProcedure, createTRPCRouter } from "../../trpc"

export const simulation = createTRPCRouter({
	simulateNext: apiKeyProcedure.input(z.object({ count: z.number() }).nullish()).mutation(async ({ input, ctx }) => {
		return await ctx.db.$transaction(async tx => {
			const executions = await tx.intent.findMany({
				where: {
					nextSimulationAt: { lte: new Date() },
					status: "active"
				},
				take: input?.count ?? 100,
				select: {
					id: true,
					chainId: true,
					actions: true,
					plug: {
						select: {
							socket: { select: { id: true, socketAddress: true } }
						}
					}
				}
			})

			await tx.intent.updateMany({
				where: { id: { in: executions.map(e => e.id) } },
				data: { status: "processing" }
			})

			executions.forEach(execution => {
				ctx.emitter.emit(subscriptions.execution.update, execution)
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
					chainId: queuedWorkflow.chainId,
					from: queuedWorkflow.plug.socket.socketAddress,
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
					success: z.boolean(),
					gasUsed: z.number().optional(),
					errorMessage: z.string().optional()
				})
			)
		)
		.mutation(async ({ input, ctx }) => {
			return await ctx.db.$transaction(async tx => {
				console.log("input", input)

				return await Promise.all(
					input.map(async simulation => {
						const execution = await tx.intent.findUnique({
							where: { id: simulation.id },
							include: { plug: true, runs: true }
						})

						if (!execution) throw new TRPCError({ code: "NOT_FOUND" })
						if (execution.status !== "processing") throw new TRPCError({ code: "BAD_REQUEST" })

						const status = simulation.success ? "success" : "failure"
						const nextSimulation = getNextSimulationAt(execution, { status })

						await tx.intent.update({
							where: { id: simulation.id },
							data: {
								status: nextSimulation?.nextSimulationAt ? "active" : "completed",
								nextSimulationAt: nextSimulation?.nextSimulationAt ?? null,
								periodEndAt: nextSimulation?.periodEndAt ?? execution.periodEndAt
							}
						})

						const { id } = await tx.run.create({
							data: {
								status,
								intentId: simulation.id,
								error: simulation.errorMessage,
								// errors: simulation.errors,
								gasEstimate: simulation.gasUsed
							},
							select: { id: true }
						})

						ctx.emitter.emit(subscriptions.execution.update, execution)

						return id
					})
				)
			})
		})
})
