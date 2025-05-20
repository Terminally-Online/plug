import { z } from "zod"

import { publicProcedure } from "@/server/api/trpc"

import { buildQueryParams, zerion, zerionApi } from "../lib"

const DetailInputSchema = z.object({
	path: z.object({
		fungibleId: z.string()
	}),
	query: z
		.object({
			currency: z.string().default("usd")
		})
		.optional()
		.default({})
})

const DetailOutputSchema = z.object({
	links: z.object({
		self: z.string()
	}),
	data: z.object({
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
				.nullable(),
			implementations: z.array(
				z.object({
					chain_id: z.string(),
					address: z.string().nullable(),
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
})

export const detail = publicProcedure
	.input(DetailInputSchema)
	.output(DetailOutputSchema)
	.query(async ({ input }) => {
		const queryParams = buildQueryParams({
			currency: input.query.currency
		})

		return zerionApi(() => zerion.get(`/fungibles/${input.path.fungibleId}${queryParams}`))
	})
