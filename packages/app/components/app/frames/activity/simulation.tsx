import { FC } from "react"

import { Bell, CircleDollarSign, Clock10, MessageCircleQuestionIcon, Share } from "lucide-react"

import { Frame } from "@/components/app/frames/base"
import { ActionPreview } from "@/components/app/plugs/actions/action-preview"
import { ActivityIcon } from "@/components/app/sockets/activity/activity-item"
import { Button } from "@/components/shared/buttons/button"
import { Counter } from "@/components/shared/utils/counter"
import { DateSince } from "@/components/shared/utils/date-since"
import { formatTitle } from "@/lib"
import { RouterOutputs } from "@/server/client"
import { useColumnStore } from "@/state/columns"

export const SimulationFrame: FC<{
	index: number
	activity: RouterOutputs["plugs"]["activity"]["get"][number] | undefined
	simulationId: string | undefined
}> = ({ index, activity, simulationId }) => {
	const { handle } = useColumnStore(index, `${activity?.id}-activity`)

	const simulation = activity?.simulations.find(sim => sim.id === simulationId)
	const actions = JSON.parse(activity?.actions ?? "[]")

	if (!activity || !simulation) return null

	return (
		<Frame
			index={index}
			icon={<ActivityIcon status={simulation.status} />}
			label={`Simulation ${formatTitle(simulation.status)}`}
			visible={simulation !== undefined}
			handleBack={() => handle.frame()}
			hasOverlay={true}
		>
			<ActionPreview
				index={index}
				item={activity.workflow.id}
				actions={actions}
				errors={simulation.errors ?? []}
			/>

			{simulation.error && (
				<p className="mx-auto mt-4 px-8 text-center text-sm font-bold text-plug-red">
					Warning: {simulation.error}
				</p>
			)}

			{simulation.status === "success" ? (
				<Button className="mt-4 flex w-full flex-row items-center justify-center gap-2 py-4" onClick={() => {}}>
					<Share size={18} className="opacity-60" />
					Share
				</Button>
			) : (
				<Button className="mt-4 flex w-full flex-row items-center justify-center gap-2 py-4" onClick={() => {}}>
					<MessageCircleQuestionIcon size={18} className="opacity-60" />
					Get Help
				</Button>
			)}

			<div className="mb-2 mt-4 flex flex-row items-center gap-4">
				<p className="font-bold opacity-40">Details</p>
				<div className="h-[2px] w-full bg-plug-green/10" />
			</div>

			<p className="flex flex-row items-center justify-between gap-4 font-bold">
				<Bell size={18} className="opacity-20" />
				<span className="opacity-40">Status</span>{" "}
				<span className="ml-auto">{formatTitle(simulation.status)}</span>
			</p>

			<p className="flex flex-row items-center justify-between gap-4 font-bold">
				<Clock10 size={18} className="opacity-20" />
				<span className="opacity-40">Simulated</span>{" "}
				<span className="ml-auto">
					<DateSince date={simulation.createdAt} />
				</span>
			</p>

			{simulation.gasEstimate && (
				<p className="flex flex-row items-center justify-between gap-4 font-bold">
					<CircleDollarSign size={18} className="opacity-20" />
					<span className="opacity-40">Fee</span>{" "}
					<span className="group ml-auto flex cursor-pointer flex-row items-center gap-4">
						<Counter count={simulation.gasEstimate} />
					</span>
				</p>
			)}
		</Frame>
	)
}
