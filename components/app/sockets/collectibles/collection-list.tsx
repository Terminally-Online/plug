import { FC, HTMLAttributes, useMemo, useState } from "react"

import { motion, MotionProps } from "framer-motion"
import { SearchIcon } from "lucide-react"

import { Search, SocketCollectionItem } from "@/components"
import { useBalances } from "@/contexts"
import { cn } from "@/lib"

export const SocketCollectionList: FC<
	HTMLAttributes<HTMLDivElement> &
		MotionProps & {
			id: string
		}
> = ({ id, className, ...props }) => {
	const { collectibles } = useBalances()

	const [search, handleSearch] = useState("")

	const visibleCollectibles = useMemo(() => {
		if (collectibles === undefined) return Array(5).fill(undefined)

		const filteredCollectibles = collectibles.filter(
			collectible =>
				collectible.name.toLowerCase().includes(search.toLowerCase()) ||
				collectible.description
					.toLowerCase()
					.includes(search.toLowerCase()) ||
				collectible.collection
					.toLowerCase()
					.includes(search.toLowerCase()) ||
				collectible.collectibles.some(
					collectionCollectible =>
						(collectionCollectible.name ?? "")
							.toLowerCase()
							.includes(search.toLowerCase()) ||
						(collectionCollectible.description ?? "")
							.toLowerCase()
							.includes(search.toLowerCase())
				)
		)

		return filteredCollectibles
	}, [search, collectibles])

	return (
		<div className={cn("flex h-full flex-col gap-2", className)} {...props}>
			<Search
				className="mb-2"
				icon={<SearchIcon size={14} className="opacity-40" />}
				placeholder="Search collectibles"
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
				{visibleCollectibles.map((collection, index) => (
					<SocketCollectionItem
						key={index}
						id={id}
						collection={collection}
						searched={search !== ""}
					/>
				))}
			</motion.div>
		</div>
	)
}
