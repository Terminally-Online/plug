import { FC } from "react"

import { Accordion, TestFrame } from "@/components"
import { useFrame } from "@/contexts"
import { RouterOutputs } from "@/server/client"

import { CollectibleFrame } from "../../frames/assets/collectible"

export const SocketCollectibleItem: FC<{
	collection: NonNullable<
		RouterOutputs["socket"]["balances"]["collectibles"]
	>[number]
	collectible?: NonNullable<
		RouterOutputs["socket"]["balances"]["collectibles"]
	>[number]["collectibles"][number]
}> = ({ collection, collectible }) => {
	const { handleFrameVisible } = useFrame()

	return (
		<>
			<div
				className="w-full rounded-md"
				style={{
					paddingTop: "100%",
					backgroundImage: `url(${collectible?.displayImageUrl || collection.imageUrl})`,
					backgroundSize: "cover",
					backgroundPosition: "center",
					backgroundRepeat: "no-repeat"
				}}
				onClick={() =>
					handleFrameVisible(
						`${collection.slug}-${collectible?.contract}-${collectible?.identifier}`
					)
				}
			/>
		</>
	)
}
