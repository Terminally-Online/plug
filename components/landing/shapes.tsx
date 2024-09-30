import { motion } from "framer-motion"

import { cn } from "@/lib"

export const HeroShapes = () => {
	const gradients = [
		"linear-gradient(45deg, #00E100, #A3F700)",
		"linear-gradient(135deg, #A3F700, #00E100)",
		"linear-gradient(45deg, #00E100, #A3F700)",
	]

	const positions = [
		{ top: "100%", left: "-15%" },
		{ top: "75%", right: "0%" },
		{ top: "-30%", right: "15%" }
	]

	return (
		<>
			<div className="absolute inset-0 bottom-0 left-0 right-0 top-0 z-0 w-screen overflow-hidden">
				<div className="absolute inset-0 blur-[80px] filter">
					{gradients.map((gradient, index) => (
						<motion.div
							key={index}
							className={cn("absolute rounded-full", `z-${index}`)}
							style={{
								background: gradient,
								width: `${30 + Math.random() * 10}%`,
								height: `${30 + Math.random() * 10}%`,
								...positions[index]
							}}
							animate={{
								rotate: [0, Math.random() < 0.5 ? 360 : -360, 0],
								translateX: [
									`${Math.random() * 10 - 5}%`,
									`${Math.random() * 10 - 5}%`,
									`${Math.random() * 10 - 5}%`
								],
								translateY: [
									`${Math.random() * 10 - 5}%`,
									`${Math.random() * 10 - 5}%`,
									`${Math.random() * 10 - 5}%`
								]
							}}
							transition={{
								duration: 30 + Math.random() * 3,
								ease: "easeInOut",
								repeat: Infinity
							}}
						/>
					))}
				</div>
			</div>

			<div className="absolute inset-0 top-0 left-0 right-0 bottom-0 z-1"
				style={{
					backgroundImage: `url(/cheese.svg)`,
					backgroundSize: "cover",
					backgroundPosition: "center",
					backgroundRepeat: "no-repeat"
				}}
			/>
		</>
	)
}
