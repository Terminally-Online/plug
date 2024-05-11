import type { FC } from "react"

import { motion } from "framer-motion"

import { Button } from "@/components/buttons"
import { Container } from "@/components/landing/container"
import { routes } from "@/lib/constants"

const EARLY_ACCESS =
	process.env.NEXT_PUBLIC_EARLY_ACCESS === "false" ? false : true

export const CallToAction: FC<{
	text: string
	description: string
	button: string
}> = ({ text, description, button }) => (
	<Container className="flex-col">
		<motion.div
			className="flex flex-col justify-center gap-[15px] rounded-lg bg-gradient-to-tr from-[#00EF35] to-[#93DF00] p-8 lg:min-h-[700px] lg:gap-[30px] lg:px-[80px]"
			initial={{ opacity: 0, y: 20 }}
			whileInView={{ opacity: 1, y: 0 }}
			transition={{ duration: 0.2 }}
		>
			<h1 className="text-[36px] font-bold text-white lg:max-w-[80%] lg:text-[72px]">
				{text}
			</h1>
			<p className="text-[18px] text-white lg:max-w-[60%] lg:text-[24px]">
				{description}
			</p>
			<Button
				variant="white"
				href={EARLY_ACCESS ? routes.earlyAccess : routes.app}
				className="mt-[30px] w-full lg:w-max"
			>
				{button}
			</Button>
		</motion.div>
	</Container>
)
