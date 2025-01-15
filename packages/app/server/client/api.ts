import { type inferRouterInputs, type inferRouterOutputs } from "@trpc/server"

import superjson from "superjson"

import { createTRPCNext } from "@trpc/next"

import { type AppRouter } from "@/server/api/root"

import { createLinks } from "./links"

export const api = createTRPCNext<AppRouter>({
	config({ ctx }) {
		return {
			transformer: superjson,
			links: createLinks({ ctx }),
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
	ssr: true
})

export type RouterInputs = inferRouterInputs<AppRouter>
export type RouterOutputs = inferRouterOutputs<AppRouter>
