import { FC } from "react"

import { TransferAmountFrame } from "@/components/app/frames/assets/transfer/amount/frame"
import { TransferNFTFrame } from "@/components/app/frames/assets/transfer/nft/frame"
import { TransferRecipientFrame } from "@/components/app/frames/assets/transfer/recipient/frame"
import { ZerionPosition } from "@/lib"
import { RouterOutputs } from "@/server/client"

type TransferFrameProps = {
	index: number
	token?: ZerionPosition
	collectible?: NonNullable<RouterOutputs["service"]["zerion"]["nfts"]["detail"]["data"]>
	included?: NonNullable<RouterOutputs["service"]["zerion"]["nfts"]["detail"]["included"]>[number]
	color: string
	textColor: string
}

export const TransferFrame: FC<TransferFrameProps> = ({ index, token, collectible, included, ...colors }) => (
	<>
		<TransferRecipientFrame index={index} token={token} collectible={collectible} included={included} />
		{token && <TransferAmountFrame index={index} token={token} {...colors} />}
		{collectible && included && (
			<TransferNFTFrame
				index={index}
				collectible={collectible}
				included={included}
				{...colors}
			/>
		)}
	</>
)
