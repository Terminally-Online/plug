import Image from "next/image"

import BlockiesSvg from "blockies-react-svg"
import { Eye } from "lucide-react"

import { ActionPreview, Button, Frame } from "@/components"
import { useFrame, usePlugs, useSockets } from "@/contexts"

export const RunFrame = () => {
	const { frameVisible, handleFrameVisible } = useFrame()
	const { socket, sockets } = useSockets()
	const { chains, chainsAvailable } = usePlugs()

	const isFrame =
		frameVisible !== undefined &&
		(frameVisible === "run" || frameVisible.split("-")[0] === "run")

	const prevFrameSuffix =
		frameVisible !== undefined && frameVisible.split("-")[1] === "schedule"
			? "schedule"
			: "run"

	const handleBack =
		prevFrameSuffix !== "schedule"
			? chainsAvailable.length === 1
				? sockets && sockets.length === 1
					? undefined
					: () => handleFrameVisible(`socket-${prevFrameSuffix}`)
				: () => handleFrameVisible(`chain-${prevFrameSuffix}`)
			: () => handleFrameVisible(`schedule`)

	return (
		<Frame
			className="z-[2]"
			handleBack={handleBack}
			icon={<Eye size={18} />}
			label={
				prevFrameSuffix === "schedule"
					? "Intent Preview"
					: "Transaction Preview"
			}
			visible={isFrame}
			hasOverlay={true}
		>
			<div className="flex flex-col gap-2">
				<p className="font-bold opacity-60">Actions</p>
				<ActionPreview />

				{socket && (
					<p className="mt-4 flex font-bold">
						<span className="mr-auto opacity-60">Use Socket</span>
						<div className="flex flex-row items-center gap-2">
							<BlockiesSvg
								address={socket.socketAddress}
								className="h-5 w-5 rounded-md"
							/>
							<p className="mr-auto font-bold">{socket.name}</p>
						</div>
					</p>
				)}

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
					onClick={() =>
						handleFrameVisible(`running-${prevFrameSuffix}`)
					}
				>
					{prevFrameSuffix === "schedule"
						? "Sign Intent"
						: "Submit Transaction"}
				</Button>
			</div>
		</Frame>
	)
}
