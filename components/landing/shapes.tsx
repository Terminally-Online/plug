import { motion } from "framer-motion"

import { cn } from "@/lib"

export const HeroShapes = () => {
	const gradients = [
		"linear-gradient(45deg, #00E100, #A3F700)",
		"linear-gradient(135deg, #A3F700, #00E100)",
		"linear-gradient(225deg, #00E100, #A3F700)",
		"linear-gradient(315deg, #A3F700, #00E100)",
		"linear-gradient(90deg, #00E100, #A3F700)",
		"linear-gradient(270deg, #A3F700, #00E100)",
		"linear-gradient(0deg, #00E100, #A3F700)",
		"linear-gradient(180deg, #A3F700, #00E100)"
	]

	const positions = [
		{ top: "0%", left: "0%" },
		{ top: "0%", right: "0%" },
		{ bottom: "0%", left: "0%" },
		{ bottom: "0%", right: "0%" },
		{ top: "0%", left: "50%", transform: "translateX(-50%)" },
		{ bottom: "0%", left: "50%", transform: "translateX(-50%)" },
		{ top: "50%", left: "0%", transform: "translateY(-50%)" },
		{ top: "50%", right: "0%", transform: "translateY(-50%)" }
	]

	return (
		<div className="absolute inset-0 bottom-0 left-0 right-0 top-0 z-0 w-screen overflow-hidden bg-plug-green">
			<div className="absolute inset-0 blur-[40px] filter">
				{gradients.map((gradient, index) => (
					<motion.div
						key={index}
						className={cn("absolute rounded-full", `z-${index}`)}
						style={{
							background: gradient,
							width: "80%",
							height: "80%",
							...positions[index]
						}}
						animate={{
							rotate: [0, 360]
						}}
						transition={{
							duration: 30,
							ease: "linear",
							repeat: Infinity,
							delay: index * 6
						}}
					/>
				))}
			</div>
		</div>
	)
}
