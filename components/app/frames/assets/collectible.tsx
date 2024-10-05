import { FC, useEffect, useState } from "react"

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

import { api, RouterOutputs } from "@/server/client"

import { CollectibleImage, Frame, Image } from "@/components"
import {
	cn,
	formatAddress,
	formatLongString,
	formatTitle,
	formatTokenStandard,
	getBlockExplorerAddress,
	getChainId,
	getChainImage,
	getTextColor
} from "@/lib"
import { useColumns } from "@/state"

type Traits = Array<{ trait_type: string; value: string }>

export const CollectibleFrame: FC<{
	index: number
	collection: NonNullable<RouterOutputs["socket"]["balances"]["collectibles"]>[number]
	collectible: NonNullable<RouterOutputs["socket"]["balances"]["collectibles"]>[number]["collectibles"][number]
}> = ({ index, collection, collectible }) => {
	const { isFrame } = useColumns(index, `${collection.address}-${collection.chain}-${collectible?.tokenId}`)

	const { data: metadata } = api.socket.balances.metadata.useQuery(
		{
			type: "ERC721",
			address: collectible.collectionAddress,
			chain: collectible.collectionChain,
			tokenId: collectible.tokenId
		},
		{
			staleTime: Infinity
		}
	)

	const [expanded, setExpanded] = useState(false)

	const textColor = getTextColor(metadata?.color ?? "#ffffff")

	const { truncated } = formatLongString(
		// collectible?.description || collection.description,
		collection.description,
		expanded === false ? 80 : undefined
	)

	useEffect(() => {
		if (isFrame === false) setExpanded(false)
	}, [isFrame])

	return (
		<Frame
			index={index}
			icon={
				<div className="relative h-8 w-10">
					<div
						className="absolute left-1/2 top-1/2 -translate-x-1/2 -translate-y-1/2 animate-fade-in rounded-full bg-grayscale-100 blur-2xl filter"
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
						className="absolute left-1/2 top-1/2 -translate-x-1/2 -translate-y-1/2 animate-fade-in rounded-full bg-grayscale-100"
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
			<div className="flex flex-col gap-2 px-6 pb-4">
				<CollectibleImage
					video={collectible?.videoUrl?.includes("mp4") ? collectible?.videoUrl : undefined}
					image={collectible?.imageUrl ?? undefined}
					fallbackImage={collection.iconUrl ?? ""}
					name={collectible?.name || collection.name}
				/>

				<p className="pt-4 text-lg font-bold">{collectible?.name}</p>

				<div
					className={cn(
						"relative flex flex-col gap-2 overflow-y-hidden transition-all duration-200 ease-in-out",
						expanded === false ? "max-h-[60px]" : "h-auto max-h-[1000px]"
					)}
				>
					{/* <p className="font-bold opacity-40">{collectible?.description || collection.description}</p> */}
					<p className="font-bold opacity-40">{collection.description}</p>

					<div
						className={cn(
							"absolute bottom-0 left-0 right-0 top-0 h-full bg-gradient-to-b from-white/0 transition-all duration-200 ease-in-out",
							truncated && expanded === false ? "to-white" : "to-white/0"
						)}
					/>
				</div>

				{(truncated || expanded) && (
					<button
						className="mr-auto flex flex-row items-center gap-2 font-bold"
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
				)}
			</div>

			<div className="flex flex-row gap-2 px-6 pb-4">
				<button
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
						<div className="h-[2px] w-full" style={{ backgroundColor: metadata?.color ?? "" }} />
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
							{collectible.tokenId.length > 11 ? formatAddress(collectible.tokenId) : collectible.tokenId}
						</p>
						<p className="flex w-full flex-row items-center gap-4">
							<Waypoints size={18} className="opacity-20" />
							<span className="mr-auto opacity-40">Chain</span>
							<span className="flex flex-row items-center gap-4">
								<Image
									className="h-4 w-4"
									src={getChainImage(collection.chain)}
									alt={collection.chain}
									width={24}
									height={24}
								/>
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
						<div
							className="h-[2px] w-full"
							style={{
								backgroundColor: metadata?.color ?? ""
							}}
						/>
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
	)
}
