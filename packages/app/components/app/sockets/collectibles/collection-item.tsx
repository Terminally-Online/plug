import { FC, memo } from "react"

import { Image } from "@/components/app/utils/image"
import { Accordion } from "@/components/shared/utils/accordion"
import { formatTitle, ZerionCollection } from "@/lib"

import { useColumnActions } from "@/state/columns"

export const SocketCollectionItem: FC<{
	index: number
	collection: ZerionCollection | undefined
	searched?: boolean
}> = memo(({ index, collection }) => {
	const { frame } = useColumnActions(index, `collection___${collection?.relationships.nft_collection.data.id ?? ""}___${collection?.attributes.collection_info.name}___${collection?.attributes.collection_info.content.icon.url}`)

	return (
		<Accordion
			loading={collection === undefined}
			className="text-left"
			onExpand={() => frame()}
		>
			{collection === undefined ? (
				<div className="invisible">
					<p>.</p>
					<p>.</p>
				</div>
			) : (
				<div className="flex w-full flex-row items-center gap-4">
					<div className="relative h-10 min-w-10">
						<Image
							src={collection.attributes.collection_info.content.icon.url ?? ""}
							alt={collection.attributes.collection_info.name}
							className="absolute left-1/2 h-48 w-48 -translate-x-1/2 rounded-full blur-2xl"
							width={240}
							height={240}
						/>
						<Image
							className="absolute left-1/2 top-1/2 h-10 w-10 min-w-10 -translate-x-1/2 -translate-y-1/2 rounded-full"
							src={collection.attributes.collection_info.content.icon.url ?? ""}
							alt={collection.attributes.collection_info.name}
							width={240}
							height={240}
						/>
					</div>

					<div className="flex w-min flex-col truncate overflow-ellipsis">
						<p className="truncate whitespace-nowrap font-bold">
							{formatTitle(collection.attributes.collection_info.name.toLowerCase())}
						</p>
						<div className="relative flex w-max flex-row items-center gap-2">
							<p className="text-sm font-bold opacity-40">
								{collection.attributes.nfts_count}{" "}
								Token
								{parseInt(collection.attributes.nfts_count) > 1 && "s"}
							</p>
						</div>
					</div>
				</div>
			)}
		</Accordion>
	)
})

SocketCollectionItem.displayName = "SocketCollectionItem"
