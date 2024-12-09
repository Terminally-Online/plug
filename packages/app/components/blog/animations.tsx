import { motion } from "framer-motion"

export const postAnimations = {
	"pulse-of-crypto": (
		<div
			className="grid items-center justify-center gap-[2px]"
			style={{
				gridTemplateColumns: "repeat(20, 1fr)"
			}}
		>
			{Array.from({ length: 160 }).map((_, index) => {
				const row = Math.floor(index / 20)
				const col = index % 20

				const frequency = 0.5
				const amplitude = 5
				const speed = 0.1

				const waveRow = Math.sin(col * frequency) * amplitude + 4

				const distanceFromWave = Math.abs(row - waveRow)
				const isPartOfWave = distanceFromWave < 2

				return (
					<motion.div
						key={index}
						className="mx-auto h-8 w-full rounded-[4px]"
						initial={{ backgroundColor: "#385842" }}
						animate={{
							backgroundColor: isPartOfWave ? ["#385842", "#D2F38A", "#385842"] : ["#385842"]
						}}
						transition={{
							duration: 2,
							repeat: Infinity,
							repeatType: "reverse",
							delay: col * speed
						}}
					/>
				)
			})}
		</div>
	)
}
