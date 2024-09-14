import { usePathname } from "next/navigation"
import { useCallback, useEffect, useRef, useState } from "react"

export const useDebounce = (
	initial: string,
	delay = 250,
	callback = (data: string) => {}
): [string, string, (value: string) => void, React.MutableRefObject<string>] => {
	const pathname = usePathname()

	const valueRef = useRef<string>(initial)

	const [value, setValue] = useState<string>(initial)
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

export default useDebounce
