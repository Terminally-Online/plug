import { FC, useState } from "react"

import { ChainFrame } from "./execute/chain"
import { RanFrame } from "./execute/ran"
import { frequencies, RecurringFrame } from "./execute/recurring"
import { RunFrame } from "./execute/run"
import { RunningFrame } from "./execute/running"
import { ScheduleFrame } from "./execute/schedule"

export const ExecuteFrame: FC<{ index: number; item: string }> = ({ index, item }) => {
	const [repeats, setRepeats] = useState<(typeof frequencies)[0]>(frequencies[0])

	return (
		<>
			<ChainFrame index={index} item={item} />
			<ScheduleFrame index={index} item={item} repeats={repeats} />
			<RecurringFrame index={index} handleRepeats={setRepeats} />
			<RunFrame index={index} item={item} />
			<RunningFrame index={index} item={item} />
			<RanFrame index={index} item={item} />
		</>
	)
}
