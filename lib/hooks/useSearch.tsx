import { useState } from "react"

import useDebounce from "./useDebounce"

export const useSearch = () => {
	const [search, debouncedSearch, handleSearch] = useDebounce("")
	const [tag, handleTag] = useState("")

	const handleReset = () => {
		handleSearch("")
		handleTag("")
	}

	return {
		search,
		debouncedSearch,
		handleSearch,
		tag,
		handleTag,
		handleReset
	}
}
