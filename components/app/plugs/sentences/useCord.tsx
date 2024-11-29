import { useCallback, useMemo, useState } from "react"

import {
	CordState,
	createInitialState,
	parseCordSentence,
	resolveSentence,
	setValue,
	shouldRenderInput,
	UseCordReturn
} from "@terminallyonline/cord"

const initialState: CordState = {
	values: createInitialState(),
	parsed: null,
	resolvedSentence: null,
	error: null,
	isDirty: false,
	validationErrors: new Map()
}

export const useCord = (sentence: string): UseCordReturn => {
	const [state, setState] = useState<CordState>(initialState)

	const parsed = useMemo(() => {
		const result = parseCordSentence(sentence)
		if (!result.success) {
			setState(prev => ({
				...prev,
				error: { type: "parse", message: result.error },
				parsed: null
			}))
			return null
		}
		setState(prev => ({ ...prev, error: null }))
		return result.value
	}, [sentence])

	const filteredInputs = useMemo(() => {
		if (!parsed) return []
		return parsed.inputs.filter(input =>
			shouldRenderInput(input.type, parsed.inputs, index => state.values.get(index))
		)
	}, [parsed, state.values])

	const parsedWithFilteredInputs = useMemo(() => {
		if (!parsed) return null
		return {
			...parsed,
			inputs: filteredInputs
		}
	}, [parsed, filteredInputs])

	const resolvedSentence = useMemo(() => {
		if (!parsed) return null

		const allInputsHaveValues = parsed.inputs.every(input => state.values.has(input.index))
		if (!allInputsHaveValues) return null

		const result = resolveSentence(parsed, state.values)
		if (!result.success) {
			setState(prev => ({
				...prev,
				error: { type: "resolution", message: result.error }
			}))
			return null
		}
		return result.value
	}, [parsed, state.values])

	const actions = {
		setValue: useCallback(
			(index: number, value: string) => {
				if (!parsed) return

				const result = setValue({
					parsedSentence: parsed,
					currentValues: state.values,
					index,
					value
				})

				setState(prev => ({
					...prev,
					values: result.value,
					isDirty: true,
					validationErrors: result.error
						? new Map(prev.validationErrors).set(index, {
								type: "validation",
								message: result.error
							})
						: new Map([...prev.validationErrors].filter(([k]) => k !== index))
				}))
			},
			[parsed, state.values]
		),

		reset: useCallback(() => {
			setState(initialState)
		}, []),

		clear: useCallback((index: number) => {
			setState(prev => {
				const newValues = new Map(prev.values)
				newValues.delete(index)

				return {
					...prev,
					values: newValues,
					validationErrors: new Map([...prev.validationErrors].filter(([k]) => k !== index)),
					isDirty: true
				}
			})
		}, []),

		clearAll: useCallback(() => {
			setState(prev => ({
				...prev,
				values: createInitialState(),
				validationErrors: new Map(),
				isDirty: true
			}))
		}, [])
	}

	const helpers = {
		getInputValue: useCallback((index: number) => state.values.get(index), [state.values]),

		getInputError: useCallback((index: number) => state.validationErrors.get(index), [state.validationErrors]),

		getDependentInputs: useCallback(
			(index: number) => {
				if (!parsed) return []
				return parsed.inputs.filter(input => input.dependentOn === index)
			},
			[parsed]
		),

		hasDependency: useCallback(
			(index: number) => {
				if (!parsed) return false
				return parsed.inputs.some(input => input.dependentOn === index)
			},
			[parsed]
		),

		isComplete: useMemo(
			() => parsed?.inputs.every(input => state.values.has(input.index)) ?? false,
			[parsed, state.values]
		),

		isValid: useMemo(() => state.validationErrors.size === 0, [state.validationErrors])
	}

	return {
		state: {
			...state,
			parsed: parsedWithFilteredInputs,
			resolvedSentence
		},
		actions,
		helpers
	}
}
