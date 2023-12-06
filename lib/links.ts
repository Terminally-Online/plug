import { NextPageContext } from "next"

import { createWSClient, httpBatchLink, loggerLink, wsLink } from "@trpc/client"

import { type AppRouter } from "@/server/api/root"

export const getBaseUrl = () => {
	// browser should use relative url
	if (typeof window !== "undefined") return ""

	// SSR should use vercel url
	if (process.env.VERCEL_URL) return `https://${process.env.VERCEL_URL}`

	const envUrl = process.env.NEXT_PUBLIC_APP_URL

	if (envUrl) return envUrl

	// dev SSR should use localhost
	return `http://localhost:${process.env.PORT ?? 3000}`
}

function getEndingLink(ctx: NextPageContext | undefined) {
	if (typeof window === "undefined") {
		return httpBatchLink({
			url: `${getBaseUrl()}/api/trpc`,
			headers() {
				if (!ctx?.req?.headers) {
					return {}
				}
				return {
					...ctx.req.headers,
					"x-ssr": "1"
				}
			}
		})
	}
	const client = createWSClient({
		url: process.env.NEXT_PUBLIC_WS_URL || `ws://localhost:3001`
	})
	return wsLink<AppRouter>({
		client
	})
}

export const createLinks = (ctx: NextPageContext | undefined) => [
	loggerLink({
		enabled: opts =>
			process.env.NODE_ENV === "development" ||
			(opts.direction === "down" && opts.result instanceof Error)
	}),
	getEndingLink(ctx)
]
