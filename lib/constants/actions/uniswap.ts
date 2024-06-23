import { ArrowLeftRightIcon, ArrowRightLeftIcon } from "lucide-react"
import { parseAbi, zeroAddress } from "viem"

import { abis } from "../abis"

export const uniswap = {
	swapExactETH: {
		address: zeroAddress,
		abi: abis.uniswap.swapExactETH,
		inputs: parseAbi([abis.uniswap.swapExactETH])[0]["inputs"],
		options: [
			// No options needed here as the user will input values directly
		],
		// {0} is the amount of ETH to swap (user input)
		// {1} is the minimum amount of tokens to receive (user input)
		// {2} is the type of token (user will select from an imported list)
		sentence: "Swap {0} ETH for at least {1} {2}.",
		info: "Swap a specific amount of ETH for a specific amount of the selected token. This can be used as a limit order.",
		icon: ArrowRightLeftIcon,
		primary: true
	},
	swapExactTokens: {
		address: zeroAddress,
		abi: abis.uniswap.swapExactTokens,
		inputs: parseAbi([abis.uniswap.swapExactTokens])[0]["inputs"],
		options: [
			// No options needed here as the user will input values directly
		],
		// {0} is the amount of the selected token to swap (user input)
		// {1} is the type of token to swap (user will select from an imported list)
		// {2} is the exact amount of ETH to receive (user input)
		sentence: "Swap {0} {1} for exactly {2} ETH.",
		info: "Swap a specific amount of the selected token for a specific amount of ETH. This can be used as a limit order.",
		icon: ArrowLeftRightIcon,
		primary: true
	},
	swapExactTokensForTokens: {
		address: zeroAddress,
		abi: abis.uniswap.swapExactTokensForTokens,
		inputs: parseAbi([abis.uniswap.swapExactTokensForTokens])[0]["inputs"],
		options: [
			// No options needed here as the user will input values directly
		],
		// {0} is the amount of the source token to swap (user input)
		// {1} is the type of source token (user will select from an imported list)
		// {2} is the exact amount of the destination token to receive (user input)
		// {3} is the type of destination token (user will select from an imported list)
		sentence: "Swap {0} {1} for exactly {2} {3}.",
		info: "Swap a specific amount of one token for a specific amount of another token. This can be used as a limit order.",
		icon: ArrowRightLeftIcon
	},
	swapTokensForExactETH: {
		address: zeroAddress,
		abi: abis.uniswap.swapTokensForExactETH,
		inputs: parseAbi([abis.uniswap.swapTokensForExactETH])[0]["inputs"],
		options: [
			// No options needed here as the user will input values directly
		],
		// {0} is the amount of the selected token to swap (user input)
		// {1} is the type of token to swap (user will select from an imported list)
		// {2} is the minimum amount of ETH to receive (user input)
		sentence: "Swap {0} {1} for at least {2} ETH.",
		info: "Swap a specific amount of the selected token for a specific amount of ETH. This can be used as a limit order.",
		icon: ArrowLeftRightIcon
	}
}

export default uniswap
