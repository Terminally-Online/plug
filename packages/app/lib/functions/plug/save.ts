import { TRPCError } from "@trpc/server"

import axios from "axios"

import { env } from "@/env"
import { Intent } from "@/lib/types"

type CreateIntentProps = Omit<Intent, "id" | "nextSimulationAt" | "periodEndAt" | "runs">
type GetIntentProps = { id?: string, address?: string }
type IntentIdProps = { id: string }

const save = async <TData>(method: "get" | "post" | "delete", path: string, input?: unknown) => {
	const url = `${env.SOLVER_URL}/solver/save${path ? `/${path}` : ''}`
	const config = {
		headers: {
			'X-Api-Key': env.SOLVER_API_KEY
		}
	}

	let response
	switch (method) {
		case "post":
			response = await axios.post<TData>(url, input ?? {}, config) 
			break
		case "get":
		case "delete":
			response = await axios[method]<TData>(url, config)
			break
		default:
			throw new TRPCError({ code: "BAD_REQUEST" })
	}

	if (response.status !== 200) throw new TRPCError({ code: "INTERNAL_SERVER_ERROR" })

	return response.data
}

export const createIntent = async (props: CreateIntentProps) => await save<Intent>("post", "", props)
export const getIntent = async ({ id, address }: GetIntentProps): Promise<Array<Intent>> => { 
	if (!id && !address) throw new TRPCError({ code: "BAD_REQUEST" })

	return await save("get", id ? `/${id}` : `/${address}`)
}
export const toggleIntent = async ({ id }: IntentIdProps): Promise<Intent> => await save("post", `/${id}/toggle`)
export const deleteIntent = async ({ id }: IntentIdProps): Promise<Intent> => await save("delete", `/${id}`)
