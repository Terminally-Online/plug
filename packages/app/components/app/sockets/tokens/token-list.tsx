import { FC, HTMLAttributes, memo, useMemo, useState } from "react"

import { SearchIcon } from "lucide-react"

import { Search } from "@/components/app/inputs/search"
import { SocketTokenItem } from "@/components/app/sockets/tokens/token-item"
import { Callout } from "@/components/app/utils/callout"
import { cn, useDebounce } from "@/lib"
import { api, RouterOutputs } from "@/server/client"
import { useSocket } from "@/state/authentication"
import { useHoldings } from "@/state/positions"

type Tokens = RouterOutputs["socket"]["balances"]["positions"]["tokens"] | RouterOutputs["solver"]["tokens"]["get"]

export const SocketTokenList: FC<
	HTMLAttributes<HTMLDivElement> & {
		index: number
		columnTokens?: Tokens
		expanded?: boolean
		count?: number
		isColumn?: boolean
	}
> = memo(({ index, columnTokens, expanded, count = 5, isColumn = true, className, ...props }) => {
	const { isAnonymous, socket } = useSocket()
	const { tokens: apiTokens } = useHoldings(socket?.socketAddress)

	const tokens = columnTokens ?? apiTokens

	const [search, debouncedSearch, handleSearch] = useDebounce("")

	const { data: searchedTokens } = api.solver.tokens.get.useQuery(debouncedSearch, {
		enabled: search !== "" && expanded,
		keepPreviousData: true
	})

	const visibleTokens = useMemo(() => {
		if (isAnonymous || tokens === undefined || (search === "" && tokens.length === 0)) {
			console.log("returning empty array")
			return Array(5).fill(undefined)
		}

		const filteredTokens = tokens.filter(
			token =>
				token.name.toLowerCase().includes(debouncedSearch.toLowerCase()) ||
				token.symbol.toLowerCase().includes(debouncedSearch.toLowerCase()) ||
				token.implementations.some(implementation =>
					implementation.contract.toLowerCase().includes(debouncedSearch.toLowerCase())
				)
		)

		if (!expanded) return filteredTokens.slice(0, count)
		if (!debouncedSearch) return filteredTokens

		const ownedSymbols = filteredTokens.map(token => token.symbol)
		const unownedTokens = searchedTokens?.filter(token => !ownedSymbols.includes(token.symbol))

		return filteredTokens.concat(unownedTokens ?? [])
	}, [isAnonymous, tokens, expanded, count, debouncedSearch, searchedTokens, search])

	return (
		<div className={cn("flex flex-col gap-2", className)} {...props}>
			{isAnonymous === false && isColumn && tokens.length > 0 && (
				<Search
					className="mb-2"
					icon={<SearchIcon size={14} className="opacity-40" />}
					placeholder="Search tokens"
					search={search}
					handleSearch={handleSearch}
					clear
				/>
			)}

			<Callout.EmptySearch
				isEmpty={search !== "" && visibleTokens.length === 0}
				search={search}
				handleSearch={handleSearch}
			/>

			<div className="flex flex-col gap-2">
				{visibleTokens.map((token, tokenIndex) => (
					<SocketTokenItem key={`${tokenIndex}-${token?.symbol}`} index={index} token={token} />
				))}
			</div>

			<Callout.Anonymous index={index} viewing="tokens" isAbsolute={true} />
			<Callout.EmptyAssets index={index} isEmpty={tokens.length === 0} isViewing="tokens" isReceivable={true} />
		</div>
	)
})

SocketTokenList.displayName = "SocketTokenList"
