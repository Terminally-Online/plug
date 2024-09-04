import { motion } from "framer-motion"
import { Wallet } from "lucide-react"

import { InfoCard } from "@/components"

export const Cardiogram = () => {
	return (
		<InfoCard
			icon={<Wallet size={24} className="opacity-40" />}
			text="Automate or Die"
			description="The market is evolving. It’s time to upgrade to a modern tool stack. You’re leaving money on the table."
			className="col-span-2 h-[280px] sm:h-[320px] 2xl:h-[300px]"
		>
			<svg className="absolute inset-0 h-[50%] w-full" viewBox="0 0 200 100" preserveAspectRatio="none">
				<defs>
					<linearGradient id="cardiogramGradient" x1="0%" y1="0%" x2="100%" y2="0%">
						<stop offset="0%" stopColor="#00E100" />
						<stop offset="50%" stopColor="#A3F700" />
						<stop offset="100%" stopColor="#00E100" />
					</linearGradient>
					<filter id="glow">
						<feGaussianBlur stdDeviation="8" result="coloredBlur" />
						<feMerge>
							<feMergeNode in="coloredBlur" />
							<feMergeNode in="SourceGraphic" />
						</feMerge>
					</filter>
				</defs>
				<motion.path
					d="M0,50 H200 L210,50 L220,30 L230,70 L240,50 L250,50 H300 L310,50 L320,20 L330,80 L340,50 L350,50 H400 L410,50 L420,40 L430,60 L440,50 L450,50 H500 L510,50 L520,25 L530,75 L540,50 L550,50 H600 L610,50 L620,35 L630,65 L640,50 L650,50 H700 L710,50 L720,15 L730,85 L740,50 L750,50 H800 L810,50 L820,45 L830,55 L840,50 L850,50 H900 L910,50 L920,30 L930,70 L940,50 L950,50 H1000 L1010,50 L1020,20 L1030,80 L1040,50 L1050,50 H1100 L1110,50 L1120,40 L1130,60 L1140,50 L1150,50 H1200 L1210,50 L1220,25 L1230,75 L1240,50 L1250,50 H1300 L1310,50 L1320,35 L1330,65 L1340,50 L1350,50 H1400 L1410,50 L1420,15 L1430,85 L1440,50 L1450,50 H1500 L1510,50 L1520,45 L1530,55 L1540,50 L1550,50 H1750"
					fill="none"
					stroke="url(#cardiogramGradient)"
					strokeWidth="2"
					filter="url(#glow)"
					initial={{ pathLength: 0, opacity: 0 }}
					animate={{
						pathLength: 1,
						opacity: 1,
						transform: ["translateX(0px)", "translateX(-1550px)"],
						transition: {
							pathLength: { duration: 2, ease: "linear" },
							opacity: { duration: 0.5 },
							transform: {
								duration: 50,
								ease: "linear",
								repeat: Infinity,
								repeatType: "loop"
							}
						}
					}}
				/>
			</svg>

			<div className="absolute bottom-[50%] left-0 right-0 top-0 bg-gradient-to-b from-grayscale-0/0 to-grayscale-0" />
			<div className="absolute bottom-0 left-0 right-0 top-[50%] bg-grayscale-0" />
		</InfoCard>
	)
}
