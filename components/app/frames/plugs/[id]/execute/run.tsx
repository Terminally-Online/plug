import Image from "next/image"

import { Eye } from "lucide-react"

import { ActionPreview, Button, Frame } from "@/components"
import { useFrame, usePlugs } from "@/contexts"

export const RunFrame = () => {
	const { id, isFrame, prevFrame, handleFrame } = useFrame({
		id: "global",
		key: "run",
		seperator: "-"
	})
	const { chains, chainsAvailable } = usePlugs()

	const handleBack =
		prevFrame !== "schedule"
			? chainsAvailable.length === 1
				? undefined
				: () => handleFrame(`chain-${prevFrame}`)
			: () => handleFrame(`schedule`)

	return (
		<Frame
			id={id}
			className="z-[2]"
			handleBack={handleBack}
			icon={<Eye size={18} />}
			label={
				prevFrame === "schedule"
					? "Intent Preview"
					: "Transaction Preview"
			}
			visible={isFrame}
			hasOverlay={true}
		>
			<div className="flex flex-col gap-2">
				<p className="font-bold opacity-60">Actions</p>
				<ActionPreview />

				<p className="flex font-bold">
					<span className="mr-auto opacity-60">Run On</span>
					{chains.map(chain => (
						<Image
							key={chain}
							className="ml-[-20px] h-6 w-6"
							src={`/blockchain/${chain}.png`}
							alt={chain}
							width={24}
							height={24}
						/>
					))}
				</p>

				<p className="flex font-bold">
					<span className="mr-auto opacity-60">Fee</span>
					<span className="flex flex-row gap-2">
						<span className="opacity-40">0.0011 ETH</span>
						<span>$4.19</span>
					</span>
				</p>

				<Button
					className="mt-4 w-full"
					onClick={() => handleFrame(`running-${prevFrame}`)}
				>
					{prevFrame === "schedule"
						? "Sign Intent"
						: "Submit Transaction"}
				</Button>
			</div>
		</Frame>
	)
}
