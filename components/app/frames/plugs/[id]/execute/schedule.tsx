import { useEffect, useState } from "react"

import { motion } from "framer-motion"
import {
	ArrowRight,
	CalendarPlus,
	ChevronDown,
	ChevronLeft,
	ChevronRight
} from "lucide-react"
import { DateRange, DayPicker } from "react-day-picker"

import { Button } from "@/components/buttons"
import { Checkbox } from "@/components/inputs"
import { useFrame, usePlugs, useSockets } from "@/contexts"
import { cn } from "@/lib"
import { formatDate } from "@/lib/functions"

import { Frame } from "../../../base"

export const ScheduleFrame = () => {
	const { frameVisible, handleFrameVisible } = useFrame()
	const { sockets } = useSockets()
	const { chainsAvailable } = usePlugs()

	const [calendarOpen, setCalendarOpen] = useState<
		"start" | "end" | undefined
	>(undefined)
	const [date, setDate] = useState<DateRange | undefined>({
		from: undefined,
		to: undefined
	})
	const [advanced, setAdvanced] = useState(false)
	const [allDay, setAllDay] = useState(false)

	const isFrame = frameVisible === "schedule"

	const handleBack =
		chainsAvailable.length === 1
			? sockets && sockets.length === 1
				? undefined
				: () => handleFrameVisible("socket-schedule")
			: () => handleFrameVisible("chain-schedule")

	useEffect(() => setAdvanced(false), [frameVisible])

	return (
		<Frame
			className="z-[2]"
			handleBack={handleBack}
			icon={<CalendarPlus size={18} className="opacity-60" />}
			label="Choose Availability"
			visible={isFrame}
			hasOverlay={true}
		>
			<div className="flex flex-col gap-4">
				<div className="flex w-full flex-row items-center justify-between">
					{date && date.from && (
						<button
							className="rounded-md bg-grayscale-100 p-1 px-2 font-bold text-plug-green"
							style={{
								background: `linear-gradient(to right, rgba(0,239,54,0.1), rgba(147,223,0,0.1))`
							}}
							onClick={() =>
								setCalendarOpen(prevCalendarOpen =>
									prevCalendarOpen === "start"
										? undefined
										: "start"
								)
							}
						>
							{formatDate(
								date && date.from ? date.from : new Date()
							)}
						</button>
					)}

					{date && date.to && (
						<ArrowRight size={14} className="opacity-60" />
					)}

					{date && date.to && (
						<button
							className="rounded-md bg-grayscale-100 p-1 px-2 font-bold text-plug-green"
							style={{
								background: `linear-gradient(to right, rgba(0,239,54,0.1), rgba(147,223,0,0.1))`
							}}
							onClick={() =>
								setCalendarOpen(prevCalendarOpen =>
									prevCalendarOpen === "end"
										? undefined
										: "end"
								)
							}
						>
							{formatDate(date.to ?? new Date())}
						</button>
					)}
				</div>

				<DayPicker
					mode="range"
					selected={date}
					onSelect={setDate}
					showOutsideDays
					fixedWeeks
					weekStartsOn={1}
					disabled={{
						before: new Date()
					}}
					className={cn(
						"select-none",
						date && (date.from || date.to) && "mt-4"
					)}
					classNames={{
						months: "flex flex-col sm:flex-row space-y-4 sm:space-x-4 sm:space-y-0",
						month: "space-y-4",
						caption:
							"flex justify-center pt-1 relative items-center",
						caption_label: "text-sm font-bold opacity-40",
						nav: "space-x-1 flex items-center",
						nav_button: cn(
							"inline-flex items-center justify-center whitespace-nowrap rounded-md text-sm ring-offset-background transition-colors focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:pointer-events-none disabled:opacity-50",
							"border border-input bg-background hover:bg-accent hover:text-accent-foreground",
							"h-7 w-7 bg-transparent p-0 opacity-50 hover:opacity-100"
						),
						nav_button_previous: "absolute left-1",
						nav_button_next: "absolute right-1",
						table: "w-full border-collapse space-y-1",
						head_row: "flex justify-between",
						head_cell:
							"rounded-md my-2 w-9 font-bold text-[0.8rem] opacity-40",
						row: "flex w-full mt-2 justify-between",
						cell: "h-9 w-9 text-center text-sm p-0 relative [&:has([aria-selected].day-range-end)]:rounded-r-md [&:has([aria-selected].day-outside)]:bg-accent/50 first:[&:has([aria-selected])]:rounded-l-md last:[&:has([aria-selected])]:rounded-r-md focus-within:relative focus-within:z-20",
						day: cn(
							"inline-flex items-center justify-center whitespace-nowrap rounded-md text-sm font-bold ring-offset-background transition-colors focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:pointer-events-none disabled:opacity-50",
							"hover:bg-accent hover:text-accent-foreground",
							"h-7 w-7 bg-transparent p-0 hover:opacity-100"
						),
						day_range_end: "day-range-end",
						day_selected:
							"bg-gradient-to-tr from-plug-green to-plug-yellow text-primary-foreground hover:bg-primary hover:text-primary-foreground focus:bg-primary focus:text-primary-foreground",
						day_today: "bg-accent text-accent-foreground",
						day_disabled: "text-black/40",
						day_range_middle:
							"aria-selected:bg-accent aria-selected:text-primary-foreground",
						day_hidden: "invisible"
					}}
					components={{
						IconLeft: () => <ChevronLeft className="h-4 w-4" />,
						IconRight: () => <ChevronRight className="h-4 w-4" />
					}}
				/>

				<button
					className="group mt-4 flex w-full cursor-pointer flex-row gap-2 font-bold opacity-40"
					onClick={() => setAdvanced(!advanced)}
				>
					<span className="mr-auto">Advanced</span>
					<Button
						variant="secondary"
						className={cn(
							"p-1 transition-all duration-200 ease-in-out group-hover:bg-grayscale-100",
							advanced && "bg-grayscale-100"
						)}
						onClick={() => setAdvanced(!advanced)}
					>
						<motion.div
							animate={{ rotate: advanced ? 180 : 0 }}
							transition={{ duration: 0.2 }}
						>
							<ChevronDown size={14} />
						</motion.div>
					</Button>
				</button>

				{advanced && (
					<>
						<div className="flex flex-row items-center gap-2">
							<p className="mr-auto font-bold">All-day</p>

							<Checkbox
								checked={allDay}
								handleChange={setAllDay}
							/>
						</div>
					</>
				)}

				<Button
					variant={date && date.from ? "primary" : "disabled"}
					className="mt-4 w-full"
					onClick={() => handleFrameVisible("run-schedule")}
					disabled={!date || !date.from}
				>
					{date && date.from ? "Next" : "Choose Dates"}
				</Button>
			</div>
		</Frame>
	)
}
