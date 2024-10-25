import { FC, useEffect } from "react"
import { DateRange } from "react-day-picker"

import { Eye } from "lucide-react"

import { api } from "@/server/client"

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

	useEffect(() => {
		console.log("RunFrame effect triggered", { isFrame, scheduleData })
		if (!isFrame) {
			console.log("Clearing schedule from RunFrame")
			clearSchedule()
		}
	}, [isFrame, clearSchedule, scheduleData])

	// Also log in main render
	console.log("RunFrame render", { isFrame, scheduleData })

	const { plug, chains } = usePlugs(item)
	const prevFrame = "NOT_IMPLEMENTED" as string

	const queueMutation = api.plug.action.queue.useMutation()

	// const handleBack =
	// 	prevFrame !== "schedule"
	// 		? chains.length === 1
	// 			? undefined
	// 			: () => {
	// 					clearSchedule()
	// 					frame(`chain-${prevFrame}`)
	// 				}
	// 		: () => frame(`schedule`)

	const handleSubmit = async () => {
		if (!plug) return

		try {
			await queueMutation.mutateAsync({
				workflowId: plug.id,
				startAt: scheduleData?.date?.from ?? new Date(),
				endAt: scheduleData?.date?.to ?? undefined,
				frequency: scheduleData ? parseInt(scheduleData.repeats.value) : -1
			})
			clearSchedule()
			frame("running")
		} catch (error) {
			console.error("Failed to queue workflow:", error)
			// Handle error (e.g., show error message to user)
		}
	}

	return (
		<Frame
			index={index}
			className="z-[2]"
			// handleBack={handleBack}
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
