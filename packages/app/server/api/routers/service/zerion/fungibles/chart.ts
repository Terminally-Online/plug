import { z } from "zod"

import { publicProcedure } from "@/server/api/trpc"

import { buildQueryParams, isoDateString, zerion, zerionApi } from "../lib"

const ChartInputSchema = z.object({
	path: z.object({
		fungibleId: z.string(),
		period: z.enum(["hour", "day", "week", "month", "year", "max"])
	}),
	query: z
		.object({
			currency: z.string().default("usd")
		})
		.optional()
		.default({})
})

const ChartOutputSchema = z.object({
	links: z.object({
		self: z.string()
	}),
	data: z.object({
		type: z.literal("fungible_charts"),
		id: z.string(),
		attributes: z.object({
			begin_at: isoDateString,
			end_at: isoDateString,
			stats: z.object({
				first: z.number().nullable(),
				min: z.number().nullable(),
				avg: z.number().nullable(),
				max: z.number().nullable(),
				last: z.number().nullable()
			}),
			points: z.array(z.tuple([z.number(), z.number()]))
		})
	})
})

export const chart = publicProcedure
	.input(ChartInputSchema)
	.output(ChartOutputSchema)
	.query(async ({ input }) => {
		const queryParams = buildQueryParams({
			currency: input.query.currency
		})

		return zerionApi(() =>
			zerion.get(`/fungibles/${input.path.fungibleId}/charts/${input.path.period}${queryParams}`)
		)
	})
