import { FC, useCallback, useEffect } from "react"
import { DateRange } from "react-day-picker"

import { Eye } from "lucide-react"

import { ActionPreview, Button, Frame, Image } from "@/components"
import { usePlugs } from "@/contexts/PlugProvider"
import { useColumns } from "@/state"

import { frequencies } from "./recurring"

export const RunFrame: FC<{
	index: number
	item: string
	scheduleData: { date: DateRange | undefined; repeats: (typeof frequencies)[0] } | null
	clearSchedule: () => void
}> = ({ index, item, scheduleData, clearSchedule }) => {
	const { isFrame, frame } = useColumns(index, "run")
	const { plug, handle } = usePlugs(item)

	useEffect(() => {
		if (!isFrame) {
			clearSchedule()
		}
	}, [isFrame, clearSchedule, scheduleData])

	const prevFrame = "NOT_IMPLEMENTED" as string

	const handleRun = useCallback(() => {
		if (!plug) return

		if (scheduleData?.date?.from) {
			handle.plug.queue({
				workflowId: plug.id,
				startAt: scheduleData.date.from,
				endAt: scheduleData.date.to,
				frequency: parseInt(scheduleData.repeats.value)
			})
		} else {
			// Immediate execution logic
		}

		clearSchedule()
		frame("running")
	}, [plug, scheduleData, clearSchedule, frame, handle.plug])

	return (
		<Frame
			index={index}
			className="z-[2]"
			icon={<Eye size={18} />}
			label={prevFrame === "schedule" ? "Intent Preview" : "Transaction Preview"}
			visible={isFrame}
			hasOverlay={true}
		>
			<div className="flex flex-col gap-2">
				<ActionPreview index={index} item={item} review={true} />

				<p className="flex font-bold">
					<span className="mr-auto opacity-40">Run On</span>
					{/* TODO: ADD CHAIN FUNCTIONALITY BACK */}
					<Image
						className="ml-[-20px] h-6 w-6"
						src={`/blockchain/ethereum.png`}
						alt="ethereum"
						width={24}
						height={24}
					/>
				</p>

				<p className="flex font-bold">
					<span className="mr-auto opacity-40">Fee</span>
					<span className="flex flex-row gap-2">
						<span className="opacity-40">0.0011 ETH</span>
						<span>$4.19</span>
					</span>
				</p>

				<Button className="mt-4 w-full" onClick={handleRun}>
					{prevFrame === "schedule" ? "Sign Intent" : "Submit Transaction"}
				</Button>
			</div>
		</Frame>
	)
}
