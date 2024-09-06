import { FC, HTMLAttributes } from "react"

import { useSession } from "next-auth/react"

import { SearchIcon } from "lucide-react"

import {
	ActionsFrame,
	ActionView,
	Button,
	Callout,
	ExecuteFrame,
	ManagePlugFrame,
	Search,
	ShareFrame
} from "@/components"
import { useFrame, usePlugs, useSockets } from "@/contexts"
import { cn } from "@/lib"

export const Plug: FC<HTMLAttributes<HTMLDivElement> & { id: string; item: string | null }> = ({
	id,
	item,
	...props
}) => {
	const { data: session } = useSession()
	const { socket } = useSockets()
	const { handleFrame } = useFrame({ id: id })
	const { plug } = usePlugs(item!)

	const own = plug !== undefined && session && session.address === plug.userAddress

	const page = socket?.columns.find(column => column.id === id)

	if (!plug || !page) return null

	return (
		<div {...props}>
			<ActionView id={id} />

			<div className="absolute bottom-0 left-0 z-[2] mb-4 flex w-full flex-col gap-2 overflow-y-visible">
				<div className="pointer-events-none absolute bottom-[100px] left-0 right-0 top-0 z-[-1] bg-gradient-to-t from-white to-white/0" />
				<div
					className={cn(
						"absolute -bottom-4 left-0 right-0 z-[-1] h-[100px] bg-white",
						page.index !== -1 && "rounded-b-lg"
					)}
				/>

				{own && (
					<Search
						className="px-4 pt-16"
						icon={<SearchIcon size={14} className="opacity-60" />}
						placeholder="Search protocols and actions"
						handleOnClick={() => handleFrame("actions")}
					/>
				)}

				<div className="relative flex flex-row gap-2 px-4">
					<Button variant="secondary" className="w-max bg-white" onClick={() => handleFrame("socket-run")}>
						Run
					</Button>

					<Button className="w-full" onClick={() => handleFrame("socket-schedule")}>
						Schedule
					</Button>
				</div>
			</div>

			<ExecuteFrame id={id} />
			<ManagePlugFrame id={id} />
			<ActionsFrame id={id} />
			<ShareFrame id={id} />
		</div>
	)
}
