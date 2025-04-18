import { useSession } from "next-auth/react"
import Image from "next/image"
import { useRouter } from "next/router"
import { FC, HTMLAttributes, PropsWithChildren } from "react"

import { AnimatePresence, motion } from "framer-motion"
import { MotionProps } from "framer-motion"
import { FileWarning, LogIn, PaintBucket, PlugIcon, Rocket, SearchIcon } from "lucide-react"

import { useAtom, useAtomValue, useSetAtom } from "jotai"

import { cn } from "@/lib"
import { api } from "@/server/client"
import { socketModelAtom } from "@/state/authentication"
import { columnByIndexAtom, COLUMNS, isFrameAtom, useColumnActions } from "@/state/columns"

import { Frame } from "../app/frames/base"
import { ConsoleSidebarPane } from "../app/sidebar"
import { Callout } from "../app/utils/callout"
import { Button } from "../shared/buttons/button"
import { Ticket } from "./ticket"

export const colors = ["#F3EF8A", "#8AF3E6", "#EB8AF3", "#9F8AF3", "#F3908A", "#F3B08A", "#8AAEF3", "#92F38A"]

const actions = [
	{ name: "Mint Plug Founders Ticket.", sentence: `Mint [Plug Founders Ticket] in [Color:${colors.join(",")}].` },
	{ name: "Early user flag.", sentence: "Began using [Plug] in [Private Alpha]." }
]

const parseColorOptions = (text: string): string[] | null => {
	const colorMatch = text.match(/Color:([\w,#]+)/)
	if (colorMatch) {
		return colorMatch[1].split(",")
	}
	return null
}

const renderActionText = (text: string, frame: (frame: string) => void, selectedColor: string) => {
	return text.split(/(\[[^\]]*\])/).map((part, index) => {
		if (part.startsWith("[") && part.endsWith("]")) {
			const innerText = part.slice(1, -1)
			const colorOptions = parseColorOptions(innerText)

			if (colorOptions) {
				return (
					<span
						key={index}
						className={cn(
							"mx-1 inline-block cursor-pointer rounded-sm px-2 py-0.5 font-bold",
							selectedColor ? "bg-plug-yellow/60" : "bg-plug-red/60"
						)}
						onClick={() => frame("onboarding-colors")}
					>
						{selectedColor ? selectedColor.replace("#", "") : "color"}
					</span>
				)
			}

			return (
				<span key={index} className="mx-1 inline-block rounded-sm bg-plug-yellow/60 px-2 py-0.5 font-bold">
					{innerText}
				</span>
			)
		}
		return <span key={index}>{part}</span>
	})
}

export const ConsoleOnboarding = () => {
	const router = useRouter()
	const step = parseInt(router.query.step as string) || 0

	const handleStepChange = () => {
		const nextStep = step + 1
		router.replace(
			{
				query: { ...router.query, step: nextStep }
			},
			undefined,
			{ shallow: true }
		)
	}

	return (
		<>
			<ConsoleSidebarPane />
			<div className="relative flex h-full w-full flex-col items-center justify-center gap-8 p-12">
				<div className="flex h-full max-h-[960px] w-[460px] items-center justify-center rounded-lg border-[1px] border-plug-green/10">
					{step < 5 && <ConsoleOnboardingStepOne step={step} handleStep={handleStepChange} />}
				</div>

				<div className="absolute bottom-0 z-[9999] h-12 w-full bg-plug-white" />
			</div>
		</>
	)
}

export const ConsoleOnboardingStepActive: FC<
	PropsWithChildren & {
		frame?: string
		active?: boolean
		shimmer?: boolean
		tooltip?: string
		left?: boolean
		delay?: number
	}
> = ({ children, frame = "", active = true, shimmer = true, tooltip, left = false, delay = 0 }) => {
	return (
		<AnimatePresence>
			<motion.div className="relative">
				{shimmer && active && (
					<motion.div
						className={cn(
							"pointer-events-none absolute h-full w-full rounded-lg border-[1px] border-plug-yellow bg-gradient-to-r from-plug-yellow/40 via-white/20 to-plug-yellow/40"
						)}
						animate={{
							backgroundPosition: ["0% 50%", "100% 50%", "0% 50%"]
						}}
						transition={{
							duration: 3,
							ease: "linear",
							repeat: Infinity
						}}
						style={{
							backgroundSize: "200% 100%"
						}}
					/>
				)}

				{!frame && active && tooltip && (
					<motion.div initial={{ opacity: 0 }} animate={{ opacity: 1 }} exit={{ opacity: 0 }}>
						<motion.div
							className={cn(
								"absolute top-1/2 z-50 border-t-[2px] border-dashed border-plug-yellow",
								left ? "right-full" : "left-full"
							)}
							initial={{ width: 0 }}
							animate={{ width: "220px" }}
							transition={{ duration: 0.25, delay: delay + 0.5 }}
						/>

						<div className={cn("absolute bottom-1/2 mb-4 w-[140px]", left ? "-left-1/2" : "-right-1/2")}>
							<motion.p
								initial={{ opacity: 0, y: 10 }}
								animate={{ opacity: 1, y: 0 }}
								exit={{ opacity: 0, y: 10 }}
								className={cn("mt-2 text-right font-bold text-black/40", left && "text-left")}
								transition={{ duration: 0.25, delay: delay + 0.75 }}
							>
								{tooltip}
							</motion.p>
						</div>

						<motion.div
							className={cn(
								"absolute top-1/2 z-[999] h-4 w-4 -translate-y-1/2 rounded-full bg-plug-yellow",
								left ? "-left-2" : "-right-2"
							)}
							initial={{ opacity: 0 }}
							animate={{ opacity: 1 }}
							transition={{ duration: 0.25, delay: delay + 0.25 }}
						/>
					</motion.div>
				)}

				{children}
			</motion.div>
		</AnimatePresence>
	)
}

export const ConsoleOnboardingStepOne: FC<
	MotionProps & HTMLAttributes<HTMLDivElement> & { step: number; handleStep: () => void }
> = ({ step, handleStep }) => {
	const { data: session } = useSession()

	const [column] = useAtom(columnByIndexAtom(COLUMNS.MOBILE_INDEX))
	const frameKey = "onboarding-actions"
	const isFrame = useAtomValue(isFrameAtom)(column, frameKey)
	const { frame } = useColumnActions(COLUMNS.MOBILE_INDEX, frameKey)

	const router = useRouter()
	const color = (router.query.color as string) || ""

	const setColor = (newColor: string) => {
		router.replace(
			{
				query: { ...router.query, color: newColor }
			},
			undefined,
			{ shallow: true }
		)
	}

	const setSocket = useSetAtom(socketModelAtom)

	const onboard = api.socket.onboard.onboard.useMutation({
		onSuccess: socket => setSocket(socket)
	})

	return (
		<div className="flex h-full w-full flex-col">
			<div className="relative flex flex-row items-center gap-4 overflow-hidden border-b-[1px] border-plug-green/10 px-4 py-6">
				<div
					className="absolute bottom-0 left-0 h-12 w-12 rounded-sm blur-2xl filter"
					style={{ backgroundColor: color }}
				/>
				<div className="relative h-6 w-6 rounded-sm bg-plug-blue" style={{ backgroundColor: color }} />
				<p className="relative text-xl font-bold">Mint Plug Founding Ticket</p>
			</div>

			<div className="relative flex h-full w-full flex-col gap-4">
				<div className="z-2 relative flex h-full flex-col gap-2 bg-plug-white p-4">
					{step === 0 && (
						<Callout
							title="No actions added yet."
							description="Let's go ahead and add one by clicking the search bar below."
						/>
					)}

					{step !== 0 && (
						<div className="mb-auto flex h-full flex-col gap-2">
							{color && <Ticket color={color} />}

							{Array.from({ length: Math.min(step, actions.length) }).map((_, actionIndex) => {
								const action = actions[actions.length - 1 - actionIndex]
								return (
									<ConsoleOnboardingStepActive
										key={actionIndex}
										active={actionIndex === step - 1}
										tooltip={
											step === 1
												? "We automatically filled in the values for you."
												: step === 2 && !color
													? "Choose a color by clicking the red input."
													: ""
										}
										frame={column?.frame ?? ""}
										delay={1}
										shimmer={false}
										left
									>
										<motion.div
											className="relative flex flex-row items-center gap-4 overflow-hidden rounded-lg border-[1px] border-plug-green/10 p-4"
											initial={{ opacity: 0, y: 10 }}
											animate={{ opacity: 1, y: 0 }}
											transition={{ duration: 0.2 }}
										>
											<Image
												className="absolute bottom-0 left-0 h-12 w-12 rounded-sm blur-2xl"
												src="/protocols/plug.png"
												alt="Plug"
												width={24}
												height={24}
											/>
											<Image
												className="relative h-6 w-6 rounded-sm"
												src="/protocols/plug.png"
												alt="Plug"
												width={24}
												height={24}
											/>
											<p className="font-bold">
												{renderActionText(action.sentence, frame, color)}
											</p>
										</motion.div>
									</ConsoleOnboardingStepActive>
								)
							})}
						</div>
					)}

					{step < 2 && (
						<ConsoleOnboardingStepActive
							tooltip={!isFrame ? `Click and add '${actions[actions.length - 1 - step].name}' ` : ""}
						>
							<motion.button
								className="mt-auto flex w-full flex-row items-center gap-2 rounded-lg border-[1px] border-plug-green/10 p-4 font-bold text-black/40"
								onClick={() => frame("onboarding-actions")}
								initial={{ opacity: 0 }}
								animate={{ opacity: 1 }}
								transition={{ duration: 0.2, delay: 0.2 }}
							>
								<SearchIcon size={16} className="opacity-60" />
								Search for an action...
							</motion.button>
						</ConsoleOnboardingStepActive>
					)}

					<ConsoleOnboardingStepActive
						active={step === 2}
						tooltip={
							column?.frame
								? ""
								: !color
									? "Please select a 'color' for your ticket to proceed."
									: session?.user.id.startsWith("0x")
										? "Click here to run your Plug and mint your Ticket!"
										: color
											? "Please log in to run your Plug."
											: ""
						}
					>
						<Button
							className="flex w-full flex-row items-center justify-center gap-2 py-4"
							variant={
								!color || !session?.user.id.startsWith("0x") || step < 2 ? "primaryDisabled" : "primary"
							}
							onClick={
								!color || !session?.user.id.startsWith("0x") || step < 2
									? () => {}
									: () => onboard.mutate({ onboardingColor: color })
							}
						>
							{step === 2 && !color ? (
								<FileWarning size={16} className="opacity-60" />
							) : step < 2 || session?.user.id.startsWith("0x") ? (
								<Rocket size={16} className="opacity-60" />
							) : (
								<LogIn size={16} className="opacity-60" />
							)}
							{step === 2 && !color
								? "Missing required inputs"
								: step < 2 || session?.user.id.startsWith("0x")
									? "Run Plug"
									: "Log in"}
						</Button>
					</ConsoleOnboardingStepActive>
				</div>

				<Frame
					index={COLUMNS.MOBILE_INDEX}
					className="overflow-visible"
					label="Founding Ticket: Color"
					visible={column?.frame === "onboarding-colors"}
					icon={<PaintBucket size={16} className="opacity-40" />}
					hasOverlay
				>
					<ConsoleOnboardingStepActive shimmer={false} tooltip="Choose a color and click 'Done'.">
						<div className="grid h-32 grid-cols-4 gap-2">
							{colors.map((color, index) => (
								<button
									key={index}
									className="h-full rounded-lg border-[1px] border-plug-white/40 text-sm font-bold text-black/60 transition-all duration-200 hover:text-black/80"
									style={{ backgroundColor: color }}
									onClick={() => setColor(color)}
								>
									{color}
								</button>
							))}
						</div>
					</ConsoleOnboardingStepActive>

					<Button
						className="mt-2 w-full py-4"
						variant={color ? "primary" : "primaryDisabled"}
						onClick={color ? () => frame("onboarding-colors") : () => {}}
						disabled={!color}
					>
						{color ? "Done" : "Select Color"}
					</Button>
				</Frame>

				<Frame
					index={COLUMNS.MOBILE_INDEX}
					className="overflow-visible"
					label="Onboarding Actions"
					visible={isFrame}
					icon={<PlugIcon size={16} className="opacity-40" />}
					hasOverlay
				>
					<div className="flex w-full flex-col gap-2">
						{Array.from({ length: actions.length - step }).map((_, actionIndex) => {
							return (
								<ConsoleOnboardingStepActive
									active={actionIndex === 0}
									key={actionIndex}
									tooltip="Click on the action to add it to your Plug."
								>
									<button
										className="relative flex w-full flex-row gap-4 overflow-hidden rounded-lg border-[1px] border-plug-green/10 p-4"
										onClick={
											actionIndex === 0
												? () => {
														handleStep()
														frame()
													}
												: undefined
										}
									>
										<div className="flex flex-row items-center gap-2">
											<Image
												className="absolute bottom-0 left-0 h-12 w-12 rounded-sm blur-2xl"
												src="/protocols/plug.png"
												alt="Plug"
												width={24}
												height={24}
											/>
											<Image
												className="h-6 w-6 rounded-sm"
												src="/protocols/plug.png"
												alt="Plug"
												width={24}
												height={24}
											/>
										</div>
										<p className="font-bold">
											{renderActionText(
												actions[actions.length - 1 - step - actionIndex].name,
												setColor,
												color
											)}
										</p>
									</button>
								</ConsoleOnboardingStepActive>
							)
						})}
					</div>
				</Frame>
			</div>
		</div>
	)
}
