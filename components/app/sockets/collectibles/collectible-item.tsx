import { FC } from "react"

import { Accordion } from "@/components"
import { RouterOutputs } from "@/server/client"

type Collectibles = NonNullable<
	RouterOutputs["socket"]["balances"]["collectibles"]
>

export const SocketCollectibleItem: FC<{
	collection: Collectibles[keyof Collectibles]
	collectible?: Collectibles[keyof Collectibles]["collectibles"][number]
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
						backgroundImage: `url(${collectible?.display_image_url || collection.imageUrl})`,
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
