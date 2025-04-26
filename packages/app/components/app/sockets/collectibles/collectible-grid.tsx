import { FC } from "react"

import { SocketCollectibleItem } from "@/components/app/sockets/collectibles/collectible-item"
import { RouterOutputs } from "@/server/client"

export const SocketCollectibleGrid: FC<{
	index: number
	collectibles: NonNullable<RouterOutputs["service"]["zerion"]["wallet"]["nftPositions"]["data"]>
}> = ({ index, collectibles }) => {
	return (
		<div className="grid grid-cols-2 gap-4 mx-6 pb-4">
			{collectibles.map((collectible, collectibleIndex) => {
				return (
					<SocketCollectibleItem
						key={collectibleIndex}
						index={index}
						collectible={collectible}
					/>
				)
			})}
		</div>
	)
}
