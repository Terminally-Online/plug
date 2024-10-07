import Image from "next/image"
import React, { useState } from "react"
import { isAndroid, isChrome, isIOS, isMacOs, isSafari, isWindows, osName } from "react-device-detect"

import { motion } from "framer-motion"
import { ChevronRight } from "lucide-react"

import { Button } from "@/components/shared"
import { useBeforeInstall } from "@/contexts"
import { cn } from "@/lib"
import { Flag, useFlags } from "@/state"

type Instructions = Record<
	"macOS" | "iOS" | "android" | "linux" | "windows" | "safari" | "chrome",
	(string | JSX.Element)[]
>

const instructions = {
	android: ["Open this page in Chrome"],
	windows: ["Open this page in Edge"],
	linux: ["Open this page in Chromium or Chrome"],
	iOS: [
		"Open this page in Safari",
		<span key="ios-share" className="flex items-center">
			Click the Share button
			<Image
				className="mx-1 inline-block h-4 w-4 opacity-60"
				src="/misc/ios-share-icon.svg"
				alt="apple-share-icon"
				width={48}
				height={48}
			/>
			in the Safari toolbar, then choose Add to home screen
		</span>,
		"Type the name that you want to use for the web app, then click Add."
	],
	macOS: [
		"Open this page in Safari or Chrome. Arc does not support native apps on macOS.",
		<span key="macos-share" className="items-center">
			From the menu bar, choose File <ChevronRight size={14} className="mb-1 inline-block h-4 w-4" /> Add to Dock.
			Or click the Share button
			<Image
				className="mx-1 mb-1 inline-block h-4 w-4 opacity-60"
				src="/misc/ios-share-icon.svg"
				alt="apple-share-icon"
				width={48}
				height={48}
			/>
			in the Safari toolbar, then choose Add to Dock
		</span>,
		"Type the name that you want to use for the web app, then click Add"
	],
	safari: [
		<span key="safari-share" className="items-center">
			Click the Share button
			<Image
				className="mx-1 mb-1 inline-block h-4 w-4 opacity-60"
				src="/misc/ios-share-icon.svg"
				alt="apple-share-icon"
				width={48}
				height={48}
			/>
			in the Safari toolbar, then choose Add to Dock
		</span>,
		"Type the name that you want to use for the web app, then click Add"
	],
	chrome: [
		"Click the three-dot menu in the top right corner to open the options menu.",
		"Select 'More tools' > 'Create shortcut'. to open the shortcut creation menu.",
		"Finish by checking 'Open as window' and clicking 'Create' to install the app."
	]
} satisfies Instructions

function getInstructions() {
	if (isMacOs) {
		if (isSafari) return instructions.safari
		if (isChrome) return instructions.chrome
		// NOTE: If the user is on Arc, the app will not be installed due to lack of support.
		return instructions.macOS
	}
	if (isIOS) return instructions.iOS
	if (isAndroid) return instructions.android
	if (osName === "Linux") return instructions.linux
	if (isWindows) return instructions.windows

	return []
}

export const ColumnApplication: React.FC<React.HTMLAttributes<HTMLDivElement> & { index: number }> = ({
	index,
	className,
	...props
}) => {
	const beforeInstall = useBeforeInstall()
	const { handleFlag } = useFlags()

	const [currentStep, setCurrentStep] = useState(0)

	if (!beforeInstall) return null

	const osInstructions = getInstructions()

	if (osInstructions.length === 0) return null

	const handleNext = () => {
		if (currentStep < osInstructions.length - 1) {
			setCurrentStep(currentStep + 1)
		}
	}

	return (
		<div className={cn("flex h-full items-center justify-center overflow-x-hidden", className)} {...props}>
			<div className="absolute top-0 h-1/4 w-full bg-gradient-to-tr from-plug-green to-plug-yellow blur-[320px] filter" />

			<div className="relative h-full w-full overflow-hidden font-bold">
				<motion.div
					className="absolute left-[-5%] top-[-5%] flex h-[110%] w-[110%] skew-x-[10deg] items-center justify-center mix-blend-overlay"
					initial={{ scale: 1.4 }}
					animate={{ scale: 2.2 }}
					transition={{
						duration: 0.3,
						ease: "easeInOut",
						repeat: Infinity,
						repeatType: "reverse",
						repeatDelay: 3
					}}
				>
					<motion.svg
						className="origin-center -translate-x-1/2 -translate-y-1/2"
						viewBox="0 0 1500 1500"
						initial={{ rotate: 0 }}
						animate={{ rotate: -360 }}
						transition={{ duration: 20, repeat: Infinity, repeatType: "loop", ease: "linear" }}
					>
						<defs>
							<path id="circle" d="M750,750 m-550,0 a550,550 0 0,1 1100,0 a550,550 0 0,1 -1100,0" />
						</defs>
						<text className="text-[232px]">
							<textPath href="#circle" startOffset="0%">
								Stay Plugged In • Save as App •
							</textPath>
						</text>
					</motion.svg>
				</motion.div>

				<div className="absolute inset-0 z-[20] flex flex-col items-center justify-center gap-2 px-4 text-center">
					<h2 className="mb-2 flex h-4 flex-wrap gap-4 text-xl">
						Add
						<motion.span
							className="relative cursor-pointer"
							initial={{ y: -25 }}
							animate={{ y: -5 }}
							transition={{
								y: {
									duration: 0.6,
									repeat: Infinity,
									repeatType: "reverse",
									ease: "easeIn"
								}
							}}
							onClick={
								beforeInstall.isNativePromptAvailable
									? async () => {
											await beforeInstall.prompt()
										}
									: undefined
							}
						>
							<Image
								className="absolute left-0 top-1/2 h-16 w-16 -translate-y-1/2 blur-lg filter"
								src="/plug-logo-green.svg"
								alt="Logo"
								width={64}
								height={64}
							/>
							<Image
								className="relative h-8"
								src="/plug-logo-green.svg"
								alt="Logo"
								width={32}
								height={32}
							/>
						</motion.span>
						to your dock.
					</h2>

					{beforeInstall.isNativePromptAvailable ? (
						<>
							<p className="mb-4 max-w-[300px] text-sm opacity-40">
								Create a shortcut to manage your onchain activity in just one click away at all times.
							</p>
							<Button
								className="truncate"
								sizing="sm"
								onClick={async () => {
									await beforeInstall.prompt()
								}}
							>
								Install
							</Button>
						</>
					) : (
						<div className="flex max-w-[300px] flex-col gap-4 text-sm">
							<span className="inline-flex flex-wrap items-center gap-x-2 gap-y-1 text-black/40">
								{osInstructions[currentStep]}
							</span>
							<div className="flex justify-center gap-2">
								{currentStep === osInstructions.length - 1 ? (
									<Button className="w-max" sizing="sm" onClick={() => setCurrentStep(0)}>
										Done
									</Button>
								) : (
									<>
										{currentStep > 0 && (
											<Button
												variant="secondary"
												className="w-max bg-white"
												sizing="sm"
												onClick={() => setCurrentStep(prev => prev - 1)}
											>
												Back
											</Button>
										)}
										<Button className="w-max" sizing="sm" onClick={handleNext}>
											Next
										</Button>
									</>
								)}
							</div>
							<p
								className="mt-8 cursor-pointer text-sm opacity-40 hover:opacity-100"
								onClick={() => handleFlag(Flag.SHOW_PWA, false)}
							>
								Never show this again.
							</p>
						</div>
					)}
				</div>
			</div>
		</div>
	)
}
