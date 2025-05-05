import { FC, HTMLAttributes } from "react"

import { Callout } from "@/components/app/utils/callout"
import { cn } from "@/lib"
import { useSocket } from "@/state/authentication"
import { SocketTokenList } from "../tokens/token-list"
import { SocketPositionList } from "../position/position-list"
import { SocketCollectionList } from "../collectibles/collection-list"

type SocketAssetsProps = HTMLAttributes<HTMLDivElement> & {
	index?: number
	address?: string
	hasTokens?: boolean
	hasPositions?: boolean
	hasCollectibles?: boolean
}

export const SocketAssets: FC<SocketAssetsProps> = ({
	index = -1,
	address,
	hasTokens = false,
	hasPositions = false,
	hasCollectibles = false,
	className,
	...props
}) => (
	<div className={cn("flex flex-col gap-2", className)} {...props}>
		<Callout.Anonymous index={index} viewing="assets" />

		{/* <Callout.EmptyAssets
			index={index}
			isEmpty={[tokens].every(basket => basket.length === 0)}
		/> */}

		<SocketTokenList
			className="h-max"
			index={index}
			address={address}
			isColumn={false}
			expanded
		/>

		<SocketCollectionList
			index={index}
			address={address}
			isColumn={false}
			expanded
		/>
	</div>
)

