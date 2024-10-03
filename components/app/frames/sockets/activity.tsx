import { FC } from "react"

import { Counter, Image } from "@/components"
import { useColumns } from "@/state"

import { getStatusIcon } from "../../sockets/activity/activity-item"
import { Frame } from "../base"

export const ActivityFrame: FC<{
	index: number
	activityIndex: number
	activity: { name: string; status: string }
}> = ({ index, activityIndex, activity }) => {
	const { isFrame } = useColumns(index, `${index}-${activityIndex}-activity`)

	return (
		<Frame
			index={index}
			icon={<div className="relative h-10 w-10">{getStatusIcon(activity.status)}</div>}
			label={activity.name}
			visible={isFrame}
			hasOverlay={true}
			hasChildrenPadding={false}
		>
			<div className="relative flex w-full flex-col gap-2 px-6 pb-4 text-left">
				<div className="flex flex-col gap-2 text-black text-opacity-40">
					<div className="flex items-center gap-4">
						<Image src="/protocols/aave.png" alt="Aave" width={32} height={32} className="h-6 w-6" />
						<p className="font-bold">
							<span
								style={{
									background: `linear-gradient(to right, rgba(0,239,54,0.1), rgba(147,223,0,0.1))`,
									color: `#00EF35`
								}}
								className="rounded px-2 py-1 font-bold"
							>
								lend{" "}
							</span>
							rate for{" "}
							<span
								style={{
									background: `linear-gradient(to right, rgba(0,239,54,0.1), rgba(147,223,0,0.1))`,
									color: `#00EF35`
								}}
								className="rounded px-2 py-1 font-bold"
							>
								ETH
							</span>{" "}
							is{" "}
							<span
								style={{
									background: `linear-gradient(to right, rgba(0,239,54,0.1), rgba(147,223,0,0.1))`,
									color: `#00EF35`
								}}
								className="rounded px-2 py-1 font-bold"
							>
								greater than
							</span>{" "}
							<span
								style={{
									background: `linear-gradient(to right, rgba(0,239,54,0.1), rgba(147,223,0,0.1))`,
									color: `#00EF35`
								}}
								className="rounded px-2 py-1 font-bold"
							>
								12%
							</span>
						</p>
					</div>
					<div className="flex items-center gap-4">
						<Image src="/protocols/uniswap.png" alt="Uniswap" width={32} height={32} className="h-6 w-6" />
						<p className="font-bold">
							Swap{" "}
							<span
								style={{
									background: `linear-gradient(to right, rgba(0,239,54,0.1), rgba(147,223,0,0.1))`,
									color: `#00EF35`
								}}
								className="rounded px-2 py-1 font-bold"
							>
								10,000
							</span>{" "}
							<span
								style={{
									background: `linear-gradient(to right, rgba(0,239,54,0.1), rgba(147,223,0,0.1))`,
									color: `#00EF35`
								}}
								className="rounded px-2 py-1 font-bold"
							>
								USDC
							</span>{" "}
							for{" "}
							<span
								style={{
									background: `linear-gradient(to right, rgba(0,239,54,0.1), rgba(147,223,0,0.1))`,
									color: `#00EF35`
								}}
								className="rounded px-2 py-1 font-bold"
							>
								2.5
							</span>{" "}
							<span
								style={{
									background: `linear-gradient(to right, rgba(0,239,54,0.1), rgba(147,223,0,0.1))`,
									color: `#00EF35`
								}}
								className="rounded px-2 py-1 font-bold"
							>
								ETH
							</span>
						</p>
					</div>
					<div className="flex items-center gap-4">
						<Image src="/protocols/aave.png" alt="Aave" width={32} height={32} className="h-6 w-6" />
						<p className="font-bold">
							Increase collateral in{" "}
							<span
								style={{
									background: `linear-gradient(to right, rgba(0,239,54,0.1), rgba(147,223,0,0.1))`,
									color: `#00EF35`
								}}
								className="rounded px-2 py-1 font-bold"
							>
								ETH
							</span>{" "}
							with{" "}
							<span
								style={{
									background: `linear-gradient(to right, rgba(0,239,54,0.1), rgba(147,223,0,0.1))`,
									color: `#00EF35`
								}}
								className="rounded px-2 py-1 font-bold"
							>
								2.5
							</span>{" "}
							<span
								style={{
									background: `linear-gradient(to right, rgba(0,239,54,0.1), rgba(147,223,0,0.1))`,
									color: `#00EF35`
								}}
								className="rounded px-2 py-1 font-bold"
							>
								ETH
							</span>
						</p>
					</div>
				</div>

				<div className="flex flex-col gap-2">
					<p className="mt-4 flex font-bold">
						<span className="mr-auto opacity-40">Run On</span>

						<Image
							className="ml-[-20px] h-6 w-6"
							src={`/blockchain/ethereum.png`}
							alt={"Ethereum"}
							width={24}
							height={24}
						/>
					</p>
					<p className="flex">
						<span className="mr-auto font-bold opacity-40">Total</span>
						<div className="flex flex-row gap-2">
							<span className="flex flex-row items-center gap-2 opacity-60">
								<Counter count={isFrame ? 0.00135 : 0} decimals={5} /> ETH
							</span>
							<span className="flex flex-row font-bold">
								$
								<Counter count={isFrame ? 6.11 : 0} />
							</span>
						</div>
					</p>
				</div>
			</div>
		</Frame>
	)
}
