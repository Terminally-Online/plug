import { z } from "zod"

import { protectedProcedure } from "@/server/api/trpc"

import { buildQueryParams, isoDateString, zerion, zerionApi } from "../lib"

const TransactionsInputSchema = z.object({
	path: z.object({
		address: z.string().optional()
	}),
	query: z
		.object({
			currency: z.string().default("usd"),
			filter: z
				.object({
					chainIds: z.array(z.string()).optional(),
					operationTypes: z.array(z.string()).optional(),
					assetTypes: z.array(z.string()).optional(),
					fungibleIds: z.array(z.string()).optional(),
					minMinedAt: isoDateString.optional(),
					maxMinedAt: isoDateString.optional(),
					trash: z.enum(["only_trash", "only_non_trash", "no_filter"]).default("no_filter"),
					searchQuery: z.string().optional(),
					fungibleImplementations: z.array(z.string()).optional()
				})
				.optional(),
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

const TransactionsOutputSchema = z.object({
	links: z.object({
		self: z.string(),
		next: z.string().optional()
	}),
	data: z.array(
		z.object({
			type: z.string(),
			id: z.string(),
			attributes: z.object({
				operation_type: z.string(),
				hash: z.string(),
				mined_at_block: z.number(),
				mined_at: isoDateString,
				sent_from: z.string(),
				sent_to: z.string(),
				status: z.string(),
				nonce: z.number(),
				fee: z.object({
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
					quantity: z.object({
						int: z.string(),
						decimals: z.number(),
						float: z.number(),
						numeric: z.string()
					}),
					price: z.number().nullable(),
					value: z.number().nullable()
				}),
				transfers: z.array(
					z.object({
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
						quantity: z.object({
							int: z.string(),
							decimals: z.number(),
							float: z.number(),
							numeric: z.string()
						}),
						price: z.number().nullable(),
						value: z.number().nullable(),
						direction: z.enum(["in", "out"]).optional(),
						counterparty: z.string().optional()
					})
				),
				changes: z.record(z.any()).optional(),
				meta: z.record(z.any()).optional()
			}),
			relationships: z
				.object({
					chain: z.object({
						links: z.object({
							related: z.string()
						}),
						data: z.object({
							type: z.string(),
							id: z.string()
						})
					})
				})
				.optional()
		})
	)
})

export const transactions = protectedProcedure
	.input(TransactionsInputSchema)
	.output(TransactionsOutputSchema)
	.query(async ({ ctx, input }) => {
		const address = input.path.address || ctx.session.address

		const queryParams = buildQueryParams({
			currency: input.query.currency,
			"filter[chain_ids]": input.query?.filter?.chainIds,
			"filter[operation_types]": input.query?.filter?.operationTypes,
			"filter[asset_types]": input.query?.filter?.assetTypes,
			"filter[fungible_ids]": input.query?.filter?.fungibleIds,
			"filter[min_mined_at]": input.query?.filter?.minMinedAt,
			"filter[max_mined_at]": input.query?.filter?.maxMinedAt,
			"filter[trash]": input.query?.filter?.trash,
			"filter[search_query]": input.query?.filter?.searchQuery,
			"filter[fungible_implementations]": input.query?.filter?.fungibleImplementations,
			"page[size]": input.query?.page?.size,
			"page[after]": input.query?.page?.after
		})

		return zerionApi(() => zerion.get(`/wallets/${address}/transactions/${queryParams}`))
	})
