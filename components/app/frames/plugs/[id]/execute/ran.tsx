import Image from "next/image"

import { CheckCircle } from "lucide-react"

import { Button, Frame } from "@/components"
import { useFrame, usePlugs } from "@/contexts"
import { formatTitle } from "@/lib"

export const RanFrame = () => {
	const { frameVisible, handleFrameVisible } = useFrame()
	const { plug, chains } = usePlugs()

	const isFrame = frameVisible ? frameVisible.split("-")[0] === "ran" : false

	const prevFrame = frameVisible ? frameVisible.split("-")[1] : undefined

	const label = frameVisible
		? prevFrame === "schedule"
			? "Intent Signed"
			: "Transaction Ran"
		: ""

	if (!plug) return null

	return (
		<Frame
			className="z-[2]"
			icon={<CheckCircle size={18} />}
			label={label}
			visible={isFrame}
		>
			<div className="flex flex-col gap-2">
				<p>
					<span className="opacity-60">Your</span>
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
						{prevFrame === "schedule"
							? "intent successfully signed. When all the constraints of your Plug are met, your transaction will be automatically submitted."
							: "transaction was successfully submit. Your balances have updated and you can now resume editing your Plug."}
					</span>
				</p>

				<p className="mt-4 flex font-bold">
					<span className="mr-auto opacity-60">Status</span>
					<span className="flex flex-row gap-2">
						<span>
							<div className="rounded-full bg-plug-green/10 p-1">
								<CheckCircle
									className="text-plug-green"
									size={16}
								/>
							</div>
						</span>
						Success
					</span>
				</p>

				<p className="flex font-bold">
					<span className="mr-auto opacity-60">
						{prevFrame === "schedule" ? "Signed For" : "Ran on"}
					</span>

					{chains.map(chain => (
						<Image
							key={chain}
							className="ml-[-20px] h-6 w-6"
							src={`/blockchain/${chain}.png`}
							alt={formatTitle(chain)}
							width={32}
							height={32}
						/>
					))}
				</p>

				{prevFrame !== "schedule" && (
					<p className="flex font-bold">
						<span className="mr-auto opacity-60">Fee</span>
						<div className="flex flex-row gap-2">
							<span className="opacity-40">0.00135 ETH</span>
							<span>$6.11</span>
						</div>
					</p>
				)}

				<Button
					variant="secondary"
					className="mt-4 w-full"
					onClick={() => handleFrameVisible("chain")}
				>
					View on Explorer
				</Button>
			</div>
		</Frame>
	)
}
