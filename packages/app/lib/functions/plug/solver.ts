import { TRPCError } from "@trpc/server"

import axios from "axios"

import { env } from "@/env"
import { IntentResponse, KillResponse, SchemasResponse } from "@/lib/types"
import { formTestnet } from "viem/chains"

let cachedSchemas: Record<string, SchemasResponse | undefined> = {}

export const schemas = async (
	protocol?: string,
	action?: string,
	chainId = 8453,
	search: Array<string> = [],
	from?: string
): Promise<SchemasResponse> => {
	const params: Record<string, any> = {
		protocol,
		action,
		from,
		chainId,
	}

	// For gorilla/schema decoder to correctly decode into a slice of structs,
	// we need to format the URL parameters like this: search.0.index=0&search.0.value=ethereum
	search.forEach((item) => {
		const [rawKey, val] = item.split("=")
		if (val !== "") {
			// Extract the index from the format "search[key]"
			const indexMatch = rawKey.match(/search\[(\d+)\]/)
			if (indexMatch && indexMatch[1]) {
				const index = parseInt(indexMatch[1])
				params[`search.${index}.index`] = index
				params[`search.${index}.value`] = val
			}
		}
	})

	const cacheKey = JSON.stringify(params)

	if (cachedSchemas[cacheKey]) return cachedSchemas[cacheKey]

	const response = await axios.get(`${env.SOLVER_URL}/solver`, {
		params,
		headers: {
			"X-Api-Key": env.SOLVER_API_KEY
		}
	})

	if (response.status !== 200) throw new TRPCError({ code: "INTERNAL_SERVER_ERROR" })

	cachedSchemas[cacheKey] = response.data as SchemasResponse

	return response.data
}
export const getIntentSchemas = schemas

export const intent = async (input: {
	chainId: number
	from: string
	inputs: Array<{
		protocol: string
		action: string
		[key: string]: string | number
	}>
}) => {
	const response = await axios.post(`${env.SOLVER_URL}/solver`, input, {
		headers: {
			"X-Api-Key": env.SOLVER_API_KEY
		}
	})

	if (response.status !== 200) throw new TRPCError({ code: "INTERNAL_SERVER_ERROR" })

	return response.data as IntentResponse
}
export const getIntentTransaction = intent

export const killed = async () => {
	const response = await axios.get(`${env.SOLVER_URL}/solver/kill`, {
		headers: {
			"X-Api-Key": env.SOLVER_API_KEY
		}
	})

	if (response.status !== 200) throw new TRPCError({ code: "INTERNAL_SERVER_ERROR" })

	return response.data as { killed: boolean }
}

export const kill = async () => {
	const response = await axios.post(
		`${env.SOLVER_URL}/solver/kill`,
		{},
		{
			headers: {
				"X-Api-Key": env.SOLVER_API_KEY
			}
		}
	)

	if (response.status !== 200) throw new TRPCError({ code: "INTERNAL_SERVER_ERROR" })

	return response.data as KillResponse
}
