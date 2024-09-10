import { useRouter } from "next/router"

import { sendGTMEvent } from "@next/third-parties/google"

export const useAnalytics = (event: string, topValue?: string, redirect = true) => {
	const router = useRouter()

	return (bottomValue?: string) => {
		const value = bottomValue ?? topValue
		sendGTMEvent({ event, value })

		if (redirect && value) router.push(value)
	}
}
