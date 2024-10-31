import { FC } from "react"

import { useDebounce } from "@/lib"
import { RouterOutputs } from "@/server/client"

import { TransferAmountFrame } from "./transfer-amount"
import { TransferRecipientFrame } from "./transfer-recipient"
import { TransferNFTFrame } from "./transfer-nft" 


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

export const TransferFrame: FC<TransferFrameProps> = ({
	index,
	token,
	collectible,
	collection,
	color,
	textColor
}) => {
	const [recipient, debouncedRecipient, setRecipient] = useDebounce("")

	const handleRecipient = (recipient: string) => {
		setRecipient(recipient.trim())
	}

	// Determine if we're dealing with an NFT and what type
	const isNFT = Boolean(collectible && collection)
	const isERC721 = isNFT && collectible?.interface === "ERC721"
	const isERC1155 = isNFT && collectible?.interface === "ERC1155"

	return (
		<>
			<TransferRecipientFrame
				index={index}
				token={token}
				collectible={collectible}
				collection={collection}
				recipient={recipient}
				debouncedRecipient={debouncedRecipient}
				handleRecipient={handleRecipient}
			/>
			
			{/* For ERC20 tokens, use existing amount frame */}
			{token && (
				<TransferAmountFrame
					index={index}
					token={token}
					recipient={recipient}
					color={color}
					textColor={textColor}
				/>
			)}

			{/* For NFTs, use appropriate confirmation frame */}
			{isNFT && (
				<TransferNFTFrame
					index={index}
					collectible={collectible}
					collection={collection}
					recipient={recipient}
					color={color}
					textColor={textColor}
					isERC1155={isERC1155}
				/>
			)}
		</>
	)
}
