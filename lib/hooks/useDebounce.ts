"use client"

import { useEffect, useRef, useState } from "react"

export const useDebounce = (
	initial: string,
	delay = 250
): [
	string,
	string,
	(value: string) => void,
	React.MutableRefObject<string>
] => {
	const valueRef = useRef<string>(initial)

	const [value, setValue] = useState<string>(initial)
	const [debounced, setDebounced] = useState<typeof value>(value)

	useEffect(() => {
		const timeout = setTimeout(() => {
			setDebounced(value)
		}, delay)

		return () => clearTimeout(timeout)
	}, [value, delay])

	return [value, debounced, setValue, valueRef]
}

export default useDebounce
