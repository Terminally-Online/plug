import { useMemo, useState } from "react"

import { SocketEarningsChartItem } from "."

const earnings = [
	[0.2, 0.1],
	[0.4, 0.6],
	[0.3, 0.4],
	[0.3, 0.1],
	[0.2, 0.5],
	[0.1, 0.7],
	[0.1, 0.4]
]

export const SocketEarningsChart = () => {
	const [active, setActive] = useState<number | undefined>(undefined)

	const [start, end] = useMemo(
		() => [
			new Date(Date.now() - 7 * 24 * 60 * 60 * 1000).toLocaleDateString(
				"en-US",
				{
					month: "numeric",
					day: "numeric"
				}
			),
			new Date().toLocaleDateString("en-US", {
				month: "numeric",
				day: "numeric"
			})
		],
		[]
	)

	return (
		<>
			<div className="flex flex-col gap-2">
				<div className="ml-auto mt-auto flex min-h-[132px] w-full flex-row">
					{earnings.map((day, index) => (
						<SocketEarningsChartItem
							key={index}
							padded={index !== 0}
							forks={day[1]}
							runs={day[0]}
							active={
								index === active ||
								(active === undefined &&
									index === earnings.length - 1)
							}
							onMouseEnter={() => setActive(index)}
							onMouseLeave={() => setActive(undefined)}
						/>
					))}
				</div>

				<div className="flex flex-row gap-2 text-sm font-bold tabular-nums opacity-40">
					<p>{start}</p>
					<p className="ml-auto">{end}</p>
				</div>
			</div>
		</>
	)
}
