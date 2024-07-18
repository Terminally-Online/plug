import { tokens } from "@/lib"

export type Search = {
	query: string
	isSearching: boolean
	asset: (typeof tokens)[0] | undefined
}
