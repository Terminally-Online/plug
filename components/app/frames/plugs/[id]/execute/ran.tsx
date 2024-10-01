import Image from "next/image"
import { FC } from "react"

import { CheckCircle } from "lucide-react"

import { Button, Frame } from "@/components"
import { usePlugs } from "@/contexts"
import { formatTitle } from "@/lib"
import { useColumns } from "@/state"

export const RanFrame: FC<{ index: number; item: string }> = ({ index, item }) => {
	const { isFrame, frame } = useColumns(index, "ran")
	const { plug, chains } = usePlugs(item)

	const prevFrame = "NOT_IMPLEMENTED" as string

	const label = isFrame ? (prevFrame === "schedule" ? "Intent Signed" : "Transaction Ran") : ""

	if (!plug) return null

	return (
		<Frame index={index} className="z-[2]" icon={<CheckCircle size={18} />} label={label} visible={isFrame}>
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
								<CheckCircle className="text-plug-green" size={16} />
							</div>
						</span>
						Success
					</span>
				</p>

				<p className="flex font-bold">
					<span className="mr-auto opacity-60">{prevFrame === "schedule" ? "Signed For" : "Ran on"}</span>

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

				<Button variant="secondary" className="mt-4 w-full" onClick={() => frame("chain")}>
					View on Explorer
				</Button>
			</div>
		</Frame>
	)
}
