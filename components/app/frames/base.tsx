import { FC, PropsWithChildren } from "react"

import { AnimatePresence, motion } from "framer-motion"
import { ChevronLeft, X } from "lucide-react"

import { Button, Header } from "@/components"
import { cn, useMediaQuery } from "@/lib"
import { useColumns } from "@/state"

type Props = React.HTMLAttributes<HTMLDivElement> &
	PropsWithChildren & {
		index?: number
		label: string
		visible: boolean
		icon?: JSX.Element
		handleBack?: () => void
		hasOverlay?: boolean
		hasChildrenPadding?: boolean
		next?: JSX.Element
	}

export const Frame: FC<Props> = ({
	index,
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
	const { frame } = useColumns(index)

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
								"bg-gradient-to-b from-black/10 to-black/30"
						)}
						onClick={() => frame()}
					/>

					<motion.div
						initial={{ y: "100%" }}
						animate={{ y: 0 }}
						exit={{ y: "100%" }}
						transition={{ duration: 0.2, ease: "easeInOut" }}
						className={cn(
							md ? "absolute" : "fixed",
							"inset-0 top-auto max-h-[100%] w-full overflow-y-auto overflow-x-hidden rounded-t-lg bg-white",
							className,
							"z-[41]",
							index !== -1 && "rounded-b-lg"
						)}
					>
						<div className="sticky top-0 z-[31] mb-4 flex flex-row items-center gap-2 overflow-hidden border-b-[1px] border-grayscale-100 px-6 py-4 bg-white">
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
								nextOnClick={() => frame()}
								nextLabel={next ?? <X size={14} className="opacity-60 hover:opacity-100" />}
								nextEmpty={next !== undefined}
							/>
						</div>

						<div className={cn(hasChildrenPadding && "overflow-hidden px-6 py-8")}>{children}</div>
					</motion.div>
				</div>
			) : null}
		</AnimatePresence>
	)
}
