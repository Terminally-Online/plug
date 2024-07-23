import { useState } from "react"

import { CircleDollarSign, ImageIcon } from "lucide-react"

import {
	Button,
	Footer,
	Header,
	SocketCollectionList,
	SocketTokenList,
	TransferFrame
} from "@/components"
import { useFrame } from "@/contexts"

export const SocketAssets = () => {
	const { handleFrameVisible } = useFrame()

	const [tokensExpanded, setTokensExpanded] = useState(false)

	return (
		<>
			<Header
				size="md"
				icon={<CircleDollarSign size={14} className="opacity-40" />}
				label="Tokens"
				nextLabel={tokensExpanded ? "Collapse" : "See All"}
				nextOnClick={() => setTokensExpanded(!tokensExpanded)}
			/>
			<SocketTokenList expanded={tokensExpanded} />

			<Header
				size="md"
				icon={<ImageIcon size={14} className="opacity-40" />}
				label="Collectibles"
			/>
			<SocketCollectionList />

			<Footer>
				<Button
					className="w-full"
					onClick={() => handleFrameVisible("transfer")}
				>
					Transfer
				</Button>
			</Footer>

			<TransferFrame />
		</>
	)
}
