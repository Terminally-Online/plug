import { FC, HTMLAttributes, memo, useMemo, useState } from "react"

import { Coins, SearchIcon } from "lucide-react"

import { Search } from "@/components/app/inputs/search"
import { SocketTokenItem } from "@/components/app/sockets/tokens/token-item"
import { Callout } from "@/components/app/utils/callout"
import { cn, NATIVE_TOKEN_ADDRESS, useDebounce } from "@/lib"
import { PLACEHOLDER_TOKENS } from "@/lib/constants/placeholder/tokens"
import { api } from "@/server/client"
import { useSocket } from "@/state/authentication"
import { Header } from "../../layout/header"
import { Animate } from "../../utils/animate"
import { Counter } from "@/components/shared/utils/counter"

export const SocketTokenList: FC<
	HTMLAttributes<HTMLDivElement> & {
		index: number
		address?: string
		isExpanded?: boolean
		count?: number
		isColumn?: boolean
	}
> = memo(({ index, address, isExpanded = false, count = 5, isColumn = true, className, ...props }) => {
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
	const [expanded, setExpanded] = useState<boolean>(isExpanded)
	const [hovering, setHovering] = useState<string | undefined>()

	const visibleTokens = useMemo(() => {
		if (search !== "" && tokens.length === 0) return Array(5).fill(undefined)

		const isEmptyResults = (search === "" && tokens.length == 0)
		const isPlaceholder = (!tokens || isAnonymous || isEmptyResults)

		if (isPlaceholder) return !isColumn ? PLACEHOLDER_TOKENS.slice(0, 5) : PLACEHOLDER_TOKENS

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
	}, [isAnonymous, tokens, expanded, count, isColumn, debouncedSearch, search])

	const visibleValue = useMemo(() => {
		let value = 0;
		for (const token of visibleTokens) {
			value += token.attributes?.value || 0;
		}
		return value;
	}, [visibleTokens])
	const value = useMemo(() => {
		let value = 0;
		for (const token of tokens) {
			value += token.attributes?.value || 0;
		}
		return value;
	}, [tokens])

	return (
		<div className={cn("flex flex-col gap-2", className)} {...props}>
			{!isColumn && <Header
				variant="frame"
				icon={<Coins size={14} className="opacity-40" />}
				label={<div className="flex w-full justify-between items-center">
					<p className="font-bold">Tokens</p>
					<p
						className="font-bold text-xs flex gap-1 opacity-40"
						onMouseEnter={() => setHovering("visible")}
						onMouseLeave={() => setHovering(undefined)}
					>
						{hovering === "visible" ? (<>
							$<Counter
								count={(visibleValue).toLocaleString(
									"en-US",
									{
										minimumFractionDigits: 2,
										maximumFractionDigits: 2
									}
								)}
								decimals={2}
							/>
							<span className="opacity-40">/</span>
							$<Counter
								count={(value).toLocaleString(
									"en-US",
									{
										minimumFractionDigits: 2,
										maximumFractionDigits: 2
									}
								)}
								decimals={2}
							/>
						</>) : (<>
							<Counter count={hovering ? hovering : visibleTokens.length} />
							<span className="opacity-40">/</span>
							<Counter count={tokens.length} />
						</>)}

					</p>
				</div>}
				nextOnClick={tokens.length > count ? () => setExpanded(prev => !prev) : undefined}
				nextLabel={expanded ? "See Less" : "See All"}
			/>}

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

			<Animate.List>
				{visibleTokens.map((token, tokenIndex) => (
					<Animate.ListItem key={tokenIndex}>
						<SocketTokenItem
							index={index}
							token={token}
							onMouseEnter={() => setHovering((tokenIndex + 1).toString())}
							onMouseLeave={() => setHovering(undefined)}
						/>
					</Animate.ListItem>
				))}
			</Animate.List>

			<Callout.Anonymous index={index} viewing="tokens" isAbsolute={true} />
			<Callout.EmptyAssets
				index={index}
				isEmpty={isColumn && !isAnonymous && search === "" && tokens.length === 0}
				isViewing="tokens"
				isReceivable
			/>
		</div >
	)
})

SocketTokenList.displayName = "SocketTokenList"
