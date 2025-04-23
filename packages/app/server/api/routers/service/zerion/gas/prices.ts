import { z } from "zod"

import { publicProcedure } from "@/server/api/trpc"

import { buildQueryParams, isoDateString, zerion, zerionApi } from "../lib"

const PricesInputSchema = z.object({
	query: z
		.object({
			filter: z
				.object({
					chainIds: z.array(z.string()).optional(),
					gasTypes: z.array(z.string()).optional()
				})
				.optional()
		})
		.optional()
		.default({})
})

const PricesOutputSchema = z.object({
	links: z.object({
		self: z.string()
	}),
	data: z.array(
		z.object({
			type: z.literal("gas-prices"),
			id: z.string(),
			attributes: z.object({
				gas_type: z.string(),
				updated_at: isoDateString,
				info: z.record(z.any())
			}),
			relationships: z.object({
				chain: z.object({
					links: z.object({
						related: z.string()
					}),
					data: z.object({
						type: z.literal("chains"),
						id: z.string()
					})
				})
			})
		})
	)
})

export const prices = publicProcedure
	.input(PricesInputSchema)
	.output(PricesOutputSchema)
	.query(async ({ input }) => {
		const queryParams = buildQueryParams({
			"filter[chain_ids]": input.query?.filter?.chainIds,
			"filter[gas_types]": input.query?.filter?.gasTypes
		})

		return zerionApi(() => zerion.get(`/gas-prices/${queryParams}`))
	})
