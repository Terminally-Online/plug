import { TRPCError } from "@trpc/server"

import axios from "axios"

import { ZerionFungibles } from "@/lib/types"

import { getZerionApiKey } from "./authentication"

type Tokens = {
	price: number
	change: number
	implementations: {
		chain: string
		contract: string
		chain_id: string
		address: string
		decimals: number
		balance?: number
	}[]
	name: string
	symbol: string
	description: string
	icon?: string
	flags: {
		verified: boolean
	}
	external_links: Array<{
		type: string
		name: string
		url: string
	}>
	market_data: {
		total_supply: number
		circulating_supply: number
		market_cap: number
		fully_diluted_valuation: number
		price: number
		changes: {
			percent_1d: number
			percent_30d: number
			percent_90d: number
			percent_365d: number
		}
	}
}[]

const SEARCH_TOKENS_CACHE_TIME = 10 * 60 * 1000
const cache = new Map<string, { timestamp: number; tokens: any[] }>()

export const getTokens = async (search: string = "", chains: string[] = ["base"]): Promise<Tokens> => {
	const cacheKey = `${chains.join(",")}-${search}`
	if (cache.has(cacheKey)) {
		const cached = cache.get(search)

		if (cached && cached.timestamp + SEARCH_TOKENS_CACHE_TIME > Date.now()) return cached.tokens
	}

	const response = await axios.get(
		`https://api.zerion.io/v1/fungibles/?currency=usd&page[size]=100&filter[search_query]=${search}&sort=-market_data.market_cap&filter[implementation_chain_id]=${chains.join(",")}`,
		{
			headers: {
				accept: "application/json",
				authorization: getZerionApiKey()
			}
		}
	)

	if (response.status !== 200) throw new TRPCError({ code: "INTERNAL_SERVER_ERROR" })

	const data: ZerionFungibles = response.data
	const tokens = data.data.map(token => ({
		...token.attributes,
		icon: token.attributes?.icon?.url || undefined,
		price: token.attributes.market_data.price,
		change: token.attributes.market_data.changes.percent_1d,
		implementations: token.attributes.implementations
			.filter(implementation => chains.includes(implementation.chain_id))
			.map(implementation => ({
				...implementation,
				chain: implementation.chain_id,
				contract: implementation.address
			}))
	}))

	cache.set(cacheKey, { timestamp: Date.now(), tokens })

	return tokens
}
