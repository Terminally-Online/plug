import Image from "next/image"
import { FC, useCallback, useEffect, useMemo, useRef, useState } from "react"

import { Counter, Frame, TokenImage } from "@/components"
import { chains, cn, formatTitle, getChainId, getTextColor } from "@/lib"
import { RouterOutputs } from "@/server/client"
import { useColumns } from "@/state"

import { TransferRecipient } from "./transfer-recipient"

type Implementation = NonNullable<
	RouterOutputs["socket"]["balances"]["positions"]
>["tokens"][number]["implementations"][number]

const ImplementationComponent: FC<{
	implementation: Implementation
	token: NonNullable<RouterOutputs["socket"]["balances"]["positions"]>["tokens"][number]
	index: number
	color: string
}> = ({ implementation, token, index, color }) => {
	const containerRef = useRef<HTMLDivElement>(null)
	const inputRef = useRef<HTMLInputElement>(null)

	const { column, transfer } = useColumns(index)

	const [isPrecise, setIsPrecise] = useState(false)

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
					const percentage = Math.min(Math.max((x / rect.width) * 100, 0), 100)

					transfer(prev => ({
						...prev,
						percentage
					}))

					const newAmount = (implementation.balance * percentage) / 100
					const formattedAmount = newAmount.toFixed(20).replace(/\.?0+$/, "")
					const [integerPart, decimalPart] = formattedAmount.split(".")
					let finalAmount = integerPart

					if (decimalPart) {
						if (decimalPart.length <= 2) {
							finalAmount += "." + decimalPart
						} else {
							const significantDecimals = decimalPart.match(/^0*[1-9]\d?|0{2,}/)?.[0] || ""
							finalAmount += "." + significantDecimals
						}
					}

					transfer(prev => ({
						...prev,
						precise: finalAmount
					}))

					setIsPrecise(true)
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
		[implementation, transfer]
	)

	const handleAmountChange = (value: string) => {
		const numericValue = value.replace(/[^0-9.]/g, "")
		const parsedValue = parseFloat(numericValue)

		if (numericValue === "") {
			transfer(prev => ({ ...prev, precise: "0" }))
		} else if (!isNaN(parsedValue)) {
			const maxBalance = implementation.balance
			const clampedValue = Math.min(Math.max(parsedValue, 0), maxBalance)
			const percentage = (clampedValue / maxBalance) * 100

			const precise = numericValue.includes(".")
				? numericValue.split(".")[0] + "." + numericValue.split(".")[1].slice(0, 18)
				: clampedValue.toString()
			transfer(prev => ({ ...prev, percentage, precise }))
		} else {
			transfer(prev => ({ ...prev, precise: numericValue }))
		}
	}

	return (
		<div
			className="relative mr-6 flex cursor-ew-resize items-center gap-4 overflow-hidden rounded-r-lg border-[1px] border-l-[0px] border-grayscale-100 p-4"
			ref={containerRef}
			onMouseDown={handleDragStart}
			onMouseEnter={() => setIsPrecise(true)}
			onMouseLeave={() => setIsPrecise(false)}
		>
			<div className="flex w-full flex-row">
				<div className="flex flex-row items-center gap-4 px-2">
					<TokenImage
						logo={
							token?.icon ||
							`https://token-icons.llamao.fi/icons/tokens/${getChainId(implementation.chain)}/${implementation.contract}?h=240&w=240`
						}
						symbol={token.symbol}
						size="sm"
					/>

					<div className="flex flex-col items-center">
						<p className="mr-auto font-bold">{formatTitle(token.symbol)}</p>
						<div className="relative flex flex-row items-center gap-2">
							<Image
								className="h-4 w-4 rounded-full"
								src={chains[getChainId(implementation.chain)].logo}
								alt=""
								width={24}
								height={24}
							/>
							<p className="flex flex-row text-sm font-bold text-black/40">
								<Counter count={column?.transfer?.percentage ?? 0} decimals={0} />%
							</p>
						</div>
					</div>
				</div>

				<div className="ml-auto flex-col items-end px-2">
					<div className="pointer-events-none relative flex h-full w-max min-w-32 flex-col items-center justify-center text-right">
						{isPrecise && (
							<input
								ref={inputRef}
								value={column?.transfer?.precise ?? 0}
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

						{!isPrecise && (
							<p className="ml-auto flex text-sm font-bold tabular-nums text-black/40">
								<span className="ml-auto">$</span>
								<Counter
									count={
										((implementation.balance * (column?.transfer?.percentage ?? 0)) / 100) *
										(token.price ?? 0)
									}
									decimals={2}
								/>
							</p>
						)}
					</div>
				</div>
			</div>

			<div
				className="absolute inset-0 z-[-1] min-w-4 rounded-r-lg opacity-20 blur-2xl filter"
				style={{ width: `${column?.transfer?.percentage ?? 0}%`, backgroundColor: color }}
			>
				<div className="absolute inset-0 rounded-r-[16px] shadow-[inset_4px_0_4px_0_rgba(255,255,255,.5)]" />
				<div className="absolute inset-0 rounded-r-[16px] shadow-[inset_0_4px_4px_0_rgba(255,255,255,0.5)]" />
				<div className="absolute inset-0 rounded-r-[16px] shadow-[inset_0_-4px_4px_0_rgba(255,255,255,0.5)]" />
			</div>
		</div>
	)
}

export const TransferAmountFrame: FC<{
	index: number
	token: NonNullable<RouterOutputs["socket"]["balances"]["positions"]>["tokens"][number]
	color: string
	textColor: string
}> = ({ index, token, color, textColor }) => {
	const { isFrame, column, frame, transfer } = useColumns(
		index,
		index === -2 ? `${token?.symbol}-transfer-deposit` : `${token?.symbol}-transfer-amount`
	)

	const isReady = useMemo(() => token && column && parseFloat(column?.transfer?.precise ?? "0") > 0, [token, column])

	if (!token || !column) return null

	return (
		<>
			<Frame
				index={index}
				icon={
					<div className="relative h-8 w-10">
						<TokenImage
							logo={
								token?.icon ||
								`https://token-icons.llamao.fi/icons/tokens/${getChainId(token.implementations[0].chain)}/${token.implementations[0].contract}?h=240&w=240`
							}
							symbol={token.symbol}
							size="sm"
						/>
					</div>
				}
				label={`${index === -2 ? "Deposit" : "Transfer"}`}
				visible={isFrame}
				handleBack={index !== -2 ? () => frame(`${token.symbol}-transfer-recipient`) : undefined}
				hasChildrenPadding={false}
				hasOverlay
			>
				<div className="mb-4 flex flex-col gap-2">
					{index !== -2 && (
						<div className="px-6">
							<TransferRecipient
								address={column?.transfer?.recipient ?? ""}
								handleSelect={() => frame(`${token.symbol}-transfer-recipient`)}
							/>
						</div>
					)}
					<div className="flex flex-col gap-2">
						<div className="flex flex-col gap-2">
							{token.implementations.map((implementation, index) => (
								<ImplementationComponent
									key={index}
									implementation={implementation}
									token={token}
									index={index}
									color={color}
								/>
							))}
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
							<p
								className="ml-auto cursor-pointer font-bold text-black/40 hover:brightness-105"
								onClick={() =>
									transfer(prev => ({
										...prev,
										percentage: 100,
										precise: token?.implementations[0].balance.toString()
									}))
								}
								style={{ color: (column?.transfer?.percentage ?? 0) < 100 ? color : undefined }}
							>
								Max
							</p>
						</div>
					</div>

					<div className="mx-6 mt-2 flex flex-col gap-4">
						<button
							className={cn(
								"flex w-full items-center justify-center gap-2 rounded-lg border-[1px] py-4 font-bold transition-all duration-200 ease-in-out hover:opacity-90 hover:brightness-105",
								isReady === false && "bg-white"
							)}
							style={{
								backgroundColor: isReady ? color : "#FFFFFF",
								color: isReady ? textColor : color,
								borderColor: isReady ? "#FFFFFF" : color
							}}
							disabled={isReady === false}
						>
							{isReady ? (index === -2 ? "Deposit" : "Send") : "Enter Amount"}
						</button>
					</div>
				</div>
			</Frame>
		</>
	)
}
