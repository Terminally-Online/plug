import { ChainFrame } from "./execute/chain"
import { RanFrame } from "./execute/ran"
import { RunFrame } from "./execute/run"
import { RunningFrame } from "./execute/running"
import { ScheduleFrame } from "./execute/schedule"
import { SocketFrame } from "./execute/socket"

export const ExecuteFrame = () => (
	<>
		<SocketFrame />
		<ChainFrame />
		<ScheduleFrame />
		<RunFrame />
		<RunningFrame />
		<RanFrame />
	</>
)
