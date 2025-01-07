import { useEffect, useMemo, useState } from "react"

import { formatUnits, getAddress, parseUnits } from "viem"

import { ArrowRight, ArrowRightLeft, Bell, Loader, TriangleRight } from "lucide-react"

import { Counter, Frame, SwapAmountInput, TokenImage } from "@/components"
import { cn, getChainId, getTextColor } from "@/lib"
import { api, RouterOutputs } from "@/server/client"
import { useColumnStore } from "@/state"

type Token =
	| NonNullable<RouterOutputs["socket"]["balances"]["positions"]>["tokens"][number]
	| NonNullable<RouterOutputs["solver"]["tokens"]["get"]>[number]

type SwapAmountFrameProps = {
	index: number
	tokenIn: Token
	tokenOut: Token
}

export const SwapAmountFrame = ({ index, tokenIn, tokenOut }: SwapAmountFrameProps) => {
	const {
		isFrame,
		handle: { frame }
	} = useColumnStore(index, `${tokenOut.symbol}-${tokenIn.symbol}-swap-amount`)

	const [tokenOutColor, setTokenOutColor] = useState("#000000")
	const [tokenInColor, setTokenInColor] = useState("#000000")

	const { tokenOutImplementation, tokenInImplementation } = useMemo(() => {
		const tokenInImplementation = tokenIn.implementations.find(
			implementation => implementation?.chain === "ethereum"
		)
		const tokenOutImplementation = tokenOut.implementations.find(
			implementation => implementation.chain === "ethereum"
		)

		return { tokenOutImplementation, tokenInImplementation }
	}, [tokenIn, tokenOut])

	const [amounts, setAmounts] = useState({
		[tokenOut.symbol]: {
			precise: ((tokenOut.implementations[0].balance ?? 0) / 2).toString(),
			percentage: 0
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

	const transaction = api.solver.actions.getTransaction.useQuery(
		{
			chainId: getChainId(tokenOutImplementation?.chain ?? "ethereum"),
			from: getAddress("0x62180042606624f02d8a130da8a3171e9b33894d"),
			inputs: [
				{
					protocol: "plug",
					action: "swap",
					tokenIn: getAddress(
						tokenInImplementation?.contract ?? "0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE"
					),
					tokenOut: getAddress(
						tokenOutImplementation?.contract ?? "0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE"
					),
					amountOut: parseUnits(
						amounts[tokenOut.symbol]?.precise,
						tokenOutImplementation?.decimals ?? 18
					).toString()
				}
			]
		},
		{
			enabled:
				isFrame &&
				!!tokenInImplementation &&
				!!tokenOutImplementation &&
				amounts[tokenOut.symbol].precise !== "0" &&
				isSufficientBalance,
			refetchInterval: 3500,
			staleTime: 1000
		}
	)

	const meta = useMemo(() => {
		if (!transaction.data) return null

		return transaction.data.plug.plugs[transaction.data.plug.plugs.length - 1].meta
	}, [transaction.data])

	const isReady =
		amounts[tokenOut.symbol].precise !== "0" && !transaction.error && !transaction.isLoading && isSufficientBalance

	useEffect(() => {
		setAmounts({
			[tokenOut.symbol]: {
				precise: ((tokenOut.implementations[0].balance ?? 0) / 2).toString(),
				percentage: (tokenOut.implementations[0].balance ?? 0) / 2 === 0 ? 0 : 50
			},
			[tokenIn.symbol]: {
				precise: "0",
				percentage: 0
			}
		})
	}, [tokenIn, tokenOut])

	return (
		<Frame
			index={index}
			icon={
				<div className="-gap-2 relative flex flex-row items-center">
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
					<div className="relative -ml-4 h-8 w-10">
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

				<div className="mb-2 flex flex-row items-center justify-between gap-4 px-6">
					{meta && (
						<p className="flex w-full flex-row items-center gap-2 font-bold tabular-nums">
							<span className="group flex w-max flex-row items-center gap-2">
								<TriangleRight size={14} className="opacity-40" />{" "}
								<span className="flex flex-row items-center whitespace-nowrap opacity-40 group-hover:hidden">
									{meta?.slippage >= 0 && "+"}
									<Counter count={meta?.slippage} />%
								</span>
								<span className="hidden opacity-40 group-hover:flex">Slippage</span>
							</span>
							<span className="opacity-20">|</span>
							<span className="group flex w-max flex-row items-center gap-2">
								<Bell size={14} className="opacity-40" />{" "}
								<span
									className={cn(
										"flex flex-row items-center whitespace-nowrap group-hover:hidden",
										meta?.priceImpact === undefined
											? "opacity-40"
											: meta?.priceImpact > 0
												? "text-plug-green"
												: "text-red-500"
									)}
								>
									{meta?.priceImpact >= 0 && "+"}
									<Counter count={meta?.priceImpact * 100} />%
								</span>
								<span className="hidden opacity-40 group-hover:flex">Price Impact</span>
							</span>
						</p>
					)}
					<button
						className="ml-auto font-bold text-black/40 hover:brightness-105"
						onClick={() =>
							setAmounts(prev => ({
								...prev,
								[tokenOut.symbol]: {
									percentage: 100,
									precise: (tokenOut.implementations[0]?.balance ?? 0).toString()
								}
							}))
						}
						style={{
							color:
								(amounts[tokenOut.symbol].percentage ?? 0) < 100 && isSufficientBalance
									? tokenInColor
									: undefined
						}}
						disabled={!isSufficientBalance}
					>
						Max
					</button>{" "}
				</div>
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
					onClick={() => frame(`${tokenOut.symbol}-${tokenIn.symbol}-swap-confirm`)}
				>
					{!isSufficientBalance ? (
						"Insufficient Balance"
					) : amounts[tokenOut.symbol].precise === "0" ? (
						"Enter Amount"
					) : transaction.isLoading ? (
						<span className="flex flex-row items-center gap-2">
							<Loader
								size={14}
								className="animate-spin opacity-60"
								style={{ color: getTextColor(tokenInColor) }}
							/>
							<span>Routing...</span>
						</span>
					) : transaction.error || !tokenInImplementation || !tokenOutImplementation ? (
						"Route could not be found"
					) : (
						"Swap"
					)}
				</button>
			</div>
		</Frame>
	)
}
