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
			onClick={() => frame()}
		>
			<Image
				src={collectible?.attributes.nft_info?.content?.detail?.url || collectible.attributes.collection_info?.content?.icon?.url || ""}
				alt={collectible?.attributes.nft_info.name ?? ""}
				width={300}
				height={300}
				className={cn(
					"rounded-t-md w-full",
					loading
						? "animate-loading bg-gradient-animated bg-[length:200%_200%]"
						: "transition-all duration-200 ease-in-out",
					loading === false ? "cursor-pointer" : "cursor-default"
				)}
				onLoad={() => setLoading(false)}
			/>

			<div className="text-left px-4 py-2 rounded-b-md border-[1px] border-plug-green/10">
				<p className="font-bold">{collectible.attributes.nft_info.name}</p>
				<p className="font-bold opacity-40 text-xs truncate">#{collectible.attributes.nft_info.token_id}</p>
			</div>
		</div>
	)
}
