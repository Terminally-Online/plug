import { FC, useState } from "react"

import { SwapAmountFrame } from "@/components/app/frames/assets/swap/amount/frame"
import { SwapTokenFrame } from "@/components/app/frames/assets/swap/token/frame"
import { ZerionFungible, ZerionPosition } from "@/lib"

type SwapFrameProps = {
	index: number
	tokenOut: ZerionPosition
	color: string
	textColor: string
}

export const SwapFrame: FC<SwapFrameProps> = ({ index, tokenOut }) => {
	const [tokenIn, setTokenIn] = useState<ZerionFungible | undefined>()

	return (
		<>
			<SwapTokenFrame index={index} tokenOut={tokenOut} handleTokenIn={setTokenIn} />

			{tokenIn && <SwapAmountFrame index={index} tokenIn={tokenIn} tokenOut={tokenOut} />}
		</>
	)
}
