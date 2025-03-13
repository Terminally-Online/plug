import { NextPageContext } from "next"

import superjson from "superjson"

import { loggerLink, splitLink, unstable_httpBatchStreamLink, unstable_httpSubscriptionLink } from "@trpc/client"

const getUrl = () => {
	const base = (() => {
		if (typeof window !== "undefined") return window.location.origin
		if (process.env.APP_URL) return process.env.APP_URL
		return `http://localhost:${process.env.PORT ?? 3000}`
	})()

	return `${base}/api/trpc`
}

export const createLinks = (ctx: NextPageContext | undefined) => [
	loggerLink({
		enabled: opts =>
			process.env.NODE_ENV === "development" || (opts.direction === "down" && opts.result instanceof Error)
	}),
	splitLink({
		condition: op => op.type === "subscription",
		true: unstable_httpSubscriptionLink({
			url: getUrl(),
			transformer: superjson
		}),
		false: unstable_httpBatchStreamLink({
			url: getUrl(),
			headers() {
				if (!ctx?.req?.headers) return {}
				return { cookie: ctx.req.headers.cookie }
			},
			transformer: superjson
		})
	})
]
