import { getSession } from "next-auth/react"

import { ACTION_REGEX } from "@/contexts"
import { actions, categories } from "@/lib"

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

export const getValues = (
	categoryName: keyof typeof categories,
	actionName: keyof (typeof actions)[keyof typeof categories]
) => {
	const staticAction = actions[categoryName][actionName]

	const fragments = staticAction
		? (staticAction["sentence"].split(ACTION_REGEX) as string[])
		: []

	const dynamic = fragments.filter(fragment => fragment.match(ACTION_REGEX))

	return Array(dynamic.length).fill(undefined)
}

export const getIndexes = (fragment: string) => {
	const sanitized = fragment.replace("{", "").replace("}", "").split("=>")

	if (sanitized.length > 1)
		return [sanitized[0], sanitized[1]].map(Number) as [number, number]

	return [null, Number(sanitized[0])] as [null, number]
}
