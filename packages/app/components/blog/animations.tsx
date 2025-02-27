import { cn } from "@/lib"
import { motion } from "framer-motion"

export const postAnimations = {
	"hello-world": (
		<div
			className="grid items-center justify-center gap-[2px]"
			style={{
				gridTemplateColumns: "repeat(19, 1fr)"
			}}
		>
			{Array.from({ length: 160 }).map((_, index) => {
				const row = Math.floor(index / 19)
				const col = index % 19

				const pattern = [
					[0, 2, 2, 0, 0, 2, 2, 0, 0, 2, 0, 0, 2, 2, 0, 0, 2, 2, 0],
					[0, 0, 2, 2, 0, 0, 2, 2, 0, 2, 0, 2, 2, 0, 0, 2, 2, 0, 0],
					[2, 0, 0, 2, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 2, 2, 0, 0, 2],
					[2, 2, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0, 2, 2],
					[0, 0, 2, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 2, 0, 0],
					[2, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 2],
					[2, 2, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 2, 2],
					[0, 2, 2, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 2, 2, 0],
				]

				const pixelType = pattern[row]?.[col] || 0
				const isPartOfSun = pixelType === 1
				const isRay = pixelType === 2

				return (
					<motion.div
						key={index}
						className={cn(
							"mx-auto h-8 w-full rounded-[4px]",
							isPartOfSun ? "bg-plug-yellow" : isRay ? "bg-plug-yellow" : "bg-transparent"
						)}
						initial={{
							opacity: 0,
							...(!isRay && { y: 40 })
						}}
						animate={{
							opacity: 1,
							y: 0,
							...(isRay && { opacity: [0.4, 1, 0.4] })
						}}
						transition={{
							duration: isRay ? 1.5 + Math.random() * 0.5 : 0.5,
							repeat: isRay ? Infinity : 0,
							repeatType: "reverse",
							ease: "easeInOut",
							delay: isPartOfSun
								? 0
								: isRay
									? 0.5 + (row * 0.1) + (Math.random() * 0.3)
									: 0
						}}
					/>
				)
			})}
		</div>
	),
	"chat-interfaces": (<div className="w-full h-1/2 flex items-center justify-center mt-8 relative">
			{Array.from({ length: 36 }).map((_, bubbleIndex) => (
				<motion.div 
					key={bubbleIndex} 
					className="border-[12px] border-plug-white absolute rounded-full bg-[#EAEEE5] p-12 flex flex-row items-center gap-4"
					style={{
						left: `${-50 + Math.random() * 150}%`,
						top: `${-50 + Math.random() * 150}%`,
					}}
				>
					{Array.from({ length: 3 }).map((_, dotIndex) => (
						<motion.div
							key={`${bubbleIndex}-${dotIndex}`}
							className="w-8 h-8 rounded-full bg-plug-green/10"
							initial={{ y: -10 }}
							animate={{ y: 10 }}
							transition={{
								repeat: Infinity,
								repeatType: "reverse",
								ease: "easeInOut",
								delay: 0.1 * dotIndex,
								repeatDelay: 0.1,
							}}
						/>
					))}
				</motion.div>
	
			))}
		</div>)
}
