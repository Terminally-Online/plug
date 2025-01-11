const MINUTE = 60 * 1000
const HOUR = 60 * MINUTE
const DAY = 24 * HOUR

export const getNextSimulationAt = (
	execution: {
		frequency: number
		nextSimulationAt: Date | null
		endAt: Date | null
		periodEndAt: Date | null
		workflow: { frequency: number }
	},
	simulation: {
		status: string
	},
	now = new Date()
) => {
	const workflowFrequency = execution.workflow.frequency * MINUTE
	const executionFrequency = execution.frequency * DAY

	if (execution.endAt && now >= execution.endAt) return null
	if (execution.frequency === 0 && simulation.status === "success") return null

	if (execution.frequency === 0) {
		return {
			periodEndAt: execution.periodEndAt,
			nextSimulationAt: new Date(now.getTime() + workflowFrequency)
		}
	}

	if (execution.frequency > 0 && execution.periodEndAt) {
		if (simulation.status === "success") {
			const nextPeriodEnd = new Date(execution.periodEndAt.getTime() + executionFrequency)

			if (execution.endAt && nextPeriodEnd.getTime() > execution.endAt.getTime()) {
				return {
					periodEndAt: execution.endAt,
					nextSimulationAt: null
				}
			}

			return {
				periodEndAt: nextPeriodEnd,
				nextSimulationAt: execution.periodEndAt
			}
		}

		return {
			periodEndAt: execution.periodEndAt,
			nextSimulationAt: new Date(now.getTime() + workflowFrequency)
		}
	}

	return null
}
