import { motion } from "framer-motion"

import { Accordion } from "@/components"
import { cn } from "@/lib"

export const SocketCollectibleItem = () => {
	const collectible = undefined
	// const loading = collectible === undefined
	const loading = true

	return (
		<motion.div
			variants={{
				hidden: { opacity: 0, y: 10 },
				visible: {
					opacity: 1,
					y: 0,
					transition: {
						type: "spring",
						stiffness: 100,
						damping: 10
					}
				}
			}}
		>
			<Accordion
				loading={loading}
				expanded={false}
				onExpand={() => {}}
				noPaddingChildren={
					<image
						className={cn(
							"min-h-40 w-full border-b-[1px] border-grayscale-100 bg-grayscale-0 ",
							loading
								? "animate-loading bg-gradient-animated bg-[length:200%_200%]"
								: "transition-all duration-200 ease-in-out group-hover:bg-grayscale-200"
						)}
						height="100%"
						width="auto"
					/>
				}
			>
				<div className="mr-auto w-full text-left">
					<p
						className={cn(
							"truncate whitespace-nowrap font-bold",
							loading && "invisible"
						)}
					>
						Token Name
					</p>

					<div className="flex flex-row items-center gap-2">
						<image
							className={cn(
								"h-4 w-4 rounded-full",
								loading
									? "animate-loading bg-gradient-animated bg-[length:200%_200%]"
									: "bg-grayscale-100"
							)}
							height="100%"
							width="auto"
						/>
						<p
							className={cn(
								"text-sm opacity-60",
								loading && "invisible"
							)}
						>
							Collection Name
						</p>
					</div>
				</div>
			</Accordion>
		</motion.div>
	)
}
