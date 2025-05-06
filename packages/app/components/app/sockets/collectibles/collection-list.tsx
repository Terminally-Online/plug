import { FC, HTMLAttributes, memo, useMemo, useState } from "react"

import { Frame, SearchIcon } from "lucide-react"

import { Search } from "@/components/app/inputs/search"
import { SocketCollectionItem } from "@/components/app/sockets/collectibles/collection-item"
import { Callout } from "@/components/app/utils/callout"
import { Counter } from "@/components/shared/utils/counter"
import { cn } from "@/lib"
import { PLACEHOLDER_COLLECTIONS } from "@/lib/constants/placeholder/collectibles"
import { api } from "@/server/client"
import { useSocket } from "@/state/authentication"

import { CollectibleFrame } from "../../frames/assets/collectible/frame"
import { CollectionFrame } from "../../frames/assets/collection/frame"
import { Header } from "../../layout/header"
import { Animate } from "../../utils/animate"

export const SocketCollectionList: FC<
	HTMLAttributes<HTMLDivElement> & {
		index: number
		address?: string
		isExpanded?: boolean
		count?: number
		isColumn?: boolean
	}
> = memo(({ index, address, isExpanded = false, count = 5, isColumn = true, className, ...props }) => {
	const { isAnonymous, socket } = useSocket()

	const { data } = api.service.zerion.wallet.nftCollections.useQuery(
		{
			path: { address: address || socket?.socketAddress }
		},
		{ enabled: !isAnonymous, retry: false, placeholderData: prev => prev }
	)
	const collections = useMemo(() => data?.data || [], [data])

	const [search, handleSearch] = useState("")
	const [expanded, setExpanded] = useState(isExpanded)
	const [hovering, setHovering] = useState<string | undefined>()

	const visibleCollectibles = useMemo(() => {
		if (search !== "" && collections.length === 0) return Array(5).fill(undefined)

		if (collections === undefined || isAnonymous || (search === "" && collections.length === 0))
			return !isColumn ? PLACEHOLDER_COLLECTIONS.slice(0, 5) : PLACEHOLDER_COLLECTIONS

		const filteredCollectibles = collections.filter(
			collection =>
				collection.attributes.collection_info.name.toLowerCase().includes(search.toLowerCase()) ||
				collection.attributes.collection_info.description.toLowerCase().includes(search.toLowerCase())
		)

		if (expanded) return filteredCollectibles

		return filteredCollectibles.slice(0, count)
	}, [isAnonymous, collections, expanded, count, search])

	return (
		<div className={cn("flex flex-col gap-2", className)} {...props}>
			{!isColumn && (
				<Header
					variant="frame"
					icon={<Frame size={14} className="opacity-40" />}
					label={
						<div className="flex w-full items-center justify-between">
							<p className="font-bold">Collectibles</p>
							<p className="flex gap-1 text-xs font-bold opacity-40">
								<Counter count={hovering ? hovering : visibleCollectibles.length} />
								<span className="opacity-40">/</span>
								<Counter count={collections.length} />
							</p>
						</div>
					}
					nextOnClick={() => setExpanded(prev => !prev)}
					nextLabel={expanded ? "See Less" : "See All"}
				/>
			)}

			{isAnonymous === false && isColumn && collections.length > 0 && (
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
				{visibleCollectibles.map((collection, collectionIndex) => (
					<Animate.ListItem key={collectionIndex}>
						<SocketCollectionItem
							index={index}
							collection={collection}
							searched={false}
							onMouseEnter={() => setHovering((collectionIndex + 1).toString())}
							onMouseLeave={() => setHovering(undefined)}
						/>
					</Animate.ListItem>
				))}
			</Animate.List>

			<Callout.Anonymous index={index} viewing="collectibles" isAbsolute={true} />
			<Callout.EmptyAssets
				index={index}
				isEmpty={isColumn && !isAnonymous && search === "" && collections.length === 0}
				isViewing="collectibles"
				isReceivable
			/>

			<CollectionFrame index={index} address={address} />
			<CollectibleFrame index={index} />
		</div>
	)
})

SocketCollectionList.displayName = "SocketCollectionList"
