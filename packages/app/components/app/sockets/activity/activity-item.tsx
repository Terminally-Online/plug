import Image from "next/image"
import { FC, useMemo } from "react"

import {
	AlertCircle,
	Bell,
	CheckCircle,
	CircleDollarSign,
	Clock10,
	Fence,
	FileWarning,
	Globe,
	Hash,
	Loader,
	Pause,
	Waypoints,
	XCircle
} from "lucide-react"

import { Accordion, Counter, DateSince, ExecutionFrame, Frame } from "@/components"
import { chains, formatTitle, getChainId } from "@/lib"
import { RouterOutputs } from "@/server/client"
import { useColumnStore } from "@/state"

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

	return (
		<>
			<Accordion loading={activity === undefined} onExpand={() => handle.frame()}>
				{activity === undefined ? (
					<div className="invisible">
						<p>.</p>
						<p>.</p>
					</div>
				) : (
					<div className="flex w-full flex-row">
						<ActivityIcon status={activity.status} />

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
				<div className="flex flex-row items-center gap-4">
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
				<p className="flex flex-row items-center justify-between gap-4 font-bold">
					<Waypoints size={18} className="opacity-20" />
					<span className="opacity-40">Chain</span>
					<span className="ml-auto flex flex-row items-center gap-2">
						<Image
							className="h-4 w-4"
							src={chains[getChainId("ethereum")].logo}
							alt="Ethereum"
							width={24}
							height={24}
						/>
						Ethereum
					</span>
				</p>
				<p className="flex flex-row items-center justify-between gap-4 font-bold">
					<CircleDollarSign size={18} className="opacity-20" />
					<span className="opacity-40">Fee</span>{" "}
					<span className="group ml-auto flex cursor-pointer flex-row items-center gap-4">
						<Counter count={simulation?.gasEstimate ?? 0} />
					</span>
				</p>
				{simulation?.error && (
					<p className="flex flex-row items-center justify-between gap-4 font-bold">
						<FileWarning size={18} className="opacity-20" />
						<span className="opacity-40">Error</span>{" "}
						<span className="group ml-auto flex cursor-pointer flex-row items-center gap-4">
							{simulation?.error}
						</span>
					</p>
				)}
			</Frame>
		</>
	)
}
