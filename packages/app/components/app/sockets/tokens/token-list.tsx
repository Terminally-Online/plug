import { FC, HTMLAttributes, memo, useMemo } from "react"

import { SearchIcon } from "lucide-react"

import { Search } from "@/components/app/inputs/search"
import { SocketTokenItem } from "@/components/app/sockets/tokens/token-item"
import { Callout } from "@/components/app/utils/callout"
import { cn, useDebounce } from "@/lib"
import { PLACEHOLDER_TOKENS } from "@/lib/constants/placeholder/tokens"
import { api } from "@/server/client"
import { useSocket } from "@/state/authentication"

import { ErrorFrame } from "../../frames/plugs/[id]/execute/error"

export const SocketTokenList: FC<
	HTMLAttributes<HTMLDivElement> & {
		index: number
		address?: string
		expanded?: boolean
		count?: number
		isColumn?: boolean
	}
> = memo(({ index, address, expanded, count = 5, isColumn = true, className, ...props }) => {
	const { isAnonymous, socket } = useSocket()

	const { data } = api.service.zerion.wallet.positions.useQuery(
		{
			path: { address: address || socket?.socketAddress },
			query: {
				aggregate: true
			}
		},
		{ enabled: !isAnonymous, retry: false, placeholderData: prev => prev }
	)
	const positions = data?.data || []
	const tokens = useMemo(() => positions.filter(pos => pos.attributes.position_type === "wallet"), [positions])

	const [search, debouncedSearch, handleSearch] = useDebounce("")

	// const { data: searchedTokens } = api.solver.tokens.get.useQuery(debouncedSearch, {
	// 	enabled: search !== "" && expanded,
	// 	placeholderData: prev => prev
	// })

	const visibleTokens = useMemo(() => {
		return tokens
		// if (search !== "" && tokens.length === 0) {
		// 	return Array(5).fill(undefined)
		// }
		//
		// if (isAnonymous || tokens === undefined || (search === "" && tokens.length === 0)) {
		// 	return PLACEHOLDER_TOKENS
		// }
		//
		// const filteredTokens = tokens.filter(
		// 	token =>
		// 		token.name.toLowerCase().includes(debouncedSearch.toLowerCase()) ||
		// 		token.symbol.toLowerCase().includes(debouncedSearch.toLowerCase()) ||
		// 		token.implementations.some(implementation =>
		// 			implementation.contract.toLowerCase().includes(debouncedSearch.toLowerCase())
		// 		)
		// )
		//
		// if (!expanded) return filteredTokens.slice(0, count)
		// if (!debouncedSearch) return filteredTokens
		//
		// const ownedSymbols = filteredTokens.map(token => token.symbol)
		// const unownedTokens = searchedTokens?.filter(token => !ownedSymbols.includes(token.symbol))
		//
		// return filteredTokens.concat(unownedTokens ?? [])
	}, [isAnonymous, tokens, expanded, count, debouncedSearch, search])

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
					<SocketTokenItem
						key={`${tokenIndex}-${token.attributes.fungible_info.symbol}`}
						index={index}
						token={token}
					/>
				))}
			</div>

			<Callout.Anonymous index={index} viewing="tokens" isAbsolute={true} />
			<Callout.EmptyAssets index={index} isEmpty={tokens.length === 0} isViewing="tokens" isReceivable={true} />

			<ErrorFrame index={index} />
		</div>
	)
})

SocketTokenList.displayName = "SocketTokenList"
