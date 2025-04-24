import { FC, useMemo, useState } from "react"

import { ArrowDownFromLine, ArrowRightLeft, ArrowUpToLine, ChartPie, RefreshCcw, Send } from "lucide-react"

import { useAtom, useAtomValue } from "jotai"

import { SocketTokenPriceChart } from "@/components/app/frames/assets/token/chart"
import { TokenFrameExternalLink } from "@/components/app/frames/assets/token/link"
import { Frame } from "@/components/app/frames/base"
import { ChainImage } from "@/components/app/sockets/chains/chain.image"
import { TokenImage } from "@/components/app/sockets/tokens/token-image"
import { Counter } from "@/components/shared/utils/counter"
import { cn, formatNumber, formatTimestamp, formatTitle, getZerionTokenIconUrl, ZerionPosition } from "@/lib"
import { api } from "@/server/client"
import { useSocket } from "@/state/authentication"
import { columnByIndexAtom, COLUMNS, isFrameAtom, useColumnActions } from "@/state/columns"

export const TokenFrame: FC<{
	index: number
	token: ZerionPosition
	color: string
	textColor: string
}> = ({ index, token, color, textColor }) => {
	const { isAnonymous } = useSocket()

	const [column] = useAtom(columnByIndexAtom(index))
	const frameKey = `${token.attributes.fungible_info.symbol}-token`
	const isFrame = useAtomValue(isFrameAtom)(column, frameKey)
	const { frame, transfer } = useColumnActions(index)

	const { data } = api.service.zerion.fungibles.detail.useQuery(
		{
			path: { fungibleId: token.relationships.fungible.data.id }
		},
		{ enabled: !isAnonymous && isFrame, retry: false, placeholderData: prev => prev }
	)
	const details = data?.data

	const [header, setHeader] = useState<{ title?: string; change?: number }>({})
	const [tooltipData, setTooltipData] = useState<
		| {
				timestamp: string
				price: number
				start: number
		  }
		| undefined
	>()

	const change = useMemo(() => {
		if (!token) return undefined

		if (header.change && !tooltipData) return header.change

		if (tooltipData) {
			const start = tooltipData.start
			const percentageChange = (tooltipData.price - start) / start
			return percentageChange * 100
		}

		return token.attributes.changes?.percent_1d
	}, [header, tooltipData, token])

	return (
		<Frame
			index={index}
			icon={
				<div className="relative h-8 w-10">
					{token.attributes.fungible_info.implementations &&
						token.attributes.fungible_info.implementations.length > 0 && (
							<TokenImage
								logo={getZerionTokenIconUrl(token)}
								symbol={token.attributes.fungible_info.symbol}
							/>
						)}
				</div>
			}
			label={token.attributes.fungible_info.name}
			visible={isFrame}
			hasChildrenPadding={false}
			hasOverlay
		>
			<div className="flex flex-row px-6 font-bold">
				<div className="flex flex-col items-center font-bold">
					<div className="mr-auto flex w-max flex-row text-lg">
						{tooltipData || token.attributes.price ? (
							<p className="flex flex-row items-center">
								$
								<Counter
									count={(tooltipData?.price || token.attributes.price || 0).toLocaleString("en-US", {
										minimumFractionDigits: 2
									})}
								/>
							</p>
						) : (
							<>-</>
						)}
					</div>
					<p className="mr-auto" style={{ color }}>
						{`$${token.attributes.fungible_info.symbol}`}
					</p>
				</div>
				<div
					className={cn(
						"ml-auto flex flex-col items-center",
						change === undefined ? "opacity-60" : change >= 0 ? "text-chart-green" : "text-plug-red"
					)}
				>
					<div className="ml-auto flex w-max flex-row text-lg font-bold">
						<Counter count={change || 0} decimals={2} />
						<p>%</p>
					</div>
					<p className="ml-auto flex flex-row items-center">
						{tooltipData ? formatTimestamp(Number(tooltipData.timestamp)) : (header.title ?? "Today")}
					</p>
				</div>
			</div>

			<SocketTokenPriceChart
				enabled={isFrame}
				token={token}
				color={color}
				handleHeader={setHeader}
				handleTooltip={setTooltipData}
			/>

			<div className="flex flex-row gap-2 px-6 pt-4">
				{token.attributes.fungible_info.implementations &&
					token.attributes.fungible_info.implementations
						.filter(implementation => implementation.chain_id === "base")
						.some(implementation => implementation.balance && implementation.balance > 0) && (
						<button
							className={cn(
								"flex items-center justify-center gap-2 rounded-lg border-[1px] px-12 py-4 font-bold transition-all duration-200 ease-in-out hover:opacity-90",
								index === COLUMNS.SIDEBAR_INDEX ? "w-full" : "w-max"
							)}
							style={{
								backgroundColor: index === COLUMNS.SIDEBAR_INDEX ? color : "",
								borderColor: color ?? "",
								color: index === COLUMNS.SIDEBAR_INDEX ? textColor : (color ?? "")
							}}
							onClick={() => {
								transfer(undefined)
								frame(
									index === COLUMNS.SIDEBAR_INDEX
										? `${token.attributes.fungible_info.symbol}-transfer-deposit`
										: `${token.attributes.fungible_info.symbol}-transfer-recipient`
								)
							}}
						>
							{index === COLUMNS.SIDEBAR_INDEX ? (
								<>
									<ArrowDownFromLine size={14} className="opacity-60" />
									Deposit
								</>
							) : (
								<>
									<Send size={14} className="opacity-60" />
									Send
								</>
							)}
						</button>
					)}

				<button
					className="flex w-full items-center justify-center gap-2 rounded-lg border-[1px] py-4 font-bold transition-all duration-200 ease-in-out hover:opacity-90"
					style={{
						backgroundColor: index !== COLUMNS.SIDEBAR_INDEX ? color : "",
						borderColor: color ?? "",
						color: index !== COLUMNS.SIDEBAR_INDEX ? textColor : (color ?? "")
					}}
					onClick={() => {
						transfer(undefined)
						frame(`${token.attributes.fungible_info.symbol}-swap-token`)
					}}
				>
					<ArrowRightLeft size={14} className="opacity-60" />
					Swap
				</button>
			</div>

			<div className="flex flex-col px-6 pb-2 pt-4 font-bold">
				<div className="flex flex-row items-center gap-4">
					<p className="mr-auto opacity-40">Balance</p>
					<div className="h-[2px] w-full bg-plug-green/10" />
					<p className="ml-auto opacity-40">Value</p>
				</div>

				<div className="mt-2 flex flex-row items-center justify-between gap-4">
					<div className="mr-auto flex h-8 items-center" style={{ color: color }}>
						<TokenImage
							logo={getZerionTokenIconUrl(token)}
							symbol={token.attributes.fungible_info.symbol}
							size="xs"
						/>
						<Counter className="ml-4 mr-2 w-max" count={token.attributes.quantity.float ?? 0} />
						<p>{token.attributes.fungible_info.symbol}</p>
					</div>
					<div className="ml-auto flex h-8 items-center text-center">
						<p>$</p>
						<Counter
							count={
								tooltipData
									? token.attributes.quantity.float * tooltipData.price
									: (token.attributes.value ?? 0)
							}
							decimals={2}
						/>
					</div>
				</div>
			</div>

			<div className="flex flex-row items-center gap-4 px-6 font-bold">
				<p className="opacity-40">Distribution</p>
				<div className="h-[2px] w-full bg-plug-green/10" />
			</div>

			<div className="relative mt-2 flex w-full flex-col gap-2 px-6 pb-4">
				{token.attributes.fungible_info.implementations
					?.filter(impl => impl.balance && impl.balance > 0)
					.map((implementation, implementationIndex) => {
						return (
							<div key={implementationIndex} className="flex flex-row items-center gap-4">
								<ChainImage chainId={implementation.chain_id} size="sm" />
								<p className="mr-auto font-bold">{formatTitle(implementation.chain_id ?? "Unknown")}</p>
								<p className="flex flex-col font-bold opacity-60">
									<Counter count={implementation.balance || 0} />
								</p>
								<p className="flex min-w-[72px] flex-row items-center text-right font-bold">
									<Counter count={implementation.percentage || 0} decimals={2} />%
								</p>
							</div>
						)
					})}
			</div>

			<div className="flex flex-col gap-2 px-6 pb-4">
				<div className="flex flex-row items-center gap-4 font-bold">
					<p className="opacity-40">Market</p>
					<div className="h-[2px] w-full bg-plug-green/10" />
				</div>

				<p className="flex w-full flex-row items-center gap-4 font-bold">
					<RefreshCcw size={18} className="opacity-20" />
					<span className="mr-auto truncate whitespace-nowrap opacity-40">Supply</span>
					<div className="flex flex-col font-bold opacity-60">
						<Counter
							count={formatNumber(details?.attributes.market_data?.circulating_supply ?? 0)}
							decimals={0}
						/>
					</div>

					<div className="flex min-w-[72px] flex-row items-center text-right font-bold">
						<Counter
							count={(
								((details?.attributes.market_data?.circulating_supply ?? 0) /
									(details?.attributes.market_data?.total_supply ?? 1)) *
								100
							).toLocaleString("en-US", {
								minimumFractionDigits: 2
							})}
							decimals={2}
						/>
						%
					</div>
				</p>

				<p className="flex w-full flex-row items-center gap-4 font-bold">
					<ArrowUpToLine size={18} className="opacity-20" />
					<span className="mr-auto truncate whitespace-nowrap opacity-40">Market Cap</span>
					<span className="flex flex-row items-center whitespace-nowrap">
						$<Counter count={formatNumber(details?.attributes.market_data?.market_cap ?? 0)} decimals={2} />
					</span>
				</p>

				<p className="flex w-full flex-row items-center gap-4 font-bold">
					<ChartPie size={18} className="opacity-20" />
					<span className="mr-auto truncate whitespace-nowrap opacity-40">Fully Diluted Valuation</span>
					<span className="flex flex-row items-center whitespace-nowrap">
						$
						<Counter
							count={formatNumber(details?.attributes.market_data?.fully_diluted_valuation ?? 0)}
							decimals={2}
						/>
					</span>
				</p>
			</div>

			<div className="flex flex-col gap-2 px-6 pb-4">
				<div className="flex flex-row items-center gap-4 font-bold">
					<p className="whitespace-nowrap opacity-40">External Links</p>
					<div className="h-[2px] w-full bg-plug-green/10" />
				</div>

				{details?.attributes.external_links?.map((link, linkIndex) => (
					<TokenFrameExternalLink key={linkIndex} link={link} />
				))}
			</div>
		</Frame>
	)
}
