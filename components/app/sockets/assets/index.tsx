import { useState } from "react"

import { CircleDollarSign, ImageIcon } from "lucide-react"

import {
	Button,
	Footer,
	Header,
	SocketTokenList,
	TransferFrame
} from "@/components"
import { useFrame } from "@/contexts"

import { SocketCollectibleGrid } from "../collectibles"

export const SocketAssets = () => {
	const { handleFrameVisible } = useFrame()

	const [tokensExpanded, setTokensExpanded] = useState(false)
	const [collectiblesExpanded, setCollectiblesExpanded] = useState(false)

	return (
		<>
			<Header
				size="md"
				icon={<CircleDollarSign size={14} className="opacity-40" />}
				label="Tokens"
				nextLabel="See all"
				nextOnClick={() => setTokensExpanded(true)}
			/>
			<SocketTokenList expanded={tokensExpanded} />

			<Header
				size="md"
				icon={<ImageIcon size={14} className="opacity-40" />}
				label="Collectibles"
				nextLabel="See all"
				nextOnClick={() => setCollectiblesExpanded(true)}
			/>
			<SocketCollectibleGrid expanded={collectiblesExpanded} />

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
