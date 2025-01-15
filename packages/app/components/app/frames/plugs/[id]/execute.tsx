import { FC } from "react"

import { RanFrame } from "@/components/app/frames/plugs/[id]/execute/ran"
import { RecurringFrame } from "@/components/app/frames/plugs/[id]/execute/recurring"
import { RunFrame } from "@/components/app/frames/plugs/[id]/execute/run"
import { ScheduleFrame } from "@/components/app/frames/plugs/[id]/execute/schedule"

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
