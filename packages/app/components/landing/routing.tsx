import Image from "next/image"

import { motion } from "framer-motion"
import { Sun } from "lucide-react"

import { InfoCard } from "@/components"

export const Routing = () => {
	const actions = [
		["Dollar Cost Average USDC:ETH", "danner"],
		["Auto-Redeem MKR Backing", "federalreserve"],
		["Bid on Noun with Pineapple Hat", "nftchance"],
		["Buy ETH on Market Dump", "federalreserve"],
		["Auto-Renew ENS", "danner"],
		["Stream 65 ETH to Team", "nftchance"],
		["Enter Gearbox at Target APY", "nftchance"],
		["Bid on Noun", "nftchance"],
		["Rebalance Memecoin Portfolio", "federalreserve"],
		["Fill Ethena Liquidity Cap to Limit", "nftchance"],
		["Compound Enjoy Staking Rewards", "nftchance"],
		["Top-Up Loan Health Factor", "federalreserve"],
		["Exit Yearn at Target APY", "nftchance"]
	]

	return (
		<InfoCard
			icon={<Sun size={24} className="opacity-40" />}
			text="Running 24 hours."
			description="Stay active in a market that never sleeps and keep your money working when you aren't."
			className="col-span-2 h-[280px] sm:h-[320px] 2xl:h-[300px]"
		>
			<motion.div
				className="relative z-[-1] mr-[-20px] flex flex-col items-end justify-end gap-2 overflow-y-hidden"
				animate={{
					transform: [`translateY(80px)`, `translateY(${80 + -1 * actions.length * 80}px)`]
				}}
				transition={{
					duration: 30,
					repeat: Infinity,
					repeatType: "reverse",
					repeatDelay: 1,
					ease: "easeInOut"
				}}
			>
				{actions.map((action, index) => (
					<motion.div
						key={index}
						className="mr-0 flex w-[320px] items-center gap-4 overflow-hidden rounded-lg border-[1px] border-plug-green/10 bg-white px-6 py-4"
					>
						<h3 className="flex w-full flex-col">
							<span className="font-bold">{action[0]}</span>
							<span className="flex w-full flex-row items-center gap-2">
								<Image
									src={`/users/${action[1]}.png`}
									alt="NFT Chance"
									width={18}
									height={18}
									className="h-4 w-4 rounded-full"
								/>
								<span className="flex w-full justify-between font-bold opacity-40">
									<span className="w-full">{action[1]}.eth</span>
									<span className="ml-auto whitespace-nowrap">
										{Math.floor(index ** 1.2) + 1} hrs. ago
									</span>
								</span>
							</span>
						</h3>
					</motion.div>
				))}
			</motion.div>

			<div className="absolute bottom-[50%] left-0 right-0 top-[40%] bg-gradient-to-b from-plug-white/0 to-plug-white" />
			<div className="absolute bottom-0 left-0 right-0 top-[50%] bg-plug-white" />
		</InfoCard>
	)
}
