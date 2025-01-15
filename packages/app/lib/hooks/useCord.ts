import { useCallback, useMemo, useState } from "react"

import {
	CordState,
	createInitialState,
	InputState,
	parseCordSentence,
	resolveSentence,
	setValue,
	shouldRenderInput,
	UseCordReturn
} from "@terminallyonline/cord"

const createStateFromValues = (values: Record<string, string | undefined>) => {
	const state = new Map<number, InputState>()
	Object.entries(values).forEach(([key, value]) => {
		if (value !== undefined) {
			state.set(Number(key), { value })
		}
	})
	return state
}

const initialState: CordState = {
	values: createInitialState(),
	parsed: null,
	resolvedSentence: null,
	error: null,
	validationErrors: new Map()
}

export const useCord = (sentence: string, values: Record<string, string | undefined>) => {
	const [state, setState] = useState<CordState>(() => ({
		...initialState,
		values: createStateFromValues(values)
	}))

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

				const dependentInputs = parsed.inputs.filter(input => input.dependentOn === index)
				const newValues = new Map(result.value)

				dependentInputs.forEach(input => {
					newValues.delete(input.index)
				})

				setState(prev => ({
					...prev,
					values: newValues,
					validationErrors: result.error
						? new Map(prev.validationErrors).set(index, {
								type: "validation",
								message: result.error
							})
						: new Map([...prev.validationErrors].filter(([k]) => k !== index))
				}))
			},
			[parsed, state.values]
		)
	}

	const helpers = {
		getInputName: useCallback((index: number) => parsed?.inputs[index].name, [parsed]),
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
			() =>
				parsedWithFilteredInputs?.inputs.every(input => {
					const value = state.values.get(input.index)
					return value !== undefined && value.value.trim() !== ""
				}) ?? false,
			[parsedWithFilteredInputs, state.values]
		),
		isValid: useMemo(() => {
			if (!parsedWithFilteredInputs) return false
			// Check for validation errors
			if (state.validationErrors.size > 0) return false
			// Check that all values that exist are non-empty strings
			const hasEmptyValues = Array.from(state.values.values()).some(
				value => !value?.value || value.value.trim() === ""
			)
			return !hasEmptyValues
		}, [parsedWithFilteredInputs, state.validationErrors, state.values])
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
