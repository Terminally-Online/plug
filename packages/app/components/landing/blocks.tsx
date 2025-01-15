import Image from "next/image"
import { useEffect, useState } from "react"

import { motion } from "framer-motion"

import { LandingContainer } from "@/components/landing/layout/container"
import { GTM_EVENTS, routes, useAnalytics, useMediaQuery } from "@/lib"

const blockchains = [
	"/blockchain/arbitrum",
	"/blockchain/avalanche",
	"/blockchain/base",
	"/blockchain/bera",
	"/blockchain/blast",
	"/blockchain/ethereum",
	"/blockchain/optimism",
	"/blockchain/polygon",
	"/blockchain/scroll",
	"/blockchain/zksync"
]

const protocols = [
	"/protocols/yearn",
	"/protocols/hop",
	"/protocols/gearbox",
	"/protocols/aerodrome",
	"/protocols/zora",
	"/protocols/sushiswap",
	"/protocols/alchemix",
	"/protocols/eigen-layer",
	"/protocols/ethena",
	"/protocols/balancer",
	"/protocols/chainlink",
	"/protocols/rocket-pool",
	"/protocols/compound",
	"/protocols/maker",
	"/protocols/fraxlend",
	"/protocols/curve",
	"/protocols/lido",
	"/protocols/synthetix",
	"/protocols/wasabi",
	"/protocols/ens",
	"/protocols/convex",
	"/protocols/paraswap",
	"/protocols/uniswap",
	"/protocols/aave"
]

function shuffleArray<T>(array: Array<T>): Array<T> {
	const shuffled = [...array]
	for (let i = shuffled.length - 1; i > 0; i--) {
		const j = Math.floor(Math.random() * (i + 1))
		;[shuffled[i], shuffled[j]] = [shuffled[j], shuffled[i]]
	}
	return shuffled
}

export const Blocks = () => {
	const { md } = useMediaQuery()
	const handleCallToAction = useAnalytics(GTM_EVENTS.CTA_CLICKED, `${routes.documentation}/introduction/integrations`)

	const [hoveredItems, setHoveredItems] = useState<Set<string>>(new Set())
	const [allHovered, setAllHovered] = useState(false)
	const [resetAll, setResetAll] = useState(false)
	const [hasFallen, setHasFallen] = useState(false)
	const [shuffledItems, setShuffledItems] = useState<string[]>([])

	const rows = md ? 3 : 4
	const columns = md ? 12 : 5

	useEffect(() => {
		const allItems = [...protocols, ...blockchains]
		const shuffled = shuffleArray(allItems)
		const totalNeeded = rows * 12
		while (shuffled.length < totalNeeded) {
			shuffled.push(...shuffleArray(allItems))
		}
		setShuffledItems(shuffled.slice(0, totalNeeded))
	}, [rows])

	const handleItemHover = (item: string) => {
		if (!hasFallen) {
			setHoveredItems(prev => {
				const newSet = new Set(prev).add(item)
				if (newSet.size === protocols.length + blockchains.length) {
					setAllHovered(true)
				}
				return newSet
			})
		}
	}

	const handleItemClick = () => {
		if (allHovered && !hasFallen) {
			setResetAll(true)
			setAllHovered(false)
			setHoveredItems(new Set([...protocols, ...blockchains]))
			setHasFallen(true)
		}
	}

	return (
		<div className="relative z-[9999] overflow-hidden">
			<LandingContainer className="relative mb-[40px] flex flex-col gap-4">
				<div className="flex flex-row items-center gap-12">
					<motion.h1
						className="max-w-[420px] text-[52px] font-black leading-tight md:max-w-[520px] lg:min-w-[480px] lg:text-[64px]"
						initial={{ transform: "translateY(-20px)", opacity: 0 }}
						whileInView={{
							transform: ["translateY(-20px)", "translateY(0px)"],
							opacity: [0, 1]
						}}
						transition={{ duration: 0.3 }}
					>
						All of Ethereum in one place.
					</motion.h1>

					<div className="hidden w-full items-center gap-4 md:visible xl:flex xl:flex-row">
						<div className="h-[2px] w-full bg-plug-green/10" />
						<button
							className="whitespace-nowrap font-bold opacity-40 transition-opacity duration-200 ease-in-out hover:opacity-100"
							onClick={() => handleCallToAction()}
						>
							Explore Integrations
						</button>
						<div className="h-[2px] w-24 bg-plug-green/10" />
					</div>
				</div>
				<motion.p
					className="max-w-[520px] text-xl font-bold text-black/40 md:max-w-[480px] lg:text-[18px]"
					initial={{ transform: "translateY(20px)", opacity: 0 }}
					whileInView={{
						transform: ["translateY(20px)", "translateY(0px)"],
						opacity: [0, 1]
					}}
					transition={{ duration: 0.3 }}
				>
					Use onchain protocols the way they were meant to be experienced. Access advanced functionality from
					the best protocols in one unified interface.
				</motion.p>

				<div className="mt-12 flex w-full flex-col gap-2">
					{[...Array(rows)].map((_, rowIndex) => {
						const rowItems = shuffledItems.slice(rowIndex * columns, (rowIndex + 1) * columns)

						return (
							<div key={rowIndex} className="flex w-full flex-wrap gap-2">
								{rowItems.map((item, index) => (
									<motion.div
										key={`${item}-${index}`}
										className="relative flex aspect-square flex-1 cursor-pointer items-center justify-center rounded-xl border-[1px] border-plug-green/10 bg-plug-white transition-all duration-200 ease-in-out"
										onMouseEnter={allHovered ? undefined : () => handleItemHover(item)}
										onClick={handleItemClick}
										animate={
											allHovered && !resetAll && !hasFallen
												? {
														transform: `translateY(${Math.random() * 500 + 200}px) rotate(${
															Math.random() < 0.5
																? Math.random() * 180
																: Math.random() * -180
														}deg)`,
														transition: { duration: 1, ease: "easeInOut" }
													}
												: {
														transform: "translateY(0) rotate(0deg)",
														transition: { duration: 0.5, ease: "easeInOut" }
													}
										}
									>
										<div className="relative h-[40%] w-[40%] transition-all duration-200 ease-in-out group-hover:scale-110">
											{((hoveredItems.has(item) && !hasFallen) || resetAll) && (
												<motion.div
													className="absolute bottom-0 left-0 right-0 top-0 blur-2xl filter"
													style={{
														backgroundImage: `url(${item}.png)`,
														backgroundSize: "contain",
														backgroundPosition: "center",
														backgroundRepeat: "no-repeat"
													}}
													initial={{ opacity: 0 }}
													whileInView={{ opacity: 1 }}
													transition={{ duration: 0.3 }}
												/>
											)}

											<Image
												src={`${item}.png`}
												alt={item.split("/").pop() || ""}
												layout="fill"
												objectFit="contain"
												className={`relative transition-all duration-200 ease-in-out ${
													(hoveredItems.has(item) && !hasFallen) || resetAll
														? ""
														: "grayscale"
												}`}
											/>
										</div>
									</motion.div>
								))}
							</div>
						)
					})}
				</div>
			</LandingContainer>
		</div>
	)
}
