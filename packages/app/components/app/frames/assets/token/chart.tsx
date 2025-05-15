import { FC, useCallback, useEffect, useState } from "react"

import { Line, LineChart, ResponsiveContainer, Tooltip, XAxis, YAxis } from "recharts"

import { Button } from "@/components/shared/buttons/button"
import { cn } from "@/lib"
import { api, RouterOutputs } from "@/server/client"

import { ActiveDot, Dot } from "./dot"

type Token = NonNullable<RouterOutputs["service"]["zerion"]["wallet"]["positions"]["data"]>[number] | undefined

const periods = [
	{
		label: "1d",
		period: "day",
		title: "Today"
	},
	{
		label: "1w",
		period: "week",
		title: "Past Week"
	},
	{
		label: "1m",
		period: "month",
		title: "Past Month"
	},
	{
		label: "1y",
		period: "year",
		title: "Past Year"
	}
] as const

export const SocketTokenPriceChart: FC<{
	enabled: boolean
	token: Token
	color: string
	handleHeader?: (data: { title: string; change?: number }) => void
	handleTooltip?: (data?: { timestamp: string; price: number; start: number }) => void
}> = ({ enabled, token, color, handleHeader = () => {}, handleTooltip = () => {} }) => {
	const [period, setPeriod] = useState<(typeof periods)[number]>(periods[0])

	const { data } = api.service.zerion.fungibles.chart.useQuery(
		{
			path: {
				fungibleId: token?.relationships.fungible.data.id ?? "",
				period: period.period
			},
			query: {
				currency: "usd"
			}
		},
		{ enabled: enabled && !!token, retry: false, placeholderData: prev => prev }
	)

	const price = data?.data
	const domain = [(price?.attributes.stats.min ?? 0) * 0.95, (price?.attributes.stats.max ?? 0) * 1.05]

	const handleMouseMove = useCallback(
		(e: any) => {
			if (!price || !e.activePayload || !e.activePayload.length) return

			handleTooltip({
				...e.activePayload[0].payload,
				start: price.attributes.stats.first
			})
		},
		[price, handleTooltip]
	)

	useEffect(() => {
		if (!price) return

		const start = price.attributes.stats.first ?? 0
		const end = price.attributes.stats.last ?? 0
		const change = ((end - start) / start) * 100

		handleHeader({ title: period.title, change })
	}, [period, price, handleHeader])

	return (
		<div className="w-full overflow-x-hidden pt-8">
			<ResponsiveContainer minHeight={140} height="100%" width="100%">
				<LineChart
					onMouseMove={handleMouseMove}
					onMouseLeave={() => handleTooltip(undefined)}
					margin={{ left: 0, right: 60, top: 0, bottom: 0 }}
				>
					<Line
						data={price?.attributes.points.map(([timestamp, price]) => ({
							timestamp,
							price
						}))}
						type="monotone"
						dataKey="price"
						stroke={color}
						strokeWidth={4}
						strokeLinecap="round"
						dot={<Dot length={price?.attributes.points.length ?? 0} color={color} />}
						activeDot={<ActiveDot color={color} />}
						isAnimationActive={false}
					/>

					<XAxis dataKey="timestamp" hide />
					<YAxis dataKey="price" domain={domain} hide />
					<Tooltip content={<></>} cursor={<></>} />
				</LineChart>
			</ResponsiveContainer>

			<div className="mt-4 flex flex-row justify-center gap-2">
				{periods.map(p => (
					<Button
						key={p.label}
						variant="secondary"
						className={cn(
							"rounded-sm p-1 px-2 text-sm",
							period.label === p.label && "active",
							period.label !== p.label && "opacity-40 hover:opacity-100"
						)}
						onClick={() => setPeriod(p)}
					>
						{p.label.toUpperCase()}
					</Button>
				))}
			</div>
		</div>
	)
}
