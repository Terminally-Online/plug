import { FC, useMemo, useState } from "react"

import Image from "next/image"

import { Send } from "lucide-react"

import { Counter, Frame, SocketTokenPriceChart } from "@/components"
import { useBalances, useFrame } from "@/contexts"
import { cn, formatTitle, getChainImage, getTextColor } from "@/lib"

import { TokenImage } from "../../sockets/tokens/token-image"

export const TokenFrame: FC<{ id: string; symbol: string }> = ({
	id,
	symbol
}) => {
	const { isFrame, handleFrame } = useFrame({ id, key: `token/${symbol}` })
	const { positions } = useBalances()
	const { tokens } = positions

	const token = useMemo(
		() =>
			tokens &&
			tokens.find(token => token.symbol === symbol && token.name),
		[tokens, symbol]
	)

	const [color, setColor] = useState("")
	const [header, setHeader] = useState<{
		title?: string
		change?: number
	}>({
		title: undefined,
		change: undefined
	})
	const [tooltipData, setTooltipData] = useState<
		| {
				timestamp: string
				price: number
				start: number
		  }
		| undefined
	>()

	const textColor = getTextColor(color ?? "#ffffff")

	const change = useMemo(() => {
		if (!token) return undefined

		if (header.change && !tooltipData) return header.change

		if (tooltipData) {
			const percentageChange =
				(tooltipData.price - tooltipData.start) / tooltipData.start

			return percentageChange * 100
		}

		return token.change
	}, [header, tooltipData, token])

	const formatTimestamp = (timestamp: number) => {
		if (isNaN(timestamp)) return

		const milliseconds =
			timestamp.toString().length === 10 ? timestamp * 1000 : timestamp
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
			id={id}
			className="overflow-x-hidden"
			icon={
				<div className="relative h-10 w-10">
					<TokenImage
						logo={token?.icon ?? ""}
						symbol={token.symbol}
						size="sm"
						handleColor={setColor}
					/>
				</div>
			}
			label=""
			visible={isFrame}
			hasChildrenPadding={false}
			hasOverlay
		>
			<div className="flex flex-row px-6 font-bold">
				<div className="flex flex-col items-center font-bold">
					<p className="mr-auto flex w-max flex-row text-lg">
						{tooltipData || token.price ? (
							<>
								$
								<Counter
									count={
										tooltipData?.price || token.price || 0
									}
								/>
							</>
						) : (
							<>-</>
						)}
					</p>
					<p
						className="mr-auto"
						style={{
							color
						}}
					>
						{token.name}
					</p>
				</div>
				<div
					className={cn(
						"ml-auto flex flex-col items-center",
						change === undefined
							? "opacity-60"
							: change > 0
								? "text-plug-green"
								: "text-red-500"
					)}
				>
					<p className="ml-auto flex w-max flex-row text-lg font-bold">
						<Counter count={change || 0} decimals={2} />%
					</p>
					<p className="ml-auto flex flex-row items-center">
						{tooltipData
							? formatTimestamp(Number(tooltipData.timestamp))
							: (header.title ?? "Past Week")}
					</p>
				</div>
			</div>

			<SocketTokenPriceChart
				enabled={isFrame}
				chain={token.implementations[0].chain}
				contract={token.implementations[0].contract}
				color={color}
				handleHeader={setHeader}
				handleTooltip={setTooltipData}
			/>

			<div className="flex flex-row gap-2 px-6 pt-4">
				<button
					className="flex w-full items-center justify-center gap-2 rounded-lg py-4 font-bold transition-all duration-200 ease-in-out hover:opacity-90"
					style={{
						backgroundColor: color ?? "",
						color: textColor
					}}
				>
					<Send size={14} className="opacity-60" />
					Send
				</button>
			</div>

			<div className="flex flex-col px-6 pb-2 pt-4 font-bold">
				<div className="flex flex-row items-center gap-4">
					<p className="mr-auto opacity-40">Balance</p>
					<div
						className="h-[2px] w-full"
						style={{ backgroundColor: color }}
					/>
					<p className="ml-auto opacity-40">Value</p>
				</div>

				<div className="mt-2 flex flex-row items-center justify-between gap-4">
					<p className="mr-auto flex flex-col items-center">
						<span
							className="mr-auto flex h-8 flex-row items-center"
							style={{ color: color }}
						>
							<TokenImage
								logo={token?.icon ?? ""}
								symbol={token.symbol}
								size="xs"
							/>
							<Counter
								className="ml-4 mr-2 w-max"
								count={token.balance}
							/>
							{token.symbol}
						</span>
					</p>
					{token.value && (
						<p className="ml-auto flex flex-col items-center text-center">
							<span className="mx-auto flex h-8 w-max items-center">
								$
								<Counter
									count={
										tooltipData
											? token.balance * tooltipData.price
											: token.value
									}
									decimals={2}
								/>
							</span>
						</p>
					)}
				</div>
			</div>

			<div className="flex flex-row items-center gap-4 px-6 font-bold">
				<p className="opacity-40">Distribution</p>
				<div
					className="h-[2px] w-full"
					style={{ backgroundColor: color }}
				/>
			</div>

			<div className="relative mt-2 flex w-full flex-col gap-2 px-6 pb-4">
				{token.implementations.map((implementation, index) => (
					<div
						key={index}
						className="flex flex-row items-center gap-4"
					>
						<Image
							src={getChainImage(implementation.chain)}
							alt={implementation.chain}
							className="h-4 w-4 rounded-full"
							width={24}
							height={24}
						/>

						<p className="mr-auto font-bold">
							{formatTitle(implementation.chain)}
						</p>

						<p className="flex flex-col font-bold opacity-60">
							<Counter
								count={isFrame ? implementation.balance : 0}
							/>
						</p>

						<p className="flex min-w-[72px] flex-row items-center text-right font-bold">
							<Counter
								count={isFrame ? implementation.percentage : 0}
							/>
							%
						</p>
					</div>
				))}
			</div>
		</Frame>
	)
}
