import { FC, HTMLAttributes, useEffect, useState } from "react"

import { Counter } from "@/components/shared/utils/counter"
import { getTimeInterval } from "@/lib/functions/time"

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

	useEffect(() => {
		const interval = setInterval(() => {
			setTimeSince(calculateTimeSince(date))
		}, 1000)

		return () => clearInterval(interval)
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
