import React, { FC, useState } from "react"

import { RouterOutputs } from "@/server/client"

import { Accordion, Image } from "@/components"
import { chains, formatTitle, getChainId } from "@/lib"

import { SocketCollectibleGrid } from "./collectible-grid"

export const SocketCollectionItem: FC<{
	index: number
	collection: NonNullable<RouterOutputs["socket"]["balances"]["collectibles"]>[number] | undefined
	searched?: boolean
}> = ({ index, collection, searched = false }) => {
	const [expanded, setExpanded] = useState(searched)
	const [error, setError] = useState(false)

	if (error) return null

	return (
		<Accordion
			loading={collection === undefined}
			className="text-left"
			expanded={expanded || searched}
			onExpand={collection === undefined || searched ? () => {} : () => setExpanded(!expanded)}
			accordion={collection && <SocketCollectibleGrid index={index} collection={collection} />}
		>
			{collection === undefined ? (
				<div className="invisible">
					<p>.</p>
					<p>.</p>
				</div>
			) : (
				<div className="flex w-full flex-row items-center gap-4">
					<div className="relative h-10 min-w-10">
						<Image
							src={collection.iconUrl ?? ""}
							alt={collection.name}
							className="absolute left-1/2 h-48 w-48 -translate-x-1/2 animate-fade-in rounded-full blur-2xl filter transition-all duration-200 ease-in-out"
							width={140}
							height={140}
							onError={() => setError(true)}
						/>
						<div
							className="absolute left-1/2 top-1/2 h-10 w-10 min-w-10 -translate-x-1/2 -translate-y-1/2 animate-fade-in rounded-full bg-grayscale-100"
							style={{
								backgroundImage: `url(${collection.iconUrl})`,
								backgroundSize: "cover",
								backgroundPosition: "center",
								backgroundRepeat: "no-repeat"
							}}
						/>
					</div>

					<div className="flex w-min flex-col truncate overflow-ellipsis">
						<p className="truncate font-bold">{formatTitle(collection.name.toLowerCase())}</p>
						<div className="flex w-max flex-row items-center gap-2">
							<Image
								src={chains[getChainId(collection.chain)].logo}
								alt={collection.name}
								className="z-1 relative h-4 w-4 rounded-full bg-grayscale-100"
								width={48}
								height={48}
							/>
							<p className="text-sm font-bold opacity-40">
								{collection.collectibles.length} Token
								{collection.collectibles.length > 1 && "s"}
							</p>
						</div>
					</div>
				</div>
			)}
		</Accordion>
	)
}
