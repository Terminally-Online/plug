import { FC, useEffect, useCallback } from "react"
import { DateRange } from "react-day-picker"

import { Eye } from "lucide-react"

import { ActionPreview, Button, Frame, Image } from "@/components"
import { usePlugs } from "@/contexts/PlugProvider"
import { useColumns } from "@/state"

export const RunFrame: FC<{
	index: number
	item: string
	scheduleData: { date: DateRange | undefined; repeats: { value: string } } | null
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



	const handleSubmit = useCallback(async () => {
		if (!plug) return
	  
		// Call the queue mutation
		handle.plug.queue({
		  workflowId: plug.id,
		  startAt: scheduleData?.date?.from ?? new Date(),
		  endAt: scheduleData?.date?.to ?? undefined,
		  frequency: scheduleData ? parseInt(scheduleData.repeats.value) : -1
		})
		
		clearSchedule()
		frame("running")
	  }, [plug, scheduleData, handle.plug.queue, clearSchedule, frame])


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
				<p className="font-bold opacity-60">Actions</p>
				<ActionPreview index={index} item={item} />

				<p className="flex font-bold">
					<span className="mr-auto opacity-60">Run On</span>
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
					<span className="mr-auto opacity-60">Fee</span>
					<span className="flex flex-row gap-2">
						<span className="opacity-40">0.0011 ETH</span>
						<span>$4.19</span>
					</span>
				</p>

				<Button className="mt-4 w-full" onClick={handleSubmit}>
					{prevFrame === "schedule" ? "Sign Intent" : "Submit Transaction"}
				</Button>
			</div>
		</Frame>
	)
}
