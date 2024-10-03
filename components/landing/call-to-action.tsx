import { FC } from "react"

import { motion } from "framer-motion"

import { Button } from "@/components"
import { GTM_EVENTS, routes, useAnalytics } from "@/lib"

import { HeroShapes } from "./shapes"

export const CallToAction: FC<{
	text: string
	description: string
	button: string
}> = ({ text, description }) => {
	const handleNavigate = useAnalytics(
		GTM_EVENTS.CTA_CLICKED,
		process.env.NEXT_PUBLIC_EARLY_ACCESS === "false" ? routes.app : routes.earlyAccess
	)

	return (
		<motion.div
			initial={{ opacity: 0, transform: "translateY(20px)" }}
			whileInView={{ opacity: 1, transform: "translateY(0px)" }}
			transition={{ duration: 0.2 }}
		>
			<div className="relative flex flex-col justify-center overflow-hidden p-8 lg:min-h-[700px] lg:gap-[30px] lg:px-[80px]">
				<HeroShapes />

				<div className="z-[2] my-[40px] flex flex-col gap-16 text-black">
					<motion.h1
						className="max-w-[520px] text-[52px] font-black leading-tight text-black md:max-w-[720px] md:text-[72px] lg:max-w-[840px] lg:text-[82px] xl:max-w-[1200px] xl:text-[96px]"
						initial={{ y: 20, opacity: 0 }}
						whileInView={{
							y: [0, 20],
							opacity: [0, 1]
						}}
						transition={{ duration: 0.3 }}
					>
						{text}
					</motion.h1>

					<motion.p
						className="max-w-[340px] text-[18px] font-bold text-black/40 md:max-w-[520px] lg:max-w-[620px] lg:text-[24px] xl:max-w-[620px] 2xl:max-w-[720px]"
						initial={{ y: -20, opacity: 0 }}
						whileInView={{
							y: [0, -20],
							opacity: [0, 1]
						}}
						transition={{
							duration: 0.3,
							delay: 0.15
						}}
					>
						{description}
					</motion.p>

					<Button
						variant="none"
						className="bg-black/2 w-max rounded-md border-[1px] border-black/20 px-8 py-3 text-center font-black text-black filter backdrop-blur-xl"
						onClick={() => handleNavigate()}
					>
						Enter App
					</Button>
				</div>
			</div>
		</motion.div>
	)
}
