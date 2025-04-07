import { FC } from "react"

import { TransferAmountFrame } from "@/components/app/frames/assets/transfer/amount/frame"
import { TransferNFTFrame } from "@/components/app/frames/assets/transfer/nft/frame"
import { TransferRecipientFrame } from "@/components/app/frames/assets/transfer/recipient/frame"
import { TransferSuccessFrame } from "@/components/app/frames/assets/transfer/success/frame"
import { RouterOutputs } from "@/server/client"

type TransferFrameProps = {
	index: number
	token?: NonNullable<RouterOutputs["socket"]["balances"]["positions"]>["tokens"][number]
	collectible?: NonNullable<RouterOutputs["socket"]["balances"]["collectibles"]>[number]["collectibles"][number]
	collection?: NonNullable<RouterOutputs["socket"]["balances"]["collectibles"]>[number]
	color: string
	textColor: string
}

export const TransferFrame: FC<TransferFrameProps> = ({ index, token, collectible, collection, ...colors }) => {
	const isNFT = Boolean(collectible && collection)

	return (
		<>
			<TransferRecipientFrame index={index} token={token} collectible={collectible} collection={collection} />

			{token && (
				<>
					<TransferAmountFrame index={index} token={token} {...colors} />
					<TransferSuccessFrame index={index} token={token} {...colors} />
				</>
			)}

			{isNFT && collectible && collection && (
				<TransferNFTFrame
					index={index}
					isERC1155={isNFT && collectible?.interface === "ERC1155"}
					collectible={collectible}
					collection={collection}
					{...colors}
				/>
			)}
		</>
	)
}
