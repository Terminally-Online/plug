import { useSession } from "next-auth/react"
import { FC, HTMLAttributes, useMemo } from "react"

import { Callout } from "@/components"
import { useActivities } from "@/contexts"
import { cn } from "@/lib"
import { useColumns, useSocket } from "@/state"

import { ActivityFrame } from "../../frames/sockets/activity"
import { ActivityItem } from "./activity-item"

export const SocketActivity: FC<HTMLAttributes<HTMLDivElement> & { index?: number }> = ({
	index = -1,
	className,
	...props
}) => {
	const { data: session } = useSession()
	const { isAnonymous } = useSocket()
	const { isExternal } = useColumns(index)
	const { activities, isLoading } = useActivities()

	const visibleActivities = useMemo(() => {
		if (!session || (isAnonymous && isExternal === false) || isLoading) return Array(10).fill(undefined)

		return activities || []
	}, [activities, session, isAnonymous, isExternal, isLoading])

	return (
		<div className={cn("flex h-full flex-col gap-2", className)} {...props}>
			<div className="flex flex-col gap-2">
				{visibleActivities.map((activity, activityIndex) => (
					<ActivityItem
						key={activity?.id || activityIndex}
						id={`${index}-${activityIndex}-activity`}
						index={activityIndex}
						activity={activity}
					/>
				))}
			</div>

			<Callout.Anonymous index={index} viewing="activity" isAbsolute={true} />
		</div>
	)
}
