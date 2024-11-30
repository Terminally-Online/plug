import { FC, HTMLAttributes, memo, useMemo, useState } from "react"

import { SearchIcon } from "lucide-react"

import { Callout, Search, SocketTokenItem } from "@/components"
import { cn } from "@/lib"
import { RouterOutputs } from "@/server/client"
import { useHoldings, useSocket } from "@/state"

export const SocketTokenList: FC<
	HTMLAttributes<HTMLDivElement> & {
		index: number
		columnTokens?: RouterOutputs["socket"]["balances"]["positions"]["tokens"]
		expanded?: boolean
		count?: number
		isColumn?: boolean
	}
> = memo(({ index, columnTokens, expanded, count = 5, isColumn = true, className, ...props }) => {
	const { isAnonymous, socket } = useSocket()
	const { tokens: apiTokens } = useHoldings(socket?.socketAddress)

	const tokens = columnTokens ?? apiTokens

	const [search, handleSearch] = useState("")

	const visibleTokens = useMemo(() => {
		if (isAnonymous || tokens === undefined || (search === "" && tokens.length === 0))
			return Array(5).fill(undefined)

		const filteredTokens = tokens.filter(
			token =>
				token.name.toLowerCase().includes(search.toLowerCase()) ||
				token.symbol.toLowerCase().includes(search.toLowerCase()) ||
				token.implementations.some(implementation =>
					implementation.contract.toLowerCase().includes(search.toLowerCase())
				)
		)

		if (expanded) return filteredTokens

		return filteredTokens.slice(0, count)
	}, [isAnonymous, tokens, expanded, count, search])

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
					<SocketTokenItem key={tokenIndex} index={index} token={token} />
				))}
			</div>

			<Callout.Anonymous index={index} viewing="tokens" isAbsolute={true} />
			<Callout.EmptyAssets
				index={index}
				isEmpty={!isAnonymous && search === "" && tokens.length === 0}
				isViewing="tokens"
				isReceivable={true}
			/>
		</div>
	)
})

SocketTokenList.displayName = "SocketTokenList"
