import { z } from "zod"

import { protectedProcedure } from "@/server/api/trpc"

import { buildQueryParams, isoDateString, zerion, zerionApi } from "../lib"

const ChartInputSchema = z.object({
	path: z.object({
		period: z.enum(["hour", "day", "week", "month", "year", "max"]),
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

const ChartOutputSchema = z.object({
	links: z.object({
		self: z.string()
	}),
	data: z.object({
		type: z.string(),
		id: z.string(),
		attributes: z.object({
			begin_at: isoDateString,
			end_at: isoDateString,
			points: z.array(z.tuple([z.number(), z.number()]))
		})
	})
})

export const chart = protectedProcedure
	.input(ChartInputSchema)
	.output(ChartOutputSchema)
	.query(async ({ ctx, input }) => {
		const address = input.path.address || ctx.session.address

		const queryParams = buildQueryParams({
			currency: input.query.currency,
			"filter[chain_ids]": input.query?.filter?.chainIds,
			"filter[fungible_ids]": input.query?.filter?.fungibleIds
		})

		return zerionApi(() => zerion.get(`/wallets/${address}/charts/${input.path.period}${queryParams}`))
	})
