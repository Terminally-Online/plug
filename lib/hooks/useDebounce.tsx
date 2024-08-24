import { useEffect, useRef, useState } from "react"

import { usePathname } from "next/navigation"

export const useDebounce = (
	initial: string,
	delay = 250,
	callback = () => {}
): [
	string,
	string,
	(value: string) => void,
	React.MutableRefObject<string>
] => {
	const pathname = usePathname()

	const valueRef = useRef<string>(initial)

	const [value, setValue] = useState<string>(initial)
	const [debounced, setDebounced] = useState<typeof value>(value)

	useEffect(() => {
		const timeout = setTimeout(() => {
			setDebounced(value)

			if (callback) callback()
		}, delay)

		return () => clearTimeout(timeout)
	}, [value, delay, callback])

	useEffect(() => {
		if (!pathname) return

		setValue(initial)
		setDebounced(initial)
	}, [pathname, initial])

	return [value, debounced, setValue, valueRef]
}

export default useDebounce
