import { useSession } from "next-auth/react"
import Image from "next/image"
import { FC, HTMLAttributes, useCallback, useState } from "react"

import { motion } from "framer-motion"

import { Button } from "@/components/shared"
import { useBeforeInstall } from "@/contexts"
import { cn, GTM_EVENTS, useAnalytics } from "@/lib"
import { Flag, useColumns, useFlags } from "@/state"

export const ColumnApplication: FC<HTMLAttributes<HTMLDivElement> & { index: number }> = ({
	index,
	className,
	...props
}) => {
	const { data: session } = useSession()
	const { prompt, isNativePromptAvailable, instructions } = useBeforeInstall()
	const { handleFlag } = useFlags()
	const { remove } = useColumns()

	const handlePWAInstall = useAnalytics(GTM_EVENTS.CTA_CLICKED, session?.user?.id, true, "/")

	const [currentStep, setCurrentStep] = useState(0)

	const handleClose = useCallback(
		(added: boolean = false) => {
			if (added) handlePWAInstall()

			handleFlag(Flag.SHOW_PWA, false)
			remove(index)
		},
		[index, remove, handleFlag, handlePWAInstall]
	)

	if (instructions.length === 0) return null

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
							onClick={isNativePromptAvailable && prompt ? async () => await prompt?.() : undefined}
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

					{isNativePromptAvailable && prompt ? (
						<>
							<p className="mb-4 max-w-[300px] text-sm opacity-40">
								Create a shortcut to manage your onchain activity in just one click away at all times.
							</p>
							<Button className="truncate" sizing="sm" onClick={async () => await prompt?.()}>
								Install
							</Button>
						</>
					) : (
						<div className="flex max-w-[300px] flex-col gap-4 text-sm">
							<p className="inline-flex flex-wrap items-center gap-x-2 gap-y-1 text-black/40">
								{instructions[currentStep]}
							</p>

							<div className="flex justify-center gap-2">
								{currentStep === instructions.length - 1 ? (
									<Button className="w-max" sizing="sm" onClick={() => handleClose(true)}>
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
										<Button
											className="w-max"
											sizing="sm"
											onClick={() => {
												if (currentStep < instructions.length - 1) {
													setCurrentStep(currentStep + 1)
												}
											}}
										>
											Continue
										</Button>
									</>
								)}
							</div>
							<p
								className="mt-8 cursor-pointer text-sm opacity-40 transition-opacity duration-200 ease-in-out hover:opacity-100"
								onClick={() => handleClose(false)}
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
