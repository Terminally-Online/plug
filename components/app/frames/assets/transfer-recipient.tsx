import { FC } from "react"

import { RouterOutputs } from "@/server/client"

import { Frame, TokenImage } from "@/components"
import { getChainId } from "@/lib"
import { useColumns } from "@/state"

export const TransferRecipientFrame: FC<{
	index: number
	token: NonNullable<RouterOutputs["socket"]["balances"]["positions"]>["tokens"][number]
	color: string
	textColor: string
}> = ({ index, token, color, textColor }) => {
	const { column, frame } = useColumns(index, `${token?.symbol}-transfer-recipient`)

	if (!token || !column) return null

	return (
		<Frame
			index={index}
			icon={
				<div className="relative h-8 w-10">
					<TokenImage
						logo={
							token?.icon ||
							`https://token-icons.llamao.fi/icons/tokens/${getChainId(token.implementations[0].chain)}/${token.implementations[0].contract}?h=240&w=240`
						}
						symbol={token.symbol}
						size="sm"
					/>
				</div>
			}
			label="Transfer Amount"
			visible={column.frame === `${token.symbol}-transfer-send`}
			handleBack={() => frame(`${token.symbol}-token`)}
			hasChildrenPadding={false}
			hasOverlay
		></Frame>
	)
}
