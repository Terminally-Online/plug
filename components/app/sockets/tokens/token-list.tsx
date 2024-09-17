import { FC, HTMLAttributes, useMemo, useState } from "react"

import { SearchIcon } from "lucide-react"

import { api, RouterOutputs } from "@/server/client"

import { Animate, Callout, Search, SocketTokenItem, TokenFrame } from "@/components"
import { cn } from "@/lib"
import { useColumns, useSocket } from "@/state"

export const SocketTokenList: FC<
	HTMLAttributes<HTMLDivElement> & {
		index: number
		positions?: RouterOutputs["socket"]["balances"]["positions"]
		expanded?: boolean
		count?: number
		isColumn?: boolean
	}
> = ({ index, positions, expanded, count = 5, isColumn = true, className, ...props }) => {
	const { socket, isAnonymous } = useSocket()
	const { column, isExternal } = useColumns(index)
	const { data: columnPositions } = api.socket.balances.positions.useQuery(
		column?.viewAs?.socketAddress ?? socket?.socketAddress,
		{
			enabled:
				(isExternal && column?.viewAs?.socketAddress !== undefined) ||
				(socket !== undefined &&
					socket.socketAddress !== undefined &&
					socket.id.startsWith("anonymous") === false)
		}
	)

	const { tokens } = positions ?? columnPositions ?? { tokens: [] }

	const [search, handleSearch] = useState("")

	const visibleTokens = useMemo(() => {
		if (tokens === undefined || (search === "" && tokens.length === 0)) return Array(5).fill(undefined)

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
	}, [tokens, expanded, count, search])

	return (
		<div className={cn("relative flex h-full flex-col gap-2", className)} {...props}>
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
				{visibleTokens.map((token, index) => (
					<Animate.ListItem key={index}>
						<SocketTokenItem index={index} token={token} />
					</Animate.ListItem>
				))}
			</Animate.List>

			<Callout.Anonymous index={index} viewing="tokens" isAbsolute={true} />
			<Callout.EmptyAssets
				isEmpty={!isAnonymous && search === "" && tokens.length === 0}
				isViewing="tokens"
				isReceivable={true}
			/>

			{visibleTokens
				.filter(token => Boolean(token))
				.map((token, index) => {
					return <TokenFrame key={index} index={index} symbol={token.symbol} />
				})}
		</div>
	)
}
