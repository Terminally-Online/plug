import { FC, HTMLAttributes, useMemo } from "react"

import { Callout } from "@/components/app/utils/callout"
import { cn } from "@/lib"
import { useSocket } from "@/state/authentication"
import { api } from "@/server/client"
import { RouterOutputs } from "@/server/client"
import { SocketTokenList } from "../tokens/token-list"
import { Header } from "../../layout/header"
import { CircleDollarSign, ImageIcon } from "lucide-react"
import { SocketPositionList } from "../position/position-list"
import { SocketCollectionList } from "../collectibles/collection-list"

type SocketAssetsProps = HTMLAttributes<HTMLDivElement> & {
	index?: number
	address?: string
	hasTokens?: boolean
	hasPositions?: boolean
	hasCollectibles?: boolean
}

type Positions = RouterOutputs["service"]["zerion"]["wallet"]["positions"]["data"]

/**
 * Group positions by their position_type attribute.
 * @param positions Array of position objects from Zerion API
 * @returns Object with position types as keys and arrays of positions as values
 */
const groupByPositionType = (positions: Positions = []) => {
	return positions.reduce<Record<string, Positions>>((groups, position) => {
		const type = position.attributes.position_type;
		if (!groups[type]) {
			groups[type] = [];
		}
		groups[type].push(position);
		return groups;
	}, {});
};

export const SocketAssets: FC<SocketAssetsProps> = ({
	index = -1,
	address,
	hasTokens = false,
	hasPositions = false,
	hasCollectibles = false,
	className,
	...props
}) => {
	const { isAnonymous } = useSocket()

	return (
		<div className={cn("flex flex-col gap-2", className)} {...props}>
			<Callout.Anonymous index={index} viewing="assets" />

			{/*
			<Callout.EmptyAssets
				index={index}
				isEmpty={[tokens].every(basket => basket.length === 0)}
			/>
			*/}

			{isAnonymous === false && (
				<>
					<SocketTokenList
						className="h-max"
						index={index}
						address={address}
						expanded={true}
						isColumn={false}
					/>

					{/*
					{hasPositions && protocols && protocols.length > 0 && (
						<>
							<Header
								size="sm"
								icon={<CircleDollarSign size={14} className="opacity-40" />}
								label="Positions"
							/SocketPositionList
								index={index}
								expanded={true}
								isColumn={false}
								columnProtocols={protocols}
							/>
						</>
					)}

					{hasCollectibles && collectibles && collectibles.length > 0 && (
						<>
							<Header
								size="sm"
								icon={<ImageIcon size={14} className="opacity-40" />}
								label="Collectibles"
							/>
							<SocketCollectionList
								index={index}
								expanded={true}
								isColumn={false}
								columnCollectibles={collectibles}
							/>
						</>
					)}
					*/}
				</>
			)}
		</div>
	)
}
