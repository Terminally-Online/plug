import { useCallback, useEffect, useMemo, useState } from "react"

import { formatUnits, getAddress } from "viem"

import { ArrowRight, Bell, CircleDollarSign, Loader, TriangleRight, Waypoints } from "lucide-react"

import { useAtom, useAtomValue } from "jotai"

import { SwapAmountInput } from "@/components/app/frames/assets/swap.amount.input"
import { Frame } from "@/components/app/frames/base"
import { TokenImage } from "@/components/app/sockets/tokens/token-image"
import { Counter } from "@/components/shared/utils/counter"
import { cn, getChainId, getChainName, getTextColor, NATIVE_TOKEN_ADDRESS, useDebounceInline } from "@/lib"
import { api, RouterOutputs } from "@/server/client"
import { useSocket } from "@/state/authentication"
import { columnByIndexAtom, COLUMNS, isFrameAtom, useColumnActions } from "@/state/columns"

import { ChainImage } from "../../sockets/chains/chain.image"
import { useSendTransaction } from "wagmi"

type Token =
	| NonNullable<RouterOutputs["socket"]["balances"]["positions"]>["tokens"][number]
	| NonNullable<RouterOutputs["solver"]["tokens"]["get"]>[number]

type SwapAmountFrameProps = {
	index: number
	tokenIn: Token
	tokenOut: Token
}

export const SwapAmountFrame = ({ index, tokenIn, tokenOut }: SwapAmountFrameProps) => {
	const [column] = useAtom(columnByIndexAtom(index))
	const frameKey = `${tokenOut.symbol}-${tokenIn.symbol}-swap-amount`
	const isFrame = useAtomValue(isFrameAtom)(column, frameKey)
	const { frame, navigate } = useColumnActions(index, frameKey)

	const { socket } = useSocket()
	const { error, sendTransaction, isPending } = useSendTransaction()

	const { tokenOutImplementation, tokenInImplementation } = useMemo(() => {
		const tokenOutImplementation = tokenOut.implementations.find(implementation => implementation.chain === "base")
		const tokenInImplementation = tokenIn.implementations.find(implementation => implementation?.chain === "base")

		return { tokenOutImplementation, tokenInImplementation }
	}, [tokenIn, tokenOut])

	const [tokenOutColor, setTokenOutColor] = useState("#000000")
	const [tokenInColor, setTokenInColor] = useState("#000000")
	const [amounts, setAmounts] = useState({
		[tokenOut.symbol]: {
			precise: ((tokenOutImplementation?.balance ?? 0) / 2).toString(),
			percentage: (tokenOutImplementation?.balance ?? 0) / 2 === 0 ? 0 : 50
		},
		[tokenIn.symbol]: {
			precise: "0",
			percentage: 0
		}
	})

	const isSufficientBalance =
		tokenOutImplementation &&
		(tokenOutImplementation?.balance ?? 0) > 0 &&
		(tokenOutImplementation?.balance ?? 0) >= Number(amounts[tokenOut.symbol].precise)

	const request = useDebounceInline<{
		chainId: number,
		from: string | `0x${string}`
		inputs: Array<{
			protocol: string,
			action: string,
			[key: string]: string
		}>, options?: { isEOA?: boolean, simulate?: boolean }
	}>({
		chainId: getChainId(tokenOutImplementation?.chain ?? "base"),
		from: socket ? (column && column.index === COLUMNS.SIDEBAR_INDEX ? socket.id : socket.socketAddress) : "",
		inputs: [
			{
				protocol: "plug",
				action: "swap",
				amount: amounts[tokenOut.symbol].precise,
				token: `${getAddress(tokenOutImplementation?.contract ?? NATIVE_TOKEN_ADDRESS)}:${tokenOutImplementation?.decimals ?? 18}:${20}`,
				tokenIn: `${getAddress(tokenInImplementation?.contract ?? NATIVE_TOKEN_ADDRESS)}:${tokenInImplementation?.decimals ?? 18}:${20}`
			}
		],
		options: {
			isEOA: column && column.index === COLUMNS.SIDEBAR_INDEX,
			simulate: true
		}
	})


	console.log('request', request)

	const { data: intent, error: intentError, isLoading } = api.solver.actions.intent.useQuery(request, {
		enabled:
			isFrame &&
			!!tokenInImplementation &&
			!!tokenOutImplementation &&
			amounts[tokenOut.symbol].precise !== "0" &&
			isSufficientBalance &&
			!!socket,
	})

	const isReady =
		amounts[tokenOut.symbol].precise !== "0" && !intentError && !isLoading && isSufficientBalance
	const meta = intent ? intent.transactions[0].meta : null

	const toggleSavedMutation = api.plugs.activity.toggleSaved.useMutation()
	const handleTransactionOffchain = useCallback(() => {
		if (!intent) return

		toggleSavedMutation.mutate({ id: intent.intentId }, {
			onSuccess: () => {
				navigate({ index, key: COLUMNS.KEYS.ACTIVITY })
				frame(`${intent.intentId}-activity`)
			}
		})
	}, [intent])

	const handleTransactionOnchain = useCallback(() => {
		if (!column || !intent) return

		if (column.index === COLUMNS.SIDEBAR_INDEX)
			sendTransaction(
				{
					to: intent.transactions[0].to,
					data: intent.transactions[0].data,
					value: intent.transactions[0].value
				},
				{
					onSuccess: handleTransactionOffchain
				}
			)
		else
			handleTransactionOffchain()
	}, [column, intent])

	return (
		<Frame
			index={index}
			icon={
				<div className="group relative flex flex-row items-center">
					<div className="relative h-8 w-10">
						<TokenImage
							logo={
								tokenOut?.icon ||
								`https://token-icons.llamao.fi/icons/tokens/${getChainId(tokenOut.implementations[0].chain)}/${tokenOut.implementations[0]?.contract ?? "0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE"}?h=240&w=240`
							}
							symbol={tokenOut.symbol}
							size="sm"
							handleColor={setTokenOutColor}
						/>
					</div>
					<div className="relative -ml-4 h-8 w-10 transition-all duration-100 group-hover:ml-0">
						<TokenImage
							logo={
								tokenIn?.icon ||
								`https://token-icons.llamao.fi/icons/tokens/${getChainId(tokenIn.implementations[0].chain)}/${tokenIn.implementations[0]?.contract ?? "0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE"}?h=240&w=240`
							}
							symbol={tokenIn.symbol}
							size="sm"
							handleColor={setTokenInColor}
						/>
					</div>
				</div>
			}
			label={
				<span className="flex flex-row items-center gap-2 text-lg font-bold">
					<span>{tokenOut.symbol}</span>
					<ArrowRight size={14} className="opacity-60" />
					<span>{tokenIn.symbol}</span>
				</span>
			}
			visible={isFrame}
			handleBack={() => frame(`${tokenOut.symbol}-swap-token`)}
			hasChildrenPadding={false}
			hasOverlay
		>
			<div>
				<div className="relative mb-2 flex flex-col gap-2">
					<SwapAmountInput
						index={index}
						token={{
							...tokenOut,
							price:
								meta?.sellTokens[
									getAddress(
										tokenOutImplementation?.contract ?? "0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE"
									)
								]?.priceUsd ?? 0
						}}
						color={tokenOutColor}
						amounts={amounts[tokenOut.symbol]}
						setAmounts={amount =>
							setAmounts(prev => {
								if (typeof amount === "function") {
									return {
										...prev,
										[tokenOut.symbol]: amount(prev[tokenOut.symbol])
									}
								}
								return {
									...prev,
									[tokenOut.symbol]: amount
								}
							})
						}
					/>

					<SwapAmountInput
						index={index}
						token={{
							...tokenIn,
							price:
								meta?.buyTokens[
									getAddress(
										tokenInImplementation?.contract ?? "0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE"
									)
								]?.priceUsd ?? 0,
							implementations: [
								// @ts-ignore
								{
									...tokenInImplementation,
									balance:
										meta?.buyTokens[
											getAddress(
												tokenInImplementation?.contract ??
												"0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE"
											)
										]?.amount ?? 0
								}
							]
						}}
						color={tokenInColor}
						amounts={{
							precise: formatUnits(
								meta?.buyTokens[
									getAddress(
										tokenInImplementation?.contract ?? "0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE"
									)
								]?.amount ?? "0",
								meta?.buyTokens[
									getAddress(
										tokenInImplementation?.contract ?? "0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE"
									)
								]?.decimals ?? 18
							).toString(),
							percentage: amounts[tokenOut.symbol].percentage
						}}
						setAmounts={amount =>
							setAmounts(prev => {
								if (typeof amount === "function") {
									return {
										...prev,
										[tokenIn.symbol]: amount(prev[tokenIn.symbol])
									}
								}
								return {
									...prev,
									[tokenIn.symbol]: amount
								}
							})
						}
					/>
				</div>
			</div>

			<div className="px-6">
				<div className="mb-2 flex flex-row items-center gap-4">
					<p className="font-bold opacity-40">Details</p>
					<div className="h-[2px] w-full bg-plug-green/10" />
				</div>

				{/* {tokenOutImplementation && tokenOutImplementation.chain && (
					<p className="flex flex-row justify-between font-bold">
						<span className="flex w-max flex-row items-center gap-4">
							<Waypoints size={18} className="opacity-20" />
							<span className="opacity-40">Chain</span>
						</span>{" "}
						<span className="flex flex-row items-center gap-2 font-bold">
							<ChainImage chainId={getChainId(tokenOutImplementation.chain)} size="xs" />
							{getChainName(getChainId(tokenOutImplementation.chain))}
						</span>
					</p>
				)} */}

				{meta && (
					<>
						<p className="flex flex-row justify-between font-bold tabular-nums">
							<span className="flex w-full flex-row items-center gap-4">
								<TriangleRight size={18} className="opacity-20" />
								<span className="opacity-40">Slippage</span>
							</span>{" "}
							<span className="flex flex-row items-center gap-1 font-bold tabular-nums">
								<span className="ml-auto flex flex-row items-center gap-1 pl-2">
									<span
										className={cn(
											"flex flex-row items-center whitespace-nowrap group-hover:hidden",
											meta?.slippage === undefined
												? "opacity-40"
												: meta?.priceImpact >= 0
													? ""
													: "text-red-500"
										)}
									>
										{meta?.slippage > 0 && "+"}
										<Counter count={meta?.slippage} />%
									</span>
								</span>
							</span>
						</p>

						<p className="flex flex-row justify-between font-bold tabular-nums">
							<span className="flex w-full flex-row items-center gap-4">
								<Bell size={18} className="opacity-20" />
								<span className="opacity-40">Price Impact</span>
							</span>{" "}
							<span className="flex flex-row items-center gap-1 font-bold tabular-nums">
								<span className="ml-auto flex flex-row items-center gap-1 pl-2">
									<span
										className={cn(
											"flex flex-row items-center whitespace-nowrap group-hover:hidden",
											meta?.priceImpact === undefined
												? "opacity-40"
												: meta?.priceImpact >= 0
													? ""
													: "text-red-500"
										)}
									>
										{meta?.priceImpact > 0 && "+"}
										<Counter count={meta?.priceImpact * 100} />%
									</span>
								</span>
							</span>
						</p>
					</>
				)}

				{tokenOutImplementation?.chain && (
					<p className="flex flex-row justify-between font-bold">
						<span className="flex w-max flex-row items-center gap-4">
							<Waypoints size={18} className="opacity-20" />
							<span className="opacity-40">Chain</span>
						</span>{" "}
						<span className="flex flex-row items-center gap-2 font-bold">
							<ChainImage chainId={getChainId(tokenOutImplementation.chain)} size="xs" />
							{getChainName(getChainId(tokenOutImplementation.chain))}
						</span>
					</p>
				)}

				<p className="flex flex-row justify-between font-bold">
					<span className="flex w-full flex-row items-center gap-4">
						<CircleDollarSign size={18} className="opacity-20" />
						<span className="opacity-40">Fee</span>
					</span>{" "}
					<span className="flex flex-row items-center gap-1 font-bold tabular-nums">
						<span className="ml-auto flex flex-row items-center gap-1 pl-2 opacity-40">
							<Counter count={0.00011} /> ETH
						</span>
						<span className="ml-2 flex flex-row items-center">
							$<Counter count={0.049} />
						</span>
					</span>
				</p>
			</div>

			<div className="mx-6 my-4 flex flex-col gap-4">
				<button
					className={cn(
						"flex w-full items-center justify-center gap-2 rounded-lg border-[1px] px-12 py-4 font-bold transition-all duration-200 ease-in-out hover:opacity-90 hover:brightness-105"
					)}
					style={{
						backgroundColor: !isReady ? "transparent" : tokenInColor,
						color: !isReady ? tokenInColor : getTextColor(tokenInColor),
						borderColor: !isReady ? tokenInColor : "transparent"
					}}
					disabled={!isReady}
					onClick={handleTransactionOffchain}
				>
					{!isSufficientBalance ? (
						"Insufficient Balance"
					) : amounts[tokenOut.symbol].precise === "0" ? (
						"Enter Amount"
					) : isLoading ? (
						<span className="flex flex-row items-center gap-2">
							<Loader
								size={14}
								className="animate-spin opacity-60"
								style={{ color: getTextColor(tokenInColor) }}
							/>
							<span>Routing...</span>
						</span>
					) : intentError || !tokenInImplementation || !tokenOutImplementation ? (
						"Route could not be found"
					) : (
						"Swap"
					)}
				</button>
			</div>
		</Frame>
	)
}
