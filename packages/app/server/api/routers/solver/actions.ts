import { z } from "zod"

import { intent, schemas } from "@/lib"
import { anonymousProtectedProcedure, createTRPCRouter } from "@/server/api/trpc"

export const events = {
	edit: "edit-plug",
	queue: "queue-plug"
} as const

export const actions = createTRPCRouter({
	schemas: anonymousProtectedProcedure
		.input(
			z.object({
				chainId: z.number(),
				protocol: z.string().optional(),
				action: z.string().optional(),
				search: z.array(z.string()).optional()
			})
		)
		.query(
			async ({ input, ctx }) =>
				await schemas(input?.protocol, input?.action, input.chainId, input.search, ctx.session.address)
		),
	intent: anonymousProtectedProcedure
		.input(
			z.object({
				chainId: z.number(),
				from: z.string(),
				inputs: z.array(
					z
						.object({
							protocol: z.string(),
							action: z.string()
						})
						.and(z.record(z.string(), z.union([z.string(), z.number()])))
				),
				options: z
					.object({
						simulate: z.boolean().optional(),
						submit: z.boolean().optional(),
						isEOA: z.boolean().optional()
					})
					.optional()
			})
		)
		.query(async ({ input }) => await intent(input))
})
