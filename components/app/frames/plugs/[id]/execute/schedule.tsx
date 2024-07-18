import { useState } from "react"

import {
	ArrowRight,
	CalendarPlus,
	ChevronLeft,
	ChevronRight,
	Clock
} from "lucide-react"
import { DateRange, DayPicker } from "react-day-picker"

import { Button, Dropdown, Frame } from "@/components"
import { useFrame, usePlugs, useSockets } from "@/contexts"
import { cn, formatDate } from "@/lib"

const frequencies = [
	{ label: "Never", value: "0" },
	{ label: "Daily", value: "1" },
	{ label: "Weekly", value: "7" },
	{ label: "Monthly", value: "30" },
	{ label: "Quarterly", value: "90" },
	{ label: "Yearly", value: "365" }
]

export const ScheduleFrame = () => {
	const { frameVisible, handleFrameVisible } = useFrame()
	const { sockets } = useSockets()
	const { chainsAvailable } = usePlugs()

	const [date, setDate] = useState<DateRange | undefined>({
		from: undefined,
		to: undefined
	})
	const [frequency, setFrequency] = useState<(typeof frequencies)[0]>(
		frequencies[0]
	)

	const isFrame = frameVisible === "schedule"

	const handleBack =
		chainsAvailable.length === 1
			? sockets && sockets.length === 1
				? undefined
				: () => handleFrameVisible("socket-schedule")
			: () => handleFrameVisible("chain-schedule")

	return (
		<>
			<Frame
				className="z-[2]"
				handleBack={handleBack}
				icon={<CalendarPlus size={18} />}
				label="Choose Availability"
				visible={isFrame}
				hasOverlay={true}
			>
				<div className="flex flex-col gap-4">
					<div className="flex w-full flex-row items-center justify-between">
						{date && date.from && (
							<div
								className="rounded-md bg-grayscale-100 p-1 px-2 font-bold text-plug-green"
								style={{
									background: `linear-gradient(to right, rgba(0,239,54,0.1), rgba(147,223,0,0.1))`
								}}
							>
								{formatDate(
									date && date.from ? date.from : new Date()
								)}
							</div>
						)}

						{date && date.to && <ArrowRight size={14} />}

						{date && date.to && (
							<div
								className="rounded-md bg-grayscale-100 p-1 px-2 font-bold text-plug-green"
								style={{
									background: `linear-gradient(to right, rgba(0,239,54,0.1), rgba(147,223,0,0.1))`
								}}
							>
								{formatDate(date.to ?? new Date())}
							</div>
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
								"inline-flex items-center justify-center whitespace-nowrap rounded-full text-sm ring-offset-background transition-colors focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:pointer-events-none disabled:opacity-50",
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
							IconRight: () => (
								<ChevronRight className="h-4 w-4" />
							)
						}}
					/>

					<Dropdown
						icon={<Clock size={14} className="opacity-60" />}
						placeholder="Repeats"
						value={frequency.label}
						options={frequencies}
						handleClick={() => handleFrameVisible("recurring")}
					/>

					<Button
						variant={date && date.from ? "primary" : "disabled"}
						className="mt-4 w-full"
						onClick={() => handleFrameVisible("run-schedule")}
						disabled={!date || !date.from}
					>
						{date && date.from ? "Next" : "Select a Date"}
					</Button>
				</div>
			</Frame>

			<Frame
				className="scrollbar-hide z-[2] max-h-[calc(100vh-80px)] overflow-y-auto"
				icon={<Clock size={18} className="opacity-60" />}
				label="Recurring Frequency"
				visible={frameVisible === "recurring"}
				handleBack={() => handleFrameVisible("schedule")}
				hasOverlay={true}
			>
				<div className="flex flex-col gap-4">
					{frequencies.map(frequency => (
						<Button
							key={frequency.label}
							variant="secondary"
							onClick={() => {
								setFrequency(frequency)
								handleFrameVisible("schedule")
							}}
						>
							{frequency.label}
						</Button>
					))}
				</div>
			</Frame>
		</>
	)
}
