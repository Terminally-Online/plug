import { InfinityIcon } from "lucide-react"
import { AbiParameter } from "viem"

import aave from "./aave"
import fraxlend from "./fraxlend"
import nouns from "./nouns"
import plug from "./plug"
import uniswap from "./uniswap"

type Options = Array<
	| Array<{ label: string; value: string }>
	| Record<string | `0x${string}`, Array<{ label: string; value: string }>>
>

type Action = {
	// Interaction metadata used to build the transaction.
	address: `0x${string}`
	abi: string
	inputs: Array<AbiParameter>
	options?: Options | undefined

	// Display metadata inside the application.
	sentence: string
	info: string
	icon: typeof InfinityIcon
	primary?: boolean
}

export const actions: Record<string, Record<string, Action>> = {
	plug,
	nouns,
	fraxlend,
	aave,
	uniswap
}
