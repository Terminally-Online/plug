"use client"

import { FC } from "react"

import { SocketCollectibleItem } from "@/components"
import { RouterOutputs } from "@/server/client"

export const SocketCollectibleGrid: FC<{
	collection: NonNullable<
		RouterOutputs["socket"]["balances"]["collectibles"]
	>[number]
}> = ({ collection }) => {
	return (
		<div className="grid grid-cols-2 gap-4">
			{collection.collectibles.map((collectible, index) => {
				return (
					<SocketCollectibleItem
						key={index}
						collection={collection}
						collectible={collectible}
					/>
				)
			})}
		</div>
	)
}
