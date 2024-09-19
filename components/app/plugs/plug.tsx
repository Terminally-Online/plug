import { useSession } from "next-auth/react"
import { FC, HTMLAttributes } from "react"

import { SearchIcon } from "lucide-react"

import { ActionsFrame, ActionView, Button, ExecuteFrame, ManagePlugFrame, Search, ShareFrame } from "@/components"
import { usePlugs } from "@/contexts"
import { cn } from "@/lib"
import { useFrame } from "@/state"

export const Plug: FC<HTMLAttributes<HTMLDivElement> & { index?: number; item?: string; from?: string }> = ({
	index = -1,
	item,
	from,
	...props
}) => {
	const { data: session } = useSession()
	const { handleFrame } = useFrame({ index })
	const { plug } = usePlugs(item)

	const own = plug !== undefined && session && session.address === plug.socketId

	if (!plug) return null

	return (
		<div {...props}>
			<ActionView index={index} />
			<div className="absolute bottom-0 left-0 z-[2] mb-4 flex w-full flex-col gap-2 overflow-y-visible">
				<div className="pointer-events-none absolute bottom-[100px] left-0 right-0 top-0 z-[-1] bg-gradient-to-t from-white to-white/0" />
				<div
					className={cn(
						"absolute -bottom-4 left-0 right-0 z-[-1] h-[100px] bg-white",
						index !== -1 && "rounded-b-lg"
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
			
			{item && (
				<>
					<ExecuteFrame index={index} item={item} />
					<ManagePlugFrame index={index} item={item} from={from} />
					<ActionsFrame index={index} item={item} />
					<ShareFrame index={index} item={item} />
				</>
			)}
		</div>
	)
}
