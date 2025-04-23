import { z } from "zod"

import { protectedProcedure } from "@/server/api/trpc"

import { buildQueryParams, zerion, zerionApi } from "../lib"

const PositionsInputSchema = z.object({
	path: z.object({
		address: z.string().optional()
	}),
	query: z
		.object({
			currency: z.string().default("usd"),
			filter: z
				.object({
					positions: z.enum(["only_simple", "only_complex", "no_filter"]).default("only_simple"),
					chainIds: z.array(z.string()).optional(),
					positionTypes: z.array(z.string()).optional(),
					fungibleIds: z.array(z.string()).optional(),
					dappIds: z.array(z.string()).optional(),
					trash: z.enum(["only_trash", "only_non_trash"]).default("only_non_trash")
				})
				.optional(),
			sort: z.enum(["value", "recent"]).default("value")
		})
		.optional()
		.default({})
})

const PositionsOutputSchema = z.object({
	links: z.object({
		self: z.string()
	}),
	data: z.array(
		z.object({
			type: z.string(),
			id: z.string(),
			attributes: z.object({
				parent: z.any().nullable(),
				protocol: z.any().nullable(),
				name: z.string(),
				position_type: z.string(),
				quantity: z.object({
					int: z.string(),
					decimals: z.number(),
					float: z.number(),
					numeric: z.string()
				}),
				value: z.number().nullable(),
				price: z.number().nullable(),
				changes: z
					.object({
						absolute_1d: z.number(),
						percent_1d: z.number()
					})
					.nullable(),
				fungible_info: z.object({
					name: z.string(),
					symbol: z.string(),
					icon: z
						.object({
							url: z.string()
						})
						.nullable(),
					flags: z.object({
						verified: z.boolean()
					}),
					implementations: z.array(
						z.object({
							chain_id: z.string(),
							address: z.string().nullable(),
							decimals: z.number()
						})
					)
				}),
				flags: z.object({
					displayable: z.boolean(),
					is_trash: z.boolean()
				}),
				updated_at: z.string(),
				updated_at_block: z.number()
			}),
			relationships: z.object({
				chain: z.object({
					links: z.object({
						related: z.string()
					}),
					data: z.object({
						type: z.string(),
						id: z.string()
					})
				}),
				fungible: z.object({
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

export const positions = protectedProcedure
	.input(PositionsInputSchema)
	.output(PositionsOutputSchema)
	.query(async ({ ctx, input }) => {
		const address = input.path.address || ctx.session.address

		const queryParams = buildQueryParams({
			currency: input.query.currency,
			"filter[positions]": input.query?.filter?.positions,
			"filter[chain_ids]": input.query?.filter?.chainIds,
			"filter[position_types]": input.query?.filter?.positionTypes,
			"filter[fungible_ids]": input.query?.filter?.fungibleIds,
			"filter[dapp_ids]": input.query?.filter?.dappIds,
			"filter[trash]": input.query?.filter?.trash,
			sort: input.query.sort
		})

		return zerionApi(() => zerion.get(`/wallets/${address}/positions/${queryParams}`))
	})
