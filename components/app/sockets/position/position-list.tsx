import { FC, HTMLAttributes, useMemo, useState } from "react"

import { SearchIcon } from "lucide-react"

import { api, RouterOutputs } from "@/server/client"

import { Animate, Callout, PositionFrame, Search, SocketPositionItem } from "@/components"
import { cn } from "@/lib"
import { useColumns, useSocket } from "@/state"

export const SocketPositionList: FC<
	HTMLAttributes<HTMLDivElement> & {
		index: number
		positions?: RouterOutputs["socket"]["balances"]["positions"]
		expanded?: boolean
		isColumn?: boolean
	}
> = ({ index, positions, expanded, isColumn = true, className, ...props }) => {
	const { socket, isAnonymous } = useSocket()
	const { column, isExternal } = useColumns(index)
	const { data: columnPositions, isLoading } = api.socket.balances.positions.useQuery(
		column?.viewAs?.socketAddress ?? socket?.socketAddress,
		{
			enabled:
				(isExternal && column?.viewAs?.socketAddress !== undefined) ||
				(socket !== undefined &&
					socket.socketAddress !== undefined &&
					socket.id.startsWith("anonymous") === false)
		}
	)

	console.log("columnPositions", columnPositions)

	const { protocols } = positions ?? columnPositions ?? { protocols: [] }

	const [search, handleSearch] = useState("")

	const visibilePositions = useMemo(() => {
		if (protocols === undefined || (search === "" && protocols.length === 0)) return Array(5).fill(undefined)

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

	if (protocols === undefined) return null

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

			{isLoading && <p>Loading...</p>}

			<Callout.EmptySearch
				isEmpty={search !== "" && visibilePositions.length === 0}
				search={search}
				handleSearch={handleSearch}
			/>

			<Animate.List>
				{visibilePositions.map((protocol, index) => (
					<Animate.ListItem key={index}>
						<SocketPositionItem index={index} protocol={protocol} />
					</Animate.ListItem>
				))}
			</Animate.List>

			<Callout.Anonymous index={index} viewing="positions" isAbsolute={true} />
			<Callout.EmptyAssets
				isEmpty={!isAnonymous && search === "" && protocols.length === 0}
				isViewing="positions"
				isReceivable={false}
			/>

			{visibilePositions
				.filter(protocol => Boolean(protocol))
				.map((protocol, protocolIndex) => {
					return <PositionFrame key={protocolIndex} index={index} protocol={protocol} />
				})}
		</div>
	)
}
