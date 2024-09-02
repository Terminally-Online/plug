import { useEffect, useRef, useState } from "react"

import Image from "next/image"
import Link from "next/link"

import { motion } from "framer-motion"

import { Animate, Blob, LandingContainer } from "@/components"
import { cn, routes } from "@/lib"

const protocols = [
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

const ProtocolLine = ({
	isGreen,
	onHover,
	strumDelay,
	stopStrumming
}: {
	isGreen: boolean
	onHover: () => void
	strumDelay: number
	stopStrumming: () => void
}) => {
	const nextId = useRef(0)
	const availableProtocolsRef = useRef(protocols)

	const [activeProtocols, setActiveProtocols] = useState<Array<{ id: number; name: string }>>([])
	const [isTabVisible, setIsTabVisible] = useState(true)
	const [isStrumming, setIsStrumming] = useState(false)

	useEffect(() => {
		const handleVisibilityChange = () => {
			setIsTabVisible(!document.hidden)
		}

		document.addEventListener("visibilitychange", handleVisibilityChange)

		return () => {
			document.removeEventListener("visibilitychange", handleVisibilityChange)
		}
	}, [])

	useEffect(() => {
		let timeoutId: NodeJS.Timeout

		const addProtocol = () => {
			if (isTabVisible && Math.random() < 0.3 && availableProtocolsRef.current.length > 0) {
				const randomIndex = Math.floor(Math.random() * availableProtocolsRef.current.length)
				const newProtocol = availableProtocolsRef.current[randomIndex]
				const timestamp = Date.now() + nextId.current
				setActiveProtocols(prev => [...prev, { id: timestamp, name: newProtocol }])
				nextId.current += 1
				availableProtocolsRef.current = availableProtocolsRef.current.filter(p => p !== newProtocol)
			}
			timeoutId = setTimeout(addProtocol, Math.random() * 500 + 200)
		}

		addProtocol()

		return () => {
			clearTimeout(timeoutId)
			setActiveProtocols([])
			availableProtocolsRef.current = protocols
		}
	}, [isTabVisible])

	useEffect(() => {
		let strumTimeoutId: NodeJS.Timeout
		if (isGreen) {
			setIsStrumming(true)
			strumTimeoutId = setTimeout(() => setIsStrumming(false), 150 + strumDelay)
		}
		return () => {
			clearTimeout(strumTimeoutId)
		}
	}, [isGreen, strumDelay])

	const removeProtocol = (id: number) => {
		setActiveProtocols(prev => {
			const removed = prev.find(p => p.id === id)
			if (removed) {
				availableProtocolsRef.current = [...availableProtocolsRef.current, removed.name]
			}
			return prev.filter(protocol => protocol.id !== id)
		})
	}

	const handleMouseEnter = () => {
		onHover()
		stopStrumming()
	}

	return (
		<div
			className={cn("group relative ml-[-25vw] flex w-[150vw] overflow-hidden py-8")}
			onMouseEnter={handleMouseEnter}
		>
			<motion.div
				className={cn(
					"absolute top-[50%] h-[2px] w-full translate-y-[-50%] bg-gradient-to-r transition-all duration-200 ease-in-out",
					isGreen
						? "from-plug-green to-plug-yellow blur filter"
						: "from-grayscale-100 to-grayscale-100/0 group-hover:from-plug-green group-hover:to-plug-yellow group-hover:blur group-hover:filter"
				)}
				animate={
					isStrumming
						? {
								scaleY: [1, 2, 1],
								translateY: ["-50%", "-60%", "-50%"],
								rotate: [0, 2, -2, 0]
							}
						: {}
				}
				transition={{ duration: 0.15, times: [0, 0.5, 1] }}
			/>
			<motion.div
				className={cn(
					"absolute top-[50%] h-[2px] w-full translate-y-[-50%] bg-gradient-to-r transition-all duration-200 ease-in-out",
					isGreen
						? "from-plug-green to-plug-yellow"
						: "from-grayscale-100 to-grayscale-100/0 group-hover:from-plug-green group-hover:to-plug-yellow"
				)}
				animate={
					isStrumming
						? {
								scaleY: [1, 2, 1],
								translateY: ["-50%", "-60%", "-50%"],
								rotate: [0, 2, -2, 0]
							}
						: {}
				}
				transition={{ duration: 0.15, times: [0, 0.5, 1] }}
			/>

			{activeProtocols.map(({ id, name }) => (
				<motion.div
					key={id}
					initial={{ transform: "translateX(-2rem)" }}
					animate={{ transform: "translateX(160vw)" }}
					transition={{
						duration: 2.5,
						ease: "linear"
					}}
					onAnimationComplete={() => removeProtocol(id)}
					className="absolute top-[25%] z-[2]"
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

					<motion.div
						className="absolute right-4"
						initial={{ width: "12rem" }}
						animate={{ width: "6rem" }}
						transition={{ duration: 2 }}
					>
						<div className="clip-path-asteroid-trail gradient-mask fade-out-trail h-8 w-full blur filter">
							<Image
								src={`/protocols/${name}.png`}
								alt={name}
								width={240}
								height={240}
								className="h-8 w-full object-cover"
							/>
						</div>
					</motion.div>

					<motion.div className="relative h-8 w-8" initial={{ opacity: 0 }} whileInView={{ opacity: 1 }}>
						<Image
							src={`/protocols/${name}.png`}
							alt={name}
							width={240}
							height={240}
							className="h-8 w-8 min-w-8 rounded-full"
						/>
					</motion.div>
				</motion.div>
			))}
		</div>
	)
}

export const Light = () => {
	const hoverTimeoutRef = useRef<NodeJS.Timeout | null>(null)
	const animationTimeoutRef = useRef<NodeJS.Timeout | null>(null)

	const [_, setHoveredLines] = useState(new Set<number>())
	const [greenLineIndex, setGreenLineIndex] = useState(-1)
	const [direction, setDirection] = useState<"up" | "down">("down")
	const [animationCount, setAnimationCount] = useState(0)
	const [lastHoveredLine, setLastHoveredLine] = useState(-1)

	const handleLineHover = (index: number) => {
		setHoveredLines(prev => {
			const newSet = new Set(prev)
			newSet.add(index)
			setLastHoveredLine(index)
			if (newSet.size === 5) {
				if (hoverTimeoutRef.current) clearTimeout(hoverTimeoutRef.current)
				hoverTimeoutRef.current = setTimeout(() => {
					setGreenLineIndex(index)
					setDirection(index === 4 ? "up" : "down")
					setAnimationCount(0)
				}, 500)
			}
			return newSet
		})
	}

	const stopStrumming = () => {
		if (greenLineIndex !== -1) {
			setGreenLineIndex(-1)
			setHoveredLines(new Set())
			if (animationTimeoutRef.current) clearTimeout(animationTimeoutRef.current)
		}
	}

	useEffect(() => {
		if (greenLineIndex >= 0) {
			animationTimeoutRef.current = setTimeout(() => {
				if (direction === "down") {
					if (greenLineIndex < 4) {
						setGreenLineIndex(prev => prev + 1)
					} else {
						setDirection("up")
						setGreenLineIndex(3)
					}
				} else {
					if (greenLineIndex > 0) {
						setGreenLineIndex(prev => prev - 1)
					} else {
						if (animationCount < 2) {
							setDirection("down")
							setGreenLineIndex(1)
							setAnimationCount(prev => prev + 1)
						} else {
							setGreenLineIndex(-1)
							setHoveredLines(new Set())
						}
					}
				}
			}, 100)
		}

		return () => {
			if (animationTimeoutRef.current) clearTimeout(animationTimeoutRef.current)
		}
	}, [greenLineIndex, direction, animationCount])

	useEffect(() => {
		return () => {
			if (hoverTimeoutRef.current) clearTimeout(hoverTimeoutRef.current)
			if (animationTimeoutRef.current) clearTimeout(animationTimeoutRef.current)
		}
	}, [])

	return (
		<div className="relative z-0 my-[80px]">
			<Blob left={"-100"} top={"-700"} width={"1000"} height={"500"} />

			<LandingContainer className="relative mb-[40px] flex flex-col gap-4">
				<div className="flex flex-row items-center gap-12">
					<motion.h1
						className="min-w-[640px] text-[64px] font-bold leading-tight"
						initial={{ transform: "translateY(-20px)", opacity: 0 }}
						whileInView={{
							transform: ["translateY(-20px)", "translateY(0px)"],
							opacity: [0, 1]
						}}
						transition={{ duration: 0.3 }}
					>
						Onchain Activity at the Speed of Light.
					</motion.h1>
					<div className="h-[2px] w-full bg-grayscale-100" />
				</div>
				<motion.p
					className="max-w-[640px] text-[18px] font-bold text-black/40"
					initial={{ transform: "translateY(20px)", opacity: 0 }}
					whileInView={{
						transform: ["translateY(20px)", "translateY(0px)"],
						opacity: [0, 1]
					}}
					transition={{
						duration: 0.3,
						delay: 0.15
					}}
				>
					Trigger and build automatically executing transactions faster than you can comprehend. Combine
					onchain and offchain data feeds to drive the most informed transactions possible.
				</motion.p>
			</LandingContainer>

			<motion.div
				className="flex flex-col"
				style={{
					transform: "rotate(-15deg)"
				}}
				initial={{ opacity: 0 }}
				whileInView={{ opacity: 1 }}
			>
				<Animate.List>
					{Array.from({ length: 5 }).map((_, index) => (
						<Animate.ListItem key={index}>
							<ProtocolLine
								isGreen={greenLineIndex === index}
								onHover={() => handleLineHover(index)}
								strumDelay={Math.abs(index - lastHoveredLine) * 25}
								stopStrumming={stopStrumming}
							/>
						</Animate.ListItem>
					))}
				</Animate.List>
			</motion.div>
		</div>
	)
}
