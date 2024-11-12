import { useMemo } from "react"

import { motion } from "framer-motion"
import { Code } from "lucide-react"

import { InfoCard } from "@/components"

export const Underperforming = () => {
	const bars = useMemo(() => {
		return Array.from({ length: 40 }).map((_, index) => ({
			key: index,
			height: `${(100 * Math.pow(1.05, index)) / Math.pow(1.05, 39)}%`,
			redHeight: `${Math.min(50, Math.max(0, 30 - (index - 20) * 1.5))}%`,
			delay: 0.05 * index
		}))
	}, [])

	const animationProps = {
		duration: 2,
		repeat: Infinity,
		repeatType: "reverse" as const,
		repeatDelay: 1.5,
		ease: "easeInOut" as const
	}

	return (
		<InfoCard
			icon={<Code size={24} className="opacity-40" />}
			text="Unmatched returns."
			description="The best tools set you up for unparalleled successes and makes winning effortless."
			className="col-span-2 h-[280px] sm:h-[320px] 2xl:h-[300px]"
		>
			<div className="relative -mx-1 flex h-[40%] flex-row gap-1 pt-2">
				{bars.map(({ key, height, redHeight, delay }) => (
					<motion.div
						key={key}
						className="relative mt-auto w-full rounded-xl bg-plug-green/10"
						initial={{ height: height }}
						animate={{ height: [height, "0%"] }}
						transition={{ ...animationProps, delay }}
					>
						<motion.div
							className="absolute bottom-0 left-0 right-0 h-[12px] rounded-xl bg-plug-red"
							initial={{ bottom: redHeight }}
							animate={{ bottom: [redHeight, "0%"] }}
							transition={{ ...animationProps, delay }}
						/>
						<motion.div
							className="absolute bottom-0 left-0 right-0 h-[12px] rounded-xl bg-plug-yellow"
							initial={{ bottom: "calc(100% - 6px)" }}
							animate={{ bottom: ["calc(100% - 6px)", "0%"] }}
							transition={{ ...animationProps, delay }}
						/>
					</motion.div>
				))}
			</div>

			<div className="absolute bottom-[45%] left-0 right-0 top-[25%] bg-gradient-to-b from-plug-white/0 to-plug-white" />
			<div className="absolute bottom-0 left-0 right-0 top-[55%] bg-plug-white" />
		</InfoCard>
	)
}
