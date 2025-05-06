import { FC, type JSX, useEffect, useMemo, useRef, useState } from "react"

import { Bell, Calendar, Eye, Pause, Play, Waypoints } from "lucide-react"

import { useAtom, useAtomValue } from "jotai"

import { Frame } from "@/components/app/frames/base"
import { ActionPreview } from "@/components/app/plugs/actions/action-preview"
import { ActivityIcon } from "@/components/app/sockets/activity/activity-item"
import { Button } from "@/components/shared/buttons/button"
import { Accordion } from "@/components/shared/utils/accordion"
import { Counter } from "@/components/shared/utils/counter"
import { DateSince } from "@/components/shared/utils/date-since"
import { TimeUntil } from "@/components/shared/utils/time-until"
import { useActivities } from "@/contexts"
import { cardColors, ChainId, cn, formatFrequency, formatTitle, getChainName } from "@/lib"
import { RouterOutputs } from "@/server/client"
import { columnByIndexAtom, COLUMNS, isFrameAtom, useColumnActions } from "@/state/columns"

import { ChainImage } from "../../../sockets/chains/chain.image"

const ITEMS_PER_PAGE = 10

export const ExecutionFrame: FC<{
	index: number
	icon: JSX.Element
	activity: RouterOutputs["plugs"]["activity"]["get"][number]
}> = ({ index, icon, activity }) => {
	const loadMoreRef = useRef<HTMLDivElement>(null)

	const [column] = useAtom(columnByIndexAtom(index))
	const frameKey = `${activity?.id}-activity`
	const isFrame = useAtomValue(isFrameAtom)(column, frameKey)
	const { frame, navigate } = useColumnActions(index, frameKey)

	const { handle: activityHandle } = useActivities()
	const [visibleCount, setVisibleCount] = useState(ITEMS_PER_PAGE)

	const actions = activity?.inputs ?? []

	const { visibleRuns, totalRuns, hasMore } = useMemo(
		() => ({
			visibleRuns: activity?.runs?.slice(0, visibleCount) ?? [],
			totalRuns: activity?.runs?.length ?? 0,
			hasMore: visibleCount < (activity?.runs?.length ?? 0)
		}),
		[activity?.runs, visibleCount]
	)

	useEffect(() => {
		if (!isFrame || !activity) return
		setVisibleCount(ITEMS_PER_PAGE)

		const observer = new IntersectionObserver(
			entries => {
				if (entries[0].isIntersecting && hasMore) {
					setVisibleCount(prev => Math.min(prev + ITEMS_PER_PAGE, totalRuns))
				}
			},
			{ threshold: 0.1 }
		)

		if (loadMoreRef.current) {
			observer.observe(loadMoreRef.current)
		}

		return () => observer.disconnect()
	}, [isFrame, activity, hasMore, totalRuns])

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
								backgroundImage: activity.plug?.color
									? cardColors[activity.plug.color]
									: cardColors["plug"]
							}}
						/>
						<span>{activity.inputs.map(input => formatTitle(input.action)).join(", ")}</span>
					</span>
				}
				visible={isFrame}
				hasOverlay={true}
			>
				<div className="flex flex-col">
					<ActionPreview
						index={index}
						item={activity.plug?.id ?? activity.id}
						actions={actions}
						errors={visibleRuns[0]?.errors ?? []}
					/>

					<div className="my-4 flex flex-row items-center gap-2">
						<Button
							className={cn("flex w-max flex-row items-center justify-center gap-2 py-4", {
								"w-full": activity.status === "completed"
							})}
							onClick={() =>
								navigate({
									index,
									key: COLUMNS.KEYS.PLUG,
									item: activity.plug?.id,
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
						<div className="h-[1px] w-full bg-plug-green/10" />
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
								<ChainImage chainId={activity.chainId as ChainId} size="xs" />
								{getChainName(activity.chainId as ChainId)}
							</span>
						</p>

						<div className="mb-2 mt-4 flex flex-row items-center gap-4">
							<p className="font-bold opacity-40">Schedule</p>
							<div className="h-[1px] w-full bg-plug-green/10" />
						</div>
						<p className="flex flex-row justify-between font-bold">
							<span className="flex w-full flex-row items-center gap-4">
								<Calendar size={18} className="opacity-20" />
								<span className="opacity-40">Frequency</span>
							</span>{" "}
							{formatFrequency(activity.frequency)}
						</p>
						<p className="flex flex-row justify-between font-bold">
							<span className="flex w-full flex-row items-center gap-4">
								<Play size={18} className="opacity-20" />
								<span className="opacity-40">Start At</span>
							</span>{" "}
							<Counter count={new Date(activity.startAt).toLocaleDateString()} />
						</p>
						{activity.endAt && (
							<p className="flex flex-row justify-between font-bold">
								<span className="flex w-full flex-row items-center gap-4">
									<Pause size={18} className="opacity-20" />
									<span className="opacity-40">Stop At</span>
								</span>{" "}
								<Counter count={new Date(activity.endAt).toLocaleDateString()} />
							</p>
						)}
					</div>

					{(activity.status !== "paused" || activity.runs.length > 0) && (
						<div className="mb-2 mt-4 flex flex-row items-center gap-4">
							<p className="font-bold opacity-40">Runs</p>
							<div className="h-[1px] w-full bg-plug-green/10" />
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
												Run{" "}
												<span className="text-sm tabular-nums opacity-40">
													(#{activity.runs.length + 1})
												</span>
											</p>
											<TimeUntil date={new Date(activity.nextSimulationAt)} />
										</div>
										<div className="flex flex-row items-center justify-between gap-2 text-sm font-bold opacity-40">
											<p>Upcoming</p>
											<p>
												<Counter
													count={new Date(activity.nextSimulationAt).toLocaleDateString()}
												/>
											</p>
										</div>
									</div>
								</div>
							</Accordion>
						)}

						{visibleRuns.map((run, index) => (
							<Accordion key={run.id} onExpand={() => frame(`${run.id}-simulation`)}>
								<div className="flex flex-row gap-2">
									<ActivityIcon status={run.status} />
									<div className="flex w-full flex-col">
										<div className="flex flex-row items-center justify-between gap-2 font-bold">
											<p>
												Run{" "}
												<span className="text-sm tabular-nums opacity-40">
													(#{activity.runs.length - index})
												</span>
											</p>
											<DateSince date={new Date(run.createdAt)} />
										</div>
										<div className="flex flex-row items-center justify-between gap-2 text-sm font-bold opacity-40">
											<p>{formatTitle(run.status)}</p>
											<p>
												<Counter count={new Date(run.createdAt).toLocaleDateString()} />
											</p>
										</div>
									</div>
								</div>
							</Accordion>
						))}

						{hasMore && <div ref={loadMoreRef} className="h-4" />}
					</div>
				</div>
			</Frame>
		</>
	)
}
