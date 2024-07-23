import { FC } from "react"

import { SocketCollectibleItem } from "@/components"
import { RouterOutputs } from "@/server/client"

type Collectibles = NonNullable<RouterOutputs["socket"]["collectibles"]>

type Props = {
	collection: Collectibles[keyof Collectibles]
}

export const SocketCollectibleGrid: FC<Props> = ({ collection }) => {
	if (!collection) return <></>

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
