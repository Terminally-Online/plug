import axios from "axios"
import { z } from "zod"

import { TRPCError } from "@trpc/server"

import { getAPIKey, getCollectibles } from "@/lib"
import { getPositions } from "@/lib/functions/zerion"
import { getDominantColor } from "@/server/color"

import { anonymousProtectedProcedure, createTRPCRouter } from "../../trpc"

export const balances = createTRPCRouter({
	collectibles: anonymousProtectedProcedure.input(z.string().optional()).query(async ({ input }) => {
		if (input === undefined) throw new TRPCError({ code: "BAD_REQUEST" })
		return await getCollectibles(input)
	}),
	positions: anonymousProtectedProcedure.input(z.string().optional()).query(async ({ input }) => {
		try {
			if (input === undefined) throw new TRPCError({ code: "BAD_REQUEST" })
			return await getPositions(input)
		} catch (error) {
			console.error(error)
			throw new TRPCError({ code: "BAD_REQUEST" })
		}
	}),
	metadata: anonymousProtectedProcedure
		.input(
			z.object({
				type: z.union([z.literal("ERC20"), z.literal("ERC721"), z.literal("ERC1155")]),
				id: z.string()
			})
		)
		.query(async ({ input, ctx }) => {
			if (input.type === "ERC20") throw new TRPCError({ code: "NOT_IMPLEMENTED" })

			const metadataCache = await ctx.db.openseaCollectibleMetadata.findUnique({
				where: { collectibleId: input.id }
			})

			if (metadataCache) return metadataCache

			const collectible = await ctx.db.openseaCollectible.findUnique({
				where: { id: input.id },
				include: { collection: true }
			})

			if (collectible === null) throw new TRPCError({ code: "NOT_FOUND" })

			const url = `https://api.opensea.io/api/v2/chain/${collectible.collection.chain}/contract/${collectible.contract}/nfts/${collectible.identifier}`
			const response = await axios.get(url, {
				headers: {
					Accept: "application/json",
					"x-api-key": getAPIKey()
				}
			})

			if (response.status !== 200) throw new TRPCError({ code: "INTERNAL_SERVER_ERROR" })

			const traits = response.data.nft.traits === null ? [] : response.data.nft.traits
			const color = await getDominantColor(collectible.displayImageUrl ?? collectible.collection.imageUrl)

			if (metadataCache !== null)
				return await ctx.db.openseaCollectibleMetadata.update({
					where: { collectibleId: input.id },
					data: {
						traits,
						color
					}
				})

			return await ctx.db.openseaCollectibleMetadata.create({
				data: {
					traits,
					color,
					collectibleId: input.id
				}
			})
		})
})
