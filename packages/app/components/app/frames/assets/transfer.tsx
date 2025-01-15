import { FC } from "react"

import { TransferAmountFrame } from "@/components/app/frames/assets/transfer-amount"
import { TransferNFTFrame } from "@/components/app/frames/assets/transfer-nft"
import { TransferRecipientFrame } from "@/components/app/frames/assets/transfer-recipient"
import { RouterOutputs } from "@/server/client"

type TransferFrameProps = {
	index: number
	token?: NonNullable<RouterOutputs["socket"]["balances"]["positions"]>["tokens"][number]
	collectible?: NonNullable<RouterOutputs["socket"]["balances"]["collectibles"]>[number]["collectibles"][number]
	collection?: NonNullable<RouterOutputs["socket"]["balances"]["collectibles"]>[number]
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
