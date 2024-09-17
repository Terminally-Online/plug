import { FC, HTMLAttributes, useMemo, useState } from "react"

import { SearchIcon } from "lucide-react"

import { api, RouterOutputs } from "@/server/client"

import { Animate, Callout, CollectibleFrame, Search, SocketCollectionItem } from "@/components"
import { cn } from "@/lib"
import { useColumns, useSocket } from "@/state"

export const SocketCollectionList: FC<
	HTMLAttributes<HTMLDivElement> & {
		index: number
		collectibles?: RouterOutputs["socket"]["balances"]["collectibles"]
		expanded?: boolean
		count?: number
		isColumn?: boolean
	}
> = ({ index, collectibles, expanded, count = 5, isColumn = true, className, ...props }) => {
	const { socket, isAnonymous } = useSocket()
	const { column, isExternal } = useColumns(index)
	const { data: columnCollectibles } = api.socket.balances.collectibles.useQuery(
		column?.viewAs?.socketAddress ?? socket?.socketAddress,
		{
			enabled:
				(isExternal && column?.viewAs?.socketAddress !== undefined) ||
				(socket !== undefined &&
					socket.socketAddress !== undefined &&
					socket.id.startsWith("anonymous") === false)
		}
	)

	collectibles = collectibles ?? columnCollectibles ?? []

	const [search, handleSearch] = useState("")

	const visibleCollectibles: RouterOutputs["socket"]["balances"]["collectibles"] | Array<undefined> = useMemo(() => {
		if (collectibles === undefined || (search === "" && collectibles.length === 0)) return Array(5).fill(undefined)

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
		<div className={cn("relative flex h-full flex-col gap-2", className)} {...props}>
			{(isAnonymous === false || isExternal) &&
				isColumn &&
				columnCollectibles &&
				columnCollectibles.length > 0 && (
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

			<Animate.List>
				{visibleCollectibles.map((collection, index) => (
					<Animate.ListItem key={index}>
						<SocketCollectionItem index={index} collection={collection} searched={search !== ""} />
					</Animate.ListItem>
				))}
			</Animate.List>

			<Callout.Anonymous index={index} viewing="collectibles" isAbsolute={true} />
			<Callout.EmptyAssets
				isEmpty={!isAnonymous && search === "" && collectibles.length === 0}
				isViewing="collectibles"
				isReceivable={false}
			/>

			{visibleCollectibles.map(
				(collection, index) =>
					collection &&
					collection.collectibles.map(collectible => (
						<CollectibleFrame
							key={`${index}-${collectible.identifier}`}
							index={index}
							collection={collection}
							collectible={collectible}
						/>
					))
			)}
		</div>
	)
}
