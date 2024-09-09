import { motion } from "framer-motion"
import { CalendarClock } from "lucide-react"

import { InfoCard } from "@/components"

export const Scheduled = () => {
	const getDayAnimation = (delay: number, active: boolean = false) => ({
		style: {
			color: active === true ? "#FFFFFF" : "rgba(0,0,0,0.40)",
			borderColor: "#D9D9D9",
			padding: "4px 8px"
		},
		whileInView: {
			background:
				active === true
					? [
							"linear-gradient(30deg, rgba(0,239,53,0.65), rgba(147,233,0,1))",
							"linear-gradient(30deg, rgba(0,239,53,1), rgba(147,233,0,1))",
							"linear-gradient(30deg, rgba(0,239,53,0.65), rgba(147,233,0,1))"
						]
					: ["rgba(217,217,217,0)", "rgba(217,217,217,0.4)", "rgba(217,217,217,0)"]
		},
		transition: {
			duration: 0.25,
			delay: 1 + delay,
			repeat: Infinity,
			repeatDelay: 7.5
		}
	})

	return (
		<InfoCard
			icon={<CalendarClock size={24} className="opacity-40" />}
			text="Scheduled transactions."
			description="Define timeframes for your transactions. You don't have to be online to be onchain."
			className="col-span-2 h-[320px] sm:h-[320px] 2xl:h-[300px]"
		>
			<div className="ml-auto grid w-full grid-cols-7 grid-rows-4 text-xs">
				<div className="h-10 border-b-[1px] border-r-[1px]" />
				<div className="border-b-[1px] border-r-[1px]" />
				<div className="border-b-[1px] border-r-[1px]" />
				<motion.div className="border-b-[1px] border-r-[1px]" {...getDayAnimation(0)}>
					1
				</motion.div>
				<motion.div className="border-b-[1px] border-r-[1px]" {...getDayAnimation(0.25)}>
					2
				</motion.div>
				<motion.div className="border-b-[1px] border-r-[1px]" {...getDayAnimation(0.5)}>
					3
				</motion.div>
				<motion.div className="border-b-[1px]" {...getDayAnimation(0.75)}>
					4
				</motion.div>
				<motion.div className="border-b-[1px] border-r-[1px]" {...getDayAnimation(1)}>
					5
				</motion.div>
				<motion.div className="border-b-[1px] border-r-[1px]" {...getDayAnimation(1.25)}>
					6
				</motion.div>
				<motion.div className="border-b-[1px] border-r-[1px]" {...getDayAnimation(1.5, true)}>
					7
				</motion.div>
				<motion.div className="border-b-[1px] border-r-[1px]" {...getDayAnimation(1.75)}>
					8
				</motion.div>
				<motion.div className="border-b-[1px] border-r-[1px]" {...getDayAnimation(2)}>
					9
				</motion.div>
				<motion.div className="border-b-[1px] border-r-[1px]" {...getDayAnimation(2.25, true)}>
					10
				</motion.div>
				<motion.div className="border-b-[1px] border-[#D9D9D9]" {...getDayAnimation(2.5)}>
					11
				</motion.div>
				<motion.div className="border-b-[1px] border-r-[1px]" {...getDayAnimation(2.75)}>
					12
				</motion.div>
				<motion.div className="border-b-[1px] border-r-[1px]" {...getDayAnimation(3)}>
					13
				</motion.div>
				<motion.div className="border-b-[1px] border-r-[1px]" {...getDayAnimation(3.25, true)}>
					14
				</motion.div>
				<motion.div className="border-b-[1px] border-r-[1px]" {...getDayAnimation(3.5)}>
					15
				</motion.div>
				<motion.div className="border-b-[1px] border-r-[1px]" {...getDayAnimation(3.75)}>
					16
				</motion.div>
				<motion.div className="border-b-[1px] border-r-[1px]" {...getDayAnimation(4)}>
					17
				</motion.div>
				<motion.div className="border-b-[1px]" {...getDayAnimation(4.25)}>
					18
				</motion.div>
				<motion.div className="border-b-[1px] border-r-[1px]" {...getDayAnimation(4.5)}>
					19
				</motion.div>
				<motion.div className="border-b-[1px] border-r-[1px]" {...getDayAnimation(4.75)}>
					20
				</motion.div>
				<motion.div className="border-b-[1px] border-r-[1px] border-[#D9D9D9]" {...getDayAnimation(5, true)}>
					21
				</motion.div>
				<motion.div className="border-b-[1px] border-r-[1px]" {...getDayAnimation(5.25)}>
					22
				</motion.div>
				<motion.div className="border-b-[1px] border-r-[1px]" {...getDayAnimation(5.5)}>
					23
				</motion.div>
				<motion.div className="border-b-[1px] border-r-[1px]" {...getDayAnimation(5.75, true)}>
					24
				</motion.div>
				<motion.div className="border-b-[1px]" {...getDayAnimation(6)}>
					25
				</motion.div>
				<motion.div className="h-10 border-r-[1px]" {...getDayAnimation(6.25)}>
					26
				</motion.div>
				<motion.div className="border-r-[1px]" {...getDayAnimation(6.5)}>
					27
				</motion.div>
				<motion.div className="border-r-[1px]" {...getDayAnimation(6.75, true)}>
					28
				</motion.div>
				<motion.div className="border-r-[1px]" {...getDayAnimation(7)}>
					29
				</motion.div>
				<motion.div className="border-r-[1px]" {...getDayAnimation(7.25)}>
					30
				</motion.div>
				<div className="border-r-[1px]" />
			</div>

			<div className="absolute bottom-[50%] left-0 right-0 top-0 bg-gradient-to-b from-grayscale-0/0 to-grayscale-0" />
			<div className="absolute bottom-0 left-0 right-0 top-[50%] bg-grayscale-0" />
		</InfoCard>
	)
}
