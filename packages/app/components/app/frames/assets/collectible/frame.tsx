import { FC, memo, useEffect, useState } from "react"
import ReactMarkdown from "react-markdown"

import { getAddress } from "viem"

import {
	BookDashed,
	BookText,
	ChevronDown,
	Globe,
	Hash,
	Instagram,
	MapIcon,
	MessageCircle,
	Send,
	Ship,
	Twitter,
	Waypoints
} from "lucide-react"

import { useAtom, useAtomValue } from "jotai"

import { TransferFrame } from "@/components/app/frames/assets/transfer/frame"
import { Frame } from "@/components/app/frames/base"
import { ChainImage } from "@/components/app/sockets/chains/chain.image"
import { CollectibleImage } from "@/components/app/sockets/collectibles/collectible-image"
import {
	cn,
	formatAddress,
	formatTitle,
	formatTokenStandard,
	getBlockExplorerAddress,
	getChainId,
	getTextColor
} from "@/lib"
import { api, RouterOutputs } from "@/server/client"
import { columnByIndexAtom, isFrameAtom, useColumnActions } from "@/state/columns"

type Traits = Array<{ trait_type: string; value: string }>

export const CollectibleFrame: FC<{
	index: number
	collection: NonNullable<RouterOutputs["socket"]["balances"]["collectibles"]>[number]
	collectible: NonNullable<RouterOutputs["socket"]["balances"]["collectibles"]>[number]["collectibles"][number]
}> = memo(({ index, collection, collectible }) => {
	const [column] = useAtom(columnByIndexAtom(index))
	const frameKey = `${collection.address}-${collection.chain}-${collectible?.tokenId}`
	const isFrame = useAtomValue(isFrameAtom)(column, frameKey)
	const { frame, transfer } = useColumnActions(index, frameKey)

	const { data: metadata } = api.socket.balances.metadata.useQuery(
		{
			type: "ERC721",
			address: collectible.collectionAddress,
			chain: collectible.collectionChain,
			tokenId: collectible.tokenId
		},
		{ enabled: isFrame }
	)

	const [expanded, setExpanded] = useState(false)

	const textColor = getTextColor(metadata?.color ?? "#ffffff")

	useEffect(() => {
		if (isFrame === false) setExpanded(false)
	}, [isFrame])

	return (
		<>
			<Frame
				index={index}
				icon={
					<div className="relative h-8 w-10">
						<div
							className="absolute left-1/2 top-1/2 -translate-x-1/2 -translate-y-1/2 animate-fade-in rounded-full bg-plug-green/10 blur-2xl filter"
							style={{
								backgroundImage: `url(${collection.iconUrl})`,
								backgroundSize: "cover",
								backgroundPosition: "center",
								backgroundRepeat: "no-repeat",
								width: "4rem",
								minWidth: "4rem",
								height: "4rem"
							}}
						/>
						<div
							className="absolute left-1/2 top-1/2 -translate-x-1/2 -translate-y-1/2 animate-fade-in rounded-full bg-plug-green/10"
							style={{
								backgroundImage: `url(${collection.iconUrl})`,
								backgroundSize: "cover",
								backgroundPosition: "center",
								backgroundRepeat: "no-repeat",
								width: "2rem",
								minWidth: "2rem",
								height: "2rem"
							}}
						/>
					</div>
				}
				label={collection.name}
				visible={isFrame}
				hasChildrenPadding={false}
				hasOverlay
			>
				<div className="flex flex-col gap-2 px-6">
					<CollectibleImage
						video={collectible?.videoUrl?.includes("mp4") ? collectible?.videoUrl : undefined}
						image={collectible?.imageUrl ?? undefined}
						fallbackImage={collection.iconUrl ?? ""}
						name={collectible?.name || collection.name}
					/>

					<p className="pt-2 text-left text-lg font-bold">{collectible?.name}</p>

					{collection.description && (
						<>
							<div className="inline-flex gap-2">
								<ReactMarkdown
									className="relative z-10 w-full text-left text-sm font-bold opacity-60"
									components={{
										p: ({ children }) => <p className="mb-4">{children}</p>,
										a: ({ node, children, ...props }) => (
											<a
												{...props}
												className="relative z-20 cursor-pointer transition-opacity duration-200 hover:opacity-80"
												style={{ color: metadata?.color ?? "" }}
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
									{expanded
										? `${collection.description.trim()}`
										: `${collection.description.trim().slice(0, 120)}...`}
								</ReactMarkdown>
							</div>

							{collection.description.trim().length > 120 && (
								<>
									<button
										className="mb-2 mr-auto flex flex-row items-center gap-2 text-sm font-bold"
										onClick={() => setExpanded(!expanded)}
									>
										{expanded ? "Read Less" : "Read More"}
										<ChevronDown
											size={18}
											className={cn(
												"ml-auto opacity-40 transition-all duration-200 ease-in-out",
												expanded && "rotate-180"
											)}
										/>
									</button>
								</>
							)}
						</>
					)}
				</div>

				<div className="flex flex-row gap-2 px-6 pb-4">
					<button
						onClick={() => {
							transfer({ percentage: 0, precise: "0", recipient: undefined })
							frame(`${collection.address}-${collection.chain}-${collectible.tokenId}-transfer-recipient`)
						}}
						className="flex w-full items-center justify-center gap-2 rounded-lg py-4 font-bold transition-all duration-200 ease-in-out hover:opacity-90"
						style={{
							backgroundColor: metadata?.color ?? "",
							color: textColor
						}}
					>
						<Send size={14} className="opacity-60" />
						Send
					</button>
				</div>

				{metadata?.traits && (metadata?.traits?.length ?? 0) > 0 && (
					<div className="flex flex-wrap gap-2 px-6 pb-4">
						{(metadata.traits as Traits).map((trait, index) => (
							<div
								key={index}
								className="flex flex-col rounded-lg border-2 px-4 py-2"
								style={{ borderColor: metadata?.color ?? "" }}
							>
								<p className="truncate overflow-ellipsis whitespace-nowrap text-sm font-bold opacity-40">
									{formatTitle(trait.trait_type)}
								</p>
								<p className="flex flex-row items-center gap-2 truncate overflow-ellipsis whitespace-nowrap font-bold">
									{trait.value}
								</p>
							</div>
						))}
					</div>
				)}

				<div className="flex flex-col gap-2 px-6 pb-4">
					<div>
						<div className="flex flex-row items-center gap-4">
							<p className="font-bold opacity-40">Details</p>
							<div className="h-[2px] w-full bg-plug-green/10" />
						</div>

						<div className="mt-2 w-full font-bold">
							<p className="flex w-full flex-row items-center gap-4">
								<BookText size={18} className="opacity-20" />
								<span className="mr-auto opacity-40">Address</span>
								{formatAddress(getAddress(collectible.collectionAddress))}
							</p>
							<p className="flex w-full flex-row items-center gap-4">
								<Hash size={18} className="opacity-20" />
								<span className="mr-auto opacity-40">Identifier</span>
								{collectible.tokenId.length > 11
									? formatAddress(collectible.tokenId)
									: collectible.tokenId}
							</p>
							{collectible.interface === "ERC1155" && (
								<p className="flex w-full flex-row items-center gap-4">
									<Hash size={18} className="opacity-20" />
									<span className="mr-auto opacity-40">Balance</span>
									{collectible.amount}
								</p>
							)}
							<p className="flex w-full flex-row items-center gap-4">
								<Waypoints size={18} className="opacity-20" />
								<span className="mr-auto opacity-40">Chain</span>
								<span className="flex flex-row items-center gap-2">
									<ChainImage chainId={getChainId(collection.chain)} size="xs" />
									{formatTitle(collection.chain)}
								</span>
							</p>
							<p className="flex w-full flex-row items-center gap-4">
								<BookText size={18} className="opacity-20" />
								<span className="mr-auto opacity-40">Token Standard</span>
								<span className="whitespace-nowrap">{formatTokenStandard(collectible.interface)}</span>
							</p>
						</div>

						<div className="mt-4 flex flex-row items-center gap-4">
							<p className="font-bold opacity-40">Links</p>
							<div className="h-[2px] w-full bg-plug-green/10" />
						</div>

						<div className="mt-2 flex flex-wrap gap-2">
							<a
								className="flex flex-row items-center gap-2 rounded-md px-4 py-2 text-xs font-bold transition-all duration-200 ease-in-out hover:opacity-90"
								style={{
									backgroundColor: metadata?.color ?? "",
									color: textColor
								}}
								href={getBlockExplorerAddress(
									getChainId(collectible.collectionChain),
									collectible.collectionAddress
								)}
								target="_blank"
								rel="noreferrer"
							>
								<MapIcon size={14} className="opacity-60" />
								Explorer
							</a>

							{collection.projectUrl && (
								<a
									className="flex flex-row items-center gap-2 rounded-md px-4 py-2 text-xs font-bold transition-all duration-200 ease-in-out hover:opacity-90"
									style={{
										backgroundColor: metadata?.color ?? "",
										color: textColor
									}}
									href={collection.projectUrl}
									target="_blank"
									rel="noreferrer"
								>
									<Globe size={14} className="opacity-60" />
									Website
								</a>
							)}

							{collection.discordUrl && (
								<a
									className="flex flex-row items-center gap-2 rounded-md px-4 py-2 text-xs font-bold transition-all duration-200 ease-in-out hover:opacity-90"
									style={{
										backgroundColor: metadata?.color ?? "",
										color: textColor
									}}
									href={collection.discordUrl}
									target="_blank"
									rel="noreferrer"
								>
									<MessageCircle size={14} className="opacity-60" />
									Discord
								</a>
							)}

							{collection.twitterUsername && (
								<a
									className="flex flex-row items-center gap-2 rounded-md px-4 py-2 text-xs font-bold transition-all duration-200 ease-in-out hover:opacity-90"
									style={{
										backgroundColor: metadata?.color ?? "",
										color: textColor
									}}
									href={`https://twitter.com/${collection.twitterUsername}`}
									target="_blank"
									rel="noreferrer"
								>
									<Twitter size={14} className="opacity-60" />
									Twitter
								</a>
							)}

							{collection.telegramUrl && (
								<a
									className="flex flex-row items-center gap-2 rounded-md px-4 py-2 text-xs font-bold transition-all duration-200 ease-in-out hover:opacity-90"
									style={{
										backgroundColor: metadata?.color ?? "",
										color: textColor
									}}
									href={collection.telegramUrl}
									target="_blank"
									rel="noreferrer"
								>
									<Globe size={14} className="opacity-60" />
									Telegram
								</a>
							)}

							{collection.instagramUsername && (
								<a
									className="flex flex-row items-center gap-2 rounded-md px-4 py-2 text-xs font-bold transition-all duration-200 ease-in-out hover:opacity-90"
									style={{
										backgroundColor: metadata?.color ?? "",
										color: textColor
									}}
									href={`https://instagram.com/${collection.instagramUsername}`}
									target="_blank"
									rel="noreferrer"
								>
									<Instagram size={14} className="opacity-60" />
									Instagram
								</a>
							)}

							{collection.openseaUrl && (
								<a
									className="flex flex-row items-center gap-2 rounded-md px-4 py-2 text-xs font-bold transition-all duration-200 ease-in-out hover:opacity-90"
									style={{
										backgroundColor: metadata?.color ?? "",
										color: textColor
									}}
									href={collection.openseaUrl}
									target="_blank"
									rel="noreferrer"
								>
									<Ship size={14} className="opacity-60" />
									Opensea
								</a>
							)}

							{collection.wikiUrl && (
								<a
									className="flex flex-row items-center gap-2 rounded-md px-4 py-2 text-xs font-bold transition-all duration-200 ease-in-out hover:opacity-90"
									style={{
										backgroundColor: metadata?.color ?? "",
										color: textColor
									}}
									href={collection.wikiUrl}
									target="_blank"
									rel="noreferrer"
								>
									<BookDashed size={14} className="opacity-60" />
									Wiki
								</a>
							)}
						</div>
					</div>
				</div>
			</Frame>

			<TransferFrame
				index={index}
				collectible={collectible}
				collection={collection}
				color={metadata?.color ?? ""}
				textColor={textColor}
			/>
		</>
	)
})

CollectibleFrame.displayName = "CollectibleFrame"
