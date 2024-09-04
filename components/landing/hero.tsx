import Image from "next/image"
import Link from "next/link"

import { motion } from "framer-motion"
import { Book, Twitter } from "lucide-react"

import { Button, HeroBarChart, HeroShapes, LandingContainer } from "@/components"
import { routes } from "@/lib"

const EARLY_ACCESS = process.env.NEXT_PUBLIC_EARLY_ACCESS !== "false"

export const Hero = () => {
	return (
		<div className="relative z-[2] flex h-full min-h-screen w-screen">
			<HeroShapes />
			<HeroBarChart />

			<div className="z-2 relative w-full">
				<LandingContainer className="flex h-full flex-col py-8 text-white">
					<div className="flex flex-row items-center gap-4">
						<Link className="mr-4" href={routes.index}>
							<Image src="/white-icon.svg" alt="Logo" width={24} height={24} />
						</Link>
						<a href={routes.documentation} target="_blank" rel="noreferrer">
							<Book size={18} className="opacity-80 transition-opacity duration-200 hover:opacity-100" />
						</a>
						<a href={routes.twitter} target="_blank" rel="noreferrer">
							<Twitter
								size={18}
								className="opacity-80 transition-opacity duration-200 hover:opacity-100"
							/>
						</a>

						<Button
							variant="none"
							className="ml-auto w-max rounded-md border-[1px] border-white/20 bg-white/20 px-4 py-2 text-center text-sm font-black text-white"
							href={EARLY_ACCESS ? routes.earlyAccess : routes.app}
						>
							Enter App
						</Button>
					</div>

					<div className="my-auto flex min-h-[calc(100vh-180px)] items-center pb-6">
						<div className="my-12 flex flex-col gap-8 md:my-auto">
							<motion.h1
								className="text-[36px] font-black leading-tight md:max-w-[95%] md:text-[64px] lg:text-[72px] xl:max-w-[80%] xl:text-[96px] 2xl:max-w-[65%]"
								initial={{ transform: "translateY(-20px)", opacity: 0 }}
								whileInView={{
									transform: ["translateY(-20px)", "translateY(0px)"],
									opacity: [0, 1]
								}}
								transition={{ duration: 0.3 }}
							>
								Supernatural returns with unparalleled control.
							</motion.h1>

							<motion.p
								className="max-w-[75%] text-[1.25rem] font-bold text-white/80 md:text-[24px] xl:max-w-[55%] 2xl:max-w-[45%]"
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
								Discover a new paradigm of human designed, bot executed transactions that kick your onchain activity into high gear. 
								Adopt and create strategies that deliver the exact outcome you seek in seconds. 
							</motion.p>

							<Button
								variant="none"
								className="mt-8 w-max rounded-md border-[1px] border-white/30 bg-white/20 px-8 py-3 text-center font-black text-white"
								href={EARLY_ACCESS ? routes.earlyAccess : routes.app}
							>
								Enter App
							</Button>
						</div>
					</div>
				</LandingContainer>
			</div>
		</div>
	)
}
