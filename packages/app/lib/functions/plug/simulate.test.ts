import { getNextSimulationAt } from "./simulate"

const NOW = new Date()
const MINUTE = 60 * 1000
const HOUR = 60 * MINUTE
const DAY = 24 * HOUR

const WORKFLOW = {
	workflowId: "test",
	workflow: {
		createdAt: NOW,
		name: "test",
		id: "test",
		updatedAt: NOW,
		actions: "",
		frequency: 10,
		isCurated: false,
		isPrivate: false,
		color: "",
		tags: [],
		workflowForkedId: null,
		socketId: ""
	}
}

const SUCCESSFUL_SIMULATION = { status: "success" }
const FAILED_SIMULATION = { status: "failed" }

const getExecution = (frequency: number, duration?: number) => ({
	id: "test",
	status: "active",
	simulations: [],
	actions: "[]",
	...WORKFLOW,
	frequency,
	nextSimulationAt: NOW,
	endAt: duration ? new Date(NOW.getTime() + DAY * duration) : null,
	periodEndAt: new Date(NOW.getTime() + frequency * DAY)
})

describe("getNextSimulationAt", () => {
	describe("zero frequency (one-time execution)", () => {
		const singleUseExecution = getExecution(0, 7)

		it("returns null if frequency is zero and experiences success", () => {
			const result = getNextSimulationAt(singleUseExecution, SUCCESSFUL_SIMULATION)
			expect(result).toBeNull()
		})

		it("returns null if end date is in the past", () => {
			const pastExecution = {
				...singleUseExecution,
				endAt: new Date(NOW.getTime() - DAY)
			}
			const result = getNextSimulationAt(pastExecution, SUCCESSFUL_SIMULATION)
			expect(result).toBeNull()
		})

		it("returns next simulation time if zero frequency experiences failure", () => {
			const result = getNextSimulationAt(singleUseExecution, FAILED_SIMULATION, NOW)
			const expectedNext = new Date(NOW.getTime() + WORKFLOW.workflow.frequency * MINUTE)
			expect(result?.nextSimulationAt).toEqual(expectedNext)
			expect(result?.periodEndAt).toEqual(singleUseExecution.periodEndAt)
		})
	})

	describe("daily frequency (recurring execution)", () => {
		const dailyExecution = getExecution(1, 30)

		describe("success cases", () => {
			it("moves to next period and sets next simulation to period start", () => {
				const result = getNextSimulationAt(dailyExecution, SUCCESSFUL_SIMULATION, NOW)

				// Next period should be current period + 1 day
				const expectedPeriodEnd = new Date(dailyExecution.periodEndAt.getTime() + DAY)
				expect(result?.periodEndAt).toEqual(expectedPeriodEnd)

				// Next simulation should be at current period end (start of next period)
				expect(result?.nextSimulationAt).toEqual(dailyExecution.periodEndAt)
			})

			it("returns null if next period would exceed end date", () => {
				const nearEndExecution = {
					...dailyExecution,
					endAt: new Date(dailyExecution.periodEndAt.getTime() + HOUR)
				}
				const result = getNextSimulationAt(nearEndExecution, SUCCESSFUL_SIMULATION, NOW)
				expect(result?.nextSimulationAt).toBeNull()
				expect(result?.periodEndAt).toEqual(nearEndExecution.endAt)
			})
		})

		describe("failure cases", () => {
			it("retries within same period using workflow frequency", () => {
				const result = getNextSimulationAt(dailyExecution, FAILED_SIMULATION, NOW)
				expect(result?.nextSimulationAt).toEqual(new Date(NOW.getTime() + WORKFLOW.workflow.frequency * MINUTE))
				expect(result?.periodEndAt).toEqual(dailyExecution.periodEndAt)
			})
		})
	})

	describe("weekly frequency (recurring execution)", () => {
		const weeklyExecution = getExecution(7, 30)

		describe("success cases", () => {
			it("moves to next period and sets next simulation to period start", () => {
				const result = getNextSimulationAt(weeklyExecution, SUCCESSFUL_SIMULATION, NOW)

				// Next period should be current period + 7 days
				const expectedPeriodEnd = new Date(weeklyExecution.periodEndAt.getTime() + 7 * DAY)
				expect(result?.periodEndAt).toEqual(expectedPeriodEnd)

				// Next simulation should be at current period end (start of next period)
				expect(result?.nextSimulationAt).toEqual(weeklyExecution.periodEndAt)
			})

			it("returns null if next period would exceed end date", () => {
				const nearEndExecution = {
					...weeklyExecution,
					endAt: new Date(weeklyExecution.periodEndAt.getTime() + DAY)
				}
				const result = getNextSimulationAt(nearEndExecution, SUCCESSFUL_SIMULATION, NOW)
				expect(result?.nextSimulationAt).toBeNull()
				expect(result?.periodEndAt).toEqual(nearEndExecution.endAt)
			})
		})

		describe("failure cases", () => {
			it("retries within same period using workflow frequency", () => {
				const result = getNextSimulationAt(weeklyExecution, FAILED_SIMULATION, NOW)
				expect(result?.nextSimulationAt).toEqual(new Date(NOW.getTime() + WORKFLOW.workflow.frequency * MINUTE))
				expect(result?.periodEndAt).toEqual(weeklyExecution.periodEndAt)
			})
		})
	})

	describe("monthly frequency (recurring execution)", () => {
		const monthlyExecution = getExecution(30, 90)

		describe("success cases", () => {
			it("moves to next period and sets next simulation to period start", () => {
				const result = getNextSimulationAt(monthlyExecution, SUCCESSFUL_SIMULATION, NOW)

				// Next period should be current period + 30 days
				const expectedPeriodEnd = new Date(monthlyExecution.periodEndAt.getTime() + 30 * DAY)
				expect(result?.periodEndAt).toEqual(expectedPeriodEnd)

				// Next simulation should be at current period end (start of next period)
				expect(result?.nextSimulationAt).toEqual(monthlyExecution.periodEndAt)
			})

			it("returns null if next period would exceed end date", () => {
				const nearEndExecution = {
					...monthlyExecution,
					endAt: new Date(monthlyExecution.periodEndAt.getTime() + DAY)
				}
				const result = getNextSimulationAt(nearEndExecution, SUCCESSFUL_SIMULATION, NOW)
				expect(result?.nextSimulationAt).toBeNull()
				expect(result?.periodEndAt).toEqual(nearEndExecution.endAt)
			})
		})

		describe("failure cases", () => {
			it("retries within same period using workflow frequency", () => {
				const result = getNextSimulationAt(monthlyExecution, FAILED_SIMULATION, NOW)
				expect(result?.nextSimulationAt).toEqual(new Date(NOW.getTime() + WORKFLOW.workflow.frequency * MINUTE))
				expect(result?.periodEndAt).toEqual(monthlyExecution.periodEndAt)
			})
		})
	})

	describe("yearly frequency (recurring execution)", () => {
		const yearlyExecution = getExecution(365, 730) // 2 year duration

		describe("success cases", () => {
			it("moves to next period and sets next simulation to period start", () => {
				const result = getNextSimulationAt(yearlyExecution, SUCCESSFUL_SIMULATION, NOW)

				// Next period should be current period + 365 days
				const expectedPeriodEnd = new Date(yearlyExecution.periodEndAt.getTime() + 365 * DAY)
				expect(result?.periodEndAt).toEqual(expectedPeriodEnd)

				// Next simulation should be at current period end (start of next period)
				expect(result?.nextSimulationAt).toEqual(yearlyExecution.periodEndAt)
			})

			it("returns null if next period would exceed end date", () => {
				const nearEndExecution = {
					...yearlyExecution,
					endAt: new Date(yearlyExecution.periodEndAt.getTime() + DAY)
				}
				const result = getNextSimulationAt(nearEndExecution, SUCCESSFUL_SIMULATION, NOW)
				expect(result?.nextSimulationAt).toBeNull()
				expect(result?.periodEndAt).toEqual(nearEndExecution.endAt)
			})
		})

		describe("failure cases", () => {
			it("retries within same period using workflow frequency", () => {
				const result = getNextSimulationAt(yearlyExecution, FAILED_SIMULATION, NOW)
				expect(result?.nextSimulationAt).toEqual(new Date(NOW.getTime() + WORKFLOW.workflow.frequency * MINUTE))
				expect(result?.periodEndAt).toEqual(yearlyExecution.periodEndAt)
			})
		})
	})
})
