import { motion } from "framer-motion"

export const HeroShapes = () => {
	const colors = ["#00E100", "#A3F700", "#00E100", "#A3F700", "#00E100", "#A3F700", "#00E100", "#A3F700"]

	const initialPositions = [
		{ x: "-10%", y: "-10%" },
		{ x: "110%", y: "-10%" },
		{ x: "-10%", y: "110%" },
		{ x: "110%", y: "110%" },
		{ x: "50%", y: "-10%" },
		{ x: "50%", y: "110%" },
		{ x: "-10%", y: "50%" },
		{ x: "110%", y: "50%" }
	]

	return (
		<div className="absolute inset-0 bottom-0 left-0 right-0 top-0 z-0 w-screen overflow-hidden bg-plug-green">
			<div className="absolute inset-0 blur-[120px] filter">
				{colors.map((color, index) => (
					<motion.div
						key={index}
						className="z-= absolute rounded-full"
						style={{
							background: color,
							width: "130%",
							height: "130%",
							x: initialPositions[index].x,
							y: initialPositions[index].y,
							top: "-65%",
							left: "-65%"
						}}
						animate={{
							x: [
								initialPositions[index].x,
								...["0%", "100%", "50%", initialPositions[index].x].filter(
									pos => pos !== initialPositions[index].x
								)
							],
							y: [
								initialPositions[index].y,
								...["0%", "100%", "50%", initialPositions[index].y].filter(
									pos => pos !== initialPositions[index].y
								)
							],
							scale: [1, 1.1, 0.9, 1]
						}}
						transition={{
							duration: 50,
							ease: "easeInOut",
							repeat: Infinity,
							delay: index * 6
						}}
					/>
				))}
			</div>
		</div>
	)
}
