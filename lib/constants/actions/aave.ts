import {
	AmbulanceIcon,
	CircleParkingIcon,
	CoinsIcon,
	GaugeIcon,
	HandCoinsIcon,
	SquarePercentIcon
} from "lucide-react"
import { parseAbi, zeroAddress } from "viem"

import { abis } from "../abis"

export const aave = {
	rate: {
		address: zeroAddress,
		abi: abis.fraxlend.rate,
		inputs: parseAbi([abis.fraxlend.rate])[0]["inputs"],
		options: [
			[
				{ label: "Borrow", value: "borrow" },
				{ label: "Lending", value: "lending" }
			],
			undefined,
			[
				{ label: "Less than", value: "<" },
				{ label: "Greater than", value: ">" }
			]
		],
		// {0} is the rate type (borrow/lending) selected by user
		// {1} is the asset type (user will select from an imported list)
		// {2} is the comparison condition (less than/greater than) selected by user
		// {3} is the rate value entered by user
		sentence: "{0} rate for {1} is {2} {3}",
		info: "Check if the borrow or lending rate of a specific asset on FraxLend is less than or greater than the value entered.",
		icon: SquarePercentIcon,
		primary: true
	},
	// utilizationRate: {
	// 	address: zeroAddress,
	// 	abi: abis.fraxlend.utilizationRate,
	// 	inputs: parseAbi([abis.fraxlend.utilizationRate])[0]["inputs"],
	// 	options: [
	// 		[
	// 			{ label: "Less than", value: "<" },
	// 			{ label: "Greater than", value: ">" }
	// 		]
	// 	],
	// 	// {0} is the asset type (user will select from an imported list)
	// 	// {1} is the comparison condition (less than/greater than) selected by user
	// 	// {2} is the utilization rate value entered by user
	// 	sentence: "Utilization rate for {0} is {1} {2}",
	// 	info: "Check the utilization rate of a specific borrow & lend pool on FraxLend to confirm that the utilization rate is less than or greater than the value entered.",
	// 	icon: GaugeIcon,
	// 	primary: true
	// },
	health: {
		address: zeroAddress,
		abi: abis.fraxlend.health,
		inputs: parseAbi([abis.aave.health])[0]["inputs"],
		options: [
			[
				{ label: "Less than", value: "<" },
				{ label: "Greater than", value: ">" }
			],
			undefined
		],
		// {0} is the asset type (user will select from an imported list)
		// {1} is the comparison condition (less than/greater than) selected by user
		// {2} is the health value entered by user
		sentence: "Health of borrow position is {0} {1}",
		info: "Check the health of a specific loan on FraxLend to determine if it is less than or greater than the value entered. This can be used to automatically add or remove collateral based on your loan health.",
		icon: AmbulanceIcon,
		primary: true
	},
	addCollateral: {
		address: zeroAddress,
		abi: abis.fraxlend.addCollateral,
		inputs: parseAbi([abis.aave.deposit])[0]["inputs"],
		options: [
			undefined,
			[
				{
					label: "ETH",
					value: "0x03"
					// imagePath: `${BACKGROUND_PATH}bg-warm.png`
				},
				{
					label: "WETH",
					value: "0x03"
					// imagePath: `${BODY_PATH}body-blue-sky.png`
				},
				{
					label: "USDC",
					value: "0x03"
					// imagePath: `${ACCESSORY_PATH}accessory-bling-anchor.png`
				}
			]
		],
		// {0} is the asset type (user will select from an imported list)
		// {1} is the amount of collateral to add (user input)
		sentence: "Increase collateral with {0} {1}",
		info: "Add the entered amount to a borrow and lend pool on FraxLend. This can be used to increase your borrow limit.",
		icon: HandCoinsIcon
	},
	borrow: {
		address: zeroAddress,
		abi: abis.fraxlend.borrow,
		inputs: parseAbi([abis.fraxlend.borrow])[0]["inputs"],
		options: [
			// No options needed here as the user will input values directly
		],
		// {0} is the amount of $FRAX to borrow (user input)
		// {1} is the asset type to borrow from (user will select from an imported list)
		// {2} is the collateral provided (user input)
		sentence: "Borrow {0} $FRAX from {1} while providing {2} as collateral",
		info: "Borrow the entered amount of $FRAX tokens against a selected token. This can be used to take a loan on FraxLend.",
		icon: CoinsIcon,
		primary: true
	},
	repay: {
		address: zeroAddress,
		abi: abis.fraxlend.repay,
		inputs: parseAbi([abis.fraxlend.repay])[0]["inputs"],
		options: [
			// No options needed here as the user will input values directly
		],
		// {0} is the pool to repay in (user will select from an imported list)
		// {1} is the amount of $FRAX to repay (user input)
		sentence: "Repay $FRAX debt in {0} with {1}",
		info: "Repay the entered amount of $FRAX in the specified pool on FraxLend. This can be used to repay debts and increase loan health, therefore protecting against potential liquidation.",
		icon: HandCoinsIcon
	},
	closePosition: {
		address: zeroAddress,
		abi: abis.fraxlend.closePosition,
		inputs: parseAbi([abis.fraxlend.closePosition])[0]["inputs"],
		options: [
			// No options needed here as the user will input values directly
		],
		// {0} is the pool to close position in (user will select from an imported list)
		sentence: "Close position in {0}",
		info: "Withdraw all assets including deposits and accrued interest from the specified pool on FraxLend.",
		icon: CircleParkingIcon
	}
	// lendFrax: {
	// 	address: zeroAddress,
	// 	abi: abis.fraxlend.lendFrax,
	// 	inputs: parseAbi([abis.fraxlend.lendFrax])[0]["inputs"],
	// 	options: [
	// 		// No options needed here as the user will input values directly
	// 	],
	// 	// {0} is the amount of $FRAX to lend (user input)
	// 	// {1} is the pool to lend in (user will select from an imported list)
	// 	sentence: "Lend {0} $FRAX to {1}",
	// 	info: "Lend the entered amount of $FRAX tokens within a specified pool on FraxLend.",
	// 	icon: HandCoinsIcon
	// },
	// withdrawFrax: {
	// 	address: zeroAddress,
	// 	abi: abis.fraxlend.withdrawFrax,
	// 	inputs: parseAbi([abis.fraxlend.withdrawFrax])[0]["inputs"],
	// 	options: [
	// 		// No options needed here as the user will input values directly
	// 	],
	// 	// {0} is the amount of $FRAX to withdraw (user input)
	// 	// {1} is the pool to withdraw from (user will select from an imported list)
	// 	sentence: "Withdraw {0} $FRAX from {1}",
	// 	info: "Withdraw a specific amount of FRAX tokens used for lending from a specified pool on FraxLend.",
	// 	icon: CoinsIcon
	// }
}

export default aave
