import { FC, HTMLAttributes, useMemo, useState } from "react"

import { SearchIcon } from "lucide-react"

import { PositionFrame } from "@/components/app/frames/assets/position/frame"
import { Search } from "@/components/app/inputs/search"
import { SocketPositionItem } from "@/components/app/sockets/position/position-item"
import { Callout } from "@/components/app/utils/callout"
import { cn } from "@/lib"
import { RouterOutputs } from "@/server/client"
import { useSocket } from "@/state/authentication"
import { useHoldings } from "@/state/positions"

export const SocketPositionList: FC<
	HTMLAttributes<HTMLDivElement> & {
		index: number
		columnProtocols?: RouterOutputs["socket"]["balances"]["positions"]["protocols"]
		expanded?: boolean
		isColumn?: boolean
	}
> = ({ index, columnProtocols, expanded, isColumn = true, className, ...props }) => {
	const { isAnonymous, socket } = useSocket()
	const { protocols: apiProtocols } = useHoldings(socket?.socketAddress)

	const protocols = columnProtocols ?? apiProtocols

	const [search, handleSearch] = useState("")

	const visibilePositions = useMemo(() => {
		if (isAnonymous || protocols === undefined || (search === "" && protocols.length === 0))
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
	}, [isAnonymous, expanded, protocols, search])

	if (protocols === undefined) return null

	return (
		<div className={cn("flex flex-col gap-2", className)} {...props}>
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
				isEmpty={search !== "" && visibilePositions.length === 0}
				search={search}
				handleSearch={handleSearch}
			/>

			<div className="flex flex-col gap-2">
				{visibilePositions.map((protocol, positionIndex) => (
					<SocketPositionItem key={positionIndex} index={index} protocol={protocol} />
				))}
			</div>

			<Callout.Anonymous index={index} viewing="positions" isAbsolute={true} />
			<Callout.EmptyAssets
				index={index}
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
