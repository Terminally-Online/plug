import { FC, useState } from "react"

import { RouterOutputs } from "@/server/client"

import { SwapAmountFrame } from "./swap.amount"
import { SwapConfirmFrame } from "./swap.confirm"
import { SwapTokenFrame } from "./swap.token"

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
			{tokenIn && <SwapConfirmFrame index={index} tokenIn={tokenIn} tokenOut={tokenOut} />}
		</>
	)
}
