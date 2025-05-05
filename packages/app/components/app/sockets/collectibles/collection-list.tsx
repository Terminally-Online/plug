import { FC, HTMLAttributes, memo, useMemo, useState } from "react"

import { SearchIcon } from "lucide-react"

import { Search } from "@/components/app/inputs/search"
import { SocketCollectionItem } from "@/components/app/sockets/collectibles/collection-item"
import { Callout } from "@/components/app/utils/callout"
import { cn } from "@/lib"
import { PLACEHOLDER_COLLECTIONS } from "@/lib/constants/placeholder/collectibles"
import { api } from "@/server/client"
import { useSocket } from "@/state/authentication"
import { CollectibleFrame } from "../../frames/assets/collectible/frame"
import { CollectionFrame } from "../../frames/assets/collection/frame"

export const SocketCollectionList: FC<
	HTMLAttributes<HTMLDivElement> & {
		index: number
		address?: string
		expanded?: boolean
		count?: number
		isColumn?: boolean
	}
> = memo(({ index, address, expanded, count = 5, isColumn = true, className, ...props }) => {
	const { isAnonymous, socket } = useSocket()

	const { data } = api.service.zerion.wallet.nftCollections.useQuery(
		{
			path: { address: address || socket?.socketAddress }
		},
		{ enabled: !isAnonymous, retry: false, placeholderData: prev => prev }
	)
	const collections = useMemo(() => data?.data || [], [data])

	const [search, handleSearch] = useState("")

	const visibleCollectibles = useMemo(() => {
		if (search !== "" && collections.length === 0) return Array(5).fill(undefined)

		if (collections === undefined || isAnonymous || (search === "" && collections.length === 0))
			return PLACEHOLDER_COLLECTIONS

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
