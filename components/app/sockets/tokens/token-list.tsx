import { FC, HTMLAttributes, useMemo, useState } from "react"

import { SearchIcon } from "lucide-react"

import { RouterOutputs } from "@/server/client"

import { Animate, Callout, Search, SocketTokenItem, TokenFrame, TransferFrame } from "@/components"
import { cn } from "@/lib"
import { useColumns, useHoldings, useSocket } from "@/state"

export const SocketTokenList: FC<
	HTMLAttributes<HTMLDivElement> & {
		index: number
		columnTokens?: RouterOutputs["socket"]["balances"]["positions"]["tokens"]
		expanded?: boolean
		count?: number
		isColumn?: boolean
	}
> = ({ index, columnTokens, expanded, count = 5, isColumn = true, className, ...props }) => {
	const { isAnonymous } = useSocket()
	const { column, isExternal } = useColumns(index)
	const { tokens: apiTokens } = useHoldings(column?.viewAs?.socketAddress)

	const tokens = columnTokens ?? apiTokens

	const [search, handleSearch] = useState("")

	const visibleTokens = useMemo(() => {
		if ((isAnonymous && isExternal === false) || tokens === undefined || (search === "" && tokens.length === 0))
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
	}, [isAnonymous, isExternal, tokens, expanded, count, search])

	return (
		<div className={cn("flex h-full flex-col gap-2", className)} {...props}>
			{(isAnonymous === false || isExternal) && isColumn && tokens.length > 0 && (
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

			<Animate.List>
				{visibleTokens.map((token, tokenIndex) => (
					<Animate.ListItem key={tokenIndex}>
						<SocketTokenItem index={index} token={token} />
					</Animate.ListItem>
				))}
			</Animate.List>

			<Callout.Anonymous index={index} viewing="tokens" isAbsolute={true} />
			<Callout.EmptyAssets
				index={index}
				isEmpty={!isAnonymous && search === "" && tokens.length === 0}
				isViewing="tokens"
				isReceivable={true}
			/>

			{visibleTokens.map((token, tokenIndex) => (
				<>
					<TokenFrame key={tokenIndex} index={index} token={token} />
					<TransferFrame key={`${tokenIndex}-transfer-send`} index={index} token={token} />
				</>
			))}
		</div>
	)
}
