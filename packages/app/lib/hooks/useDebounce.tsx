import { usePathname } from "next/navigation"
import { useCallback, useEffect, useRef, useState } from "react"

const DEFAULT_DELAY = 250

export const useDebounce = <T,>(
	initial: T,
	delay = DEFAULT_DELAY,
	callback = (_: T) => {}
): [T, T, (value: T) => void, React.ForwardedRef<T>] => {
	const pathname = usePathname()

	const valueRef = useRef<T>(initial)

	const [value, setValue] = useState<T>(initial)
	const [debounced, setDebounced] = useState<typeof value>(value)

	const debouncedCallback = useCallback(() => {
		setDebounced(value)
		if (callback) callback(value)
	}, [value, callback])

	useEffect(() => {
		const timeout = setTimeout(debouncedCallback, delay)

		return () => clearTimeout(timeout)
	}, [value, delay, debouncedCallback])

	useEffect(() => {
		if (!pathname) return

		setValue(initial)
		setDebounced(initial)
	}, [pathname, initial])

	return [value, debounced, setValue, valueRef]
}

export const useDebounceInline = <T,>(
	value: T,
	delay = DEFAULT_DELAY,
	callback?: (value: T) => void
): T => {
	const [debouncedValue, setDebouncedValue] = useState<T>(value)

	useEffect(() => {
		const timeout = setTimeout(() => {
			setDebouncedValue(value)
			callback?.(value)
		}, delay)

		return () => clearTimeout(timeout)
	}, [value, delay, callback])

	return debouncedValue
}
