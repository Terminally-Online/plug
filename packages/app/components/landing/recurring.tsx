import { motion } from "framer-motion"
import { Wallet } from "lucide-react"

import { InfoCard } from "@/components/landing/cards/info"

export const Recurring = () => {
	return (
		<InfoCard
			icon={<Wallet size={24} className="opacity-40" />}
			text="Recurring runs."
			description="Set it and forget it with transactions that auto-execute on the frequency you have defined."
			className="relative z-[99999] col-span-2 h-[320px] sm:h-[320px] 2xl:h-[300px]"
		>
			<div className="ml-[-4px] grid h-[50%] w-[102%] gap-[2px]">
				{Array.from({ length: 5 }).map((_, rowIndex) => (
					<div key={rowIndex} className="flex flex-row gap-[2px]">
						{Array.from({ length: 16 }).map((_, colIndex) => (
							<motion.div
								key={colIndex}
								className="h-full w-full rounded-[2px]"
								style={{
									background: Math.random() < 0.5 ? "rgba(56,88,66,0.1)" : "#D2F38A"
								}}
								initial={{ opacity: 1 }}
								animate={{ opacity: [1, 0] }}
								transition={{
									duration: 0.5,
									delay: 0.05 * colIndex + 0.1 * rowIndex,
									repeat: Infinity,
									repeatType: "reverse"
								}}
							/>
						))}
					</div>
				))}
			</div>

			<div className="absolute bottom-[50%] left-0 right-0 top-0 bg-gradient-to-b from-plug-white/0 to-plug-white" />
			<div className="absolute bottom-0 left-0 right-0 top-[50%] bg-plug-white" />
		</InfoCard>
	)
}
