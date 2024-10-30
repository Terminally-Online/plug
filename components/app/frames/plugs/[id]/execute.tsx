import { FC, useState, useCallback } from "react"
import { DateRange } from "react-day-picker"

import { frequencies } from "@/lib"

import { ChainFrame } from "./execute/chain"
import { RanFrame } from "./execute/ran"
import { RecurringFrame } from "./execute/recurring"
import { RunFrame } from "./execute/run"
import { ScheduleFrame } from "./execute/schedule"

export const ExecuteFrame: FC<{ index: number; item: string }> = ({ index, item }) => {
	const [scheduleData, setScheduleData] = useState<{
		date: DateRange | undefined
		repeats: (typeof frequencies)[0]
	}>({
		date: undefined,
		repeats: frequencies[0]
	})

	const clearSchedule = useCallback(() => {
		setScheduleData({
			date: undefined,
			repeats: frequencies[0]
		})
	}, [])

	const handleRepeats = useCallback((repeats: (typeof frequencies)[0]) => {
		setScheduleData(prev => ({
			date: prev.date,
			repeats
		}))
	}, [])

	return (
		<>
			<ChainFrame index={index} item={item} />
			<ScheduleFrame 
				index={index} 
				item={item} 
				scheduleData={scheduleData} 
				setScheduleData={setScheduleData} 
			/>
			<RecurringFrame index={index} handleRepeats={handleRepeats} />
			<RunFrame 
				index={index} 
				item={item} 
				scheduleData={scheduleData} 
				clearSchedule={clearSchedule} 
			/>
			<RanFrame index={index} item={item} />
		</>
	)
}
