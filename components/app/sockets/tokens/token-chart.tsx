import { FC, useCallback, useEffect, useMemo, useState } from "react"

import {
	Line,
	LineChart,
	ResponsiveContainer,
	Tooltip,
	XAxis,
	YAxis
} from "recharts"

import { Button } from "@/components/shared"
import { cn } from "@/lib"

const periods = [
	{
		label: "1d",
		period: "30m",
		span: "48",
		title: "Today"
	},
	{
		label: "1w",
		period: "6h",
		span: "42",
		title: "Past Week"
	},
	{
		label: "1m",
		period: "12h",
		span: "60",
		title: "Past Month"
	},
	{
		label: "1y",
		period: "3d",
		span: "121",
		title: "Past Year"
	}
]

export const SocketTokenPriceChart: FC<{
	enabled: boolean
	chain: string
	contract: string
	color: string
	handleHeader: (data: { title: string; change?: number }) => void
	handleTooltip: (data?: {
		timestamp: string
		price: number
		start: number
	}) => void
}> = ({ enabled, chain, contract, color, handleHeader, handleTooltip }) => {
	const [isLoading, setIsLoading] = useState(false)
	const [historicalPriceData, setHistoricalPriceData] = useState<
		Record<
			string,
			| Array<{ timestamp: string; price: number; error?: string }>
			| undefined
		>
	>({})
	const [period, setPeriod] = useState(periods[0])

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
				<g>
					<circle cx={cx} cy={cy} r={12} fill={color}>
						<animate
							attributeName="r"
							values="8;5;8"
							dur="1s"
							repeatCount="indefinite"
							keyTimes="0;0.5;1"
							keySplines="0.1 0.8 0.2 1; 0.1 0.8 0.2 1"
							calcMode="spline"
							begin="0s;animate2.end+1s"
						/>
						<animate
							id="animate2"
							attributeName="opacity"
							values="0;0.5;0"
							dur="1s"
							repeatCount="indefinite"
							keyTimes="0;0.5;1"
							keySplines="0.4 0 0.6 1; 0.4 0 0.6 1"
							calcMode="spline"
						/>
					</circle>
					<circle cx={cx} cy={cy} r={5} fill={color} />
				</g>
			)
		}

		return null
	}

	const ActiveDot: FC<{
		cx?: number
		cy?: number
	}> = ({ cx, cy }) => <circle cx={cx} cy={cy} r={5} fill={color} />

	const handleMouseMove = (e: any) => {
		if (e.activePayload && e.activePayload.length) {
			handleTooltip({
				...e.activePayload[0].payload,
				start: priceData && priceData[0].price
			})
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

			const key = `${chain}:${contract}`
			const url = `https://coins.llama.fi/chart/${key}?span=${period.span}&period=${period.period}`

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
		[enabled, period, chain, contract]
	)

	useEffect(() => {
		fetchHistoricalPriceData()
	}, [period, fetchHistoricalPriceData])

	useEffect(() => {
		const start = priceData?.[0]?.price
		const end = priceData?.[priceData.length - 1]?.price
		const change = start && end ? ((end - start) / start) * 100 : undefined

		handleHeader({ title: period.title, change })
	}, [period, historicalPriceData, handleHeader, priceData])

	return (
		<div className="w-full pt-8">
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
							activeDot={<ActiveDot />}
							isAnimationActive={false}
						/>
						<XAxis
							dataKey="timestamp"
							axisLine={false}
							tickLine={false}
							tick={false}
							padding={{ right: 20 }}
						/>
						<YAxis
							domain={domain}
							axisLine={false}
							tickLine={false}
							tick={false}
						/>
						<Tooltip content={<></>} cursor={<></>} />
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
