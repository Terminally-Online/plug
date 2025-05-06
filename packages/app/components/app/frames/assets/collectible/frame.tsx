import { FC, memo, useState } from "react"
import ReactMarkdown from "react-markdown"


import {
	ArrowDownFromLine,
	BookText,
	Hash,
	Send,
	Waypoints
} from "lucide-react"

import { useAtom, useAtomValue } from "jotai"

import { TransferFrame } from "@/components/app/frames/assets/transfer/frame"
import { Frame } from "@/components/app/frames/base"
import { ChainImage } from "@/components/app/sockets/chains/chain.image"
import { CollectibleImage } from "@/components/app/sockets/collectibles/collectible-image"
import { formatAddress, formatTitle, formatTokenStandard, getTextColor } from "@/lib"
import { api } from "@/server/client"
import { columnByIndexAtom, COLUMNS, isFrameAtom, useColumnActions } from "@/state/columns"
import { TokenFrameExternalLink } from "../token/link"
import { TokenImage } from "@/components/app/sockets/tokens/token-image"

type CollectibleFrameProps = { index: number }
export const CollectibleFrame: FC<CollectibleFrameProps> = memo(({ index }) => {
	const [column] = useAtom(columnByIndexAtom(index))
	const frameKey = "collectible"
	const [, collectibleId] = column?.frame?.split("___") || []
	const isFrame = useAtomValue(isFrameAtom)(column, `${frameKey}___${collectibleId}`)

	const { frame, transfer } = useColumnActions(index, frameKey)

	const { data: { data: collectible, included } = {} } = api.service.zerion.nfts.detail.useQuery(
		{ path: { nftId: collectibleId }, query: { include: "nft_collections" } },
		{ enabled: isFrame }
	)

	const [color, setColor] = useState("#FFFFFF")
	const textColor = getTextColor(color)

	return (
		<>
			<Frame
				index={index}
				icon={
					<div className="relative h-8 w-10 rounded-full">
						<TokenImage
							logo={included?.at(0)?.attributes?.metadata?.icon?.url}
							symbol={included?.at(0)?.attributes?.metadata?.name ?? ""}
						/>
					</div>
				}
				label={collectible?.attributes?.metadata?.name ?? ""}
				visible={isFrame}
				handleBack={() => frame(`collection___${collectible?.relationships.nft_collection.data.id}`)}
				hasOverlay
			>
				<div className="flex flex-col gap-2 pt-4">
					<TokenImage
						className="absolute"
						logo={collectible?.attributes?.metadata?.content?.detail?.url ?? ""}
						symbol={collectible?.attributes?.metadata?.name ?? ""}
						handleColor={setColor}
						blur={false}
					/>

					<CollectibleImage
						video={collectible?.attributes?.metadata?.content?.video?.url}
						image={collectible?.attributes?.metadata?.content?.detail?.url}
						name={collectible?.attributes?.metadata?.name ?? ""}
					/>

					<div className="flex flex-row gap-2 py-2">
						<button
							className="flex w-full items-center justify-center gap-2 rounded-lg py-4 font-bold transition-all duration-200 ease-in-out hover:opacity-90"
							style={{
								backgroundColor: color ?? "",
								color: textColor
							}}
							onClick={() => {
								transfer({ percentage: 0, precise: "0", recipient: undefined })
								frame(
									index === COLUMNS.SIDEBAR_INDEX
										? `collectible___${collectible?.id}___transfer-amount`
										: `collectible___${collectible?.id}___transfer-recipient`
								)
							}}
						>
							{index === COLUMNS.SIDEBAR_INDEX ? (
								<>
									<ArrowDownFromLine size={14} className="opacity-60" />
									Deposit
								</>
							) : (
								<>
									<Send size={14} className="opacity-60" />
									Send
								</>
							)}
						</button>
					</div>

					{collectible?.attributes?.metadata?.attributes &&
						(collectible?.attributes?.metadata?.attributes?.length ?? 0) > 0 && (
							<div className="flex flex-wrap gap-2 pb-4">
								{collectible?.attributes?.metadata?.attributes?.map((trait, index) => (
									<div
										key={index}
										className="flex flex-col rounded-lg border-2 px-4 py-2"
										style={{ borderColor: color ?? "" }}

									>
										<p
											className="truncate overflow-ellipsis whitespace-nowrap text-sm font-bold opacity-40 text-left"
											style={{ color: color ?? "" }}
										>
											{formatTitle(trait.key)}
										</p>
										<p className="flex flex-row items-center gap-2 truncate overflow-ellipsis whitespace-nowrap font-bold">
											{trait.value}
										</p>
									</div>
								))}
							</div>
						)}

					<div className="flex flex-col gap-2">
						<div>
							<div className="flex flex-row items-center gap-4">
								<p className="font-bold opacity-40">Details</p>
								<div className="h-[1px] w-full bg-plug-green/10" />
							</div>

							<div className="mt-2 w-full font-bold">
								<p className="flex w-full flex-row items-center gap-4">
									<BookText size={18} className="opacity-20" />
									<span className="mr-auto opacity-40">Address</span>
									{formatAddress(collectible?.attributes?.contract_address ?? "")}
								</p>
								<p className="flex w-full flex-row items-center gap-4">
									<Hash size={18} className="opacity-20" />
									<span className="mr-auto opacity-40">Identifier</span>
									{(collectible?.attributes?.token_id?.length ?? 0) > 11
										? formatAddress(collectible?.attributes?.token_id ?? "")
										: collectible?.attributes?.token_id}
								</p>
								<p className="flex w-full flex-row items-center gap-4">
									<Waypoints size={18} className="opacity-20" />
									<span className="mr-auto opacity-40">Chain</span>
									<span className="flex flex-row items-center gap-2">
										<ChainImage
											chainId={collectible?.relationships.chain.data.id ?? ""}
											size="xs"
										/>
										{formatTitle(collectible?.relationships.chain.data.id ?? "")}
									</span>
								</p>
								<p className="flex w-full flex-row items-center gap-4">
									<BookText size={18} className="opacity-20" />
									<span className="mr-auto opacity-40">Token Standard</span>
									<span className="whitespace-nowrap">
										{formatTokenStandard(collectible?.attributes?.interface ?? "")}
									</span>
								</p>
							</div>


							{collectible?.attributes?.external_links && collectible?.attributes?.external_links?.length > 0 && <div className="flex flex-col px-6">
								<div className="flex flex-row items-center gap-4 font-bold">
									<p className="whitespace-nowrap opacity-40">External Links</p>
									<div className="h-[1px] w-full bg-plug-green/10" />
								</div>

								{collectible?.attributes?.external_links?.map((link, linkIndex) => (
									<TokenFrameExternalLink key={linkIndex} link={link} />
								))}
							</div>}
						</div>
					</div>
				</div>

				{collectible?.attributes?.metadata?.description && (
					<>
						<div className="mb-2 flex flex-row items-center gap-4 pt-4">
							<p className="font-bold opacity-40">Description</p>
							<div className="h-[1px] w-full bg-plug-green/10" />
						</div>

						<ReactMarkdown
							className="relative w-full text-justify font-bold"
							components={{
								p: ({ children }) => <p>{children}</p>,
								a: ({ node, children, ...props }) => (
									<a
										{...props}
										className="relative cursor-pointer transition-opacity duration-200 hover:opacity-80"
										style={{ color: color ?? "" }}
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
							{collectible?.attributes?.metadata?.description.trim()}
						</ReactMarkdown>
					</>
				)}
			</Frame>

			<TransferFrame
				index={index}
				collectible={collectible}
				included={included?.at(0)}
				color={color}
				textColor={textColor}
			/>
		</>
	)
})

CollectibleFrame.displayName = "CollectibleFrame"
