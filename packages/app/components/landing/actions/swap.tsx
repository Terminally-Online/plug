import { motion } from "framer-motion"
import { ReplaceAll } from "lucide-react"

import { InfoCard } from "@/components/landing/cards/info"

const SYMBOLS = ["ETH", "BTC", "SOL", "AVAX", "MATIC", "UNI", "AAVE", "SUSHI", "YFI", "COMP", "MKR", "CRV"]

export const ActionSwap = () => {
	return (
		<InfoCard
			icon={<ReplaceAll size={24} className="opacity-40" />}
			text="Swap."
			description="No matter the tokens, we will find the best route."
			className="relative z-[99999] col-span-2 h-[320px] sm:h-[320px] xl:col-span-1 2xl:h-[300px]"
		>
			<div className="absolute inset-0 flex items-center justify-center">
				{SYMBOLS.map((symbol, index) => (
					<motion.div
						key={index}
						className="absolute -top-8 flex h-64 w-64 items-center justify-center rounded-full border-[1px] border-dashed border-plug-green bg-plug-yellow font-bold text-plug-green"
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
