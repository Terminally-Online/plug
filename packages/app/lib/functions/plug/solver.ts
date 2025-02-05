import { TRPCError } from "@trpc/server"

import axios from "axios"

import { env } from "@/env"
import { ActionSchemas } from "@/lib/types"

let cachedSchemas: Record<string, ActionSchemas | undefined> = {}

export const schemas = async (protocol?: string, action?: string, chainId: number = 8453): Promise<ActionSchemas> => {
	const cacheKey = `${protocol}-${action}`

	if (cachedSchemas[cacheKey]) return cachedSchemas[cacheKey]

	const response = await axios.get(`${env.SOLVER_URL}/solver`, {
		params: {
			protocol,
			action,
			chainId
		},
		headers: {
			'X-Api-Key': env.SOLVER_API_KEY
		}
	})

	if (response.status !== 200) throw new TRPCError({ code: "INTERNAL_SERVER_ERROR" })

	cachedSchemas[cacheKey] = response.data as ActionSchemas

	return response.data
}

export const intent = async (input: {
	chainId: number
	from: string
	inputs: Array<{
		protocol: string
		action: string
		[key: string]: string | number
	}>
}) => {
	const response = await axios.post(`${env.SOLVER_URL}/solver`, input)

	if (response.status !== 200) throw new TRPCError({ code: "INTERNAL_SERVER_ERROR" })

	return response.data
}

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
