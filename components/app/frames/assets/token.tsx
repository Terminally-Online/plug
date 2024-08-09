import { FC, useMemo, useState } from "react"

import Image from "next/image"

import { Counter, Frame, SocketTokenPriceChart } from "@/components"
import { useBalances, useFrame } from "@/contexts"
import { cn, formatTitle, getChainImage } from "@/lib"

import { TokenImage } from "../../sockets/tokens/token-image"

export const TokenFrame: FC<{ id: string; symbol: string }> = ({
	id,
	symbol
}) => {
	const { isFrame } = useFrame({ id, key: `token/${symbol}` })
	const { tokens } = useBalances()

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

	const change = useMemo(() => {
		if (header.change && !tooltipData) return header.change

		if (tooltipData) {
			const percentageChange =
				(tooltipData.price - tooltipData.start) / tooltipData.start

			return percentageChange * 100
		}

		return token?.chains[0].change
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
			icon={
				<div className="relative h-10 w-10">
					<TokenImage
						logo={token.logo}
						symbol={token.symbol}
						size="sm"
						handleColor={setColor}
					/>
				</div>
			}
			label=""
			visible={isFrame}
			hasOverlay={true}
			hasChildrenPadding={false}
		>
			<div className="flex flex-row px-6 text-lg font-bold">
				<div className="flex flex-col items-center font-bold">
					<p className="mr-auto flex w-max flex-row">
						$
						<Counter
							count={tooltipData?.price || token.chains[0].price}
						/>
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
					<p className="ml-auto flex w-max flex-row font-bold">
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
				chain={token.chains[0].chain}
				contract={token.chains[0].contract}
				color={color}
				handleHeader={setHeader}
				handleTooltip={setTooltipData}
			/>

			<div className="mt-4 flex flex-row items-center justify-between border-t-[1px] border-grayscale-100 px-6 py-4 font-bold">
				<p className="mr-auto flex flex-col items-center">
					<span className="mr-auto opacity-40">Balance</span>
					<span
						className="mr-auto flex h-8 flex-row items-center text-lg"
						style={{ color: color }}
					>
						<TokenImage
							logo={token.logo}
							symbol={token.symbol}
							size="xs"
						/>
						<Counter
							className="ml-4 mr-2 w-max"
							count={token.balance}
							decimals={2}
						/>
						{token.symbol}
					</span>
				</p>
				<p className="ml-auto flex flex-col items-center text-center">
					<span className="ml-auto opacity-40">Value</span>
					<span className="mx-auto flex h-8 w-max items-center text-lg">
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
			</div>

			<div className="relative flex w-full flex-col gap-2 border-t-[1px] border-grayscale-100 px-6 pb-8 pt-4 text-lg">
				{token.chains.map((chain, index) => (
					<div
						key={index}
						className="flex flex-row items-center gap-4"
					>
						<Image
							src={getChainImage(chain.chain)}
							alt={chain.chain}
							className="h-6 w-6 rounded-full"
							width={32}
							height={32}
						/>

						<p className="mr-auto font-bold">
							{formatTitle(chain.chain)}
						</p>

						<p className="flex flex-col font-bold opacity-60">
							<Counter count={isFrame ? chain.balance : 0} />
						</p>

						<p className="flex min-w-[72px] flex-row items-center text-right font-bold">
							<Counter count={isFrame ? chain.percentage : 0} />%
						</p>
					</div>
				))}
			</div>
		</Frame>
	)
}
