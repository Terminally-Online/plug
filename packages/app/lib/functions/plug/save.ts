import { TRPCError } from "@trpc/server"

import axios from "axios"

import { env } from "@/env"
import { IntentResponseIntent } from "@/lib/types"

type CreateIntentProps = Omit<IntentResponseIntent, "id" | "nextSimulationAt" | "periodEndAt" | "runs" | "createdAt">
type GetIntentProps = { id?: string; address?: string; addresses?: string[] }
type IntentIdProps = { id: string }

const save = async <TData>(method: "get" | "post" | "delete", path: string, input?: unknown) => {
	const url = `${env.SOLVER_URL}/solver/save${path ? `/${path}` : ""}`
	const config = {
		headers: {
			"X-Api-Key": env.SOLVER_API_KEY
		}
	}

	try {
		let response
		switch (method) {
			case "post":
				response = await axios.post<TData>(url, input ?? {}, config)
				break
			case "get":
				response = await axios[method]<TData>(url, config)
				break
			case "delete":
				response = await axios[method]<TData>(url, config)
				break
			default:
				throw new TRPCError({ code: "METHOD_NOT_SUPPORTED" })
		}

		if (response.status !== 200) throw new TRPCError({ code: "BAD_REQUEST" })

		return response.data
	} catch (error) {
		throw new TRPCError({
			code: "INTERNAL_SERVER_ERROR",
			cause: error
		})
	}
}

export const createIntent = async (props: CreateIntentProps) => await save<IntentResponseIntent>("post", "", props)
export const getIntent = async ({ id, address, addresses }: GetIntentProps): Promise<Array<IntentResponseIntent>> => {
	if (!id && !address && !addresses) throw new TRPCError({ code: "BAD_REQUEST" })

	return await save("get", id ? `/${id}` : addresses ? `/${addresses.join(",")}` : `/${address}`)
}
export const toggleIntent = async ({ id }: IntentIdProps): Promise<IntentResponseIntent> => await save("post", `/${id}`)
export const deleteIntent = async ({ id }: IntentIdProps): Promise<IntentResponseIntent> =>
	await save("delete", `/${id}`)
export const toggleIntentStatus = async ({ id }: IntentIdProps): Promise<IntentResponseIntent> =>
	await save("post", `/${id}/status`)
