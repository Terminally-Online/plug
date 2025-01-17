import { TRPCError } from "@trpc/server"

import { z } from "zod"

import { anonymousProtectedProcedure, createTRPCRouter } from "@/server/api/trpc"
import { Actions } from "@/lib/types"

export const events = {
	edit: "edit-plug",
	queue: "queue-plug"
} as const

export const action = createTRPCRouter({
	edit: anonymousProtectedProcedure
		.input(
			z.object({
				id: z.string().optional(),
				actions: z.string()
			})
		)
		.mutation(async ({ input, ctx }) => {
			if (input.id === undefined) throw new TRPCError({ code: "BAD_REQUEST" })

			try {
				const actions = JSON.parse(input.actions) as Actions
				const dominantProtocol = getDominantProtocol(actions)

				const plug = await ctx.db.workflow.update({
					where: { id: input.id, socketId: ctx.session.address },
					data: {
						actions: input.actions,
						color: dominantProtocol, // Store protocol name instead of hex
						updatedAt: new Date()
					}
				})

				ctx.emitter.emit(events.edit, plug)
				return plug
			} catch {
				throw new TRPCError({ code: "BAD_REQUEST" })
			}
		})
})

// Helper function to get dominant protocol
const getDominantProtocol = (actions: Actions): string => {
	if (!actions?.length) return 'plug'

	// Count protocol frequency
	const protocolFrequency: Record<string, number> = {}

	for (const action of actions) {
		if (!action?.protocol) continue

		// Normalize protocol name
		const normalizedProtocol = action.protocol
			.split('_')[0]  // Remove version numbers
			.toLowerCase()

		protocolFrequency[normalizedProtocol] = 
			(protocolFrequency[normalizedProtocol] || 0) + 1
	}

	// Find protocol with highest frequency
	const entries = Object.entries(protocolFrequency)
	if (!entries.length) return 'plug'

	return entries.reduce((a, b) => a[1] > b[1] ? a : b)[0]
}
