import { FC, useCallback, useEffect, useMemo } from "react"
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
	const { plug, actions, handle } = usePlugs(item)

	const isReady = useMemo(
		() => plug && actions && actions.every(action => action.values.every(value => Boolean(value))),
		[plug, actions]
	)

	const handleRun = useCallback(() => {
		if (!plug) return

		handle.plug.queue({
			workflowId: plug.id,
			startAt: scheduleData?.date?.from ?? new Date(),
			endAt: scheduleData?.date?.to ?? new Date(),
			frequency: parseInt(scheduleData?.repeats?.value ?? "0")
		})

		clearSchedule()
		frame("running")
	}, [plug, scheduleData, clearSchedule, frame, handle.plug])

	useEffect(() => {
		if (!isFrame) clearSchedule()
	}, [isFrame, clearSchedule])

	return (
		<Frame
			index={index}
			className="z-[2]"
			icon={<Eye size={18} />}
			label="Preview"
			visible={isFrame}
			hasOverlay={true}
		>
			<div className="flex flex-col">
				<ActionPreview index={index} item={item} review={true} />

				<p className="mt-2 flex font-bold">
					<span className="mr-auto opacity-40">Run On</span>
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

				<Button
					variant={isReady ? "primary" : "disabled"}
					className="mt-4 w-full py-4"
					onClick={handleRun}
					disabled={!isReady}
				>
					{isReady ? "Run" : "Missing required values"}
				</Button>
			</div>
		</Frame>
	)
}
