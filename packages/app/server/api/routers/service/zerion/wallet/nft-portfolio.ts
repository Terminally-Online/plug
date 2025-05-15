import { z } from "zod"

import { protectedProcedure } from "@/server/api/trpc"

import { buildQueryParams, zerion, zerionApi } from "../lib"

const NftPortfolioInputSchema = z.object({
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

const NftPortfolioOutputSchema = z.object({
	links: z.object({
		self: z.string()
	}),
	data: z.object({
		type: z.literal("wallet_nft_portfolio"),
		id: z.string(),
		attributes: z.object({
			positions_distribution_by_chain: z.record(z.number())
		})
	})
})

export const nftPortfolio = protectedProcedure
	.input(NftPortfolioInputSchema)
	.output(NftPortfolioOutputSchema)
	.query(async ({ ctx, input }) => {
		const address = input.path.address || ctx.session.address

		const queryParams = buildQueryParams({
			currency: input.query.currency,
			"filter[chain_ids]": input.query?.filter?.chainIds,
			"filter[fungible_ids]": input.query?.filter?.fungibleIds
		})

		return zerionApi(() => zerion.get(`/wallets/${address}/nft-portfolio/${queryParams}`))
	})
