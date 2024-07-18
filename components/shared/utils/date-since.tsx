import { FC, PropsWithChildren, useEffect, useState } from "react"

import { Counter } from "@/components"

const getTimeSince = (date: Date) => {
	const now = new Date()
	const seconds = Math.floor((now.getTime() - date.getTime()) / 1000)

	let interval = Math.floor(seconds / 31536000)
	if (interval >= 1) {
		return { interval, unit: "y" }
	}

	interval = Math.floor(seconds / 2592000)
	if (interval >= 1) {
		return { interval, unit: "mo" }
	}

	interval = Math.floor(seconds / 86400)
	if (interval >= 1) {
		return { interval, unit: "d" }
	}

	interval = Math.floor(seconds / 3600)
	if (interval >= 1) {
		return { interval, unit: "h" }
	}

	interval = Math.floor(seconds / 60)
	if (interval >= 1) {
		return { interval, unit: "m" }
	}

	return { interval: Math.floor(seconds), unit: "s" }
}

export const DateSince: FC<
	{ date: Date } & PropsWithChildren & React.HTMLProps<HTMLSpanElement>
> = ({ date, ...props }) => {
	const [timeSince, setTimeSince] = useState(getTimeSince(date))

	useEffect(() => {
		const interval = setInterval(() => {
			setTimeSince(getTimeSince(date))
		}, 1000)

		return () => clearInterval(interval)
	}, [date])

	return (
		<span className="flex flex-row" {...props}>
			<span className="ml-auto">
				<Counter count={timeSince.interval} decimals={0} />
			</span>
			{timeSince.unit} ago
		</span>
	)
}
