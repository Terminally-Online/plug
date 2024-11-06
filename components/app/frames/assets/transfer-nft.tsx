import { FC, useCallback, useMemo, useRef, useState } from "react"

import { Accordion, CollectibleImage, Counter, Frame, Image } from "@/components"
import { chains, cn, formatTitle } from "@/lib"
import { RouterOutputs } from "@/server/client"
import { useColumns } from "@/state"

import { TransferRecipient } from "./transfer-recipient"

type CollectibleType = NonNullable<RouterOutputs["socket"]["balances"]["collectibles"]>[number]["collectibles"][number]

type CollectionType = NonNullable<RouterOutputs["socket"]["balances"]["collectibles"]>[number]

type TransferNFTFrameProps = {
	index: number
	collectible: CollectibleType
	collection: CollectionType
	color: string
	textColor: string
	isERC1155: boolean
}

export const TransferNFTFrame: FC<TransferNFTFrameProps> = ({
	index,
	collectible,
	collection,
	color,
	textColor,
	isERC1155
}) => {
	const containerRef = useRef<HTMLDivElement>(null)
	const inputRef = useRef<HTMLInputElement>(null)

	const { isFrame, frame } = useColumns(
		index,
		`${collection.address}-${collection.chain}-${collectible.tokenId}-transfer-amount`
	)

	const { column, transfer } = useColumns(index)
	const [isPrecise, setIsPrecise] = useState(false)

	const maxAmount = parseInt(collectible.amount)

	const handleDragStart = useCallback(
		(e: React.MouseEvent<HTMLDivElement>) => {
			e.preventDefault()

			const activeElement = document.activeElement as HTMLElement
			if (activeElement && activeElement.tagName === "INPUT") {
				activeElement.blur()
			}

			const handleDrag = (e: MouseEvent) => {
				if (containerRef.current) {
					const rect = containerRef.current.getBoundingClientRect()
					const x = e.clientX - rect.left
					const rawPercentage = x / rect.width

					// First calculate what whole number we should be at
					const totalOptions = maxAmount + 1 // +1 because 0 is an option
					const stepSize = 1 / maxAmount // Each whole number takes this much percentage

					// Calculate which step we're closest to
					const nearestWholeNumber = Math.round(rawPercentage * maxAmount)

					// Convert back to percentage
					const snappedPercentage = (nearestWholeNumber / maxAmount) * 100

					transfer(prev => ({
						...prev,
						percentage: snappedPercentage,
						precise: nearestWholeNumber.toString()
					}))
				}
			}

			const handleDragEnd = () => {
				document.removeEventListener("mousemove", handleDrag)
				document.removeEventListener("mouseup", handleDragEnd)
				if (inputRef.current) {
					inputRef.current.focus()
				}
			}

			document.addEventListener("mousemove", handleDrag)
			document.addEventListener("mouseup", handleDragEnd)
		},
		[maxAmount]
	)

	const handleAmountChange = (value: string) => {
		const numericValue = value.replace(/[^0-9]/g, "")

		if (numericValue === "") {
			transfer(prev => ({ ...prev, precise: "0" }))
		} else {
			transfer(prev => ({ ...prev, precise: numericValue }))
		}
	}
	const handleMaxClick = useCallback(() => {
		if (isERC1155) {
			transfer(prev => ({ ...prev, percentage: 100, precise: collectible.amount }))
		}
	}, [isERC1155, maxAmount])

	const isReady = useMemo(() => {
		if (!isERC1155) return true
		const numAmount = parseInt(column?.transfer?.precise ?? "0")
		return !isNaN(numAmount) && numAmount > 0 && numAmount <= maxAmount
	}, [column?.transfer?.precise, isERC1155, maxAmount])

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
			label="Transfer"
			visible={isFrame}
			handleBack={() =>
				frame(`${collection.address}-${collection.chain}-${collectible.tokenId}-transfer-recipient`)
			}
			hasChildrenPadding={false}
			hasOverlay
		>
			<div className="relative mb-4 flex flex-col gap-2">
				<div className="flex flex-col gap-2">
					<div className="px-6">
						<TransferRecipient
							address={column?.transfer?.recipient ?? ""}
							handleSelect={() =>
								frame(`${collection.address}-${collection.chain}-${collectible.tokenId}-recipient`)
							}
						/>
					</div>

					{!isERC1155 ? (
						<div className="px-6">
							<Accordion>
								<div className="flex w-full flex-row items-center gap-4">
									<div className="relative h-10 w-10 min-w-10">
										<CollectibleImage
											className="rounded-md"
											video={
												collectible.videoUrl?.includes("mp4") ? collectible.videoUrl : undefined
											}
											image={collectible.imageUrl ?? undefined}
											fallbackImage={collection.iconUrl ?? undefined}
											name={collectible.name || collection.name}
										/>
									</div>
									<div className="flex w-full flex-col truncate overflow-ellipsis">
										<div className="flex flex-row items-center justify-between">
											<p className="mr-auto font-bold">
												{formatTitle(collectible.name || collection.name)}
											</p>
										</div>
										<div className="flex flex-row items-center justify-between">
											<p className="mr-auto truncate overflow-ellipsis text-sm font-bold tabular-nums opacity-40">
												#{collectible.tokenId}
											</p>
										</div>
									</div>
								</div>
							</Accordion>
						</div>
					) : (
						<div className="relative z-[5] flex flex-col gap-4">
							<div
								className="relative mr-6 flex cursor-ew-resize items-center gap-4 overflow-hidden rounded-r-lg border-[1px] border-l-[0px] border-grayscale-100 p-4"
								ref={containerRef}
								onMouseDown={handleDragStart}
								onMouseEnter={() => setIsPrecise(true)}
								onMouseLeave={() => setIsPrecise(false)}
							>
								<div className="flex w-full flex-row">
									<div className="flex flex-row items-center gap-4 px-2">
										<div className="h-8 w-8 min-w-8 overflow-hidden">
											<CollectibleImage
												className="rounded-sm"
												video={
													collectible.videoUrl?.includes("mp4")
														? collectible.videoUrl
														: undefined
												}
												image={collectible.imageUrl ?? undefined}
												fallbackImage={collection.iconUrl ?? undefined}
												name={collectible.name || collection.name}
												size="sm"
											/>
										</div>

										<div className="flex flex-col">
											<p className="mr-auto truncate overflow-ellipsis font-bold">{`${collectible.name}`}</p>
											<p className="flex w-max flex-row text-sm font-bold text-black/40">
												<Counter count={column?.transfer?.percentage ?? 0} decimals={0} />%
											</p>
										</div>
									</div>

									<div className="ml-auto flex-col items-end px-2">
										<div className="pointer-events-none relative flex h-full w-max min-w-32 flex-col items-center justify-center text-right">
											{isPrecise && (
												<input
													ref={inputRef}
													value={column?.transfer?.precise}
													onChange={e => handleAmountChange(e.target.value)}
													className="sr-only pointer-events-none absolute inset-0"
													autoFocus
												/>
											)}

											<p
												className="my-auto ml-auto flex flex-row font-bold tabular-nums transition-all duration-200 ease-in-out"
												style={{ color: isPrecise ? color : undefined }}
											>
												<Counter count={column?.transfer?.precise ?? "0"} />

												{isPrecise && (
													<div
														className="absolute -right-2 bottom-3 top-3 w-[3px] animate-pulse rounded-full"
														style={{ backgroundColor: color }}
													/>
												)}
											</p>
										</div>
									</div>
								</div>

								<div
									className="absolute inset-0 z-[-1] min-w-4 rounded-r-lg opacity-20 blur-2xl filter"
									style={{ width: `${column?.transfer?.percentage}%`, backgroundColor: color }}
								>
									<div className="absolute inset-0 rounded-r-[16px] shadow-[inset_4px_0_4px_0_rgba(255,255,255,.5)]" />
									<div className="absolute inset-0 rounded-r-[16px] shadow-[inset_0_4px_4px_0_rgba(255,255,255,0.5)]" />
									<div className="absolute inset-0 rounded-r-[16px] shadow-[inset_0_-4px_4px_0_rgba(255,255,255,0.5)]" />
								</div>
							</div>
						</div>
					)}
				</div>

				<div className="flex flex-row items-center justify-between gap-4 px-6">
					<p className="flex flex-row items-center gap-1 font-bold tabular-nums">
						<Image
							src={chains[1].logo}
							alt={"ethereum"}
							className="mr-2 h-4 w-4 rounded-full"
							width={24}
							height={24}
						/>
						$0.50
						<span className="ml-auto pl-2 opacity-40">~11 secs</span>
					</p>

					{isERC1155 && (
						<p
							className="ml-auto cursor-pointer font-bold text-black/40 hover:brightness-105"
							onClick={handleMaxClick}
							style={{ color: column?.transfer?.precise !== collectible.amount ? color : undefined }}
						>
							Max
						</p>
					)}
				</div>

				<div className="relative mx-6 mt-2">
					<button
						className={cn(
							"flex w-full items-center justify-center gap-2 rounded-lg border-[1px] py-4 font-bold transition-all duration-200 ease-in-out hover:opacity-90 hover:brightness-105",
							!isReady && "bg-white"
						)}
						style={{
							backgroundColor: isReady ? color : "#FFFFFF",
							color: isReady ? textColor : color,
							borderColor: isReady ? "#FFFFFF" : color
						}}
						disabled={!isReady}
					>
						{isReady ? "Send" : "Enter Amount"}
					</button>
				</div>
			</div>
		</Frame>
	)
}
