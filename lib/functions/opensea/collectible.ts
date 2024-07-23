import axios from "axios"

import { OpenseaCollection } from "@prisma/client"

import { OpenseaCollectible } from "@/lib/types"
import { db } from "@/server/db"

type Collectibles = Array<OpenseaCollectible>

const getAPIKey = () => {
	const keys = process.env.OPENSEA_API_KEY?.split(",")
	return keys?.[Math.floor(Math.random() * keys.length)]
}

export const getCollection = async (
	slug: string,
	chain: string
): Promise<OpenseaCollection> => {
	const collection = await db.openseaCollection.findUnique({
		where: { slug }
	})

	if (collection) return collection

	const response = await axios.get(
		`https://api.opensea.io/api/v2/collections/${slug}`,
		{
			headers: {
				Accept: "application/json",
				"x-api-key": getAPIKey()
			}
		}
	)

	const data = response.data

	const transformed = {
		collection: data.collection,
		name: data.name,
		description: data.description,
		imageUrl: data.image_url,
		bannerImageUrl: data.banner_image_url,
		owner: data.owner,
		category: data.category,
		isDisabled: data.is_disabled,
		isNsfw: data.is_nsfw,
		traitOffersEnabled: data.trait_offers_enabled,
		collectionOffersEnabled: data.collection_offers_enabled,
		openseaUrl: data.opensea_url,
		projectUrl: data.project_url,
		wikiUrl: data.wiki_url,
		discordUrl: data.discord_url,
		telegramUrl: data.telegram_url,
		twitterUsername: data.twitter_username,
		instagramUsername: data.instagram_username,
		totalSupply: data.total_supply,
		chain,
		createdDate: new Date(data.created_date)
	}

	return await db.openseaCollection.upsert({
		where: { slug },
		create: {
			slug,
			...transformed
		},
		update: {
			...transformed
		}
	})
}

export const getCollectiblesForChain = async (
	address: string,
	chain: string,
	limit = 200,
	next?: string,
	collectibles: Collectibles = []
): Promise<Collectibles> => {
	const response = await axios.get(
		`https://api.opensea.io/api/v2/chain/${chain}/account/${address}/nfts?limit=${limit}${next ? `&next=${next}` : ""}`,
		{
			headers: {
				Accept: "application/json",
				"x-api-key": getAPIKey()
			}
		}
	)

	const responseCollectibles = await Promise.all(
		response.data.nfts.map(async (nft: OpenseaCollectible) => {
			return {
				...nft,
				collection: await getCollection(
					nft.collection as unknown as string,
					chain
				)
			}
		})
	)

	collectibles = [...collectibles, ...responseCollectibles]

	return response.data.next
		? getCollectiblesForChain(
				address,
				chain,
				limit,
				response.data.next,
				collectibles
			)
		: collectibles
}

export const getCollectibles = async (
	address: string,
	chains = ["ethereum", "optimism", "base"]
) => {
	const responses = await Promise.allSettled(
		chains.map(chain => getCollectiblesForChain(address, chain))
	)

	const collectibles = responses.flatMap(response => {
		if (response.status === "fulfilled") {
			return response.value
		}
		return []
	})

	const cleanedCollectibles = collectibles
		.filter(nft => nft.is_disabled === false)
		.sort(
			(a, b) =>
				new Date(b.updated_at).getTime() -
				new Date(a.updated_at).getTime()
		)

	// move collection out of the collectible into the collectibles object
	const groupedCollectibles = cleanedCollectibles.reduce(
		(acc, collectible) => {
			const slug = collectible.collection.slug

			if (!acc[slug])
				acc[slug] = {
					...collectible.collection,
					collectibles: []
				}

			const { collection, ...data } = collectible

			acc[slug].collectibles.push(data)

			return acc
		},
		{} as Record<
			string,
			OpenseaCollection & {
				collectibles: Array<Omit<OpenseaCollectible, "collection">>
			}
		>
	)

	return groupedCollectibles
}
