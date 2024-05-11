import type { FC } from "react"

import Image from "next/image"

import { motion } from "framer-motion"
import { Check, PowerOff, Sparkles, Unplug } from "lucide-react"

import { InfoCard } from "@/components/cards"
import { Container } from "@/components/landing/container"

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
		"01011001 01101111 01110101 00100000 01101110 01100101 01110110 01100101 01110010 00100000 01101000 01100001 01110110 01100101 00100000 01110100 01101111 00100000 01100011 01101111 01100100 01100101 00100000 01111001 01101111 01110101 01110010 00100000 01101111 01101110 01100011 01101000 01100001 01101001 01101110 00100000 01110011 01110100 01110010 01100001 01110100 01100101 01100111 01111001 00100000 01100001 01100111 01100001 01101001 01101110 00101110"

	const getRandomDelay = (min: number, max: number) => {
		return Math.random() * (max - min) + min
	}

	return (
		<Container className="mb-8">
			<div className="flex flex-col gap-8 lg:grid lg:grid-cols-6 lg:grid-rows-2">
				<InfoCard
					text={
						<>
							<PowerOff size={24} className="opacity-40" />
							<span>Plug In and Log Off</span>
						</>
					}
					description="With your strategies running you can finally step away from the computer. You can put the phone down. You can live with certainty that when all the conditions of your transaction have been met it will be run."
					className="h-[540px] lg:col-span-4 lg:row-span-2 lg:h-full"
					initial={{ opacity: 0, y: 20 }}
					animate={{ opacity: 1, y: 0 }}
					transition={{ duration: 0.2 }}
				>
					<motion.div
						className="flex flex-col items-end justify-end gap-4 overflow-y-hidden lg:mr-12"
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
								className="mr-[-30%] flex w-[460px] items-center gap-6 rounded-lg bg-white px-6 py-2 lg:mr-0 lg:w-[520px]"
							>
								<div className="flex h-6 w-6 items-center justify-center rounded-full bg-[#00EF35]/10">
									<Check
										size={18}
										className="text-[#00EF35]"
									/>
								</div>
								<h3 className="flex flex-col gap-1">
									<span className="text-xl font-bold">
										{action[0]}
									</span>
									<div className="flex w-full flex-row items-center gap-2 text-lg">
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
					<div className="absolute bottom-1/2 left-0 right-0 top-1/4 bg-gradient-to-b from-[#FBFBFB]/0 to-[#FBFBFB] lg:bottom-1/4" />
					<div className="absolute bottom-0 left-0 right-0 top-1/2 bg-[#FBFBFB] lg:top-3/4" />
				</InfoCard>
				<InfoCard
					text={
						<>
							<Unplug size={24} className="opacity-40" />
							<span>No Code Required</span>
						</>
					}
					description="Get started in seconds without writing a single line of code or needing any technical knowledge. Plug and play with building blocks already built for you."
					className="h-[420px] lg:col-span-2 lg:h-full"
					initial={{ opacity: 0, y: -20 }}
					whileInView={{ opacity: 1, y: 0 }}
					transition={{ duration: 0.2, delay: 0.2 }}
				>
					<div className="flex flex-wrap items-end justify-end">
						{immutableFactoryBytecode
							.slice(0, 240)
							.replaceAll(" ", "")
							.split("")
							.map((char, index) => {
								const random = Math.random()
								const color =
									random < 0.5 ? "transparent" : "#00EF35"

								return (
									<motion.span
										key={index}
										className="font-bold"
										style={{
											display: "inline-block",
											color
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
								)
							})}
					</div>
					<div className="absolute bottom-1/2 left-0 right-0 top-0 bg-gradient-to-b from-[#FBFBFB]/0 to-[#FBFBFB] lg:bottom-1/4" />
					<div className="absolute bottom-0 left-0 right-0 top-1/2 bg-[#FBFBFB] lg:top-3/4" />
				</InfoCard>
				<InfoCard
					text={
						<>
							<Sparkles size={24} className="opacity-40" />
							<span>Execute With The Best</span>
						</>
					}
					description="You don't have to be an expert. Choose from a curated catalog of Plugs to get started. Fork them and make changes with a simple interface."
					className="h-[320px] lg:col-span-2 lg:h-full"
					initial={{ opacity: 0, y: 20 }}
					whileInView={{ opacity: 1, y: 0 }}
					transition={{ duration: 0.2, delay: 0.4 }}
				/>
			</div>
		</Container>
	)
}
