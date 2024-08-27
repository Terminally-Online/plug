import { FC, HTMLAttributes, useMemo, useState } from "react"

import { motion, MotionProps } from "framer-motion"
import { SearchIcon } from "lucide-react"

import { Search, SocketCollectionItem } from "@/components"
import { useBalances, useSockets } from "@/contexts"
import { cn } from "@/lib"
import { RouterOutputs } from "@/server/client"

export const SocketCollectionList: FC<
	HTMLAttributes<HTMLDivElement> &
		MotionProps & {
			id: string
			collectibles?: RouterOutputs["socket"]["balances"]["collectibles"]
			expanded?: boolean
			count?: number
			column?: boolean
		}
> = ({ id, collectibles, expanded, count = 5, column = true, className, ...props }) => {
	const { anonymous } = useSockets()
	const { collectibles: apiCollectibles } = useBalances()
	collectibles = collectibles ?? apiCollectibles

	const [search, handleSearch] = useState("")

	const visibleCollectibles = useMemo(() => {
		if (collectibles === undefined) return Array(5).fill(undefined)

		const filteredCollectibles = collectibles.filter(
			collectible =>
				collectible.name.toLowerCase().includes(search.toLowerCase()) ||
				collectible.description.toLowerCase().includes(search.toLowerCase()) ||
				collectible.collection.toLowerCase().includes(search.toLowerCase()) ||
				collectible.collectibles.some(
					collectionCollectible =>
						(collectionCollectible.name ?? "").toLowerCase().includes(search.toLowerCase()) ||
						(collectionCollectible.description ?? "").toLowerCase().includes(search.toLowerCase())
				)
		)

		if (expanded) return filteredCollectibles

		return filteredCollectibles.slice(0, count)
	}, [collectibles, expanded, count, search])

	return (
		<div className={cn("flex h-full flex-col gap-2", className)} {...props}>
			{anonymous && (
				<div className="flex h-full flex-col items-center justify-center text-center font-bold">
					<p>You are anonymous.</p>
					<p className="max-w-[320px] opacity-40">To view the collectibles you are holding you must authenticate a wallet.</p>
				</div>
			)}

			{anonymous === false && column && (
				<Search
					className="mb-2"
					icon={<SearchIcon size={14} className="opacity-40" />}
					placeholder="Search collectibles"
					search={search}
					handleSearch={handleSearch}
					clear
				/>
			)}

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
			>
				{visibleCollectibles.map((collection, index) => (
					<SocketCollectionItem key={index} id={id} collection={collection} searched={search !== ""} />
				))}
			</motion.div>
		</div>
	)
}
