import { FC } from "react"

import { RouterOutputs } from "@/server/client"

import { useDebounce } from "@/lib"

import { TransferAmountFrame } from "./transfer-amount"
import { TransferRecipientFrame } from "./transfer-recipient"

export const TransferFrame: FC<{
	index: number
	token: NonNullable<RouterOutputs["socket"]["balances"]["positions"]>["tokens"][number]
	color: string
	textColor: string
}> = ({ index, token, color, textColor }) => {
	const [recipient, debouncedRecipient, setRecipient] = useDebounce("")

	const handleRecipient = (recipient: string) => {
		setRecipient(recipient.trim())
	}

	return (
		<>
			<TransferRecipientFrame
				index={index}
				token={token}
				recipient={recipient}
				debouncedRecipient={debouncedRecipient}
				handleRecipient={handleRecipient}
			/>
			<TransferAmountFrame
				index={index}
				token={token}
				recipient={recipient}
				color={color}
				textColor={textColor}
			/>
		</>
	)
}
