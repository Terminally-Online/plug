export const categories: Record<
	string,
	{
		image: string
		gradientFrom: string
		gradientTo: string
		chains: Array<"ethereum" | "optimism" | "base" | "arbitrum">
		tags: Array<
			"defi" | "nft" | "consumer" | "degen" | "mev" | "social" | "trading"
		>
	}
> = {
	plug: {
		image: "/protocols/plug.png",
		gradientFrom: "#00E100",
		gradientTo: "#A3F700",
		chains: ["ethereum", "optimism", "base", "arbitrum"],
		tags: ["consumer"]
	},
	nouns: {
		image: "/protocols/nouns.png",
		gradientFrom: "#E9C80B",
		gradientTo: "#FFE02C",
		chains: ["ethereum"],
		tags: ["nft", "consumer"]
	},
	fraxlend: {
		image: "/protocols/fraxlend.png",
		gradientFrom: "#000000",
		gradientTo: "#323232",
		chains: ["ethereum"],
		tags: ["defi"]
	},
	aave: {
		image: "/protocols/aave.png",
		gradientFrom: "#33B7C5",
		gradientTo: "#B254A0",
		chains: ["ethereum"],
		tags: ["defi"]
	},
	uniswap: {
		image: "/protocols/uniswap.png",
		gradientFrom: "#FF007A",
		gradientTo: "#FF9BCB",
		chains: ["ethereum", "optimism", "base"],
		tags: ["defi", "trading"]
	}
	// chainlink: {
	// 	image: "/protocols/chainlink.png",
	// 	gradientFrom: "#000000",
	// 	gradientTo: "#000000",
	// 	chains: ["ethereum", "optimism", "base"]
	// },
	// ,velodrome: {
	// 	image: "/protocols/velodrome.png",
	// 	gradientFrom: "#000000",
	// 	gradientTo: "#000000",
	// 	chains: ["optimism"]
	// }
	// aerodrome: {
	// 	image: "/protocols/aerodrome.png",
	// 	gradientFrom: "#000000",
	// 	gradientTo: "#000000",
	// 	chains: ["base"]
	// }
	// ,ethena: {
	// 	image: "/protocols/ethena.png",
	// 	gradientFrom: "#000000",
	// 	gradientTo: "#000000",
	// 	chains: ["ethereum", "optimism", "base"]
	// }
	// ,enjoy: {
	// 	image: "/protocols/enjoy.png",
	// 	gradientFrom: "#000000",
	// 	gradientTo: "#000000",
	// 	chains: ["base"]
	// }
	// ,ens: {
	// 	image: "/protocols/ens.png",
	// 	gradientFrom: "#000000",
	// 	gradientTo: "#000000",
	// 	chains: ["ethereum"]
	// }
	// pendle: {
	// 	image: "/protocols/pendle.png",
	// 	gradientFrom: "#000000",
	// 	gradientTo: "#000000",
	// 	chains: ["ethereum", "optimism", "base"]
	// }
	// ,gearbox: {
	// 	image: "/protocols/gearbox.png",
	// 	gradientFrom: "#000000",
	// 	gradientTo: "#000000",
	// 	chains: ["ethereum", "optimism", "base"]
	// }
	// ,relaylink: {
	// 	image: "/protocols/relaylink.png",
	// 	gradientFrom: "#000000",
	// 	gradientTo: "#000000",
	// 	chains: ["ethereum", "optimism", "base"]
	// }
}

export const abis: Record<string, Record<string, string>> = {
	plug: {
		baseFee:
			"function encode(uint8 $lessThanOrGreaterThan, uint256 $threshold) public pure returns (bytes memory $data)",
		blockNumber:
			"function encode(uint8 $lessThanOrGreaterThan, uint256 $block) public pure returns (bytes memory $data)",
		timestamp:
			"function encode(uint8 $beforeOrAfter, uint256 $timestamp) public pure returns (bytes memory $data)",
		tokenBalance:
			"function encode(address $holder, address $asset, uint8 $lessThanOrGreaterThan, uint256 $amount) public pure returns (bytes memory $data)",
		rateLimit:
			"function encode(uint128 $numberOf, uint32 $frequency, uint32 $duration)",
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
		health: "function encode(uint8 $lessThanOrGreaterThan, uint256 $health) public pure returns (bytes memory $data)",
		borrowRate:
			"function encode(address $asset, uint8 $lessThanOrGreaterThan, uint256 $rate) public pure returns (bytes memory $data)",
		depositRate:
			"function encode(address $asset, uint8 $lessThanOrGreaterThan, uint256 $rate) public pure returns (bytes memory $data)",
		deposit:
			"function encode(uint256 $amount, address $asset) public pure returns (bytes memory $data)",
		loanHealth:
			"function encode(uint8 $lessThanOrGreaterThan, uint256 $health) public pure returns (bytes memory $data)",
		rewardClaim:
			"function encode(uint256 $amount, address $asset) public pure returns (bytes memory $data)"
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
	// chainlink: {
	// 	functionName:
	// 		"function encode(address $variable) public pure returns (bytes memory $data)"
	// },
	// ,velodrome: {
	// 	functionName:
	// 		"function encode(address $variable) public pure returns (bytes memory $data)"
	// }
	// aerodrome: {
	// 	functionName:
	// 		"function encode(address $variable) public pure returns (bytes memory $data)"
	// }
	//,ethena: {
	// 		functionName:
	// 			"function encode(address $variable) public pure returns (bytes memory $data)"
	// 	}
	//,enjoy: {
	// 		functionName:
	// 			"function encode(address $variable) public pure returns (bytes memory $data)"
	// 	}
	//,ens: {
	// 		functionName:
	// 			"function encode(address $variable) public pure returns (bytes memory $data)"
	// 	}
	// pendle: {
	// 	functionName:
	// 		"function encode(address $variable) public pure returns (bytes memory $data)"
	// }
	// ,gearbox: {
	// 		functionName:
	// 			"function encode(address $variable) public pure returns (bytes memory $data)"
	// 	}
	//,relaylink: {
	// 		functionName:
	// 			"function encode(address $variable) public pure returns (bytes memory $data)"
	// 	}
}
