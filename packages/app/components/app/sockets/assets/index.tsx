import { FC, HTMLAttributes } from "react"

import { CircleDollarSign, ImageIcon } from "lucide-react"

import { Header } from "@/components/app/layout/header"
import { SocketCollectionList } from "@/components/app/sockets/collectibles/collection-list"
import { SocketPositionList } from "@/components/app/sockets/position/position-list"
import { SocketTokenList } from "@/components/app/sockets/tokens/token-list"
import { Callout } from "@/components/app/utils/callout"
import { cn } from "@/lib"
import { useSocket } from "@/state/authentication"
import { useHoldings } from "@/state/positions"

export const SocketAssets: FC<
	HTMLAttributes<HTMLDivElement> & {
		index?: number
		address?: string
		hasTokens?: boolean
		hasPositions?: boolean
		hasCollectibles?: boolean
	}
> = ({
	index = -1,
	address,
	hasTokens = false,
	hasPositions = false,
	hasCollectibles = false,
	className,
	...props
}) => {
	const { isAnonymous } = useSocket()
	const { collectibles, tokens, protocols } = useHoldings(address)

	return (
		<div className={cn("flex flex-col gap-2", className)} {...props}>
			<Callout.Anonymous index={index} viewing="assets" />
			<Callout.EmptyAssets
				index={index}
				isEmpty={[collectibles, tokens, protocols].every(basket => basket.length === 0)}
			/>
			{isAnonymous === false && (
				<>
					{hasTokens && tokens && tokens.length > 0 && (
						<SocketTokenList className="h-max" index={index} expanded={true} isColumn={false} />
					)}

					{hasPositions && protocols && protocols.length > 0 && (
						<>
							<Header
								size="sm"
								icon={<CircleDollarSign size={14} className="opacity-40" />}
								label="Positions"
							/>
							<SocketPositionList index={index} expanded={true} isColumn={false} />
						</>
					)}

					{hasCollectibles && collectibles && collectibles.length > 0 && (
						<>
							<Header
								size="sm"
								icon={<ImageIcon size={14} className="opacity-40" />}
								label="Collectibles"
							/>
							<SocketCollectionList index={index} expanded={true} isColumn={false} />
						</>
					)}
				</>
			)}
		</div>
	)
}
