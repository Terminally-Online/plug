import { FC, useCallback, useEffect, useMemo } from "react"
import { DateRange } from "react-day-picker"

import { CircleDollarSign, Eye, Globe, Waypoints } from "lucide-react"

import { ActionPreview, Button, Frame, Image } from "@/components"
import { usePlugs } from "@/contexts/PlugProvider"
import { chains } from "@/lib"
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
			icon={<Eye size={18} className="opacity-40" />}
			label="Preview"
			visible={isFrame}
			hasOverlay={true}
		>
			<div className="flex flex-col">
				<ActionPreview index={index} item={item} />

				<div className="mb-2 mt-4 flex flex-row items-center gap-4">
					<p className="font-bold opacity-40">Details</p>
					<div className="h-[2px] w-full bg-grayscale-100" />
				</div>

				<p className="flex w-full flex-row items-center gap-4 font-bold">
					<Waypoints size={18} className="opacity-20" />
					<span className="mr-auto opacity-40">Chain</span>
					<span className="flex flex-row items-center gap-2">
						<Image className="h-4 w-4" src={chains[1].logo} alt="ethereum" width={24} height={24} />
						Ethereum
					</span>
				</p>

				<p className="flex flex-row justify-between font-bold">
					<span className="flex w-full flex-row items-center gap-4">
						<CircleDollarSign size={18} className="opacity-20" />
						<span className="opacity-40">Fee</span>
					</span>{" "}
					<span className="flex w-full flex-row justify-end gap-2">
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
