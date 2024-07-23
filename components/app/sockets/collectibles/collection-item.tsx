import { FC, useState } from "react"

import Image from "next/image"

import { Accordion } from "@/components"
import { getChainImage } from "@/lib"
import { RouterOutputs } from "@/server/client"

import { SocketCollectibleGrid } from "./collectible-grid"

type Collectibles = NonNullable<RouterOutputs["socket"]["collectibles"]>

type Props = {
	collection: Collectibles[keyof Collectibles]
}

export const SocketCollectionItem: FC<Props> = ({ collection }) => {
	const [expanded, setExpanded] = useState(false)

	return (
		<Accordion
			className="text-left"
			expanded={expanded}
			onExpand={() => setExpanded(!expanded)}
			accordion={<SocketCollectibleGrid collection={collection} />}
		>
			<div className="flex flex-row items-center gap-4">
				<div className="relative h-10 w-10">
					<Image
						src={collection.imageUrl}
						alt={collection.name}
						className="absolute left-1/2 top-1/2 h-48 w-48 -translate-x-1/2 -translate-y-1/2 animate-fade-in rounded-full blur-2xl filter transition-all duration-200 ease-in-out"
						width={140}
						height={140}
					/>
					<Image
						src={collection.imageUrl}
						alt={collection.name}
						className="absolute left-1/2 top-1/2 h-10 w-10 min-w-10 -translate-x-1/2 -translate-y-1/2 animate-fade-in rounded-full bg-grayscale-100"
						width={140}
						height={140}
					/>
				</div>

				<div className="flex flex-col">
					<p className="w-72 overflow-hidden overflow-ellipsis whitespace-nowrap font-bold">
						{collection.name}
					</p>
					<div className="flex flex-row items-center gap-2">
						<Image
							src={getChainImage(collection.chain)}
							alt={collection.name}
							className="z-1 relative h-4 w-4 rounded-full bg-grayscale-100"
							width={16}
							height={16}
						/>
						<p className="overflow-hidden overflow-ellipsis whitespace-nowrap text-sm opacity-60">
							{collection.collectibles.length} Token
							{collection.collectibles.length > 1 && "s"}
						</p>
					</div>
				</div>
			</div>
		</Accordion>
	)
}
