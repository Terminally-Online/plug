import { FC } from "react"

import { tags } from "@/lib/constants"
import { cn } from "@/lib/utils"

import { Button } from "../buttons"

export const Tags: FC<{
	tag?: string
	handleTag: (tag: string) => void
}> = ({ tag, handleTag }) => {
	return (
		<div className="relative mb-4 mt-2">
			<div className="scrollbar-hide flex flex-row gap-2 overflow-x-auto whitespace-nowrap">
				{tags.map((tagItem, index) => {
					const tagFormatted = tagItem.toLowerCase()

					const tagActive =
						(tag === "" && tagFormatted === "all") ||
						tag === tagFormatted

					return (
						<Button
							key={tagFormatted}
							variant="secondary"
							sizing="sm"
							className={cn(
								tagActive === true ? "bg-grayscale-100" : "",
								index === 0 && "ml-4",
								index === tags.length - 1 && "mr-24"
							)}
							onClick={() =>
								handleTag(
									tagFormatted === "all" ? "" : tagFormatted
								)
							}
						>
							{tagItem}
						</Button>
					)
				})}
			</div>

			<div className="absolute bottom-0 right-0 top-0 w-24 bg-gradient-to-r from-white/0 to-white" />
		</div>
	)
}
