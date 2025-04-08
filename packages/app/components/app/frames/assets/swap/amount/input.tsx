import { FC, useCallback, useRef, useState } from "react"

import { TokenImage } from "@/components/app/sockets/tokens/token-image"
import { Counter } from "@/components/shared/utils/counter"
import { getChainId } from "@/lib"
import { RouterOutputs } from "@/server/client"
import { ChainImage } from "@/components/app/sockets/chains/chain.image"


type Token =
	| RouterOutputs["socket"]["balances"]["positions"]["tokens"][number]
	| RouterOutputs["solver"]["tokens"]["get"][number]

type SwapAmountInputProps = {
	index: number
	token: Token
	color: string
	amounts: {
		precise: string
		percentage: number
	}
	setAmounts: (
		amounts:
			| { precise: string; percentage: number }
			| ((prev: { precise: string; percentage: number }) => { precise: string; percentage: number })
	) => void
}

export const SwapAmountInput: FC<SwapAmountInputProps> = ({ token, color, amounts, setAmounts }) => {
	const containerRef = useRef<HTMLDivElement>(null)
	const inputRef = useRef<HTMLInputElement>(null)

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

					// @ts-ignore
					const newAmount = ((token?.balance ?? token.implementations[0]?.balance ?? 0) * percentage) / 100

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

					setAmounts({
						precise: finalAmount,
						percentage
					})
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
		[token, amounts, setAmounts]
	)

	const handleAmountChange = (value: string) => {
		const numericValue = value.replace(/[^0-9.]/g, "")
		const parsedValue = parseFloat(numericValue)

		if (numericValue === "") {
			setAmounts({ ...amounts, precise: "0" })
		} else if (!isNaN(parsedValue)) {
			// @ts-ignore
			const maxBalance = token?.balance ?? token?.implementations[0]?.balance ?? 0
			const clampedValue = Math.min(Math.max(parsedValue, 0), maxBalance)
			const percentage = (clampedValue / maxBalance) * 100

			const precise = numericValue.includes(".")
				? numericValue.split(".")[0] + "." + numericValue.split(".")[1].slice(0, 18)
				: clampedValue.toString()
			setAmounts({ ...amounts, percentage, precise })
		} else {
			setAmounts({ ...amounts, precise: numericValue })
		}
	}

	return (
		<>
			<div
				className="relative mr-6 flex cursor-ew-resize items-center gap-4 overflow-hidden rounded-r-lg border-[1px] border-l-[0px] border-plug-green/10 p-4"
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
								`https://token-icons.llamao.fi/icons/tokens/${getChainId(token.implementations[0].chain)}/${token.implementations[0].contract}?h=240&w=240`
							}
							symbol={token.symbol}
							size="sm"
						/>

						<div className="flex flex-col items-center">
							<p className="mr-auto font-bold">{token.symbol}</p>
							<div className="relative flex flex-row items-center gap-2">
								<ChainImage chainId={getChainId(token.implementations[0].chain)} size="xs" />
								<p className="flex flex-row text-sm font-bold text-black/40">
									<Counter count={amounts?.percentage ?? 0} decimals={0} />%
								</p>
							</div>
						</div>
					</div>

					<div className="ml-auto flex-col items-end px-2">
						<div className="pointer-events-none relative flex h-full w-max min-w-32 flex-col items-center justify-center text-right">
							{isPrecise && (
								<input
									ref={inputRef}
									value={amounts?.precise ?? 0}
									onChange={e => handleAmountChange(e.target.value)}
									className="sr-only pointer-events-none absolute inset-0"
									autoFocus
								/>
							)}

							<p
								className="my-auto ml-auto flex flex-row font-bold tabular-nums transition-all duration-200 ease-in-out"
								style={{ color: isPrecise ? color : undefined }}
							>
								<Counter
									count={
										Number(amounts?.precise).toLocaleString("en-US", {
											maximumFractionDigits: 40
										}) ?? "0"
									}
								/>

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
										count={(Number(amounts?.precise) * (token.price ?? 0)).toLocaleString("en-US", {
											minimumFractionDigits: 2,
											maximumFractionDigits: 2
										})}
										decimals={2}
									/>
								</p>
							)}
						</div>
					</div>
				</div>

				<div
					className="absolute inset-0 z-[-1] min-w-4 rounded-r-lg opacity-20 blur-2xl filter"
					style={{ width: `${amounts.percentage ?? 0}%`, backgroundColor: color }}
				>
					<div className="absolute inset-0 rounded-r-[16px] shadow-[inset_4px_0_4px_0_rgba(255,255,255,.5)]" />
					<div className="absolute inset-0 rounded-r-[16px] shadow-[inset_0_4px_4px_0_rgba(255,255,255,0.5)]" />
					<div className="absolute inset-0 rounded-r-[16px] shadow-[inset_0_-4px_4px_0_rgba(255,255,255,0.5)]" />
				</div>
			</div>
		</>
	)
}
