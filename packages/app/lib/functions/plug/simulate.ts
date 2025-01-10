import { Prisma } from "@prisma/client"

const MINUTE = 60 * 1000
const HOUR = 60 * MINUTE
const DAY = 24 * HOUR

const execution = Prisma.validator<Prisma.ExecutionDefaultArgs>()({
	include: {
		workflow: true,
		simulations: { orderBy: { createdAt: "desc" } }
	}
})
type Execution = Prisma.ExecutionGetPayload<typeof execution>
type Simulation = { id: string; status: string; error?: string; errors?: string[]; gasEstimate?: number }

export const getNextSimulationAt = (execution: Execution, simulation: Simulation) => {
	const now = new Date()
	const workflowFrequency = execution.workflow.frequency * MINUTE
	const executionFrequency = execution.frequency * DAY

	// If we have passed the end date then stop everything.
	if (execution.endAt && now >= execution.endAt) return null
	// If we have a success and the execution was "once" then stop everything.
	if (execution.frequency <= 0 && simulation.status === "success") return null

	// Since we know we haven't reached the end date, calculate the adjusted period end date
	// if we were to move it to the next period.
	const nextPeriodEnd = execution.periodEndAt ? new Date(execution.periodEndAt.getTime() + executionFrequency) : null

	// Handle zero/negative frequency failure case
	if (execution.frequency <= 0) {
		const potentialNext = new Date(now.getTime() + workflowFrequency)

		if (!execution.endAt || potentialNext <= execution.endAt)
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
