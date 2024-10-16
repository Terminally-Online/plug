import { FC } from "react"

import { RouterOutputs } from "@/server/client"

import { TransferAmountFrame } from "./transfer-amount"
import { TransferRecipientFrame } from "./transfer-recipient"

export const TransferFrame: FC<{
	index: number
	token: NonNullable<RouterOutputs["socket"]["balances"]["positions"]>["tokens"][number]
	color: string
	textColor: string
}> = ({ index, token, color, textColor }) => {
	return (
		<>
			<TransferRecipientFrame index={index} token={token} color={color} textColor={textColor} />
			<TransferAmountFrame index={index} token={token} color={color} textColor={textColor} />
		</>
	)
}
