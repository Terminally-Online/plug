import { NextPageContext } from "next"
import { signIn } from "next-auth/react"

import superjson from "superjson"

import { httpBatchLink, loggerLink } from "@trpc/client"
import { TRPCClientError } from "@trpc/client"

export function createLinks(ctx: { ctx?: NextPageContext }) {
	const url = process.env.NEXT_PUBLIC_APP_URL
		? `${process.env.NEXT_PUBLIC_APP_URL}/api/trpc`
		: "http://localhost:3000/api/trpc"

	return [
		loggerLink({
			enabled: opts =>
				process.env.NODE_ENV === "development" || (opts.direction === "down" && opts.result instanceof Error)
		}),
		httpBatchLink({
			url,
			headers() {
				if (ctx.ctx?.req) {
					// Pass headers from SSR request
					return {
						...ctx.ctx.req.headers,
						"x-ssr": "1"
					}
				}
				return {}
			},
			async onError(err) {
				if (err.data?.code === "UNAUTHORIZED") {
					// Instead of letting the error redirect to error page,
					// we can handle session expiry by redirecting to sign in
					if (typeof window !== "undefined") {
						await signIn()
					}
				}
			}
		})
	]
}
