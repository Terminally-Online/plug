import { useState } from "react"

export const useSearch = () => {
	const [search, handleSearch] = useState("")
	const [tag, handleTag] = useState("")

	const handleReset = () => {
		handleSearch("")
		handleTag("")
	}

	return {
		search,
		handleSearch,
		tag,
		handleTag,
		handleReset
	}
}
