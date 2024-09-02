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
		{ top: "5%", left: "10%" },
		{ top: "15%", right: "5%" },
		{ bottom: "8%", left: "12%" },
		{ bottom: "20%", right: "15%" },
		{ top: "30%", left: "45%" },
		{ bottom: "25%", left: "60%" },
		{ top: "55%", left: "8%" },
		{ top: "40%", right: "20%" }
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
							width: `${70 + Math.random() * 20}%`,
							height: `${70 + Math.random() * 20}%`,
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
							duration: 30 + Math.random() * 2,
							ease: "easeInOut",
							repeat: Infinity
						}}
					/>
				))}
			</div>
		</div>
	)
}
