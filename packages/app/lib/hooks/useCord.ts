import { useCallback, useMemo, useState } from "react"

import {
	CordState,
	createInitialState,
	InputError,
	InputState,
	parseCordSentence,
	resolveSentence,
	setValue,
	shouldRenderInput,
} from "@terminallyonline/cord"

// Helper to create internal Map from object values
const createStateFromValues = (values: Record<string, string | undefined>) => {
	const state = new Map<number, InputState>()
	Object.entries(values).forEach(([key, value]) => {
		if (value !== undefined) {
			state.set(Number(key), { value })
		}
	})
	return state
}

/**
 * A stateless version of useCord that works with external state management
 * Use this when you want to store values in your global state
 */
export const useCordStateless = (
	sentence: string, 
	values: Record<string, string | undefined>,
	onUpdateValue?: (index: number, value: string | undefined, error?: string) => void
) => {
	// Convert object values to internal Map
	const valuesMap = useMemo(() => createStateFromValues(values), [values])
	
	// Track validation errors separately
	const validationErrors = useMemo(() => {
		const errors = new Map<number, InputError>()
		return errors
	}, [])

	// Parse the sentence
	const parsed = useMemo(() => {
		const result = parseCordSentence(sentence)
		if (!result.success) {
			return null
		}
		return result.value
	}, [sentence])

	// Filter inputs based on dependencies
	const filteredInputs = useMemo(() => {
		if (!parsed) return []
		return parsed.inputs.filter(input => 
			shouldRenderInput(input.type, parsed.inputs, index => valuesMap.get(index))
		)
	}, [parsed, valuesMap])

	// Create parsed sentence with filtered inputs
	const parsedWithFilteredInputs = useMemo(() => {
		if (!parsed) return null
		return {
			...parsed,
			inputs: filteredInputs
		}
	}, [parsed, filteredInputs])

	// Split sentence into parts for rendering
	const parts = useMemo(() => parsedWithFilteredInputs
		? parsedWithFilteredInputs.template
			.split(/(\{[^}]+\})/g)
			.map(part => {
				if (part.match(/\{[^}]+\}/)) return [part]
				return part.split(/(\s+)/g)
			})
			.flat()
		: [], [parsedWithFilteredInputs])

	// Attempt to resolve the complete sentence
	const resolvedSentence = useMemo(() => {
		if (!parsed) return null

		const allInputsHaveValues = parsed.inputs.every(input => valuesMap.has(input.index))
		if (!allInputsHaveValues) return null

		const result = resolveSentence(parsed, valuesMap)
		return result.success ? result.value : null
	}, [parsed, valuesMap])

	// Actions
	const actions = {
		// Process a value change and notify parent component
		setValue: useCallback(
			(index: number, value: string | undefined) => {
				if (!parsed || !onUpdateValue) return

				// Validate the value using cord's built-in validator
				const result = setValue({
					parsedSentence: parsed,
					currentValues: valuesMap,
					index,
					value: value ?? ""
				})

				// Handle dependent inputs that need to be cleared
				const dependentInputs = parsed.inputs.filter(input => input.dependentOn === index)
				dependentInputs.forEach(input => {
					onUpdateValue(input.index, undefined)
				})

				// Notify parent component of the update with any errors
				onUpdateValue(index, value, result.error)
			},
			[parsed, valuesMap, onUpdateValue]
		)
	}

	// Helper functions
	const helpers = {
		getInputName: useCallback((index: number) => parsed?.inputs[index].name, [parsed]),
		getInputValue: useCallback((index: number) => valuesMap.get(index), [valuesMap]),
		getInputError: useCallback((index: number) => validationErrors.get(index), [validationErrors]),
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
			() => parsedWithFilteredInputs?.inputs.every(input => {
				const value = valuesMap.get(input.index)
				return value !== undefined
			}) ?? false,
			[parsedWithFilteredInputs, valuesMap]
		),
		isValid: useMemo(() => {
			if (!parsedWithFilteredInputs) return false
			// Check for validation errors
			if (validationErrors.size > 0) return false
			// Check that all values exist and are non-empty
			return !Array.from(valuesMap.values()).some(value => !value || value.value === "")
		}, [parsedWithFilteredInputs, validationErrors, valuesMap])
	}

	return {
		state: {
			parsed: parsedWithFilteredInputs,
			parts,
			resolvedSentence,
			values: valuesMap,
			error: null,
			validationErrors
		},
		actions,
		helpers
	}
}

// Keep original useCord for backward compatibility
export const useCord = (sentence: string, values: Record<string, string | undefined>) => {
	// Internal state storage
	const initialState: CordState = {
		values: createInitialState(),
		parsed: null,
		resolvedSentence: null,
		error: null,
		validationErrors: new Map()
	}

	// Internal state to track current values
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

	const parts = useMemo(() => parsedWithFilteredInputs
		? parsedWithFilteredInputs.template
			.split(/(\{[^}]+\})/g)
			.map(part => {
				if (part.match(/\{[^}]+\}/)) return [part]
				return part.split(/(\s+)/g)
			})
			.flat()
		: [], [parsedWithFilteredInputs])

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
			(index: number, value: string | undefined) => {
				if (!parsed) return

				const result = setValue({
					parsedSentence: parsed,
					currentValues: state.values,
					index,
					value: value ?? ""
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
					return value !== undefined
					// return value !== undefined && value.value.trim() !== ""
				}) ?? false,
			[parsedWithFilteredInputs, state.values]
		),
		isValid: useMemo(() => {
			if (!parsedWithFilteredInputs) return false
			// Check for validation errors
			if (state.validationErrors.size > 0) return false
			// Check that all values that exist are non-empty strings
			return !Array.from(state.values.values()).some(value => value?.value === "")
		}, [parsedWithFilteredInputs, state.validationErrors, state.values])
	}

	return {
		state: {
			...state,
			parsed: parsedWithFilteredInputs,
			parts,
			resolvedSentence
		},
		actions,
		helpers
	}
}
