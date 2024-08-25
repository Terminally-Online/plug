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
		<LandingContainer className="relative flex-col">
			<motion.div initial={{ opacity: 0, y: 20 }} whileInView={{ opacity: 1, y: 0 }} transition={{ duration: 0.2 }}>
				<div className="relative flex flex-col justify-center overflow-hidden rounded-lg p-8 lg:min-h-[700px] lg:gap-[30px] lg:px-[80px]">
					<HeroShapes />

					<div className="z-[2] flex flex-col gap-4">
						<h1 className="text-[36px] font-black text-white lg:max-w-[90%] lg:text-[72px] 2xl:max-w-[50%]">{text}</h1>
						<p className="text-[18px] font-bold text-white/80 md:max-w-[75%] lg:max-w-[60%] lg:text-[24px]">{description}</p>
						<Button variant="white" href={EARLY_ACCESS ? routes.earlyAccess : routes.app} className="mt-[30px] w-max">
							{button}
						</Button>
					</div>
				</div>
			</motion.div>
		</LandingContainer>
	)
}
