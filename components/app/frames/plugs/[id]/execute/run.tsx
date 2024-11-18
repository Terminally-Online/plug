import { useSession } from "next-auth/react"
import { FC, useCallback, useMemo } from "react"

import { Calendar, CircleDollarSign, Eye, Pause, Play, Waypoints } from "lucide-react"

import { ActionPreview, Button, Counter, Frame, Image } from "@/components"
import { chains } from "@/lib"
import { useColumnStore, usePlugStore } from "@/state"

export const RunFrame: FC<{
	index: number
	item: string
}> = ({ index, item }) => {
	const { data: session } = useSession()
	const { column, isFrame, handle } = useColumnStore(index, "run")
	const { plug, actions, handle: plugHandle } = usePlugStore(item)

	const isReady = useMemo(
		() => plug && actions && actions.length > 0 && actions.every(action => action.values.every(value => Boolean(value))),
		[plug, actions]
	)

	const handleRun = useCallback(() => {
		if (!column || !column.item) return

		plugHandle.plug.queue({
			workflowId: column.item,
			startAt: column.schedule?.date?.from ?? new Date(),
			endAt: column.schedule?.date?.to ?? new Date(),
			frequency: parseInt(column.schedule?.repeats?.value ?? "0")
		})

		handle.frame("ran")
	}, [column, handle, plugHandle.plug])

	if (!column) return null

	return (
		<Frame
			index={index}
			className="z-[2]"
			icon={<Eye size={18} className="opacity-40" />}
			label="Preview"
			visible={(isFrame && session && session.user.anonymous === false) || false}
			hasOverlay={true}
			handleBack={column.schedule ? () => handle.frame("schedule") : undefined}
		>
			<div className="flex flex-col">
				{actions && actions.length > 0 ? (
					<>
						<ActionPreview index={index} item={item} />
						{!isReady && (
							<p className="text-sm font-medium text-grayscale-300 text-center py-2">
								Some actions have missing required values.
							</p>
						)}
					</>
				) : (
					<p className="text-sm font-medium text-grayscale-300 text-center py-4">
						No actions configured yet. Add some actions to run this Plug.
					</p>
				)}

				<div className="mb-2 mt-4 flex flex-row items-center gap-4">
					<p className="font-bold opacity-40">Transaction</p>
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

				{column.schedule && (
					<>
						<div className="mb-2 mt-4 flex flex-row items-center gap-4">
							<p className="font-bold opacity-40">Schedule</p>
							<div className="h-[2px] w-full bg-grayscale-100" />
						</div>

						<p className="flex flex-row justify-between font-bold">
							<span className="flex w-full flex-row items-center gap-4">
								<Calendar size={18} className="opacity-20" />
								<span className="opacity-40">Frequency</span>
							</span>{" "}
							{column.schedule.repeats.label}
						</p>

						{column.schedule.date && column.schedule.date.from instanceof Date && (
							<p className="flex flex-row justify-between font-bold">
								<span className="flex w-full flex-row items-center gap-4">
									<Play size={18} className="opacity-20" />
									<span className="opacity-40">Start At</span>
								</span>{" "}
								<Counter count={column.schedule.date.from.toLocaleDateString()} />
							</p>
						)}

						{column.schedule.date && column.schedule.date.to instanceof Date && (
							<p className="flex flex-row justify-between font-bold">
								<span className="flex w-full flex-row items-center gap-4">
									<Pause size={18} className="opacity-20" />
									<span className="opacity-40">Stop At</span>
								</span>{" "}
								<Counter count={column.schedule.date.to.toLocaleDateString()} />
							</p>
						)}
					</>
				)}

				<Button
					variant={isReady ? "primary" : "primaryDisabled"}
					className="mt-4 w-full py-4"
					onClick={handleRun}
					disabled={!isReady}
				>
					{isReady 
						? "Run" 
						: actions?.length === 0 
							? "No Actions Added" 
							: "Required Inputs Incomplete"}
				</Button>
			</div>
		</Frame>
	)
}
