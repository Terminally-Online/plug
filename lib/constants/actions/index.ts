import { AbiParameter } from "viem"

import { InfinityIcon } from "lucide-react"

import aave from "./aave"
import chainlink from "./chainlink"
import fraxlend from "./fraxlend"
import nouns from "./nouns"
import pendle from "./pendle"
import plug from "./plug"
import uniswap from "./uniswap"

type Options = Array<
	| Array<{ label: string; value: string } | undefined>
	| Record<string | `0x${string}`, Array<{ label: string; value: string } | undefined>>
	| undefined
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
