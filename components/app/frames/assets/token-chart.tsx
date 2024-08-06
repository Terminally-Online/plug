import React, { FC, useCallback, useEffect, useMemo, useState } from "react"

import { Line, LineChart, ResponsiveContainer, XAxis, YAxis } from "recharts"

import { Button } from "@/components/shared"
import { cn } from "@/lib"

const periods = [
	{
		label: "1d",
		period: "1h",
		span: "24"
	},
	{
		label: "1w",
		period: "6h",
		span: "42"
	},
	{
		label: "1m",
		period: "12h",
		span: "60"
	},
	{
		label: "1y",
		period: "3d",
		span: "121"
	}
]

export const TokenPriceChart: FC<{
	enabled: boolean
	chain: string
	contract: string
	color: string
	handleTooltip: (data?: { timestamp: string; price: number }) => void
}> = ({ enabled, chain, contract, color, handleTooltip }) => {
	const [isLoading, setIsLoading] = useState(false)
	const [historicalPriceData, setHistoricalPriceData] = useState<
		Record<
			string,
			| Array<{ timestamp: string; price: number; error?: string }>
			| undefined
		>
	>({})
	const [period, setPeriod] = useState(periods[1])

	const key = `${chain}:${contract}`
	const url = `https://coins.llama.fi/chart/${key}?span=${period.span}&period=${period.period}`
	const priceData = historicalPriceData[period.label]

	const domain = useMemo(() => {
		const periodData = historicalPriceData[period.label]

		if (periodData === undefined || periodData[0]?.error) return [0, 1]

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
		if (priceData && index === priceData.length - 1) {
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

	const handleRetry = () => {
		setHistoricalPriceData(prevData => ({
			...prevData,
			[period.label]: undefined
		}))
		fetchHistoricalPriceData(true)
	}

	const fetchHistoricalPriceData = useCallback(
		async (isRetry = false) => {
			if (!enabled) return
			setIsLoading(true)

			const failure = {
				timestamp: new Date().toISOString(),
				price: 0,
				error: "Could not load price data."
			}

			try {
				if (isRetry) {
					await new Promise(resolve => setTimeout(resolve, 2000))
				}

				const response = await fetch(url)
				if (!response.ok) {
					throw new Error("Network response was not ok")
				}

				const data = await response.json()
				const prices = data.coins[key]?.prices

				if (prices === undefined) {
					throw new Error("No price data available")
				}

				setHistoricalPriceData(prevData => ({
					...prevData,
					[period.label]: prices
				}))
			} catch (error) {
				setHistoricalPriceData(prevData => ({
					...prevData,
					[period.label]: [failure]
				}))
			} finally {
				setIsLoading(false)
			}
		},
		[enabled, key, period, url]
	)

	useEffect(() => {
		fetchHistoricalPriceData()
	}, [key, period, url])

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

			{isLoading ? (
				<div className="flex min-h-[240px] flex-col items-center justify-center">
					<p className="font-bold opacity-40">
						Loading price data...
					</p>
				</div>
			) : priceData?.[0]?.error ? (
				<div className="flex min-h-[240px] flex-col items-center justify-center gap-2">
					<p className="font-bold opacity-40">
						{priceData?.[0]?.error}
					</p>

					<Button
						onClick={handleRetry}
						variant="secondary"
						sizing="md"
					>
						Retry
					</Button>
				</div>
			) : (
				<ResponsiveContainer
					minHeight={240}
					height="100%"
					width="100%"
					style={{ marginLeft: "-15%" }}
				>
					<LineChart
						data={priceData}
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
			)}

			<div className="mt-4 flex flex-row justify-center gap-2">
				{periods.map(p => (
					<Button
						key={p.label}
						variant="secondary"
						className={cn(
							"rounded-sm p-1 px-2",
							period.label === p.label && "active"
						)}
						onClick={() => setPeriod(p)}
					>
						<span
							className={cn(
								"text-black transition-all duration-200 ease-in-out",
								period.label === p.label
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
