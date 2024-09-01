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

				<div className="z-[2] flex flex-col gap-16">
					<motion.h1
						className="max-w-[60%] text-[3.5rem] font-black leading-tight text-white md:text-[72px] lg:text-[96px]"
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
						className="max-w-[52%] text-[1.25rem] font-bold text-white/80 md:text-[24px]"
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
						variant="white"
						href={EARLY_ACCESS ? routes.earlyAccess : routes.app}
						className="mt-[30px] w-max"
					>
						{button}
					</Button>
				</div>
			</div>
		</motion.div>
	)
}
