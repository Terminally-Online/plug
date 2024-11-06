import { FC } from "react"

import { RouterOutputs } from "@/server/client"

import { TransferAmountFrame } from "./transfer-amount"
import { TransferNFTFrame } from "./transfer-nft"
import { TransferRecipientFrame } from "./transfer-recipient"

type TokenType = NonNullable<RouterOutputs["socket"]["balances"]["positions"]>["tokens"][number]
type CollectibleType = NonNullable<RouterOutputs["socket"]["balances"]["collectibles"]>[number]["collectibles"][number]
type CollectionType = NonNullable<RouterOutputs["socket"]["balances"]["collectibles"]>[number]

interface TransferFrameProps {
	index: number
	token?: TokenType
	collectible?: CollectibleType
	collection?: CollectionType
	color: string
	textColor: string
}

export const TransferFrame: FC<TransferFrameProps> = ({ index, token, collectible, collection, color, textColor }) => {
	const isNFT = Boolean(collectible && collection)
	const isERC1155 = isNFT && collectible?.interface === "ERC1155"

	return (
		<>
			<TransferRecipientFrame index={index} token={token} collectible={collectible} collection={collection} />

			{token && <TransferAmountFrame index={index} token={token} color={color} textColor={textColor} />}

			{isNFT && collectible && collection && (
				<TransferNFTFrame
					index={index}
					collectible={collectible}
					collection={collection}
					color={color}
					textColor={textColor}
					isERC1155={isERC1155}
				/>
			)}
		</>
	)
}
