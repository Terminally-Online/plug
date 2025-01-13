import Image from "next/image"
import { FC, useEffect, useMemo, useState } from "react"

import { Bell, Calendar, Eye, Pause, Play, Waypoints } from "lucide-react"

import { Frame } from "@/components/app/frames/base"
import { ActionPreview } from "@/components/app/plugs/actions/action-preview"
import { ActivityIcon } from "@/components/app/sockets/activity/activity-item"
import { Button } from "@/components/shared/buttons/button"
import { Accordion } from "@/components/shared/utils/accordion"
import { Counter } from "@/components/shared/utils/counter"
import { DateSince } from "@/components/shared/utils/date-since"
import { TimeUntil } from "@/components/shared/utils/time-until"
import { useActivities } from "@/contexts"
import { cardColors, chains, cn, formatFrequency, formatTitle } from "@/lib"
import { RouterOutputs } from "@/server/client"
import { COLUMNS, useColumnStore } from "@/state/columns"

const ITEMS_PER_PAGE = 10

export const ExecutionFrame: FC<{
	index: number
	icon: JSX.Element
	activity: RouterOutputs["plugs"]["activity"]["get"][number] | undefined
}> = ({ index, icon, activity }) => {
	const { isFrame, handle } = useColumnStore(index, `${activity?.id}-activity`)
	const { handle: activityHandle } = useActivities()
	const [visibleCount, setVisibleCount] = useState(ITEMS_PER_PAGE)

	const actions = useMemo(() => JSON.parse(activity?.actions ?? "[]"), [activity])

	const visibleSimulations = useMemo(() => {
		if (!activity?.simulations) return []
		return activity.simulations.slice(0, visibleCount)
	}, [activity?.simulations, visibleCount])

	const totalSimulations = activity?.simulations?.length ?? 0
	const hasMore = visibleCount < totalSimulations

	// Reset visible count when activity changes
	useEffect(() => {
		setVisibleCount(ITEMS_PER_PAGE)
	}, [activity?.id])

	if (!activity) return null

	return (
		<>
			<Frame
				index={index}
				icon={icon}
				label={
					<span className="flex flex-row items-center gap-2 text-lg font-bold">
						<div
							className="mr-4 flex h-8 w-8 min-w-8 items-center justify-center rounded-sm bg-plug-green/10 text-white/60"
							style={{
								backgroundImage: cardColors[activity.workflow.color]
							}}
						/>
						<span>{activity.workflow.name}</span>
					</span>
				}
				visible={isFrame}
				hasOverlay={true}
			>
				<div className="flex flex-col">
					<ActionPreview index={index} item={activity.workflow.id} actions={actions} />

					<div className="my-4 flex flex-row items-center gap-2">
						<Button
							className={cn("flex w-max flex-row items-center justify-center gap-2 py-4", {
								"w-full": activity.status === "completed"
							})}
							onClick={() =>
								handle.navigate({
									index,
									key: COLUMNS.KEYS.PLUG,
									item: activity.workflow.id,
									from: COLUMNS.KEYS.ACTIVITY
								})
							}
						>
							<Eye size={14} className="opacity-60" />
							View
						</Button>
						{activity.status !== "completed" && (
							<Button
								variant="secondary"
								className="flex w-full flex-row items-center justify-center gap-2 py-4"
								onClick={() => activityHandle.toggle({ id: activity.id })}
							>
								{activity.status === "active" ? (
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
						)}
					</div>

					<div className="mb-2 flex flex-row items-center gap-4">
						<p className="font-bold opacity-40">Details</p>
						<div className="h-[2px] w-full bg-plug-green/10" />
					</div>

					<div className="flex flex-col">
						<p className="flex flex-row justify-between font-bold">
							<span className="flex w-full flex-row items-center gap-4">
								<Bell size={18} className="opacity-20" />
								<span className="opacity-40">Status</span>
							</span>{" "}
							{formatTitle(activity.status)}
						</p>
						<p className="flex flex-row justify-between font-bold">
							<span className="flex w-full flex-row items-center gap-4">
								<Calendar size={18} className="opacity-20" />
								<span className="opacity-40">Frequency</span>
							</span>{" "}
							{formatFrequency(activity.frequency)}
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
					</div>

					{(activity.status !== "paused" || activity.simulations.length > 0) && (
						<div className="mb-2 mt-4 flex flex-row items-center gap-4">
							<p className="font-bold opacity-40">Simulations</p>
							<div className="h-[2px] w-full bg-plug-green/10" />
						</div>
					)}

					<div className="flex flex-col gap-2">
						{activity.status !== "paused" && activity.nextSimulationAt && (
							<Accordion>
								<div className="flex flex-row gap-2">
									<ActivityIcon status="upcoming" />
									<div className="flex w-full flex-col">
										<div className="flex flex-row items-center justify-between gap-2 font-bold">
											<p>
												Simulation{" "}
												<span className="text-sm tabular-nums opacity-40">
													(#{activity.simulations.length + 1})
												</span>
											</p>
											<TimeUntil date={activity.nextSimulationAt} />
										</div>
										<div className="flex flex-row items-center justify-between gap-2 text-sm font-bold opacity-40">
											<p>Upcoming</p>
											<p>
												<Counter count={activity.nextSimulationAt.toLocaleDateString()} />
											</p>
										</div>
									</div>
								</div>
							</Accordion>
						)}

						{visibleSimulations.map((simulation, index) => (
							<Accordion key={index} onExpand={() => handle.frame(`${simulation.id}-simulation`)}>
								<div className="flex flex-row gap-2">
									<ActivityIcon status={simulation.status} />
									<div className="flex w-full flex-col">
										<div className="flex flex-row items-center justify-between gap-2 font-bold">
											<p>
												Simulation{" "}
												<span className="text-sm tabular-nums opacity-40">
													(#{activity.simulations.length - index})
												</span>
											</p>
											<DateSince date={simulation.createdAt} />
										</div>
										<div className="flex flex-row items-center justify-between gap-2 text-sm font-bold opacity-40">
											<p>{formatTitle(simulation.status)}</p>
											<p>
												<Counter count={simulation.createdAt.toLocaleDateString()} />
											</p>
										</div>
									</div>
								</div>
							</Accordion>
						))}

						{hasMore && (
							<Button
								variant="secondary"
								className="mt-2 py-4"
								onClick={() =>
									setVisibleCount(prev => Math.min(prev + ITEMS_PER_PAGE, totalSimulations))
								}
							>
								Load More ({totalSimulations - visibleCount} remaining)
							</Button>
						)}
					</div>
				</div>
			</Frame>
		</>
	)
}
