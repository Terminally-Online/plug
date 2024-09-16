import { FC } from "react"

import { ChainFrame } from "./execute/chain"
import { RanFrame } from "./execute/ran"
import { RecurringFrame } from "./execute/recurring"
import { RunFrame } from "./execute/run"
import { RunningFrame } from "./execute/running"
import { ScheduleFrame } from "./execute/schedule"

export const ExecuteFrame: FC<{ index: number; item: string }> = ({ index, item }) => (
	<>
		<ChainFrame index={index} item={item} />
		<ScheduleFrame index={index} item={item} />
		<RecurringFrame index={index} />
		<RunFrame index={index} item={item} />
		<RunningFrame index={index} item={item} />
		<RanFrame index={index} item={item}  />
	</>
)
