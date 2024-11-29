import { useSession } from "next-auth/react"
import { FC, useCallback } from "react"

import { Calendar, CircleDollarSign, Eye, Pause, Play, Waypoints } from "lucide-react"

import { ActionPreview, Button, Counter, Frame, Image } from "@/components"
import { chains } from "@/lib"
import { useColumnStore, usePlugStore } from "@/state"

export const RunFrame: FC<{
	index: number
	item: string
}> = ({ index, item }) => {
	const { data: session } = useSession()
	const { column, isFrame, handle: { frame } } = useColumnStore(index, "run")
	const { actions, handle: { plug: { queue }} } = usePlugStore(item)

	const isReady = false

	// TODO: Disabled this while working on implementing cord. Need to re-implement it.
	// const isReady = useMemo(
	// 	() =>
	// 		plug &&
	// 		actions &&
	// 		actions.length > 0 &&
	// 		actions.every(action => action.values.every(value => Boolean(value))),
	// 	[plug, actions]
	// )

	const handleRun = useCallback(() => {
		if (!column || !column.item) return

		queue({
			workflowId: column.item,
			startAt: column.schedule?.date?.from ?? new Date(),
			endAt: column.schedule?.date?.to ?? new Date(),
			frequency: parseInt(column.schedule?.repeats?.value ?? "0")
		})

		frame("ran")
	}, [column, queue, frame])

	if (!column) return null

	return (
		<Frame
			index={index}
			className="z-[2]"
			icon={<Eye size={18} className="opacity-40" />}
			label="Preview"
			visible={(isFrame && session && session.user.anonymous === false) || false}
			hasOverlay={true}
			handleBack={column.schedule ? () => frame("schedule") : undefined}
		>
			<div className="flex flex-col">
				{actions && actions.length > 0 ? (
					<ActionPreview index={index} item={item} />
				) : (
					<div className="flex rounded-lg border-[1px] border-plug-green/10 p-4 py-4 text-center font-bold text-black/40">
						<p className="mx-auto max-w-[380px]">
							No actions added and configured on this Plug yet. Add some actions to run and schedule it.
						</p>
					</div>
				)}

				<div className="mb-2 mt-4 flex flex-row items-center gap-4">
					<p className="font-bold opacity-40">Transaction</p>
					<div className="h-[2px] w-full bg-plug-green/10" />
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
							<div className="h-[2px] w-full bg-plug-green/10" />
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
					{isReady ? "Run" : actions?.length === 0 ? "No Actions Added" : "Required Inputs Incomplete"}
				</Button>
			</div>
		</Frame>
	)
}
