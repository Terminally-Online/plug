import Image from "next/image"
import { FC, useCallback, useEffect, useRef, useState } from "react"

import { RouterOutputs } from "@/server/client"

import { Counter, Frame, TokenImage } from "@/components"
import { chains, cn, getChainId, getTextColor } from "@/lib"
import { useColumns } from "@/state"

type Implementation = NonNullable<
	RouterOutputs["socket"]["balances"]["positions"]
>["tokens"][number]["implementations"][number]

const ImplementationComponent: FC<{
	implementation: Implementation
	token: NonNullable<RouterOutputs["socket"]["balances"]["positions"]>["tokens"][number]
	index: number
	color: string
	dragPercentage: number
	preciseAmount: string
	onDragPercentageChange: (index: number, percentage: number) => void
	onPreciseAmountChange: (index: number, amount: string) => void
}> = ({
	implementation,
	token,
	index,
	color,
	dragPercentage,
	preciseAmount,
	onDragPercentageChange,
	onPreciseAmountChange
}) => {
	const [isPrecise, setIsPrecise] = useState(false)
	const containerRef = useRef<HTMLDivElement>(null)
	const inputRef = useRef<HTMLInputElement>(null)

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
					onDragPercentageChange(index, percentage)

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

					onPreciseAmountChange(index, finalAmount)
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
		[implementation, index, onDragPercentageChange, onPreciseAmountChange]
	)

	const handleAmountChange = (value: string) => {
		const numericValue = value.replace(/[^0-9.]/g, '');
		const parsedValue = parseFloat(numericValue);

		if (numericValue === "") {
			onPreciseAmountChange(index, "0")
		} else if (!isNaN(parsedValue)) {
			const maxBalance = implementation.balance
			const clampedValue = Math.min(Math.max(parsedValue, 0), maxBalance)
			const newPercentage = (clampedValue / maxBalance) * 100

			const formattedValue = numericValue.includes('.') ? numericValue.split('.')[0] + '.' + numericValue.split('.')[1].slice(0, 18) : clampedValue.toString()
			onPreciseAmountChange(index, formattedValue)
			onDragPercentageChange(index, newPercentage)
		} else {
			onPreciseAmountChange(index, numericValue)
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
						<p className="mr-auto font-bold">{token.symbol}</p>
						<div className="relative flex flex-row items-center gap-2">
							<Image
								className="h-4 w-4 rounded-full"
								src={chains[getChainId(implementation.chain)].logo}
								alt=""
								width={24}
								height={24}
							/>
							<p className="flex flex-row text-sm font-bold text-black/40">
								<Counter count={dragPercentage ?? 0} decimals={0} />%
							</p>
						</div>
					</div>
				</div>

				<div className="ml-auto flex-col items-end px-2">
					<div className="pointer-events-none relative flex h-full w-max min-w-32 flex-col items-center justify-center text-right">
						{isPrecise && (
							<input
								ref={inputRef}
								value={preciseAmount}
								onChange={e => handleAmountChange(e.target.value)}
								className="sr-only pointer-events-none absolute inset-0"
								autoFocus
							/>
						)}

						<p
							className="my-auto ml-auto flex flex-row font-bold tabular-nums transition-all duration-200 ease-in-out"
							style={{ color: isPrecise ? color : undefined }}
						>
							<Counter count={preciseAmount ?? "0"} />

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
									count={((implementation.balance * dragPercentage) / 100) * (token.price ?? 0)}
									decimals={2}
								/>
							</p>
						)}
					</div>
				</div>
			</div>

			<div
				className="absolute inset-0 z-[-1] min-w-4 rounded-r-lg opacity-20 blur-2xl filter"
				style={{ width: `${dragPercentage}%`, backgroundColor: color }}
			>
				<div className="absolute inset-0 rounded-r-[16px] shadow-[inset_4px_0_4px_0_rgba(255,255,255,.5)]" />
				<div className="absolute inset-0 rounded-r-[16px] shadow-[inset_0_4px_4px_0_rgba(255,255,255,0.5)]" />
				<div className="absolute inset-0 rounded-r-[16px] shadow-[inset_0_-4px_4px_0_rgba(255,255,255,0.5)]" />
			</div>
		</div>
	)
}

export const TransferFrame: FC<{
	index: number
	token: NonNullable<RouterOutputs["socket"]["balances"]["positions"]>["tokens"][number]
}> = ({ index, token }) => {
	const { column, frame } = useColumns(index, `${token?.symbol}-transfer-send`)

	const [dragPercentages, setDragPercentages] = useState<number[]>([])
	const [preciseAmounts, setPreciseAmounts] = useState<string[]>([])
	const [color, setColor] = useState("")

	const textColor = getTextColor(color ?? "#ffffff")

	useEffect(() => {
		if (token) {
			setDragPercentages(new Array(token.implementations.length).fill(0))
			setPreciseAmounts(new Array(token.implementations.length).fill("0"))
		}
	}, [token])

	const handleDragPercentageChange = useCallback((index: number, percentage: number) => {
		setDragPercentages(prev => {
			const newState = [...prev]
			newState[index] = percentage
			return newState
		})
	}, [])

	const handlePreciseAmountChange = useCallback((index: number, amount: string) => {
		setPreciseAmounts(prev => {
			const newState = [...prev]
			newState[index] = amount
			return newState
		})
	}, [])

	const handleMaxClick = useCallback(() => {
		if (dragPercentages.some(p => p < 100)) {
			const newDragPercentages = new Array(token.implementations.length).fill(100)
			const newPreciseAmounts = token.implementations.map(impl => impl.balance.toString())
			setDragPercentages(newDragPercentages)
			setPreciseAmounts(newPreciseAmounts)
		}
	}, [token, dragPercentages])

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
							handleColor={setColor}
						/>
					</div>
				}
				label="Transfer Amount"
				visible={column.frame === `${token.symbol}-transfer-send`}
				handleBack={() => frame(`${token.symbol}-token`)}
				hasChildrenPadding={false}
				hasOverlay
			>
				<div className="mb-4 flex flex-col gap-4">
					<div className="flex flex-col gap-2">
						<div className="flex flex-col gap-2">
							{token.implementations.map((implementation, index) => (
								<ImplementationComponent
									key={index}
									implementation={implementation}
									token={token}
									index={index}
									color={color}
									dragPercentage={dragPercentages[index]}
									preciseAmount={preciseAmounts[index]}
									onDragPercentageChange={handleDragPercentageChange}
									onPreciseAmountChange={handlePreciseAmountChange}
								/>
							))}
						</div>

						<div className="flex flex-row items-center justify-between gap-4 px-6">
							<p className="flex flex-row font-bold tabular-nums text-black/40">11 seconds</p>
							<p
								className="ml-auto cursor-pointer font-bold text-black/40 hover:brightness-105"
								onClick={handleMaxClick}
								style={{ color: dragPercentages.some(p => p < 100) ? color : undefined }}
							>
								Max
							</p>
						</div>
					</div>

					<div className="mx-6 flex flex-col gap-4">
						<button
							className="flex w-full items-center justify-center gap-2 rounded-lg py-4 font-bold transition-all duration-200 ease-in-out hover:opacity-90 hover:brightness-105"
							style={{
								backgroundColor: color ?? "",
								color: textColor
							}}
						>
							Confirm
						</button>
					</div>
				</div>
			</Frame>
		</>
	)
}
