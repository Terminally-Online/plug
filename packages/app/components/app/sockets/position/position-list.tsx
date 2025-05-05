import { FC, HTMLAttributes, useMemo, useState } from "react"

import { SearchIcon } from "lucide-react"

import { Search } from "@/components/app/inputs/search"
import { SocketPositionItem } from "@/components/app/sockets/position/position-item"
import { Callout } from "@/components/app/utils/callout"
import { cn } from "@/lib"
import { api } from "@/server/client"
import { useSocket } from "@/state/authentication"
import { PLACEHOLDER_POSITIONS } from "@/lib/constants/placeholder/positions"

export const SocketPositionList: FC<
	HTMLAttributes<HTMLDivElement> & {
		index: number
		address?: string
		isColumn?: boolean
		expanded?: boolean
	}
> = ({ index, address, expanded, isColumn = true, className, ...props }) => {
	const { isAnonymous, socket } = useSocket()

	const { data } = api.service.zerion.wallet.positions.useQuery({
		path: { address: address || socket?.socketAddress },
		query: {
			aggregate: true,
			filter: {
				positions: "only_complex"
			}
		}
	}, {
		enabled: !isAnonymous,
		placeholderData: prev => prev
	})
	const protocols = useMemo(() => data?.data || [], [data])

	const [search, handleSearch] = useState("")

	const visibilePositions = useMemo(() => {
		if (search !== "" && protocols.length === 0)
			return Array(5).fill(undefined)

		const isEmptyResults = (search === "" && protocols.length == 0)
		const isPlaceholder = isColumn && (!protocols || isAnonymous || isEmptyResults)

		if (isPlaceholder) return PLACEHOLDER_POSITIONS
		if (search === "") return protocols

		const filteredProtocols = protocols.filter(
			protocol =>
				protocol.attributes.name.toLowerCase().includes(search.toLowerCase())
			// protocol.positions.some(
			// 	position =>
			// 		position.fungible.name.toLowerCase().includes(search.toLowerCase()) ||
			// 		position.fungible.symbol.toLowerCase().includes(search.toLowerCase()) ||
			// 		position.fungible.implementations.some(implementation =>
			// 			implementation.contract.toLowerCase().includes(search.toLowerCase())
			// 		)
			// )
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
				isEmpty={isColumn && !isAnonymous && search === "" && protocols.length === 0}
				isViewing="positions"
				isReceivable={false}
			/>
		</div>
	)
}
