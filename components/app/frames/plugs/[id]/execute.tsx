import { FC } from "react"

import { RanFrame } from "./execute/ran"
import { RecurringFrame } from "./execute/recurring"
import { RunFrame } from "./execute/run"
import { ScheduleFrame } from "./execute/schedule"

export const ExecuteFrame: FC<{ index: number; item: string }> = ({ index, item }) => {
	return (
		<>
			<ScheduleFrame index={index} item={item} />
			<RecurringFrame index={index} />
			<RunFrame index={index} item={item} />
			<RanFrame index={index} item={item} />
		</>
	)
}
