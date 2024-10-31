import { FC, useEffect } from "react"

import { CheckCircle, CircleDollarSign, Waypoints } from "lucide-react"

import { Frame, Image } from "@/components"
import { usePlugs } from "@/contexts"
import { chains } from "@/lib"
import { useColumns } from "@/state"

export const RanFrame: FC<{ index: number; item: string }> = ({ index, item }) => {
	const { isFrame, schedule } = useColumns(index, "ran")
	const { plug } = usePlugs(item)

	// NOTE: If a user has made it here the transaction / scheduling process has concluded.
	useEffect(() => {
		if (isFrame) schedule()
	}, [isFrame, schedule])

	if (!plug) return null

	return (
		<Frame
			index={index}
			className="z-[2]"
			icon={<CheckCircle size={18} className="opacity-40" />}
			label="Execution Queued"
			visible={isFrame}
		>
			<div className="flex flex-col">
				<p className="mb-2">
					<span className="opacity-60">The execution of</span>
					<span
						className="rounded-lg bg-gradient-to-tr px-2 py-1 font-bold"
						style={{
							background: `linear-gradient(to right, rgba(0,239,54,0.1), rgba(147,223,0,0.1))`,
							color: `#00E100`
						}}
					>
						{plug.name}
					</span>{" "}
					<span className="opacity-60">
						was successfully scheduled. The transaction will be automatically executed when it is ready.
					</span>
				</p>

				<div className="my-2 flex flex-row items-center gap-4">
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
			</div>
		</Frame>
	)
}
