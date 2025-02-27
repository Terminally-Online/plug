import { FC, HTMLAttributes, useEffect, useRef, useState } from "react"

import { Counter } from "@/components/shared/utils/counter"
import { getTimeInterval } from "@/lib/functions/time"

const getUpdateInterval = (unit: string): number => {
	switch (unit) {
		case "s":
			return 1000
		case "m":
			return 60000
		case "h":
			return 300000
		case "d":
			return 3600000
		case "mo":
			return 3600000 * 6
		case "y":
			return 3600000 * 12
		default:
			return 1000
	}
}

const calculateTimeSince = (date: Date) => {
	const now = new Date()
	const seconds = Math.floor((now.getTime() - date.getTime()) / 1000)
	return getTimeInterval(seconds)
}

export const DateSince: FC<{ date: Date; ago?: boolean } & HTMLAttributes<HTMLParagraphElement>> = ({
	date,
	ago = true,
	...props
}) => {
	const [timeSince, setTimeSince] = useState(calculateTimeSince(date))
	const intervalRef = useRef<NodeJS.Timeout | null>(null)

	useEffect(() => {
		const updateDisplay = () => {
			const newTimeSince = calculateTimeSince(date)
			setTimeSince(newTimeSince)

			const nextInterval = getUpdateInterval(newTimeSince.unit)

			if (intervalRef.current) {
				clearInterval(intervalRef.current)
			}

			intervalRef.current = setTimeout(updateDisplay, nextInterval)
		}

		updateDisplay()

		return () => {
			if (intervalRef.current) {
				clearInterval(intervalRef.current)
			}
		}
	}, [date])

	return (
		<p className="flex w-max flex-row whitespace-nowrap" {...props}>
			<span>
				<Counter count={timeSince.interval} decimals={0} />
			</span>
			{timeSince.unit}
			{ago && <span className="ml-1">ago</span>}
		</p>
	)
}
