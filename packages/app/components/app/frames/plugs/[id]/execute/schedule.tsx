import { FC, lazy, Suspense, useMemo } from "react"

import { CalendarPlus, ChevronLeft, ChevronRight, Clock } from "lucide-react"

import { Frame } from "@/components/app/frames/base"
import { Dropdown } from "@/components/app/inputs/dropdown"
import { Button } from "@/components/shared/buttons/button"
import { cn, frequencies, useConnect } from "@/lib"
import { columnByIndexAtom, isFrameAtom, useColumnActions } from "@/state/columns"
import { useAtom, useAtomValue } from "jotai"

const DayPicker = lazy(() => import("react-day-picker").then(mod => ({ default: mod.DayPicker })))

export const ScheduleFrame: FC<{
	index: number
	item: string
}> = ({ index }) => {
	const {
		account: { isAuthenticated }
	} = useConnect()

	const [column] = useAtom(columnByIndexAtom(index))
	const frameKey="schedule"
	const isFrame = useAtomValue(isFrameAtom)(column, frameKey)
	const { frame, schedule } = useColumnActions(index, frameKey)

	const MemoizedDayPicker = useMemo(() => {
		if (!column) return null

		return (
			<Suspense
				fallback={<div className="flex h-72 w-full items-center justify-center">Loading calendar...</div>}
			>
				<DayPicker
					mode="range"
					selected={column?.schedule?.date}
					onSelect={date => schedule({ date, repeats: column?.schedule?.repeats || frequencies[0] })}
					showOutsideDays
					fixedWeeks
					weekStartsOn={1}
					disabled={{
						before: new Date()
					}}
					className={cn(
						"select-none",
						column.schedule &&
						column.schedule.date &&
						(column.schedule.date.from || column.schedule.date.to)
					)}
					classNames={{
						months: "flex flex-col relative",
						month: "space-y-4 w-full",
						month_caption: "flex justify-center pt-1 relative items-center",
						caption_label: "text-sm font-bold opacity-40",
						nav: "flex items-center",
						button_previous: cn(
							"rounded-sm p-1 bg-white border-[1px] border-plug-green/10 text-black hover:bg-plug-green/10 items-center flex justify-center text-opacity-60 whitespace-nowrap [&.active]:bg-plug-green/5 [&.active]:text-opacity-100 [&.active]:hover:bg-plug-green/10 [&.active]:hover:border-plug-green/5",
							"absolute left-0 top-1 z-[20]",
						),
						button_next: cn(
							"rounded-sm p-1 bg-white border-[1px] border-plug-green/10 text-black hover:bg-plug-green/10 items-center flex justify-center text-opacity-60 whitespace-nowrap [&.active]:bg-plug-green/5 [&.active]:text-opacity-100 [&.active]:hover:bg-plug-green/10 [&.active]:hover:border-plug-green/5",
							"absolute right-0 top-1 z-[20]",
						),
						month_grid: "w-full border-collapse",
						weekdays: "flex gap-2 items-center w-full",
						weekday: "rounded-sm my-2 w-9 font-bold text-sm opacity-40 w-full",
						week: "flex w-full justify-between gap-2 mt-2",
						day: "rounded-sm h-full w-full text-center text-sm p-0 relative [&:has([aria-selected].day-outside)]:bg-accent/50 focus-within:relative focus-within:z-20",
						day_button: cn(
							"min-w-7 min-h-7",
							"inline-flex w-full items-center justify-center whitespace-nowrap rounded-md text-sm font-bold transition-colors focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:pointer-events-none disabled:opacity-50",
							"hover:brightness-105 focus:brightness-105 p-0"
						),
						range_middle: "aria-selected:bg-plug-yellow aria-selected:text-text-plug-green",
						range_end: "day-range-end",
						selected:
							"bg-plug-yellow text-plug-green hover:brightness-105 focus:brightness-105",
						today: "bg-plug-green/10 text-black/40",
						disabled: "text-black/40",
						hidden: "invisible"
					}}
					components={{
						Chevron: (props) => {
							if (props.orientation === "left")
								return <ChevronLeft size={14} className="h-4 w-4 opacity-60" />
							return <ChevronRight size={14} className="h-4 w-4 opacity-60" />
						}
					}}
				/>
			</Suspense>
		)
	}, [column, schedule])

	if (!column) return null

	return (
		<Frame
			index={index}
			icon={<CalendarPlus size={18} className="opacity-40" />}
			label={
				<span className="text-lg font-bold">
					<span className="opacity-40">Run:</span> Schedule
				</span>
			}
			visible={(isFrame && isAuthenticated) || false}
			hasOverlay={true}
		>
			<div className="flex flex-col gap-4">
				{MemoizedDayPicker}

				<Dropdown
					icon={<Clock size={14} className="opacity-60" />}
					placeholder="Frequency"
					value={column.schedule?.repeats.label || "Once"}
					options={frequencies}
					handleClick={() => frame("recurring")}
				/>

				<Button
					variant={
						column.schedule && column.schedule.date && column.schedule.date.from
							? "primary"
							: "primaryDisabled"
					}
					className="w-full py-4"
					onClick={() => frame("run")}
					disabled={!column.schedule || !column.schedule.date || !column.schedule.date.from}
				>
					{column.schedule && column.schedule.date && column.schedule.date.from ? "Next" : "Select a Date"}
				</Button>
			</div>
		</Frame>
	)
}
