import { useState } from "react"

import { cn } from "@/lib"

export const SocketPositionItem = () => {
	const [expanded, setExpanded] = useState(false)

	return (
		<button
			className={cn(
				"group group flex h-min w-full cursor-pointer flex-col items-center overflow-hidden rounded-[16px] border-[1px] border-grayscale-0 p-4 transition-all duration-200 ease-in-out",
				expanded
					? "bg-grayscale-0 hover:bg-white"
					: "bg-white hover:border-white hover:bg-grayscale-0"
			)}
			onClick={() => setExpanded(!expanded)}
		></button>
	)
}
