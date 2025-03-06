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
			{Array.from({ length: 24 }).map((_, bubbleIndex) => {
				const row = Math.floor(bubbleIndex / 6);
				const col = bubbleIndex % 6;
				return (
				<motion.div 
					key={bubbleIndex} 
					className="border-[12px] border-plug-white absolute rounded-full bg-[#EAEEE5] p-12 flex flex-row items-center gap-4"
					style={{
						left: `${(col * 33.33) - 50 + (Math.random() * 20)}%`,
						top: `${(row * 33.33) - 50 + (Math.random() * 20)}%`,
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
				);
			})}
		</div>),
	"infra-race": (<div className="grid grid-cols-11 grid-rows-10 w-full h-full">
		{Array.from({ length: 110 }).map((_, index) => {
			const row = Math.floor(index / 11);
			const col = index % 11;
			return (
				<motion.div 
					key={index} 
					className={cn(
						"w-full h-full",
						index % 2 == 0 ? "bg-plug-green/10" : ""
					)}
					animate={{
						y: [-2, 2, -2],
						x: [-1, 1, -1]
					}}
					transition={{
						duration: 2,
						repeat: Infinity,
						ease: "easeInOut",
						delay: (row * 0.1) + (col * 0.1)
					}}
				/>
			)
		})}
	</div>),
	"why-we-built-plug": (<div className="h-full flex flex-col items-center gap-4 mt-4 mb-4">
			{Array.from({ length: 12 }).map((_, lineIndex) => (
				<div key={lineIndex} className="w-full h-1 relative bg-plug-green/10 rounded-full">
					{Array.from({ length: 20 }).map((_, dotIndex) => {
						const opacity = 0.3 + (Math.sin(dotIndex * 0.5) * 0.7);
						return (
							<motion.div
								key={`${lineIndex}-${dotIndex}`}
								className={`bg-plug-green absolute h-3 w-3 rounded-full -mt-1`}
								style={{ 
									left: "-5%",
									opacity
								}}
								animate={{ 
									left: "105%",
									scale: [1, 1 + opacity * 0.7, 1]
								}}
								transition={{
									duration: 1 + Math.random(),
									repeat: Infinity,
									ease: "linear",
									delay: (dotIndex * (1.5 + Math.random())) + (lineIndex * (0.1 + Math.random() * 0.3))
								}}
							/>
						)
					})}
				</div>
			))}
		</div>)
}
