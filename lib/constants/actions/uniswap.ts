import { ArrowLeftRightIcon, ArrowRightLeftIcon } from "lucide-react"
import { parseAbi, zeroAddress } from "viem"

import { abis } from "../abis"

export const uniswap = {
	swapExactETH: {
		address: zeroAddress,
		abi: abis.uniswap.swapExactETH,
		inputs: parseAbi([abis.uniswap.swapExactETH])[0]["inputs"],
		sentence: "Swap {0} ETH for at least {1} {2}.",
		info: "Swap a specific amount of ETH for a specific amount of the selected token. This can be used as a limit order.",
		icon: ArrowRightLeftIcon,
		primary: true
	},
	swapExactTokens: {
		address: zeroAddress,
		abi: abis.uniswap.swapExactTokens,
		inputs: parseAbi([abis.uniswap.swapExactTokens])[0]["inputs"],
		sentence: "Swap {0} {1} for exactly {2} ETH.",
		info: "Swap a specific amount of the selected token for a specific amount of ETH. This can be used as a limit order.",
		icon: ArrowLeftRightIcon,
		primary: true
	},
	swapExactTokensForTokens: {
		address: zeroAddress,
		abi: abis.uniswap.swapExactTokensForTokens,
		inputs: parseAbi([abis.uniswap.swapExactTokensForTokens])[0]["inputs"],
		sentence: "Swap {0} {1} for exactly {2} {3}.",
		info: "Swap a specific amount of one token for a specific amount of another token. This can be used as a limit order.",
		icon: ArrowRightLeftIcon
	},
	swapTokensForExactETH: {
		address: zeroAddress,
		abi: abis.uniswap.swapTokensForExactETH,
		inputs: parseAbi([abis.uniswap.swapTokensForExactETH])[0]["inputs"],
		sentence: "Swap {0} {1} for at least {2} ETH.",
		info: "Swap a specific amount of the selected token for a specific amount of ETH. This can be used as a limit order.",
		icon: ArrowLeftRightIcon
	}
}

export default uniswap
