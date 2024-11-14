import { motion } from "framer-motion"
import { ReplaceAll } from "lucide-react"

import { InfoCard } from "../cards"

const SYMBOLS = ["ETH", "BTC", "SOL", "AVAX", "MATIC", "UNI", "AAVE", "SUSHI", "YFI", "COMP", "MKR", "CRV"]

export const ActionSwap = () => {
	return (
		<InfoCard
			icon={<ReplaceAll size={24} className="opacity-40" />}
			text="Swap."
			description="Simple built in swaps in just a few clicks."
			className="relative z-[99999] col-span-2 h-[320px] sm:h-[320px] xl:col-span-1 2xl:h-[300px]"
		>
			<div className="absolute inset-0 flex flex-col items-center justify-center">
				{SYMBOLS.map((symbol, index) => (
					<motion.div
						key={index}
						className="origin-bottom-center absolute -top-8 left-[12%] mx-auto flex h-3/4 w-3/4 translate-x-1/4 items-center justify-center rounded-full border-[2px] border-dashed border-plug-green bg-plug-yellow font-bold text-plug-green"
						initial={{
							y: "0rem"
						}}
						animate={{
							y: ["12rem", "0rem", "12rem"]
						}}
						style={{ zIndex: -1 * index }}
						transition={{
							duration: 1,
							repeat: Infinity,
							ease: "easeInOut",
							delay: index * 0.6,
							repeatDelay: SYMBOLS.length * 0.6
						}}
					>
						<p className="text-[48px] font-black">${symbol}</p>
					</motion.div>
				))}
			</div>

			<div className="absolute bottom-[50%] left-0 right-0 top-0 bg-gradient-to-b from-plug-white/0 to-plug-white" />
			<div className="absolute bottom-0 left-0 right-0 top-[50%] bg-plug-white" />
		</InfoCard>
	)
}
