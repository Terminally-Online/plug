import { useSession } from "next-auth/react"
import { FC, HTMLAttributes, useMemo } from "react"

import { Callout } from "@/components"
import { useActivities } from "@/contexts"
import { cn } from "@/lib"
import { useSocket } from "@/state"

import { ActivityItem } from "./activity-item"

export const SocketActivity: FC<HTMLAttributes<HTMLDivElement> & { index?: number }> = ({
	index = -1,
	className,
	...props
}) => {
	const { data: session } = useSession()
	const { isAnonymous } = useSocket()
	const { activities, isLoading } = useActivities()

	const visibleActivities = useMemo(() => {
		if (!session || isAnonymous || isLoading) return Array(10).fill(undefined)

		return activities || []
	}, [activities, session, isAnonymous, isLoading])

	return (
		<div className={cn("flex h-full flex-col gap-2", className)} {...props}>
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
