import { motion } from "framer-motion"
import { Telescope } from "lucide-react"

import { InfoCard } from "../cards"

const PLUGS = [
	"Auto Compound",
	"Exit Loans",
	"Redeem Aero",
	"Bid on Noun",
	"DCA USDC:ETH",
	"Beta Volatility",
	"Rolling Dump PEPE",
	"Balance Health Factor",
	"Auto Compound",
	"Exit Loans",
	"Redeem Aero",
	"Bid on Noun",
	"DCA USDC:ETH",
	"Beta Volatility",
	"Rolling Dump PEPE",
	"Balance Health Factor"
]

const COLORS = [
	"#F3EF8A",
	"#8AF3E6",
	"#EB8AF3",
	"#9F8AF3",
	"#F3908A",
	"#F3B08A",
	"#D2F38A",
	"#F3EF8A",
	"#8AF3E6",
	"#EB8AF3",
	"#9F8AF3",
	"#F3908A",
	"#F3B08A",
	"#D2F38A"
]

export const ActionDiscover = () => {
	return (
		<InfoCard
			icon={<Telescope size={24} className="opacity-40" />}
			text="Discover Opportunities."
			description="Stay on top of the latest crypto opportunities and trends by exploring curated and commmunity plugs."
			className="relative z-[99999] col-span-2 h-[320px] sm:h-[320px] 2xl:h-[300px]"
		>
			<div className="absolute inset-0 -mx-8 grid grid-cols-4 gap-4">
				{PLUGS.map((plug, index) => {
					const columnIndex = index % 4

					return (
						<motion.div
							key={index}
							animate={{ y: [0, -40] }}
							transition={{
								repeat: Infinity,
								repeatType: "reverse",
								ease: "easeInOut",
								duration: 2,
								delay: columnIndex * 0.5
							}}
							className="relative flex h-20 h-full w-full items-center overflow-hidden rounded-lg rounded-sm border-[1px] border-plug-green/10 p-2 text-plug-white"
						>
							<div style={{ backgroundColor: COLORS[index] }} className="absolute inset-0 -z-[1]" />
							<p className="mt-auto font-bold">{plug}</p>
						</motion.div>
					)
				})}
			</div>

			<div className="absolute bottom-[50%] left-0 right-0 top-0 bg-gradient-to-b from-plug-white/0 to-plug-white" />
			<div className="absolute bottom-0 left-0 right-0 top-[50%] bg-plug-white" />
		</InfoCard>
	)
}
