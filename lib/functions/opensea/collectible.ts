import axios from "axios"

import { db } from "@/server/db"
import { TRPCError } from "@trpc/server"

import { OpenseaCollectible } from "@/lib/types"

const MINUTE = 60 * 1000
const HOUR = 60 * MINUTE
const COLLECTION_CACHE_TIME = 24 * HOUR * 7
const COLLECTIBLES_CACHE_TIME = 1 * HOUR

export const getAPIKey = () => {
	const keys = process.env.OPENSEA_API_KEY?.split(",")
	return keys?.[Math.floor(Math.random() * keys.length)]
}

export const getOpenseaCollection = async (slug: string, chain: string) => {
	const cachedCollection = await db.openseaCollection.findUnique({
		where: { slug }
	})

	const cache = cachedCollection && cachedCollection.updatedAt > new Date(Date.now() - COLLECTION_CACHE_TIME)

	if (cache) return cachedCollection

	const response = await axios.get(`https://api.opensea.io/api/v2/collections/${slug}`, {
		headers: {
			Accept: "application/json",
			"x-api-key": getAPIKey()
		}
	})

	if (response.status !== 200)
		throw new TRPCError({
			code: "INTERNAL_SERVER_ERROR",
			message: "An upstream service is unavailable."
		})

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
		createdAt: new Date(data.created_date)
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

export const getOpenseaCollectiblesForChain = async (
	address: string,
	chain: string,
	limit = 200,
	next?: string,
	collectibles: Array<OpenseaCollectible> = []
): Promise<Array<OpenseaCollectible>> => {
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
		response.data.nfts.map(async (collectible: OpenseaCollectible) => ({
			...collectible,
			collection: await getOpenseaCollection(collectible.collection as unknown as string, chain)
		}))
	)

	collectibles = [...collectibles, ...responseCollectibles]

	return response.data.next
		? getOpenseaCollectiblesForChain(address, chain, limit, response.data.next, collectibles)
		: collectibles
}

export const getCollectiblesForChain = async (address: string, chain: string, limit = 200, next?: string) => {
	const socket = await db.userSocket.findFirst({
		where: { socketAddress: address }
	})
	if (!socket) return []

	const cachedCollectibles = await db.openseaCollectibleCache.findUnique({
		where: { socketId_chain: { socketId: socket.id, chain } },
		include: { collectibles: { include: { collection: true } } }
	})

	if (cachedCollectibles && cachedCollectibles?.updatedAt > new Date(Date.now() - COLLECTIBLES_CACHE_TIME))
		return cachedCollectibles.collectibles

	const newCollectibles = await getOpenseaCollectiblesForChain(address, chain, limit, next)

	const existingCollectibles = await db.openseaCollectible.findMany({
		where: { owner: address, cacheChain: chain },
		select: { id: true, identifier: true, collectionSlug: true }
	})

	const newCollectiblesMap = new Map(
		newCollectibles.map(c => [
			`${c.identifier}:${c.collection.slug}`,
			{
				identifier: c.identifier,
				collectionSlug: c.collection.slug,
				contract: c.contract,
				tokenStandard: c.token_standard,
				name: c.name,
				description: c.description,
				imageUrl: c.image_url,
				displayImageUrl: c.display_image_url,
				displayAnimationUrl: c.display_animation_url,
				metadataUrl: c.metadata_url,
				openseaUrl: c.opensea_url,
				updatedAt: new Date(c.updated_at),
				isDisabled: c.is_disabled,
				isNsfw: c.is_nsfw,
				owner: address,
				cacheChain: chain
			}
		])
	)

	const existingCollectiblesMap = new Map(existingCollectibles.map(c => [`${c.identifier}:${c.collectionSlug}`, c]))

	const toDelete = existingCollectibles.filter(c => !newCollectiblesMap.has(`${c.identifier}:${c.collectionSlug}`))
	const toCreate = Array.from(newCollectiblesMap.values()).filter(
		c => !existingCollectiblesMap.has(`${c.identifier}:${c.collectionSlug}`)
	)
	const toUpdate = Array.from(newCollectiblesMap.values()).filter(c =>
		existingCollectiblesMap.has(`${c.identifier}:${c.collectionSlug}`)
	)

	await db.$transaction(async prisma => {
		if (toDelete.length)
			await prisma.openseaCollectible.deleteMany({
				where: { id: { in: toDelete.map(c => c.id) } }
			})

		await prisma.openseaCollectibleCache.upsert({
			where: { socketId_chain: { socketId: socket.id, chain } },
			create: { chain, socketId: socket.id, updatedAt: new Date() },
			update: { updatedAt: new Date() }
		})

		if (toCreate.length)
			await prisma.openseaCollectible.createMany({
				data: toCreate.map(c => ({
					...c,
					cacheSocketId: socket.id,
					cacheChain: chain
				}))
			})

		for (const collectible of toUpdate) {
			await prisma.openseaCollectible.updateMany({
				where: {
					identifier: collectible.identifier,
					collectionSlug: collectible.collectionSlug,
					cacheChain: chain,
					owner: address
				},
				data: collectible
			})
		}
	})

	return Array.from(newCollectiblesMap.values())
}

export const getCollectibles = async (address: string, chains = ["ethereum", "optimism", "base"]) => {
	const socket = await db.userSocket.findFirst({
		where: { socketAddress: address }
	})

	if (socket === null) return []

	await Promise.all(chains.map(async chain => await getCollectiblesForChain(address, chain)))

	return await db.openseaCollection.findMany({
		where: { collectibles: { some: { cacheSocketId: socket.id } } },
		include: {
			collectibles: {
				where: { cacheSocketId: socket.id },
				orderBy: { updatedAt: "desc" }
			}
		},
		orderBy: { createdAt: "desc" }
	})
}
