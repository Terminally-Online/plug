import { FC, useCallback, useEffect, useMemo, useState } from "react"

import axios from "axios"
import { Line, LineChart, ResponsiveContainer, Tooltip, XAxis, YAxis } from "recharts"

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

const Dot: FC<{
	length: number
	color: string
	cx?: number
	cy?: number
	index?: number
}> = ({ length, color, cx, cy, index }) => {
	if (index === length - 1) {
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
	color: string
	cx?: number
	cy?: number
}> = ({ color, cx, cy }) => <circle cx={cx} cy={cy} r={5} fill={color} />

export const SocketTokenPriceChart: FC<{
	enabled: boolean
	keys: Array<string>
	colors: Record<string, string>
	handleHeader?: (data: { title: string; change?: Array<number> }) => void
	handleTooltip?: (data?: { timestamp: string; price: number; start: Array<number> }) => void
}> = ({ enabled, keys, colors, handleHeader = () => {}, handleTooltip = () => {} }) => {
	const [isLoading, setIsLoading] = useState(false)
	const [queriedKeys, setQueriedKeys] = useState("")
	const [historicalPriceData, setHistoricalPriceData] = useState<
		Record<
			string,
			Record<
				string,
				{
					prices: Array<{
						timestamp: string
						price: number
						error?: string
					}>
				}
			>
		>
	>({})
	const [period, setPeriod] = useState(periods[0])

	const priceData = historicalPriceData[period.label]

	const domain = useMemo(() => {
		const periodData = historicalPriceData[period.label]

		if (!periodData || Object.keys(periodData).length === 0) return [0, 1]

		const allPrices = Object.values(periodData).flatMap(token => token.prices.map(price => price.price))

		if (allPrices.length === 0) return [0, 1]

		const min = Math.min(...allPrices)
		const max = Math.max(...allPrices)

		return [min * 0.98, max * 1.02]
	}, [period, historicalPriceData])

	const handleMouseMove = useCallback(
		(e: any) => {
			const start = priceData && Object.values(priceData).map(token => token.prices[0].price)

			if (e.activePayload && e.activePayload.length) {
				handleTooltip({
					...e.activePayload[0].payload,
					start
				})
			}
		},
		[priceData, handleTooltip]
	)

	useEffect(() => {
		const fetchHistoricalPriceData = async () => {
			if (!enabled || keys.join(",") === queriedKeys) return

			setIsLoading(true)
			setQueriedKeys(keys.join(","))

			const failure = {
				timestamp: new Date().toISOString(),
				price: 0,
				error: "Could not load price data."
			}

			await Promise.allSettled(
				periods.map(async period => {
					const url = `https://coins.llama.fi/chart/${keys.join(",")}?span=${period.span}&period=${period.period}&searchWidth=1200`

					try {
						const response = await axios.get(url)

						if (response.status !== 200) throw new Error("Network response was not ok")

						const prices = response.data?.coins ?? {}

						setHistoricalPriceData(prevData => ({
							...prevData,
							[period.label]: prices
						}))
					} catch (error) {
						console.error(error)
						// setHistoricalPriceData(prevData => ({
						// 	...prevData,
						// 	[period.label]: [failure]
						// }))
					}
				})
			)

			setIsLoading(false)
		}

		fetchHistoricalPriceData()
	}, [enabled, keys, historicalPriceData])

	useEffect(() => {
		if (!priceData) return

		const change = Object.values(priceData).map(token => {
			const start = token.prices[0].price
			const end = token.prices[token.prices.length - 1].price
			return ((end - start) / start) * 100
		})

		handleHeader({ title: period.title, change })
	}, [period, historicalPriceData, handleHeader, priceData])

	return (
		<div className="w-full overflow-x-hidden pt-8">
			{isLoading ? (
				<div className="flex min-h-[240px] flex-col items-center justify-center">
					<p className="font-bold opacity-40">Loading price data...</p>
				</div>
			) : (
				// : priceData?.[0]?.error ? (
				// <div className="flex min-h-[240px] flex-col items-center justify-center gap-2">
				// 	<p className="font-bold opacity-40">
				// 		{priceData?.[0]?.error}
				// 	</p>
				// </div>
				// )
				<>
					<ResponsiveContainer minHeight={240} height="100%" width="100%" style={{ marginLeft: "-16%" }}>
						<LineChart onMouseMove={handleMouseMove} onMouseLeave={() => handleTooltip(undefined)}>
							{priceData &&
								keys.map(key => {
									if (!priceData[key]) return null

									const data = priceData[key]?.prices
									const color = colors[key]

									return (
										<Line
											data={data}
											type="monotone"
											dataKey="price"
											stroke={color}
											strokeWidth={4}
											strokeLinecap="round"
											name={key}
											key={key}
											dot={<Dot length={data ? data.length : 0} color={color} />}
											activeDot={<ActiveDot color={color} />}
											isAnimationActive={false}
										/>
									)
								})}

							<XAxis dataKey="timestamp" axisLine={false} tickLine={false} tick={false} padding={{ right: 20 }} />
							<YAxis dataKey="price" domain={domain} axisLine={false} tickLine={false} tick={false} />
							<Tooltip content={<></>} cursor={<></>} />
						</LineChart>
					</ResponsiveContainer>
				</>
			)}

			<div className="mt-4 flex flex-row justify-center gap-2">
				{periods.map(p => (
					<Button
						key={p.label}
						variant="secondary"
						className={cn("rounded-sm p-1 px-2", period.label === p.label && "active")}
						onClick={() => setPeriod(p)}
					>
						<span
							className={cn(
								"text-black transition-all duration-200 ease-in-out",
								period.label === p.label ? "text-opacity-100" : "text-opacity-40"
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
