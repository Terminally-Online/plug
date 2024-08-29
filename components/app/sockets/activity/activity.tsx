import { FC, HTMLAttributes } from "react"

import { ActivityList, Callout } from "@/components"
import { useSockets } from "@/contexts"
import { cn } from "@/lib"

export const SocketActivity: FC<HTMLAttributes<HTMLDivElement> & { id: string }> = ({ id, className, ...props }) => {
	const { isAnonymous: anonymous } = useSockets()

	return (
		<div className={cn("flex h-full flex-col gap-2", className)} {...props}>
			<Callout.Anonymous viewing="activity" />

			{anonymous === false && <ActivityList id={id} />}
		</div>
	)
}
