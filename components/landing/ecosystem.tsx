import { FC, useMemo } from "react"

import Image from "next/image"

import { motion } from "framer-motion"

const getRandomDelay = (min = 0, max = 2) => {
	return Math.random() * (max - min) + min
}

const Line: FC<{ size: number; index: number; imagePath: string }> = ({
	size,
	index,
	imagePath
}) => {
	const active = true

	const transition = useMemo(
		() => ({
			duration: 4,
			repeat: Infinity,
			repeatDelay: getRandomDelay(0.5, 3.5 * 3),
			delay: getRandomDelay(0, 3.5 * 3)
		}),
		[]
	)

	return (
		<motion.div
			key={index}
			className="absolute left-[50%] top-0 origin-bottom border-l-[2px] border-dashed"
			style={{
				height: `${size / 2 - 1}px`,
				transform: `rotate(${index * 15}deg) translateX(-1px)`
			}}
			animate={{
				borderColor: active
					? ["#D9D9D9", "#00E100", "#00E100", "#00E100", "#D9D9D9"]
					: "#D9D9D9"
			}}
			transition={transition}
		>
			<motion.div
				className="absolute left-[50%] top-[-32px] z-[10] flex h-16 w-16 translate-x-[-50%] transform items-center justify-center rounded-lg border-[2px] border-dashed bg-white"
				animate={{
					borderColor: active
						? [
								"#D9D9D9",
								"#D9D9D9",
								"#00E100",
								"#D9D9D9",
								"#D9D9D9"
							]
						: "#D9D9D9"
				}}
				transition={transition}
			>
				<motion.div
					animate={{
						rotate: [index * -15, index * -15 - 360]
					}}
					transition={{
						duration: 200,
						repeat: Infinity,
						repeatDelay: 0
					}}
					className="relative p-4"
				>
					<motion.div
						initial={{ opacity: 0 }}
						whileInView={{ opacity: 1 }}
						transition={{ duration: 1 }}
					>
						<Image
							src={imagePath}
							alt="Logo for the ecosystem"
							width={64}
							height={32}
							className="absolute left-0 top-0 blur-sm filter"
						/>

						<Image
							src={imagePath}
							alt="Logo for the ecosystem"
							width={64}
							height={32}
							className="absolute left-0 top-0"
						/>
					</motion.div>
				</motion.div>
			</motion.div>

			<motion.div
				className="absolute bottom-0 h-3 w-3 rounded-full border-[1px]"
				style={{ backgroundColor: "#00E100" }}
				initial={{ x: -7, y: -(size / 2), scale: 0.6 }}
				animate={{ y: [0, -(size / 2), 0], scale: [0.6, 1, 0.6] }}
				transition={{ ...transition }}
			/>
		</motion.div>
	)
}

// Create lines from the center to the edge of the circle.
const Lines: FC<{ size: number }> = ({ size = 900 }) => {
	const protocols = [
		"yearn.png",
		"hop.png",
		"gearbox.png",
		"aerodrome.png",
		"zora.png",
		"sushiswap.png",
		"alchemix.png",
		"eigen-layer.png",
		"ethena.png",
		"balancer.png",
		"chainlink.png",
		"rocket-pool.png",
		"compound.png",
		"maker.png",
		"fraxlend.png",
		"curve.png",
		"lido.png",
		"synthetix.png",
		"wasabi.png",
		"ens.png",
		"convex.png",
		"paraswap.png",
		"uniswap.png",
		"aave.png"
	]

	return (
		<>
			{Array.from({ length: 24 }).map((_, index) => (
				<Line
					key={index}
					size={size}
					index={index}
					imagePath={`/protocols/${protocols[index]}`}
				/>
			))}
		</>
	)
}

// Visualize the ecosystem.
export const Ecosystem: FC = () => {
	const size: number = 700

	const rotationTransition = {
		duration: 200,
		repeat: Infinity
	}

	return (
		<motion.div
			style={{
				minHeight: `${size + 200}px`,
				minWidth: `${size + 200}px`
			}}
			className={`absolute bottom-[-42%] left-[calc(50%_-_450px)] flex items-center justify-center lg:bottom-[5%] lg:left-[65%] lg:right-[-30%]`}
		>
			<div className="absolute bottom-0 left-1/2 right-[35%] top-0 z-[20] hidden bg-gradient-to-r from-white/0 to-white lg:visible" />

			<motion.div
				className="relative rounded-full border-[2px] border-dashed border-[#D9D9D9]"
				style={{
					height: `${size}px`,
					width: `${size}px`
				}}
				animate={{ rotate: [0, 360] }}
				transition={rotationTransition}
			>
				<Lines size={size} />

				<motion.div
					className="absolute z-[10] h-16 w-16 origin-center translate-x-[-50%] transform"
					style={{
						top: "calc(50% - 32px)",
						left: "calc(50% - 32px)"
					}}
					animate={{ rotate: [0, -360] }}
					transition={rotationTransition}
				>
					<motion.div
						className="flex h-full w-full items-center justify-center rounded-lg border-[2px] border-dashed bg-white"
						style={{ borderColor: "#D9D9D9" }}
						animate={{
							borderColor: ["#D9D9D9", "#00E100", "#D9D9D9"]
						}}
						transition={{
							duration: 0.15,
							repeat: Infinity,
							repeatDelay: 0.1,
							delay: 1.5
						}}
					>
						<Image
							src="/black-icon.svg"
							alt="Logo"
							width={64}
							height={32}
							className="p-4"
						/>
					</motion.div>
				</motion.div>
			</motion.div>
		</motion.div>
	)
}

export default Ecosystem
