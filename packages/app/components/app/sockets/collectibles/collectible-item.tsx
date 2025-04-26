import { FC, useState } from "react"

import { Image } from "@/components/app/utils/image"
import { cn } from "@/lib"
import { RouterOutputs } from "@/server/client"
import { useColumnActions } from "@/state/columns"

export const SocketCollectibleItem: FC<{
	index: number
	collectible: NonNullable<RouterOutputs["service"]["zerion"]["wallet"]["nftPositions"]["data"]>[number]
}> = ({ index, collectible }) => {
	const { frame } = useColumnActions(index, `collectible___${collectible?.relationships.nft.data.id}`)

	const [loading, setLoading] = useState(true)

	return (
		<div
			className="relative z-[4] w-full rounded-md"
			style={{
				paddingTop: "100%"
			}}
			onClick={() => frame()}
		>
			<Image
				src={collectible?.attributes.nft_info?.content?.detail?.url || collectible.attributes.collection_info?.content?.icon?.url || ""}
				alt={collectible?.attributes.nft_info.name ?? ""}
				fill
				style={{
					objectFit: "cover",
					objectPosition: "center"
				}}
				className={cn(
					"rounded-md",
					loading
						? "animate-loading bg-gradient-animated bg-[length:200%_200%]"
						: "transition-all duration-200 ease-in-out",
					loading === false ? "cursor-pointer" : "cursor-default"
				)}
				onLoad={() => setLoading(false)}
			/>
		</div>
	)
}
