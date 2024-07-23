import { FC } from "react"

import { motion } from "framer-motion"

import { Accordion } from "@/components"
import { RouterOutputs } from "@/server/client"

type Collectibles = NonNullable<RouterOutputs["socket"]["collectibles"]>

type Props = {
	collection: Collectibles[keyof Collectibles]
	collectible?: Collectibles[keyof Collectibles]["collectibles"][number]
}

export const SocketCollectibleItem: FC<Props> = ({
	collection,
	collectible
}) => {
	const loading = collectible === undefined

	return (
		<motion.div
			variants={{
				hidden: { opacity: 0, y: 10 },
				visible: {
					opacity: 1,
					y: 0,
					transition: {
						type: "spring",
						stiffness: 100,
						damping: 10
					}
				}
			}}
		>
			<Accordion
				loading={loading}
				expanded={false}
				onExpand={() => {}}
				noPaddingChildren={
					<div
						style={{
							position: "relative",
							width: "100%",
							paddingTop: "100%", // This creates a square based on the width
							backgroundImage: `url(${collectible?.display_image_url || collection.imageUrl})`,
							backgroundSize: "contain",
							backgroundPosition: "center",
							backgroundRepeat: "no-repeat"
						}}
					/>
				}
				noPadding={true}
			/>
		</motion.div>
	)
}
