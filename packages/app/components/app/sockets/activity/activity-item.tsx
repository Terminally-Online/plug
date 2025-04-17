import { FC } from "react"

import { AlertCircle, CheckCircle, Loader, Pause, Play, XCircle } from "lucide-react"

import { ExecutionFrame } from "@/components/app/frames/activity/execution/frame"
import { SimulationFrame } from "@/components/app/frames/activity/simulation/frame"
import { Accordion } from "@/components/shared/utils/accordion"
import { Counter } from "@/components/shared/utils/counter"
import { DateSince } from "@/components/shared/utils/date-since"
import { ChainId, formatTitle } from "@/lib"
import { RouterOutputs } from "@/server/client"
import { useColumnActions } from "@/state/columns"

import { ChainImage } from "../chains/chain.image"

const STATUS_CONFIG: Record<
	string,
	{ color: string; Icon: FC<{ className?: string; size?: number }> }
> = {
	upcoming: { color: "text-blue-400", Icon: Loader },
	completed: { color: "text-blue-400", Icon: CheckCircle },
	active: { color: "text-blue-400", Icon: Play },
	paused: { color: "text-plug-green", Icon: Pause },
	success: { color: "text-plug-green", Icon: CheckCircle },
	failure: { color: "text-plug-red", Icon: XCircle },
	default: { color: "text-yellow-400", Icon: AlertCircle },
}

export const ActivityIcon: FC<{ status: string }> = ({ status }) => {
	const { color, Icon } = STATUS_CONFIG[status] ?? STATUS_CONFIG.default
	const bgColor = color.replace("text-", "bg-")

	return (
		<div className="relative h-10 min-w-12">
			<div className={`absolute mt-8 h-48 w-16 rounded-full ${bgColor} blur-2xl filter`} />
			<Icon
				className={`-ml-1 absolute top-1/2 left-1/2 h-4 w-6 -translate-y-1/2 -translate-x-1/2 text-center ${color}`}
				size={16}
			/>
		</div>
	)
}

export const ActivityItem: FC<{
	index: number
	activity: RouterOutputs["plugs"]["activity"]["get"][number] | undefined
	simulationId: string | undefined
}> = ({ index, activity, simulationId }) => {
	const { frame } = useColumnActions(index, `${activity?.id}-activity`)

	return (
		<>
			<Accordion loading={activity === undefined} onExpand={() => frame()}>
				{activity === undefined ? (
					<div className="invisible">
						<p>.</p>
						<p>.</p>
					</div>
				) : (
					<div className="flex w-full flex-row items-center">
						<ActivityIcon status={activity.status} />

						<div className="relative flex w-full flex-col overflow-hidden">
							<div className="flex flex-row items-center justify-between gap-2 font-bold">
								<p className="mr-2 truncate overflow-ellipsis whitespace-nowrap">
									{activity.plug?.name || activity.inputs.map(input => formatTitle(input.action)).join(", ")}
								</p>
								<div className="flex-shrink-0">
									<DateSince date={new Date(activity.createdAt)} />
								</div>
							</div>
							<div className="flex w-full flex-row items-center justify-between gap-2 text-sm font-bold text-black text-opacity-40">
								<p className="flex flex-row items-center gap-2">
									<ChainImage chainId={activity.chainId as ChainId} size="xs" />
									<span className="truncate overflow-ellipsis whitespace-nowrap">
										{formatTitle(activity.status)}
									</span>
								</p>
								<div className="flex flex-row gap-2">
									<Counter count={new Date(activity.startAt).toLocaleDateString()} />
									{activity.endAt ? (
										<>
											<span className="opacity-60">→</span>
											<Counter count={new Date(activity.endAt).toLocaleDateString()} />
										</>
									) : activity.frequency !== 0 ? (
										<>
											<span className="opacity-60">→</span>∞
										</>
									) : (
										""
									)}
								</div>
							</div>
						</div>
					</div>
				)}
			</Accordion>

			<ExecutionFrame
				index={index}
				icon={<ActivityIcon status={activity?.status ?? "pending"} />}
				activity={activity!!}
			/>
			<SimulationFrame index={index} activity={activity} simulationId={simulationId} />
		</>
	)
}
