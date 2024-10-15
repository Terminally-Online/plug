import Image from "next/image"
import { FC, useCallback, useEffect, useRef, useState } from "react"

import { User } from "lucide-react"

import { RouterOutputs } from "@/server/client"

import { Button, Counter, Frame, Search, TokenImage } from "@/components"
import { chains, cn, formatForDisplay, formatTitle, getChainId, getTextColor } from "@/lib"
import { useColumns } from "@/state"

const DEFAULT_TRANSFER = {
	token: undefined,
	chain: undefined,
	amount: undefined,
	to: undefined
}

export const TransferFrame: FC<{
	index: number
	token: NonNullable<RouterOutputs["socket"]["balances"]["positions"]>["tokens"][number]
}> = ({ index, token }) => {
	const { column } = useColumns(index, `${token?.symbol}-transfer`)
	const [dragPercentage, setDragPercentage] = useState(100)
	const containerRef = useRef<HTMLDivElement>(null)
	const [totalDollarValue, setTotalDollarValue] = useState<number>(0)
	const [color, setColor] = useState("")
	const [isPreciseAmount, setIsPreciseAmount] = useState<boolean>(false)
	const [preciseAmount, setPreciseAmount] = useState<string>("")

	const textColor = getTextColor(color ?? "#ffffff")

	const handleAmountClick = (index: number) => {
		setIsPreciseAmount(prev => {
			if (!prev) {
				const currentAmount = (token.implementations[index].balance * dragPercentage) / 100
				setPreciseAmount(formatForDisplay(currentAmount, true).toString())
			}
			return !prev
		})
	}

	const handleAmountChange = (value: string, index: number) => {
		const parsedValue = parseFloat(value)
		if (!isNaN(parsedValue)) {
			const maxBalance = token.implementations[index].balance
			const clampedValue = Math.min(Math.max(parsedValue, 0), maxBalance)
			const newPercentage = (clampedValue / maxBalance) * 100

			setPreciseAmount(clampedValue.toString())
			setDragPercentage(newPercentage)

			const newTotalDollarValue = token.implementations.reduce((total, implementation, i) => {
				const amount = i === index ? clampedValue : (implementation.balance * newPercentage) / 100
				return total + amount * (token.price ?? 0)
			}, 0)
			setTotalDollarValue(newTotalDollarValue)
		} else {
			setPreciseAmount(value)
		}
	}

	const handleDragStart = useCallback((e: React.MouseEvent<HTMLDivElement>) => {
		e.preventDefault()

		console.log("handleDragStart")

		const activeElement = document.activeElement as HTMLElement
		if (activeElement && activeElement.tagName === "INPUT") {
			activeElement.blur()
		}

		const startX = e.clientX
		let hasDragged = false

		const handleDrag = (e: MouseEvent) => {
			const currentX = e.clientX
			const diff = Math.abs(currentX - startX)

			if (diff > 5) {
				// 5px threshold
				hasDragged = true
				setIsPreciseAmount(false)
			}

			if (hasDragged && containerRef.current) {
				const rect = containerRef.current.getBoundingClientRect()
				const x = e.clientX - rect.left
				const percentage = Math.min(Math.max((x / rect.width) * 100, 0), 100)
				setDragPercentage(percentage)
			}
		}

		const handleDragEnd = () => {
			document.removeEventListener("mousemove", handleDrag)
			document.removeEventListener("mouseup", handleDragEnd)
		}

		document.addEventListener("mousemove", handleDrag)
		document.addEventListener("mouseup", handleDragEnd)
	}, [])

	const handleMaxClick = useCallback(() => {
		setDragPercentage(100)
		setTotalDollarValue(
			token.implementations.reduce((total, implementation) => {
				return total + implementation.balance * (token.price ?? 0)
			}, 0)
		)
		setIsPreciseAmount(true)
	}, [token])

	useEffect(() => {
		if (!token) return

		const newTotalDollarValue = token.implementations.reduce((total, implementation) => {
			const amount = (implementation.balance * dragPercentage) / 100
			return total + amount * (token.price ?? 0)
		}, 0)

		setTotalDollarValue(newTotalDollarValue)
	}, [dragPercentage, token])

	useEffect(() => {
		if (isPreciseAmount) {
			const inputElement = document.querySelector('input[type="number"]') as HTMLInputElement
			if (inputElement) {
				inputElement.focus()
			}
		}
	}, [isPreciseAmount])

	if (!token || !column) return null

	return (
		<>
			<Frame
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
				label="Transfer"
				visible={column.frame === `${token.symbol}-transfer-send`}
				hasChildrenPadding={false}
			>
				<div className="flex flex-col gap-4 pb-4">
					<div className="flex flex-col gap-2">
						<div className="flex flex-col gap-2" ref={containerRef}>
							{token.implementations.map((implementation, index) => (
								<div
									key={index}
									className="relative mr-6 flex cursor-ew-resize items-center gap-4 overflow-hidden rounded-r-lg border-[1px] border-l-[0px] border-grayscale-100 p-4"
									onMouseDown={handleDragStart}
								>
									<div className="flex w-full flex-row">
										<div className="flex flex-row items-center gap-4 px-2">
											<TokenImage
												logo={
													token?.icon ||
													`https://token-icons.llamao.fi/icons/tokens/${getChainId(token.implementations[0].chain)}/${token.implementations[0].contract}?h=240&w=240`
												}
												symbol={token.symbol}
												size="sm"
												handleColor={setColor}
											/>

											<div className="flex flex-col items-center">
												<p className="mr-auto font-bold">{token.symbol}</p>
												<div className="flex flex-row items-center gap-2">
													<Image
														className="h-4 w-4 rounded-full"
														src={chains[getChainId(implementation.chain)].logo}
														alt=""
														width={24}
														height={24}
													/>
													<p className="text-sm font-bold text-black/40">
														{formatTitle(implementation.chain)}
													</p>
												</div>
											</div>
										</div>

										<div
											className="ml-auto flex cursor-pointer flex-col items-end px-2"
											onClick={() => handleAmountClick(index)}
										>
											<div
												className={cn(
													"pointer-events-none flex h-full min-w-32 flex-col items-center justify-center text-right"
												)}
											>
												{isPreciseAmount && (
													<input
														type="number"
														value={preciseAmount}
														onChange={e => handleAmountChange(e.target.value, index)}
														onFocus={() =>
															setPreciseAmount(
																((implementation.balance * dragPercentage) / 100)
																	.toFixed(3)
																	.toString()
															)
														}
														className="sr-only pointer-events-none"
														autoFocus
													/>
												)}

												<p
													className={cn(
														"my-auto ml-auto flex flex-row font-bold tabular-nums"
														// isPreciseAmount && "text-xl"
													)}
													style={{ color: isPreciseAmount ? color : undefined }}
												>
													<Counter
														count={
															isPreciseAmount
																? parseFloat(preciseAmount)
																: (implementation.balance * dragPercentage) / 100
														}
														decimals={isPreciseAmount ? 3 : undefined}
													/>

													{isPreciseAmount && (
														<div
															className="ml-2 h-full w-[4px] animate-pulse rounded-full"
															style={{ backgroundColor: color }}
														/>
													)}
												</p>

												{!isPreciseAmount && (
													<>
														<p className="ml-auto flex text-sm font-bold tabular-nums text-black/40">
															<span className="ml-auto">$</span>
															<Counter count={totalDollarValue} decimals={2} />
														</p>
													</>
												)}
											</div>
										</div>
									</div>

									<div
										className="absolute inset-0 z-[-1] min-w-4 rounded-r-lg bg-grayscale-100"
										style={{ width: `${dragPercentage}%` }}
									>
										<div className="absolute inset-0 rounded-r-[16px] shadow-[inset_4px_0_4px_0_rgba(255,255,255,.5)]" />
										<div className="absolute inset-0 rounded-r-[16px] shadow-[inset_0_4px_4px_0_rgba(255,255,255,0.5)]" />
										<div className="absolute inset-0 rounded-r-[16px] shadow-[inset_0_-4px_4px_0_rgba(255,255,255,0.5)]" />
									</div>
								</div>
							))}
						</div>

						<div className="flex flex-row items-center justify-between gap-4 px-6">
							<p className="flex flex-row font-bold tabular-nums text-black/40">11 seconds</p>
							<p
								className="ml-auto cursor-pointer font-bold text-black/40 hover:text-black/60"
								onClick={handleMaxClick}
								style={{ color: dragPercentage < 100 ? color : undefined }}
							>
								Max
							</p>
						</div>
					</div>

					<div className="mx-6 flex flex-col gap-4">
						<button
							className="flex w-full items-center justify-center gap-2 rounded-lg py-4 font-bold transition-all duration-200 ease-in-out hover:opacity-90"
							style={{
								backgroundColor: color ?? "",
								color: textColor
							}}
							// onClick={() => frame(`${token.symbol}-transfer-send`)}
						>
							Confirm
						</button>
					</div>
				</div>
			</Frame>
		</>
	)
}
