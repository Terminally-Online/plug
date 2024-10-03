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
			className="z-[4] w-full rounded-md"
			style={{
				paddingTop: "100%",
				backgroundImage: `url(${collectible?.displayImageUrl || collection.imageUrl})`,
				backgroundSize: "cover",
				backgroundPosition: "center",
				backgroundRepeat: "no-repeat"
			}}
			onClick={() => frame(`${collection.slug}-${collectible?.contract}-${collectible?.identifier}`)}
		/>
	)
}
