import { parseAbi, zeroAddress } from "viem"

import { AmbulanceIcon, CircleParkingIcon, CoinsIcon, GaugeIcon, HandCoinsIcon, SquarePercentIcon } from "lucide-react"

import { abis } from "../abis"

export const pendle = {
	functionName: {
		address: zeroAddress,
		abi: abis.pendle.functionName,
		inputs: parseAbi([abis.pendle.functionName])[0]["inputs"],
		// options: [],
		sentence: "SENTENCE {0]",
		info: "PLACEHOLDER FOR LONGER DESC",
		icon: SquarePercentIcon,
		primary: true
	}
}

export default pendle
