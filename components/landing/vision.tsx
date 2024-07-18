import { FC } from "react"

import Image from "next/image"

import { motion } from "framer-motion"
import { Check, Code, PowerOff, Wallet } from "lucide-react"

import { InfoCard, LandingContainer } from "@/components"
import { greenGradientStyle } from "@/lib"

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
	const immutableFactoryBytecode =
		"01011001 01101111 01110101 00100000 01101110 01100101 01110110 01100101 01110010 00100000 01101000 01100001 01110110 01100101 00100000 01110100 01101111 00100000 01100011 01101111 01100100 01100101 00100000 01111001 01101111 01110101 01110010 00100000 01101111 01101110 01100011 01101000 01100001 01101001 01101110 00100000 01110011 01110100 01110010 01100001 01110100 01100101 01100111 01111001 00100000 01100001 01100111 01100001 01101001 01101110 00101110 01011001 01101111 01110101 00100000 01101110 01100101 01110110 01100101 01110010 00100000 01101000 01100001 01110110 01100101 00100000 01110100 01101111 00100000 01100011 01101111 01100100 01100101 00100000 01111001 01101111 01110101 01110010 00100000 01101111 01101110 01100011 01101000 01100001 01101001 01101110 00100000 01110011 01110100 01110010 01100001 01110100 01100101 01100111 01111001 00100000 01100001 01100111 01100001 01101001 01101110 00101110 01011001 01101111 01110101 00100000 01101110 01100101 01110110 01100101 01110010 00100000 01101000 01100001 01110110 01100101 00100000 01110100 01101111 00100000 01100011 01101111 01100100 01100101 00100000 01111001 01101111 01110101 01110010 00100000 01101111 01101110 01100011 01101000 01100001 01101001 01101110 00100000 01110011 01110100 01110010 01100001 01110100 01100101 01100111 01111001 00100000 01100001 01100111 01100001 01101001 01101110 00101110"

	const getRandomDelay = (min: number, max: number) => {
		return Math.random() * (max - min) + min
	}

	return (
		<LandingContainer className="mt-8">
			<div className="grid w-full grid-cols-2 gap-8 xl:grid-cols-6 xl:grid-rows-2">
				<InfoCard
					text={
						<div className="flex flex-row items-center gap-4">
							<PowerOff size={24} className="opacity-40" />
							<span>Plug In and Log Off</span>
						</div>
					}
					description="With your strategies running you can finally step away from the computer. You can put the phone down. You can live with certainty that when all the conditions of your transaction have been met it will be run."
					className="col-span-2 h-[540px] xl:col-span-4 xl:row-span-2 xl:h-full"
					initial={{ opacity: 0, y: 20 }}
					animate={{ opacity: 1, y: 0 }}
					transition={{ duration: 0.2 }}
				>
					<motion.div
						className="mr-[-120px] flex flex-col items-end justify-end gap-4 overflow-y-hidden sm:mr-6 lg:mr-12"
						animate={{
							y: [240, 240 + -1 * actions.length * 80]
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
									<Check
										size={18}
										className="text-[#00E100]"
									/>
								</div>
								<h3 className="flex flex-col gap-1">
									<span className="font-bold md:text-xl">
										{action[0]}
									</span>
									<div className="flex w-full flex-row items-center gap-2 md:text-lg">
										<Image
											src={`/wallets/${action[1]}.png`}
											alt="NFT Chance"
											width={18}
											height={18}
											className="h-4 w-4 rounded-full"
										/>
										<span className="opacity-40">
											{action[1]}.eth
										</span>
									</div>
								</h3>
								<h4 className="mb-auto ml-auto opacity-40">
									{Math.floor(index ** 1.2) + 1} hrs. ago
								</h4>
							</motion.div>
						))}
					</motion.div>
					<div className="absolute bottom-[30%] left-0 right-0 top-1/4 bg-gradient-to-b from-[#FBFBFB]/0 to-[#FBFBFB]" />
					<div className="absolute bottom-0 left-0 right-0 top-[70%] bg-[#FBFBFB]" />
				</InfoCard>
				<InfoCard
					text={
						<div className="flex flex-row items-center gap-4">
							<Code size={24} className="opacity-40" />
							<span>No Code Needed</span>
						</div>
					}
					description="Get started in seconds without writing a single line of code or needing any technical knowledge. Plug and play with building blocks already built for you."
					className="col-span-2 h-[280px] sm:h-[320px] 2xl:h-[300px]"
					initial={{ opacity: 0, y: -20 }}
					whileInView={{ opacity: 1, y: 0 }}
					transition={{ duration: 0.2, delay: 0.2 }}
				>
					<div className="flex flex-wrap items-end justify-end">
						{immutableFactoryBytecode
							.replaceAll(" ", "")
							.split("")
							.map((char, index) => (
								<motion.span
									key={index}
									className="font-bold"
									style={{
										display: "inline-block",
										...greenGradientStyle
									}}
									initial={{ opacity: 0 }}
									whileInView={{
										opacity: [0, 1]
									}}
									transition={{
										duration: getRandomDelay(0.5, 2),
										repeat: Infinity,
										repeatType: "reverse",
										delay: getRandomDelay(0, 2)
									}}
								>
									{char}
								</motion.span>
							))}
					</div>
					<div className="absolute bottom-[45%] left-0 right-0 top-0 bg-gradient-to-b from-[#FBFBFB]/0 to-[#FBFBFB] xl:bottom-[60%]" />
					<div className="absolute bottom-0 left-0 right-0 top-[55%] bg-[#FBFBFB] xl:top-[40%]" />
				</InfoCard>
				<InfoCard
					text={
						<div className="flex flex-row items-center gap-4">
							<Wallet size={24} className="opacity-40" />
							<span>No Upfront Costs</span>
						</div>
					}
					description="With gasless signatures you only have to pay for the transactions that successfully run. No more having to pay for failed transactions."
					className="col-span-2 h-[280px] sm:h-[320px] 2xl:h-[300px]"
					initial={{ opacity: 0, y: 20 }}
					whileInView={{ opacity: 1, y: 0 }}
					transition={{ duration: 0.2, delay: 0.4 }}
				>
					<div className="xs:py-20 grid grid-cols-9 items-center gap-0 py-16 sm:py-20">
						<div className="h-[2px] bg-[#D9D9D9]" />
						<motion.div
							className="h-4 w-full rounded-full"
							animate={{
								backgroundImage: [
									"linear-gradient(30deg, #00E100, #A3F700)",
									"linear-gradient(30deg, #D9D9D9, #D9D9D9)"
								]
							}}
							transition={{
								duration: 1,
								delay: 0,
								repeat: Infinity,
								repeatType: "reverse"
							}}
						/>
						<div className="h-[2px] bg-[#D9D9D9]" />
						<motion.div
							className="h-4 w-full rounded-full"
							animate={{
								backgroundImage: [
									"linear-gradient(30deg, #00E100, #A3F700)",
									"linear-gradient(30deg, #D9D9D9, #D9D9D9)"
								]
							}}
							transition={{
								duration: 1,
								delay: 0.4,
								repeat: Infinity,
								repeatType: "reverse"
							}}
						/>
						<div className="h-[2px] bg-[#D9D9D9]" />
						<motion.div
							className="h-4 w-full rounded-full"
							animate={{
								backgroundImage: [
									"linear-gradient(30deg, #00E100, #A3F700)",
									"linear-gradient(30deg, #D9D9D9, #D9D9D9)"
								]
							}}
							transition={{
								duration: 1,
								delay: 0.8,
								repeat: Infinity,
								repeatType: "reverse"
							}}
						/>
						<div className="h-[2px] bg-[#D9D9D9]" />
						<motion.div
							className="h-4 w-full rounded-full"
							animate={{
								backgroundImage: [
									"linear-gradient(30deg, #00E100, #A3F700)",
									"linear-gradient(30deg, #D9D9D9, #D9D9D9)"
								]
							}}
							transition={{
								duration: 1,
								delay: 1.2,
								repeat: Infinity,
								repeatType: "reverse"
							}}
						/>
						<div className="h-[2px] bg-[#D9D9D9]" />
						<div />
						<p className="mx-auto mt-2 text-xs opacity-60">Sign</p>
						<div />
						<p className="mx-auto mt-2 text-xs opacity-60">
							Simulate
						</p>
						<div />
						<p className="mx-auto mt-2 text-xs opacity-60">
							Execute
						</p>
						<div />
						<p className="mx-auto mt-2 text-xs opacity-60">Pay</p>
					</div>
				</InfoCard>
			</div>
		</LandingContainer>
	)
}
