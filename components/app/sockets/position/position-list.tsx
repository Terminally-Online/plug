import { FC, HTMLAttributes, useMemo, useState } from "react"

import { motion, MotionProps } from "framer-motion"
import { SearchIcon } from "lucide-react"

import { Animate, Callout, PositionFrame, Search, SocketPositionItem } from "@/components"
import { useSockets } from "@/contexts"
import { cn } from "@/lib"

export const SocketPositionList: FC<HTMLAttributes<HTMLDivElement> & { id: string; expanded?: boolean }> = ({
	id,
	expanded,
	className,
	...props
}) => {
	const { isAnonymous: anonymous, positions } = useSockets()
	const { protocols } = positions

	const [search, handleSearch] = useState("")

	const visibilePositions = useMemo(() => {
		if (protocols === undefined) return Array(3).fill(undefined)

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
		<div className={cn("flex h-full flex-col gap-2", className)} {...props}>
			<Callout.Anonymous viewing="positions" />

			{anonymous === false && (
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

			{visibilePositions.map((protocol, index) => {
				return <PositionFrame key={index} id={id} protocol={protocol} />
			})}
		</div>
	)
}
