import { z } from "zod"

import { protectedProcedure } from "@/server/api/trpc"

import { buildQueryParams, zerion, zerionApi } from "../lib"

const PortfolioInputSchema = z.object({
	path: z.object({
		address: z.string().optional()
	}),
	query: z
		.object({
			currency: z.string().default("usd"),
			filter: z
				.object({
					positions: z.enum(["only_simple", "only_complex", "no_filter"]).default("only_simple")
				})
				.optional()
		})
		.optional()
		.default({})
})

const PortfolioSchema = z.object({
	links: z.object({
		self: z.string()
	}),
	data: z.object({
		type: z.string(),
		id: z.string(),
		attributes: z.object({
			positions_distribution_by_type: z.object({
				wallet: z.number(),
				deposited: z.number(),
				borrowed: z.number(),
				locked: z.number(),
				staked: z.number()
			}),
			positions_distribution_by_chain: z.record(z.string(), z.number()),
			total: z.object({
				positions: z.number()
			}),
			changes: z.object({
				absolute_1d: z.number(),
				percent_1d: z.number()
			})
		})
	})
})

export const portfolio = protectedProcedure
	.input(PortfolioInputSchema)
	.output(PortfolioSchema)
	.query(async ({ ctx, input }) => {
		const address = input.path.address || ctx.session.address

		const queryParams = buildQueryParams({
			currency: input.query.currency,
			"filter[positions]": input.query?.filter?.positions
		})

		return zerionApi(() => zerion.get(`/wallets/${address}/portfolio${queryParams}`))
	})
