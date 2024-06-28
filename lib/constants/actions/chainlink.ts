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

export const chainlink = {
	functionName: {
		address: zeroAddress,
		abi: abis.chainlink.functionName,
		inputs: parseAbi([abis.chainlink.functionName])[0]["inputs"],
		// options: [],
		sentence: "SENTENCE {0}",
		info: "PLACEHOLDER FOR LONGER DESC",
		icon: SquarePercentIcon,
		primary: true
	}
}

export default chainlink
