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
			id: z.string(),
			attributes: z.object({
				min_changed_at: z.string(),
				max_changed_at: z.string(),
				nfts_count: z.string(),
				total_floor_price: z.number(),
				collection_info: z.object({
					name: z.string(),
					description: z.string(),
					content: z.object({
						icon: z.object({
							url: z.string()
						}),
						banner: z
							.object({
								url: z.string(),
								content_type: z.string().optional()
							})
							.optional()
					})
				})
			}),
			relationships: z.object({
				chains: z.object({
					data: z.array(
						z.object({
							type: z.literal("chains"),
							id: z.string()
						})
					)
				}),
				nft_collection: z.object({
					data: z.object({
						type: z.literal("nft_collections"),
						id: z.string()
					})
				})
			})
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
