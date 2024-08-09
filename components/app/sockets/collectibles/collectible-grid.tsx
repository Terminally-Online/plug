"use client"

import { FC } from "react"

import { SocketCollectibleItem } from "@/components"
import { RouterOutputs } from "@/server/client"

export const SocketCollectibleGrid: FC<{
	id: string
	collection: NonNullable<
		RouterOutputs["socket"]["balances"]["collectibles"]
	>[number]
}> = ({ id, collection }) => {
	return (
		<div className="grid grid-cols-2 gap-4">
			{collection.collectibles.map((collectible, index) => {
				return (
					<SocketCollectibleItem
						key={index}
						id={id}
						collection={collection}
						collectible={collectible}
					/>
				)
			})}
		</div>
	)
}
