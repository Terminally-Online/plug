import { motion } from "framer-motion"
import { Wallet } from "lucide-react"

import { InfoCard } from "@/components"

export const Recurring = () => {
	return (
		<InfoCard
			icon={<Wallet size={24} className="opacity-40" />}
			text="Recurring runs made easy."
			description="Consistent outcomes with a click of a button. Set it and forget it with transactions that auto-execute when you want."
			className="col-span-2 h-[280px] sm:h-[320px] 2xl:h-[300px]"
		>
			<div className="ml-[-4px] grid h-[50%] w-[102%] gap-[2px]">
				{Array.from({ length: 5 }).map((_, rowIndex) => (
					<div key={rowIndex} className="flex flex-row gap-[2px]">
						{Array.from({ length: 16 }).map((_, colIndex) => (
							<motion.div
								key={colIndex}
								className="h-full w-full rounded-sm"
								style={{
									background:
										Math.random() < 0.5 ? "#D9D9D9" : "linear-gradient(30deg, #00E100, #A3F700)"
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

			<div className="absolute bottom-[50%] left-0 right-0 top-0 bg-gradient-to-b from-grayscale-0/0 to-grayscale-0" />
			<div className="absolute bottom-0 left-0 right-0 top-[50%] bg-grayscale-0" />
		</InfoCard>
	)
}
