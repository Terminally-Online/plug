import { FC, useState } from "react"

import { SwapAmountFrame } from "@/components/app/frames/assets/swap/amount/frame"
import { SwapTokenFrame } from "@/components/app/frames/assets/swap/token/frame"
import { RouterOutputs } from "@/server/client"

type Token =
	| NonNullable<RouterOutputs["socket"]["balances"]["positions"]>["tokens"][number]
	| NonNullable<RouterOutputs["solver"]["tokens"]["get"]>[number]

type SwapFrameProps = {
	index: number
	tokenOut: Token
	color: string
	textColor: string
}

export const SwapFrame: FC<SwapFrameProps> = ({ index, tokenOut }) => {
	const [tokenIn, setTokenIn] = useState<Token | undefined>(undefined)

	return (
		<>
			<SwapTokenFrame index={index} tokenOut={tokenOut} handleTokenIn={setTokenIn} />

			{tokenIn && <SwapAmountFrame index={index} tokenIn={tokenIn} tokenOut={tokenOut} />}
		</>
	)
}
