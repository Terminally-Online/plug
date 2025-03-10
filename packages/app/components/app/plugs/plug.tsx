import { useSession } from "next-auth/react"
import { FC, HTMLAttributes, useEffect, useState } from "react"

import { SearchIcon } from "lucide-react"

import { AuthRequiredFrame } from "@/components/app/frames/misc/auth-required"
import { ActionsFrame } from "@/components/app/frames/plugs/[id]/actions"
import { ExecuteFrame } from "@/components/app/frames/plugs/[id]/execute"
import { ShareFrame } from "@/components/app/frames/plugs/[id]/share"
import { Search } from "@/components/app/inputs/search"
import { ActionView } from "@/components/app/plugs/actions/action-view"
import { Button } from "@/components/shared/buttons/button"
import { cn } from "@/lib"
import { columnByIndexAtom, COLUMNS, useColumnActions } from "@/state/columns"
import { plugByIdAtom } from "@/state/plugs"
import { useAtom } from "jotai"

export const Plug: FC<HTMLAttributes<HTMLDivElement> & { index?: number; item: string; from?: string }> = ({
	index = COLUMNS.MOBILE_INDEX,
	item,
	from,
	...props
}) => {
	const { data: session } = useSession()

	const [column] = useAtom(columnByIndexAtom(index))
	const { frame } = useColumnActions(index)
	const [plug] = useAtom(plugByIdAtom(item))

	const [hasOpenedActions, setHasOpenedActions] = useState(false)

	const own = plug !== undefined && session && session.address === plug.socketId

	useEffect(() => {
		if (!plug || plug.actions.length !== 0 || hasOpenedActions) return

		frame(`${item}-actions`)
		setHasOpenedActions(true)
	}, [item, plug, hasOpenedActions, frame])

	if (!column || !plug) return null

	return (
		<div {...props}>
			<ActionView index={index} />

			<div className="absolute bottom-0 left-0 z-[2] mb-4 flex w-full flex-col gap-2 overflow-y-visible">
				<div className="pointer-events-none absolute bottom-[120px] left-0 right-0 top-0 z-[-1] bg-gradient-to-t from-white to-white/0" />
				<div
					className={cn(
						"absolute -bottom-4 left-0 right-0 z-[-1] h-[140px] bg-white",
						index !== COLUMNS.MOBILE_INDEX && "rounded-b-lg"
					)}
				/>

				{own && (
					<Search
						className="px-4 pt-16"
						icon={<SearchIcon size={14} className="opacity-60" />}
						placeholder="Search protocols and actions"
						handleOnClick={() => frame(`${item}-actions`)}
					/>
				)}

				<div className="relative flex flex-row gap-2 px-4">
					<Button
						variant="secondary"
						className="w-max bg-white py-4"
						onClick={() => {
							// schedule() // NOTE: Clear the schedule when we have a one-off run use.
							frame("run")
						}}
					>
						Run
					</Button>

					<Button className="w-full py-4" onClick={() => frame("schedule")}>
						Schedule
					</Button>
				</div>
			</div>

			{item && (
				<>
					<ActionsFrame index={index} item={item} />
					<AuthRequiredFrame index={index} />
					<ExecuteFrame index={index} item={item} />
					<ShareFrame index={index} item={item} />
				</>
			)}
		</div>
	)
}