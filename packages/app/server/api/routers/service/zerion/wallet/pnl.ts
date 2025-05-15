import { z } from "zod"

import { protectedProcedure } from "@/server/api/trpc"

import { buildQueryParams, zerion, zerionApi } from "../lib"

const PnlInputSchema = z.object({
	path: z.object({
		address: z.string().optional()
	}),
	query: z
		.object({
			currency: z.string().default("usd"),
			filter: z
				.object({
					chainIds: z.array(z.string()).optional(),
					fungibleIds: z.array(z.string()).optional()
				})
				.optional()
		})
		.optional()
		.default({})
})

const PnlOutputSchema = z.object({
	links: z.object({
		self: z.string()
	}),
	data: z.object({
		type: z.literal("pnl"),
		id: z.string(),
		attributes: z.object({
			realized_gain: z.number(),
			unrealized_gain: z.number(),
			total_fee: z.number(),
			net_invested: z.number(),
			received_external: z.number(),
			sent_external: z.number(),
			sent_for_nfts: z.number(),
			received_for_nfts: z.number()
		})
	})
})

export const pnl = protectedProcedure
	.input(PnlInputSchema)
	.output(PnlOutputSchema)
	.query(async ({ ctx, input }) => {
		const address = input.path.address || ctx.session.address

		const queryParams = buildQueryParams({
			currency: input.query.currency,
			"filter[chain_ids]": input.query?.filter?.chainIds,
			"filter[fungible_ids]": input.query?.filter?.fungibleIds
		})

		return zerionApi(() => zerion.get(`/wallets/${address}/pnl/${queryParams}`))
	})
