"use client"

import { FC, useEffect, useRef, useState } from "react"

import { motion } from "framer-motion"

import { cn } from "@/lib/utils"

const isDate = (input: string) => input.includes("/")

const formatForDisplay = (
	input: number | string,
	formatDecimals: boolean,
	targetDecimals?: number
) => {
	if (typeof input === "number") {
		const absCount = Math.abs(input)
		let formattedNumber: string[]

		if (formatDecimals) {
			if (Number.isInteger(absCount) && targetDecimals === undefined) {
				// For whole numbers without specified decimals, don't add decimal places
				formattedNumber = absCount.toLocaleString("en-US").split("")
			} else {
				// Convert to string to preserve all decimal places
				let stringNumber = absCount.toString()

				let decimalPlacesToKeep: number

				if (targetDecimals !== undefined) {
					// Use the manually specified number of decimal places
					decimalPlacesToKeep = targetDecimals
				} else {
					// Find the index of the first non-zero digit after decimal point
					const decimalIndex = stringNumber.indexOf(".")
					let significantDigitIndex = decimalIndex + 1
					while (stringNumber[significantDigitIndex] === "0") {
						significantDigitIndex++
					}

					// Calculate how many decimal places to keep
					decimalPlacesToKeep = Math.max(
						3,
						significantDigitIndex - decimalIndex + 2
					)
				}

				// Round to the calculated number of decimal places
				const rounded = Number(absCount.toFixed(decimalPlacesToKeep))

				// Format the number, ensuring at least 2 decimal places or the targetDecimals
				let formatted = rounded.toFixed(
					Math.max(2, decimalPlacesToKeep)
				)

				// Trim trailing zeros, but keep at least 2 decimal places or the targetDecimals
				if (targetDecimals === undefined) {
					formatted = formatted.replace(/\.?0+$/, "")
					if (
						formatted.includes(".") &&
						formatted.split(".")[1].length < 2
					) {
						formatted += "0"
					}
				}

				formattedNumber = formatted.split("")
			}
		} else {
			formattedNumber = absCount
				.toLocaleString("en-US", {
					maximumFractionDigits: 20
				})
				.split("")
		}

		if (input < 0) formattedNumber.unshift("-")
		return formattedNumber.reverse()
	} else if (typeof input === "string" && isDate(input)) {
		return input.split("").reverse()
	}
	return []
}

const MinusColumn = () => (
	<span>
		<span>-</span>
	</span>
)

const CommaColumn = () => (
	<span>
		<span>,</span>
	</span>
)

const DecimalColumn = () => (
	<span>
		<span>.</span>
	</span>
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
		formatDecimals?: boolean
		targetDecimals?: number
	} & React.HTMLAttributes<HTMLDivElement>
> = ({
	count,
	formatDecimals = true,
	targetDecimals,
	className = "",
	...props
}) => {
	const numArray = formatForDisplay(count, formatDecimals, targetDecimals)

	return (
		<span
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
		</span>
	)
}
