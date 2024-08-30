import { FC, useEffect, useMemo, useState } from "react"

import Image from "next/image"
import Link from "next/link"

import { AnimatePresence, motion } from "framer-motion"
import { Book, Twitter } from "lucide-react"

import { Button, LandingContainer } from "@/components"
import { cn, routes } from "@/lib"

import { HeroShapes } from "./shapes"

const EARLY_ACCESS = process.env.NEXT_PUBLIC_EARLY_ACCESS === "false" ? false : true

export const Hero: FC<{ handleExpand: () => void }> = ({ handleExpand }) => {
	const slides = useMemo(
		() => [
			{
				title: "Less Clicks. More Crypto.",
				subtitle:
					"Plug-and-play protocols with constraint driven transactions and watch everything happen like magic."
			},
			{
				title: "Less Worrying. More Certainty.",
				subtitle:
					"Take risky positions and make complex trades without worrying about getting stuck on the way or ending up in a bad situation."
			},
			{
				title: "Less Decisions. More Options.",
				subtitle:
					"Take advantage of strategies in the Plug marketplace. Adopt top strategies or customize them to fit your circumstances."
			}
		],
		[]
	)

	const [slideIndex, setSlideIndex] = useState(0)
	const [slidesChanged, setSlidesChanged] = useState(false)

	useEffect(() => {
		if (slidesChanged) return

		const interval = setInterval(() => {
			setSlideIndex(prevIndex => (prevIndex === slides.length - 1 ? 0 : prevIndex + 1))
		}, 7500)

		return () => clearInterval(interval)
	}, [slides, slidesChanged])

	return (
		<div className="relative flex h-full min-h-screen w-screen">
			<HeroShapes />

			<div className="z-2 relative w-full">
				<LandingContainer className="flex h-full flex-col py-8 text-white">
					<div className="flex flex-row items-center gap-4">
						<Link href={routes.index}>
							<Image src="/white-icon.svg" alt="Logo" width={24} height={24} />
						</Link>

						<a href={routes.documentation} target="_blank" rel="noreferrer" className="ml-auto">
							<Book size={18} className="opacity-80 transition-opacity duration-200 hover:opacity-100" />
						</a>

						<a href={routes.twitter} target="_blank" rel="noreferrer">
							<Twitter
								size={18}
								className="opacity-80 transition-opacity duration-200 hover:opacity-100"
							/>
						</a>
					</div>

					<div className="my-auto flex min-h-[calc(100vh-180px)] items-center pb-6">
						<AnimatePresence>
							{Array.from({
								length: slides.length
							}).map((_, index) => (
								<>
									{index === slideIndex && (
										<div key={index} className="my-auto flex flex-col gap-16">
											<motion.h1
												className="text-[3.5rem] font-black text-white md:max-w-[580px] md:text-[72px] lg:max-w-[800px] lg:text-[96px] xl:max-w-[1200px] xl:text-[144px]"
												initial={{ y: 20, opacity: 0 }}
												animate={{
													y: [0, 20],
													opacity: [0, 1]
												}}
												transition={{ duration: 0.3 }}
											>
												{slides[index].title}
											</motion.h1>

											<motion.p
												className="max-w-[440px] text-[1.25rem] font-bold text-white/80 md:max-w-[500px] md:text-[24px]"
												initial={{ y: -20, opacity: 0 }}
												animate={{
													y: [0, -20],
													opacity: [0, 1]
												}}
												transition={{
													duration: 0.3,
													delay: 0.15
												}}
											>
												{slides[index].subtitle}
											</motion.p>
										</div>
									)}
								</>
							))}
						</AnimatePresence>
					</div>

					<div className="flex flex-col gap-8 font-bold">
						<div className="mx-auto flex flex-row gap-4">
							{Array.from({ length: slides.length }).map((_, index) => (
								<button
									key={index}
									className={cn(
										"h-2 w-2 rounded-full transition-all duration-200 ease-in-out hover:bg-white",
										index === slideIndex ? "bg-white" : "bg-white/60"
									)}
									onClick={() => {
										setSlidesChanged(true)
										setSlideIndex(index)
									}}
								/>
							))}
						</div>

						<div className="flex flex-row gap-2">
							<button
								className="w-max whitespace-nowrap rounded-md border-[1px] border-white bg-gradient-to-tr from-white/10 to-plug-green/0 px-8 py-3 font-black"
								onClick={handleExpand}
							>
								Learn More
							</button>
							<Button
								variant="none"
								className="w-full rounded-md border-[1px] border-white bg-white px-8 py-3 text-center font-black text-black"
								href={EARLY_ACCESS ? routes.earlyAccess : routes.app}
							>
								{EARLY_ACCESS ? "Get Early Access" : "Enter App"}
							</Button>
						</div>
					</div>
				</LandingContainer>
			</div>
		</div>
	)
}
