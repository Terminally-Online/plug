import { FC } from "react"

import Image from "next/image"

import { motion } from "framer-motion"
import {
	CalendarClock,
	Check,
	Clock,
	FileStack,
	HandCoins,
	Rotate3d,
	RotateCw,
	Ruler,
	ShieldPlus,
	TestTubeDiagonal,
	X
} from "lucide-react"

import { Container, Fees, InfoCard, Mitigation } from "@/components"

export const Value: FC = () => {
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
			padding: "4px 8px"
		},
		whileInView: {
			background:
				active === true
					? [
							"linear-gradient(30deg, rgba(0,239,53,0.4), rgba(147,233,0,0.9))",
							"linear-gradient(30deg, rgba(0,239,53,1), rgba(147,233,0,1))",
							"linear-gradient(30deg, rgba(0,239,53,0.4), rgba(147,233,0,0.9))"
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
		<Container>
			<div className="grid w-full gap-8 md:grid-cols-2 lg:grid-rows-3 xl:grid-cols-3">
				<InfoCard
					text={
						<div className="flex flex-row gap-4">
							<Clock size={24} className="opacity-40" />
							<span>24/7 Execution</span>
						</div>
					}
					description="Plug keeps your strategies running at all times. Whether you're at dinner, asleep, or on vacation, your transactions will be executed."
					{...animation}
					className={`${animation.className} md:col-span-2 xl:col-span-1`}
				>
					<div
						className="mx-8 grid h-44 gap-[2px] 2xl:h-48"
						style={{ gridTemplateColumns: "repeat(14, 1fr)" }}
					>
						{Array.from({ length: 14 }).map((_, index) => (
							<motion.div
								key={index}
								className="mt-auto w-full origin-bottom rounded-lg"
								style={{
									background:
										index === 13
											? "linear-gradient(30deg, #00E100, #A3F700)"
											: "#D9D9D9"
								}}
								initial={{ height: 10 }}
								animate={{
									height: [
										10,
										25 * 2 ** (0.04 * (index / 2) * 8) +
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
				</InfoCard>

				<InfoCard
					text={
						<div className="flex flex-row gap-4">
							<ShieldPlus size={24} className="opacity-40" />
							<span>Loss Mitigation</span>
						</div>
					}
					description="Automatically protect yourself from losses in the volatile market and instantly react to every change of the ecosystem."
					{...animation}
				>
					<Mitigation />
				</InfoCard>

				<InfoCard
					text={
						<div className="flex flex-row gap-4">
							<CalendarClock size={24} className="opacity-40" />
							<span>Scheduled Transactions</span>
						</div>
					}
					description="Execute your transaction with the same scheduling methods as traditional calendar apps you're used to."
					{...animation}
				>
					<div className="ml-auto grid w-full grid-cols-7 grid-rows-4 text-xs">
						<div className="h-10 border-b-[1px] border-r-[1px]" />
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
							className="h-10 border-r-[1px]"
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
						<div className="flex flex-row gap-4">
							<FileStack size={24} className="opacity-40" />
							<span>Multichain Signatures</span>
						</div>
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
						<div className="flex flex-row gap-4">
							<RotateCw size={24} className="opacity-40" />
							<span>Recurring Outcomes</span>
						</div>
					}
					description="With Plug your signed intents can be reused to run again when the conditions are met without needing to sign a new transaction."
					{...animation}
				>
					<div
						className="ml-[-4px] mt-[-9px] grid w-[102%] grid-rows-3 gap-[2px]"
						style={{ gridTemplateColumns: "repeat(28, 1fr)" }}
					>
						{Array.from({ length: 28 * 7 }).map((_, index) => {
							const background =
								Math.random() < 0.5
									? "#D9D9D9"
									: "linear-gradient(30deg, #00E100, #A3F700)"
							return (
								<motion.div
									key={index}
									className="h-6 w-full rounded-[2px]"
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
						<div className="flex flex-row gap-4">
							<Ruler size={24} className="opacity-40" />
							<span>Atomic Constraints</span>
						</div>
					}
					description="Rules of execution are deeply embedded to enable automation that only results in what you want, when you want, how you want."
					{...animation}
				>
					<svg
						viewBox="0 0 479 126"
						fill="none"
						xmlns="http://www.w3.org/2000/svg"
						className="mt-8 lg:mx-auto lg:mt-8 2xl:mx-0 2xl:w-[100%]"
					>
						<motion.path
							d="M0 18.9805H112.217C112.217 18.9805 119.584 19.2795 123.438 22.3844C129.266 27.0795 134.228 40.6805 134.228 40.6805C134.228 40.6805 139.19 54.2814 145.018 58.9765C148.872 62.0814 156.24 62.3805 156.24 62.3805"
							strokeWidth="2"
							strokeDasharray="4 4"
							animate={{
								stroke: [
									"#D9D9D9",
									"#00E100",
									"#D9D9D9",
									"#FF0000",
									"#D9D9D9",
									"#00E100",
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
									"#00E100",
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
									"#00E100",
									"#D9D9D9",
									"#FF0000",
									"#D9D9D9",
									"#00E100",
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
									"#00E100",
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
									"#00E100",
									"#D9D9D9",
									"#00E100",
									"#D9D9D9",
									"#00E100",
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
									"#00E100",
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
									"#00E100",
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
									"#00E100",
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
									"#00E100",
									"#D9D9D9",
									"#FF0000",
									"#D9D9D9",
									"#00E100",
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
									"#00E100",
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
									"#00E100",
									"#D9D9D9",
									"#FF0000",
									"#D9D9D9",
									"#00E100",
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
									"#00E100",
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
									"#00E100",
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
							<HandCoins size={24} className="opacity-40" />
							<span>Creator Fees</span>
						</>
					}
					description="As a strategy creator, when others fork and use your strategy you automatically receive a portion of the fees generated."
					{...animation}
				>
					<Fees />
				</InfoCard>

				<InfoCard
					text={
						<div className="flex flex-row gap-4">
							<TestTubeDiagonal
								size={24}
								className="opacity-40"
							/>
							<span>Constant Simulation</span>
						</div>
					}
					description="To achieve near-instant responses your transaction is constantly simulated to check the expected outcomes can be achieved."
					{...animation}
				>
					<motion.div
						className="mx-auto ml-auto mt-20 flex w-[420px] items-center gap-6 rounded-lg bg-white px-6 py-2"
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
								animate={{ color: ["#00E100", "#FF5154"] }}
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
							<Rotate3d size={24} className="opacity-40" />
							<span>Transferrable Accounts</span>
						</>
					}
					description="Transfer your account to a different wallet or sell it on any major NFT marketplace. Your account is yours to do with as you please."
					{...animation}
				>
					<div className="flex w-full flex-row items-center">
						<motion.div
							className="ml-[-65px] mt-12 flex h-[130px] w-[130px] items-center justify-center rounded-full border-[2px] border-dashed"
							animate={{
								borderColor: ["#D9D9D9", "#00E100"]
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
										"linear-gradient(90deg, #00E100 25%, #A3F700 50%, transparent 50%)",
									backgroundRepeat: "repeat",
									backgroundSize: "4px 6px"
								}}
							/>
							<motion.div
								className="absolute top-[-5px] h-3 w-3 rounded-full"
								style={{
									background:
										"linear-gradient(30deg, #00E100, #A3F700)"
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
								borderColor: ["#D9D9D9", "#00E100"]
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
			</div>
		</Container>
	)
}
