import { FC, useMemo, useState } from "react"

import { ArrowDownFromLine, ArrowRightLeft, MapIcon, Send } from "lucide-react"

import { useAtom, useAtomValue } from "jotai"

import { Frame } from "@/components/app/frames/base"
import { SocketTokenPriceChart } from "@/components/app/sockets/tokens/token-chart"
import { TokenImage } from "@/components/app/sockets/tokens/token-image"
import { Image } from "@/components/app/utils/image"
import { Counter } from "@/components/shared/utils/counter"
import { chains, cn, formatTitle, getBlockExplorerAddress, getChainId } from "@/lib"
import { RouterOutputs } from "@/server/client"
import { columnByIndexAtom, COLUMNS, isFrameAtom, useColumnActions } from "@/state/columns"

import { ChainImage } from "../../sockets/chains/chain.image"

export const TokenFrame: FC<{
	index: number
	token?: NonNullable<RouterOutputs["socket"]["balances"]["positions"]>["tokens"][number]
	color: string
	textColor: string
}> = ({ index, token, color, textColor }) => {
	const [column] = useAtom(columnByIndexAtom(index))
	const frameKey = `${token?.symbol}-token`
	const isFrame = useAtomValue(isFrameAtom)(column, frameKey)
	const { frame, transfer } = useColumnActions(index)

	const [header, setHeader] = useState<{
		title?: string
		change?: Array<number>
	}>({
		title: undefined,
		change: undefined
	})
	const [tooltipData, setTooltipData] = useState<
		| {
				timestamp: string
				price: number
				start: Array<number>
		  }
		| undefined
	>()

	const change = useMemo(() => {
		if (!token) return undefined

		if (header.change && !tooltipData) return header.change[0]

		if (tooltipData) {
			const start = tooltipData.start[0]
			const percentageChange = (tooltipData.price - start) / start
			return percentageChange * 100
		}

		return token.change
	}, [header, tooltipData, token])

	const formatTimestamp = (timestamp: number) => {
		if (isNaN(timestamp)) return

		const milliseconds = timestamp.toString().length === 10 ? timestamp * 1000 : timestamp
		const date = new Date(milliseconds)

		const formatter = new Intl.DateTimeFormat("en-US", {
			year: "numeric",
			month: "short",
			day: "numeric",
			hour: "2-digit",
			minute: "2-digit",
			hour12: true
		})

		return formatter.format(date)
	}

	if (token === undefined) return null

	return (
		<Frame
			index={index}
			icon={
				<div className="relative h-8 w-10">
					<TokenImage
						logo={
							// @ts-ignore
							token?.icon?.url ||
							token?.icon ||
							`https://token-icons.llamao.fi/icons/tokens/${getChainId(token.implementations[0].chain)}/${token.implementations[0].contract}?h=240&w=240`
						}
						symbol={token.symbol}
						size="sm"
					/>
				</div>
			}
			label={token.name}
			visible={isFrame}
			hasChildrenPadding={false}
			hasOverlay
		>
			<div className="flex flex-row px-6 font-bold">
				<div className="flex flex-col items-center font-bold">
					<div className="mr-auto flex w-max flex-row text-lg">
						{tooltipData || token.price ? (
							<>
								<p>$</p>
								<Counter
									count={(tooltipData?.price || token.price || 0).toLocaleString("en-US", {
										minimumFractionDigits: 2
									})}
								/>
							</>
						) : (
							<>-</>
						)}
					</div>
					<p
						className="mr-auto"
						style={{
							color
						}}
					>
						{`$${token.symbol}`}
					</p>
				</div>
				<div
					className={cn(
						"ml-auto flex flex-col items-center",
						change === undefined ? "opacity-60" : change >= 0 ? "text-plug-green" : "text-red-500"
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
				keys={[`${token.implementations[0].chain}:${token.implementations[0].contract}`]}
				colors={{
					[`${token.implementations[0].chain}:${token.implementations[0].contract}`]: color
				}}
				handleHeader={setHeader}
				handleTooltip={setTooltipData}
			/>

			<div className="flex flex-row gap-2 px-6 pt-4">
				{token.implementations.some(implementation => implementation.balance) && (
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
									? `${token.symbol}-transfer-deposit`
									: `${token.symbol}-transfer-recipient`
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
						frame(`${token.symbol}-swap-token`)
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
							logo={
								// @ts-ignore
								token?.icon?.url ||
								token?.icon ||
								`https://token-icons.llamao.fi/icons/tokens/${getChainId(token.implementations[0].chain)}/${token.implementations[0].contract}?h=240&w=240`
							}
							symbol={token.symbol}
							size="xs"
						/>
						<Counter className="ml-4 mr-2 w-max" count={token.balance ?? 0} />
						<p>{token.symbol}</p>
					</div>
					<div className="ml-auto flex h-8 items-center text-center">
						<p>$</p>
						<Counter
							count={tooltipData ? token.balance * tooltipData.price : (token.value ?? 0)}
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
				{token.implementations.map((implementation, index) => (
					<div key={index} className="flex flex-row items-center gap-4">
						<ChainImage chainId={getChainId(implementation.chain)} size="sm" />

						<p className="mr-auto font-bold">{formatTitle(implementation.chain ?? "Unknown")}</p>

						<div className="flex flex-col font-bold opacity-60">
							<Counter count={implementation.balance ?? 0} />
						</div>

						<div className="flex min-w-[72px] flex-row items-center text-right font-bold">
							<Counter count={implementation.percentage ?? 0} />%
						</div>
					</div>
				))}
			</div>

			<div className="flex flex-row items-center gap-4 px-6 font-bold">
				<p className="opacity-40">Links</p>
				<div className="h-[2px] w-full bg-plug-green/10" />
			</div>

			<div className="relative mt-2 flex w-full flex-wrap gap-2 px-6 pb-4">
				{token.implementations
					// @ts-ignore
					.map((implementation, index) => (
						<div key={index} className="flex flex-row items-center gap-4">
							<a
								className="flex flex-row items-center gap-2 rounded-md px-4 py-2 text-xs font-bold transition-all duration-200 ease-in-out hover:opacity-90"
								style={{
									backgroundColor: color ?? "",
									color: textColor
								}}
								href={getBlockExplorerAddress(
									getChainId(implementation.chain),
									implementation.contract
								)}
								target="_blank"
								rel="noreferrer"
							>
								{token.implementations.length > 1 ? (
									<Image
										src={chains[getChainId(implementation.chain)].logo}
										alt={implementation.chain}
										className="h-4 w-4 rounded-full"
										width={24}
										height={24}
									/>
								) : (
									<MapIcon size={14} className="opacity-60" />
								)}
								Explorer
							</a>
						</div>
					))}
			</div>
		</Frame>
	)
}
