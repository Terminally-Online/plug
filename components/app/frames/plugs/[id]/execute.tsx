import { FC, useCallback, useState } from "react"
import { DateRange } from "react-day-picker"

import { frequencies } from "@/lib"

import { ChainFrame } from "./execute/chain"
import { RanFrame } from "./execute/ran"
import { RecurringFrame } from "./execute/recurring"
import { RunFrame } from "./execute/run"
import { ScheduleFrame } from "./execute/schedule"

export const ExecuteFrame: FC<{ index: number; item: string }> = ({ index, item }) => {
	const [schedule, setSchedule] = useState<{
		date: DateRange | undefined
		repeats: (typeof frequencies)[0]
	}>({
		date: undefined,
		repeats: frequencies[0]
	})

	return (
		<>
			<ChainFrame index={index} item={item} />
			<ScheduleFrame index={index} item={item} schedule={schedule} onSchedule={setSchedule} />
			<RecurringFrame
				index={index}
				handleRepeats={repeats =>
					setSchedule(prev => ({
						date: prev.date,
						repeats
					}))
				}
			/>
			<RunFrame
				index={index}
				item={item}
				scheduleData={schedule}
				clearSchedule={() =>
					setSchedule({
						date: undefined,
						repeats: frequencies[0]
					})
				}
			/>
			<RanFrame index={index} item={item} />
		</>
	)
}
