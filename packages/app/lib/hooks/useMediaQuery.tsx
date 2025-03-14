import { useEffect, useMemo, useState } from "react"

const BREAKPOINTS = {
	sm: 640,
	md: 768,
	lg: 1024,
	xl: 1280,
	"2xl": 1536
} as const

type Breakpoint = keyof typeof BREAKPOINTS

type BreakpointState = {
	[K in Breakpoint]: boolean
}

export const useMediaQuery = () => {
	const mediaQueries = useMemo(
		() => Object.entries(BREAKPOINTS).map(([key, value]) => [key, `(min-width: ${value}px)`] as const),
		[]
	)

	const [state, setState] = useState<BreakpointState>(
		() =>
			Object.keys(BREAKPOINTS).reduce(
				(acc, key) => ({
					...acc,
					[key]: true
				}),
				{}
			) as BreakpointState
	)

	useEffect(() => {
		const mediaQueryLists = mediaQueries.map(([_, query]) => window.matchMedia(query))

		const handleChange = () => {
			const newState = mediaQueryLists.reduce(
				(acc, mql, index) => ({
					...acc,
					[mediaQueries[index][0]]: mql.matches
				}),
				{} as BreakpointState
			)

			setState(newState)
		}

		mediaQueryLists.forEach(mql => mql.addListener(handleChange))
		handleChange() // Set initial state

		return () => {
			mediaQueryLists.forEach(mql => mql.removeListener(handleChange))
		}
	}, [mediaQueries])

	return state
}
