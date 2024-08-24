import { FC, HTMLAttributes, useMemo, useState } from "react"

import { motion, MotionProps } from "framer-motion"
import { SearchIcon } from "lucide-react"

import { useBalances } from "@/contexts"
import { cn } from "@/lib"

import { PositionFrame } from "../../frames/assets/position"
import { Search } from "../../inputs"
import { SocketPositionItem } from "./position-item"

export const SocketPositionList: FC<
	HTMLAttributes<HTMLDivElement> &
		MotionProps & { id: string; expanded?: boolean }
> = ({ id, expanded, className, ...props }) => {
	const { positions } = useBalances()
	const { protocols } = positions

	const [search, handleSearch] = useState("")

	const visibilePositions = useMemo(() => {
		if (protocols === undefined) return Array(3).fill(undefined)

		const filteredProtocols = protocols.filter(
			protocol =>
				protocol.name.toLowerCase().includes(search.toLowerCase()) ||
				protocol.positions.some(
					position =>
						position.fungible.name
							.toLowerCase()
							.includes(search.toLowerCase()) ||
						position.fungible.symbol
							.toLowerCase()
							.includes(search.toLowerCase()) ||
						position.fungible.implementations.some(implementation =>
							implementation.contract
								.toLowerCase()
								.includes(search.toLowerCase())
						)
				)
		)

		if (expanded) return filteredProtocols

		return filteredProtocols.slice(0, 3)
	}, [expanded, protocols, search])

	if (positions === undefined) return null

	return (
		<div className={cn("flex h-full flex-col gap-2", className)} {...props}>
			<Search
				className="mb-2"
				icon={<SearchIcon size={14} className="opacity-40" />}
				placeholder="Search positions"
				search={search}
				handleSearch={handleSearch}
				clear
			/>

			<motion.div
				className="flex flex-col gap-2"
				initial="hidden"
				animate="visible"
				variants={{
					hidden: { opacity: 0 },
					visible: {
						opacity: 1,
						transition: {
							staggerChildren: 0.05
						}
					}
				}}
				{...(props as MotionProps)}
			>
				{visibilePositions.map(protocol => (
					<SocketPositionItem
						key={protocol}
						id={id}
						protocol={protocol}
					/>
				))}
			</motion.div>

			{visibilePositions.map((protocol, index) => {
				return <PositionFrame key={index} id={id} protocol={protocol} />
			})}
		</div>
	)
}
