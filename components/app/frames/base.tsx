import { FC, PropsWithChildren, useEffect } from "react"

import { AnimatePresence, motion } from "framer-motion"
import { ChevronLeft, X } from "lucide-react"

import { Button, Header } from "@/components"
import { useFrame } from "@/contexts"
import { cn } from "@/lib"

type Props = PropsWithChildren & {
	label: string
	visible: boolean
	icon?: JSX.Element
	handleBack?: () => void
	hasOverlay?: boolean
} & React.HTMLAttributes<HTMLDivElement>

export const Frame: FC<Props> = ({
	children,
	className,
	label,
	visible,
	icon,
	handleBack,
	hasOverlay = false
}) => {
	const { handleFrameVisible } = useFrame()

	useEffect(() => {
		const handleKeyDown = (event: KeyboardEvent) => {
			if (event.key === "Escape") handleFrameVisible(undefined)
		}

		if (visible) document.addEventListener("keydown", handleKeyDown)

		return () => document.removeEventListener("keydown", handleKeyDown)
	}, [visible, handleFrameVisible])

	return (
		<AnimatePresence>
			{visible ? (
				<>
					<motion.div
						initial={{ opacity: 0 }}
						animate={{ opacity: 1 }}
						exit={{ opacity: 0 }}
						transition={{
							duration: 0.2,
							ease: "easeInOut",
							delay: 0.1
						}}
						className={cn(
							"fixed bottom-0 left-0 right-0 top-0 z-[10] cursor-pointer",
							(handleBack === undefined || hasOverlay === true) &&
								"bg-gradient-to-b from-black/10 to-black/30"
						)}
						onClick={() => handleFrameVisible(undefined)}
					/>

					<motion.div
						initial={{ y: "100%" }}
						animate={{ y: 0 }}
						exit={{ y: "100%" }}
						transition={{ duration: 0.2, ease: "easeInOut" }}
						className={cn(
							"fixed bottom-0 left-0 w-full rounded-t-lg bg-white px-6 py-8",
							className,
							"z-[11]"
						)}
					>
						<div className="flex flex-row items-center gap-2">
							{handleBack && (
								<Button
									variant="secondary"
									onClick={handleBack}
									className="mb-[20px] mr-2 h-min p-1"
								>
									<ChevronLeft size={14} />
								</Button>
							)}

							<Header
								variant="frame"
								size="md"
								icon={icon}
								label={label}
								nextPadded={false}
								nextOnClick={() =>
									handleFrameVisible(undefined)
								}
								nextLabel={<X size={14} />}
							/>
						</div>

						{children}
					</motion.div>
				</>
			) : null}
		</AnimatePresence>
	)
}
