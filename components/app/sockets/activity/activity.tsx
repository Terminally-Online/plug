import { FC, HTMLAttributes, useMemo } from "react"
import { useSession } from "next-auth/react"
import { cn } from "@/lib"
import { useColumns, useSocket } from "@/state"
import { api } from "@/server/client"
import { Callout } from "@/components"
import { ActivityItem } from "./activity-item"
import { ActivityFrame } from "../../frames/sockets/activity"

export const SocketActivity: FC<HTMLAttributes<HTMLDivElement> & { index?: number }> = ({
    index = -1,
    className,
    ...props
}) => {
    const { data: session } = useSession()
    const { isAnonymous } = useSocket()
    const { isExternal } = useColumns(index)

    const { data: queuedWorkflows, isLoading } = api.plug.action.getQueued.useQuery(
        undefined,
        {
            enabled: Boolean(session?.address) && (!isAnonymous || isExternal)
        }
    )

    const visibleActivities = useMemo(() => {
        if (!session || (isAnonymous && isExternal === false)) {
            return Array(10).fill(undefined)
        }

        if (isLoading) {
            return Array(10).fill(undefined)
        }

        return queuedWorkflows || []
    }, [queuedWorkflows, session, isAnonymous, isExternal, isLoading])

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

            {visibleActivities
                .filter(activity => Boolean(activity))
                .map((activity, activityIndex) => (
                    <ActivityFrame
                        key={activity?.id || activityIndex}
                        index={index}
                        activityIndex={activityIndex}
                        activity={{
                            name: activity?.text || '',
                            status: 'pending'
                        }}
                    />
                ))}
        </div>
    )
}
