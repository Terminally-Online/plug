import { FC, useCallback, useRef, useState } from "react"

import { useAtom } from "jotai"

import { RouterOutputs } from "@/server/client"
import { columnByIndexAtom, useColumnActions } from "@/state/columns"
import { CollectibleImage } from "@/components/app/sockets/collectibles/collectible-image"
import { Counter } from "@/components/shared/utils/counter"

type TransferSFTAmountProps = {
	index: number
	collectible?: NonNullable<RouterOutputs["service"]["zerion"]["nfts"]["detail"]["data"]>
	included?: NonNullable<RouterOutputs["service"]["zerion"]["nfts"]["detail"]["included"]>[number]
	color: string
}
export const TransferSFTAmount: FC<TransferSFTAmountProps> = ({ index, collectible, included, color }) => {
	const containerRef = useRef<HTMLDivElement>(null)
	const inputRef = useRef<HTMLInputElement>(null)

	const [column] = useAtom(columnByIndexAtom(index))
	const frameKey = `collectible___${collectible?.id}___transfer-amount`
	const { transfer } = useColumnActions(index, frameKey)
	const [isPrecise, setIsPrecise] = useState(false)

	const balance = 1

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
					const nearestWholeNumber = Math.round(rawPercentage * balance)
					const snappedPercentage = (nearestWholeNumber / balance) * 100

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
		[transfer, balance, containerRef]
	)

	const handleAmountChange = (value: string) => {
		const numericValue = value.replace(/[^0-9]/g, "")

		if (numericValue === "") {
			transfer(prev => ({
				...prev,
				precise: "0",
				percentage: 0
			}))
		} else {
			const parsedValue = parseInt(numericValue)
			// const maxAmount = parseInt(collectible.amount)
			const maxAmount = 1
			const clampedValue = Math.min(Math.max(0, parsedValue), maxAmount)
			const percentage = (clampedValue / maxAmount) * 100

			transfer(prev => ({
				...prev,
				precise: clampedValue.toString(),
				percentage
			}))
		}
	}

	return <div className="relative z-[5] flex flex-col gap-4">
		<div
			className="relative mr-6 flex cursor-ew-resize items-center gap-4 overflow-hidden rounded-r-lg border-[1px] border-l-[0px] border-plug-green/10 p-4"
			ref={containerRef}
			onMouseDown={handleDragStart}
			onMouseEnter={() => setIsPrecise(true)}
			onMouseLeave={() => setIsPrecise(false)}
		>
			<div className="flex w-full flex-row">
				<div className="flex flex-row items-center gap-4 px-2">
					<div className="h-8 w-8 min-w-8 overflow-hidden">
						<CollectibleImage
							video={collectible?.attributes?.metadata?.content?.video?.url}
							image={collectible?.attributes?.metadata?.content?.detail?.url}
							name={collectible?.attributes?.metadata?.name ?? ""}
						/>
					</div>

					<div className="flex flex-col">
						<p className="mr-auto truncate overflow-ellipsis font-bold">{`${collectible?.attributes?.metadata?.name}`}</p>
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
								value={
									column?.transfer?.precise === "0"
										? ""
										: column?.transfer?.precise
								}
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
}
