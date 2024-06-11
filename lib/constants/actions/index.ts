import aave from "./aave"
import fraxlend from "./fraxlend"
import nouns from "./nouns"
import plug from "./plug"
import uniswap from "./uniswap"

export const actions = {
	plug,
	nouns,
	fraxlend,
	aave,
	uniswap
} as const

export type Action<
	TCategory extends keyof typeof actions,
	TAction extends keyof (typeof actions)[TCategory]
> = (typeof actions)[TCategory][TAction]

export type Actions = Readonly<{
	[KCategory in keyof typeof actions]: {
		[KAction in keyof (typeof actions)[KCategory]]: Action<
			KCategory,
			KAction
		>
	}
}>
