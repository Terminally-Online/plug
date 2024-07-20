import {
	createContext,
	FC,
	PropsWithChildren,
	useContext,
	useMemo
} from "react"

import { useDebounce } from "@/lib"
import { api } from "@/server/client"

export const BalancesContext = createContext<{
	search: string
	debouncedSearch: string
	handleSearch: (search: string) => void
}>({
	search: "",
	debouncedSearch: "",
	handleSearch: () => {}
})

export const BalancesProvider: FC<PropsWithChildren> = ({ children }) => {
	const [search, debouncedSearch, handleSearch] = useDebounce("")

	return (
		<BalancesContext.Provider
			value={{
				search,
				debouncedSearch,
				handleSearch
			}}
		>
			{children}
		</BalancesContext.Provider>
	)
}

export const useBalances = ({ address }: { address: string }) => {
	const { search, debouncedSearch, handleSearch } =
		useContext(BalancesContext)

	const { data: apiBalances } = api.socket.balances.useQuery(address)

	const balances = useMemo(() => {
		if (apiBalances === undefined) return undefined

		if (search === "") return apiBalances

		return apiBalances.filter(
			token =>
				token?.symbol.toLowerCase().includes(search.toLowerCase()) ||
				token?.name.toLowerCase().includes(search.toLowerCase())
		)
	}, [apiBalances, search])

	return {
		address,
		search,
		debouncedSearch,
		balances,
		handleSearch
	}
}
