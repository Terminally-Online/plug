import { motion } from "framer-motion"
import { Wallet } from "lucide-react"

import { InfoCard } from "@/components"

export const Recurring = () => {
	return (
		<InfoCard
			icon={<Wallet size={24} className="opacity-40" />}
			text="Recurring runs made easy."
			description="No need to sign a new transaction every time you want it to run. Set your transactions repeat with one click."
			className="col-span-2 h-[280px] sm:h-[320px] 2xl:h-[300px]"
		>
			<div
				className="ml-[-4px] mt-[-9px] grid w-[102%] grid-rows-3 gap-[2px]"
				style={{ gridTemplateColumns: "repeat(28, 1fr)" }}
			>
				{Array.from({ length: 28 * 7 }).map((_, index) => {
					const background = Math.random() < 0.5 ? "#D9D9D9" : "linear-gradient(30deg, #00E100, #A3F700)"
					return (
						<motion.div
							key={index}
							className="h-6 w-full rounded-[2px]"
							style={{ background }}
							initial={{ opacity: 1 }}
							animate={{ opacity: [1, 0] }}
							transition={{
								duration: 0.5,
								delay: 0.05 * index,
								repeat: Infinity,
								repeatType: "reverse"
							}}
						/>
					)
				})}
			</div>

			<div className="absolute bottom-[50%] left-0 right-0 top-0 bg-gradient-to-b from-grayscale-0/0 to-grayscale-0" />
			<div className="absolute bottom-0 left-0 right-0 top-[50%] bg-grayscale-0" />
		</InfoCard>
	)
}
