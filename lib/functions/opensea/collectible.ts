const MINUTE = 60 * 1000
const HOUR = 60 * MINUTE

export const getAPIKey = () => {
	const keys = process.env.OPENSEA_API_KEY?.split(",")
	return keys?.[Math.floor(Math.random() * keys.length)]
}

// export const getOpenseaCollection = async (slug: string, chain: string) => {
// 	const cachedCollection = await db.openseaCollection.findUnique({
// 		where: { slug }
// 	})

// 	const cache = cachedCollection && cachedCollection.updatedAt > new Date(Date.now() - COLLECTION_CACHE_TIME)

// 	if (cache) return cachedCollection

// 	const response = await axios.get(`https://api.opensea.io/api/v2/collections/${slug}`, {
// 		headers: {
// 			Accept: "application/json",
// 			"x-api-key": getAPIKey()
// 		}
// 	})

// 	if (response.status !== 200)
// 		throw new TRPCError({
// 			code: "INTERNAL_SERVER_ERROR",
// 			message: "An upstream service is unavailable."
// 		})

// 	const data = response.data

// 	const transformed = {
// 		collection: data.collection,
// 		name: data.name,
// 		description: data.description,
// 		imageUrl: data.image_url,
// 		bannerImageUrl: data.banner_image_url,
// 		owner: data.owner,
// 		category: data.category,
// 		isDisabled: data.is_disabled,
// 		isNsfw: data.is_nsfw,
// 		traitOffersEnabled: data.trait_offers_enabled,
// 		collectionOffersEnabled: data.collection_offers_enabled,
// 		openseaUrl: data.opensea_url,
// 		projectUrl: data.project_url,
// 		wikiUrl: data.wiki_url,
// 		discordUrl: data.discord_url,
// 		telegramUrl: data.telegram_url,
// 		twitterUsername: data.twitter_username,
// 		instagramUsername: data.instagram_username,
// 		totalSupply: data.total_supply,
// 		chain,
// 		createdAt: new Date(data.created_date)
// 	}

// 	return await db.openseaCollection.upsert({
// 		where: { slug },
// 		create: {
// 			slug,
// 			...transformed
// 		},
// 		update: {
// 			...transformed
// 		}
// 	})
// }
