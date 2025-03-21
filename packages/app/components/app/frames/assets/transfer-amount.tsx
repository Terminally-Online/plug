import { FC, useCallback, useRef, useState } from "react"

import { getAddress, isAddress } from "viem"
import { useSendTransaction } from "wagmi"

import { useAtom, useAtomValue } from "jotai"

import { Frame } from "@/components/app/frames/base"
import { TokenImage } from "@/components/app/sockets/tokens/token-image"
import { Counter } from "@/components/shared/utils/counter"
import { cn, formatTitle, getChainId, NATIVE_TOKEN_ADDRESS, useConnect, useDebounceInline } from "@/lib"
import { api, RouterOutputs } from "@/server/client"
import { useSocket } from "@/state/authentication"
import { columnByIndexAtom, COLUMNS, isFrameAtom, useColumnActions } from "@/state/columns"

import { ChainImage } from "../../sockets/chains/chain.image"
import { Marquee } from "../../utils/marquee"
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

	const [column] = useAtom(columnByIndexAtom(index))
	const { transfer } = useColumnActions(index)

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
					const percentage = Math.floor(Math.min(Math.max((x / rect.width) * 100, 0), 100))

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
						precise: percentage >= 99.5 ? implementation.balance.toString() : finalAmount
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
							`https://token-icons.llamao.fi/icons/tokens/${getChainId(implementation.chain)}/${implementation.contract}?h=240&w=240`
						}
						symbol={token.symbol}
						size="sm"
					/>

					<div className="flex flex-col items-center">
						<p className="mr-auto font-bold">{formatTitle(token.symbol)}</p>
						<div className="relative flex flex-row items-center gap-2">
							<ChainImage chainId={getChainId(implementation.chain)} size="xs" />
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
							<Counter
								count={
									Number(column?.transfer?.precise ?? 0).toLocaleString("en-US", {
										maximumFractionDigits: 18
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

const ScrollingError = ({ error }: { error: string | undefined }) => {
	if (!error) return null

	return (
		<div className="relative min-h-6 overflow-x-hidden">
			<div className="absolute bottom-0 left-0 top-0 z-[20] w-12 bg-gradient-to-r from-plug-white to-plug-white/0" />
			<div className="absolute bottom-0 right-0 top-0 z-[20] w-12 bg-gradient-to-l from-plug-white to-plug-white/0" />
			<Marquee className="-z-1 relative max-w-full whitespace-nowrap font-bold text-plug-red">{error}</Marquee>
		</div>
	)
}

export const TransferAmountFrame: FC<{
	index: number
	token: NonNullable<RouterOutputs["socket"]["balances"]["positions"]>["tokens"][number]
	color: string
	textColor: string
}> = ({ index, token, color, textColor }) => {
	const [column] = useAtom(columnByIndexAtom(index))
	const frameKey = `${token?.symbol}-transfer-${index === COLUMNS.SIDEBAR_INDEX ? "deposit" : "amount"}`
	const isFrame = useAtomValue(isFrameAtom)(column, frameKey)
	const { frame, navigate } = useColumnActions(index, frameKey)

	const {
		account: { isAuthenticated }
	} = useConnect()
	const { socket } = useSocket()
	// const {
	// 	handle: {
	// 		plug: { queue }
	// 	}
	// } = usePlugStore(item)
	const { error, sendTransaction, isPending } = useSendTransaction()

	const isReady = token && column && parseFloat(column?.transfer?.precise ?? "0") > 0 && !isPending
	const from = socket
		? index === COLUMNS.SIDEBAR_INDEX
			? getAddress(socket.id)
			: getAddress(socket.socketAddress)
		: ""
	const recipient =
		column && socket
			? index === COLUMNS.SIDEBAR_INDEX
				? getAddress(socket.socketAddress)
				: column.transfer?.recipient && isAddress(column.transfer?.recipient)
					? getAddress(column.transfer?.recipient)
					: ""
			: ""

	// TODO: This is just hard-coded to base for now. It should support having multiple
	//       chains in the same execution at once.
	const chain = "base"
	const chainId = getChainId(chain)
	const implementation = token?.implementations.find(implementation => implementation.chain === chain)
	const request = useDebounceInline({
		chainId,
		from,
		inputs: [
			{
				protocol: "plug",
				action: "transfer",
				amount: `${column?.transfer?.precise ?? 0}`,
				token: `${implementation?.contract ?? NATIVE_TOKEN_ADDRESS}:${implementation?.decimals ?? 18}:20`,
				recipient
			}
		],
		options: {
			isEOA: column && column.index === COLUMNS.SIDEBAR_INDEX,
			simulate: true
		}
	})
	const { data: intent } = api.solver.actions.intent.useQuery(request, {
		enabled: isFrame && isReady && !!column && !!socket && !!implementation
	})

	const handleTransaction = () => {
		if (!column || !intent) return

		if (column.index === COLUMNS.SIDEBAR_INDEX)
			sendTransaction(
				{
					to: intent.transactions[0].to,
					data: intent.transactions[0].data,
					value: intent.transactions[0].value
				},
				{
					onSuccess: data => {
						navigate({ index, key: COLUMNS.KEYS.ACTIVITY })
						frame(`${data}-activity`)
					}
				}
			)
		else
			// TODO: Implement the socket side logic for transfers.
			return
	}

	const handleRun = () => {
		if (!column || !column.item || !chainId) return

		const intent = {
			plugId: column.item,
			chainId,
			startAt: column.schedule?.date?.from ?? new Date(),
			endAt: column.schedule?.date?.to,
			frequency: parseInt(column.schedule?.repeats?.value ?? "0")
		}

		// TODO: Transfers do not have a plug id to associate things to -- Not sure what to do about this.

		// queue(
		// 	intent,
		// 	{
		// 		onError: data => console.error(data),
		// 		onSuccess: data => {
		// 			navigate({ index, key: COLUMNS.KEYS.ACTIVITY })
		// 			frame(`${data.id}-activity`)
		// 		}
		// 	}
		// )
	}

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
				label={`${index === COLUMNS.SIDEBAR_INDEX ? "Deposit" : "Transfer"}`}
				visible={isFrame}
				handleBack={() =>
					frame(
						index !== COLUMNS.SIDEBAR_INDEX ? `${token.symbol}-transfer-recipient` : `${token.symbol}-token`
					)
				}
				hasChildrenPadding={false}
				hasOverlay
			>
				<div className="mb-4 flex flex-col gap-2">
					{index !== COLUMNS.SIDEBAR_INDEX && (
						<div className="px-6">
							<TransferRecipient
								address={column?.transfer?.recipient ?? ""}
								handleSelect={() => frame(`${token.symbol}-transfer-recipient`)}
							/>
						</div>
					)}

					<div className="flex flex-col gap-2">
						{token.implementations.map((implementation, implementationIndex) => (
							<ImplementationComponent
								key={implementationIndex}
								index={index}
								implementation={implementation}
								token={token}
								color={color}
							/>
						))}
					</div>

					<div className="mx-6 mt-2 flex flex-col gap-4">
						<ScrollingError error={error?.message ?? ""} />

						<button
							className={cn(
								"flex w-full items-center justify-center gap-2 rounded-lg border-[1px] py-4 font-bold transition-all duration-200 ease-in-out hover:opacity-90 hover:brightness-105",
								isReady === false && "transparent"
							)}
							style={{
								backgroundColor: isReady ? color : "transparent",
								color: isReady ? textColor : color,
								borderColor: isReady ? "#FFFFFF" : color
							}}
							disabled={(intent && isPending) || isReady === false}
							onClick={intent && !isPending && isReady ? handleTransaction : () => {}}
						>
							{!isAuthenticated
								? "Connect Wallet"
								: isPending
									? "Transfering..."
									: isReady
										? index === COLUMNS.SIDEBAR_INDEX
											? "Deposit"
											: "Send"
										: "Enter Amount"}
						</button>
					</div>
				</div>
			</Frame>
		</>
	)
}
