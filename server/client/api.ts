import { type AppRouter } from "@/server/api/root"
import { type inferRouterInputs, type inferRouterOutputs } from "@trpc/server"

import superjson from "superjson"

import { createTRPCNext } from "@trpc/next"

import { createLinks } from "./links"

export const api = createTRPCNext<AppRouter>({
	config({ ctx }) {
		return {
			transformer: superjson,
			links: createLinks(ctx)
		}
	},
	ssr: true
})

export type RouterInputs = inferRouterInputs<AppRouter>
export type RouterOutputs = inferRouterOutputs<AppRouter>
