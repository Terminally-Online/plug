import { z } from "zod"

import { publicProcedure } from "@/server/api/trpc"

import { buildQueryParams, zerion, zerionApi } from "../lib"

const ListInputSchema = z.object({
	query: z
		.object({
			currency: z.string().default("usd"),
			filter: z
				.object({
					searchQuery: z.string().min(2).max(64).optional(),
					implementationChainId: z.string().optional(),
					implementationAddress: z.string().min(32).max(44).optional(),
					fungibleIds: z.array(z.string()).max(50).optional()
				})
				.optional(),
			sort: z.string().optional().default("-market_data.market_cap"),
			page: z
				.object({
					size: z.number().optional(),
					after: z.string().optional()
				})
				.optional()
		})
		.optional()
		.default({})
})

const ListOutputSchema = z.object({
	links: z.object({
		self: z.string(),
		first: z.string().optional(),
		next: z.string().optional(),
		prev: z.string().optional()
	}),
	data: z.array(
		z.object({
			type: z.literal("fungibles"),
			id: z.string(),
			attributes: z.object({
				name: z.string(),
				symbol: z.string(),
				description: z.string().nullable(),
				icon: z
					.object({
						url: z.string()
					})
					.nullable(),
				flags: z.object({
					verified: z.boolean()
				}),
				external_links: z
					.array(
						z.object({
							type: z.string(),
							name: z.string(),
							url: z.string()
						})
					)
					.optional(),
				implementations: z.array(
					z.object({
						chain_id: z.string(),
						address: z.string(),
						decimals: z.number()
					})
				),
				market_data: z
					.object({
						total_supply: z.number().nullable(),
						circulating_supply: z.number().nullable(),
						fully_diluted_valuation: z.number().nullable(),
						market_cap: z.number().nullable(),
						price: z.number().nullable(),
						changes: z
							.object({
								percent_1d: z.number().nullable(),
								percent_30d: z.number().nullable(),
								percent_90d: z.number().nullable(),
								percent_365d: z.number().nullable()
							})
							.nullable()
					})
					.nullable()
			}),
			relationships: z.object({
				chart_hour: z.object({
					links: z.object({
						related: z.string()
					}),
					data: z.object({
						type: z.string(),
						id: z.string()
					})
				}),
				chart_day: z.object({
					links: z.object({
						related: z.string()
					}),
					data: z.object({
						type: z.string(),
						id: z.string()
					})
				}),
				chart_week: z.object({
					links: z.object({
						related: z.string()
					}),
					data: z.object({
						type: z.string(),
						id: z.string()
					})
				}),
				chart_month: z.object({
					links: z.object({
						related: z.string()
					}),
					data: z.object({
						type: z.string(),
						id: z.string()
					})
				}),
				chart_year: z.object({
					links: z.object({
						related: z.string()
					}),
					data: z.object({
						type: z.string(),
						id: z.string()
					})
				}),
				chart_max: z.object({
					links: z.object({
						related: z.string()
					}),
					data: z.object({
						type: z.string(),
						id: z.string()
					})
				})
			})
		})
	)
})

export const list = publicProcedure
	.input(ListInputSchema)
	.output(ListOutputSchema)
	.query(async ({ input }) => {
		const queryParams = buildQueryParams({
			currency: input.query.currency,
			"filter[search_query]": input.query?.filter?.searchQuery,
			"filter[implementation_chain_id]": input.query?.filter?.implementationChainId,
			"filter[implementation_address]": input.query?.filter?.implementationAddress,
			"filter[fungible_ids]": input.query?.filter?.fungibleIds,
			sort: input.query.sort,
			"page[size]": input.query?.page?.size,
			"page[after]": input.query?.page?.after
		})

		return zerionApi(() => zerion.get(`/fungibles/${queryParams}`))
	})
