import { FC, HTMLAttributes, memo, useMemo, useState } from "react"

import { SearchIcon } from "lucide-react"

import { CollectibleFrame } from "@/components/app/frames/assets/collectible/frame"
import { Search } from "@/components/app/inputs/search"
import { SocketCollectionItem } from "@/components/app/sockets/collectibles/collection-item"
import { Callout } from "@/components/app/utils/callout"
import { cn } from "@/lib"
import { RouterOutputs } from "@/server/client"
import { useSocket } from "@/state/authentication"
import { useHoldings } from "@/state/positions"

export const SocketCollectionList: FC<
	HTMLAttributes<HTMLDivElement> & {
		index: number
		columnCollectibles?: RouterOutputs["socket"]["balances"]["collectibles"]
		expanded?: boolean
		count?: number
		isColumn?: boolean
	}
> = memo(({ index, columnCollectibles, expanded, count = 5, isColumn = true, className, ...props }) => {
	const { isAnonymous, socket } = useSocket()
	const { collectibles: apiCollectibles } = useHoldings(socket?.socketAddress)

	const collectibles = columnCollectibles ?? apiCollectibles

	const [search, handleSearch] = useState("")

	const visibleCollectibles: RouterOutputs["socket"]["balances"]["collectibles"] | Array<undefined> = useMemo(() => {
		if (collectibles === undefined || isAnonymous || (search === "" && collectibles.length === 0))
			return Array(5).fill(undefined)

		const filteredCollectibles = collectibles.filter(
			collectible =>
				collectible.name.toLowerCase().includes(search.toLowerCase()) ||
				collectible.description.toLowerCase().includes(search.toLowerCase()) ||
				collectible.collectibles.some(collectionCollectible =>
					(collectionCollectible.name ?? "").toLowerCase().includes(search.toLowerCase())
				)
		)

		if (expanded) return filteredCollectibles

		return filteredCollectibles.slice(0, count)
	}, [isAnonymous, collectibles, expanded, count, search])

	return (
		<div className={cn("flex flex-col gap-2", className)} {...props}>
			{isAnonymous === false && isColumn && collectibles && collectibles.length > 0 && (
				<Search
					className="mb-2"
					icon={<SearchIcon size={14} className="opacity-40" />}
					placeholder="Search collectibles"
					search={search}
					handleSearch={handleSearch}
					clear
				/>
			)}

			<Callout.EmptySearch
				isEmpty={search !== "" && visibleCollectibles.length === 0}
				search={search}
				handleSearch={handleSearch}
			/>

			<div className="flex flex-col gap-2">
				{visibleCollectibles.map((collection, collectionIndex) => (
					<SocketCollectionItem
						key={collectionIndex}
						index={index}
						collection={collection}
						searched={false}
					/>
				))}
			</div>

			<Callout.Anonymous index={index} viewing="collectibles" isAbsolute={true} />
			<Callout.EmptyAssets
				index={index}
				isEmpty={!isAnonymous && search === "" && collectibles.length === 0}
				isViewing="collectibles"
				isReceivable={true}
			/>

			{visibleCollectibles.map(
				collection =>
					collection &&
					collection.collectibles.map(collectible => (
						<CollectibleFrame
							key={`${collection.address}-${collection.chain}-${collectible.tokenId}`}
							index={index}
							collection={collection}
							collectible={collectible}
						/>
					))
			)}
		</div>
	)
})

SocketCollectionList.displayName = "SocketCollectionList"
