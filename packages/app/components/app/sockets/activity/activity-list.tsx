import { FC, HTMLAttributes, memo, useMemo } from "react"

import { ActivityItem } from "@/components/app/sockets/activity/activity-item"
import { Callout } from "@/components/app/utils/callout"
import { useActivities } from "@/contexts"
import { cn } from "@/lib"
import { useSocket } from "@/state/authentication"
import { COLUMNS, useColumnStore } from "@/state/columns"

const SocketActivityList: FC<{
	activities: any[] | undefined
	index: number
	simulationId?: string
}> = memo(({ activities, index, simulationId }) => {
	const visibleActivities = useMemo(() => {
		if (activities === undefined || activities.length === 0) return Array(10).fill(undefined)
		return activities
	}, [activities])

	return (
		<div className="flex flex-col gap-2">
			{visibleActivities.map((activity, activityIndex) => (
				<ActivityItem
					key={activity?.id || activityIndex}
					index={index}
					activity={activity}
					simulationId={simulationId}
				/>
			))}
		</div>
	)
})
SocketActivityList.displayName = "SocketActivityList"

export const SocketActivity: FC<HTMLAttributes<HTMLDivElement> & { index?: number }> = ({
	index = COLUMNS.MOBILE_INDEX,
	className,
	...props
}) => {
	const { column } = useColumnStore(index)
	const { isAnonymous } = useSocket()
	const { activities } = useActivities()

	const simulationId = useMemo(() => {
		const prefix = "-simulation"
		const isSimulation = column?.frame?.endsWith(prefix)

		if (!isSimulation || !column?.frame) return

		return column?.frame.split(prefix)[0]
	}, [column?.frame])

	return (
		<div className={cn("flex flex-col gap-2", className)} {...props}>
			<Callout.Anonymous index={index} viewing="activity" isAbsolute={true} />
			<Callout.EmptyActivity index={index} isEmpty={!isAnonymous && activities?.length === 0} />

			<SocketActivityList activities={activities} index={index} simulationId={simulationId} />
		</div>
	)
}
