import { FC, HTMLAttributes, useMemo, useState } from "react"

import { Coins, SearchIcon } from "lucide-react"

import { Search } from "@/components/app/inputs/search"
import { SocketPositionItem } from "@/components/app/sockets/position/position-item"
import { Callout } from "@/components/app/utils/callout"
import { Counter } from "@/components/shared/utils/counter"
import { cn } from "@/lib"
import { PLACEHOLDER_POSITIONS } from "@/lib/constants/placeholder/positions"
import { api } from "@/server/client"
import { useSocket } from "@/state/authentication"

import { Header } from "../../layout/header"
import { Animate } from "../../utils/animate"

export const SocketPositionList: FC<
	HTMLAttributes<HTMLDivElement> & {
		index: number
		address?: string
		isColumn?: boolean
		isExpanded?: boolean
	}
> = ({ index, address, isExpanded, isColumn = true, className, ...props }) => {
	const { isAnonymous, socket } = useSocket()

	const { data } = api.service.zerion.wallet.positions.useQuery(
		{
			path: { address: address || socket?.socketAddress },
			query: {
				aggregate: true,
				filter: {
					positions: "only_complex"
				}
			}
		},
		{
			enabled: !isAnonymous,
			placeholderData: prev => prev
		}
	)
	const protocols = useMemo(() => data?.data || [], [data])

	const [search, handleSearch] = useState("")
	const [expanded, setExpanded] = useState(isExpanded)
	const [hovering, setHovering] = useState<string | undefined>()

	const groupedProtocols = useMemo(() => {
		const groups: Record<string, { name: string; positions: typeof protocols }> = {}
		protocols.forEach(position => {
			const dappId = position.relationships.dapp?.data.id || "unknown"
			const protocolName = position.attributes.application_metadata?.name || dappId
			if (!groups[dappId]) {
				groups[dappId] = { name: protocolName, positions: [] }
			}
			groups[dappId].positions.push(position)
		})
		return groups
	}, [protocols])

	const visibleProtocols = useMemo(() => {
		if (search !== "" && protocols.length === 0) {
			return {}
		}

		const isEmptyResults = search === "" && protocols.length === 0
		const isPlaceholder = (!protocols || isAnonymous || isEmptyResults)

		if (isPlaceholder)
			return !isColumn ? Object.fromEntries(Object.entries(PLACEHOLDER_POSITIONS).slice(0, 3)) : PLACEHOLDER_POSITIONS


		if (search === "") {
			if (expanded) {
				return groupedProtocols
			}
			return Object.fromEntries(Object.entries(groupedProtocols).slice(0, 3))
		}

		const filteredGroups: Record<string, { name: string; positions: typeof protocols }> = {}
		Object.entries(groupedProtocols).forEach(([dappId, group]) => {
			if (group.name.toLowerCase().includes(search.toLowerCase())) {
				filteredGroups[dappId] = group
			}
		})

		if (expanded) {
			return filteredGroups
		}

		return Object.fromEntries(Object.entries(filteredGroups).slice(0, 3))
	}, [groupedProtocols, expanded, search, isColumn, isAnonymous, protocols])

	if (protocols === undefined) return null

	return (
		<div className={cn("flex flex-col gap-2", className)} {...props}>
			{!isColumn && (
				<Header
					variant="frame"
					icon={<Coins size={14} className="opacity-40" />}
					label={
						<div className="flex w-full items-center justify-between">
							<p className="font-bold">Positions</p>
							<p className="flex gap-1 text-xs font-bold opacity-40">
								<Counter count={hovering ? hovering : Object.keys(visibleProtocols).length} />
								<span className="opacity-40">/</span>
								<Counter count={Object.keys(groupedProtocols).length} />
							</p>
						</div>
					}
					nextOnClick={
						Object.keys(groupedProtocols).length > 3 ? () => setExpanded(prev => !prev) : undefined
					}
					nextLabel={expanded ? "See Less" : "See All"}
				/>
			)}

			{isAnonymous === false && isColumn && protocols.length > 0 && (
				<Search
					className="mb-2"
					icon={<SearchIcon size={14} className="opacity-40" />}
					placeholder="Search positions"
					search={search}
					handleSearch={handleSearch}
					clear
				/>
			)}

			<Callout.EmptySearch
				isEmpty={search !== "" && Object.values(visibleProtocols).length === 0}
				search={search}
				handleSearch={handleSearch}
			/>

			<Animate.List>
				{Object.entries(visibleProtocols).map(([dappId, group], idx) => (
					<Animate.ListItem key={dappId}>
						<SocketPositionItem 
							index={index} 
							protocols={group.positions} 
							onMouseEnter={() => setHovering(String(idx + 1))}
							onMouseLeave={() => setHovering(undefined)}
						/>
					</Animate.ListItem>
				))}
			</Animate.List>

			<Callout.Anonymous index={index} viewing="positions" isAbsolute={true} />
			<Callout.EmptyAssets
				index={index}
				isEmpty={isColumn && !isAnonymous && search === "" && protocols.length === 0}
				isViewing="positions"
				isReceivable={false}
			/>
		</div>
	)
}
