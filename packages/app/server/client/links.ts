import { NextPageContext } from "next"

import { createWSClient, httpBatchLink, loggerLink, wsLink } from "@trpc/client"

import { env } from "@/env"
import { type AppRouter } from "@/server/api/root"

export const getBaseUrl = () => {
	if (typeof window !== "undefined") return ""

	if (process.env.VERCEL_URL) return `https://${process.env.VERCEL_URL}`

	const envUrl = env.NEXT_PUBLIC_APP_URL

	if (envUrl) return envUrl

	return `http://localhost:${env.PORT}`
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
		url: env.NEXT_PUBLIC_WS_URL
	})
	return wsLink<AppRouter>({
		client
	})
}

export const createLinks = (ctx: NextPageContext | undefined) => [
	loggerLink({
		enabled: opts =>
			process.env.NODE_ENV === "development" || (opts.direction === "down" && opts.result instanceof Error)
	}),
	getEndingLink(ctx)
]
