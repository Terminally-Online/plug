import { TRPCError } from "@trpc/server"

import { z } from "zod"

import Anthropic from "@anthropic-ai/sdk"

import { env } from "@/env"
import { Actions } from "@/lib/types"
import { anonymousProtectedProcedure, createTRPCRouter } from "@/server/api/trpc"

export const events = {
	edit: "edit-plug",
	queue: "queue-plug"
} as const

const anthropic = new Anthropic({
	apiKey: env.ANTHROPIC_KEY
})

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

			const plug = await ctx.db.workflow.findUnique({
				where: { id: input.id, socketId: ctx.session.address }
			})

			if (!plug) throw new TRPCError({ code: "NOT_FOUND" })

			// TODO: (#611) Update the tags of the Plug based on the protocol/action pair.
			const tags: string[] = []
			const actions = JSON.parse(input.actions)
			const dominantProtocol = getDominantProtocol(actions)

			void (async (name: string, namedAt: Date | null, updatedAt: Date | null) => {
				// NOTE: If the plug has been updated in the last 30 minutes.
				if (updatedAt && Date.now() - updatedAt.getTime() < 30 * 60 * 1000 && name != "Untitled Plug") return
				// NOTE: If the user has ever named this plug disable auto-naming completely.
				if (namedAt) return
				// NOTE: If it does not yet have action context we want to set the name to "Untitled Plug".
				if (actions.length <= 1) return

				const prompt = `You are an expert at creating clear, intuitive names for automation workflows.
					These names will be used for the algorithm and placement in the discovery feed so the name must
					convey the purpose and intent of the user when designing the workflow.

					Create a concise name (maximum 24 characters) for this workflow based on its actions:

					${JSON.stringify(actions, null, 2)}

					Guidelines:
					- Name MUST describe the specific strategy or outcome
					- Avoid generic terms like "Workflow", "Setup", or "Automation"
					- Use concrete, specific language that tells users exactly what it does
					- Think: "How would a user describe this to a friend?"
					- If it involves trading/investing, name the specific strategy

					Examples of bad vs good names:
					- ❌ "Crypto Workflow" → ✅ "ETH Yield Maximizer"
					- ❌ "Token Setup" → ✅ "USDC to ETH Swing Trade"
					- ❌ "NFT Automation" → ✅ "NFT Mint Sniper"
					- ❌ "DeFi Manager" → ✅ "Stablecoin Yield Farm"

					For DeFi/Web3 workflows:
					- Name the specific assets or protocols if they're key to understanding
					- Describe the financial strategy (e.g., "Yield", "Hedge", "Stake")
					- Make the end benefit clear (e.g., "Profit", "Income", "Growth")
					- Look at the serious of actions/constraints and determine what the collective action is performing.

					After coming up with names, take a moment to consider:
					- Does this name focus on the intent?
					- Is the name too focused on the actions?

					Remember... we want our users to see semantic but colloquial names! 

					Respond with only the name, no explanation or punctuation.`

				const message = await anthropic.messages.create({
					model: "claude-3-haiku-20240307",
					max_tokens: 50,
					messages: [
						{
							role: "user",
							content: prompt
						}
					]
				})
				const generated = message.content[0].type === "text" ? message.content[0].text : undefined
				const namedPlug = await ctx.db.workflow.update({
					where: { id: input.id },
					data: { name: generated, renamedAt: new Date(), updatedAt: new Date() }
				})

				ctx.emitter.emit(events.edit, namedPlug)
			})(plug.name, plug.namedAt, plug.updatedAt)

			const updated = await ctx.db.workflow.update({
				where: { id: input.id, socketId: ctx.session.address },
				data: {
					actions: input.actions,
					tags,
					color: dominantProtocol,
					updatedAt: new Date()
				}
			})

			ctx.emitter.emit(events.edit, updated)

			return updated
		})
})

// Helper function to get dominant protocol
const getDominantProtocol = (actions: Actions): string => {
	if (!actions?.length) return "plug"

	const protocolFrequency: Record<string, number> = {}

	for (const action of actions) {
		if (!action?.protocol) continue

		const normalizedProtocol = action.protocol.split("_")[0].toLowerCase()

		protocolFrequency[normalizedProtocol] = (protocolFrequency[normalizedProtocol] || 0) + 1
	}

	const entries = Object.entries(protocolFrequency)
	if (!entries.length) return "plug"

	return entries.reduce((a, b) => (a[1] > b[1] ? a : b))[0]
}
