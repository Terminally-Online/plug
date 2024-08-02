import { ChainFrame } from "./execute/chain"
import { RanFrame } from "./execute/ran"
import { RunFrame } from "./execute/run"
import { RunningFrame } from "./execute/running"
import { ScheduleFrame } from "./execute/schedule"

export const ExecuteFrame = () => (
	<>
		<ChainFrame />
		<ScheduleFrame />
		<RunFrame />
		<RunningFrame />
		<RanFrame />
	</>
)
