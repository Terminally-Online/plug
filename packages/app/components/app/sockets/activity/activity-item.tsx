import { FC, useMemo } from "react"

import {
	AlertCircle,
	Bell,
	CheckCircle,
	CircleDollarSign,
	Clock10,
	Loader,
	MessageCircleQuestionIcon,
	Pause,
	Share,
	XCircle
} from "lucide-react"

import { ExecutionFrame } from "@/components/app/frames/activity/execution"
import { Frame } from "@/components/app/frames/base"
import { ActionPreview } from "@/components/app/plugs/actions/action-preview"
import { Button } from "@/components/shared/buttons/button"
import { Accordion } from "@/components/shared/utils/accordion"
import { Counter } from "@/components/shared/utils/counter"
import { DateSince } from "@/components/shared/utils/date-since"
import { cardColors, formatTitle } from "@/lib"
import { RouterOutputs } from "@/server/client"
import { useColumnStore } from "@/state/columns"

export const ActivityIcon: FC<{ status: string }> = ({ status }) => {
	switch (status) {
		case "upcoming":
			return (
				<div className="relative h-10 min-w-10">
					<div className="absolute mt-8 h-48 w-10 rounded-full bg-blue-400 blur-2xl filter" />
					<Loader
						className="absolute top-1/2 ml-auto h-4 w-6 -translate-y-1/2 text-center text-blue-400"
						size={16}
					/>
				</div>
			)
		case "completed":
			return (
				<div className="relative h-10 min-w-10">
					<div className="absolute mt-8 h-48 w-10 rounded-full bg-blue-400 blur-2xl filter" />
					<CheckCircle
						className="absolute top-1/2 ml-auto h-4 w-6 -translate-y-1/2 text-center text-blue-400"
						size={16}
					/>
				</div>
			)
		case "active":
			return (
				<div className="relative h-10 min-w-10">
					<div className="absolute mt-8 h-48 w-10 rounded-full bg-plug-green blur-2xl filter" />
					<CheckCircle
						className="absolute top-1/2 ml-auto h-4 w-6 -translate-y-1/2 text-center text-plug-green"
						size={16}
					/>
				</div>
			)
		case "paused":
			return (
				<div className="relative h-10 min-w-10">
					<div className="bg-text-plug-green/20 absolute mt-8 h-48 w-10 rounded-full blur-2xl filter" />
					<Pause
						className="absolute top-1/2 ml-auto h-4 w-6 -translate-y-1/2 text-center text-plug-green/20"
						size={16}
					/>
				</div>
			)
		case "success":
			return (
				<div className="relative h-10 min-w-10">
					<div className="absolute mt-8 h-48 w-10 rounded-full bg-plug-green blur-2xl filter" />
					<CheckCircle
						className="absolute top-1/2 ml-auto h-4 w-6 -translate-y-1/2 text-center text-plug-green"
						size={16}
					/>
				</div>
			)
		case "failure":
			return (
				<div className="relative h-10 min-w-10">
					<div className="absolute mt-8 h-48 w-10 rounded-full bg-plug-red blur-2xl filter" />
					<XCircle
						className="absolute top-1/2 h-4 w-6 -translate-y-1/2 text-center text-plug-red"
						size={16}
					/>
				</div>
			)
		default:
			return (
				<div className="relative h-10 min-w-10">
					<div className="absolute mt-8 h-48 w-10 rounded-full bg-yellow-400 blur-2xl filter" />
					<AlertCircle
						className="absolute top-1/2 h-4 w-6 -translate-y-1/2 text-center text-yellow-400"
						size={16}
					/>
				</div>
			)
	}
}

export const ActivityItem: FC<{
	index: number
	activity: RouterOutputs["plugs"]["activity"]["get"][number] | undefined
	simulationId: string | undefined
}> = ({ index, activity, simulationId }) => {
	const { handle } = useColumnStore(index, `${activity?.id}-activity`)

	const simulation = useMemo(
		() => activity?.simulations.find(sim => sim.id === simulationId),
		[activity, simulationId]
	)

	const actions = useMemo(() => JSON.parse(activity?.actions ?? "[]"), [activity])

	return (
		<>
			<Accordion loading={activity === undefined} onExpand={() => handle.frame()}>
				{activity === undefined ? (
					<div className="invisible">
						<p>.</p>
						<p>.</p>
					</div>
				) : (
					<div className="flex w-full flex-row items-center">
						<ActivityIcon status={activity.status} />

						<div
							className="mr-4 h-10 w-10 min-w-10 rounded-sm bg-plug-green/10"
							style={{
								backgroundImage: cardColors[activity.workflow.color]
							}}
						/>

						<div className="relative flex w-full flex-col overflow-hidden">
							<div className="flex flex-row items-center justify-between gap-2 font-bold">
								<p className="mr-2 truncate overflow-ellipsis whitespace-nowrap">
									{activity.workflow.name}
								</p>
								<div className="flex-shrink-0">
									<DateSince date={new Date(activity.createdAt)} />
								</div>
							</div>
							<div className="flex w-full flex-row items-center justify-between text-sm font-bold text-black text-opacity-40">
								<p>{formatTitle(activity.status)}</p>
								<p className="flex flex-row gap-2">
									<Counter count={activity.startAt.toLocaleDateString()} />
									<span className="opacity-60">→</span>
									{activity.endAt ? <Counter count={activity.endAt.toLocaleDateString()} /> : "∞"}
								</p>
							</div>
						</div>
					</div>
				)}
			</Accordion>

			<ExecutionFrame
				index={index}
				icon={<ActivityIcon status={activity?.status ?? "pending"} />}
				activity={activity}
			/>

			<Frame
				index={index}
				icon={<ActivityIcon status={simulation?.status ?? "pending"} />}
				label={`Simulation ${formatTitle(simulation?.status ?? "Pending")}`}
				visible={simulation !== undefined}
				handleBack={() => handle.frame()}
				hasOverlay={true}
			>
				{activity && (
					<ActionPreview
						index={index}
						item={activity.workflow.id}
						actions={actions}
						errors={simulation?.errors ?? []}
					/>
				)}

				{simulation?.error && (
					<p className="mx-auto mt-4 px-8 text-center text-sm font-bold text-plug-red">
						Warning: {simulation?.error}
					</p>
				)}

				{simulation?.status === "success" ? (
					<Button
						className="mt-4 flex w-full flex-row items-center justify-center gap-2 py-4"
						onClick={() => {}}
					>
						<Share size={18} className="opacity-60" />
						Share
					</Button>
				) : (
					<Button
						className="mt-4 flex w-full flex-row items-center justify-center gap-2 py-4"
						onClick={() => {}}
					>
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
					<span className="ml-auto">{formatTitle(simulation?.status ?? "pending")}</span>
				</p>
				<p className="flex flex-row items-center justify-between gap-4 font-bold">
					<Clock10 size={18} className="opacity-20" />
					<span className="opacity-40">Simulated</span>{" "}
					<span className="ml-auto">
						<DateSince date={new Date(simulation?.createdAt ?? new Date())} />
					</span>
				</p>
				{simulation?.gasEstimate && (
					<p className="flex flex-row items-center justify-between gap-4 font-bold">
						<CircleDollarSign size={18} className="opacity-20" />
						<span className="opacity-40">Fee</span>{" "}
						<span className="group ml-auto flex cursor-pointer flex-row items-center gap-4">
							<Counter count={simulation?.gasEstimate ?? 0} />
						</span>
					</p>
				)}
			</Frame>
		</>
	)
}
