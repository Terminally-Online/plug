import { FC } from "react"

import { AlertCircle, CheckCircle, XCircle } from "lucide-react"

import { Accordion, Counter, DateSince } from "@/components"
import { formatTitle } from "@/lib"
import { useFrame } from "@/state"

export const getStatusIcon = (status: string) => {
	switch (status) {
		case "success":
			return (
				<div className="relative h-10 w-10 min-w-10">
					<div className="absolute top-1/2 h-12 w-12 -translate-y-1/2 rounded-full bg-plug-green blur-2xl filter" />
					<CheckCircle
						className="absolute top-1/2 ml-auto h-4 w-6 -translate-y-1/2 text-center text-plug-green"
						size={16}
					/>
				</div>
			)
		case "error":
			return (
				<div className="relative h-10 w-10 min-w-10">
					<div className="absolute top-1/2 h-12 w-12 -translate-y-1/2 rounded-full bg-plug-red blur-2xl filter" />
					<XCircle
						className="absolute top-1/2 h-4 w-6 -translate-y-1/2 text-center text-plug-red"
						size={16}
					/>
				</div>
			)
		default:
			return (
				<div className="relative h-10 w-10 min-w-10">
					<div className="absolute top-1/2 h-12 w-12 -translate-y-1/2 rounded-full bg-yellow-400 blur-2xl filter" />
					<AlertCircle
						className="absolute top-1/2 h-4 w-6 -translate-y-1/2 text-center text-yellow-400"
						size={16}
					/>
				</div>
			)
	}
}

export const ActivityItem: FC<{
	index: number
	activityIndex: number
	activity?: {
		text: string
		color: string
		status: string
		time: string
	}
}> = ({ index, activityIndex, activity }) => {
	const { handleFrame } = useFrame({
		index,
		key: `${index}-${activityIndex}-activity`
	})

	const pastDate = new Date(Date.now() - 60000 * 0.2)

	return (
		<Accordion loading={activity === undefined} onExpand={() => handleFrame()}>
			{activity === undefined ? (
				<div className="invisible">
					<p>.</p>
					<p>.</p>
				</div>
			) : (
				<div className="flex w-full flex-row">
					{getStatusIcon(activity.status)}
					<div className="flex w-full flex-col overflow-hidden">
						<div className="flex flex-row items-center justify-between gap-2 font-bold">
							<p className="mr-2 truncate overflow-ellipsis whitespace-nowrap">{activity.text}</p>
							<div className="flex-shrink-0">
								<DateSince date={pastDate} />
							</div>
						</div>
						<div className="flex w-full flex-row items-center justify-between text-sm font-bold text-black text-opacity-40">
							<p>{formatTitle(activity.status)}</p>
							<p>
								<Counter count={pastDate.toLocaleDateString()} />
							</p>
						</div>
					</div>
				</div>
			)}
		</Accordion>
	)
}
