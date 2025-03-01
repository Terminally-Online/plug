import { FC } from "react"

import { Calendar, CheckCircle, CircleDollarSign, Pause, Play, Waypoints } from "lucide-react"

import { Frame } from "@/components/app/frames/base"
import { ChainImage } from "@/components/app/sockets/chains/chain.image"
import { Button } from "@/components/shared/buttons/button"
import { Counter } from "@/components/shared/utils/counter"
import { columnByIndexAtom, isFrameAtom, useColumnActions } from "@/state/columns"
import { workflowByIdAtom } from "@/state/plugs"
import { useAtom, useAtomValue } from "jotai"

export const RanFrame: FC<{ index: number; item: string }> = ({ index, item }) => {
	const [column] = useAtom(columnByIndexAtom(index))
	const frameKey = "ran"
	const isFrame = useAtomValue(isFrameAtom)(column, frameKey)
	const { frame } = useColumnActions(index, frameKey)

	const plug = useAtomValue(workflowByIdAtom)(item)

	if (!plug || !column) return null

	return (
		<Frame
			index={index}
			className="z-[2]"
			icon={<CheckCircle size={18} className="opacity-40" />}
			label="Execution Queued"
			visible={isFrame}
		>
			<div className="flex flex-col">
				<p className="mb-2 font-bold">
					<span className="opacity-40">The run of</span>{" "}
					<span
						className="rounded-sm bg-gradient-to-tr px-2 font-bold"
						style={{
							background: `linear-gradient(to right, rgba(56, 88, 66, 0.1), rgba(210, 243, 138, 0.1))`,
							color: `#385842`
						}}
					>
						{plug.name}
					</span>{" "}
					<span className="opacity-40">
						was successfully scheduled. Your execution will be regularly simulated and automatically
						executed when as soon as it is ready.
					</span>
				</p>

				<div className="mb-2 mt-4 flex flex-row items-center gap-4">
					<p className="font-bold opacity-40">Transaction</p>
					<div className="h-[2px] w-full bg-plug-green/10" />
				</div>

				<p className="flex w-full flex-row items-center gap-4 font-bold">
					<Waypoints size={18} className="opacity-20" />
					<span className="mr-auto opacity-40">Chain</span>
					<span className="flex flex-row items-center gap-2">
						<ChainImage chainId={1} size="xs" />
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

				<Button className="mt-4 py-4" onClick={() => frame()}>
					Done
				</Button>
			</div>
		</Frame>
	)
}
