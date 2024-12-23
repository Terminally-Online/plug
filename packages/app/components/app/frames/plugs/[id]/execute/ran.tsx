import { FC } from "react"

import { Calendar, CheckCircle, CircleDollarSign, Pause, Play, Waypoints } from "lucide-react"

import { Counter, Frame, Image } from "@/components"
import { chains } from "@/lib"
import { useColumnStore, usePlugData } from "@/state"

export const RanFrame: FC<{ index: number; item: string }> = ({ index, item }) => {
	const { column, isFrame } = useColumnStore(index, "ran")
	const { plug } = usePlugData(item)

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
					<span className="opacity-40">The run of</span>
					<span
						className="rounded-lg bg-gradient-to-tr px-2 py-1 font-bold"
						style={{
							background: `linear-gradient(to right, rgba(0,239,54,0.1), rgba(147,223,0,0.1))`,
							color: `#00E100`
						}}
					>
						{plug.name}
					</span>{" "}
					<span className="opacity-40">
						was successfully scheduled. Your intent will be regularly simulated and automatically executed
						when as soon as it is ready.
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
			</div>
		</Frame>
	)
}
