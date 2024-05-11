import type { FC } from "react"

import Image from "next/image"

import { motion } from "framer-motion"

import { Container } from "@/components/landing/container"

export const Letter: FC = () => (
	<>
		<Container className="my-[90px] flex-col items-center gap-4">
			<motion.h2
				className="text-center text-[28px] font-bold lg:w-[65%] lg:text-[64px]"
				initial={{ opacity: 0, y: 20 }}
				whileInView={{ opacity: 1, y: 0 }}
				transition={{ duration: 0.2 }}
			>
				Letter From The Team
			</motion.h2>
			<motion.p
				className="text-center text-[18px] font-light opacity-40 lg:w-[40%] lg:text-[24px]"
				initial={{ opacity: 0, y: 20 }}
				whileInView={{ opacity: 0.4, y: 0 }}
				transition={{ duration: 0.2, delay: 0.2 }}
			>
				We’re here to empower humans to benefit from the blockchain to
				the maximum extent and provide the ability to log off instead of
				being terminally online.
			</motion.p>

			<motion.div
				className="mt-[40px] grid lg:grid-cols-12"
				initial={{ opacity: 0, y: 20 }}
				whileInView={{ opacity: 1, y: 0 }}
				transition={{ duration: 0.2, delay: 0.4 }}
			>
				<div className="col-span-6 col-start-4 bg-[#D9D9D9]/10 px-8 lg:mx-8 lg:px-24">
					<div className="mb-16 flex flex-row justify-between gap-2">
						{Array.from({ length: 32 }).map((_, index) => (
							<div
								key={index}
								className="h-24 w-[2px] bg-gradient-to-b from-[#00EF35] to-[#93DF00]"
							/>
						))}
					</div>

					<div className="flex flex-col gap-6 opacity-65">
						<p>Dear anon,</p>
						<p>
							Every turn of the market cycle we find ourselves
							spending more time online. Yet, we go to sleep and
							miss the opportunity we had been trying to time for
							weeks.
						</p>
						<p>
							You were not born to watch numbers go up and down.
							You were not born to obsess over the smallest
							details of blockchain transactions. You were not
							born to live like a robot in a world of abundance.
						</p>
						<p>
							Plug is designed to give you your life back. To give
							you the reality we all dream of where your money
							works even when you don’t. To let you unlock the
							fully power of onchain financial primitives and the
							composability between the many options.
						</p>
						<p>
							You can choose to stay in the past and get worse
							execution with unexpected outcomes or you can use
							Plug and always be certain the outcomes will be
							generated when your conditions have been met.
						</p>
						<p>Love,</p>
						<Image
							src="/landing/signature.svg"
							alt="Signature"
							width={280}
							height={120}
						/>
					</div>

					<div className="mt-16 flex flex-row justify-between gap-2">
						{Array.from({ length: 32 }).map((_, index) => (
							<div
								key={index}
								className="h-8 w-[2px] bg-gradient-to-b from-[#00EF35] to-[#93DF00]"
							/>
						))}
					</div>
				</div>
			</motion.div>
		</Container>
	</>
)
