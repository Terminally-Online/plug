import { FC, useState } from "react"

import { Image } from "@/components/app/utils/image"
import { cn } from "@/lib"
import { RouterOutputs } from "@/server/client"
import { useColumnActions } from "@/state/columns"

export const SocketCollectibleItem: FC<{
	index: number
	collection: NonNullable<RouterOutputs["socket"]["balances"]["collectibles"]>[number]
	collectible?: NonNullable<RouterOutputs["socket"]["balances"]["collectibles"]>[number]["collectibles"][number]
}> = ({ index, collection, collectible }) => {
	const { frame } = useColumnActions(index, `${collection.address}-${collection.chain}-${collectible?.tokenId}`)

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
