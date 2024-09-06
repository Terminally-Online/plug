import { FC } from "react"

import { motion } from "framer-motion"

import { Button, LandingContainer } from "@/components"
import { routes } from "@/lib"

import { HeroShapes } from "./shapes"

const EARLY_ACCESS = process.env.NEXT_PUBLIC_EARLY_ACCESS === "false" ? false : true

export const CallToAction: FC<{
	text: string
	description: string
	button: string
}> = ({ text, description, button }) => {
	return (
		<motion.div
			initial={{ opacity: 0, transform: "translateY(20px)" }}
			whileInView={{ opacity: 1, transform: "translateY(0px)" }}
			transition={{ duration: 0.2 }}
		>
			<div className="relative flex flex-col justify-center overflow-hidden p-8 lg:min-h-[700px] lg:gap-[30px] lg:px-[80px]">
				<HeroShapes />

				<div className="z-[2] my-[40px] flex flex-col gap-16 text-white">
					<motion.h1
						className="max-w-[520px] text-[52px] font-black leading-tight md:max-w-[720px] md:text-[72px] lg:max-w-[840px] lg:text-[82px] xl:max-w-[980px] xl:text-[96px]"
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
						className="max-w-[340px] text-[18px] font-bold text-white/80 md:max-w-[520px] lg:max-w-[620px] lg:text-[24px] xl:max-w-[620px] 2xl:max-w-[720px]"
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
						className="w-max rounded-md border-[1px] border-white/30 bg-white/20 px-8 py-3 text-center font-black text-white filter backdrop-blur-sm"
						href={EARLY_ACCESS ? routes.earlyAccess : routes.app}
					>
						Enter App
					</Button>
				</div>
			</div>
		</motion.div>
	)
}
