import { FC, HTMLAttributes, useMemo } from "react"

import { ActivityItem, Callout } from "@/components"
import { useActivities } from "@/contexts"
import { cn } from "@/lib"
import { COLUMNS, useSocket } from "@/state"

export const SocketActivity: FC<HTMLAttributes<HTMLDivElement> & { index?: number }> = ({
	index = COLUMNS.MOBILE_INDEX,
	className,
	...props
}) => {
	const { isAnonymous } = useSocket()
	const { activities, isLoading } = useActivities()

	const visibleActivities = useMemo(() => {
		if (activities === undefined || isLoading || activities.length === 0) return Array(10).fill(undefined)
		return activities
	}, [activities, isLoading])

	return (
		<div className={cn("flex flex-col gap-2", className)} {...props}>
			<Callout.Anonymous index={index} viewing="activity" isAbsolute={true} />
			<Callout.EmptyActivity index={index} isEmpty={!isAnonymous && activities?.length === 0} />

			<div className="flex flex-col gap-2">
				{visibleActivities.map((activity, activityIndex) => (
					<ActivityItem key={activity?.id || activityIndex} index={index} activity={activity} />
				))}
			</div>
		</div>
	)
}
