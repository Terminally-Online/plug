import { type inferRouterInputs, type inferRouterOutputs } from "@trpc/server"

import superjson from "superjson"
import { ssrPrepass } from '@trpc/next/ssrPrepass';

import { createTRPCNext } from "@trpc/next"

import { type AppRouter } from "@/server/api/root"

import { createLinks } from "./links"

export const api = createTRPCNext<AppRouter>({
	transformer: superjson,
	ssrPrepass,
	ssr: true,
	config(opts) {
		return {
			links: createLinks(opts.ctx),
			queryClientConfig: {
				defaultOptions: {
					queries: {
						retry: (failureCount, error: any) => {
							if (error?.data?.code === "UNAUTHORIZED") return false
							return failureCount < 2
						},
						staleTime: 5000
					}
				}
			}
		}
	},
	responseMeta(opts) {
		const { clientErrors } = opts;
		if (clientErrors.length) {
			return {
				status: clientErrors[0].data?.httpStatus ?? 500,
			};
		}
		const ONE_DAY_IN_SECONDS = 60 * 60 * 24;
		return {
			headers: new Headers([
				[
					'cache-control',
					`s-maxage=1, stale-while-revalidate=${ONE_DAY_IN_SECONDS}`,
				],
			]),
		};
	},
})

export type RouterInputs = inferRouterInputs<AppRouter>
export type RouterOutputs = inferRouterOutputs<AppRouter>
