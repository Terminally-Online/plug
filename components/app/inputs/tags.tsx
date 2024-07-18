import { createRef, FC, useEffect, useRef } from "react"

import { motion } from "framer-motion"

import { Button } from "@/components"
import { cn, tags } from "@/lib"

export const Tags: FC<{
	tag?: string
	handleTag: (tag: string) => void
}> = ({ tag, handleTag }) => {
	const scrollContainerRef = useRef<HTMLDivElement>(null)
	const tagRefs = useRef(tags.map(() => createRef<HTMLDivElement>()))

	useEffect(() => {
		if (scrollContainerRef.current && tagRefs.current) {
			const activeIndex = tags.findIndex(
				tagItem => tagItem.toLowerCase() === (tag || "all")
			)

			if (tagRefs.current[activeIndex].current) {
				const activeTag = tagRefs.current[activeIndex].current!
				const activeTagOffsetLeft = activeTag.offsetLeft
				const remInPixels = 16

				scrollContainerRef.current.scrollTo({
					left: activeTagOffsetLeft - remInPixels,
					behavior: "smooth"
				})
			}
		}
	}, [tag])

	return (
		<div className="relative mb-4 mt-2">
			<div
				className="scrollbar-hide overflow-x-auto"
				ref={scrollContainerRef}
			>
				<motion.div
					className="flex flex-row gap-2"
					initial={false}
					transition={{ type: "spring", stiffness: 300, damping: 30 }}
				>
					{tags.map((tagItem, index) => {
						const tagFormatted = tagItem.toLowerCase()
						const tagActive =
							(tag === "" && tagFormatted === "all") ||
							tag === tagFormatted

						return (
							<>
								<motion.div
									key={tagFormatted}
									layout
									transition={{
										type: "spring",
										stiffness: 300,
										damping: 30
									}}
									ref={tagRefs.current[index]}
								>
									<Button
										variant="secondary"
										sizing="sm"
										className={cn(
											"w-max rounded-sm",
											tagActive ? "active" : "",
											index === 0 && "ml-4",
											index === tags.length - 1 &&
												"mr-[60vw]"
										)}
										onClick={() =>
											handleTag(
												tagFormatted === "all" ||
													tagFormatted === tag
													? ""
													: tagFormatted
											)
										}
									>
										{tagItem}
									</Button>
								</motion.div>
							</>
						)
					})}
				</motion.div>
			</div>

			<div className="absolute bottom-0 right-0 top-0 w-24 bg-gradient-to-r from-white/0 to-white" />
		</div>
	)
}
