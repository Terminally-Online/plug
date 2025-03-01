import { DateRange } from "react-day-picker"

import { frequencies } from "@/lib"
import { connectedChains } from "@/contexts"

export type Options = {
	key: string
	value: string
	name: string
	label: string
	icon: { default: string; secondary: string;[key: string]: string }
	info?: { label: string, value: string }
}[]

export type ActionSchema = {
	metadata: {
		icon: string
		tags: Array<string>
		chains: Array<{
			name: string,
			chainIds: [typeof connectedChains[number]['id']],
			explorer: string,
			icon: { default: string; secondary: string;[key: string]: string }
		}>
	}
	schema: Record<
		string,
		{
			type: string
			sentence: string
			options?: Record<string, Options | Record<string, Options>>
		}
	>
}

export type ActionSchemas = {
	[protocol: string]: ActionSchema
}


type InputValue = { key: string; value: string; name: string } & Partial<Options[number]> | undefined
type InputValues = Record<string, InputValue>

export type Action = {
	protocol: string
	action: string
	id: number
	values: InputValues
}
export type Actions = Array<Action>



export type Schedule = {
	date: DateRange | undefined
	repeats: (typeof frequencies)[0]
}

export type Transfer = {
	recipient?: string
	percentage?: number
	precise?: string
	tokenId?: string
}

export type Run = {
	id: string
	status: string
	result?: string
	error?: string
	errors?: string[]
	gasEstimate?: number
	data: {
		raw: Uint8Array
		decoded?: unknown
	}
	intentId?: string
	intent?: Intent
	createdAt: string
}

export type Intent = {
	id: string
	status: string
	chainId: number
	from: string
	inputs: Actions
	frequency: number
	startAt: string
	endAt?: string
	periodEndAt?: string | null
	nextSimulationAt?: string | null
	runs: Array<Run>
	createdAt: string
}

export type IntentResponse<TDecoded extends {} = Record<string, unknown>> = {
	status: { success: boolean }
	transactions: Array<{ to: string, data: string, value: number, gas: number, meta: any }>
	intent: Intent
	simulation: {
		status: string
		from: string
		to: string
		value: number
		gasEstimate: number
		resultData: { raw?: string, decoded?: TDecoded }
	}
	transaction: {
		from: `0x${string}`
		to: `0x${string}`
		chainId: number
		value: bigint
		data: `0x${string}`
		gas: number
	}
}
