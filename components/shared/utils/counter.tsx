import { FC, useEffect, useRef, useState } from "react"

import { motion } from "framer-motion"

import { cn } from "@/lib/utils"

const isDate = (input: string) => input.includes("/")

const formatForDisplay = (input: number | string, decimals = 2) => {
	if (typeof input === "number") {
		const absCount = Math.abs(input)
		const formattedNumber = absCount
			.toLocaleString("en-US", {
				minimumFractionDigits: decimals,
				maximumFractionDigits: decimals
			})
			.split("")
		if (input < 0) formattedNumber.unshift("-")
		return formattedNumber.reverse()
	} else if (typeof input === "string" && isDate(input)) {
		return input.split("").reverse()
	}
	return []
}

const MinusColumn = () => (
	<div>
		<span>-</span>
	</div>
)

const CommaColumn = () => (
	<div>
		<span>,</span>
	</div>
)

const DecimalColumn = () => (
	<div>
		<span>.</span>
	</div>
)

const SlashColumn = () => (
	<div>
		<span>/</span>
	</div>
)

const NumberColumn: FC<{ digit: string }> = ({ digit }) => {
	const [y, setY] = useState(0)
	const columnContainer = useRef<HTMLDivElement>(null)

	const setColumnToNumber = (number: string) => {
		if (columnContainer.current === null) return

		setY(columnContainer.current.clientHeight * parseInt(number))
	}

	useEffect(() => setColumnToNumber(digit), [digit])

	return (
		<div className="relative text-center" ref={columnContainer}>
			<motion.div className="absolute bottom-0 h-[1000%]" animate={{ y }}>
				{[9, 8, 7, 6, 5, 4, 3, 2, 1, 0].map(i => (
					<div key={i} className="h-[10%]">
						<span>{i}</span>
					</div>
				))}
			</motion.div>

			<span className="invisible">0</span>
		</div>
	)
}

export const Counter: FC<
	{
		count: number | string
		decimals?: number
	} & React.HTMLAttributes<HTMLDivElement>
> = ({ count, decimals = 2, className = "", ...props }) => {
	const numArray = formatForDisplay(count, decimals)

	return (
		<div
			className={cn(
				"relative flex w-full flex-row-reverse overflow-hidden",
				className
			)}
			{...props}
		>
			{numArray.map((number, index) =>
				number === "." ? (
					<DecimalColumn key={index} />
				) : number === "-" ? (
					<MinusColumn key={index} />
				) : number === "," ? (
					<CommaColumn key={index} />
				) : number === "/" ? (
					<SlashColumn key={index} />
				) : (
					<NumberColumn key={index} digit={number} />
				)
			)}
		</div>
	)
}
