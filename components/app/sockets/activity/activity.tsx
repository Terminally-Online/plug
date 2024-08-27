import { FC, HTMLAttributes } from "react"

import { ActivityList } from "@/components"
import { useSockets } from "@/contexts"
import { cn } from "@/lib"

export const SocketActivity: FC<HTMLAttributes<HTMLDivElement> & { id: string }> = ({ id, className, ...props }) => {
	const { anonymous } = useSockets()

	return (
		<div className={cn("flex h-full flex-col gap-2", className)} {...props}>
			{anonymous && (
				<div className="flex h-full flex-col items-center justify-center text-center font-bold">
					<p>You are anonymous.</p>
					<p className="max-w-[320px] opacity-40">
						To view the collectibles you are holding you must authenticate a wallet.
					</p>
				</div>
			)}

			{anonymous === false && <ActivityList id={id} />}
		</div>
	)
}
