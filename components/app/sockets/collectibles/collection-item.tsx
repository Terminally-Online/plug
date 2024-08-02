import { FC, useState } from "react"

import Image from "next/image"

import { motion } from "framer-motion"

import { Accordion } from "@/components"
import { getChainImage } from "@/lib"
import { RouterOutputs } from "@/server/client"

import { SocketCollectibleGrid } from "./collectible-grid"

type Collectibles = NonNullable<
	RouterOutputs["socket"]["balances"]["collectibles"]
>

export const SocketCollectionItem: FC<{
	collection: Collectibles[keyof Collectibles] | undefined
}> = ({ collection }) => {
	const [expanded, setExpanded] = useState(false)
	const [error, setError] = useState(false)

	if (error) return <></>

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
				loading={collection === undefined}
				className="text-left"
				expanded={expanded}
				onExpand={
					collection === undefined
						? () => {}
						: () => setExpanded(!expanded)
				}
				accordion={
					collection && (
						<SocketCollectibleGrid collection={collection} />
					)
				}
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
								src={collection.imageUrl}
								alt={collection.name}
								className="absolute left-1/2 top-1/2 h-48 w-48 -translate-x-1/2 -translate-y-1/2 animate-fade-in rounded-full blur-2xl filter transition-all duration-200 ease-in-out"
								width={140}
								height={140}
								onError={() => setError(true)}
							/>
							<div
								className="absolute left-1/2 top-1/2 h-10 w-10 min-w-10 -translate-x-1/2 -translate-y-1/2 animate-fade-in rounded-full bg-grayscale-100"
								style={{
									backgroundImage: `url(${collection.imageUrl})`,
									backgroundSize: "cover",
									backgroundPosition: "center",
									backgroundRepeat: "no-repeat"
								}}
							/>
						</div>

						<div className="flex w-min flex-col truncate overflow-ellipsis">
							<p className="truncate font-bold">
								{collection.name}
							</p>
							<div className="flex w-max flex-row items-center gap-2">
								<Image
									src={getChainImage(collection.chain)}
									alt={collection.name}
									className="z-1 relative h-4 w-4 rounded-full bg-grayscale-100"
									width={48}
									height={48}
								/>
								<p className="text-sm opacity-60">
									{collection.collectibles.length} Token
									{collection.collectibles.length > 1 && "s"}
								</p>
							</div>
						</div>
					</div>
				)}
			</Accordion>
		</motion.div>
	)
}
