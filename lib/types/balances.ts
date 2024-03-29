import { TOKENS } from "@/lib/tokens"

export type Search = {
	query: string
	isSearching: boolean
	asset: (typeof TOKENS)[0] | undefined
}
