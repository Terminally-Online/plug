import { FC } from "react"

import Image from "next/image"

import { motion } from "framer-motion"
import { Check, Code, PowerOff, Wallet } from "lucide-react"

import { Cardiogram, InfoCard, LandingContainer } from "@/components"

export const Vision: FC = () => {
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
		<LandingContainer className="mb-[80px]">
			<div className="grid w-full grid-cols-2 gap-8 xl:grid-cols-6 xl:grid-rows-2">
				<InfoCard
					icon={<PowerOff size={24} className="opacity-40" />}
					text="Book profit even when you're sleeping."
					description="Stop missing opportunities and have your transactions execute no matter where you are or what you're doing. If you want it done, it will be delivered on a silver platter."
					className="col-span-2 h-[540px] xl:col-span-4 xl:row-span-2 xl:h-full"
				>
					<motion.div
						className="mr-[-120px] flex flex-col items-end justify-end gap-4 overflow-y-hidden sm:mr-6 lg:mr-12"
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
								className="mr-0 flex w-[460px] items-center gap-6 rounded-lg bg-white px-6 py-2 md:w-[640px] md:gap-12"
							>
								<div className="flex h-6 w-6 items-center justify-center rounded-full bg-[#00E100]/10">
									<Check size={18} className="text-[#00E100]" />
								</div>
								<h3 className="flex flex-col gap-1">
									<span className="font-bold md:text-xl">{action[0]}</span>
									<div className="flex w-full flex-row items-center gap-2 md:text-lg">
										<Image
											src={`/wallets/${action[1]}.png`}
											alt="NFT Chance"
											width={18}
											height={18}
											className="h-4 w-4 rounded-full"
										/>
										<span className="opacity-40">{action[1]}.eth</span>
									</div>
								</h3>
								<h4 className="mb-auto ml-auto opacity-40">{Math.floor(index ** 1.2) + 1} hrs. ago</h4>
							</motion.div>
						))}
					</motion.div>
					<div className="absolute bottom-[30%] left-0 right-0 top-1/4 bg-gradient-to-b from-[#FBFBFB]/0 to-[#FBFBFB]" />
					<div className="absolute bottom-0 left-0 right-0 top-[70%] bg-[#FBFBFB]" />
				</InfoCard>

				<InfoCard
					icon={<Code size={24} className="opacity-40" />}
					text="You're underperforming."
					description="You don't need to be a rocket scientist to outcompete. Top performers no longer execute transactions manually."
					className="col-span-2 h-[280px] sm:h-[320px] 2xl:h-[300px]"
				/>

				<Cardiogram />
			</div>
		</LandingContainer>
	)
}
