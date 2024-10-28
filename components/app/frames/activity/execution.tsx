import Image from "next/image"
import { FC, useMemo } from "react"

import { Bell, Calendar, Pause, Play, TestTubeDiagonal, Trash, Waypoints } from "lucide-react"

import { ActionPreview, Button, Counter, Frame, TimeUntil } from "@/components"
import { useActivities } from "@/contexts"
import { chains, formatTitle } from "@/lib"
import { RouterOutputs } from "@/server/client"
import { useColumns } from "@/state"

export const ExecutionFrame: FC<{
	index: number
	icon: JSX.Element
	activity: RouterOutputs["plugs"]["activity"]["get"][number]
}> = ({ index, icon, activity }) => {
	const { isFrame } = useColumns(index, `${activity.id}-activity`)
	const { handle } = useActivities()

	const actions = useMemo(() => JSON.parse(activity.actions), [activity])

	return (
		<Frame index={index} icon={icon} label={activity.workflow.name} visible={isFrame} hasOverlay={true}>
			<div className="flex flex-col">
				<ActionPreview index={index} item={activity.workflow.id} actions={actions} />

				<div className="flex flex-row items-center gap-2">
					<Button
						variant="destructive"
						className="my-4 flex flex-row items-center justify-center gap-2 py-4"
						onClick={() => handle.delete({ id: activity.id })}
					>
						<Trash size={14} className="opacity-60" />
						Delete
					</Button>

					<Button
						variant="secondary"
						className="my-4 flex w-full flex-row items-center justify-center gap-2 py-4"
						onClick={() => handle.toggle({ id: activity.id })}
					>
						{activity.status === "pending" ? (
							<>
								<Pause size={14} className="opacity-60" />
								Pause
							</>
						) : (
							<>
								<Play size={14} className="opacity-60" />
								Resume
							</>
						)}
					</Button>
				</div>

				<div className="mb-2 flex flex-row items-center gap-4">
					<p className="font-bold opacity-40">Details</p>
					<div className="h-[2px] w-full bg-grayscale-100" />
				</div>

				<div className="flex flex-col">
					<p className="flex flex-row justify-between font-bold">
						<span className="flex w-full flex-row items-center gap-4">
							<Bell size={18} className="opacity-20" />
							<span className="opacity-40">Status</span>
						</span>{" "}
						{formatTitle(activity.status)}
					</p>
					<p className="flex w-full flex-row items-center gap-4 font-bold">
						<Waypoints size={18} className="opacity-20" />
						<span className="mr-auto opacity-40">Chain</span>
						<span className="flex flex-row items-center gap-2">
							<Image className="h-4 w-4" src={chains[1].logo} alt="ethereum" width={24} height={24} />
							Ethereum
						</span>
					</p>
					<p className="flex flex-row justify-between font-bold">
						<span className="flex w-full flex-row items-center gap-4">
							<Calendar size={18} className="opacity-20" />
							<span className="opacity-40">Scheduled At</span>
						</span>{" "}
						<Counter count={activity.createdAt.toLocaleDateString()} />
					</p>
					<p className="flex flex-row justify-between font-bold">
						<span className="flex w-full flex-row items-center gap-4">
							<Play size={18} className="opacity-20" />
							<span className="opacity-40">Start At</span>
						</span>{" "}
						<Counter count={activity.startAt.toLocaleDateString()} />
					</p>
					{activity.endAt && (
						<p className="flex flex-row justify-between font-bold">
							<span className="flex w-full flex-row items-center gap-4">
								<Pause size={18} className="opacity-20" />
								<span className="opacity-40">Stop At</span>
							</span>{" "}
							<Counter count={activity.endAt.toLocaleDateString()} />
						</p>
					)}
					<p className="flex flex-row justify-between font-bold">
						<span className="flex w-full flex-row items-center gap-4">
							<TestTubeDiagonal size={18} className="opacity-20" />
							<span className="opacity-40">Next Simulation</span>
						</span>{" "}
						<TimeUntil date={activity.nextSimulationAt} />
					</p>
				</div>
			</div>
		</Frame>
	)
}
