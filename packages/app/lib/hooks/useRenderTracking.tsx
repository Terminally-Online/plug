import { useEffect, useRef } from "react"

export const useRenderTracking = (componentName: string) => {
	const renderCount = useRef(0)

	useEffect(() => {
		renderCount.current += 1
		console.log(`${componentName} has rendered ${renderCount.current} times`)
	})

	return renderCount
}
