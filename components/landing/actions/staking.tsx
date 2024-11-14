import { useEffect, useState } from "react"

import { motion } from "framer-motion"
import { BookUp2 } from "lucide-react"

import { InfoCard } from "@/components"

const PAIRS = [
	["DAI", "sDAI"],
	["USDC", "sUSDC"],
	["USDT", "sUSDT"],
	["ETH", "stETH"],
	["ETH", "rETH"],
	["BTC", "wBTC"],
	["SOL", "mSOL"],
	["MATIC", "stMATIC"],
	["AVAX", "sAVAX"],
	["CRV", "yCRV"],
	["ATOM", "stATOM"],
	["BNB", "stBNB"]
]

export const ActionStaking = () => {
	const [currentPairIndex, setCurrentPairIndex] = useState(0)
	const [baseToken, stakedToken] = PAIRS[currentPairIndex]

	useEffect(() => {
		const interval = setInterval(() => {
			setCurrentPairIndex(prev => (prev + 1) % PAIRS.length)
		}, 7000)
		return () => clearInterval(interval)
	}, [])

	return (
		<InfoCard
			icon={<BookUp2 size={24} className="opacity-40" />}
			text="Stake."
			description="Earn rewards by staking your tokens in pools."
			className="relative z-[99999] col-span-2 row-span-2 overflow-hidden xl:col-span-1"
		>
			<div className="absolute inset-0 bottom-1/2 overflow-hidden">
				<motion.div
					className="absolute left-1/2 top-[-24rem] flex h-24 w-24 -translate-x-1/2 items-center justify-center rounded-full border-[2px] border-dashed border-plug-green/40 font-bold text-plug-green"
					animate={{
						top: ["-35%", "100%"]
					}}
					transition={{
						duration: 4,
						repeat: Infinity,
						repeatDelay: 3,
						ease: "linear"
					}}
				>
					<p className="relative">${baseToken}</p>
				</motion.div>
			</div>
			<div className="absolute inset-0 top-1/2 z-[999] h-[2px] overflow-hidden bg-plug-yellow" />
			<div className="absolute inset-0 top-1/2 overflow-hidden">
				<motion.div
					className="absolute left-1/2 top-[-24rem] flex h-24 w-24 -translate-x-1/2 items-center justify-center rounded-full border-[2px] border-dashed border-plug-green bg-plug-yellow font-bold text-plug-green"
					animate={{
						top: ["-35%", "100%"]
					}}
					transition={{
						duration: 4,
						repeat: Infinity,
						repeatDelay: 3,
						ease: "linear",
						delay: 3
					}}
				>
					<p className="relative">${stakedToken}</p>
				</motion.div>
			</div>

			<div className="absolute bottom-[20%] left-0 right-0 top-0 bg-gradient-to-b from-plug-white/0 to-plug-white" />
			<div className="absolute bottom-0 left-0 right-0 top-[80%] bg-plug-white" />
		</InfoCard>
	)
}

export default ActionStaking
