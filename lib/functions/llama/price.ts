import axios from "axios"

import { db } from "@/server/db"

const PRICE_CACHE_TIME = 3 * 60 * 1000

export const getPriceKey = (chain: string, address: string) => `${chain.toLowerCase()}:${address}`

export const getPrices = async (queries: string[]) => {
	const cachedPrices = await db.tokenPrice.findMany({
		where: { id: { in: queries } }
	})

	const uncachedQueries = queries.filter(query => {
		const cache = cachedPrices.find(price => price.id === query)

		return cache === undefined || cache.updatedAt < new Date(Date.now() - PRICE_CACHE_TIME)
	})

	if (uncachedQueries.length === 0) return cachedPrices

	const query = uncachedQueries.join(",")

	const currentPricesResponse = await axios.get(`https://coins.llama.fi/chart/${query}?span=48&period=30m&searchWidth=1200`)

	if (currentPricesResponse.status !== 200) return cachedPrices

	const currentPrices: Record<
		`${string}:${string}`,
		Partial<{
			decimals: number
			symbol: string
			prices: Array<{
				timestamp: number
				price: number
			}>
			confidence: number
		}>
	> = currentPricesResponse.data.coins ?? {}

	if (Object.keys(currentPrices).length === 0) return cachedPrices

	const transformed = Object.entries(currentPrices).map(([key, price]) => {
		const { prices, ...data } = price

		const [chain, address] = key.split(":") as [string, string]

		const start = prices?.[0].price
		const end = prices?.[prices.length - 1].price
		const change = start && end ? ((end - start) / start) * 100 : undefined

		const timestamp = prices?.[prices.length - 1].timestamp

		return {
			...data,
			id: key,
			chain,
			address,
			timestamp,
			price: end,
			change
		}
	})

	await Promise.all(
		transformed.map(async price => {
			await db.tokenPrice.upsert({
				where: { id: price.id },
				create: price,
				update: price
			})
		})
	)

	return await db.tokenPrice.findMany({
		where: { id: { in: queries } },
		select: { id: true, price: true, change: true }
	})
}
