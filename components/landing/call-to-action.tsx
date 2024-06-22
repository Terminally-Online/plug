import type { FC } from "react"

import Image from "next/image"

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
}> = ({ text, description, button }) => {
	return (
		<Container className="relative flex-col">
			<motion.div
				initial={{ opacity: 0, y: 20 }}
				whileInView={{ opacity: 1, y: 0 }}
				transition={{ duration: 0.2 }}
			>
				<motion.div className="absolute left-0 top-0 z-[0] h-full w-full overflow-hidden rounded-lg">
					{Array.from({ length: 8 }).map((_, i) => {
						const direction =
							Math.random() > 0.5 ? [0, 90] : [90, 0]

						return (
							<motion.div
								key={i}
								className="absolute h-[140%] w-[140%]"
								initial={{
									left: `${Math.random() * ((i * 40) % 160) - 60}%`,
									top: `${Math.random() * ((i * 40) % 160) - 60}%`,
									filter: "blur(60px)"
								}}
							>
								<motion.div
									className="relative origin-center"
									style={{
										backgroundImage:
											"linear-gradient(45deg, #00E100, #A3F700)",
										width: `${160}%`,
										height: `${160}%`
									}}
									animate={{ rotateZ: direction }}
									transition={{
										duration: 20,
										repeat: Infinity,
										repeatDelay: 0,
										repeatType: "reverse",
										ease: "linear"
									}}
								/>
							</motion.div>
						)
					})}
				</motion.div>

				<div className="relative z-[1] flex flex-col justify-center gap-[15px] rounded-lg bg-transparent p-8 lg:min-h-[700px] lg:gap-[30px] lg:px-[80px]">
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
		</Container>
	)
}
