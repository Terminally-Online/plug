import { z } from "zod"

import { getSchemas } from "@/lib"
import { anonymousProtectedProcedure, createTRPCRouter } from "@/server/api/trpc"

export const events = {
	edit: "edit-plug",
	queue: "queue-plug"
} as const

export const actions = createTRPCRouter({
	get: anonymousProtectedProcedure
		.input(
			z
				.object({
					protocol: z.string().optional(),
					action: z.string().optional()
				})
				.optional()
		)
		.query(async ({ input }) => await getSchemas(input?.protocol, input?.action))
})
