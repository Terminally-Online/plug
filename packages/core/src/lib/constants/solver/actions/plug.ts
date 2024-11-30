import {
	InfinityIcon,
	ClockIcon,
	FuelIcon,
	GemIcon,
	HashIcon,
	SnowflakeIcon,
	Tally5Icon
} from "lucide-react"
import { parseAbi, zeroAddress } from "viem"

import { ActionProvider } from "@/src/lib/types"

import { mainnet } from "../network"

const abis = {
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
}

export const plug: ActionProvider = {
	info: {
		image: "/protocols/plug.png",
		gradient: ["#00E100", "#A3F700"],
		tags: ["consumer"]
	},
	actions: {
		baseFee: {
			address: `${mainnet}:${zeroAddress}`,
			abi: abis.baseFee,
			inputs: parseAbi([abis.baseFee])[0]["inputs"],
			options: [
				[
					{ label: "Greater than", value: ">" },
					{ label: "Less than", value: "<" }
				],
				[
					{ label: "Low", value: "low" },
					{ label: "Medium", value: "medium" },
					{ label: "High", value: "high" }
				]
			],
			sentence: "Base gas fee is {0} {1}",
			info: "Only allow this Plug to be executed when the base gas fee is greater or less than the value entered. This can be used to schedule future executions when gas is low.",
			icon: FuelIcon,
			primary: true
		},
		tokenBalance: {
			address: `${mainnet}:${zeroAddress}`,
			abi: abis.tokenBalance,
			inputs: parseAbi([abis.tokenBalance])[0]["inputs"],
			options: [
				undefined,
				[
					{ label: "Greater than", value: ">" },
					{ label: "Less than", value: "<" },
					{ label: "Equal to", value: "=" }
				],
				undefined
			],
			sentence: "Balance of {1} is {2} {3}",
			info: "Only allow this Plug to be executed when the balance of the selected token is greater or lower than the value entered for the address entered. This can be used to auto-allocate funds that are received in this Socket or to act based on token changes in another address.",
			icon: GemIcon,
			primary: true
		},
		timestamp: {
			address: `${mainnet}:${zeroAddress}`,
			abi: abis.timestamp,
			inputs: parseAbi([abis.timestamp])[0]["inputs"],
			options: [
				[
					{ label: "Greater than", value: ">" },
					{ label: "Less than", value: "<" },
					{ label: "Equal to", value: "=" }
				],
				undefined
			],
			sentence: "Is {0} {1}",
			info: "Only allow this Plug to be executed when the current time is before or after the value entered. This can be used to schedule a transaction to occur before or after a specified date and time.",
			icon: ClockIcon,
			primary: true
		},
		limitedCalls: {
			address: `${mainnet}:${zeroAddress}`,
			abi: abis.limitedCalls,
			inputs: parseAbi([abis.limitedCalls])[0]["inputs"],
			options: undefined,
			sentence: "Can only ever be called {0} times",
			info: "Only allow this Plug to be executed a certain number of times before expiring. This can be used to execute a transaction a pre-determined amount of times.",
			icon: Tally5Icon,
			primary: true
		},
		cooldown: {
			address: `${mainnet}:${zeroAddress}`,
			abi: abis.cooldown,
			inputs: parseAbi([abis.cooldown])[0]["inputs"],
			options: [
				undefined,
				[
					{ label: "Seconds", value: "seconds" },
					{ label: "Minutes", value: "minutes" },
					{ label: "Hours", value: "hours" },
					{ label: "Days", value: "days" },
					{ label: "Weeks", value: "weeks" },
					{ label: "Months", value: "months" },
					{ label: "Quarters", value: "quarters" },
					{ label: "Years", value: "years" }
				]
			],
			sentence: "Time between calls is at least {0} {1}",
			info: "Only allow this Plug to be executed after a certain amount of time has passed since the last execution. This can be used to set up recurring transactions with a pre-determined gap of time between them.",
			icon: SnowflakeIcon
		},
		rateLimit: {
			address: `${mainnet}:${zeroAddress}`,
			abi: abis.rateLimit,
			inputs: parseAbi([abis.rateLimit])[0]["inputs"],
			options: [
				undefined,
				undefined,
				[
					{ label: "Seconds", value: "seconds" },
					{ label: "Minutes", value: "minutes" },
					{ label: "Hours", value: "hours" },
					{ label: "Days", value: "days" },
					{ label: "Weeks", value: "weeks" },
					{ label: "Months", value: "months" },
					{ label: "Quarters", value: "quarters" },
					{ label: "Years", value: "years" }
				]
			],
			sentence: "Can be called {0} times every {1} {2}",
			info: "Control the frequency at which this Plug can be called. This can be used to set up a pre-determined number of recurring transactions with a pre-determined gap of time between them.",
			icon: InfinityIcon
		},

		blockNumber: {
			address: `${mainnet}:${zeroAddress}`,
			abi: abis.blockNumber,
			inputs: parseAbi([abis.blockNumber])[0]["inputs"],
			options: [
				[
					{ label: "Greater than", value: ">" },
					{ label: "Less than", value: "<" },
					{ label: "Equal to", value: "=" }
				],
				undefined
			],
			sentence: "Block number must be {0} {1}",
			info: "Only allow this Plug to be executed when the block number is greater or lower than the value entered. This is an advanced feature, you may prefer to use the timestamp condition.",
			icon: HashIcon
		}
	}
}

export default plug
