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
	"chat-interfaces": (<div className="w-full h-full flex items-center justify-center mt-8 relative">
		{Array.from({ length: 36 }).map((_, bubbleIndex) => {
			const row = Math.floor(bubbleIndex / 6);
			const col = bubbleIndex % 6;
			return (
				<motion.div
					key={bubbleIndex}
					className="border-[12px] border-plug-white absolute rounded-full bg-[#EAEEE5] p-12 flex flex-row items-center gap-4"
					style={{
						left: `${(col * 20) - 50 + (Math.random() * 15)}%`,
						top: `${(row * 20) - 50 + (Math.random() * 15)}%`,
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

	"one-app": (<motion.div
		className="mx-auto min-h-[320px] h-[50%] bg-plug-white rounded-2xl border-[10px] border-[#EBEEEC] flex items-center justify-center"
		animate={{
			marginTop: ["10%", "5%"],
			width: ["40%", "80%"]
		}}
		transition={{
			duration: 4,
			repeat: Infinity,
			repeatType: "reverse",
			repeatDelay: 2,
			ease: "easeInOut",
		}}
	>
		<div className="flex flex-col gap-4 w-[30%] text-center mx-auto">
			<svg xmlns="http://www.w3.org/2000/svg" xmlSpace="preserve" viewBox="0 0 814 1000" className="w-[40%] h-auto mx-auto">
				<path fill="#385842" d="M788.1 340.9c-5.8 4.5-108.2 62.2-108.2 190.5 0 148.4 130.3 200.9 134.2 202.2-.6 3.2-20.7 71.9-68.7 141.9-42.8 61.6-87.5 123.1-155.5 123.1s-85.5-39.5-164-39.5c-76.5 0-103.7 40.8-165.9 40.8s-105.6-57-155.5-127C46.7 790.7 0 663 0 541.8c0-194.4 126.4-297.5 250.8-297.5 66.1 0 121.2 43.4 162.7 43.4 39.5 0 101.1-46 176.3-46 28.5 0 130.9 2.6 198.3 99.2zm-234-181.5c31.1-36.9 53.1-88.1 53.1-139.3 0-7.1-.6-14.3-1.9-20.1-50.6 1.9-110.8 33.7-147.1 75.8-28.5 32.4-55.1 83.6-55.1 135.5 0 7.8 1.3 15.6 1.9 18.1 3.2.6 8.4 1.3 13.6 1.3 45.4 0 102.5-30.4 135.5-71.3z" />
			</svg>
			<div className="h-1 bg-plug-green/10 rounded-lg w-full relative">
				<motion.div
					className="absolute left-0 top-0 bottom-0 bg-plug-green rounded-lg"
					animate={{ width: ["0%", "100%"] }}
					transition={{
						duration: 10,
						repeat: Infinity,
						repeatDelay: 2,
						ease: "easeInOut"
					}}
				/>
			</div>
		</div>
	</motion.div>),
	"why-we-built-plug": (<div className="h-full flex flex-col items-center justify-between mt-4 mb-4">
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
