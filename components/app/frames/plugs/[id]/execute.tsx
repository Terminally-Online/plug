import { FC } from "react"

import { ChainFrame } from "./execute/chain"
import { RanFrame } from "./execute/ran"
import { RecurringFrame } from "./execute/recurring"
import { RunFrame } from "./execute/run"
import { RunningFrame } from "./execute/running"
import { ScheduleFrame } from "./execute/schedule"

export const ExecuteFrame: FC<{ id: string }> = ({ id }) => (
	<>
		<ChainFrame id={id} />
		<ScheduleFrame id={id} />
		<RecurringFrame id={id} />
		<RunFrame id={id} />
		<RunningFrame id={id} />
		<RanFrame id={id} />
	</>
)
