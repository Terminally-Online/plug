import { FC } from "react"

import { Bell, Calendar, Pause, Play, Sparkle, TestTubeDiagonal } from "lucide-react"

import { RouterOutputs } from "@/server/client"

import { ActionPreview, Counter, Frame } from "@/components"
import { useColumns } from "@/state"

export const ExecutionFrame: FC<{
	index: number
	icon: JSX.Element
	activity: RouterOutputs["plug"]["action"]["activity"][number]
}> = ({ index, icon, activity }) => {
	const { isFrame } = useColumns(index, `${activity.id}-activity`)

	return (
		<Frame index={index} icon={icon} label={activity.workflow.name} visible={isFrame} hasOverlay={true}>
			<ActionPreview index={index} item={activity.workflow.id} />

			<div className="mb-2 flex flex-row items-center gap-4">
				<p className="font-bold opacity-40">Details</p>
				<div className="h-[2px] w-full bg-grayscale-100" />
			</div>

			<p className="flex flex-row justify-between font-bold">
				<span className="flex w-full flex-row items-center gap-4">
					<Bell size={18} className="opacity-20" />
					<span className="opacity-40">Status</span>
				</span>{" "}
				Pending
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
					<span className="opacity-40">Next Simulation At</span>
				</span>{" "}
				<Counter count={activity.nextSimulationAt.toLocaleDateString()} />
			</p>
		</Frame>
	)
}
