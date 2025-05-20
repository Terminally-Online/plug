import { FC, useState } from "react"
import ReactMarkdown from "react-markdown"

import { CircleDollarSign, Tag, Waypoints } from "lucide-react"

import { useAtom } from "jotai"

import { ChainImage } from "@/components/app/sockets/chains/chain.image"
import { SocketCollectibleGrid } from "@/components/app/sockets/collectibles/collectible-grid"
import { Image } from "@/components/app/utils/image"
import { Counter } from "@/components/shared/utils/counter"
import { cn, formatTitle } from "@/lib"
import { api } from "@/server/client"
import { useSocket } from "@/state/authentication"
import { columnByIndexAtom } from "@/state/columns"

import { Frame } from "../../base"
import { TokenFrameExternalLink } from "../token/link"

export const CollectionFrame: FC<{ index: number; address?: string }> = ({ index, address }) => {
	const { socket, isAnonymous } = useSocket()

	const [column] = useAtom(columnByIndexAtom(index))
	const frameKey = "collection"
	const [columnFrameKey, collectionId] = column?.frame?.split("___") || []
	const isFrame = frameKey === columnFrameKey

	const { data: { data: collectibles = [], included } = {} } = api.service.zerion.wallet.nftPositions.useQuery(
		{
			path: { address: address ?? socket.socketAddress },
			query: { filter: { collectionIds: [collectionId] }, include: ["nft_collections", "nfts"] }
		},
		{ enabled: isFrame && !!collectionId && !isAnonymous }
	)

	const collection = included?.find(resource => resource.type === "nft_collections")
	const collectionName = collection?.attributes?.metadata?.name ?? ""

	const nft = included?.find(resource => resource.type === "nfts")

	return (
		<Frame
			index={index}
			icon={
				<div className="relative h-10 min-w-10">
					<Image
						src={collection?.attributes?.metadata?.icon?.url ?? ""}
						alt={collectionName}
						className="absolute left-1/2 h-48 w-48 -translate-x-1/2 rounded-full blur-2xl"
						width={240}
						height={240}
					/>
					<Image
						className="absolute left-1/2 top-1/2 h-10 w-10 min-w-10 -translate-x-1/2 -translate-y-1/2 rounded-full"
						src={collection?.attributes?.metadata?.icon?.url ?? ""}
						alt={collectionName}
						width={240}
						height={240}
					/>
				</div>
			}
			label={collectionName}
			visible={isFrame}
			hasOverlay
		>
			<SocketCollectibleGrid index={index} collectibles={collectibles} />

			<div className="mb-2 flex flex-row items-center gap-4">
				<p className="font-bold opacity-40">Details</p>
				<div className="h-[1px] w-full bg-plug-green/10" />
			</div>
			<p className="flex flex-row justify-between font-bold">
				<span className="flex flex-row items-center gap-4">
					<Waypoints size={18} className="opacity-20" />
					<span className="opacity-40">Chain</span>
				</span>{" "}
				<span className="flex flex-row items-center gap-4 font-bold">
					<ChainImage chainId={collectibles[0]?.relationships.chain.data.id ?? ""} size="xs" />
					{formatTitle(collectibles[0]?.relationships.chain.data.id ?? "")}
				</span>
			</p>

			<p className="flex flex-row justify-between font-bold">
				<span className="flex w-full flex-row items-center gap-4">
					<Tag size={18} className="opacity-20" />
					<span className="opacity-40">Floor Price</span>
				</span>{" "}
				<span className="flex flex-row items-center gap-1 font-bold tabular-nums">
					<Counter count={collection?.attributes?.market_data?.prices?.floor ?? "-"} />
					{collection?.attributes?.metadata?.payment_token_symbol}
				</span>
			</p>

			{nft?.attributes?.external_links && nft?.attributes?.external_links.length > 0 && (
				<>
					<div className="flex flex-row items-center gap-4 pt-4 font-bold">
						<p className="whitespace-nowrap opacity-40">External Links</p>
						<div className="h-[1px] w-full bg-plug-green/10" />
					</div>

					{nft.attributes.external_links?.map((link, linkIndex) => (
						<TokenFrameExternalLink key={linkIndex} link={link} />
					))}
				</>
			)}

			{collection?.attributes?.metadata?.description && (
				<>
					<div className="mb-2 flex flex-row items-center gap-4 pt-4">
						<p className="font-bold opacity-40">Description</p>
						<div className="h-[1px] w-full bg-plug-green/10" />
					</div>

					<div className="relative">
						<ReactMarkdown
							className="relative w-full text-justify font-bold"
							components={{
								p: ({ children }) => <p>{children}</p>,
								a: ({ node, children, ...props }) => (
									<a
										{...props}
										className="relative cursor-pointer transition-opacity duration-200 hover:opacity-80"
										// style={{ color: metadata?.color ?? "" }}
										target="_blank"
										rel="noopener noreferrer"
										onClick={e => {
											e.preventDefault()
											e.stopPropagation()
											if (props.href) {
												window.open(props.href, "_blank", "noopener,noreferrer")
											}
										}}
									>
										{children}
									</a>
								)
							}}
						>
							{collection?.attributes?.metadata?.description.trim()}
						</ReactMarkdown>
					</div>
				</>
			)}
		</Frame>
	)
}
