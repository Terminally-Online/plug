import { NextPageContext } from 'next'

import superjson from 'superjson'

import { createWSClient, httpBatchLink, loggerLink, wsLink } from '@trpc/client'
import { createTRPCNext } from '@trpc/next'
import { type inferRouterInputs, type inferRouterOutputs } from '@trpc/server'

import { type AppRouter } from '@/server/api/root'

const getBaseUrl = () => {
	if (typeof window !== 'undefined') return '' // browser should use relative url

	if (process.env.VERCEL_URL) return `https://${process.env.VERCEL_URL}` // SSR should use vercel url

	return `http://localhost:${process.env.PORT ?? 3000}` // dev SSR should use localhost
}

function getEndingLink(ctx: NextPageContext | undefined) {
	if (typeof window === 'undefined') {
		return httpBatchLink({
			url: `${getBaseUrl()}/api/trpc`,
			headers() {
				if (!ctx?.req?.headers) {
					return {}
				}
				// on ssr, forward client's headers to the server
				return {
					...ctx.req.headers,
					'x-ssr': '1'
				}
			}
		})
	}
	const client = createWSClient({
		url: `ws://localhost:3001`
	})
	return wsLink<AppRouter>({
		client
	})
}

/** A set of type-safe react-query hooks for your tRPC API. */
export const api = createTRPCNext<AppRouter>({
	config({ ctx }) {
		return {
			/**
			 * Transformer used for data de-serialization from the server.
			 *
			 * @see https://trpc.io/docs/data-transformers
			 */
			transformer: superjson,

			/**
			 * Links used to determine request flow from client to server.
			 *
			 * @see https://trpc.io/docs/links
			 */
			links: [
				loggerLink({
					enabled: opts =>
						process.env.NODE_ENV === 'development' ||
						(opts.direction === 'down' &&
							opts.result instanceof Error)
				}),
				getEndingLink(ctx)
			]
		}
	},
	/**
	 * Whether tRPC should await queries when server rendering pages.
	 *
	 * @see https://trpc.io/docs/nextjs#ssr-boolean-default-false
	 */
	ssr: true
})

/**
 * Inference helper for inputs.
 *
 * @example type HelloInput = RouterInputs['example']['hello']
 */
export type RouterInputs = inferRouterInputs<AppRouter>

/**
 * Inference helper for outputs.
 *
 * @example type HelloOutput = RouterOutputs['example']['hello']
 */
export type RouterOutputs = inferRouterOutputs<AppRouter>
