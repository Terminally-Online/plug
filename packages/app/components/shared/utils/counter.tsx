import { FC, useEffect, useRef, useState } from "react"

import { motion } from "framer-motion"

import { formatForDisplay, useClient } from "@/lib"
import { cn } from "@/lib/utils"

const NumberColumn: FC<{ digit: string }> = ({ digit }) => {
	const [y, setY] = useState(0)
	const columnContainer = useRef<HTMLSpanElement>(null)

	const setColumnToNumber = (number: string) => {
		if (columnContainer.current === null) return
		setY(columnContainer.current.clientHeight * parseInt(number))
	}

	useEffect(() => setColumnToNumber(digit), [digit])

	return (
		<span className="relative text-center" ref={columnContainer}>
			<motion.span className="absolute bottom-0 h-[1000%]" whileInView={{ translateY: y }}>
				{[9, 8, 7, 6, 5, 4, 3, 2, 1, 0].map(i => (
					<div key={i} className="h-[10%]">
						{i}
					</div>
				))}
			</motion.span>
			<span className="invisible">0</span>
		</span>
	)
}

export const Counter: FC<
	{
		count: number | string
		formatDecimals?: boolean
		decimals?: number
	} & React.HTMLAttributes<HTMLDivElement>
> = ({ count, formatDecimals = true, decimals, className = "", ...props }) => {
	const isClient = useClient()

	const numArray = formatForDisplay(count, formatDecimals, decimals)

	if (!isClient) return null

	return (
		<span className={cn("relative flex w-full flex-row-reverse overflow-hidden", className)} {...props}>
			{count !== undefined && count !== null ? (
				numArray.map((char, index) =>
					char === " " ? (
						<span key={index}>&nbsp;</span>
					) : !isNaN(Number(char)) ? (
						<NumberColumn key={index} digit={char} />
					) : (
						<span key={index}>{char}</span>
					)
				)
			) : (
				"-"
			)}
		</span>
	)
}
