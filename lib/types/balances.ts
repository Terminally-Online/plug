import { tokens } from "@/lib/constants/tokens"

export type Search = {
	query: string
	isSearching: boolean
	asset: (typeof tokens)[0] | undefined
}
