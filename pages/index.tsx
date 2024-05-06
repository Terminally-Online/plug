import type { FC } from "react"
import { useState } from "react"

import Head from "next/head"
import Image from "next/image"

import { motion } from "framer-motion"
import {
	ArrowDownWideNarrow,
	CalendarClock,
	Check,
	ChevronDown,
	Clock,
	FilePen,
	FileStack,
	PowerOff,
	Puzzle,
	Rotate3d,
	RotateCw,
	Ruler,
	ShieldPlus,
	SlidersHorizontal,
	Sparkles,
	TestTubeDiagonal,
	Trophy,
	Unplug,
	X
} from "lucide-react"

import { Footer, Navbar } from "@/components/base"
import { MainButton } from "@/components/buttons"
import { ActionCard, InfoCard, StepCard } from "@/components/cards"
import { LandingContainer as Container } from "@/components/container"
import { Ecosystem, Glitter, Version } from "@/components/effect"
import { colors, greenGradientStyle } from "@/lib/constants"
import { routes } from "@/lib/routes"

const Hero = () => (
	<div className="relative flex h-[1050px] w-full w-full overflow-hidden lg:h-[900px] lg:items-center">
		<Container>
			<div className="mt-4 flex flex-col gap-[15px] lg:mt-0 lg:max-w-[70%] lg:gap-[30px]">
				<h1 className="text-[42px] font-bold lg:text-[72px] 2xl:text-[96px]">
					Automate your onchain activity with an{" "}
					<span style={{ ...greenGradientStyle }}>
						“if this, then that”
					</span>{" "}
					interface.
				</h1>
				<p className="text-[18px] font-light text-black/40 lg:max-w-[85%] lg:text-[24px]">
					Use Plug to build your own transaction workflows or choose
					from community generated strategies. Let our bots execute
					your transactions and never worry about missing an
					opportunity again. No code needed.
				</p>
				<div>
					<a
						href={routes.earlyAccess}
						target="_blank"
						rel="noreferrer"
					>
						<MainButton
							text="Get Early Access"
							className="mt-[30px] w-max"
						/>
					</a>
				</div>
			</div>
		</Container>

		<Ecosystem />
	</div>
)

const CallToAction: FC<{
	text: string
	description: string
	button: string
}> = ({ text, description, button }) => (
	<Container className="flex-col">
		<motion.div
			className="flex flex-col justify-center gap-[15px] rounded-lg bg-gradient-to-tr from-[#00EF35] to-[#93DF00] p-8 lg:min-h-[700px] lg:gap-[30px] lg:px-[80px]"
			initial={{ opacity: 0, y: 20 }}
			whileInView={{ opacity: 1, y: 0 }}
			transition={{ duration: 0.2 }}
		>
			<h1 className="text-[36px] font-bold text-white lg:max-w-[80%] lg:text-[72px] lg:text-[96px]">
				{text}
			</h1>
			<p className="text-[18px] text-white lg:max-w-[60%] lg:text-[24px]">
				{description}
			</p>
			<a href={routes.earlyAccess} target="_blank" rel="noreferrer">
				<MainButton
					variant="white"
					text={button}
					className="mt-[30px] w-full lg:w-max"
				/>
			</a>
		</motion.div>
	</Container>
)
const Steps = () => (
	<Container className="grid gap-8 lg:grid-cols-3">
		<StepCard
			index={1}
			title="Set Rules"
			description="Choose a set of conditions to determine when your transaction can be executed."
			initial={{ opacity: 0, y: 20 }}
			whileInView={{ opacity: 1, y: 0 }}
			transition={{ duration: 0.2 }}
		>
			<SlidersHorizontal size={24} />
		</StepCard>
		<StepCard
			index={2}
			title="Define Actions"
			description="Bundle the actions that will automatically execute once all of your rules are satisfied."
			initial={{ opacity: 0, y: 20 }}
			whileInView={{ opacity: 1, y: 0 }}
			transition={{ duration: 0.2, delay: 0.2 }}
		>
			<ArrowDownWideNarrow size={24} />
		</StepCard>
		<StepCard
			index={3}
			title="Declare Intent"
			description="Sign a gasless signature to signal your intent of execution then sit back and relax."
			initial={{ opacity: 0, y: 20 }}
			whileInView={{ opacity: 1, y: 0 }}
			transition={{ duration: 0.4, delay: 0.4 }}
		>
			<FilePen size={24} />
		</StepCard>
	</Container>
)

const Templates = () => {
	const templates = [
		"Bridge to Optimism, Base, and Bera",
		"Top-Up Gearbox Loan Health Factor",
		"Bid on Noun with Pineapple Hat",
		"Buy Beta When Majors Move",
		"Fill Ethena Liquidty Cap to Limit",
		"Renew ENS Annually at Low Gas",
		"Enter Yearn When Above 65% APY",
		"Withdraw ETH:USDC Liquidity Rewards",
		"Rebalance Portfolio Monthly"
	]

	return (
		<Container className="my-[90px] flex-col items-center gap-4">
			<h2 className="text-center text-[28px] font-bold lg:w-[60%] lg:text-[64px] 2xl:w-[50%]">
				Start today with best-practice templates
			</h2>
			<p className="text-center text-[18px] font-light opacity-40 lg:w-[45%] lg:text-[24px]">
				No need to start from scratch. In just a few minutes, you can
				deploy a strategy that has been battle-tested by the Plug team
				and industry experts.
			</p>

			<div className="grid grid-cols-1 2xl:grid-cols-12">
				<a
					href={routes.earlyAccess}
					target="_blank"
					rel="noreferrer"
					className="mt-[40px] grid gap-4 md:grid-cols-2 2xl:col-start-4 2xl:col-end-10 2xl:grid-cols-3"
				>
					{templates.map((template, index) => (
						<ActionCard
							key={index}
							size="lg"
							color={
								Object.keys(colors)[
									index % Object.keys(colors).length
								] as keyof typeof colors
							}
							glow={false}
							title={template}
							initial={{ opacity: 0, y: 20 }}
							whileInView={{ opacity: 1, y: 0 }}
							transition={{
								duration: 0.2,
								delay: 0.1 * index
							}}
							className={index > 7 ? "hidden 2xl:flex" : ""}
						/>
					))}
				</a>
			</div>
		</Container>
	)
}

const Examples = () => (
	<>
		<Container className="mt-[90px] flex-col items-center gap-4">
			<h2 className="text-center text-[28px] font-bold lg:w-[60%] lg:text-[64px] 2xl:w-[45%]">
				A framework that lets you plug-and-play.
			</h2>
			<p className="text-center text-[18px] font-light opacity-40 lg:w-[45%] lg:text-[24px]">
				Declare intents with a state of the art no-code plug builder in
				seconds. Combine the power of top protocols into a single
				transaction that are all simulatenously settled.
			</p>
		</Container>

		<div
			className="border-green my-[45px] border-y-[2px] lg:my-[90px]"
			style={{
				borderImage: "linear-gradient(45deg, #00EF35, #93DF00) 1"
			}}
		>
			<div className="grid lg:grid-cols-12 lg:grid-rows-2">
				<div className="lg:col-span-6">
					<div className="p-4 py-8 lg:p-12 lg:pl-96">
						<div className="flex flex-col gap-4">
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
									<span className="rounded-md bg-[#00EF35]/10 p-2 py-1 text-[#00EF35]">
										1
									</span>{" "}
									time a{" "}
									<span className="rounded-md bg-[#00EF35]/10 p-2 py-1 text-[#00EF35]">
										year
									</span>{" "}
									.
								</p>
							</div>
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
									<span className="rounded-md bg-[#00EF35]/10 p-2 py-1 text-[#00EF35]">
										100
									</span>{" "}
									<span className="rounded-md bg-[#00EF35]/10 p-2 py-1 text-[#00EF35]">
										$USDC
									</span>{" "}
									to{" "}
									<span className="rounded-md bg-[#00EF35]/10 p-2 py-1 text-[#00EF35]">
										$ETH
									</span>{" "}
									.
								</p>
							</div>
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
									<span className="rounded-md bg-[#00EF35]/10 p-2 py-1 text-[#00EF35]">
										0.2
									</span>{" "}
									<span className="rounded-md bg-[#00EF35]/10 p-2 py-1 text-[#00EF35]">
										$ETH
									</span>{" "}
									or greater.
								</p>
							</div>
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
									<span className="rounded-md bg-[#00EF35]/10 p-2 py-1 text-[#00EF35]">
										nftchance.eth
									</span>{" "}
									.
								</p>
							</div>
						</div>
					</div>
				</div>
				<div
					className="flex border-t-[2px] lg:col-span-5 lg:col-start-7 lg:row-span-2 lg:border-l-[2px] lg:border-t-[0px] lg:pr-24"
					style={{
						borderImage:
							"linear-gradient(45deg, #00EF35, #93DF00) 1"
					}}
				>
					<div className="my-auto p-4 py-8 pr-0 lg:p-16 lg:pl-24">
						<div className="flex flex-col gap-4">
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
									<span className="rounded-md bg-[#00EF35]/10 p-2 py-1 text-[#00EF35]">
										1
									</span>{" "}
									time a{" "}
									<span className="rounded-md bg-[#00EF35]/10 p-2 py-1 text-[#00EF35]">
										day
									</span>{" "}
									.
								</p>
							</div>
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
									<span className="rounded-md bg-[#00EF35]/10 p-2 py-1 text-[#00EF35]">
										10
									</span>{" "}
									<span className="rounded-md bg-[#00EF35]/10 p-2 py-1 text-[#00EF35]">
										PM
									</span>{" "}
									<span className="rounded-md bg-[#00EF35]/10 p-2 py-1 text-[#00EF35]">
										UTC
									</span>{" "}
									.
								</p>
							</div>
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
									<span className="rounded-md bg-[#00EF35]/10 p-2 py-1 text-[#00EF35]">
										12/31/2024
									</span>{" "}
									.
								</p>
							</div>
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
									<span className="rounded-md bg-[#00EF35]/10 p-2 py-1 text-[#00EF35]">
										36,000
									</span>{" "}
									<span className="rounded-md bg-[#00EF35]/10 p-2 py-1 text-[#00EF35]">
										$USDC
									</span>{" "}
									to{" "}
									<span className="rounded-md bg-[#00EF35]/10 p-2 py-1 text-[#00EF35]">
										$ETH
									</span>
									.
								</p>
							</div>
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
									<span className="rounded-md bg-[#00EF35]/10 p-2 py-1 text-[#00EF35]">
										9
									</span>{" "}
									<span className="rounded-md bg-[#00EF35]/10 p-2 py-1 text-[#00EF35]">
										$ETH
									</span>{" "}
									.
								</p>
							</div>
							<div className="flex flex-row items-center gap-4">
								<motion.div
									className="flex h-6 w-6 items-center justify-center rounded-full bg-[#D9D9D9]/40"
									animate={{
										background: ["#EAEAEA", "#D9D9D9"]
									}}
									transition={{
										duration: 0.25,
										delay: 2.5,
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
								<p>Can bid without settling auction.</p>
							</div>
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
									<p className="text-xs">7</p>
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
									<span className="rounded-md bg-[#00EF35]/10 p-2 py-1 text-[#00EF35]">
										Pineapple Hat
									</span>{" "}
									.
								</p>
							</div>
						</div>
					</div>
				</div>
				<div
					className="border-t-[2px] lg:col-span-6"
					style={{
						borderImage:
							"linear-gradient(45deg, #00EF35, #93DF00) 1"
					}}
				>
					<div className="p-4 py-8 lg:p-16 lg:pl-96">
						<div className="flex flex-col gap-4">
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
									<span className="rounded-md bg-[#00EF35]/10 p-2 py-1 text-[#00EF35]">
										1
									</span>{" "}
									time a{" "}
									<span className="rounded-md bg-[#00EF35]/10 p-2 py-1 text-[#00EF35]">
										hour
									</span>{" "}
									.
								</p>
							</div>
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
									<span className="rounded-md bg-[#00EF35]/10 p-2 py-1 text-[#00EF35]">
										$USDC
									</span>{" "}
									pool is above{" "}
									<span className="rounded-md bg-[#00EF35]/10 p-2 py-1 text-[#00EF35]">
										72%
									</span>{" "}
									APY.
								</p>
							</div>
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
									<span className="rounded-md bg-[#00EF35]/10 p-2 py-1 text-[#00EF35]">
										10,000
									</span>{" "}
									<span className="rounded-md bg-[#00EF35]/10 p-2 py-1 text-[#00EF35]">
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

const Value = () => {
	const animation = {
		className: "w-full min-h-[360px]",
		initial: { opacity: 0, y: 20 },
		whileInView: { opacity: 1, y: 0 },
		transition: { duration: 0.2 }
	}

	const getDayAnimation = (delay: number, active: boolean = false) => ({
		style: {
			color: active === true ? "#FFFFFF" : "rgba(0,0,0,0.40)",
			borderColor: "#D9D9D9",
			padding: "4px 8px",
			height: "2rem"
		},
		whileInView: {
			background:
				active === true
					? [
							"linear-gradient(45deg, rgba(0,239,53,0.4), rgba(147,233,0,0.9))",
							"linear-gradient(45deg, rgba(0,239,53,1), rgba(147,233,0,1))",
							"linear-gradient(45deg, rgba(0,239,53,0.4), rgba(147,233,0,0.9))"
						]
					: [
							"rgba(217,217,217,0)",
							"rgba(217,217,217,0.4)",
							"rgba(217,217,217,0)"
						]
		},
		transition: {
			duration: 0.25,
			delay: 1 + delay,
			repeat: Infinity,
			repeatDelay: 7.5
		}
	})

	return (
		<Container className="mb-8">
			<div className="grid w-full gap-8 lg:grid-cols-2 lg:grid-rows-3 2xl:grid-cols-3">
				<InfoCard
					text={
						<>
							<Clock size={24} className="opacity-40" />
							<span>24/7 Execution</span>
						</>
					}
					description="Plug keeps your strategies running at all times. Whether you're at dinner, asleep, or on vacation, your transactions will be executed."
					{...animation}
					className={`${animation.className} lg:col-span-2 2xl:col-span-1`}
				>
					<div
						className="mx-8 grid h-36 gap-[2px]"
						style={{ gridTemplateColumns: "repeat(14, 1fr)" }}
					>
						{Array.from({ length: 14 }).map((_, index) => (
							<motion.div
								key={index}
								className="mt-auto w-full origin-bottom rounded-lg"
								style={{
									background:
										index === 13
											? "linear-gradient(45deg, #00EF35, #93DF00)"
											: "#D9D9D9"
								}}
								initial={{ height: 10 }}
								animate={{
									height: [
										10,
										20 * 2 ** (0.04 * (index / 2) * 8) +
											Math.random() * 40
									]
								}}
								transition={{
									duration: 0.25,
									delay: 0.25 * index,
									repeat: Infinity,
									repeatType: "reverse",
									repeatDelay: 3
								}}
							/>
						))}
					</div>
					<div className="absolute bottom-1/4 left-0 right-0 top-0 bg-gradient-to-bl from-[#FBFBFB]/0 to-[#FBFBFB]" />
					<div className="absolute bottom-0 left-0 right-0 top-1/2 bg-[#FBFBFB]" />
				</InfoCard>
				<InfoCard
					text={
						<>
							<CalendarClock size={24} className="opacity-40" />
							<span>Scheduled Transactions</span>
						</>
					}
					description="Have a specific period of time you want to execute your transaction? Use the same scheduling methods you are used to on calendars."
					{...animation}
				>
					<div className="ml-auto grid w-full grid-cols-7 grid-rows-4 text-xs">
						<div className="border-b-[1px] border-r-[1px]" />
						<div className="border-b-[1px] border-r-[1px]" />
						<div className="border-b-[1px] border-r-[1px]" />
						<motion.div
							className="border-b-[1px] border-r-[1px]"
							{...getDayAnimation(0)}
						>
							1
						</motion.div>
						<motion.div
							className="border-b-[1px] border-r-[1px]"
							{...getDayAnimation(0.25)}
						>
							2
						</motion.div>
						<motion.div
							className="border-b-[1px] border-r-[1px]"
							{...getDayAnimation(0.5)}
						>
							3
						</motion.div>
						<motion.div
							className="border-b-[1px]"
							{...getDayAnimation(0.75)}
						>
							4
						</motion.div>
						<motion.div
							className="border-b-[1px] border-r-[1px]"
							{...getDayAnimation(1)}
						>
							5
						</motion.div>
						<motion.div
							className="border-b-[1px] border-r-[1px]"
							{...getDayAnimation(1.25)}
						>
							6
						</motion.div>
						<motion.div
							className="border-b-[1px] border-r-[1px]"
							{...getDayAnimation(1.5, true)}
						>
							7
						</motion.div>
						<motion.div
							className="border-b-[1px] border-r-[1px]"
							{...getDayAnimation(1.75)}
						>
							8
						</motion.div>
						<motion.div
							className="border-b-[1px] border-r-[1px]"
							{...getDayAnimation(2)}
						>
							9
						</motion.div>
						<motion.div
							className="border-b-[1px] border-r-[1px]"
							{...getDayAnimation(2.25, true)}
						>
							10
						</motion.div>
						<motion.div
							className="border-b-[1px] border-[#D9D9D9]"
							{...getDayAnimation(2.5)}
						>
							11
						</motion.div>
						<motion.div
							className="border-b-[1px] border-r-[1px]"
							{...getDayAnimation(2.75)}
						>
							12
						</motion.div>
						<motion.div
							className="border-b-[1px] border-r-[1px]"
							{...getDayAnimation(3)}
						>
							13
						</motion.div>
						<motion.div
							className="border-b-[1px] border-r-[1px]"
							{...getDayAnimation(3.25, true)}
						>
							14
						</motion.div>
						<motion.div
							className="border-b-[1px] border-r-[1px]"
							{...getDayAnimation(3.5)}
						>
							15
						</motion.div>
						<motion.div
							className="border-b-[1px] border-r-[1px]"
							{...getDayAnimation(3.75)}
						>
							16
						</motion.div>
						<motion.div
							className="border-b-[1px] border-r-[1px]"
							{...getDayAnimation(4)}
						>
							17
						</motion.div>
						<motion.div
							className="border-b-[1px]"
							{...getDayAnimation(4.25)}
						>
							18
						</motion.div>
						<motion.div
							className="border-b-[1px] border-r-[1px]"
							{...getDayAnimation(4.5)}
						>
							19
						</motion.div>
						<motion.div
							className="border-b-[1px] border-r-[1px]"
							{...getDayAnimation(4.75)}
						>
							20
						</motion.div>
						<motion.div
							className="border-b-[1px] border-r-[1px] border-[#D9D9D9]"
							{...getDayAnimation(5, true)}
						>
							21
						</motion.div>
						<motion.div
							className="border-b-[1px] border-r-[1px]"
							{...getDayAnimation(5.25)}
						>
							22
						</motion.div>
						<motion.div
							className="border-b-[1px] border-r-[1px]"
							{...getDayAnimation(5.5)}
						>
							23
						</motion.div>
						<motion.div
							className="border-b-[1px] border-r-[1px]"
							{...getDayAnimation(5.75, true)}
						>
							24
						</motion.div>
						<motion.div
							className="border-b-[1px]"
							{...getDayAnimation(6)}
						>
							25
						</motion.div>
						<motion.div
							className="border-r-[1px]"
							{...getDayAnimation(6.25)}
						>
							26
						</motion.div>
						<motion.div
							className="border-r-[1px]"
							{...getDayAnimation(6.5)}
						>
							27
						</motion.div>
						<motion.div
							className="border-r-[1px]"
							{...getDayAnimation(6.75, true)}
						>
							28
						</motion.div>
						<motion.div
							className="border-r-[1px]"
							{...getDayAnimation(7)}
						>
							29
						</motion.div>
						<motion.div
							className="border-r-[1px]"
							{...getDayAnimation(7.25)}
						>
							30
						</motion.div>
						<div className="border-r-[1px]" />
					</div>
				</InfoCard>
				<InfoCard
					text={
						<>
							<FileStack size={24} className="opacity-40" />
							<span>Multichain Signatures</span>
						</>
					}
					description="Declare multiple actions across multiple chains with a single gasless signature. Stop signing the same signature over and over."
					{...animation}
				>
					<div className="grayscale-100 ml-16 mt-16 flex w-[100%] flex-row saturate-0 filter">
						<motion.div
							whileInView={{ x: [20, -10] }}
							transition={{
								duration: 0.5,
								repeat: Infinity,
								repeatType: "reverse",
								repeatDelay: 2
							}}
						>
							<Image
								src="/blockchain/ethereum.png"
								alt="Ethereum"
								width={120}
								height={120}
							/>
						</motion.div>
						<motion.div
							whileInView={{ x: [-10, -30] }}
							transition={{
								duration: 0.5,
								delay: 0.25,
								repeat: Infinity,
								repeatType: "reverse",
								repeatDelay: 2
							}}
						>
							<Image
								src="/blockchain/base.png"
								alt="Base"
								width={120}
								height={120}
							/>
						</motion.div>
						<motion.div
							whileInView={{ x: [-30, -50] }}
							transition={{
								duration: 0.5,
								delay: 0.5,
								repeat: Infinity,
								repeatType: "reverse",
								repeatDelay: 2
							}}
						>
							<Image
								src="/blockchain/optimism.png"
								alt="Optimism"
								width={120}
								height={120}
							/>
						</motion.div>
						<motion.div
							whileInView={{ x: [-50, -80] }}
							transition={{
								duration: 0.5,
								delay: 0.75,
								repeat: Infinity,
								repeatType: "reverse",
								repeatDelay: 2
							}}
						>
							<Image
								src="/blockchain/bera.png"
								alt="Bera"
								width={120}
								height={120}
							/>
						</motion.div>
						<motion.div
							whileInView={{ x: [-80, -110] }}
							transition={{
								duration: 0.5,
								delay: 1,
								repeat: Infinity,
								repeatType: "reverse",
								repeatDelay: 2
							}}
						>
							<Image
								src="/blockchain/arbitrum.png"
								alt="Arbitrum"
								width={120}
								height={120}
							/>
						</motion.div>
						<motion.div
							whileInView={{ x: [-110, -140] }}
							transition={{
								duration: 0.5,
								delay: 1.25,
								repeat: Infinity,
								repeatType: "reverse",
								repeatDelay: 2
							}}
						>
							<Image
								src="/blockchain/polygon.png"
								alt="Polygon"
								width={120}
								height={120}
							/>
						</motion.div>
					</div>
				</InfoCard>
				<InfoCard
					text={
						<>
							<Ruler size={24} className="opacity-40" />
							<span>Atomic Rules</span>
						</>
					}
					description="Rules of execution are embedded into your transaction and validated onchain during simulation and execution to enable completely trustless automation."
					{...animation}
				>
					<svg
						width="479"
						height="126"
						viewBox="0 0 479 126"
						fill="none"
						xmlns="http://www.w3.org/2000/svg"
						className="mt-8 lg:mx-auto lg:mt-8 lg:w-[150%] 2xl:mx-0 2xl:w-[100%]"
					>
						<motion.path
							d="M0 18.9805H112.217C112.217 18.9805 119.584 19.2795 123.438 22.3844C129.266 27.0795 134.228 40.6805 134.228 40.6805C134.228 40.6805 139.19 54.2814 145.018 58.9765C148.872 62.0814 156.24 62.3805 156.24 62.3805"
							strokeWidth="2"
							strokeDasharray="4 4"
							animate={{
								stroke: [
									"#D9D9D9",
									"#00EF35",
									"#D9D9D9",
									"#FF0000",
									"#D9D9D9",
									"#00EF35",
									"#D9D9D9"
								],
								strokeDashoffset: [
									0, -4, -8, -12, -16, -20, -24, -28
								]
							}}
							transition={{
								duration: 2,
								repeat: Infinity
							}}
						/>
						<motion.path
							d="M478.02 18.9805H366.248C366.248 18.9805 358.91 19.2795 355.071 22.3844C349.266 27.0795 344.324 40.6805 344.324 40.6805C344.324 40.6805 339.382 54.2814 333.577 58.9765C329.738 62.0814 322.4 62.3805 322.4 62.3805"
							strokeWidth="2"
							strokeDasharray="4 4"
							animate={{
								stroke: [
									"#D9D9D9",
									"#00EF35",
									"#D9D9D9",
									"#FF0000",
									"#D9D9D9",
									"#FF0000",
									"#D9D9D9"
								],
								strokeDashoffset: [0, 4, 8, 12, 16, 20, 24, 28]
							}}
							transition={{
								duration: 2,
								repeat: Infinity
							}}
						/>
						<motion.path
							d="M0 1H131.365C131.365 1 139.989 1.42296 144.501 5.81412C151.324 12.4544 157.132 31.69 157.132 31.69C157.132 31.69 162.941 50.9256 169.764 57.5659C174.275 61.957 182.9 62.38 182.9 62.38"
							strokeWidth="2"
							strokeDasharray="4 4"
							animate={{
								stroke: [
									"#D9D9D9",
									"#00EF35",
									"#D9D9D9",
									"#FF0000",
									"#D9D9D9",
									"#00EF35",
									"#D9D9D9"
								],
								strokeDashoffset: [
									0, -4, -8, -12, -16, -20, -24, -28
								]
							}}
							transition={{
								duration: 2,
								repeat: Infinity
							}}
						/>
						<motion.path
							d="M478.02 1H346.655C346.655 1 338.03 1.42296 333.518 5.81412C326.696 12.4544 320.887 31.69 320.887 31.69C320.887 31.69 315.079 50.9256 308.256 57.5659C303.744 61.957 295.12 62.38 295.12 62.38"
							strokeWidth="2"
							strokeDasharray="4 4"
							animate={{
								stroke: [
									"#D9D9D9",
									"#00EF35",
									"#D9D9D9",
									"#FF0000",
									"#D9D9D9",
									"#FF0000",
									"#D9D9D9"
								],
								strokeDashoffset: [0, 4, 8, 12, 16, 20, 24, 28]
							}}
							transition={{
								duration: 2,
								repeat: Infinity
							}}
						/>
						<motion.path
							d="M0 107.02H112.217C112.217 107.02 119.584 106.712 123.438 103.518C129.266 98.6891 134.228 84.6995 134.228 84.6995C134.228 84.6995 139.19 70.71 145.018 65.8807C148.872 62.6871 156.24 62.3795 156.24 62.3795"
							strokeWidth="2"
							strokeDasharray="4 4"
							animate={{
								stroke: [
									"#D9D9D9",
									"#00EF35",
									"#D9D9D9",
									"#00EF35",
									"#D9D9D9",
									"#00EF35",
									"#D9D9D9"
								],
								strokeDashoffset: [
									0, -4, -8, -12, -16, -20, -24, -28
								]
							}}
							transition={{
								duration: 2,
								repeat: Infinity
							}}
						/>
						<motion.path
							d="M478.02 107.02H366.248C366.248 107.02 358.91 106.712 355.071 103.518C349.266 98.6891 344.324 84.6995 344.324 84.6995C344.324 84.6995 339.382 70.71 333.577 65.8807C329.738 62.6871 322.4 62.3795 322.4 62.3795"
							strokeWidth="2"
							strokeDasharray="4 4"
							animate={{
								stroke: [
									"#D9D9D9",
									"#00EF35",
									"#D9D9D9",
									"#FF0000",
									"#D9D9D9",
									"#FF0000",
									"#D9D9D9"
								],
								strokeDashoffset: [0, 4, 8, 12, 16, 20, 24, 28]
							}}
							transition={{
								duration: 2,
								repeat: Infinity
							}}
						/>
						<motion.path
							d="M0 125H131.365C131.365 125 139.989 124.568 144.501 120.089C151.324 113.314 157.132 93.69 157.132 93.69C157.132 93.69 162.941 74.0658 169.764 67.2914C174.275 62.8115 182.9 62.38 182.9 62.38"
							strokeWidth="2"
							strokeDasharray="4 4"
							animate={{
								stroke: [
									"#D9D9D9",
									"#00EF35",
									"#D9D9D9",
									"#FF0000",
									"#D9D9D9",
									"#FF0000",
									"#D9D9D9"
								],
								strokeDashoffset: [
									0, -4, -8, -12, -16, -20, -24, -28
								]
							}}
							transition={{
								duration: 2,
								repeat: Infinity
							}}
						/>
						<motion.path
							d="M478.02 125H346.655C346.655 125 338.03 124.568 333.518 120.089C326.696 113.314 320.887 93.69 320.887 93.69C320.887 93.69 315.079 74.0658 308.256 67.2914C303.744 62.8115 295.12 62.38 295.12 62.38"
							strokeWidth="2"
							strokeDasharray="4 4"
							animate={{
								stroke: [
									"#D9D9D9",
									"#00EF35",
									"#D9D9D9",
									"#FF0000",
									"#D9D9D9",
									"#FF0000",
									"#D9D9D9"
								],
								strokeDashoffset: [0, 4, 8, 12, 16, 20, 24, 28]
							}}
							transition={{
								duration: 2,
								repeat: Infinity
							}}
						/>
						<motion.line
							x1="156.238"
							y1="62.1396"
							x2="-0.00170898"
							y2="62.1396"
							strokeWidth="2"
							strokeDasharray="4 4"
							animate={{
								stroke: [
									"#D9D9D9",
									"#00EF35",
									"#D9D9D9",
									"#FF0000",
									"#D9D9D9",
									"#00EF35",
									"#D9D9D9"
								],
								strokeDashoffset: [0, 4, 8, 12, 16, 20, 24, 28]
							}}
							transition={{
								duration: 2,
								repeat: Infinity
							}}
						/>
						<motion.line
							y1="-1"
							x2="155.62"
							y2="-1"
							transform="matrix(1 0 0 -1 322.398 61.1396)"
							strokeWidth="2"
							strokeDasharray="4 4"
							animate={{
								stroke: [
									"#D9D9D9",
									"#00EF35",
									"#D9D9D9",
									"#FF0000",
									"#D9D9D9",
									"#FF0000",
									"#D9D9D9"
								],
								strokeDashoffset: [
									0, -4, -8, -12, -16, -20, -24, -28
								]
							}}
							transition={{
								duration: 2,
								repeat: Infinity
							}}
						/>
						<motion.line
							x1="182.898"
							y1="62.1396"
							x2="156.238"
							y2="62.1396"
							strokeWidth="2"
							strokeDasharray="4 4"
							animate={{
								stroke: [
									"#D9D9D9",
									"#00EF35",
									"#D9D9D9",
									"#FF0000",
									"#D9D9D9",
									"#00EF35",
									"#D9D9D9"
								],
								strokeDashoffset: [0, 4, 8, 12, 16, 20, 24, 28]
							}}
							transition={{
								duration: 2,
								repeat: Infinity
							}}
						/>
						<motion.line
							y1="-1"
							x2="27.28"
							y2="-1"
							transform="matrix(1 0 0 -1 295.121 61.1396)"
							strokeWidth="2"
							strokeDasharray="4 4"
							animate={{
								stroke: [
									"#D9D9D9",
									"#00EF35",
									"#D9D9D9",
									"#FF0000",
									"#D9D9D9",
									"#FF0000",
									"#D9D9D9"
								],
								strokeDashoffset: [
									0, -4, -8, -12, -16, -20, -24, -28
								]
							}}
							transition={{
								duration: 2,
								repeat: Infinity
							}}
						/>
						<motion.line
							x1="295.121"
							y1="62.1396"
							x2="182.901"
							y2="62.1396"
							strokeWidth="2"
							strokeDasharray="4 4"
							animate={{
								stroke: [
									"#D9D9D9",
									"#00EF35",
									"#D9D9D9",
									"#FF0000",
									"#D9D9D9",
									"#FF0000",
									"#D9D9D9"
								],
								strokeDashoffset: [0, 4, 8, 12, 16, 20, 24, 28]
							}}
							transition={{
								duration: 2,
								repeat: Infinity
							}}
						/>
					</svg>
				</InfoCard>

				<InfoCard
					text={
						<>
							<TestTubeDiagonal
								size={24}
								className="opacity-40"
							/>
							<span>Constant Simulation</span>
						</>
					}
					description="Plug constantly simulates your transaction to check if all the conditions are met. When they are, the transaction is executed."
					{...animation}
				>
					<motion.div
						className="mx-auto ml-auto mt-16 flex w-[420px] items-center gap-6 rounded-lg bg-white px-6 py-2"
						animate={{
							x: [0, 4, -4, 2, -1, 3, -3, 4, -2, 0],
							y: [0, 4, -4, 2, -1, 3, -3, 4, -2, 0]
						}}
						transition={{
							duration: 0.25,
							delay: 0.25,
							repeat: Infinity,
							repeatType: "reverse",
							repeatDelay: 8.25
						}}
					>
						<motion.div
							className="relative flex h-8 w-8 items-center justify-center rounded-full"
							animate={{
								background: [
									"rgba(0,239,53,0.1)",
									"rgba(255,81,84,0.1)"
								]
							}}
							transition={{
								duration: 0.25,
								repeat: Infinity,
								repeatType: "reverse",
								repeatDelay: 4
							}}
						>
							<motion.span
								animate={{ color: ["#00EF35", "#FF5154"] }}
								transition={{
									duration: 0.25,
									repeat: Infinity,
									repeatType: "reverse",
									repeatDelay: 4
								}}
							>
								<motion.span
									className="absolute left-1/4 top-1/4"
									animate={{ opacity: [1, 0] }}
									transition={{
										duration: 0.25,
										repeat: Infinity,
										repeatType: "reverse",
										repeatDelay: 4
									}}
								>
									<Check size={18} />
								</motion.span>
								<motion.span
									className="absolute left-1/4 top-1/4"
									animate={{ opacity: [0, 1] }}
									transition={{
										duration: 0.25,
										repeat: Infinity,
										repeatType: "reverse",
										repeatDelay: 4
									}}
								>
									<X size={18} />
								</motion.span>
							</motion.span>
						</motion.div>
						<h3 className="flex flex-col gap-1">
							<span className="text-xl font-bold">
								Bid in Nouns Auction
							</span>
							<div className="flex w-full flex-row items-center gap-2 text-lg">
								<Image
									src={`/wallets/madison.png`}
									alt="NFT Chance"
									width={18}
									height={18}
									className="h-4 w-4 rounded-full"
								/>
								<span className="opacity-40">madison.eth</span>
							</div>
						</h3>
						<h4 className="mb-auto ml-auto opacity-40">
							{Math.floor(Math.random() ** 1.2) + 1} secs. ago
						</h4>
					</motion.div>
				</InfoCard>
				<InfoCard
					text={
						<>
							<ShieldPlus size={24} className="opacity-40" />
							<span>Upgradeable Accounts</span>
						</>
					}
					description="As new features and versions are released you can upgrade your account without hassle, paying additional fees or needing to re-deploy."
					{...animation}
				>
					<Version />
				</InfoCard>
				<InfoCard
					text={
						<>
							<RotateCw size={24} className="opacity-40" />
							<span>Recurring Outcomes</span>
						</>
					}
					description="With Plug your signed intents can be reused to run again when the conditions are met without needing to sign a new transaction."
					{...animation}
				>
					<div
						className="ml-[-4px] mt-[-9px] grid w-max grid-rows-3 gap-[2px]"
						style={{ gridTemplateColumns: "repeat(28, 1fr)" }}
					>
						{Array.from({ length: 28 * 8 }).map((_, index) => {
							const background =
								Math.random() < 0.5
									? "#D9D9D9"
									: "linear-gradient(45deg, #00EF35, #93DF00)"
							return (
								<motion.div
									key={index}
									className="h-5 w-5 rounded-[2px]"
									style={{ background }}
									initial={{ opacity: 1 }}
									animate={{ opacity: [1, 0] }}
									transition={{
										duration: 0.5,
										delay: 0.05 * index,
										repeat: Infinity,
										repeatType: "reverse"
									}}
								/>
							)
						})}
					</div>
				</InfoCard>
				<InfoCard
					text={
						<>
							<Rotate3d size={24} className="opacity-40" />
							<span>Transferrable Accounts</span>
						</>
					}
					description="Transfer your account to a different wallet or sell it on any major NFT marketplace. Your account is yours to do with as you please."
					{...animation}
				>
					<div className="flex w-full flex-row items-center">
						<motion.div
							className="ml-[-65px] mt-4 flex h-[130px] w-[130px] items-center justify-center rounded-full border-[2px] border-dashed"
							animate={{
								borderColor: ["#D9D9D9", "#00EF35"]
							}}
							transition={{
								duration: 1,
								repeat: Infinity,
								repeatDelay: 1,
								repeatType: "reverse"
							}}
						>
							<motion.div className="grayscale-100 relative h-full w-full saturate-0 filter">
								<motion.div
									className="absolute bottom-0 left-[3px] right-0 top-[3px]"
									animate={{ opacity: [0, 1] }}
									transition={{
										duration: 0.25,
										repeat: Infinity,
										repeatType: "reverse",
										repeatDelay: 3.75
									}}
								>
									<Image
										src="/wallets/nftchance.png"
										alt="NFT Chance"
										width={120}
										height={120}
									/>
								</motion.div>
								<motion.div
									className="absolute bottom-0 left-[3px] right-0 top-[3px] h-[120px] w-[120px]"
									animate={{ opacity: [1, 0] }}
									transition={{
										duration: 0.25,
										delay: 0,
										repeat: Infinity,
										repeatType: "reverse",
										repeatDelay: 3.75
									}}
								>
									<Image
										src="/wallets/deeze.png"
										alt="Deeze"
										width={120}
										height={120}
									/>
								</motion.div>
							</motion.div>
						</motion.div>
						<div
							className="relative mt-4"
							style={{ width: "calc(100% - 130px)" }}
						>
							<motion.div
								className="absolute bottom-0 left-0 right-0 top-0 h-[2px]"
								style={{
									background:
										"linear-gradient(90deg, #00EF35 25%, #93DF00 50%, transparent 50%)",
									backgroundRepeat: "repeat",
									backgroundSize: "4px 6px"
								}}
							/>
							<motion.div
								className="absolute top-[-5px] h-3 w-3 rounded-full"
								style={{
									background:
										"linear-gradient(45deg, #00EF35, #93DF00)"
								}}
								animate={{
									left: ["97%", "0%"]
								}}
								transition={{
									duration: 2,
									repeat: Infinity,
									repeatType: "reverse",
									ease: "linear"
								}}
							/>
						</div>
						<motion.div
							className="mr-[-65px] mt-4 flex h-[130px] w-[130px] items-center justify-center rounded-full border-[2px] border-dashed border-[#D9D9D9]"
							animate={{
								borderColor: ["#D9D9D9", "#00EF35"]
							}}
							transition={{
								duration: 1,
								delay: 2,
								repeat: Infinity,
								repeatDelay: 1,
								repeatType: "reverse"
							}}
						>
							<motion.div className="grayscale-100 relative h-full w-full saturate-0 filter">
								<motion.div
									className="absolute bottom-0 left-[3px] right-0 top-[3px]"
									animate={{ opacity: [0, 1] }}
									transition={{
										duration: 0.25,
										delay: 2,
										repeat: Infinity,
										repeatType: "reverse",
										repeatDelay: 3.75
									}}
								>
									<Image
										src="/wallets/danner.png"
										alt="NFT Chance"
										width={120}
										height={120}
									/>
								</motion.div>
								<motion.div
									className="absolute bottom-0 left-[3px] right-0 top-[3px] h-[120px] w-[120px]"
									animate={{ opacity: [1, 0] }}
									transition={{
										duration: 0.25,
										delay: 2,
										repeat: Infinity,
										repeatType: "reverse",
										repeatDelay: 3.75
									}}
								>
									<Image
										src="/wallets/blob.png"
										alt="Blob"
										width={120}
										height={120}
									/>
								</motion.div>
							</motion.div>
						</motion.div>
					</div>
				</InfoCard>
				<InfoCard
					text={
						<>
							<Puzzle size={24} className="opacity-40" />
							<span>Generalized Intents</span>
						</>
					}
					description="Plug is designed to support everything right out of the box. Whether you want to swap, lend, borrow, stake, or farm... Plug has you covered."
					{...animation}
				>
					<svg
						width="320"
						height="256"
						viewBox="0 0 420 256"
						fill="none"
						xmlns="http://www.w3.org/2000/svg"
						className="ml-auto mr-[-30px] mt-[-30px]"
					>
						<defs>
							<linearGradient
								id="gradient"
								x1="0%"
								y1="100%"
								x2="100%"
								y2="0%"
							>
								<stop
									offset="0%"
									style={{
										stopColor: "#00EF35",
										stopOpacity: 1
									}}
								/>
								<stop
									offset="100%"
									style={{
										stopColor: "#93DF00",
										stopOpacity: 1
									}}
								/>
							</linearGradient>
						</defs>
						<motion.path
							fill-rule="evenodd"
							clip-rule="evenodd"
							d="M99.2387 81.7617C103.74 82.5367 105.587 82.4107 110.222 80.7897C118.898 77.7547 121.154 87.9027 117.764 94.0617C113.084 102.564 105.743 90.0157 100.262 95.2617C92.7817 102.421 96.1837 111.636 99.3367 120.575C99.7247 120.495 100.112 120.415 100.499 120.336C108.449 118.711 116.769 117.166 124.807 119.136C129.33 120.245 132.786 121.659 129.625 126.592C126.576 131.35 127.067 133.621 133.062 134.996C136.615 135.811 156.03 134.753 149.561 128.121C140.964 119.308 159.277 117.767 164.893 118.015C169.728 118.228 174.462 119.156 179.187 120.148C176.079 111.343 172.901 102.307 180.262 95.2617C185.743 90.0157 193.084 102.565 197.764 94.0617C201.154 87.9027 198.898 77.7547 190.222 80.7897C185.586 82.4117 183.74 82.5367 179.239 81.7617C177.453 79.0177 176.598 75.8447 176.657 72.5797C176.733 68.3137 178.048 64.2407 179.444 60.2047C171.912 58.6217 164.231 57.1417 156.57 58.5087C153.769 59.0087 146.71 59.9847 147.031 63.9647C147.302 67.3167 151.595 68.2167 150.77 71.7757C145.699 75.4657 131.148 78.8637 127.292 71.6137C125.948 69.0867 131.675 64.9887 130.717 61.2827C126.989 59.4577 122.92 58.4877 118.784 58.2597C112.291 57.9007 105.838 59.1267 99.4067 60.4357C98.0427 64.4087 96.7727 68.4207 96.6897 72.6157C96.6257 75.8657 97.4607 79.0277 99.2387 81.7617Z"
							stroke-linecap="round"
							stroke="#FBFBFB"
							fill="url(#gradient)"
							animate={{
								opacity: [0, 1]
							}}
							transition={{
								duration: 0.5,
								delay: Math.random() * (5 - 0.25) + 0.25,
								repeat: Infinity,
								repeatType: "reverse",
								repeatDelay: 1
							}}
						/>
						<motion.path
							fill-rule="evenodd"
							clip-rule="evenodd"
							d="M307.307 60.6966C305.666 56.1556 311.805 54.7816 310.772 50.2286C307.709 47.7126 302.878 47.0376 298.978 46.9106C295.663 46.8026 289.19 46.9606 287.659 50.5336C286.259 53.7996 291.949 56.6456 290.702 60.7016C281.336 62.0226 271.876 62.5846 262.42 62.6166C262.978 71.2326 263.246 82.8576 258.044 82.9526C253.309 83.0386 248 78.1696 243.94 81.4406C241.06 83.7606 240.406 87.0786 241.357 90.8846C242.503 95.4786 245.094 97.4206 249.793 96.7906C251.897 94.7826 257.21 92.5026 259.703 94.9836C265.682 103.527 262.056 111.944 259.16 120.516C267.672 118.732 276.171 117.053 284.762 119.141C289.295 120.243 292.803 121.634 289.627 126.591C287.336 130.166 286.01 132.811 291.114 134.459C294.798 135.648 316.41 135.14 309.562 128.121C301.564 119.923 317.083 118.006 322.769 117.995C328.313 117.984 333.735 119.03 339.144 120.162C336.039 111.357 332.878 102.329 340.263 95.2616C345.744 90.0156 353.085 102.565 357.765 94.0616C361.155 87.9026 358.899 77.7546 350.223 80.7896C345.587 82.4116 343.741 82.5366 339.24 81.7616C337.463 79.0286 336.627 75.8656 336.691 72.6156C336.774 68.4726 338.013 64.5086 339.358 60.5836C337.681 60.9346 336.009 61.2856 334.339 61.6086C325.345 63.3496 316.006 64.2256 307.307 60.6966Z"
							stroke-linecap="round"
							stroke="#FBFBFB"
							fill="#D9D9D9"
							animate={{
								opacity: [0, 1]
							}}
							transition={{
								duration: 0.5,
								delay: Math.random() * (5 - 0.25) + 0.25,
								repeat: Infinity,
								repeatType: "reverse",
								repeatDelay: 1
							}}
						/>
						<motion.path
							fill-rule="evenodd"
							clip-rule="evenodd"
							d="M22.2318 76.6984C21.2798 81.7664 19.2058 83.8914 14.0118 82.2884C11.8438 81.6194 9.77081 80.5974 7.50381 80.2704C4.37081 79.8184 4.92181 81.5624 2.29881 82.3194C-1.06019 88.9524 2.52282 100.125 11.2268 95.8284C15.7818 93.5794 19.2498 92.9704 21.5498 98.2844C22.9018 101.409 23.2168 104.82 22.8048 108.178C22.2948 112.331 20.8628 116.245 19.4958 120.179C23.6338 121.048 27.7658 121.959 31.9548 122.483C36.1008 123.001 40.3208 123.126 44.4548 122.43C49.4788 121.584 52.5958 120.295 49.2838 115.033C47.0438 111.473 46.9378 110.678 50.2088 108.225C55.1168 106.425 60.5908 106.581 65.5938 107.892C70.4728 109.17 72.6318 111.205 68.8148 115.71C64.8428 120.399 68.8698 121.65 73.5348 122.535C82.1608 124.17 90.8708 122.318 99.3358 120.575C96.1828 111.636 92.7798 102.421 100.261 95.2624C105.742 90.0164 113.083 102.565 117.763 94.0624C121.153 87.9034 118.897 77.7554 110.221 80.7904C105.585 82.4124 103.739 82.5384 99.2378 81.7624C97.4598 79.0294 96.6248 75.8664 96.6888 72.6164C96.7718 68.4214 98.0418 64.4094 99.4058 60.4364C88.6708 62.6214 77.9978 65.0384 67.3048 60.6974C65.5858 55.9384 73.6968 53.5594 69.7188 49.2804C65.6398 47.5884 61.1588 46.7014 56.7388 46.9314C53.8628 47.0804 48.9418 47.5374 47.6568 50.5334C46.2578 53.7994 51.9458 56.6454 50.7008 60.7014C40.2598 64.8454 29.7838 62.3364 19.2568 60.1524C21.2848 65.5624 23.2988 71.0124 22.2318 76.6984Z"
							stroke-linecap="round"
							stroke="#FBFBFB"
							fill="url(#gradient)"
							animate={{
								opacity: [0, 1]
							}}
							transition={{
								duration: 0.5,
								delay: Math.random() * (5 - 0.25) + 0.25,
								repeat: Infinity,
								repeatType: "reverse",
								repeatDelay: 1
							}}
						/>
						<motion.path
							fill-rule="evenodd"
							clip-rule="evenodd"
							d="M179.239 81.7624C183.74 82.5374 185.587 82.4114 190.222 80.7904C198.898 77.7554 201.154 87.9034 197.764 94.0624C193.084 102.564 185.743 90.0164 180.262 95.2624C172.901 102.307 176.078 111.344 179.187 120.148C182.73 120.892 186.267 121.671 189.837 122.217C193.967 122.849 198.172 123.169 202.339 122.726C207.66 122.162 213.195 121.249 209.284 115.032C207.044 111.471 206.938 110.678 210.209 108.224C214.417 106.681 219.054 106.583 223.426 107.404C227.236 108.119 232.874 109.007 230.037 113.955C228.861 116.005 227.021 116.523 227.051 119.047C227.085 121.901 231.298 122.107 233.586 122.541C242.191 124.172 250.681 122.292 259.158 120.515C262.054 111.943 265.679 103.527 259.701 94.9834C257.208 92.5024 251.895 94.7814 249.791 96.7904C245.093 97.4194 242.502 95.4774 241.355 90.8844C240.404 87.0774 241.058 83.7604 243.938 81.4404C247.999 78.1694 253.307 83.0384 258.042 82.9524C263.244 82.8574 262.976 71.2324 262.418 62.6164C261.332 62.6194 260.247 62.6204 259.161 62.6104C254.745 62.5684 227.096 62.9204 227.05 59.0484C227.006 55.2244 233.383 53.2224 229.719 49.2794C225.64 47.5874 221.159 46.7004 216.739 46.9304C213.168 47.1164 207.865 47.7234 207.228 51.9774C208.95 54.7104 213.514 59.9034 208.608 61.4384C200.411 64.0034 191.83 62.7564 183.597 61.0734C182.218 60.7914 180.832 60.4964 179.443 60.2044C178.047 64.2394 176.732 68.3134 176.656 72.5794C176.599 75.8454 177.453 79.0184 179.239 81.7624Z"
							stroke-linecap="round"
							stroke="#FBFBFB"
							fill="url(#gradient)"
							animate={{
								opacity: [0, 1]
							}}
							transition={{
								duration: 0.5,
								delay: Math.random() * (5 - 0.25) + 0.25,
								repeat: Infinity,
								repeatType: "reverse",
								repeatDelay: 1
							}}
						/>
						<motion.path
							fill-rule="evenodd"
							clip-rule="evenodd"
							d="M387.478 65.4782C388.677 67.8492 391.467 68.7622 390.768 71.7762C388.532 73.4032 385.943 74.5072 383.245 75.0912C378.662 76.0832 371.808 76.0972 368.228 72.7222C365.042 69.7202 371.796 65.4672 370.715 61.2832C366.975 59.4682 362.899 58.5012 358.755 58.2772C352.206 57.9222 345.753 59.2452 339.355 60.5842C338.01 64.5092 336.771 68.4742 336.688 72.6162C336.624 75.8662 337.459 79.0292 339.237 81.7622C343.738 82.5372 345.585 82.4112 350.22 80.7902C358.896 77.7552 361.152 87.9032 357.762 94.0622C353.082 102.564 345.741 90.0162 340.26 95.2622C332.875 102.328 336.035 111.357 339.141 120.162C342.711 120.91 346.275 121.694 349.865 122.24C353.499 122.792 372.261 124.551 370.677 118.306C369.956 114.503 366.292 113.192 368.25 109.319C371.991 108.394 374.801 106.775 378.977 106.911C382.877 107.038 387.707 107.714 390.771 110.229C391.804 114.783 385.666 116.156 387.306 120.697C393.133 124.113 407.797 126.949 419.499 125.666V60.4832C406.342 59.1712 384.413 59.4182 387.478 65.4782Z"
							stroke-linecap="round"
							stroke="#FBFBFB"
							fill="url(#gradient)"
							animate={{
								opacity: [0, 1]
							}}
							transition={{
								duration: 0.5,
								delay: Math.random() * (5 - 0.25) + 0.25,
								repeat: Infinity,
								repeatType: "reverse",
								repeatDelay: 1
							}}
						/>
						<motion.path
							fill-rule="evenodd"
							clip-rule="evenodd"
							d="M399.666 182.34C396.366 181.887 393.082 181.319 389.772 180.939C381.832 180.833 391.832 173.166 387.7 168.522C382.851 166.889 368.38 164.334 367.228 171.979C368.503 174.004 370.217 175.875 370.677 178.306C371.636 182.083 368.033 181.081 365.498 181.367C360.266 181.957 355.01 182.322 349.748 182.502C346.625 182.609 343.5 182.644 340.376 182.629C340.405 183.891 340.527 185.153 340.811 186.418C342.564 194.244 345.113 205.185 332.935 202.027C327.521 200.624 327.362 199.855 322.297 202.284C320.285 205.388 320.274 209.424 322.055 212.632C326.916 221.39 332.661 212.568 339.009 214.306C343.922 220.556 343.386 227.927 341.165 235.12C340.647 236.799 340.046 238.489 339.454 240.192C347.749 241.937 356.062 243.843 364.454 242.43C368.33 241.778 372.282 241.431 370.134 236.595C368.77 233.525 366.542 232.697 368.25 229.319C375.25 227.59 380.306 226.032 387.699 228.522C391.868 229.927 391.469 232.578 388.815 235.71C385.13 240.059 387.765 240.636 392.241 241.253C398.666 242.138 410.679 242.126 419.499 241V182.55C412.453 183.01 403.585 182.878 399.666 182.34Z"
							stroke-linecap="round"
							stroke="#FBFBFB"
							fill="#D9D9D9"
							animate={{
								opacity: [0, 1]
							}}
							transition={{
								duration: 0.5,
								delay: Math.random() * (5 - 0.25) + 0.25,
								repeat: Infinity,
								repeatType: "reverse",
								repeatDelay: 1
							}}
						/>
						<motion.path
							fill-rule="evenodd"
							clip-rule="evenodd"
							d="M95.1911 22.465C90.5851 21.831 88.2281 18.844 83.9381 21.438C80.5691 23.475 80.0391 27.656 81.4351 31.298C82.4161 33.859 84.4971 36.221 87.2691 36.819C91.1481 37.654 92.9851 34.05 96.5331 34.015C104.622 33.936 103.04 48.581 101.674 53.501C101.03 55.818 100.203 58.12 99.4081 60.436C105.839 59.127 112.292 57.901 118.785 58.26C122.921 58.489 126.99 59.459 130.718 61.283C131.676 64.989 125.949 69.087 127.293 71.614C131.148 78.864 145.699 75.466 150.771 71.776C151.596 68.218 147.302 67.317 147.032 63.965C146.711 59.986 153.77 59.009 156.571 58.509C164.232 57.142 171.913 58.622 179.445 60.205C180.218 57.97 181.015 55.747 181.642 53.511C183.018 48.599 184.617 33.936 176.534 34.015C172.986 34.05 171.149 37.654 167.27 36.819C164.498 36.222 162.417 33.859 161.436 31.298C160.04 27.656 160.57 23.474 163.939 21.438C168.229 18.845 170.586 21.832 175.192 22.465C183.635 23.627 182.448 10.076 181.078 0.5H100.727C102.191 10.432 103.245 23.573 95.1911 22.465Z"
							stroke-linecap="round"
							stroke="#FBFBFB"
							fill="#D9D9D9"
							animate={{
								opacity: [0, 1]
							}}
							transition={{
								duration: 0.5,
								delay: Math.random() * (5 - 0.25) + 0.25,
								repeat: Infinity,
								repeatType: "reverse",
								repeatDelay: 1
							}}
						/>
						<motion.path
							fill-rule="evenodd"
							clip-rule="evenodd"
							d="M261.964 19.086C259.639 25.311 254.447 22.261 249.701 20.788C238.245 17.231 237.966 38.375 249.79 36.792C253.962 32.812 259.49 32.903 261.051 38.722C262.401 43.756 262.257 49.089 261.956 54.244C261.875 55.639 262.174 58.84 262.418 62.617C271.874 62.586 281.334 62.024 290.7 60.702C291.947 56.646 286.257 53.799 287.657 50.534C289.188 46.962 295.661 46.803 298.976 46.911C302.876 47.038 307.706 47.714 310.77 50.229C311.803 54.783 305.665 56.156 307.305 60.697C316.004 64.226 325.343 63.349 334.336 61.609C336.006 61.286 337.678 60.934 339.355 60.584C340.166 58.219 341.014 55.868 341.671 53.501C343.038 48.581 344.619 33.937 336.53 34.015C332.982 34.05 331.145 37.654 327.266 36.819C324.494 36.222 322.413 33.859 321.432 31.298C320.036 27.656 320.566 23.474 323.935 21.438C328.225 18.845 330.582 21.832 335.188 22.465C344.175 23.701 342.387 8.982 340.548 0.5H260.84C261.869 6.596 263.46 15.08 261.964 19.086Z"
							stroke-linecap="round"
							stroke="#FBFBFB"
							fill="url(#gradient)"
							animate={{
								opacity: [0, 1]
							}}
							transition={{
								duration: 0.5,
								delay: Math.random() * (5 - 0.25) + 0.25,
								repeat: Infinity,
								repeatType: "reverse",
								repeatDelay: 1
							}}
						/>
						<motion.path
							fill-rule="evenodd"
							clip-rule="evenodd"
							d="M340.551 0.5C342.39 8.982 344.178 23.702 335.191 22.465C330.586 21.831 328.228 18.844 323.938 21.438C320.569 23.475 320.039 27.656 321.435 31.298C322.416 33.859 324.497 36.221 327.269 36.819C331.148 37.654 332.985 34.05 336.533 34.015C344.622 33.936 343.041 48.581 341.674 53.501C341.017 55.868 340.168 58.218 339.358 60.584C345.756 59.245 352.21 57.922 358.758 58.277C362.902 58.501 366.978 59.468 370.718 61.283C371.799 65.467 365.045 69.72 368.231 72.722C371.811 76.097 378.665 76.083 383.248 75.091C385.946 74.507 388.535 73.403 390.771 71.776C391.47 68.761 388.68 67.849 387.481 65.478C384.416 59.418 406.344 59.172 419.502 60.483V0.5H340.551Z"
							stroke-linecap="round"
							stroke="#FBFBFB"
							fill="url(#gradient)"
							animate={{
								opacity: [0, 1]
							}}
							transition={{
								duration: 0.5,
								delay: Math.random() * (5 - 0.25) + 0.25,
								repeat: Infinity,
								repeatType: "reverse",
								repeatDelay: 1
							}}
						/>
						<motion.path
							fill-rule="evenodd"
							clip-rule="evenodd"
							d="M309.562 128.121C316.41 135.141 294.797 135.648 291.114 134.459C286.01 132.812 287.336 130.167 289.627 126.591C292.803 121.634 289.295 120.242 284.762 119.141C276.171 117.053 267.672 118.732 259.16 120.516C257.485 125.471 256.055 130.477 256.864 135.591C258.718 147.291 266.466 140.207 273.868 140.026C281.995 139.828 279.529 153.368 275.702 156.7C268.86 159.554 263.422 149.067 258.259 157.879C254.752 163.864 255.849 170.511 257.96 176.758C258.607 178.673 259.371 180.627 260.088 182.612C262.422 182.626 264.756 182.613 267.09 182.561C272.358 182.442 277.621 182.146 282.864 181.634C285.483 181.379 288.097 181.069 290.703 180.701C291.482 178.165 289.611 175.451 288.27 173.524C282.987 165.928 302.258 167.017 305.596 167.891C310.476 169.169 312.634 171.204 308.818 175.709C303.281 182.245 313.295 181.298 317.917 181.656C325.39 182.235 332.884 182.591 340.379 182.627C340.206 175.128 343.737 167.668 341.861 160.152C340.632 155.228 337.519 151.98 332.801 155.448C330.395 157.218 327.347 157.466 324.841 155.754C319.612 152.183 318.722 140.194 326.96 140.051C331.101 139.979 335.164 142.758 339.718 142.715C341.532 140.034 342.479 136.924 342.458 133.688C342.432 129.461 340.795 124.842 339.144 120.16C333.735 119.028 328.313 117.982 322.769 117.993C317.082 118.006 301.564 119.923 309.562 128.121Z"
							stroke-linecap="round"
							stroke="#FBFBFB"
							fill="url(#gradient)"
							animate={{
								opacity: [0, 1]
							}}
							transition={{
								duration: 0.5,
								delay: Math.random() * (5 - 0.25) + 0.25,
								repeat: Infinity,
								repeatType: "reverse",
								repeatDelay: 1
							}}
						/>
						<motion.path
							fill-rule="evenodd"
							clip-rule="evenodd"
							d="M99.7176 142.717C95.1636 142.76 91.1005 139.981 86.9605 140.053C78.7215 140.196 79.6125 152.184 84.8405 155.756C87.3455 157.468 90.3936 157.219 92.7996 155.45C97.0976 152.29 100.053 154.655 101.769 159.068C104.464 165.995 101.832 173.293 99.5725 179.932C99.5445 180.014 99.5166 180.097 99.4896 180.178C102.256 180.76 105.019 181.361 107.788 181.848C111.909 182.574 116.099 183.07 120.29 182.872C124.557 182.67 132.808 182.61 130.135 176.596C128.79 173.57 125.459 171.257 129.332 168.856C132.722 166.756 137.324 166.639 141.215 167.075C144.539 167.447 150.737 167.897 150.975 172.018C150.152 174.029 148.977 175.892 147.225 177.225C146.353 181.669 149.048 180.886 152.626 181.207C161.853 182.036 171.113 182.576 180.378 182.619C180.208 175.124 183.736 167.667 181.86 160.154C180.631 155.231 177.518 151.982 172.8 155.45C170.394 157.22 167.346 157.468 164.84 155.756C159.611 152.185 158.721 140.196 166.959 140.053C171.1 139.981 175.163 142.76 179.717 142.717C181.542 140.034 182.503 136.92 182.491 133.675C182.475 129.455 180.841 124.834 179.186 120.148C174.461 119.156 169.727 118.228 164.892 118.015C159.276 117.768 140.963 119.308 149.56 128.121C156.029 134.753 136.614 135.811 133.061 134.996C127.067 133.621 126.576 131.35 129.624 126.592C132.784 121.66 129.329 120.245 124.806 119.136C116.768 117.166 108.448 118.711 100.498 120.336C100.112 120.415 99.7235 120.495 99.3355 120.575C100.938 125.115 102.476 129.584 102.49 133.676C102.504 136.919 101.543 140.034 99.7176 142.717Z"
							stroke-linecap="round"
							stroke="#FBFBFB"
							fill="#D9D9D9"
							animate={{
								opacity: [0, 1]
							}}
							transition={{
								duration: 0.5,
								delay: Math.random() * (5 - 0.25) + 0.25,
								repeat: Infinity,
								repeatType: "reverse",
								repeatDelay: 1
							}}
						/>
						<motion.path
							fill-rule="evenodd"
							clip-rule="evenodd"
							d="M387.307 120.697C385.666 116.156 391.805 114.782 390.772 110.229C387.709 107.713 382.878 107.038 378.978 106.911C374.802 106.774 371.993 108.394 368.251 109.319C366.293 113.192 369.957 114.503 370.678 118.306C372.263 124.551 353.5 122.792 349.866 122.24C346.276 121.695 342.713 120.91 339.142 120.162C340.793 124.845 342.43 129.464 342.456 133.69C342.477 136.926 341.53 140.036 339.716 142.717C335.162 142.76 331.099 139.981 326.958 140.053C318.72 140.196 319.61 152.184 324.839 155.756C327.345 157.468 330.393 157.219 332.799 155.45C337.517 151.982 340.629 155.23 341.859 160.154C343.735 167.67 340.204 175.13 340.377 182.629C343.501 182.644 346.626 182.609 349.749 182.502C355.011 182.322 360.267 181.958 365.499 181.367C368.034 181.081 371.637 182.083 370.678 178.306C370.218 175.875 368.504 174.004 367.229 171.979C368.38 164.334 382.852 166.889 387.701 168.522C391.833 173.166 381.833 180.833 389.773 180.939C393.083 181.318 396.367 181.887 399.667 182.34C403.586 182.878 412.454 183.01 419.501 182.55V125.666C407.798 126.949 393.134 124.113 387.307 120.697Z"
							stroke-linecap="round"
							stroke="#FBFBFB"
							fill="url(#gradient)"
							animate={{
								opacity: [0, 1]
							}}
							transition={{
								duration: 0.5,
								delay: Math.random() * (5 - 0.25) + 0.25,
								repeat: Infinity,
								repeatType: "reverse",
								repeatDelay: 1
							}}
						/>
					</svg>
				</InfoCard>
			</div>
		</Container>
	)
}

const Letter = () => (
	<>
		<Container className="my-[90px] flex-col items-center gap-4">
			<motion.h2
				className="text-center text-[28px] font-bold lg:w-[65%] lg:text-[64px]"
				initial={{ opacity: 0, y: 20 }}
				whileInView={{ opacity: 1, y: 0 }}
				transition={{ duration: 0.2 }}
			>
				Letter From The Team
			</motion.h2>
			<motion.p
				className="text-center text-[18px] font-light opacity-40 lg:w-[40%] lg:text-[24px]"
				initial={{ opacity: 0, y: 20 }}
				whileInView={{ opacity: 0.4, y: 0 }}
				transition={{ duration: 0.2, delay: 0.2 }}
			>
				We’re here to empower humans to benefit from the blockchain to
				the maximum extent and provide the ability to log off instead of
				being terminally online.
			</motion.p>

			<motion.div
				className="mt-[40px] grid lg:grid-cols-12"
				initial={{ opacity: 0, y: 20 }}
				whileInView={{ opacity: 1, y: 0 }}
				transition={{ duration: 0.2, delay: 0.4 }}
			>
				<div className="col-span-6 col-start-4 bg-[#D9D9D9]/10 px-8 lg:mx-8 lg:px-24">
					<div className="mb-16 flex flex-row justify-between gap-2">
						{Array.from({ length: 32 }).map((_, index) => (
							<div
								key={index}
								className="h-24 w-[2px] bg-gradient-to-b from-[#00EF35] to-[#93DF00]"
							/>
						))}
					</div>

					<div className="flex flex-col gap-6 opacity-65">
						<p>Dear anon,</p>
						<p>
							Every turn of the market cycle we find ourselves
							spending more time online. Yet, we go to sleep and
							miss the opportunity we had been trying to time for
							weeks.
						</p>
						<p>
							You were not born to watch numbers go up and down.
							You were not born to obsess over the smallest
							details of blockchain transactions. You were not
							born to live like a robot in a world of abundance.
						</p>
						<p>
							Plug is designed to give you your life back. To give
							you the reality we all dream of where your money
							works even when you don’t. To let you unlock the
							fully power of onchain financial primitives and the
							composability between the many options.
						</p>
						<p>
							You can choose to stay in the past and get worse
							execution with unexpected outcomes or you can use
							Plug and always be certain the outcomes will be
							generated when your conditions have been met.
						</p>
						<p>Love,</p>
						<Image
							src="/landing/signature.svg"
							alt="Signature"
							width={280}
							height={120}
						/>
					</div>

					<div className="mt-16 flex flex-row justify-between gap-2">
						{Array.from({ length: 32 }).map((_, index) => (
							<div
								key={index}
								className="h-8 w-[2px] bg-gradient-to-b from-[#00EF35] to-[#93DF00]"
							/>
						))}
					</div>
				</div>
			</motion.div>
		</Container>
	</>
)

const FrequentlyAskedQuestion: FC<{ text: string; description: string }> = ({
	text,
	description
}) => {
	const [collapsed, setCollapsed] = useState(true)

	return (
		<div className="flex flex-col gap-2 border-b-[1px] border-[#D9D9D9]/40 pb-4">
			<button
				onClick={() => setCollapsed(!collapsed)}
				className="z-[30] flex w-full items-center text-[24px] font-bold"
			>
				{text}
				<motion.span
					className="ml-auto transform rounded-full bg-[#FBFBFB] p-1 transition-transform"
					initial={{ rotate: 0 }}
					animate={{ rotate: collapsed ? 0 : 180 }}
					transition={{ duration: 0.2 }}
				>
					<ChevronDown size={24} className="opacity-40" />
				</motion.span>
			</button>
			<motion.p
				className="text-black/65 opacity-40 lg:mr-16"
				initial={{ height: 0, opacity: 0 }}
				animate={{
					height: collapsed ? 0 : "auto",
					opacity: collapsed ? 0 : 1
				}}
				transition={{ duration: 0.2 }}
			>
				{description}
			</motion.p>
		</div>
	)
}

const FrequentlyAskedQuestions = () => (
	<Container>
		<div className="my-[90px] grid lg:grid-cols-12">
			<h3 className="mb-8 text-[28px] font-bold lg:col-span-4 lg:col-start-2 lg:mb-0 lg:w-[60%] lg:text-[64px] 2xl:w-[50%]">
				Frequently Asked Questions
			</h3>

			<div className="flex flex-col gap-8 lg:col-span-5 lg:col-start-7">
				<FrequentlyAskedQuestion
					text="How does Plug work?"
					description="With a simple plug-and-play interface, you can establish conditions and outcomes across top Ethereum based protocols. Execute immediately, schedule for the future, or set up a recurring transaction. Plug’s bots take care of the hard work."
				/>
				<FrequentlyAskedQuestion
					text="Who is Plug for?"
					description="Plug was made with you in mind. Whether you’re managing on behalf of yourself, a venture fund, a market maker, or a DAO, Plug enables you to operate securely and in ways never before possible without having to write a single line of code. Even if you’re not an expert you can utilize the Plugs curated by the team to get started in seconds."
				/>
				<FrequentlyAskedQuestion
					text="What should I automate?"
					description="All across there are actions that you find boring or difficult and struggle to scale. Automation serves as a way to minimize the impact of each of those downsides and replace them with significant upside and potential. If there’s something you don’t like doing or can’t do, you should automate it."
				/>
				<FrequentlyAskedQuestion
					text="Is Plug really trustless?"
					description="Yes! Plug utilizes intents with embedded transaction conditions that are checked during the time of simulation and onchain execution. You don’t have to trust us or a centralized oracle of any kind. All data is sourced and verified onchain during simulation and execution."
				/>
			</div>
		</div>
	</Container>
)

const Vision = () => {
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

const Page = () => {
	return (
		<>
			<Head>
				<title>Home | Plug</title>
			</Head>
			<Navbar />
			<Hero />
			<Steps />
			<Examples />
			<Value />
			<CallToAction
				text="Get what you want from every transaction."
				description="Simultaneous settlement ensures transactions only execute when the conditions and expected outcomes you set in your intent can be met. No fees are paid and tokens move unless everything happens as expected."
				button="Get Started"
			/>
			<Templates />
			<Vision />
			<CallToAction
				text="Blockchains weren’t built for humans."
				description="Instead of being trapped inside waiting to click the buttons you can go outside and live the life you want without worrying about missing every opportunity. Your capital can finally manage itself."
				button="Get Early Access"
			/>
			<Letter />
			<FrequentlyAskedQuestions />
			<Footer />
		</>
	)
}

export default Page
