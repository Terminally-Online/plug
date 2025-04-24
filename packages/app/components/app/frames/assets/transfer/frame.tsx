import { FC } from "react"

import { TransferAmountFrame } from "@/components/app/frames/assets/transfer/amount/frame"
import { TransferNFTFrame } from "@/components/app/frames/assets/transfer/nft/frame"
import { TransferRecipientFrame } from "@/components/app/frames/assets/transfer/recipient/frame"
import { ZerionPosition } from "@/lib"
import { RouterOutputs } from "@/server/client"

type TransferFrameProps = {
	index: number
	token: ZerionPosition
	// collectible?: NonNullable<RouterOutputs["socket"]["balances"]["collectibles"]>[number]["collectibles"][number]
	// collection?: NonNullable<RouterOutputs["socket"]["balances"]["collectibles"]>[number]
	color: string
	textColor: string
}

export const TransferFrame: FC<TransferFrameProps> = ({ index, token, ...colors }) => (
	<>
		{token && <TransferAmountFrame index={index} token={token} {...colors} />}
		{/* {collectible && collection && (
				<TransferNFTFrame index={index} collectible={collectible} collection={collection} {...colors} />
			)} */}
	</>
)
