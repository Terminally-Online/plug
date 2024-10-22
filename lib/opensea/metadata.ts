import { getDominantColor } from "@/server/color"
import { db } from "@/server/db"
import { TRPCError } from "@trpc/server"

import axios from "axios"

import { getAPIKey } from "@/lib"

export const getMetadataForToken = async ({
	address,
	chain,
	tokenId
}: {
	address: string
	chain: string
	tokenId: string
}) => {
	const collectible = await db.collectible.findFirst({
		where: {
			tokenId,
			collectionAddress: address,
			collectionChain: chain
		},
		include: {
			collectibleMetadata: true,
			collection: true
		}
	})

	if (!collectible) {
		throw new TRPCError({ code: "NOT_FOUND", message: "Collectible not found" })
	}

	if (collectible.collectibleMetadata) {
		return collectible.collectibleMetadata
	}

	const url = `https://api.opensea.io/api/v2/chain/${chain}/contract/${address}/nfts/${tokenId}`
	const response = await axios.get(url, {
		headers: {
			Accept: "application/json",
			"x-api-key": getAPIKey()
		}
	})

	if (response.status !== 200) throw new TRPCError({ code: "INTERNAL_SERVER_ERROR" })

	const traits = response.data.nft.traits === null ? [] : response.data.nft.traits
	const colorUrl = collectible.previewUrl ?? collectible.collection.iconUrl ?? ""
	const color = await getDominantColor(colorUrl)

	const metadata = await db.collectibleMetadata.create({
		data: {
			tokenId,
			collectionAddress: address,
			collectionChain: chain,
			traits,
			color,
			collectible: {
				connect: {
					cacheId_tokenId_collectionAddress_collectionChain: {
						cacheId: collectible.cacheId,
						tokenId,
						collectionAddress: address,
						collectionChain: chain
					}
				}
			}
		}
	})

	return metadata
}
