import { FC, useState } from "react"
import { DateRange } from "react-day-picker"

import { ChainFrame } from "./execute/chain"
import { RanFrame } from "./execute/ran"
import { frequencies, RecurringFrame } from "./execute/recurring"
import { RunFrame } from "./execute/run"
import { RunningFrame } from "./execute/running"
import { ScheduleFrame } from "./execute/schedule"

export const ExecuteFrame: FC<{ index: number; item: string }> = ({ index, item }) => {
	const [scheduleData, setScheduleData] = useState<{
		date: DateRange | undefined
		repeats: (typeof frequencies)[0]
	} | null>(null)

	const clearSchedule = () => setScheduleData(null)

	const handleRepeats = (repeats: (typeof frequencies)[0]) => {
		setScheduleData(prev => (prev ? { ...prev, repeats } : { date: undefined, repeats }))
	}

	return (
		<>
			<ChainFrame index={index} item={item} />
			<ScheduleFrame index={index} item={item} scheduleData={scheduleData} setScheduleData={setScheduleData} />
			<RecurringFrame index={index} handleRepeats={handleRepeats} />
			<RunFrame index={index} item={item} scheduleData={scheduleData} clearSchedule={clearSchedule} />
			<RanFrame index={index} item={item} />
		</>
	)
}
