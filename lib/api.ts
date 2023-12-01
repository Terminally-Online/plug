import superjson from 'superjson'

import { createTRPCNext } from '@trpc/next'
import { type inferRouterInputs, type inferRouterOutputs } from '@trpc/server'

import { type AppRouter } from '@/server/api/root'

import { createLinks } from './links'

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
			links: createLinks(ctx)
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
