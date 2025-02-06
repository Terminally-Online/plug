import { z } from "zod"

import { schemas, intent } from "@/lib"
import { anonymousProtectedProcedure, createTRPCRouter } from "@/server/api/trpc"

export const events = {
	edit: "edit-plug",
	queue: "queue-plug"
} as const

export const actions = createTRPCRouter({
	schemas: anonymousProtectedProcedure
		.input(
			z.object({
				protocol: z.string().optional(),
				action: z.string().optional(),
				chainId: z.number()
			})
		)
		.query(async ({ input, ctx }) => await schemas(input?.protocol, input?.action, input.chainId, ctx.session.address)),
	intent: anonymousProtectedProcedure
		.input(
			z.object({
				chainId: z.number(),
				from: z.string(),
				inputs: z.array(
					z.object({
						protocol: z.string(),
						action: z.string(),
						tokenIn: z.string(),
						tokenOut: z.string(),
						amountOut: z.string()
					})
				)
			})
		)
		.query(async ({ input }) => await intent(input))
})
