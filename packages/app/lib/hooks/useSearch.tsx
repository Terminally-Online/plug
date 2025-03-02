import { useDebounce } from "./useDebounce"

export const useSearch = () => {
	const [search, debouncedSearch, handleSearch] = useDebounce("")
	const [tag, debouncedTag, handleTag] = useDebounce("")

	const handleReset = () => {
		handleSearch("")
		handleTag("")
	}

	return {
		search,
		debouncedSearch,
		handleSearch,
		tag,
		debouncedTag, 
		handleTag,
		handleReset
	}
}
