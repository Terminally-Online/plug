import { useSession } from "next-auth/react"
import { FC } from "react"
import { DayPicker } from "react-day-picker"

import { CalendarPlus, ChevronLeft, ChevronRight, Clock } from "lucide-react"

import { Button, Dropdown, Frame } from "@/components"
import { cn, frequencies } from "@/lib"
import { useColumnStore } from "@/state"

export const ScheduleFrame: FC<{
	index: number
	item: string
}> = ({ index }) => {
	const { data: session } = useSession()
	const { column, isFrame, handle } = useColumnStore(index, "schedule")

	if (!column) return null

	return (
		<Frame
			index={index}
			icon={<CalendarPlus size={18} className="opacity-40" />}
			label="Schedule"
			visible={(isFrame && session && session.user.anonymous === false) || false}
			hasOverlay={true}
		>
			<div className="flex flex-col gap-4">
				<DayPicker
					mode="range"
					selected={column?.schedule?.date}
					onSelect={date => handle.schedule({ date, repeats: column?.schedule?.repeats || frequencies[0] })}
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
						months: "flex flex-col sm:flex-row space-y-4 sm:space-x-4 sm:space-y-0",
						month: "space-y-4 w-full",
						caption: "flex justify-center pt-1 relative items-center",
						caption_label: "text-sm font-bold opacity-40",
						nav: "space-x-1 flex items-center",
						nav_button: cn(
							"inline-flex items-center justify-center whitespace-nowrap rounded-sm text-sm ring-offset-background transition-colors focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:pointer-events-none disabled:opacity-50",
							"border border-input bg-background hover:bg-accent hover:text-accent-foreground",
							"p-1 bg-transparent opacity-40 hover:opacity-100"
						),
						nav_button_previous: "absolute left-1",
						nav_button_next: "absolute right-1",
						table: "w-full border-collapse space-y-1",
						head_row: "flex justify-between",
						head_cell: "rounded-sm my-2 w-9 font-bold text-sm opacity-40",
						row: "flex w-full mt-2 justify-between",
						cell: "h-9 w-9 text-center text-sm p-0 relative [&:has([aria-selected].day-range-end)]:rounded-r-sm [&:has([aria-selected].day-outside)]:bg-accent/50 first:[&:has([aria-selected])]:rounded-l-md last:[&:has([aria-selected])]:rounded-r-md focus-within:relative focus-within:z-20",
						day: cn(
							"inline-flex items-center justify-center whitespace-nowrap rounded-md text-sm font-bold ring-offset-background transition-colors focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:pointer-events-none disabled:opacity-50",
							"hover:bg-accent hover:text-accent-foreground h-7 w-7 bg-transparent p-0 hover:opacity-100"
						),
						day_range_end: "day-range-end",
						day_selected:
							"bg-gradient-to-tr from-plug-green to-plug-yellow text-primary-foreground hover:bg-primary hover:text-primary-foreground focus:bg-primary focus:text-primary-foreground",
						day_today: "bg-accent text-accent-foreground",
						day_disabled: "text-black/40",
						day_range_middle: "aria-selected:bg-accent aria-selected:text-primary-foreground",
						day_hidden: "invisible"
					}}
					components={{
						IconLeft: () => <ChevronLeft size={14} className="h-4 w-4" />,
						IconRight: () => <ChevronRight size={14} className="h-4 w-4" />
					}}
				/>

				<Dropdown
					icon={<Clock size={14} className="opacity-60" />}
					placeholder="Frequency"
					value={column.schedule?.repeats.label || "Once"}
					options={frequencies}
					handleClick={() => handle.frame("recurring")}
				/>

				<Button
					variant={
						column.schedule && column.schedule.date && column.schedule.date.from ? "primary" : "disabled"
					}
					className="w-full py-4"
					onClick={() => handle.frame("run")}
					disabled={!column.schedule || !column.schedule.date || !column.schedule.date.from}
				>
					{column.schedule && column.schedule.date && column.schedule.date.from ? "Next" : "Select a Date"}
				</Button>
			</div>
		</Frame>
	)
}
