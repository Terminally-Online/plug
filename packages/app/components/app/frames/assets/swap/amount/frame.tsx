import { useCallback, useMemo, useState } from "react"

import { formatUnits, getAddress } from "viem"
import { useSendTransaction } from "wagmi"

import { ArrowRight, Bell, CheckCircle, CircleDollarSign, Loader, TriangleRight, Waypoints } from "lucide-react"

import { useAtom, useAtomValue } from "jotai"

import { SwapAmountInput } from "@/components/app/frames/assets/swap/amount/input"
import { Frame } from "@/components/app/frames/base"
import { ChainImage } from "@/components/app/sockets/chains/chain.image"
import { TokenImage } from "@/components/app/sockets/tokens/token-image"
import { Counter } from "@/components/shared/utils/counter"
import {
	cn,
	getChainId,
	getChainName,
	getTextColor,
	getZerionTokenIconUrl,
	NATIVE_TOKEN_ADDRESS,
	useStateDebounce,
	ZerionFungible,
	ZerionPosition
} from "@/lib"
import { useAllowance } from "@/lib/hooks/chain/useApproval"
import { api } from "@/server/client"
import { useSocket } from "@/state/authentication"
import { columnByIndexAtom, COLUMNS, isFrameAtom, useColumnActions } from "@/state/columns"

import { ScrollingError } from "../../scrolling-error"

type SwapAmountFrameProps = {
	index: number
	tokenIn: ZerionFungible
	tokenOut: ZerionPosition
}

const getFungibleImplementation = (
	implementations: ZerionPosition["attributes"]["fungible_info"]["implementations"],
	chainName: string
) => {
	const implementation = implementations.find(implementation => implementation.chain_id === chainName)
	const address = getAddress(implementation?.address ?? NATIVE_TOKEN_ADDRESS)
	const decimals = implementation?.decimals ?? 18
	const lookup = `${address}:${decimals}:${20}`

	return { ...implementation, address, decimals, lookup }
}

export const SwapAmountFrame = ({ index, tokenIn, tokenOut }: SwapAmountFrameProps) => {
	const implementations = useMemo(
		() => ({
			tokenOut: getFungibleImplementation(tokenOut.attributes.fungible_info.implementations, "base"),
			tokenIn: getFungibleImplementation(tokenIn.attributes.implementations, "base")
		}),
		[tokenIn, tokenOut]
	)

	const [column] = useAtom(columnByIndexAtom(index))
	const frameKey = `${tokenOut.attributes.fungible_info.symbol}-${tokenIn.attributes.symbol}-swap-amount`
	const isFrame = useAtomValue(isFrameAtom)(column, frameKey)
	const { frame, navigate } = useColumnActions(index, frameKey)

	const { socket } = useSocket()
	const { error, sendTransaction, isPending } = useSendTransaction()

	const balance = implementations.tokenOut?.balance ?? 0

	const [amounts, setAmounts] = useState({
		[tokenOut.attributes.fungible_info.symbol]: {
			precise: (balance / 2).toString(),
			percentage: balance / 2 === 0 ? 0 : 50
		},
		[tokenIn.attributes.symbol]: {
			precise: "0",
			percentage: 0
		}
	})
	const [step, setStep] = useState(0)
	const [tokenOutColor, setTokenOutColor] = useState("#000000")
	const [tokenInColor, setTokenInColor] = useState("#000000")

	const isSufficientBalance =
		balance > 0 && balance >= Number(amounts[tokenOut.attributes.fungible_info.symbol].precise)

	const isEOA = column && column.index === COLUMNS.SIDEBAR_INDEX
	const from = socket ? (column && column.index === COLUMNS.SIDEBAR_INDEX ? socket.id : socket.socketAddress) : ""

	const amount = useStateDebounce(amounts[tokenOut.attributes.fungible_info.symbol].precise)

	const {
		data: intent,
		error: intentError,
		isLoading,
		isFetching
	} = api.solver.actions.intent.useQuery(
		{
			chainId: getChainId(implementations.tokenOut?.chain_id ?? "base"),
			from,
			inputs: [
				{
					protocol: "plug",
					action: "swap",
					amount,
					token: implementations.tokenOut.lookup,
					tokenIn: implementations.tokenIn.lookup
				}
			],
			options: {
				isEOA: isEOA,
				simulate: true
			}
		},
		{
			enabled:
				isFrame &&
				!!socket &&
				!!implementations.tokenIn &&
				!!implementations.tokenOut &&
				amounts[tokenOut.attributes.fungible_info.symbol].precise !== "0" &&
				isSufficientBalance,
			placeholderData: prev => prev
		}
	)

	const isReady =
		amounts[tokenOut.attributes.fungible_info.symbol].precise !== "0" &&
		!intentError &&
		!isLoading &&
		!isFetching &&
		isSufficientBalance &&
		!isPending
	const meta = intent ? intent.transactions[intent.transactions.length - 1].meta : null

	const { approval } = useAllowance({
		token: implementations.tokenOut.address,
		owner: from,
		spender: meta?.settlementAddress
	})

	const toggleSavedMutation = api.plugs.activity.toggleSaved.useMutation()
	const handleTransactionOffchain = useCallback(async () => {
		if (!intent) return

		if (step === 0) toggleSavedMutation.mutateAsync({ id: intent.intentId })

		const handleTransactionRedirect = () => {
			navigate({ index, key: COLUMNS.KEYS.ACTIVITY })
			frame(`${intent.intentId}-activity`)
		}

		if (step === intent.transactions.length - 1) handleTransactionRedirect()
		else setStep(prev => prev + 1)
	}, [intent, step, toggleSavedMutation, navigate, frame, index])

	const isApproved = approval > BigInt(meta?.sellTokens[implementations.tokenOut?.address]?.amount ?? "0")

	const handleTransactionOnchain = useCallback(async () => {
		if (!column || !intent) return

		const transactionAtStep = isApproved && intent.transactions.length > 1 ? 1 : step

		if (column.index === COLUMNS.SIDEBAR_INDEX)
			sendTransaction(
				{
					to: intent.transactions[transactionAtStep].to,
					data: intent.transactions[transactionAtStep].data,
					value: intent.transactions[transactionAtStep].value
				},
				{
					onSuccess: handleTransactionOffchain
				}
			)
	}, [column, intent, step, isApproved, sendTransaction, handleTransactionOffchain])

	return (
		<Frame
			index={index}
			icon={
				<div className="group relative flex flex-row items-center">
					<div className="relative h-8 w-10">
						<TokenImage
							logo={getZerionTokenIconUrl(tokenOut)}
							symbol={tokenOut.attributes.fungible_info.symbol}
							size="sm"
							handleColor={setTokenOutColor}
						/>
					</div>
					<div className="relative -ml-4 h-8 w-10 transition-all duration-100 group-hover:ml-0">
						<TokenImage
							logo={getZerionTokenIconUrl(tokenIn.attributes.icon?.url)}
							symbol={tokenIn.attributes.symbol}
							size="sm"
							handleColor={setTokenInColor}
						/>
					</div>
				</div>
			}
			label={
				<span className="flex flex-row items-center gap-2 text-lg font-bold">
					<span>{tokenOut.attributes.fungible_info.symbol}</span>
					<ArrowRight size={14} className="opacity-60" />
					<span>{tokenIn.attributes.symbol}</span>
				</span>
			}
			visible={isFrame}
			handleBack={() => frame(`${tokenOut.attributes.fungible_info.symbol}-swap-token`)}
			hasChildrenPadding={false}
			hasOverlay
		>
			<div className="relative mb-2 flex flex-col gap-2">
				<SwapAmountInput
					index={index}
					token={{
						...tokenOut,
						attributes: {
							...tokenOut.attributes,
							price: meta?.sellTokens[implementations.tokenOut.address]?.priceUsd ?? 0
						}
					}}
					color={tokenOutColor}
					amounts={amounts[tokenOut.attributes.fungible_info.symbol]}
					setAmounts={amount =>
						setAmounts(prev => {
							if (typeof amount === "function") {
								return {
									...prev,
									[tokenOut.attributes.fungible_info.symbol]: amount(
										prev[tokenOut.attributes.fungible_info.symbol]
									)
								}
							}
							return {
								...prev,
								[tokenOut.attributes.fungible_info.symbol]: amount
							}
						})
					}
				/>

				<SwapAmountInput
					index={index}
					token={{
						...tokenIn,
						id: tokenIn.id,
						type: "positions",
						attributes: {
							parent: null,
							protocol: null,
							name: tokenIn.attributes.name,
							position_type: "wallet",
							quantity: {
								int: "0",
								decimals: implementations.tokenIn.decimals,
								float: 0,
								numeric: "0"
							},
							value: 0,
							price: meta?.buyTokens[implementations.tokenIn?.address]?.priceUsd ?? null,
							changes: {
								absolute_1d: 0,
								percent_1d: 0
							},
							flags: {
								displayable: true,
								is_trash: false
							},
							updated_at: "",
							updated_at_block: 0,
							fungible_info: tokenIn.attributes
						},
						relationships: {
							fungible: {
								data: {
									type: "fungibles",
									id: tokenIn.id
								},
								links: {
									related: ""
								}
							},
							chain: {
								data: {
									type: "chains",
									id: implementations.tokenIn?.chain_id ?? "base"
								},
								links: {
									related: ""
								}
							}
						}
					}}
					color={tokenInColor}
					amounts={{
						precise: formatUnits(
							meta?.buyTokens[implementations.tokenIn.address]?.amount ?? "0",
							implementations.tokenIn.decimals
						).toString(),
						percentage: amounts[tokenIn.attributes.symbol].percentage
					}}
					setAmounts={amount =>
						setAmounts(prev => {
							if (typeof amount === "function") {
								return {
									...prev,
									[tokenIn.attributes.symbol]: amount(prev[tokenIn.attributes.symbol])
								}
							}
							return {
								...prev,
								[tokenIn.attributes.symbol]: amount
							}
						})
					}
				/>
			</div>

			<div className="px-6 pt-2">
				<div className="mb-2 flex flex-row items-center gap-4">
					<p className="font-bold opacity-40">Details</p>
					<div className="h-[1px] w-full bg-plug-green/10" />
				</div>

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
										: meta?.slippage >= 0
											? "text-chart-green"
											: "text-plug-red"
								)}
							>
								{meta?.slippage > 0 && "+"}
								<Counter count={meta?.slippage ?? "-"} />%
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
											? "text-chart-green"
											: "text-plug-red"
								)}
							>
								{meta?.priceImpact > 0 && "+"}
								<Counter count={meta?.priceImpact ? meta.priceImpact * 100 : "-"} />%
							</span>
						</span>
					</span>
				</p>

				{implementations.tokenOut?.chain_id && (
					<p className="flex flex-row justify-between font-bold">
						<span className="flex w-max flex-row items-center gap-4">
							<Waypoints size={18} className="opacity-20" />
							<span className="opacity-40">Chain</span>
						</span>{" "}
						<span className="flex flex-row items-center gap-2 font-bold">
							<ChainImage chainId={getChainId(implementations.tokenOut.chain_id)} size="xs" />
							{getChainName(getChainId(implementations.tokenOut.chain_id))}
						</span>
					</p>
				)}

				<p className="flex flex-row justify-between font-bold">
					<span className="flex w-full flex-row items-center gap-4">
						<CheckCircle size={18} className="opacity-20" />
						<span className="opacity-40">Approval</span>
					</span>{" "}
					<span className="flex flex-row items-center gap-1 font-bold tabular-nums">
						{implementations.tokenOut?.address === NATIVE_TOKEN_ADDRESS || isApproved
							? "Sufficient"
							: "Insufficient"}
					</span>
				</p>

				<p className="flex flex-row justify-between font-bold">
					<span className="flex w-full flex-row items-center gap-4">
						<CircleDollarSign size={18} className="opacity-20" />
						<span className="opacity-40">Fee</span>
					</span>{" "}
					<span className="flex flex-row items-center gap-1 font-bold tabular-nums">
						<span className="ml-auto flex flex-row items-center gap-1 pl-2 opacity-40">
							<Counter count={0.0} /> ETH
						</span>
						<span className="ml-2 flex flex-row items-center">Free</span>
					</span>
				</p>
			</div>

			<div className="mx-6 my-4 flex flex-col gap-4">
				<ScrollingError error={error?.message ?? ""} />

				<button
					className={cn(
						"flex w-full items-center justify-center gap-2 rounded-lg border-[1px] px-12 py-4 font-bold transition-all duration-200 ease-in-out hover:opacity-90 hover:brightness-105"
					)}
					style={{
						backgroundColor: !isReady ? "transparent" : tokenInColor,
						color: !isReady ? tokenInColor : getTextColor(tokenInColor),
						borderColor: !isReady ? tokenInColor : "transparent"
					}}
					disabled={!isReady || isPending || isLoading || isFetching}
					onClick={handleTransactionOnchain}
				>
					{!isSufficientBalance ? (
						"Insufficient Balance"
					) : amounts[tokenOut.attributes.fungible_info.symbol].precise === "0" ? (
						"Enter Amount"
					) : isLoading || isFetching ? (
						<span className="flex flex-row items-center gap-2">
							<Loader
								size={14}
								className="animate-spin opacity-60"
								style={{ color: getTextColor(tokenInColor) }}
							/>
							<span>Routing...</span>
						</span>
					) : intentError || !implementations.tokenOut || !implementations.tokenIn ? (
						"Route could not be found"
					) : isEOA && intent && intent.transactions.length > 1 && step === 0 && !isApproved ? (
						"Approve"
					) : (
						"Swap"
					)}
				</button>
			</div>
		</Frame>
	)
}
