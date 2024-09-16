"use client"

import { FC } from "react"

import { RouterOutputs } from "@/server/client"

import { SocketCollectibleItem } from "@/components"

export const SocketCollectibleGrid: FC<{
	index: number
	collection: NonNullable<RouterOutputs["socket"]["balances"]["collectibles"]>[number]
}> = ({ index, collection }) => {
	return (
		<div className="grid grid-cols-2 gap-4">
			{collection.collectibles.map((collectible, collectibleIndex) => {
				return (
					<SocketCollectibleItem
						key={collectibleIndex}
						index={index}
						collection={collection}
						collectible={collectible}
					/>
				)
			})}
		</div>
	)
}
