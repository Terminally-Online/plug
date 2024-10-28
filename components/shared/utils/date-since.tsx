import { FC, HTMLAttributes, useEffect, useState } from "react"

import { Counter } from "@/components"
import { getTimeInterval } from "@/lib/functions/time"

export const DateSince: FC<{ date: Date; ago?: boolean } & HTMLAttributes<HTMLParagraphElement>> = ({
	date,
	ago = true,
	...props
}) => {
	const [timeSince, setTimeSince] = useState(getTimeInterval(0))

	useEffect(() => {
		const interval = setInterval(() => {
			const now = new Date()
			const seconds = Math.floor((now.getTime() - date.getTime()) / 1000)
			setTimeSince(getTimeInterval(seconds))
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
