import { FC } from "react"

import Image from "next/image"

import { motion } from "framer-motion"

import { LandingContainer } from "@/components"

export const Examples: FC = () => (
	<>
		<LandingContainer className="mt-[90px] flex-col items-center gap-4">
			<h2 className="text-center text-[28px] font-bold lg:w-[60%] lg:text-[64px] 2xl:w-[40%]">
				A simple way to plug-and-play.
			</h2>
			<p className="text-center text-[18px] font-light opacity-40 lg:w-[45%] lg:text-[24px] 2xl:w-[35%]">
				Declare multi-outcome transactions with a state of the art
				no-code builder in seconds and combine the power of top
				protocols.
			</p>
		</LandingContainer>

		<div className="my-[45px] border-y-[2px] border-[#D9D9D9] lg:my-[90px]">
			<div className="grid lg:grid-cols-12 lg:grid-rows-2">
				<div className="lg:col-span-6">
					<div className="p-4 py-8 lg:p-12 lg:pl-16 xl:pl-64 2xl:pl-80">
						<div className="flex flex-col">
							<div className="flex flex-row items-center gap-4">
								<motion.div
									className="flex h-6 w-6 items-center justify-center rounded-full bg-[#D9D9D9]/40"
									animate={{
										background: ["#EAEAEA", "#D9D9D9"]
									}}
									transition={{
										duration: 0.25,
										repeat: Infinity,
										repeatType: "reverse",
										repeatDelay: 2
									}}
								>
									<p className="text-xs">1</p>
								</motion.div>
								<Image
									src="/protocols/plug.png"
									alt="Plug"
									width={24}
									height={24}
									className="rounded-full"
								/>
								<p>
									Run{" "}
									<span className="rounded-md bg-[#00E100]/10 p-2 py-1 text-[#00E100]">
										1
									</span>{" "}
									time a{" "}
									<span className="rounded-md bg-[#00E100]/10 p-2 py-1 text-[#00E100]">
										year
									</span>{" "}
									.
								</p>
							</div>
							<div className="ml-[11px] h-4 w-[2px] bg-[#D9D9D9]" />
							<div className="flex flex-row items-center gap-4">
								<motion.div
									className="flex h-6 w-6 items-center justify-center rounded-full bg-[#D9D9D9]/40"
									animate={{
										background: ["#EAEAEA", "#D9D9D9"]
									}}
									transition={{
										duration: 0.25,
										delay: 0.5,
										repeat: Infinity,
										repeatType: "reverse",
										repeatDelay: 2
									}}
								>
									<p className="text-xs">2</p>
								</motion.div>
								<Image
									src="/protocols/uniswap.png"
									alt="Plug"
									width={24}
									height={24}
									className="rounded-full"
								/>
								<p>
									Can swap{" "}
									<span className="rounded-md bg-[#00E100]/10 p-2 py-1 text-[#00E100]">
										100
									</span>{" "}
									<span className="rounded-md bg-[#00E100]/10 p-2 py-1 text-[#00E100]">
										$USDC
									</span>{" "}
									to{" "}
									<span className="rounded-md bg-[#00E100]/10 p-2 py-1 text-[#00E100]">
										$ETH
									</span>{" "}
									.
								</p>
							</div>
							<div className="ml-[11px] h-4 w-[2px] bg-[#D9D9D9]" />
							<div className="flex flex-row items-center gap-4">
								<motion.div
									className="flex h-6 w-6 items-center justify-center rounded-full bg-[#D9D9D9]/40"
									animate={{
										background: ["#EAEAEA", "#D9D9D9"]
									}}
									transition={{
										duration: 0.25,
										delay: 1,
										repeat: Infinity,
										repeatType: "reverse",
										repeatDelay: 2
									}}
								>
									<p className="text-xs">3</p>
								</motion.div>
								<Image
									src="/protocols/plug.png"
									alt="Plug"
									width={24}
									height={24}
									className="rounded-full"
								/>
								<p>
									Has{" "}
									<span className="rounded-md bg-[#00E100]/10 p-2 py-1 text-[#00E100]">
										0.2
									</span>{" "}
									<span className="rounded-md bg-[#00E100]/10 p-2 py-1 text-[#00E100]">
										$ETH
									</span>{" "}
									or greater.
								</p>
							</div>
							<div className="ml-[11px] h-4 w-[2px] bg-[#D9D9D9]" />
							<div className="flex flex-row items-center gap-4">
								<motion.div
									className="flex h-6 w-6 items-center justify-center rounded-full bg-[#D9D9D9]/40"
									animate={{
										background: ["#EAEAEA", "#D9D9D9"]
									}}
									transition={{
										duration: 0.25,
										delay: 1.5,
										repeat: Infinity,
										repeatType: "reverse",
										repeatDelay: 2
									}}
								>
									<p className="text-xs">4</p>
								</motion.div>
								<Image
									src="/protocols/ens.png"
									alt="Plug"
									width={24}
									height={24}
									className="rounded-full"
								/>
								<p>
									Can renew{" "}
									<span className="rounded-md bg-[#00E100]/10 p-2 py-1 text-[#00E100]">
										nftchance.eth
									</span>{" "}
									.
								</p>
							</div>
						</div>
					</div>
				</div>
				<div className="flex border-t-[2px] border-[#D9D9D9] lg:col-span-5 lg:col-start-7 lg:row-span-2 lg:border-l-[2px] lg:border-t-[0px] 2xl:pl-48">
					<div className="my-auto p-4 py-8 pr-0 lg:pl-24 xl:p-16">
						<div className="flex flex-col">
							<div className="flex flex-row items-center gap-4">
								<motion.div
									className="flex h-6 w-6 items-center justify-center rounded-full bg-[#D9D9D9]/40"
									animate={{
										background: ["#EAEAEA", "#D9D9D9"]
									}}
									transition={{
										duration: 0.25,
										repeat: Infinity,
										repeatType: "reverse",
										repeatDelay: 3.5
									}}
								>
									<p className="text-xs">1</p>
								</motion.div>
								<Image
									src="/protocols/plug.png"
									alt="Plug"
									width={24}
									height={24}
									className="rounded-full"
								/>
								<p>
									Run{" "}
									<span className="rounded-md bg-[#00E100]/10 p-2 py-1 text-[#00E100]">
										1
									</span>{" "}
									time a{" "}
									<span className="rounded-md bg-[#00E100]/10 p-2 py-1 text-[#00E100]">
										day
									</span>{" "}
									.
								</p>
							</div>
							<div className="ml-[11px] h-4 w-[2px] bg-[#D9D9D9]" />
							<div className="flex flex-row items-center gap-4">
								<motion.div
									className="flex h-6 w-6 items-center justify-center rounded-full bg-[#D9D9D9]/40"
									animate={{
										background: ["#EAEAEA", "#D9D9D9"]
									}}
									transition={{
										duration: 0.25,
										delay: 0.5,
										repeat: Infinity,
										repeatType: "reverse",
										repeatDelay: 3.5
									}}
								>
									<p className="text-xs">2</p>
								</motion.div>
								<Image
									src="/protocols/plug.png"
									alt="Plug"
									width={24}
									height={24}
									className="rounded-full"
								/>
								<p>
									Run after{" "}
									<span className="rounded-md bg-[#00E100]/10 p-2 py-1 text-[#00E100]">
										10
									</span>{" "}
									<span className="rounded-md bg-[#00E100]/10 p-2 py-1 text-[#00E100]">
										PM
									</span>{" "}
									<span className="rounded-md bg-[#00E100]/10 p-2 py-1 text-[#00E100]">
										UTC
									</span>{" "}
									.
								</p>
							</div>
							<div className="ml-[11px] h-4 w-[2px] bg-[#D9D9D9]" />
							<div className="flex flex-row items-center gap-4">
								<motion.div
									className="flex h-6 w-6 items-center justify-center rounded-full bg-[#D9D9D9]/40"
									animate={{
										background: ["#EAEAEA", "#D9D9D9"]
									}}
									transition={{
										duration: 0.25,
										delay: 1,
										repeat: Infinity,
										repeatType: "reverse",
										repeatDelay: 3.5
									}}
								>
									<p className="text-xs">3</p>
								</motion.div>
								<Image
									src="/protocols/plug.png"
									alt="Plug"
									width={24}
									height={24}
									className="rounded-full"
								/>
								<p>
									Run before{" "}
									<span className="rounded-md bg-[#00E100]/10 p-2 py-1 text-[#00E100]">
										12/31/2024
									</span>{" "}
									.
								</p>
							</div>
							<div className="ml-[11px] h-4 w-[2px] bg-[#D9D9D9]" />
							<div className="flex flex-row items-center gap-4">
								<motion.div
									className="flex h-6 w-6 items-center justify-center rounded-full bg-[#D9D9D9]/40"
									animate={{
										background: ["#EAEAEA", "#D9D9D9"]
									}}
									transition={{
										duration: 0.25,
										delay: 1.5,
										repeat: Infinity,
										repeatType: "reverse",
										repeatDelay: 3.5
									}}
								>
									<p className="text-xs">4</p>
								</motion.div>
								<Image
									src="/protocols/uniswap.png"
									alt="Plug"
									width={24}
									height={24}
									className="rounded-full"
								/>
								<p>
									Swap{" "}
									<span className="rounded-md bg-[#00E100]/10 p-2 py-1 text-[#00E100]">
										36,000
									</span>{" "}
									<span className="rounded-md bg-[#00E100]/10 p-2 py-1 text-[#00E100]">
										$USDC
									</span>{" "}
									to{" "}
									<span className="rounded-md bg-[#00E100]/10 p-2 py-1 text-[#00E100]">
										$ETH
									</span>
									.
								</p>
							</div>
							<div className="ml-[11px] h-4 w-[2px] bg-[#D9D9D9]" />
							<div className="flex flex-row items-center gap-4">
								<motion.div
									className="flex h-6 w-6 items-center justify-center rounded-full bg-[#D9D9D9]/40"
									animate={{
										background: ["#EAEAEA", "#D9D9D9"]
									}}
									transition={{
										duration: 0.25,
										delay: 2,
										repeat: Infinity,
										repeatType: "reverse",
										repeatDelay: 3.5
									}}
								>
									<p className="text-xs">5</p>
								</motion.div>
								<Image
									src="/protocols/nouns.png"
									alt="Plug"
									width={24}
									height={24}
									className="rounded-full"
								/>
								<p>
									Can bid with{" "}
									<span className="rounded-md bg-[#00E100]/10 p-2 py-1 text-[#00E100]">
										9
									</span>{" "}
									<span className="rounded-md bg-[#00E100]/10 p-2 py-1 text-[#00E100]">
										$ETH
									</span>{" "}
									.
								</p>
							</div>
							<div className="ml-[11px] h-4 w-[2px] bg-[#D9D9D9]" />
							<div className="flex flex-row items-center gap-4">
								<motion.div
									className="flex h-6 w-6 items-center justify-center rounded-full bg-[#D9D9D9]/40"
									animate={{
										background: ["#EAEAEA", "#D9D9D9"]
									}}
									transition={{
										duration: 0.25,
										delay: 3,
										repeat: Infinity,
										repeatType: "reverse",
										repeatDelay: 3.5
									}}
								>
									<p className="text-xs">6</p>
								</motion.div>
								<Image
									src="/protocols/nouns.png"
									alt="Plug"
									width={24}
									height={24}
									className="rounded-full"
								/>
								<p>
									Bid on Noun with{" "}
									<span className="rounded-md bg-[#00E100]/10 p-2 py-1 text-[#00E100]">
										Pineapple Hat
									</span>{" "}
									.
								</p>
							</div>
						</div>
					</div>
				</div>
				<div className="border-t-[2px] border-[#D9D9D9] lg:col-span-6">
					<div className="flex h-full items-center p-4 py-8 lg:p-12 lg:pl-16 xl:pl-64 2xl:pl-80">
						<div className="flex flex-col">
							<div className="flex flex-row items-center gap-4">
								<motion.div
									className="flex h-6 w-6 items-center justify-center rounded-full bg-[#D9D9D9]/40"
									animate={{
										background: ["#EAEAEA", "#D9D9D9"]
									}}
									transition={{
										duration: 0.25,
										repeat: Infinity,
										repeatType: "reverse",
										repeatDelay: 1.5
									}}
								>
									<p className="text-xs">1</p>
								</motion.div>
								<Image
									src="/protocols/plug.png"
									alt="Plug"
									width={24}
									height={24}
									className="rounded-full"
								/>
								<p>
									Run{" "}
									<span className="rounded-md bg-[#00E100]/10 p-2 py-1 text-[#00E100]">
										1
									</span>{" "}
									time a{" "}
									<span className="rounded-md bg-[#00E100]/10 p-2 py-1 text-[#00E100]">
										hour
									</span>{" "}
									.
								</p>
							</div>
							<div className="ml-[11px] h-4 w-[2px] bg-[#D9D9D9]" />
							<div className="flex flex-row items-center gap-4">
								<motion.div
									className="flex h-6 w-6 items-center justify-center rounded-full bg-[#D9D9D9]/40"
									animate={{
										background: ["#EAEAEA", "#D9D9D9"]
									}}
									transition={{
										duration: 0.25,
										delay: 0.5,
										repeat: Infinity,
										repeatType: "reverse",
										repeatDelay: 1.5
									}}
								>
									<p className="text-xs">2</p>
								</motion.div>
								<Image
									src="/protocols/yearn.png"
									alt="Plug"
									width={24}
									height={24}
									className="rounded-full"
								/>
								<p>
									<span className="rounded-md bg-[#00E100]/10 p-2 py-1 text-[#00E100]">
										$USDC
									</span>{" "}
									pool is above{" "}
									<span className="rounded-md bg-[#00E100]/10 p-2 py-1 text-[#00E100]">
										72%
									</span>{" "}
									APY.
								</p>
							</div>
							<div className="ml-[11px] h-4 w-[2px] bg-[#D9D9D9]" />
							<div className="flex flex-row items-center gap-4">
								<motion.div
									className="flex h-6 w-6 items-center justify-center rounded-full bg-[#D9D9D9]/40"
									animate={{
										background: ["#EAEAEA", "#D9D9D9"]
									}}
									transition={{
										duration: 0.25,
										delay: 1,
										repeat: Infinity,
										repeatType: "reverse",
										repeatDelay: 1.5
									}}
								>
									<p className="text-xs">3</p>
								</motion.div>
								<Image
									src="/protocols/yearn.png"
									alt="Plug"
									width={24}
									height={24}
									className="rounded-full"
								/>
								<p>
									Can deposit{" "}
									<span className="rounded-md bg-[#00E100]/10 p-2 py-1 text-[#00E100]">
										10,000
									</span>{" "}
									<span className="rounded-md bg-[#00E100]/10 p-2 py-1 text-[#00E100]">
										$USDC
									</span>{" "}
									.
								</p>
							</div>
						</div>
					</div>
				</div>
			</div>
		</div>
	</>
)
