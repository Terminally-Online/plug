import { z } from "zod"

import { publicProcedure } from "@/server/api/trpc"

import { buildQueryParams, zerion, zerionApi } from "../lib"

const DetailInputSchema = z.object({
	path: z.object({
		nftId: z.string()
	}),
	query: z
		.object({
			currency: z.string().default("usd"),
			include: z.string().optional()
		})
		.optional()
		.default({})
})

const DetailOutputSchema = z.object({
	links: z.object({
		self: z.string()
	}),
	data: z.object({
		type: z.literal("nfts"),
		id: z.string(),
		attributes: z.object({
			contract_address: z.string(),
			token_id: z.string(),
			interface: z.string(),
			metadata: z
				.object({
					name: z.string().nullable(),
					description: z.string().nullable().optional(),
					tags: z.array(z.string()).optional(),
					content: z
						.object({
							preview: z
								.object({
									url: z.string(),
									content_type: z.string().optional()
								})
								.nullable()
								.optional(),
							detail: z
								.object({
									url: z.string(),
									content_type: z.string().optional()
								})
								.nullable()
								.optional(),
							audio: z
								.object({
									url: z.string(),
									content_type: z.string().optional()
								})
								.nullable()
								.optional(),
							video: z
								.object({
									url: z.string(),
									content_type: z.string().optional()
								})
								.nullable()
								.optional()
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
	included: z
		.array(
			z.object({
				type: z.literal("nft_collections"),
				id: z.union([z.string(), z.number()]),
				attributes: z.object({
					metadata: z
						.object({
							name: z.string().nullable(),
							description: z.string().nullable().optional(),
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
						.optional()
				})
			})
		)
		.optional()
})

export const detail = publicProcedure
	.input(DetailInputSchema)
	.output(DetailOutputSchema)
	.query(async ({ input }) => {
		const queryParams = buildQueryParams({
			currency: input.query.currency,
			include: input.query.include
		})

		return zerionApi(() => zerion.get(`/nfts/${input.path.nftId}${queryParams}`))
	})
