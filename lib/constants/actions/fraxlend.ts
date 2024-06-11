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

export const fraxlend = {
	rate: {
		address: zeroAddress,
		abi: abis.fraxlend.rate,
		inputs: parseAbi([abis.fraxlend.rate])[0]["inputs"],
		sentence: "{0} rate for {1} is {2} {3}",
		info: "Check if the borrow or lending rate of a specific asset on FraxLend is less than or greater than the value entered.",
		icon: SquarePercentIcon,
		primary: true
	},
	utilizationRate: {
		address: zeroAddress,
		abi: abis.fraxlend.utilizationRate,
		inputs: parseAbi([abis.fraxlend.utilizationRate])[0]["inputs"],
		sentence: "Utilization rate for {0} is {1 2}",
		info: "Check the utilization rate of a specific borrow & lend pool on FraxLend to confirm that the utilization rate is less than or greater than the value entered.",
		icon: GaugeIcon,
		primary: true
	},
	health: {
		address: zeroAddress,
		abi: abis.fraxlend.health,
		inputs: parseAbi([abis.fraxlend.health])[0]["inputs"],
		sentence: "Health of borrow position in {0} is {1} {2}",
		info: "Check the health of a specific loan on FraxLend to determine if it is less than or greater than the value entered. This can be used to automatically add or remove collateral based on your loan health.",
		icon: AmbulanceIcon,
		primary: true
	},
	addCollateral: {
		address: zeroAddress,
		abi: abis.fraxlend.addCollateral,
		inputs: parseAbi([abis.fraxlend.addCollateral])[0]["inputs"],
		sentence: "Increase collateral in {0} with {1}",
		info: "Add the entered amount to a borrow and lend pool on FraxLend. This can be used to increase your borrow limit.",
		icon: HandCoinsIcon
	},
	borrow: {
		address: zeroAddress,
		abi: abis.fraxlend.borrow,
		inputs: parseAbi([abis.fraxlend.borrow])[0]["inputs"],
		sentence: "Borrow {0} $FRAX from {1} while providing {2} as collateral",
		info: "Borrow the entered amount of $FRAX tokens against a selected token. This can be used to take a loan on FraxLend.",
		icon: CoinsIcon,
		primary: true
	},
	repay: {
		address: zeroAddress,
		abi: abis.fraxlend.repay,
		inputs: parseAbi([abis.fraxlend.repay])[0]["inputs"],
		sentence: "Repay $FRAX debt in {0} with {1}",
		info: "Repay a the entered amount of $FRAX in the specifed pool on FraxLend. This can be used to repay debts and increase loan health therefore protecting against potential liquidation.",
		icon: HandCoinsIcon
	},
	closePosition: {
		address: zeroAddress,
		abi: abis.fraxlend.closePosition,
		inputs: parseAbi([abis.fraxlend.closePosition])[0]["inputs"],
		sentence: "Close position in {0}",
		info: "Withdraw all assets including deposits and accrued interest from the specifed pool on FraxLend.",
		icon: CircleParkingIcon
	},
	lendFrax: {
		address: zeroAddress,
		abi: abis.fraxlend.lendFrax,
		inputs: parseAbi([abis.fraxlend.lendFrax])[0]["inputs"],
		sentence: "Lend {0} $FRAX to {1}",
		info: "Lend the entered amount of $FRAX tokens within a specified pool on FraxLend.",
		icon: HandCoinsIcon
	},
	withdrawFrax: {
		address: zeroAddress,
		abi: abis.fraxlend.withdrawFrax,
		inputs: parseAbi([abis.fraxlend.withdrawFrax])[0]["inputs"],
		sentence: "Withdraw {0} $FRAX from {1}",
		info: "Withdraw a specific amount of FRAX tokens used for lending from a specified pool on FraxLend.",
		icon: CoinsIcon
	}
} as const

export default fraxlend
