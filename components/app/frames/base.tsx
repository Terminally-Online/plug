import { FC, PropsWithChildren } from "react"

import { AnimatePresence, motion } from "framer-motion"
import { ChevronLeft, X } from "lucide-react"

import { Button, Header } from "@/components"
import { useFrame, useSockets } from "@/contexts"
import { cn, useMediaQuery } from "@/lib"

type Props = React.HTMLAttributes<HTMLDivElement> &
	PropsWithChildren & {
		id?: string
		label: string
		visible: boolean
		icon?: JSX.Element
		handleBack?: () => void
		hasOverlay?: boolean
		hasChildrenPadding?: boolean
		next?: JSX.Element
	}

export const Frame: FC<Props> = ({
	id,
	label,
	visible,
	icon,
	handleBack,
	hasOverlay = false,
	hasChildrenPadding = true,
	children,
	className,
	next
}) => {
	const { md } = useMediaQuery()
	const { socket } = useSockets()
	const { handleFrame } = useFrame({ id })

	const page = socket?.columns.find(column => column.id === id)

	if (page === undefined) return null

	return (
		<AnimatePresence>
			{visible ? (
				<div>
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
							md ? "absolute" : "fixed",
							"bottom-0 left-0 right-0 top-0 z-[10] cursor-pointer",
							(handleBack === undefined || hasOverlay === true) &&
								"bg-gradient-to-b from-black/10 to-black/30",
							page.index !== -1 && "rounded-lg"
						)}
						onClick={() => handleFrame()}
					/>

					<motion.div
						initial={{ y: "100%" }}
						animate={{ y: 0 }}
						exit={{ y: "100%" }}
						transition={{ duration: 0.2, ease: "easeInOut" }}
						className={cn(
							md ? "absolute" : "fixed",
							"bottom-0 left-0 w-full rounded-t-lg bg-white",
							className,
							"z-[11]",
							page.index !== -1 && "rounded-b-lg"
						)}
					>
						<div className="flex flex-row items-center gap-2 px-6 py-4">
							{handleBack && (
								<Button variant="secondary" onClick={handleBack} className="mr-2 h-min rounded-sm p-1">
									<ChevronLeft size={14} />
								</Button>
							)}

							<Header
								variant="frame"
								size="md"
								icon={icon}
								label={label}
								nextPadded={false}
								nextOnClick={() => handleFrame()}
								nextLabel={next ?? <X size={14} />}
								nextEmpty={next !== undefined}
							/>
						</div>

						<div className={cn(hasChildrenPadding && "px-6 py-8")}>{children}</div>
					</motion.div>
				</div>
			) : null}
		</AnimatePresence>
	)
}
