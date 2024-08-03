import { FC } from "react"

import { Accordion } from "@/components"
import { RouterOutputs } from "@/server/client"

export const SocketCollectibleItem: FC<{
	collection: NonNullable<
		RouterOutputs["socket"]["balances"]["collectibles"]
	>[number]
	collectible?: NonNullable<
		RouterOutputs["socket"]["balances"]["collectibles"]
	>[number]["collectibles"][number]
}> = ({ collection, collectible }) => {
	const loading = collectible === undefined

	return (
		<Accordion
			loading={loading}
			expanded={false}
			onExpand={() => {}}
			noPaddingChildren={
				<div
					style={{
						position: "relative",
						width: "100%",
						paddingTop: "100%",
						backgroundImage: `url(${collectible?.displayImageUrl || collection.imageUrl})`,
						backgroundSize: "cover",
						backgroundPosition: "center",
						backgroundRepeat: "no-repeat"
					}}
				/>
			}
			noPadding={true}
		/>
	)
}
