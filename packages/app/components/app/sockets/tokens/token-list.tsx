import { FC, HTMLAttributes, memo, useMemo } from "react"

import { SearchIcon } from "lucide-react"

import { Search } from "@/components/app/inputs/search"
import { SocketTokenItem } from "@/components/app/sockets/tokens/token-item"
import { Callout } from "@/components/app/utils/callout"
import { cn, NATIVE_TOKEN_ADDRESS, useDebounce } from "@/lib"
import { PLACEHOLDER_TOKENS } from "@/lib/constants/placeholder/tokens"
import { api } from "@/server/client"
import { useSocket } from "@/state/authentication"

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
		{ enabled: !isAnonymous, placeholderData: prev => prev }
	)
	const positions = useMemo(() => data?.data || [], [data])
	const tokens = useMemo(() => positions.filter(pos => pos.attributes.position_type === "wallet"), [positions])

	const [search, debouncedSearch, handleSearch] = useDebounce("")

	const visibleTokens = useMemo(() => {
		if (search !== "" && tokens.length === 0) return Array(5).fill(undefined)

		const isEmptyResults = (search === "" && tokens.length == 0)
		const isPlaceholder = isColumn && (!tokens || isAnonymous || isEmptyResults)

		if (isPlaceholder) return PLACEHOLDER_TOKENS
		if (search === "") return tokens

		const filteredTokens = tokens.filter(
			token =>
				token.attributes.fungible_info.name.toLowerCase().includes(debouncedSearch.toLowerCase()) ||
				token.attributes.fungible_info.symbol.toLowerCase().includes(debouncedSearch.toLowerCase()) ||
				token.attributes.fungible_info.implementations.some(implementation =>
					(implementation?.address ?? NATIVE_TOKEN_ADDRESS).toLowerCase().includes(debouncedSearch.toLowerCase())
				)
		)

		if (!expanded) return filteredTokens.slice(0, count)

		return filteredTokens
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
			<Callout.EmptyAssets
				index={index}
				isEmpty={isColumn && !isAnonymous && search === "" && tokens.length === 0}
				isViewing="tokens"
				isReceivable
			/>
		</div>
	)
})

SocketTokenList.displayName = "SocketTokenList"
