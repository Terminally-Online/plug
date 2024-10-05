import Image from "next/image"
import { FC } from "react"

import { RouterOutputs } from "@/server/client"

import { useColumns } from "@/state"

export const SocketCollectibleItem: FC<{
	index: number
	collection: NonNullable<RouterOutputs["socket"]["balances"]["collectibles"]>[number]
	collectible?: NonNullable<RouterOutputs["socket"]["balances"]["collectibles"]>[number]["collectibles"][number]
}> = ({ index, collection, collectible }) => {
	const { frame } = useColumns(index)

	return (
		<div
			className="relative z-[4] w-full rounded-md"
			style={{
				paddingTop: "100%"
			}}
			onClick={() => frame(`${collection.address}-${collection.chain}-${collectible?.tokenId}`)}
		>
			<Image
				src={collectible?.imageUrl || collection.iconUrl || ""}
				alt={collectible?.name ?? ""}
				fill
				style={{
					objectFit: "cover",
					objectPosition: "center"
				}}
				className="rounded-md"
			/>
		</div>
	)
}
