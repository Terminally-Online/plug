import { z } from "zod"

import { protectedProcedure } from "@/server/api/trpc"

import { buildQueryParams, isoDateString, zerion, zerionApi } from "../lib"

const NftPositionsInputSchema = z.object({
	path: z.object({
		address: z.string().optional()
	}),
	query: z
		.object({
			currency: z.string().default("usd"),
			filter: z
				.object({
					chainIds: z.array(z.string()).optional(),
					collectionIds: z.array(z.string()).optional()
				})
				.optional(),
			sort: z.enum(["-floor_price", "floor_price", "created_at", "-created_at"]).optional(),
			include: z.array(z.enum(["nft_collections", "nfts"])).default([]),
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

const ContentSchema = z.object({
	url: z.string(),
	content_type: z.string().optional()
})

const NftPositionsOutputSchema = z.object({
	links: z.object({
		self: z.string(),
		next: z.string().optional()
	}),
	data: z.array(
		z.object({
			type: z.string(),
			id: z.string(),
			attributes: z.object({
				changed_at: isoDateString,
				amount: z.preprocess(v => Number(v), z.number()),
				price: z.number().nullable().optional(),
				value: z.number().nullable().optional(),
				nft_info: z.object({
					contract_address: z.string(),
					token_id: z.union([z.string(), z.number()]),
					name: z.string().nullable(),
					interface: z.string(),
					content: z
						.object({
							preview: z
								.object({
									url: z.string(),
									content_type: z.string().optional()
								})
								.optional(),
							detail: z
								.object({
									url: z.string(),
									content_type: z.string().optional()
								})
								.optional(),
							audio: z
								.object({
									url: z.string(),
									content_type: z.string().optional()
								})
								.optional(),
							video: z
								.object({
									url: z.string(),
									content_type: z.string().optional()
								})
								.optional()
						})
						.nullable()
						.optional(),
					flags: z
						.object({
							is_spam: z.boolean().optional()
						})
						.optional()
				}),
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
									.nullable()
									.optional(),
								banner: z
									.object({
										url: z.string(),
										content_type: z.string().optional()
									})
									.nullable()
									.optional()
							})
							.nullable()
					})
					.nullable()
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
				nft: z.object({
					data: z.object({
						type: z.string(),
						id: z.string()
					})
				}),
				nft_collection: z.object({
					data: z.object({
						type: z.string(),
						id: z.union([z.string(), z.number()])
					})
				}),
				wallet_nft_collection: z.object({
					data: z.object({
						type: z.string(),
						id: z.string()
					})
				})
			})
		})
	),
	included: z.array(
		z.union([z.object({
			type: z.literal("nft_collections"),
			id: z.union([z.string(), z.number()]),
			attributes: z.object({
				metadata: z.object({
					name: z.string(),
					description: z.string().optional(),
					icon: ContentSchema.nullable().optional(),
					banner: ContentSchema.nullable().optional(),
					payment_token_symbol: z.string().nullable().optional(),
				}).nullable().optional(),
				market_data: z.object({
					prices: z.object({
						floor: z.number().nullable()
					}).nullable().optional()
				}).nullable().optional()
			}).nullable().optional()
		}), z.object({
			type: z.literal("nfts"),
			id: z.union([z.string(), z.number()]),
			attributes: z.object({
				contract_address: z.string(),
				token_id: z.string().nullable().optional(),
				interface: z.string(),
				metadata: z.object({
					name: z.string(),
					description: z.string().optional(),
					content: z.object({
						preview: ContentSchema.nullable().optional(),
						detail: ContentSchema.nullable().optional(),
						audio: ContentSchema.nullable().optional(),
						video: ContentSchema.nullable().optional(),
					}).nullable().optional(),
					attributes: z.array(z.object({
						key: z.string(),
						value: z.string()
					})).nullable().optional()
				}).nullable().optional(),
				market_data: z.object({
					prices: z.object({
						floor: z.number().nullable()
					}).nullable().optional(),
					last_sale: z.object({
						price: z.number(),
						quantity: z.object({
							int: z.string(),
							decimals: z.number(),
							float: z.number(),
							numeric: z.string()
						})
					}).nullable().optional()
				}).nullable().optional(),
				external_links: z.array(z.object({
					type: z.string(),
					name: z.string(),
					url: z.string(),
				})).nullable().optional(),
			}).nullable().optional()
		})])
	)
})

export const nftPositions = protectedProcedure
	.input(NftPositionsInputSchema)
	.output(NftPositionsOutputSchema)
	.query(async ({ ctx, input }) => {
		const address = input.path.address || ctx.session.address

		const queryParams = buildQueryParams({
			currency: input.query.currency,
			"filter[chain_ids]": input.query?.filter?.chainIds,
			"filter[collections_ids]": input.query?.filter?.collectionIds,
			sort: input.query.sort,
			include: input.query.include.join(","),
			"page[size]": input.query?.page?.size,
			"page[after]": input.query?.page?.after
		})

		return zerionApi(() => zerion.get(`/wallets/${address}/nft-positions/${queryParams}`))
	})
