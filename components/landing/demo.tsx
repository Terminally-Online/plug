import Link from "next/link"

import { motion } from "framer-motion"

import { routes } from "@/lib"

import { Blob } from "./blob"
import { LandingContainer } from "./layout"

export const Demo = () => {
	return (
		<div className="relative z-[1] my-[80px]">
			<LandingContainer className="relative mb-[40px] flex flex-col gap-4">
				<div className="flex flex-row items-center gap-12">
					<motion.h1
						className="min-w-[640px] text-[64px] font-bold leading-tight"
						initial={{ transform: "translateY(-20px)", opacity: 0 }}
						whileInView={{
							transform: ["translateY(-20px)", "translateY(0px)"],
							opacity: [0, 1]
						}}
						transition={{ duration: 0.3 }}
					>
						The home of all your onchain activity.
					</motion.h1>
					<div className="h-[2px] w-full bg-grayscale-100" />
					<Link
						className="whitespace-nowrap font-bold opacity-40 transition-opacity duration-200 ease-in-out hover:opacity-100"
						href={`${routes.documentation}/introduction/integrations`}
						target="_blank"
						rel="noreferrer"
					>
						Experience the Difference
					</Link>
					<div className="h-[2px] w-24 bg-grayscale-100" />
				</div>
				<motion.p
					className="max-w-[540px] text-[18px] font-bold text-black/40"
					initial={{ transform: "translateY(20px)", opacity: 0 }}
					whileInView={{
						transform: ["translateY(20px)", "translateY(0px)"],
						opacity: [0, 1]
					}}
					transition={{
						duration: 0.3,
						delay: 0.15
					}}
				>
					Build and discover strategies, manage your portfolio, and more with an all-in-one experience you can 
					customize to fit your exact needs and wants.

				</motion.p>
			</LandingContainer>

			<Blob left={"55%"} top={"350"} width={"1000"} height={"500"} />

			<div className="z-[120] h-[800px] w-full bg-grayscale-0" />
		</div>
	)
}
