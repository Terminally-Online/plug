import { getSession } from "next-auth/react"

import { ACTION_REGEX } from "@/contexts"

export const isConnected = async (ctx: any, callback: () => any) => {
	if (!(await getSession(ctx))) {
		return {
			redirect: {
				destination: `?connect=true`,
				permanent: false
			}
		}
	}

	return callback()
}

export const getValues = (sentence: string) => {
	const fragments = sentence.split(ACTION_REGEX) as string[]

	const dynamic = fragments.filter(fragment => fragment.match(ACTION_REGEX))

	return Array(dynamic.length).fill(undefined)
}

export const getIndexes = (fragment: string) => {
	const sanitized = fragment.replace("{", "").replace("}", "").split("=>")

	if (sanitized.length > 1) return [sanitized[0], sanitized[1]].map(Number) as [number, number]

	return [null, Number(sanitized[0])] as [null, number]
}
