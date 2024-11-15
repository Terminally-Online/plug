import Image from "next/image"
import { useRef } from "react"

import { motion, useScroll, useTransform } from "framer-motion"
import { LayoutDashboard, Sparkles, Wallet } from "lucide-react"

import { InfoCard } from "./cards"
import { LandingContainer } from "./layout"

const ITEMS = 24
const PROTOCOLS = [
	"yearn",
	"hop",
	"gearbox",
	"aerodrome",
	"zora",
	"sushiswap",
	"alchemix",
	"eigen-layer",
	"ethena",
	"balancer",
	"chainlink",
	"rocket-pool",
	"compound",
	"maker",
	"fraxlend",
	"curve",
	"lido",
	"synthetix",
	"wasabi",
	"ens",
	"convex",
	"paraswap",
	"uniswap",
	"aave"
]

export const Platform = () => {
	const containerRef = useRef<HTMLDivElement>(null)
	const { scrollYProgress } = useScroll({
		target: containerRef,
		offset: ["start end", "end start"]
	})
	const pathLength = useTransform(scrollYProgress, [0, 0.7], [1, 0])

	return (
		<div className="relative overflow-visible" ref={containerRef}>
			<div className="absolute inset-0 -right-[5%] top-[-10%] hidden overflow-visible xl:flex">
				<svg viewBox="0 0 1827 976" fill="none" className="absolute inset-0 overflow-visible">
					<g clip-path="url(#clip0_4624_28608)">
						<motion.path
							d="M1737.75 371C1482.5 233 1506.38 483.689 1349.5 476C1171.88 467.295 1181 296.5 931 296.5C681 296.5 618 811 350.5 720.5C83 630 387.75 196.5 -67.5 134.5"
							stroke="url(#paint0_linear_4624_28608)"
							strokeWidth="60"
						/>
						<motion.path
							d="M1737.75 371C1482.5 233 1506.38 483.689 1349.5 476C1171.88 467.295 1181 296.5 931 296.5C681 296.5 618 811 350.5 720.5C83 630 387.75 196.5 -67.5 134.5"
							stroke="#FEFFF7"
							strokeWidth="60"
							stroke-dasharray="4 4"
							animate={{ strokeDashoffset: [0, 60] }}
							transition={{
								duration: 0.5,
								repeat: Infinity,
								ease: "linear"
							}}
						/>
					</g>
					<defs>
						<linearGradient
							id="paint0_linear_4624_28608"
							x1="1362.25"
							y1="232.5"
							x2="467.75"
							y2="504"
							gradientUnits="userSpaceOnUse"
						>
							<stop stop-color="#D2F38A" />
							<stop offset="1" stop-color="#385842" />
						</linearGradient>
						<clipPath id="clip0_4624_28608">
							<rect width="1827" height="976" fill="white" />
						</clipPath>
					</defs>
				</svg>

				<svg viewBox="0 0 1827 976" fill="none" className="absolute inset-0 hidden overflow-visible xl:flex">
					<g clip-path="url(#clip0_4624_28608)">
						<motion.path
							style={{ pathLength }}
							d="M1737.75 371C1482.5 233 1506.38 483.689 1349.5 476C1171.88 467.295 1181 296.5 931 296.5C681 296.5 618 811 350.5 720.5C83 630 387.75 196.5 -67.5 134.5"
							stroke="url(#paint0_linear_4624_28608)"
							strokeWidth="60"
						/>
					</g>
					<defs>
						<linearGradient
							id="paint0_linear_4624_28608"
							x1="1362.25"
							y1="232.5"
							x2="467.75"
							y2="504"
							gradientUnits="userSpaceOnUse"
						>
							<stop stop-color="#D2F38A" />
							<stop offset="1" stop-color="#385842" />
						</linearGradient>
						<clipPath id="clip0_4624_28608">
							<rect width="1827" height="976" fill="white" />
						</clipPath>
					</defs>
				</svg>

				<svg
					viewBox="0 0 1827 976"
					fill="none"
					className="absolute inset-0 z-[99999] hidden overflow-visible xl:flex"
				>
					<mask
						id="mask0_4624_28614"
						style={{ maskType: "alpha" }}
						maskUnits="userSpaceOnUse"
						x="487"
						y="86"
						width="1251"
						height="693"
					>
						<rect x="487" y="86" width="1251" height="693" fill="#D9D9D9" />
					</mask>
					<g mask="url(#mask0_4624_28614)">
						<motion.path
							style={{ pathLength }}
							d="M1737.75 371C1482.5 233 1506.38 483.689 1349.5 476C1171.88 467.295 1181 296.5 931 296.5C681 296.5 618 811 350.5 720.5C83 630 387.75 196.5 -67.5 134.5"
							stroke="url(#paint0_linear_4624_28614)"
							strokeWidth="60"
							strokeLinecap="round"
						/>
					</g>
					<defs>
						<linearGradient
							id="paint0_linear_4624_28614"
							x1="1362.25"
							y1="232.5"
							x2="467.75"
							y2="504"
							gradientUnits="userSpaceOnUse"
						>
							<stop stop-color="#D2F38A" />
							<stop offset="1" stop-color="#385842" />
						</linearGradient>
					</defs>
				</svg>
			</div>

			<LandingContainer className="relative z-[9] mb-[80px] grid grid-rows-2 gap-2 xl:grid-cols-3">
				<InfoCard
					icon={<Sparkles size={24} className="opacity-40" />}
					text="Protocol Abstraction."
					description="All the functionality of major protocols available without having to bounce between apps."
					className="relative z-[99999] h-[320px] sm:h-[320px] 2xl:h-[300px]"
				>
					<style jsx>{`
						.clip-path-asteroid-trail {
							clip-path: polygon(40% 25%, 100% 0%, 100% 100%, 20% 75%);
						}
						.gradient-mask {
							mask-image: linear-gradient(to left, black, black, transparent, transparent, transparent);
							-webkit-mask-image: linear-gradient(
								to left,
								black,
								black,
								transparent,
								transparent,
								transparent
							);
						}
					`}</style>

					<div className="absolute inset-0 bottom-1/2">
						{Array.from({ length: PROTOCOLS.length }).map((_, index) => (
							<motion.div
								key={index}
								className="absolute inset-0"
								initial={{ y: `${Math.random() * 100}%`, x: "-20%" }}
								animate={{ x: ["-20%", "120%"] }}
								transition={{
									duration: 1,
									repeat: Infinity,
									delay: Math.random() * 2 + index * 0.4,
									ease: "linear",
									repeatDelay: 4
								}}
							>
								<motion.div className="absolute -ml-[5rem]" initial={{ width: "6rem" }}>
									<div className="clip-path-asteroid-trail gradient-mask fade-out-trail h-8 w-full blur filter">
										<Image
											src={`/protocols/${PROTOCOLS[index]}.png`}
											alt={PROTOCOLS[index]}
											width={240}
											height={240}
											className="h-8 w-full object-cover"
										/>
									</div>
								</motion.div>

								<Image
									src={`/protocols/${PROTOCOLS[index]}.png`}
									alt="Aave"
									width={240}
									height={240}
									className="absolute h-8 w-8 rounded-full"
								/>
							</motion.div>
						))}
					</div>
					<div className="absolute bottom-[50%] left-0 right-0 top-0 bg-gradient-to-b from-plug-white/0 to-plug-white" />
					<div className="absolute bottom-0 left-0 right-0 top-[50%] bg-plug-white" />
				</InfoCard>

				<InfoCard
					icon={<LayoutDashboard size={24} className="opacity-40" />}
					text="Modular Components."
					description="However you prefer it, the layout can be resized, reordered, and refocused to what matters."
					className="relative z-[99999] row-span-2 h-full xl:col-span-2"
				>
					<div className="absolute inset-0 flex flex-row items-center justify-center">
						<motion.div
							className="h-full border-r-[2px] border-plug-green/10"
							initial={{ width: "20%" }}
							animate={{ width: "30%" }}
							transition={{
								duration: 0.5,
								repeat: Infinity,
								repeatType: "reverse",
								ease: "easeInOut",
								delay: 3,
								repeatDelay: 6
							}}
						>
							<p className="truncate p-2 pl-4 font-bold">Discover</p>
							<div className="h-[2px] w-full bg-plug-green/10" />
							<div className="flex grid h-full grid-cols-2 gap-1 px-2 pt-2">
								{Array.from({ length: ITEMS }).map((_, index) => (
									<div
										className="gradient-animated h-[52px] w-full rounded-lg bg-grayscale-100"
										key={index}
									/>
								))}
							</div>
						</motion.div>
						<motion.div
							className="h-full border-r-[2px] border-plug-green/10"
							initial={{ width: "20%" }}
							animate={{ width: "30%" }}
							transition={{
								duration: 0.5,
								repeat: Infinity,
								repeatType: "reverse",
								ease: "easeInOut",
								delay: 2,
								repeatDelay: 10
							}}
						>
							<p className="truncate p-2 pl-4 font-bold">Tokens</p>
							<div className="h-[2px] w-full bg-plug-green/10" />
							<div className="flex h-full flex-col gap-1 px-2 pt-2">
								{Array.from({ length: ITEMS / 2 }).map((_, index) => (
									<div
										className="gradient-animated h-[48px] w-full rounded-lg bg-grayscale-100"
										key={index}
									/>
								))}
							</div>
						</motion.div>
						<motion.div
							className="h-full border-r-[2px] border-plug-green/10"
							initial={{ width: "30%" }}
							animate={{ width: "20%" }}
							transition={{
								duration: 0.5,
								repeat: Infinity,
								repeatType: "reverse",
								ease: "easeInOut",
								delay: 4,
								repeatDelay: 10
							}}
						>
							<p className="truncate p-2 pl-4 font-bold">Collectibles</p>
							<div className="h-[2px] w-full bg-plug-green/10" />
							<div className="flex h-full flex-col gap-1 px-2 pt-2">
								{Array.from({ length: ITEMS / 2 }).map((_, index) => (
									<div
										className="gradient-animated h-[48px] w-full rounded-lg bg-grayscale-100"
										key={index}
									/>
								))}
							</div>
						</motion.div>
						<motion.div
							className="h-full border-r-[2px] border-plug-green/10"
							initial={{ width: "30%" }}
							animate={{ width: "20%" }}
							transition={{
								duration: 0.5,
								repeat: Infinity,
								repeatType: "reverse",
								ease: "easeInOut",
								delay: 6,
								repeatDelay: 6
							}}
						>
							<p className="truncate p-2 pl-4 font-bold">Plugs</p>
							<div className="h-[2px] w-full bg-plug-green/10" />
							<div className="flex grid h-full grid-cols-2 gap-1 px-2 pt-2">
								{Array.from({ length: ITEMS }).map((_, index) => (
									<div
										className="gradient-animated h-[52px] w-full rounded-lg bg-grayscale-100"
										key={index}
									/>
								))}
							</div>
						</motion.div>

						<motion.div
							className="h-full"
							initial={{ width: "20%" }}
							animate={{ width: "40%" }}
							transition={{
								duration: 0.5,
								repeat: Infinity,
								repeatType: "reverse",
								ease: "easeInOut",
								delay: 6,
								repeatDelay: 6
							}}
						>
							<p className="truncate p-2 pl-4 font-bold">Activity</p>
							<div className="h-[2px] w-full bg-plug-green/10" />
							<div className="flex h-full flex-col gap-1 px-2 pt-2">
								{Array.from({ length: ITEMS / 2 }).map((_, index) => (
									<div
										className="gradient-animated h-[48px] w-full rounded-lg bg-grayscale-100"
										key={index}
									/>
								))}
							</div>
						</motion.div>
					</div>

					<div className="absolute bottom-[30%] left-0 right-0 top-0 bg-gradient-to-b from-plug-white/0 to-plug-white" />
					<div className="absolute bottom-0 left-0 right-0 top-[70%] bg-plug-white" />
				</InfoCard>
				<InfoCard
					icon={<Wallet size={24} className="opacity-40" />}
					text="Focused Context."
					description="With a global dashboard you can zoom all the way in or out for the context you want."
					className="relative z-[99999] h-[320px] sm:h-[320px] 2xl:h-[300px]"
				>
					<div className="absolute inset-0">
						<motion.div
							className="absolute bottom-1/2 left-[20%] right-[20%] top-8 rounded-lg border-[2px] border-plug-green/10"
							initial={{ y: "100%" }}
							animate={{ y: ["100%", "0%", "0%", "0%", "100%"] }}
							transition={{
								duration: 2,
								repeat: Infinity,
								ease: "easeInOut",
								repeatDelay: 4
							}}
						>
							<motion.div className="mx-6 flex flex-row items-center gap-1">
								<div className="h-6 w-6 rounded-full bg-plug-yellow" />
								<p className="p-2 pl-4 font-bold">Transfer $USDC</p>
							</motion.div>

							<div className="mb-2 h-[2px] w-full bg-plug-green/10" />
							<div className="gradient-animated mx-6 mb-2 h-16 rounded-lg bg-plug-green/10" />
							<div className="gradient-animated mx-6 mb-2 h-16 rounded-lg bg-plug-green/10" />
						</motion.div>

						<motion.div
							className="absolute bottom-1/2 left-[20%] right-[20%] top-8 rounded-lg border-[2px] border-plug-green/10"
							initial={{ y: "100%" }}
							animate={{ y: ["100%", "0%", "0%", "0%", "100%"] }}
							transition={{
								duration: 2,
								repeat: Infinity,
								ease: "easeInOut",
								delay: 2,
								repeatDelay: 4
							}}
						>
							<motion.div className="mx-6 flex flex-row items-center gap-1">
								<div className="h-6 w-6 rounded-full bg-plug-yellow" />
								<p className="p-2 pl-4 font-bold">Simulation Results</p>
							</motion.div>

							<div className="mb-2 h-[2px] w-full bg-plug-green/10" />
							<div className="flex flex-row items-end gap-2 px-6">
								<div className="gradient-animated mb-2 h-4 w-full rounded-lg bg-plug-green/10" />
								<div className="gradient-animated mb-2 h-8 w-full rounded-lg bg-plug-green/10" />
								<div className="gradient-animated mb-2 h-16 w-full rounded-lg bg-plug-green/10" />
								<div className="gradient-animated mb-2 h-12 w-full rounded-lg bg-plug-green/10" />
							</div>
						</motion.div>

						<motion.div
							className="absolute bottom-1/2 left-[20%] right-[20%] top-8 rounded-lg border-[2px] border-plug-green/10"
							initial={{ y: "100%" }}
							animate={{ y: ["100%", "0%", "0%", "0%", "100%"] }}
							transition={{
								duration: 2,
								repeat: Infinity,
								ease: "easeInOut",
								delay: 4,
								repeatDelay: 4
							}}
						>
							<motion.div className="mx-6 flex flex-row items-center gap-1">
								<div className="h-6 w-6 rounded-full bg-plug-yellow" />
								<p className="p-2 pl-4 font-bold">Schedule Run</p>
							</motion.div>

							<div className="mb-2 h-[2px] w-full bg-plug-green/10" />
							<div className="mx-6 mb-1 flex flex-row justify-between text-xs font-bold">
								<p>01/24</p>
								<p>02/24</p>
							</div>
							<div className="mx-6 grid h-full grid-cols-7 grid-rows-4 gap-1">
								{Array.from({ length: 28 }).map((_, index) => (
									<div key={index} className="h-full w-full rounded-sm bg-plug-green/10" />
								))}
							</div>
						</motion.div>
					</div>

					<div className="absolute bottom-[50%] left-0 right-0 top-0 bg-gradient-to-b from-plug-white/0 to-plug-white" />
					<div className="absolute bottom-0 left-0 right-0 top-[50%] bg-plug-white" />
				</InfoCard>
			</LandingContainer>
		</div>
	)
}
