import { z } from "zod"

import { protectedProcedure } from "@/server/api/trpc"

import { buildQueryParams, isoDateString, zerion, zerionApi } from "../lib"

const NftCollectionsInputSchema = z.object({
	path: z.object({
		address: z.string().optional()
	}),
	query: z
		.object({
			currency: z.string().default("usd"),
			filter: z
				.object({
					chainIds: z.array(z.string()).optional()
				})
				.optional(),
			sort: z.enum(["-total_floor_price", "total_floor_price"]).optional(),
			include: z.array(z.string()).optional()
		})
		.optional()
		.default({})
})

const NftCollectionsOutputSchema = z.object({
	links: z.object({
		self: z.string()
	}),
	data: z.array(
		z.object({
			type: z.literal("wallet_nft_collections"),
			id: z.union([z.string(), z.number()]),
			attributes: z.object({
				min_changed_at: isoDateString.nullable(),
				max_changed_at: isoDateString.nullable(),
				nfts_count: z.string(),
				total_floor_price: z.number().nullable(),
				collection_info: z
					.object({
						name: z.string(),
						description: z.string().nullable(),
						content: z
							.object({
								icon: z
									.object({
										url: z.string()
									})
									.nullable(),
								banner: z
									.object({
										url: z.string(),
										content_type: z.string().optional()
									})
									.nullable()
							})
							.nullable()
					})
					.nullable()
			}),
			relationships: z.object({
				chains: z.array(
					z.object({
						links: z.object({
							related: z.string()
						}),
						data: z.object({
							type: z.string(),
							id: z.string()
						})
					})
				),
				nft_collection: z.object({
					data: z.object({
						type: z.string(),
						id: z.union([z.string(), z.number()])
					})
				})
			}),
			included: z
				.array(
					z.object({
						type: z.literal("nft_collections"),
						id: z.union([z.string(), z.number()]),
						attributes: z.object({
							metadata: z
								.object({
									name: z.string(),
									description: z.string().nullable(),
									content: z
										.object({
											icon: z
												.object({
													url: z.string()
												})
												.nullable(),
											banner: z
												.object({
													url: z.string(),
													content_type: z.string().optional()
												})
												.nullable()
										})
										.nullable(),
									payment_token_symbol: z.string().optional()
								})
								.nullable(),
							market_data: z
								.object({
									prices: z.object({
										floor: z.number().nullable()
									})
								})
								.optional()
						})
					})
				)
				.optional()
		})
	)
})

export const nftCollections = protectedProcedure
	.input(NftCollectionsInputSchema)
	.output(NftCollectionsOutputSchema)
	.query(async ({ ctx, input }) => {
		const address = input.path.address || ctx.session.address

		const queryParams = buildQueryParams({
			currency: input.query.currency,
			"filter[chain_ids]": input.query?.filter?.chainIds,
			sort: input.query.sort,
			include: input.query.include
		})

		return zerionApi(() => zerion.get(`/wallets/${address}/nft-collections/${queryParams}`))
	})
