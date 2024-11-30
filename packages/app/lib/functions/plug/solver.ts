import { TRPCError } from "@trpc/server"

import axios from "axios"

import { env } from "@/env"
import { ActionSchemas } from "@/lib/types"

let cachedSchemas: Record<string, ActionSchemas | undefined> = {}

export const getSchemas = async (protocol?: string, action?: string): Promise<ActionSchemas> => {
	const cacheKey = `${protocol}-${action}`

	if (cachedSchemas[cacheKey]) return cachedSchemas[cacheKey]

	const url = `${env.SOLVER_URL}/intent`

	const response = await axios.get(url, {
		params: {
			protocol,
			action
		}
	})

	if (response.status !== 200) throw new TRPCError({ code: "INTERNAL_SERVER_ERROR" })

	cachedSchemas[cacheKey] = response.data as ActionSchemas

	return response.data
}
