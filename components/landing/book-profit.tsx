import Image from "next/image"

import { motion } from "framer-motion"
import { CheckCircle, PowerOff } from "lucide-react"

import { InfoCard } from "@/components"

export const BookProfit = () => {
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
			icon={<PowerOff size={24} className="opacity-40" />}
			text="Book profit even when you're sleeping."
			description="Keep your money working even when you aren’t. Step away from the computer knowing 
						that your transactions are executing no matter where you are or what you’re doing. "
			className="col-span-2 h-[540px] xl:col-span-4 xl:row-span-2 xl:h-full"
		>
			<motion.div
				className="mr-[-120px] flex flex-col items-end justify-end gap-2 overflow-y-hidden sm:mr-6 lg:mr-12"
				animate={{
					transform: [`translateY(240px)`, `translateY(${240 + -1 * actions.length * 80}px)`]
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
						className="mr-0 flex w-[460px] items-center gap-4 overflow-hidden rounded-lg bg-white px-6 py-4 md:w-[640px]"
					>
						<div className="relative flex h-10 w-16 items-center justify-center">
							<div className="absolute top-1/2 h-16 w-16 -translate-y-1/2 rounded-full bg-gradient-to-tr from-plug-green to-plug-yellow blur-[60px] filter" />
							<CheckCircle
								className="absolute top-1/2 ml-auto h-5 w-5 -translate-y-1/2 text-center"
								size={24}
								style={{
									stroke: "url(#plug-gradient)"
								}}
							/>
							<svg width="0" height="0">
								<linearGradient id="plug-gradient" x1="0%" y1="0%" x2="100%" y2="100%">
									<stop stopColor="#00E100" offset="0%" />
									<stop stopColor="#A3F700" offset="100%" />
								</linearGradient>
							</svg>
						</div>

						<h3 className="flex w-full flex-col">
							<span className="font-bold md:text-xl">{action[0]}</span>
							<span className="flex w-full flex-row items-center gap-2 md:text-lg">
								<Image
									src={`/wallets/${action[1]}.png`}
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
			<div className="absolute bottom-[30%] left-0 right-0 top-1/4 bg-gradient-to-b from-grayscale-0/0 to-grayscale-0" />
			<div className="absolute bottom-0 left-0 right-0 top-[70%] bg-grayscale-0" />
		</InfoCard>
	)
}
