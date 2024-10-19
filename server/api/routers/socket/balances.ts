import { getDominantColor } from "@/server/color"
import { TRPCError } from "@trpc/server"

import axios from "axios"
import { z } from "zod"

import { getAPIKey} from "@/lib"
import { getCollectibles, getPositions } from "@/lib/functions/zerion"

import { anonymousProtectedProcedure, createTRPCRouter } from "../../trpc"

export const balances = createTRPCRouter({
	collectibles: anonymousProtectedProcedure
		.input(z.string().optional())
		.query(async ({ input, ctx }) => await getCollectibles(ctx.session.address, input)),
	positions: anonymousProtectedProcedure
		.input(z.string().optional())
		.query(async ({ input, ctx }) => await getPositions(ctx.session.address, input)),
	metadata: anonymousProtectedProcedure
		.input(
			z.object({
				type: z.union([z.literal("ERC20"), z.literal("ERC721"), z.literal("ERC1155")]),
				address: z.string(),
				tokenId: z.string(),
				chain: z.string()
			})
		)
		.query(async ({ input, ctx }) => {
			if (input.type === "ERC20") throw new TRPCError({ code: "NOT_IMPLEMENTED" })

			const metadataCache = await ctx.db.collectibleMetadata.findUnique({
				where: {
					tokenId_collectionAddress_collectionChain: {
						tokenId: input.tokenId,
						collectionAddress: input.address,
						collectionChain: input.chain
					}
				}
			})

			if (metadataCache) return metadataCache

			const url = `https://api.opensea.io/api/v2/chain/${input.chain}/contract/${input.address}/nfts/${input.tokenId}`
			const response = await axios.get(url, {
				headers: {
					Accept: "application/json",
					"x-api-key": getAPIKey()
				}
			})

			if (response.status !== 200) throw new TRPCError({ code: "INTERNAL_SERVER_ERROR" })

			const collectible = await ctx.db.collectible.findFirst({
				where: {
					tokenId: input.tokenId,
					collectionAddress: input.address,
					collectionChain: input.chain
				},
				include: { collection: true }
			})

			if (collectible === null) throw new TRPCError({ code: "NOT_FOUND" })

			const traits = response.data.nft.traits === null ? [] : response.data.nft.traits
			const colorUrl = collectible.previewUrl ?? collectible.collection.iconUrl ?? ""
			const color = await getDominantColor(colorUrl)

			return await ctx.db.collectibleMetadata.upsert({
				where: {
					tokenId_collectionAddress_collectionChain: {
						tokenId: input.tokenId,
						collectionAddress: input.address,
						collectionChain: input.chain
					}
				},
				create: {
					tokenId: input.tokenId,
					collectionAddress: input.address,
					collectionChain: input.chain,
					traits,
					color
				},
				update: {
					traits,
					color
				}
			})
		})
})
