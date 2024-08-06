import React, { FC, useEffect, useMemo, useState } from "react"

import { Line, LineChart, ResponsiveContainer, XAxis, YAxis } from "recharts"

import { Button } from "@/components/shared"
import { cn } from "@/lib"

const periods = [
	{
		label: "1h",
		value: "1h"
	},
	{
		label: "1d",
		value: "1d"
	},
	{
		label: "1w",
		value: "7d"
	},
	{
		label: "1m",
		value: "30d"
	},
	{
		label: "1y",
		value: "365d"
	}
]

export const TokenPriceChart: FC<{
	enabled: boolean
	chain: string
	contract: string
	color: string
	handleTooltip: (data?: { timestamp: string; price: number }) => void
}> = ({ enabled, chain, contract, color, handleTooltip }) => {
	const [historicalPriceData, setHistoricalPriceData] = useState<
		Record<string, Array<{ timestamp: string; price: number }>>
	>({})
	const [period, setPeriod] = useState<string>("1h")

	const key = `${chain}:${contract}`
	const url = `https://coins.llama.fi/chart/${key}?span=100&period=${period}`

	const domain = useMemo(() => {
		const periodData = historicalPriceData[period]

		if (!periodData || periodData.length === 0) return [0, 1]

		const min = Math.min(...periodData.map(price => price.price))
		const max = Math.max(...periodData.map(price => price.price))

		return [min * 0.98, max * 1.02]
	}, [period, historicalPriceData])

	const Dot: FC<{
		cx?: number
		cy?: number
		index?: number
	}> = ({ cx, cy, index }) => {
		if (index === historicalPriceData[period]?.length - 1) {
			return (
				<circle
					cx={cx}
					cy={cy}
					r={4}
					fill={color}
					className="blinking-dot"
				/>
			)
		}
		return null
	}

	const handleMouseMove = (e: any) => {
		if (e.activePayload && e.activePayload.length) {
			handleTooltip(e.activePayload[0].payload)
		}
	}

	const handleMouseLeave = () => {
		handleTooltip(undefined)
	}

	useEffect(() => {
		if (!enabled) return

		const fetchHistoricalPriceData = async () => {
			if (historicalPriceData[period]) return historicalPriceData[period]

			const response = await fetch(url)
			const data = await response.json()

			setHistoricalPriceData({
				[period]: data.coins[key].prices
			})
		}

		fetchHistoricalPriceData()
	}, [enabled, key, period, url])

	return (
		<div className="w-full pt-8">
			<style>
				{`
                    @keyframes blink {
                        0% { opacity: 0; }
                        50% { opacity: 1; }
                        100% { opacity: 0; }
                    }
                    .blinking-dot {
                        animation: blink 1.5s infinite;
                    }
                `}
			</style>

			{historicalPriceData[period] ? (
				<ResponsiveContainer
					minHeight={240}
					height="100%"
					width="100%"
					style={{ marginLeft: "-15%" }}
				>
					<LineChart
						data={historicalPriceData[period]}
						onMouseMove={handleMouseMove}
						onMouseLeave={handleMouseLeave}
					>
						<Line
							type="monotone"
							dataKey="price"
							stroke={color}
							strokeWidth={4}
							strokeLinecap="round"
							dot={<Dot />}
							isAnimationActive={false}
						/>
						<XAxis
							dataKey="timestamp"
							axisLine={false}
							tickLine={false}
							tick={false}
						/>
						<YAxis
							domain={domain}
							axisLine={false}
							tickLine={false}
							tick={false}
						/>
					</LineChart>
				</ResponsiveContainer>
			) : (
				<div className="flex min-h-[240px] flex-col items-center justify-center">
					<p className="font-bold opacity-40">
						Loading price data...
					</p>
				</div>
			)}

			<div className="mt-4 flex flex-row justify-center gap-2">
				{periods.map(p => (
					<Button
						key={p.label}
						variant="secondary"
						className={cn(
							"rounded-sm p-1 px-2",
							period === p.value && "active"
						)}
						onClick={() => setPeriod(p.value)}
					>
						<span
							className={cn(
								"text-black transition-all duration-200 ease-in-out",
								period === p.value
									? "text-opacity-100"
									: "text-opacity-40"
							)}
						>
							{p.label.toUpperCase()}
						</span>
					</Button>
				))}
			</div>
		</div>
	)
}
