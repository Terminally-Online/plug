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
					collectionsIds: z.array(z.number()).optional()
				})
				.optional(),
			sort: z.enum(["-floor_price", "floor_price", "created_at", "-created_at"]).optional(),
			include: z.array(z.string()).optional(),
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
				amount: z.number(),
				price: z.number().nullable(),
				value: z.number().nullable(),
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
								.nullable(),
							detail: z
								.object({
									url: z.string(),
									content_type: z.string().optional()
								})
								.nullable(),
							audio: z
								.object({
									url: z.string(),
									content_type: z.string().optional()
								})
								.nullable(),
							video: z
								.object({
									url: z.string(),
									content_type: z.string().optional()
								})
								.nullable()
						})
						.nullable(),
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
				})
			}),
			included: z
				.array(
					z.union([
						// NFT Item
						z.object({
							type: z.literal("nfts"),
							id: z.string(),
							attributes: z.object({
								contract_address: z.string(),
								token_id: z.string(),
								interface: z.string(),
								metadata: z
									.object({
										name: z.string().nullable(),
										description: z.string().nullable(),
										tags: z.array(z.string()).optional(),
										content: z
											.object({
												preview: z
													.object({
														url: z.string(),
														content_type: z.string().optional()
													})
													.nullable(),
												detail: z
													.object({
														url: z.string(),
														content_type: z.string().optional()
													})
													.nullable(),
												audio: z
													.object({
														url: z.string(),
														content_type: z.string().optional()
													})
													.nullable(),
												video: z
													.object({
														url: z.string(),
														content_type: z.string().optional()
													})
													.nullable()
											})
											.nullable(),
										attributes: z
											.array(
												z.object({
													key: z.string(),
													value: z.string()
												})
											)
											.optional()
									})
									.nullable(),
								market_data: z
									.object({
										prices: z
											.object({
												floor: z.number().nullable()
											})
											.optional(),
										last_sale: z
											.object({
												price: z.number().nullable(),
												quantity: z
													.object({
														int: z.string(),
														decimals: z.number(),
														float: z.number(),
														numeric: z.string()
													})
													.optional()
											})
											.optional()
									})
									.optional(),
								external_links: z
									.array(
										z.object({
											type: z.string(),
											name: z.string(),
											url: z.string()
										})
									)
									.optional(),
								flags: z
									.object({
										is_spam: z.boolean().optional()
									})
									.optional()
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
								nft_collection: z.object({
									data: z.object({
										type: z.string(),
										id: z.union([z.string(), z.number()])
									})
								})
							})
						}),
						// NFT Collection
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
						}),
						// Wallet NFT Collection
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
							})
						})
					])
				)
				.optional()
		})
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
			"filter[collections_ids]": input.query?.filter?.collectionsIds,
			sort: input.query.sort,
			include: input.query.include,
			"page[size]": input.query?.page?.size,
			"page[after]": input.query?.page?.after
		})

		return zerionApi(() => zerion.get(`/wallets/${address}/nft-positions/${queryParams}`))
	})
