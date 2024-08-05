import { db } from "@/server/db"

const PRICE_CACHE_TIME = 3 * 60 * 1000

export const getPriceKey = (chain: string, address: string) =>
	`${chain.toLowerCase()}:${address}`

export const getPrices = async (queries: string[]) => {
	const cachedPrices = await db.tokenPrice.findMany({
		where: { id: { in: queries } }
	})

	const uncachedQueries = queries.filter(query => {
		const cache = cachedPrices.find(price => price.id === query)

		return (
			cache === undefined ||
			cache.updatedAt < new Date(Date.now() - PRICE_CACHE_TIME)
		)
	})

	if (uncachedQueries.length === 0) return cachedPrices

	const query = uncachedQueries.join(",")

	const currentPricesResponse = await fetch(
		`https://coins.llama.fi/prices/current/${query}`
	)

	if (currentPricesResponse.ok === false) return cachedPrices
	const currentPricesJson = await currentPricesResponse.json()
	const currentPrices =
		(currentPricesJson.coins as Record<
			`${string}:${string}`,
			Partial<{
				decimals: number
				symbol: string
				price: number
				timestamp: number
				confidence: number
				change: number
			}>
		>) ?? {}

	if (Object.keys(currentPrices).length === 0) return cachedPrices

	const percentagesResponse = await fetch(
		`https://coins.llama.fi/percentage/${query}`
	)

	let percentages: Record<string, number> = {}
	if (percentagesResponse.ok === false) return cachedPrices

	const percentagesJson = await percentagesResponse.json()
	percentages = percentagesJson.coins ?? {}

	const transformed = Object.entries(currentPrices).map(([key, price]) => {
		const [chain, address] = key.split(":") as [string, string]

		return {
			...price,
			id: key,
			chain,
			address,
			change: percentages[key] as number | undefined
		}
	})

	await db.tokenPrice.deleteMany({ where: { id: { in: uncachedQueries } } })
	await db.tokenPrice.createMany({
		data: transformed
	})

	return await db.tokenPrice.findMany({
		where: { id: { in: queries } }
	})
}
