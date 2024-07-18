import { FC } from "react"

import { motion } from "framer-motion"

import { Button, LandingContainer } from "@/components"
import { routes } from "@/lib"

const EARLY_ACCESS =
	process.env.NEXT_PUBLIC_EARLY_ACCESS === "false" ? false : true

export const CallToAction: FC<{
	text: string
	description: string
	button: string
}> = ({ text, description, button }) => {
	return (
		<LandingContainer className="relative flex-col">
			<motion.div
				initial={{ opacity: 0, y: 20 }}
				whileInView={{ opacity: 1, y: 0 }}
				transition={{ duration: 0.2 }}
			>
				<div
					className="relative z-[1] flex flex-col justify-center gap-[15px] rounded-lg bg-transparent p-8 lg:min-h-[700px] lg:gap-[30px] lg:px-[80px]"
					style={{
						backgroundImage:
							"linear-gradient(30deg, #00E100, #A3F700)"
					}}
				>
					<h1 className="text-[36px] font-bold text-white lg:max-w-[90%] lg:text-[72px] 2xl:max-w-[50%]">
						{text}
					</h1>
					<p className="text-[18px] text-white md:max-w-[75%] lg:max-w-[60%] lg:text-[24px]">
						{description}
					</p>
					<Button
						variant="white"
						href={
							EARLY_ACCESS ? routes.earlyAccess : routes.app.index
						}
						className="mt-[30px] w-max"
					>
						{button}
					</Button>
				</div>
			</motion.div>
		</LandingContainer>
	)
}
