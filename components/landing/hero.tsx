import { FC, useEffect, useMemo, useState } from "react"

import Image from "next/image"
import Link from "next/link"

import { AnimatePresence, motion } from "framer-motion"
import { Activity, Book, Twitter } from "lucide-react"

import { Button, Ecosystem, LandingContainer } from "@/components"
import { cn, greenGradientStyle, routes } from "@/lib"

const EARLY_ACCESS =
	process.env.NEXT_PUBLIC_EARLY_ACCESS === "false" ? false : true

const HeroShapes = () => {
	const colors = [
		"#00E100",
		"#A3F700",
		"#00E100",
		"#A3F700",
		"#00E100",
		"#A3F700",
		"#00E100",
		"#A3F700"
	]

	const initialPositions = [
		{ x: "-10%", y: "-10%" },
		{ x: "110%", y: "-10%" },
		{ x: "-10%", y: "110%" },
		{ x: "110%", y: "110%" },
		{ x: "50%", y: "-10%" },
		{ x: "50%", y: "110%" },
		{ x: "-10%", y: "50%" },
		{ x: "110%", y: "50%" }
	]

	return (
		<div className="absolute inset-0 bottom-0 left-0 right-0 top-0 z-0 w-screen overflow-hidden bg-plug-green">
			<div className="absolute inset-0 blur-[120px] filter">
				{colors.map((color, index) => (
					<motion.div
						key={index}
						className="z-= absolute rounded-full"
						style={{
							background: color,
							width: "130%",
							height: "130%",
							x: initialPositions[index].x,
							y: initialPositions[index].y,
							top: "-65%",
							left: "-65%"
						}}
						animate={{
							x: [
								initialPositions[index].x,
								...[
									"0%",
									"100%",
									"50%",
									initialPositions[index].x
								].filter(
									pos => pos !== initialPositions[index].x
								)
							],
							y: [
								initialPositions[index].y,
								...[
									"0%",
									"100%",
									"50%",
									initialPositions[index].y
								].filter(
									pos => pos !== initialPositions[index].y
								)
							],
							scale: [1, 1.1, 0.9, 1]
						}}
						transition={{
							duration: 50,
							ease: "easeInOut",
							repeat: Infinity,
							delay: index * 6
						}}
					/>
				))}
			</div>
		</div>
	)
}

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

	const slide = slides[slideIndex]

	useEffect(() => {
		const interval = setInterval(() => {
			setSlideIndex(prevIndex =>
				prevIndex === slides.length - 1 ? 0 : prevIndex + 1
			)
		}, 5000)

		return () => clearInterval(interval)
	}, [slides])

	return (
		<div className="relative flex h-full min-h-screen">
			<HeroShapes />

			<div className="z-2 relative w-full">
				<LandingContainer className="flex h-full flex-col py-8 text-white">
					<div className="flex flex-row items-center gap-4">
						<Link href={routes.index}>
							<Image
								src="/white-icon.svg"
								alt="Logo"
								width={24}
								height={24}
							/>
						</Link>

						<a
							href={routes.documentation}
							target="_blank"
							rel="noreferrer"
							className="ml-auto"
						>
							<Book
								size={18}
								className="opacity-80 transition-opacity duration-200 hover:opacity-100"
							/>
						</a>

						<a
							href={routes.twitter}
							target="_blank"
							rel="noreferrer"
						>
							<Twitter
								size={18}
								className="opacity-80 transition-opacity duration-200 hover:opacity-100"
							/>
						</a>
					</div>

					<div className="my-auto flex min-h-[640px] items-center pb-6">
						<AnimatePresence>
							<div
								key={slideIndex}
								className="my-auto flex flex-col gap-16"
							>
								<motion.h1
									className="text-[72px] font-black text-white md:max-w-[580px] md:text-[72px] lg:max-w-[800px] lg:text-[96px] xl:max-w-[1200px] xl:text-[144px]"
									initial={{ y: 20, opacity: 0 }}
									animate={{ y: [0, 20], opacity: [0, 1] }}
									transition={{ duration: 0.3 }}
								>
									{slide.title}
								</motion.h1>

								<motion.p
									className="max-w-[440px] text-[24px] font-bold text-white/80 md:max-w-[500px] md:text-[24px]"
									initial={{ y: -20, opacity: 0 }}
									animate={{ y: [0, -20], opacity: [0, 1] }}
									transition={{ duration: 0.3, delay: 0.15 }}
								>
									{slide.subtitle}
								</motion.p>
							</div>
						</AnimatePresence>
					</div>

					<div className="flex flex-col gap-8 font-bold">
						<div className="mx-auto flex flex-row gap-4">
							{Array.from({ length: slides.length }).map(
								(_, index) => (
									<button
										key={index}
										className={cn(
											"h-2 w-2 rounded-full transition-all duration-200 ease-in-out hover:bg-white",
											index === slideIndex
												? "bg-white"
												: "bg-white/60"
										)}
										onClick={() => setSlideIndex(index)}
									/>
								)
							)}
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
								href={
									EARLY_ACCESS
										? routes.earlyAccess
										: routes.app
								}
							>
								{EARLY_ACCESS
									? "Get Early Access"
									: "Enter App"}
							</Button>
						</div>
					</div>
				</LandingContainer>
			</div>
		</div>
	)
}
