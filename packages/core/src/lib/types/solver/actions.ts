import { InfinityIcon } from "lucide-react"
import { AbiParameter } from "viem"

export type Options = Array<
	| Array<{ label: string; value: string } | undefined>
	| Record<
			string | `0x${string}`,
			Array<{ label: string; value: string } | undefined>
	  >
	| undefined
>

export type Action = {
	address: `${number}:0x${string}` | Array<`${number}:0x${string}`>
	abi: string
	inputs: Array<AbiParameter>
	sentence: string
	info: string
	icon: typeof InfinityIcon
	options?: Options
	primary?: boolean
}

export type ActionProvider = {
	info: {
		image: string
		gradient: [`#${string}`, `#${string}`]
		tags: Array<string>
	}
	actions: Record<string, Action>
}

export type Actions = Record<string, Record<string, ActionProvider>>
