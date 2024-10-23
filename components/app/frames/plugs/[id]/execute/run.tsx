import { FC } from "react"

import { Eye } from "lucide-react"

import { ActionPreview, Button, Frame, Image } from "@/components"
import { usePlugs } from "@/contexts"
import { useColumns } from "@/state"

export const RunFrame: FC<{ index: number; item: string }> = ({ index, item }) => {
	const { isFrame, frame } = useColumns(index, "run")
	const { chains } = usePlugs(item)

	const prevFrame = "NOT_IMPLEMENTED" as string

	const handleBack =
		prevFrame !== "schedule"
			? chains.length === 1
				? undefined
				: () => frame(`chain-${prevFrame}`)
			: () => frame(`schedule`)

	return (
		<Frame
			index={index}
			className="z-[2]"
			handleBack={handleBack}
			icon={<Eye size={18} />}
			label={prevFrame === "schedule" ? "Intent Preview" : "Transaction Preview"}
			visible={isFrame}
			hasOverlay={true}
		>
			<div className="flex flex-col gap-2">
				<p className="font-bold opacity-60">Actions</p>
				<ActionPreview index={index} item={item} />

				<p className="flex font-bold">
					<span className="mr-auto opacity-60">Run On</span>
					{/* {chains.map(chain => (
                        <Image
                            key={chain}
                            className="ml-[-20px] h-6 w-6"
                            src={`/blockchain/${chain}.png`}
                            alt={chain}
                            width={24}
                            height={24}
                        />
                    ))} */}
					<Image className="h-6 w-6" src="/blockchain/ethereum.png" alt="Ethereum" width={24} height={24} />
				</p>

				<p className="flex font-bold">
					<span className="mr-auto opacity-60">Fee</span>
					<span className="flex flex-row gap-2">
						<span className="opacity-40">0.0011 ETH</span>
						<span>$4.19</span>
					</span>
				</p>

				<Button className="mt-4 w-full" onClick={() => frame("running")}>
					{prevFrame === "schedule" ? "Sign Intent" : "Submit Transaction"}
				</Button>
			</div>
		</Frame>
	)
}
