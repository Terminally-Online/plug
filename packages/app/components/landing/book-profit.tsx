import { motion } from "framer-motion"
import { Radar } from "lucide-react"

import { InfoCard } from "@/components"

export const BookProfit = () => {
	return (
		<InfoCard
			icon={<Radar size={24} className="opacity-40" />}
			text="Constantly pinging the market."
			description="With Plug monitoring every onchain transaction, you can immediately respond no matter the state of the market."
			className="col-span-2 h-[540px] xl:col-span-4 xl:row-span-2 xl:h-full"
		>
			<div className="absolute inset-0 z-[-1] flex items-center justify-center overflow-hidden">
				{Array.from({ length: 20 }).map((_, index) => (
					<motion.div
						key={index}
						className="absolute origin-center rounded-full border-[1px] border-plug-yellow"
						style={{
							width: 20 + Math.sin(index * 0.5) * 20,
							height: 20 + Math.cos(index * 0.5) * 20
						}}
						animate={{
							width: [0, 20 + index * 200],
							height: [0, 20 + index * 200]
						}}
						transition={{
							duration: 1,
							repeat: Infinity,
							repeatType: "reverse",
							ease: "easeInOut",
							delay: index * 0.1
						}}
					/>
				))}
				{Array.from({ length: 12 }).map((_, index) => (
					<motion.div
						key={index}
						className="absolute origin-center rounded-full border-[1px] border-plug-yellow"
						animate={{
							height: [0, 2000],
							rotate: [index * 15, index * 15 + 90]
						}}
						transition={{
							duration: 1,
							repeat: Infinity,
							repeatType: "reverse",
							ease: "easeInOut"
						}}
					/>
				))}
			</div>
			<div className="absolute bottom-[30%] left-0 right-0 top-[50%] bg-gradient-to-b from-plug-white/0 to-plug-white" />
			<div className="absolute bottom-0 left-0 right-0 top-[70%] bg-plug-white" />
		</InfoCard>
	)
}
