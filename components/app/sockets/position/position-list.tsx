import { FC, HTMLAttributes, useMemo, useState } from "react"

import { SearchIcon } from "lucide-react"

import { Animate, Callout, PositionFrame, Search, SocketPositionItem } from "@/components"
import { useSockets } from "@/contexts"
import { cn } from "@/lib"

export const SocketPositionList: FC<
	HTMLAttributes<HTMLDivElement> & { id: string; expanded?: boolean; isColumn?: boolean }
> = ({ id, expanded, isColumn = true, className, ...props }) => {
	const { isAnonymous, isExternal, positions } = useSockets(id)
	const { protocols } = positions

	const [search, handleSearch] = useState("")

	const visibilePositions = useMemo(() => {
		if (
			(isAnonymous && isExternal === false) ||
			protocols === undefined ||
			(search === "" && protocols.length === 0)
		)
			return Array(5).fill(undefined)

		const filteredProtocols = protocols.filter(
			protocol =>
				protocol.name.toLowerCase().includes(search.toLowerCase()) ||
				protocol.positions.some(
					position =>
						position.fungible.name.toLowerCase().includes(search.toLowerCase()) ||
						position.fungible.symbol.toLowerCase().includes(search.toLowerCase()) ||
						position.fungible.implementations.some(implementation =>
							implementation.contract.toLowerCase().includes(search.toLowerCase())
						)
				)
		)

		if (expanded) return filteredProtocols

		return filteredProtocols.slice(0, 3)
	}, [expanded, protocols, search])

	if (positions === undefined) return null

	return (
		<div className={cn("relative flex h-full flex-col gap-2", className)} {...props}>
			{(isAnonymous === false || isExternal) && isColumn && protocols.length > 0 && (
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
				isEmpty={search !== "" && visibilePositions.length === 0}
				search={search}
				handleSearch={handleSearch}
			/>

			<Animate.List>
				{visibilePositions.map((protocol, index) => (
					<Animate.ListItem key={index}>
						<SocketPositionItem id={id} protocol={protocol} />
					</Animate.ListItem>
				))}
			</Animate.List>

			<Callout.Anonymous id={id} viewing="positions" isAbsolute={true} />
			<Callout.EmptyAssets
				isEmpty={!isAnonymous && search === "" && protocols.length === 0}
				isViewing="positions"
				isReceivable={false}
			/>

			{visibilePositions
				.filter(protocol => Boolean(protocol))
				.map((protocol, index) => {
					return <PositionFrame key={index} id={id} protocol={protocol} />
				})}
		</div>
	)
}
