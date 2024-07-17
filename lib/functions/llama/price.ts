export type PriceData = Record<
	`${string}:${string}`,
	{
		decimals: number
		symbol: string
		price: number
		timestamp: number
		confidence: number
		change: number | undefined
	}
>

export const getPrices = async (
	coinsKey: string | string[]
): Promise<PriceData> => {
	try {
		const currentPricesResponse = await fetch(
			`https://coins.llama.fi/prices/current/${coinsKey}`
		)
		if (currentPricesResponse.ok === false) return {}
		const currentPricesJson = await currentPricesResponse.json()
		const currentPrices: PriceData = currentPricesJson.coins ?? {}

		const percentagesResponse = await fetch(
			`https://coins.llama.fi/percentage/${coinsKey}`
		)
		if (percentagesResponse.ok === false) return currentPrices
		const percentagesJson = await percentagesResponse.json()
		const percentages = percentagesJson.coins ?? {}

		return Object.fromEntries(
			Object.entries(currentPrices).map(([key, value]) => {
				return [
					key,
					{
						...value,
						change: percentages[key]
					}
				]
			})
		)
	} catch (error) {
		return {}
	}
}
