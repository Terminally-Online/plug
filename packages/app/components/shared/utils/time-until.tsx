import { FC, HTMLAttributes, useEffect, useState } from "react"

import { Counter } from "@/components"
import { getTimeInterval } from "@/lib"

export const TimeUntil: FC<{ date: Date } & HTMLAttributes<HTMLParagraphElement>> = ({ date, ...props }) => {
	const [timeUntil, setTimeUntil] = useState(getTimeInterval(0))

	useEffect(() => {
		const interval = setInterval(() => {
			const now = new Date()
			const seconds = Math.floor((date.getTime() - now.getTime()) / 1000)
			setTimeUntil(getTimeInterval(Math.max(0, seconds)))
		}, 1000)

		return () => clearInterval(interval)
	}, [date])

	return (
		<p className="flex w-max flex-row whitespace-nowrap" {...props}>
			{timeUntil.interval > 0 ? (
				<>
					<span>
						<Counter count={timeUntil.interval} decimals={0} />
					</span>
					{timeUntil.unit}
				</>
			) : (
				"Now"
			)}
		</p>
	)
}
