import { useSession } from "next-auth/react"
import { FC, HTMLAttributes, useEffect, useState } from "react"

import { SearchIcon } from "lucide-react"

import {
	ActionsFrame,
	ActionView,
	AuthRequiredFrame,
	Button,
	ExecuteFrame,
	ManagePlugFrame,
	Search,
	ShareFrame
} from "@/components"
import { cn } from "@/lib"
import { COLUMNS, useColumnStore, usePlugData } from "@/state"

export const Plug: FC<HTMLAttributes<HTMLDivElement> & { index?: number; item?: string; from?: string }> = ({
	index = COLUMNS.MOBILE_INDEX,
	item,
	from,
	...props
}) => {
	const { data: session } = useSession()
	const { handle } = useColumnStore(index)
	const { plug } = usePlugData(item)

	const [hasOpenedActions, setHasOpenedActions] = useState(false)

	const own = plug !== undefined && session && session.address === plug.socketId

	useEffect(() => {
		if (!plug || plug.actions !== "[]" || hasOpenedActions) return

		handle.frame(`${item}-actions`)
		setHasOpenedActions(true)
	}, [item, handle, plug, hasOpenedActions])

	if (!plug || !session) return null

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
						handleOnClick={() => handle.frame(`${item}-actions`)}
					/>
				)}

				<div className="relative flex flex-row gap-2 px-4">
					<Button
						variant="secondary"
						className="w-max bg-white py-4"
						onClick={() => {
							handle.schedule() // NOTE: Clear the schedule when we have a one-off run use.
							handle.frame("run")
						}}
					>
						Run
					</Button>

					<Button className="w-full py-4" onClick={() => handle.frame("schedule")}>
						Schedule
					</Button>
				</div>
			</div>

			{item && (
				<>
					<AuthRequiredFrame index={index} />
					<ExecuteFrame index={index} item={item} />
					<ManagePlugFrame index={index} item={item} from={from} />
					<ActionsFrame index={index} item={item} />
					<ShareFrame index={index} item={item} />
				</>
			)}
		</div>
	)
}
