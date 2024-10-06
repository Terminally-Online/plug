import Image from "next/image"
import { FC, useState } from "react"

import { RouterOutputs } from "@/server/client"

import { cn } from "@/lib"
import { useColumns } from "@/state"

export const SocketCollectibleItem: FC<{
	index: number
	collection: NonNullable<RouterOutputs["socket"]["balances"]["collectibles"]>[number]
	collectible?: NonNullable<RouterOutputs["socket"]["balances"]["collectibles"]>[number]["collectibles"][number]
}> = ({ index, collection, collectible }) => {
	const { frame } = useColumns(index)

	const [loading, setLoading] = useState(true)

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
