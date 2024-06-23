export const actionCategories: Record<
	string,
	{
		image: string
		gradientFrom: string
		gradientTo: string
		chains: Array<"ethereum" | "optimism" | "base" | "arbitrum">
	}
> = {
	plug: {
		image: "/protocols/plug.png",
		gradientFrom: "#00E100",
		gradientTo: "#A3F700",
		chains: ["ethereum", "optimism", "base", "arbitrum"]
	},
	nouns: {
		image: "/protocols/nouns.png",
		gradientFrom: "#E9C80B",
		gradientTo: "#FFE02C",
		chains: ["ethereum"]
	},
	fraxlend: {
		image: "/protocols/frax-lend.png",
		gradientFrom: "#000000",
		gradientTo: "#323232",
		chains: ["ethereum"]
	},
	aave: {
		image: "/protocols/aave.png",
		gradientFrom: "#33B7C5",
		gradientTo: "#B254A0",
		chains: ["ethereum"]
	},
	uniswap: {
		image: "/protocols/uniswap.png",
		gradientFrom: "#FF007A",
		gradientTo: "#FF9BCB",
		chains: ["ethereum", "optimism", "base"]
	}
}

export const abis = {
	plug: {
		baseFee:
			"function encode(uint8 $lessThanOrGreaterThan, uint256 $threshold) public pure returns (bytes memory $data)",
		blockNumber:
			"function encode(uint8 $lessThanOrGreaterThan, uint256 $block) public pure returns (bytes memory $data)",
		timestamp:
			"function encode(uint8 $lessThanOrGreaterThan, uint256 $datetime) public pure returns (bytes memory $data)",
		tokenBalance:
			"function encode(address $holder, address $asset, uint8 $lessThanOrGreaterThan, uint256 $amount) public pure returns (bytes memory $data)",
		rateLimit: "function encode(uint128 $max, uint32 $frequency)",
		limitedCalls:
			"function encode(uint256 $count) public pure returns (bytes memory $terms)",
		cooldown:
			"function encode(uint256 $duration) public pure returns (bytes memory $terms)"
	},
	nouns: {
		bid: "function encode(uint256 $bid) public pure returns (bytes memory $live)",
		isTokenId:
			"function encode(uint256 $id) public pure returns (bytes memory $data)",
		hasTrait:
			"function encode(bytes32 $traitType, bytes32 $trait) public pure returns (bytes memory)"
	},
	fraxlend: {
		rate: "function encode(uint8 $borrowOrLend, address $pool, uint8 $lessThanOrGreaterThan, uint256 $rate) public pure returns (bytes memory $data)",
		utilizationRate:
			"function encode(address $pool, uint8 $lessThanOrGreaterThan, uint256 $rate) public pure returns (bytes memory $data)",
		health: "function encode(address $pool, uint8 $lessThanOrGreaterThan, uint256 $health) public pure returns (bytes memory $data)",
		addCollateral:
			"function encode(address $pool, uint256 $collateral) public pure returns (bytes memory $data)",
		borrow: "function encode(uint256 $amount, address $pool, uint256 $collateral) public pure returns (bytes memory $data)",
		repay: "function encode(address $pool, uint256 $amount) public pure returns (bytes memory $data)",
		closePosition:
			"function encode(address $pool) public pure returns (bytes memory $data)",
		lendFrax:
			"function encode(uint256 $amount, address $pool) public pure returns (bytes memory $data)",
		withdrawFrax:
			"function encode(uint256 $amount, address $pool) public pure returns (bytes memory $data)"
	},
	aave: {
		borrowRate:
			"function encode(address $pool, uint8 $lessThanOrGreaterThan, uint256 $rate) public pure returns (bytes memory $data)",
		depositRate:
			"function encode(address $pool, uint8 $lessThanOrGreaterThan, uint256 $rate) public pure returns (bytes memory $data)",
		loanHealth:
			"function encode(address $pool, uint8 $lessThanOrGreaterThan, uint256 $health) public pure returns (bytes memory $data)",
		rewardClaim:
			"function encode(address $pool, uint256 $amount) public pure returns (bytes memory $data)"
	},
	uniswap: {
		swapExactETH:
			"function encode(uint256 $amount, uint256 $amount, address $token) public pure returns (bytes memory $data)",
		swapExactTokens:
			"function encode(uint256 $amount, address $token, uint256 $amount) public pure returns (bytes memory $data)",
		swapExactTokensForTokens:
			"function encode(uint256 $amount, address $token, uint256 $amount, address $token) public pure returns (bytes memory $data)",
		swapTokensForExactETH:
			"function encode(uint256 $amount, address $token, uint256 $amount) public pure returns (bytes memory $data)"
	}
}
