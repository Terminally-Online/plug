import { FC, HTMLAttributes, memo, useMemo } from "react"

import { useAtom } from "jotai"

import { ActivityItem } from "@/components/app/sockets/activity/activity-item"
import { Callout } from "@/components/app/utils/callout"
import { useActivities } from "@/contexts"
import { cn } from "@/lib"
import { useSocket } from "@/state/authentication"
import { columnByIndexAtom, COLUMNS } from "@/state/columns"
import { PLACEHOLDER_ACTIVITIES } from "@/lib/constants/placeholder/activity"

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
	const [column] = useAtom(columnByIndexAtom(index))

	const { isAnonymous } = useSocket()
	const { activities } = useActivities()

	const simulationId = useMemo(() => {
		const prefix = "-simulation"
		const isSimulation = column?.frame?.endsWith(prefix)

		if (!isSimulation || !column?.frame) return

		return column?.frame.split(prefix)[0]
	}, [column?.frame])

	const visibleActivities = !isAnonymous && activities?.length === 0 ? PLACEHOLDER_ACTIVITIES : activities

	return (
		<div className={cn("flex flex-col gap-2", className)} {...props}>
			<SocketActivityList activities={visibleActivities} index={index} simulationId={simulationId} />

			<Callout.Anonymous index={index} viewing="activity" isAbsolute={true} />
			<Callout.EmptyActivity index={index} isEmpty={!isAnonymous && activities?.length === 0} />
		</div>
	)
}
