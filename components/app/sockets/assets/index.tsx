import { FC, HTMLAttributes, useState } from "react"

import { CircleDollarSign, ImageIcon } from "lucide-react"

import {
	Header,
	SocketCollectionList,
	SocketPositionList,
	SocketTokenList
} from "@/components"
import { useBalances } from "@/contexts"

export const SocketAssets: FC<
	HTMLAttributes<HTMLDivElement> & { id: string }
> = ({ id, ...props }) => {
	const { positions } = useBalances()
	const { tokens } = positions

	const [tokensExpanded, setTokensExpanded] = useState(false)
	const [positionsExpanded, setPositionsExpanded] = useState(false)

	return (
		<div {...props}>
			<Header
				size="sm"
				icon={<CircleDollarSign size={14} className="opacity-40" />}
				label="Tokens"
				nextLabel={
					tokens.length < 5
						? undefined
						: tokensExpanded
							? "Collapse"
							: "See All"
				}
				nextOnClick={() => setTokensExpanded(!tokensExpanded)}
			/>
			<SocketTokenList id={id} expanded={tokensExpanded} />

			<Header
				size="sm"
				icon={<CircleDollarSign size={14} className="opacity-40" />}
				label="Positions"
				nextLabel={
					positions.protocols.length < 3
						? undefined
						: positionsExpanded
							? "Collapse"
							: "See All"
				}
				nextOnClick={() => setPositionsExpanded(!positionsExpanded)}
			/>
			<SocketPositionList id={id} expanded={positionsExpanded} />

			<Header
				size="sm"
				icon={<ImageIcon size={14} className="opacity-40" />}
				label="Collectibles"
			/>
			<SocketCollectionList id={id} />
		</div>
	)
}
