"use client"

import { useEffect, useState } from "react"

export const useDebounce = ({
	initial,
	delay = 500
}: {
	initial: any
	delay?: number
}) => {
	const [value, setValue] = useState<typeof initial>(initial)
	const [debounced, setDebounced] = useState<typeof value>(value)

	useEffect(() => {
		const timeout = setTimeout(() => {
			setDebounced(value)
		}, delay)

		return () => {
			clearTimeout(timeout)
		}
	}, [value, delay])

	return { value, debounced, debounce: setValue }
}

export default useDebounce
