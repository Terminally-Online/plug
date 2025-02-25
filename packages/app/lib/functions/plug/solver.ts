import { TRPCError } from "@trpc/server"

import axios from "axios"

import { env } from "@/env"
import { ActionSchemas, Intent } from "@/lib/types"

let cachedSchemas: Record<string, ActionSchemas | undefined> = {}

export const schemas = async (
	protocol?: string,
	action?: string,
	chainId = 8453,
	search: Array<string> = [],
	from?: string
): Promise<ActionSchemas> => {
	const params = {
		protocol,
		action,
		from,
		chainId,
		...search.reduce((acc, value) => {
			const [key, val] = value.split("=")
			if (val === "") return acc
			return { ...acc, [key]: val }
		}, {})
	}

	const cacheKey = JSON.stringify(params)

	if (cachedSchemas[cacheKey]) return cachedSchemas[cacheKey]

	const response = await axios.get(`${env.SOLVER_URL}/solver`, {
		params,
		headers: {
			'X-Api-Key': env.SOLVER_API_KEY
		}
	})

	if (response.status !== 200) throw new TRPCError({ code: "INTERNAL_SERVER_ERROR" })

	cachedSchemas[cacheKey] = response.data as ActionSchemas

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
			'X-Api-Key': env.SOLVER_API_KEY
		}
	})

	if (response.status !== 200) throw new TRPCError({ code: "INTERNAL_SERVER_ERROR" })

	return response.data
}
export const getIntentTransaction = intent





export const killed = async () => {
	const response = await axios.get(`${env.SOLVER_URL}/solver/kill`, {
		headers: {
			'X-Api-Key': env.SOLVER_API_KEY
		}
	})

	if (response.status !== 200) throw new TRPCError({ code: "INTERNAL_SERVER_ERROR" })

	return response.data as { killed: boolean }
}

export const kill = async () => {
	const response = await axios.post(`${env.SOLVER_URL}/solver/kill`, {}, {
		headers: {
			'X-Api-Key': env.SOLVER_API_KEY
		}
	})

	if (response.status !== 200) throw new TRPCError({ code: "INTERNAL_SERVER_ERROR" })

	return response.data as { killed: boolean }
}
