import { useState } from "react"

import { CircleDollarSign, ImageIcon } from "lucide-react"

import { Header, SocketTokenList, TransferFrame } from "@/components"

export const SocketTokens = () => {
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
			{/* TODO: Add collectibles into the API and build a list */}

			<TransferFrame />
		</>
	)
}
