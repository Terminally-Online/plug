import type { FC, PropsWithChildren } from "react"
import { createContext, useContext, useMemo } from "react"

import { useDebounce } from "@/lib/hooks/useDebounce"
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
	const {
		debounce: handleSearch,
		value: search,
		debounced: debouncedSearch
	} = useDebounce({ initial: "" })

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
